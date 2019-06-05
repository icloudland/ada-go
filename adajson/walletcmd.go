package adajson

type CreateWalletCmd struct {
	Operation        string   `json:"operation"`
	BackupPhrase     []string `json:"backupPhrase"`
	AssuranceLevel   string   `json:"assuranceLevel"`
	Name             string   `json:"name"`
	SpendingPassword string   `json:"spendingPassword"`
}

func NewCreateWalletCmd(operation string, backupPhrase []string,
	assuranceLevel, name, spendingPassword string) *CreateWalletCmd {
	return &CreateWalletCmd{
		Operation:        operation,
		BackupPhrase:     backupPhrase,
		AssuranceLevel:   assuranceLevel,
		Name:             name,
		SpendingPassword: spendingPassword,
	}
}

type CreateAddressCmd struct {
	WalletId         string `json:"walletId"`
	AccountIndex     int    `json:"accountIndex"`
	SpendingPassword string `json:"spendingPassword"`
}

func NewCreateAddressCmd(walletId, spendingPassword string, accountIndex int) *CreateAddressCmd {
	return &CreateAddressCmd{
		WalletId:         walletId,
		AccountIndex:     accountIndex,
		SpendingPassword: spendingPassword,
	}
}

type GetAddressesCmd struct {
	Page    int `qstring:"page"`
	PerPage int `qstring:"per_page"`
}

func NewGetAddressesCmd(page, perPage int) *GetAddressesCmd {
	return &GetAddressesCmd{
		Page:    page,
		PerPage: perPage,
	}
}

type GetAddressCmd struct {
	AddressId string `path:"addressId" qstring:"-"`
}

func NewGetAddressCmd(addressId string) *GetAddressCmd {
	return &GetAddressCmd{
		AddressId: addressId,
	}
}

type CreateAccountCmd struct {
	WalletId         string `json:"-" path:"walletId"`
	Name             string `json:"name"`
	SpendingPassword string `json:"spendingPassword"`
}

func NewCreateAccountCmd(walletId, name, spendingPassword string) *CreateAccountCmd {
	return &CreateAccountCmd{
		WalletId:         walletId,
		Name:             name,
		SpendingPassword: spendingPassword,
	}
}

type GetAccountCmd struct {
	WalletId string `path:"walletId" qstring:"-"`
	Page     int    `qstring:"page"`
	PerPage  int    `qstring:"per_page"`
}

func NewGetAccountCmd(walletId string, page, perPage int) *GetAccountCmd {
	return &GetAccountCmd{
		WalletId: walletId,
		Page:     page,
		PerPage:  perPage,
	}
}

type NodeInfoCmd struct{}

func NewNodeInfoCmd() *NodeInfoCmd {
	return &NodeInfoCmd{}
}

type NodeSettingsCmd struct{}

func NewNodeSettingsCmd() *NodeSettingsCmd {
	return &NodeSettingsCmd{}
}

type GetTransactionsCmd struct {
	WalletId     string `json:"wallet_id" qstring:"wallet_id"`
	AccountIndex int    `json:"account_index" qstring:"account_index"`
	SortBy       string `json:"sort_by" qstring:"sort_by"`
	Page         int    `json:"page" qstring:"page"`
	PerPage      int    `json:"per_page" qstring:"per_page"`
	ID           string `json:"id" qstring:"id"`
	Address      string `json:"address" qstring:"address"`
}

func NewGetTransactionsCmd(walletId string, accountIndex int,
	sortBy string, page, perPage int, id string, address string) *GetTransactionsCmd {
	return &GetTransactionsCmd{
		WalletId:     walletId,
		AccountIndex: accountIndex,
		SortBy:       sortBy,
		Page:         page,
		PerPage:      perPage,
		ID:           id,
		Address:      address,
	}
}

type CreateTransactionCmd struct {
	Destinations     []Destination `json:"destinations"`
	Source           Source        `json:"source"`
	SpendingPassword string        `json:"spendingPassword"`
	GroupingPolicy   string        `json:"groupingPolicy"`
}

func NewCreateTransactionCmd(destinations []Destination, source Source,
	pwd string, gp string) *CreateTransactionCmd {
	return &CreateTransactionCmd{
		Destinations:     destinations,
		Source:           source,
		SpendingPassword: pwd,
		GroupingPolicy:   gp,
	}
}

type EstimatingTxFeesCmd struct {
	Destinations     []Destination `json:"destinations"`
	Source           Source        `json:"source"`
	SpendingPassword string        `json:"spendingPassword"`
	GroupingPolicy   *string       `json:"groupingPolicy"`
}

func NewEstimatingTxFeesCmd(destinations []Destination, source Source,
	pwd string, groupingPolicy *string) *EstimatingTxFeesCmd {
	return &EstimatingTxFeesCmd{
		Destinations:     destinations,
		Source:           source,
		SpendingPassword: pwd,
		GroupingPolicy:   groupingPolicy,
	}
}

type Destination struct {
	Amount  int64  `json:"amount"`
	Address string `json:"address"`
}

type Source struct {
	AccountIndex int    `json:"accountIndex"`
	WalletId     string `json:"walletId"`
}

type GetWalletCmd struct {
	WalletId string `json:"walletId" qstring:"-" path:"walletId"`
}

func NewGetWalletCmd(walletId string) *GetWalletCmd {
	return &GetWalletCmd{
		WalletId: walletId,
	}
}

type GetWalletsCmd struct {
	Page    int `qstring:"page"`
	PerPage int `qstring:"per_page"`
}

func NewGetWalletsCmd(page, perPage int) *GetWalletsCmd {
	return &GetWalletsCmd{
		Page:    page,
		PerPage: perPage,
	}
}

type UpdatePwdCmd struct {
	WalletId string `json:"-" path:"walletId"`
	Old      string `json:"old"`
	New      string `json:"new"`
}

func NewUpdatePwdCmd(walletId, oldPwd, newPwd string) *UpdatePwdCmd {
	return &UpdatePwdCmd{
		WalletId: walletId,
		Old:      oldPwd,
		New:      newPwd,
	}
}

type UpdateWalletInfoCmd struct {
	WalletId       string `json:"-" path:"walletId"`
	AssuranceLevel string `json:"assuranceLevel"`
	Name           string `json:"name"`
}

func NewUpdateWalletInfoCmd(walletId, assuranceLevel, name string) *UpdateWalletInfoCmd {
	return &UpdateWalletInfoCmd{
		WalletId:       walletId,
		AssuranceLevel: assuranceLevel,
		Name:           name,
	}
}

type DeleteWalletCmd struct {
	WalletId string `json:"-" path:"walletId"`
}

func NewDeleteWalletCmd(walletId string) *DeleteWalletCmd {
	return &DeleteWalletCmd{
		WalletId: walletId,
	}
}

func init() {
	// No special flags for commands in this file.
	flags := UsageFlag(0)
	MustRegisterCmd("addresses:post", (*CreateAddressCmd)(nil), flags)
	MustRegisterCmd("addresses:get", (*GetAddressesCmd)(nil), flags)
	MustRegisterCmd("addresses/{{addressId}}:get", (*GetAddressCmd)(nil), flags)

	MustRegisterCmd("node-info:get", (*NodeInfoCmd)(nil), flags)
	MustRegisterCmd("node-settings:get", (*NodeSettingsCmd)(nil), flags)

	MustRegisterCmd("wallets:post", (*CreateWalletCmd)(nil), flags)
	MustRegisterCmd("wallets:get", (*GetWalletsCmd)(nil), flags)
	MustRegisterCmd("wallets/{{walletId}}:get", (*GetWalletCmd)(nil), flags)
	MustRegisterCmd("wallets/{{walletId}}/accounts:post", (*CreateAccountCmd)(nil), flags)
	MustRegisterCmd("wallets/{{walletId}}/accounts:get", (*GetAccountCmd)(nil), flags)
	MustRegisterCmd("wallets/{{walletId}}/accounts:delete", (*GetAccountCmd)(nil), flags)
	MustRegisterCmd("wallets/{{walletId}}/password:put", (*UpdatePwdCmd)(nil), flags)
	MustRegisterCmd("wallets/{{walletId}}:put", (*UpdateWalletInfoCmd)(nil), flags)
	MustRegisterCmd("wallets/{{walletId}}:delete", (*DeleteWalletCmd)(nil), flags)

	MustRegisterCmd("transactions:get", (*GetTransactionsCmd)(nil), flags)
	MustRegisterCmd("transactions:post", (*CreateTransactionCmd)(nil), flags)
	MustRegisterCmd("transactions/fees:post", (*EstimatingTxFeesCmd)(nil), flags)

}
