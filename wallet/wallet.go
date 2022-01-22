package wallet

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Wallet struct {
	PrivateKey    string
	PublicKey     string
	PublicAddress common.Address
	Balance       float64
}

func (w Wallet) String() string {
	return fmt.Sprintf(
		"{\n\tEvent Index: %f,\n\tEvent Hash: %s\n}",
		w.Balance,
		w.PrivateKey,
	)
}
