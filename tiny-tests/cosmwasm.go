package tinytest

// TODO: investigate nocopy optimizations

// basic, standard struct (with embedded structs)
type Env struct {
	Block    BlockInfo
	Contract ContractInfo
}

type BlockInfo struct {
	Height int64
	Time   int64 `json:",string"`
}

type ContractInfo struct {
	Address string
}

// another important struct that includes a slice of structs (which caused issues with another parser)
type MessageInfo struct {
	Signer string
	// TODO: how to ensure empty funds -> [] not nil
	Funds []Coin
}

type Coin struct {
	Denom  string
	Amount string
}

// emulate Rust enum, only one should ever be set
type ExecuteMsg struct {
	Deposit  *DepositMsg  `json:",omitempty"`
	Withdraw *WithdrawMsg `json:",omitempty"`
}

type DepositMsg struct {
	ToAccount string
	Amount    string
}

// withdraws all funds
type WithdrawMsg struct {
	// use a different field here to ensure we have proper types (json must not overlap)
	FromAccount string
}
