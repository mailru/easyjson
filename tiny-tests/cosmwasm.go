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

/**** Test Helpers ****/

// For Testing
func (a *MessageInfo) Equals(b *MessageInfo) bool {
	if a.Signer != b.Signer {
		return false
	}
	if len(a.Funds) != len(b.Funds) {
		return false
	}
	for i, af := range a.Funds {
		bf := b.Funds[i]
		if af != bf {
			return false
		}
	}
	return true
}

// For Testing
func (a *ExecuteMsg) Equals(b *ExecuteMsg) bool {
	if a.Deposit == nil && b.Deposit != nil {
		return false
	}
	if a.Deposit != nil {
		if b.Deposit == nil {
			return false
		}
		if *a.Deposit != *b.Deposit {
			return false
		}
	}

	if a.Withdraw == nil && b.Withdraw != nil {
		return false
	}
	if a.Withdraw != nil {
		if b.Withdraw == nil {
			return false
		}
		if *a.Withdraw != *b.Withdraw {
			return false
		}
	}

	return true
}
