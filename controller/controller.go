package controller

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/hardeepnarang10/large-ecdsa-collider/client"
	"github.com/hardeepnarang10/large-ecdsa-collider/flagparser"
	"github.com/hardeepnarang10/large-ecdsa-collider/wallet"
)

func Controller(flags flagparser.Flags, ethClient *client.Client) {
	var wallets []wallet.Wallet

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		recordMinedWallets(wallets)
		os.Exit(0)
	}()

	workerChannel := make(chan wallet.Wallet, flags.NumberOfWorkers)
	outputChannel := make(chan wallet.Wallet)

	for i := 0; i < cap(workerChannel); i++ {
		go miner(ethClient, workerChannel, outputChannel)
	}

	go func() {
		for i := 0; i < flags.NumberOfWallets; i++ {
			workerChannel <- wallet.Wallet{}
		}
	}()

	for i := 0; i < flags.NumberOfWallets; i++ {
		if flags.DebugMode {
			wallets = append(wallets, <-outputChannel)
		} else if output := <-outputChannel; output.Balance != 0 {
			wallets = append(wallets, output)
		}
	}

	close(workerChannel)
	close(outputChannel)
	recordMinedWallets(wallets)
}
