package ada

import (
	"github.com/icloudland/ada-go/adajson"
	"encoding/json"
)

type FutureNodeInfo chan *response

func (r FutureNodeInfo) Receive() (*adajson.NodeInfoResult, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var info adajson.NodeInfoResult
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *Client) NodeInfoAsync() FutureNodeInfo {

	cmd := adajson.NewNodeInfoCmd()
	return c.sendCmd(cmd)
}
// Retrieves the dynamic information for this node
func (c *Client) NodeInfo() (*adajson.NodeInfoResult, error) {
	return c.NodeInfoAsync().Receive()
}


type FutureNodeSettings chan *response

func (r FutureNodeSettings) Receive() (*adajson.NodeSettings, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var info adajson.NodeSettings
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *Client) NodeSettingsAsync() FutureNodeSettings {

	cmd := adajson.NewNodeSettingsCmd()
	return c.sendCmd(cmd)
}
// Retrieves the static settings for this node.
func (c *Client) NodeSettings() (*adajson.NodeSettings, error) {
	return c.NodeSettingsAsync().Receive()
}

type FutureCreateWallet chan *response

func (r FutureCreateWallet) Receive() (*adajson.WalletInfo, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var info adajson.WalletInfo
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *Client) CreateWalletAsync(backupPhrase []string,
	assuranceLevel, name, spendingPassword string) FutureCreateWallet {

	cmd := adajson.NewCreateWalletCmd("create", backupPhrase,
		assuranceLevel, name, spendingPassword)

	return c.sendCmd(cmd)
}

// Creates a new  Wallet
func (c *Client) CreateWallet(backupPhrase []string,
	assuranceLevel, name, spendingPassword string) (*adajson.WalletInfo, error) {

	return c.CreateWalletAsync(backupPhrase, assuranceLevel, name, spendingPassword).Receive()
}

func (c *Client) RestoreWalletAsync(backupPhrase []string,
	assuranceLevel, name, spendingPassword string) FutureCreateWallet {

	cmd := adajson.NewCreateWalletCmd("restore", backupPhrase,
		assuranceLevel, name, spendingPassword)

	return c.sendCmd(cmd)
}
// Restores an existing Wallet
func (c *Client) RestoreWallet(backupPhrase []string,
	assuranceLevel, name, spendingPassword string) (*adajson.WalletInfo, error) {

	return c.RestoreWalletAsync(backupPhrase, assuranceLevel, name, spendingPassword).Receive()
}

func (c *Client) UpdatePwdAsync(walletId, op, np string) FutureWalletInfo {

	cmd := adajson.NewUpdatePwdCmd(walletId, op, np)
	return c.sendCmd(cmd)
}
// Updates the password for the given Wallet.
func (c *Client) UpdatePwd(walletId, op, np string) (*adajson.WalletInfo, error) {
	return c.UpdatePwdAsync(walletId, op, np).Receive()
}

type FutureWalletInfo chan *response

func (r FutureWalletInfo) Receive() (*adajson.WalletInfo, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var info adajson.WalletInfo
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *Client) UpdateWalletInfoAsync(walletId, assuranceLevel, name string) FutureWalletInfo {

	cmd := adajson.NewUpdateWalletInfoCmd(walletId, assuranceLevel, name)
	return c.sendCmd(cmd)
}
// Update the Wallet identified by the given walletId.
func (c *Client) UpdateWalletInfo(walletId, assuranceLevel, name string) (*adajson.WalletInfo, error) {
	return c.UpdateWalletInfoAsync(walletId, assuranceLevel, name).Receive()
}


type FutureGetWallets chan *response

func (r FutureGetWallets) Receive() ([]adajson.WalletInfo, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var info []adajson.WalletInfo
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (c *Client) GetWalletsAsync(page, perPage int) FutureGetWallets {

	cmd := adajson.NewGetWalletsCmd(page, perPage)

	return c.sendCmd(cmd)
}
// Returns a list of the available wallets
func (c *Client) GetWallets(page, perPage int) ([]adajson.WalletInfo, error) {
	return c.GetWalletsAsync(page, perPage).Receive()
}

type FutureGetWallet chan *response

func (r FutureGetWallet) Receive() (*adajson.WalletInfo, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var info adajson.WalletInfo
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *Client) GetWalletAsync(walletId string) FutureCreateWallet {

	cmd := adajson.NewGetWalletCmd(walletId)
	return c.sendCmd(cmd)
}
// Returns the Wallet identified by the given walletId
func (c *Client) GetWallet(walletId string) (*adajson.WalletInfo, error) {
	return c.GetWalletAsync(walletId).Receive()
}

type FutureCreateAddress chan *response

func (r FutureCreateAddress) Receive() (*adajson.Address, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var info adajson.Address
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *Client) CreateAddressAsync(walletId, spendingPassword string, accountIndex int) FutureCreateAddress {

	cmd := adajson.NewCreateAddressCmd(walletId, spendingPassword, accountIndex)

	return c.sendCmd(cmd)
}
// Creates a new Address
func (c *Client) CreateAddress(walletId, spendingPassword string, accountIndex int) (*adajson.Address, error) {
	return c.CreateAddressAsync(walletId, spendingPassword, accountIndex).Receive()
}

type FutureGetAddress chan *response

func (r FutureGetAddress) Receive() (*adajson.Address, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var info adajson.Address
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *Client) GetAddressAsync(addressId string) FutureGetAddress {

	cmd := adajson.NewGetAddressCmd(addressId)

	return c.sendCmd(cmd)
}
// Returns interesting information about an address, if available and valid.
func (c *Client) GetAddress(addressId string) (*adajson.Address, error) {
	return c.GetAddressAsync(addressId).Receive()
}

type FutureGetAddresses chan *response

func (r FutureGetAddresses) Receive() ([]adajson.Address, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var info []adajson.Address
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (c *Client) GetAddressesAsync(page, perPage int) FutureGetAddresses {

	cmd := adajson.NewGetAddressesCmd(page, perPage)
	return c.sendCmd(cmd)
}
// Returns a list of the addresses
func (c *Client) GetAddresses(page, perPage int) ([]adajson.Address, error) {
	return c.GetAddressesAsync(page, perPage).Receive()
}

type FutureCreateAccount chan *response

func (r FutureCreateAccount) Receive() (*adajson.Account, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var info adajson.Account
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *Client) CreateAccountAsync(walletId, name, pwd string) FutureCreateAccount {

	cmd := adajson.NewCreateAccountCmd(walletId, name, pwd)

	return c.sendCmd(cmd)
}
// Creates a new Account for the given Wallet
func (c *Client) CreateAccount(walletId, name, pwd string) (*adajson.Account, error) {
	return c.CreateAccountAsync(walletId, name, pwd).Receive()
}

type FutureGetAccount chan *response

func (r FutureGetAccount) Receive() ([]adajson.Account, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var info []adajson.Account
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return info, nil
}

func (c *Client) GetAccountAsync(walletId string, page, perPage int) FutureGetAccount {

	cmd := adajson.NewGetAccountCmd(walletId, page, perPage)

	return c.sendCmd(cmd)
}
// Retrieves the full list of Accounts
func (c *Client) GetAccount(walletId string, page, perPage int) ([]adajson.Account, error) {
	return c.GetAccountAsync(walletId, page, perPage).Receive()
}

type FutureGetTransactions chan *response

func (r FutureGetTransactions) Receive() ([]*adajson.Transaction, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var txs []*adajson.Transaction
	err = json.Unmarshal(res, &txs)
	if err != nil {
		return nil, err
	}

	return txs, nil
}

func (c *Client) GetTransactionsAsync(walletId string, accountIndex int,
	sortBy string, page, pageSize int, id string, address string) FutureGetTransactions {

	cmd := adajson.NewGetTransactionsCmd(walletId, accountIndex,
		sortBy, page, pageSize, id, address)

	return c.sendCmd(cmd)
}
// Returns the transaction history, i.e the list of all the past transactions.
func (c *Client) GetTransactions(walletId string, accountIndex int,
	sortBy string, page, pageSiz int, id string, address string) ([]*adajson.Transaction, error) {

	return c.GetTransactionsAsync(walletId, accountIndex,
		sortBy, page, pageSiz, id, address).Receive()
}

type FutureCreateTransaction chan *response

func (r FutureCreateTransaction) Receive() (*adajson.Transaction, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var tx adajson.Transaction
	err = json.Unmarshal(res, &tx)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (c *Client) CreateTransactionAsync(destinations []adajson.Destination,
	source adajson.Source, pwd string, policy string) FutureCreateTransaction {

	cmd := adajson.NewCreateTransactionCmd(destinations, source, pwd, policy)
	return c.sendCmd(cmd)
}
// Generates a new transaction from the source to
// one or multiple target addresses.
func (c *Client) CreateTransaction(destinations []adajson.Destination, source adajson.Source,
	pwd string, policy string) (*adajson.Transaction, error) {

	return c.CreateTransactionAsync(destinations, source, pwd, policy).Receive()
}

type FutureEstimatingTxFees chan *response

func (r FutureEstimatingTxFees) Receive() (int64, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	type tmp struct {
		EstimatedAmount int64 `json:"estimatedAmount"`
	}
	var t tmp
	err = json.Unmarshal(res, &t)
	if err != nil {
		return 0, err
	}

	return t.EstimatedAmount, nil
}

func (c *Client) EstimatingTxFeesAsync(destinations []adajson.Destination,
	source adajson.Source, pwd string, policy *string) FutureEstimatingTxFees {

	cmd := adajson.NewEstimatingTxFeesCmd(destinations, source, pwd, policy)
	return c.sendCmd(cmd)
}
// Estimate the fees which would originate from the payment.
func (c *Client) EstimatingTxFees(destinations []adajson.Destination, source adajson.Source,
	pwd string, policy *string) (int64, error) {

	return c.EstimatingTxFeesAsync(destinations, source, pwd, policy).Receive()
}
