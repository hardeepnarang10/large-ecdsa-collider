package controller

import (
	"fmt"

	"github.com/hardeepnarang10/large-ecdsa-collider/client"
	"github.com/hardeepnarang10/large-ecdsa-collider/enum"
	"github.com/hardeepnarang10/large-ecdsa-collider/generator"
	"github.com/hardeepnarang10/large-ecdsa-collider/wallet"
)

func miner(ethClient *client.Client, workerChannel <-chan wallet.Wallet, outputChannel chan<- wallet.Wallet) {
	for range workerChannel {
		outputChannel <- ethClient.QueryWalletBalance(generator.GenerateWallet)
	}
}

func recordMinedWallets(wallets []wallet.Wallet) {
	if len(wallets) != 0 {
		enum.LOGGER.Warn("\n\n\t\t-----------------Failure occurred on events-----------------\n\n")
		for _, w := range wallets {
			enum.LOGGER.Error(fmt.Sprint(w))
		}
	}
}
