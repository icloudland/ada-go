package adajson

type ShelleyCreateWalletCmd struct {
	BackupPhrase     []string `json:"mnemonic_sentence"`
	Name             string   `json:"name"`
	SpendingPassword string   `json:"passphrase"`
}

func NewShelleyCreateWalletCmd(backupPhrase []string, name, spendingPassword string) *ShelleyCreateWalletCmd {
	return &ShelleyCreateWalletCmd{
		BackupPhrase:     backupPhrase,
		Name:             name,
		SpendingPassword: spendingPassword,
	}
}

type GetShelleyWalletCmd struct {
	WalletId string `json:"walletId" qstring:"-" path:"walletId"`
}

func NewGetShelleyWalletCmd(walletId string) *GetShelleyWalletCmd {
	return &GetShelleyWalletCmd{
		WalletId: walletId,
	}
}

// 时间是大于等于，小于等于
type GetShelleyTransactionsCmd struct {
	WalletId string `json:"walletId" qstring:"-" path:"walletId"`
	SortBy   string `json:"order" qstring:"order"` // descending, ascending
	Start    string `json:"start" qstring:"start"`
	End      string `json:"end" qstring:"end"`
}

func NewGetShelleyTransactionsCmd(walletId, sortBy, startTime, EndTime string) *GetShelleyTransactionsCmd {
	return &GetShelleyTransactionsCmd{
		WalletId: walletId,
		SortBy:   sortBy,
		Start:    startTime,
		End:      EndTime,
	}
}

type GetShelleyAddressesCmd struct {
	WalletId string `json:"walletId" qstring:"-" path:"walletId"`
	State    string `json:"state" qstring:"state"` //used,unused
}

func NewGetShelleyAddressesCmd(walletId, state string) *GetShelleyAddressesCmd {
	return &GetShelleyAddressesCmd{
		WalletId: walletId,
		State:    state,
	}
}

type CreateShelleyTransactionCmd struct {
	WalletId   string                      `json:"walletId" qstring:"-" path:"walletId"`
	Passphrase string                      `json:"passphrase"`
	Payments   []ShelleyTransactionPayment `json:"payments"`
	Withdrawal *string                      `json:"withdrawal"`
}

type ShelleyTransactionPayment struct {
	Address string        `json:"address"`
	Amount  ShelleyAmount `json:"amount"`
}

type ShelleyAmount struct {
	Quantity int64  `json:"quantity"`
	Unit     string `json:"unit"`
}

type ShelleyAddress struct {
	Id    string `json:"id"`
	State string `json:"state"`
}

func NewCreateShelleyTransactionCmd(walletId, passphrase, address, withdrawal string, amount int64) *CreateShelleyTransactionCmd {
	pm := ShelleyTransactionPayment{
		Address: address,
		Amount:  ShelleyAmount{Quantity: amount, Unit: "lovelace",},
	}
	return &CreateShelleyTransactionCmd{
		WalletId:   walletId,
		Passphrase: passphrase,
		Payments: []ShelleyTransactionPayment{
			pm,
		},
	}
}

type GetShelleyTransactionCmd struct {
	WalletId      string `json:"walletId" qstring:"-" path:"walletId"`
	TransactionId string `json:"transactionId" qstring:"-" path:"transactionId"`
}

func NewGetShelleyTransactionCmd(walletId, transactionId string) *GetShelleyTransactionCmd {

	return &GetShelleyTransactionCmd{
		WalletId:      walletId,
		TransactionId: transactionId,
	}
}

type EstimateFeeCmd struct {
	WalletId   string                      `json:"walletId" qstring:"-" path:"walletId"`
	Payments   []ShelleyTransactionPayment `json:"payments"`
	Withdrawal *string                      `json:"withdrawal"`
}

func NewEstimateFeeCmd(walletId, address string, withdrawal *string, amount int64) *EstimateFeeCmd {
	pm := ShelleyTransactionPayment{
		Address: address,
		Amount:  ShelleyAmount{Quantity: amount, Unit: "lovelace",},
	}
	return &EstimateFeeCmd{
		WalletId:   walletId,
		Withdrawal: withdrawal,
		Payments: []ShelleyTransactionPayment{
			pm,
		},
	}
}

type ShelleyWalletInfo struct {
	AddressPoolGap int64 `json:"address_pool_gap"`
	Balance        struct {
		Available struct {
			Quantity int64  `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"available"`
		Reward struct {
			Quantity int64  `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"reward"`
		Total struct {
			Quantity int64  `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"total"`
	} `json:"balance"`
	Delegation struct {
		Active struct {
			Status string `json:"status"`
			Target string `json:"target"`
		} `json:"active"`
		Next []struct {
			ChangesAt struct {
				EpochNumber    int64  `json:"epoch_number"`
				EpochStartTime string `json:"epoch_start_time"`
			} `json:"changes_at"`
			Status string `json:"status"`
		} `json:"next"`
	} `json:"delegation"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Passphrase struct {
		LastUpdatedAt string `json:"last_updated_at"`
	} `json:"passphrase"`
	State struct {
		Status string `json:"status"`
	} `json:"state"`
	Tip struct {
		EpochNumber int64 `json:"epoch_number"`
		Height      struct {
			Quantity int64  `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"height"`
		SlotNumber int64 `json:"slot_number"`
	} `json:"tip"`
}

type ShelleyTransaction struct {
	Amount struct {
		Quantity int64  `json:"quantity"`
		Unit     string `json:"unit"`
	} `json:"amount"`
	Depth struct {
		Quantity int64  `json:"quantity"`
		Unit     string `json:"unit"`
	} `json:"depth"`
	Direction string `json:"direction"`
	ID        string `json:"id"`
	Inputs    []struct {
		Address string `json:"address"`
		Amount  struct {
			Quantity int64  `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"amount"`
		ID    string `json:"id"`
		Index int64  `json:"index"`
	} `json:"inputs"`
	InsertedAt struct {
		Block struct {
			EpochNumber int64 `json:"epoch_number"`
			Height      struct {
				Quantity int64  `json:"quantity"`
				Unit     string `json:"unit"`
			} `json:"height"`
			SlotNumber int64 `json:"slot_number"`
		} `json:"block"`
		Time string `json:"time"`
	} `json:"inserted_at"`
	Metadata struct {
		Zero  string  `json:"0"`
		One   int64   `json:"1"`
		Two   string  `json:"2"`
		Three []int64 `json:"3"`
		Four  struct {
			One4 int64  `json:"14"`
			Key  string `json:"key"`
		} `json:"4"`
	} `json:"metadata"`
	Outputs []struct {
		Address string `json:"address"`
		Amount  struct {
			Quantity int64  `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"amount"`
	} `json:"outputs"`
	PendingSince struct {
		Block struct {
			EpochNumber int64 `json:"epoch_number"`
			Height      struct {
				Quantity int64  `json:"quantity"`
				Unit     string `json:"unit"`
			} `json:"height"`
			SlotNumber int64 `json:"slot_number"`
		} `json:"block"`
		Time string `json:"time"`
	} `json:"pending_since"`
	Status      string `json:"status"`
	Withdrawals []struct {
		Amount struct {
			Quantity int64  `json:"quantity"`
			Unit     string `json:"unit"`
		} `json:"amount"`
		StakeAddress string `json:"stake_address"`
	} `json:"withdrawals"`
}

type EstimateFee struct {
	EstimatedMax struct {
		Quantity int64  `json:"quantity"`
		Unit     string `json:"unit"`
	} `json:"estimated_max"`
	EstimatedMin struct {
		Quantity int64  `json:"quantity"`
		Unit     string `json:"unit"`
	} `json:"estimated_min"`
}
