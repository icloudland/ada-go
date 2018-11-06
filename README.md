Cardano Wallet API library for Go
=========================

This library provides simple access to data structures (binary packing
and JSON interface) and API calls to an Wallet RPC server, running
remotely or locally.  


Basic usage
-----------

```go
package main
import "github.com/icloudland/ada-go"
import "fmt"

func main() {
	connCfg := &ada.ConnConfig{
    		Host:         "192.168.0.1:8090",
    		DisableTLS:   true,
    	}
    client, _ := ada.New(connCfg)
    defer client.Shutdown()
    
    nodeInfo, _ := client.NodeInfo()
    fmt.Println(nodeInfo)
}
```
