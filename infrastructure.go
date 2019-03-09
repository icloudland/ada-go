package ada

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"errors"
	"strings"
	"github.com/icloudland/ada-go/adajson"
	"github.com/dyninc/qstring"
	"reflect"
)

var (
	ErrClientShutdown = errors.New("the client has been shutdown")

	ErrNotEnoughMoney = errors.New("NotEnoughMoney")
)

const (
	sendBufferSize  = 100
	defaultBasePath = "api/v1/"
)

type sendRequestDetails struct {
	httpRequest *http.Request
	jsonRequest *jsonRequest
}

type jsonRequest struct {
	path           string
	requestType    string
	cmd            interface{}
	marshalledJSON []byte
	responseChan   chan *response
}

type Client struct {
	config *ConnConfig

	httpClient *http.Client

	sendRequestChan chan *sendRequestDetails
	shutdown        chan struct{}
	wg              sync.WaitGroup
}

type (
	rawResponse struct {
		Data       json.RawMessage `json:"data"`
		Status     string          `json:"status"`
		Diagnostic json.RawMessage `json:"diagnostic"`
		Message    string          `json:"message"`
		Meta       json.RawMessage `json:"meta"`
	}
)

type response struct {
	result []byte
	meta   []byte
	err    error
}

func (r rawResponse) result() (result, meta []byte, err error) {
	if r.Status == "error" {
		return nil, nil, errors.New(r.Message)
	}
	return r.Data, r.Meta, nil
}

func (c *Client) handleSendRequestMessage(details *sendRequestDetails) {
	jReq := details.jsonRequest
	httpResponse, err := c.httpClient.Do(details.httpRequest)
	if err != nil {
		jReq.responseChan <- &response{err: err}
		return
	}

	// Read the raw bytes and close the response.
	respBytes, err := ioutil.ReadAll(httpResponse.Body)
	httpResponse.Body.Close()
	if err != nil {
		err = fmt.Errorf("error reading json reply: %v", err)
		jReq.responseChan <- &response{err: err}
		return
	}

	// Try to unmarshal the response as a regular JSON-RPC response.
	var resp rawResponse
	err = json.Unmarshal(respBytes, &resp)
	fmt.Println(string(respBytes[:]))
	if err != nil {
		// When the response itself isn't a valid JSON-RPC response
		// return an error which includes the HTTP status code and raw
		// response bytes.
		err = fmt.Errorf("status code: %d, response: %q",
			httpResponse.StatusCode, string(respBytes))
		jReq.responseChan <- &response{err: err}
		return
	}

	res, meta, err := resp.result()
	jReq.responseChan <- &response{result: res, err: err, meta: meta,}
}

func (c *Client) sendPostHandler() {
out:
	for {
		// Send any messages ready for send until the shutdown channel
		// is closed.
		select {
		case details := <-c.sendRequestChan:
			c.handleSendRequestMessage(details)

		case <-c.shutdown:
			break out
		}
	}

	// Drain any wait channels before exiting so nothing is left waiting
	// around to send.
cleanup:
	for {
		select {
		case details := <-c.sendRequestChan:
			details.jsonRequest.responseChan <- &response{
				result: nil,
				err:    ErrClientShutdown,
			}

		default:
			break cleanup
		}
	}
	c.wg.Done()

}

// sendPostRequest sends the passed HTTP request to the RPC server using the
// HTTP client associated with the client.  It is backed by a buffered channel,
// so it will not block until the send channel is full.
func (c *Client) sendPostRequest(httpReq *http.Request, jReq *jsonRequest) {
	// Don't send the message if shutting down.
	select {
	case <-c.shutdown:
		jReq.responseChan <- &response{result: nil, err: ErrClientShutdown}
	default:
	}

	c.sendRequestChan <- &sendRequestDetails{
		jsonRequest: jReq,
		httpRequest: httpReq,
	}
}

// newFutureError returns a new future result channel that already has the
// passed error waitin on the channel with the reply set to nil.  This is useful
// to easily return errors from the various Async functions.
func newFutureError(err error) chan *response {
	responseChan := make(chan *response, 1)
	responseChan <- &response{err: err}
	return responseChan
}

// receiveFuture receives from the passed futureResult channel to extract a
// reply or any errors.  The examined errors include an error in the
// futureResult and the error in the reply from the server.  This will block
// until the result is available on the passed channel.
func receiveFuture(f chan *response) ([]byte, []byte, error) {
	// Wait for a response on the returned channel.
	r := <-f
	fmt.Println(string(r.result[:]))
	return r.result, r.meta, r.err
}

// sendRequest sends the passed json request to the associated server using the
// provided response channel for the reply.  It handles both websocket and HTTP
// POST mode depending on the configuration of the client.
func (c *Client) sendRequest(jReq *jsonRequest) {

	// Generate a request to the configured RPC server.
	protocol := "http"
	if !c.config.DisableTLS {
		protocol = "https"
	}
	url := protocol + "://" + c.config.Host + "/" + c.config.BasePath + jReq.path
	bodyReader := bytes.NewReader(jReq.marshalledJSON)
	fmt.Println(url)
	fmt.Println(string(jReq.marshalledJSON[:]))
	httpReq, err := http.NewRequest(jReq.requestType, url, bodyReader)
	if err != nil {
		jReq.responseChan <- &response{result: nil, err: err}
		return
	}
	httpReq.Close = true
	httpReq.Header.Set("Content-Type", "application/json; charset=utf-8")
	httpReq.Header.Set("Accept", "application/json; charset=utf-8")

	// Configure basic access authorization.
	httpReq.SetBasicAuth(c.config.User, c.config.Pass)

	c.sendPostRequest(httpReq, jReq)

}

// sendCmd sends the passed command to the associated server and returns a
// response channel on which the reply will be delivered at some point in the
// future.  It handles both websocket and HTTP POST mode depending on the
// configuration of the client.
func (c *Client) sendCmd(cmd interface{}) chan *response {
	// Get the method associated with the command.
	method, err := adajson.CmdMethod(cmd)
	if err != nil {
		return newFutureError(err)
	}

	methods := strings.Split(method, ":")
	path := methods[0]
	requestType := strings.ToUpper(methods[1])

	path = setUrlPath(cmd, path)
	if requestType == "GET" {
		queryString, err := qstring.MarshalString(cmd)
		if err != nil {
			return newFutureError(err)
		}
		if queryString != "" {
			path = path + "?" + queryString
		}
	}

	// Marshal the command.
	marshalledJSON, err := adajson.MarshalCmdSimple(cmd)
	if err != nil {
		return newFutureError(err)
	}

	// Generate the request and send it along with a channel to respond on.
	responseChan := make(chan *response, 1)
	jReq := &jsonRequest{
		path:           path,
		requestType:    requestType,
		cmd:            cmd,
		marshalledJSON: marshalledJSON,
		responseChan:   responseChan,
	}
	c.sendRequest(jReq)

	return responseChan
}

func setUrlPath(cmd interface{}, path string) string {
	nPath := path
	t := reflect.TypeOf(cmd)
	v := reflect.ValueOf(cmd)
	for i := 0; i < t.Elem().NumField(); i++ {
		tf := t.Elem().Field(i)
		pathTag := tf.Tag.Get("path")
		if pathTag != "" {
			vf := v.Elem().FieldByName(tf.Name).String()
			nPath = strings.Replace(nPath, "{{"+pathTag+"}}", vf, 1)
		}
	}
	return nPath
}

// doShutdown closes the shutdown channel and logs the shutdown unless shutdown
// is already in progress.  It will return false if the shutdown is not needed.
//
// This function is safe for concurrent access.
func (c *Client) doShutdown() bool {
	// Ignore the shutdown request if the client is already in the process
	// of shutting down or already shutdown.
	select {
	case <-c.shutdown:
		return false
	default:
	}

	close(c.shutdown)
	return true
}

// Shutdown shuts down the client by disconnecting any connections associated
// with the client and, when automatic reconnect is enabled, preventing future
// attempts to reconnect.  It also stops all goroutines.
func (c *Client) Shutdown() {

	// Ignore the shutdown request if the client is already in the process
	// of shutting down or already shutdown.
	if !c.doShutdown() {
		return
	}

}

// start begins processing input and output messages.
func (c *Client) start() {

	c.wg.Add(1)
	go c.sendPostHandler()
}

// WaitForShutdown blocks until the client goroutines are stopped and the
// connection is closed.
func (c *Client) WaitForShutdown() {
	c.wg.Wait()
}

// ConnConfig describes the connection configuration parameters for the client.
// This
type ConnConfig struct {
	// Host is the IP address and port of the RPC server you want to connect
	// to.
	Host string

	// User is the username to use to authenticate to the RPC server.
	User string

	// Pass is the passphrase to use to authenticate to the RPC server.
	Pass string

	// DisableTLS specifies whether transport layer security should be
	// disabled.  It is recommended to always use TLS if the RPC server
	// supports it as otherwise your username and password is sent across
	// the wire in cleartext.
	DisableTLS bool

	// Certificates are the bytes for a PEM-encoded certificate chain used
	// for the TLS connection.  It has no effect if the DisableTLS parameter
	// is true.
	Certificates []byte

	// Proxy specifies to connect through a SOCKS 5 proxy server.  It may
	// be an empty string if a proxy is not required.
	Proxy string

	// ProxyUser is an optional username to use for the proxy server if it
	// requires authentication.  It has no effect if the Proxy parameter
	// is not set.
	ProxyUser string

	// ProxyPass is an optional password to use for the proxy server if it
	// requires authentication.  It has no effect if the Proxy parameter
	// is not set.
	ProxyPass string

	BasePath string
}

// newHTTPClient returns a new http client that is configured according to the
// proxy and TLS settings in the associated connection configuration.
func newHTTPClient(config *ConnConfig) (*http.Client, error) {
	// Set proxy function if there is a proxy configured.
	var proxyFunc func(*http.Request) (*url.URL, error)
	if config.Proxy != "" {
		proxyURL, err := url.Parse(config.Proxy)
		if err != nil {
			return nil, err
		}
		proxyFunc = http.ProxyURL(proxyURL)
	}

	// Configure TLS if needed.
	var tlsConfig *tls.Config
	if !config.DisableTLS {
		if len(config.Certificates) > 0 {
			pool := x509.NewCertPool()
			pool.AppendCertsFromPEM(config.Certificates)
			tlsConfig = &tls.Config{
				RootCAs: pool,
			}
		}
		if tlsConfig == nil {
			tlsConfig = &tls.Config{}
		}
		tlsConfig.InsecureSkipVerify = true
	}

	client := http.Client{
		Transport: &http.Transport{
			Proxy:           proxyFunc,
			TLSClientConfig: tlsConfig,
		},
	}

	return &client, nil
}

func New(config *ConnConfig) (*Client, error) {

	if config.BasePath == "" {
		config.BasePath = defaultBasePath
	}
	var httpClient *http.Client

	var err error
	httpClient, err = newHTTPClient(config)
	if err != nil {
		return nil, err
	}

	client := &Client{
		config:          config,
		httpClient:      httpClient,
		sendRequestChan: make(chan *sendRequestDetails, sendBufferSize),
		shutdown:        make(chan struct{}),
	}

	client.start()

	return client, nil
}
