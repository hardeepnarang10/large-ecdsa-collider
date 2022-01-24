package client

import (
	"context"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hardeepnarang10/large-ecdsa-collider/enum"
	"github.com/hardeepnarang10/large-ecdsa-collider/wallet"
)

type Client struct {
	client      *ethclient.Client
	httpTimeout time.Duration
}

func NewClient(ethClient *ethclient.Client, httpTimeoutSeconds int) *Client {
	return &Client{
		client:      ethClient,
		httpTimeout: time.Duration(httpTimeoutSeconds) * time.Second,
	}
}

func (c Client) QueryWalletBalance(generator func() wallet.Wallet) wallet.Wallet {
	generatedWallet := generator()

	ctx, cancel := context.WithTimeout(context.Background(), c.httpTimeout)
	currentBalance, err := c.client.BalanceAt(ctx, generatedWallet.PublicAddress, nil)
	if err != nil {
		enum.LOGGER.Error("Error occurred while querying wallet balance: " + err.Error())
	}
	defer cancel()

	// Convert balance from wei to coin units
	fbalance := new(big.Float)
	fbalance.SetString(currentBalance.String())
	coinValueBigFloat := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	coinValueFloat, _ := coinValueBigFloat.Float64()

	generatedWallet.Balance = coinValueFloat
	return generatedWallet
}
