package adajson

type Address struct {
	Used          bool   `json:"used"`
	ChangeAddress bool   `json:"changeAddress"`
	Id            string `json:"id"`
}

type Account struct {
	Addresses []Address `json:"addresses"`
	Amount    int64     `json:"amount"`
	Index     int       `json:"index"`
	Name      string    `json:"name"`
	WalletID  string    `json:"walletId"`
}

type NodeInfoResult struct {
	BlockchainHeight struct {
		Quantity int    `json:"quantity"`
		Unit     string `json:"unit"`
	} `json:"blockchainHeight"`
	LocalBlockchainHeight struct {
		Quantity int    `json:"quantity"`
		Unit     string `json:"unit"`
	} `json:"localBlockchainHeight"`
	LocalTimeInformation struct {
		DifferenceFromNtpServer struct {
			Quantity int    `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"differenceFromNtpServer"`
	} `json:"localTimeInformation"`
	SyncProgress struct {
		Quantity int    `json:"quantity"`
		Unit     string `json:"unit"`
	} `json:"syncProgress"`
}

type NodeSettings struct {
	GitRevision    string `json:"gitRevision"`
	ProjectVersion string `json:"projectVersion"`
	SlotDuration   struct {
		Quantity int    `json:"quantity"`
		Unit     string `json:"unit"`
	} `json:"slotDuration"`
	SoftwareInfo struct {
		ApplicationName string `json:"applicationName"`
		Version         int    `json:"version"`
	} `json:"softwareInfo"`
}

type WalletInfo struct {
	AssuranceLevel             string `json:"assuranceLevel"`
	Balance                    int64  `json:"balance"`
	CreatedAt                  string `json:"createdAt"`
	HasSpendingPassword        bool   `json:"hasSpendingPassword"`
	ID                         string `json:"id"`
	Name                       string `json:"name"`
	SpendingPasswordLastUpdate string `json:"spendingPasswordLastUpdate"`
	SyncState struct {
		Data interface{} `json:"data"`
		Tag  string      `json:"tag"`
	} `json:"syncState"`
}

type Transaction struct {
	Amount        int    `json:"amount"`
	Confirmations int    `json:"confirmations"`
	CreationTime  string `json:"creationTime"`
	Direction     string `json:"direction"`
	ID            string `json:"id"`
	Inputs []struct {
		Address string `json:"address"`
		Amount  int64  `json:"amount"`
	} `json:"inputs"`
	Outputs []struct {
		Address string `json:"address"`
		Amount  int64  `json:"amount"`
	} `json:"outputs"`
	Status struct {
		Data struct{} `json:"data"`
		Tag  string   `json:"tag"`
	} `json:"status"`
	Type string `json:"type"`
}
