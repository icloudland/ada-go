package ada

import (
	"encoding/json"
	"github.com/icloudland/ada-go/adajson"
)

type FutureShelleyCreateWallet chan *response

func (r FutureShelleyCreateWallet) Receive() (*adajson.WalletInfo, error) {
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

func (c *Client) CreateShelleyWalletAsync(backupPhrase []string,
	name, spendingPassword string) FutureShelleyCreateWallet {

	cmd := adajson.NewShelleyCreateWalletCmd(backupPhrase,
		name, spendingPassword)

	return c.sendCmd(cmd)
}

// Creates a new  Wallet
func (c *Client) CreateShelleyWallet(backupPhrase []string,
	name, spendingPassword string) (*adajson.WalletInfo, error) {

	return c.CreateShelleyWalletAsync(backupPhrase, name, spendingPassword).Receive()
}

type FutureGetShelleyWallet chan *response

func (r FutureGetShelleyWallet) Receive() (*adajson.ShelleyWalletInfo, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var info adajson.ShelleyWalletInfo
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func (c *Client) GetShelleyWalletAsync(walletId string) FutureGetShelleyWallet {

	cmd := adajson.NewGetShelleyWalletCmd(walletId)
	return c.sendCmd(cmd)
}
func (c *Client) GetShelleyWallet(walletId string) (*adajson.ShelleyWalletInfo, error) {
	return c.GetShelleyWalletAsync(walletId).Receive()
}

type FutureGetShelleyTransactions chan *response

func (r FutureGetShelleyTransactions) Receive() ([]*adajson.ShelleyTransaction, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var txs []*adajson.ShelleyTransaction
	err = json.Unmarshal(res, &txs)
	if err != nil {
		return nil, err
	}

	return txs, nil
}

func (c *Client) GetShelleyTransactionsAsync(walletId, sortBy, startTime, EndTime string) FutureGetShelleyTransactions {

	cmd := adajson.NewGetShelleyTransactionsCmd(walletId, sortBy, startTime, EndTime)
	return c.sendCmd(cmd)
}
func (c *Client) GetShelleyTransactions(walletId, sortBy, startTime, EndTime string) ([]*adajson.ShelleyTransaction, error) {
	return c.GetShelleyTransactionsAsync(walletId, sortBy, startTime, EndTime).Receive()
}

type FutureGetShelleyAddresses chan *response

func (r FutureGetShelleyAddresses) Receive() ([]adajson.ShelleyAddress, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var addrs []adajson.ShelleyAddress
	err = json.Unmarshal(res, &addrs)
	if err != nil {
		return nil, err
	}

	return addrs, nil
}

func (c *Client) GetShelleyAddressesAsync(walletId, state string) FutureGetShelleyAddresses {

	cmd := adajson.NewGetShelleyAddressesCmd(walletId, state)
	return c.sendCmd(cmd)
}
func (c *Client) GetShelleyAddresses(walletId, state string) ([]adajson.ShelleyAddress, error) {
	return c.GetShelleyAddressesAsync(walletId, state).Receive()
}

type FutureCreateShelleyTransaction chan *response

func (r FutureCreateShelleyTransaction) Receive() (*adajson.ShelleyTransaction, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var tx adajson.ShelleyTransaction
	err = json.Unmarshal(res, &tx)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (c *Client) CreateShelleyTransactionAsync(walletId, passphrase, address, withdrawal string, amount int64) FutureCreateShelleyTransaction {

	cmd := adajson.NewCreateShelleyTransactionCmd(walletId, passphrase, address, withdrawal, amount)
	return c.sendCmd(cmd)
}
func (c *Client) CreateShelleyTransaction(walletId, passphrase, address, withdrawal string, amount int64) (*adajson.ShelleyTransaction, error) {
	return c.CreateShelleyTransactionAsync(walletId, passphrase, address, withdrawal, amount).Receive()
}

type FutureGetShelleyTransaction chan *response

func (r FutureGetShelleyTransaction) Receive() (*adajson.ShelleyTransaction, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var tx adajson.ShelleyTransaction
	err = json.Unmarshal(res, &tx)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (c *Client) GetShelleyTransactionAsync(walletId, txId string) FutureGetShelleyTransaction {

	cmd := adajson.NewGetShelleyTransactionCmd(walletId, txId)
	return c.sendCmd(cmd)
}
func (c *Client) GetShelleyTransaction(walletId, txId string) (*adajson.ShelleyTransaction, error) {
	return c.GetShelleyTransactionAsync(walletId, txId).Receive()
}



type FutureEstimateFee chan *response

func (r FutureEstimateFee) Receive() (*adajson.EstimateFee, error) {
	res, _, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	var tx adajson.EstimateFee
	err = json.Unmarshal(res, &tx)
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (c *Client) EstimateFeeAsync(walletId, address string, withdrawal *string, amount int64) FutureEstimateFee {
	cmd := adajson.NewEstimateFeeCmd(walletId, address, withdrawal, amount)
	return c.sendCmd(cmd)
}
func (c *Client) EstimateFee(walletId, address string, withdrawal *string, amount int64) (*adajson.EstimateFee, error) {
	return c.EstimateFeeAsync(walletId, address, withdrawal, amount).Receive()
}
