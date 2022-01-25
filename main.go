package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hardeepnarang10/large-ecdsa-collider/client"
	"github.com/hardeepnarang10/large-ecdsa-collider/controller"
	"github.com/hardeepnarang10/large-ecdsa-collider/enum"
	"github.com/hardeepnarang10/large-ecdsa-collider/flagparser"
	"github.com/hardeepnarang10/large-ecdsa-collider/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Parse flags
	var flags flagparser.Flags
	flagparser.ParseFlags(&flags)

	// Initialize logger
	var zapLoggerLevel zapcore.Level
	if flags.DebugMode {
		zapLoggerLevel = zapcore.DebugLevel
	} else {
		zapLoggerLevel = zapcore.InfoLevel
	}
	enum.LOGGER = zap.InitZapLogger(zapLoggerLevel)
	defer enum.LOGGER.Sync()

	// Initialize ethereum client
	ethereumClient, err := ethclient.Dial(flags.NodeAddress)
	if err != nil {
		enum.LOGGER.Fatal("Failure in initializing ethereum client: " + err.Error())
	}
	ethClient := client.NewClient(ethereumClient, flags.TimeoutBatch)

	// Initialize controller
	controller.Controller(flags, ethClient)
}
