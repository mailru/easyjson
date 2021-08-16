package tinytest

// TODO: investigate nocopy optimizations

// basic, standard struct (with embedded structs)
type Env struct {
	Block    BlockInfo    `json:"block"`
	Contract ContractInfo `json:"contract"`
}

type BlockInfo struct {
	Height int64 `json:"height"`
	Time   int64 `json:"time,string"`
}

type ContractInfo struct {
	Address string `json:"address"`
}

// another important struct that includes a slice of structs (which caused issues with another parser)
type MessageInfo struct {
	Signer string `json:"signer"`
	// TODO: how to ensure empty funds -> [] not nil
	Funds []Coin `json:"funds"`
}

type Coin struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

// emulate Rust enum, only one should ever be set
type ExecuteMsg struct {
	Deposit  *DepositMsg  `json:"deposit,omitempty"`
	Withdraw *WithdrawMsg `json:"withdraw,omitempty"`
}

type DepositMsg struct {
	ToAccount string `json:"to_account"`
	Amount    string `json:"amount"`
}

// withdraws all funds
type WithdrawMsg struct {
	// use a different field here to ensure we have proper types (json must not overlap)
	FromAccount string `json:"from_account"`
}
