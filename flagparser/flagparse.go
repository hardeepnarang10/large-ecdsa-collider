package flagparser

import (
	"flag"
	"fmt"

	"github.com/hardeepnarang10/large-ecdsa-collider/enum"
)

type Flags struct {
	NodeAddress     string
	NumberOfWallets int
	NumberOfWorkers int
	TimeoutBatch    int
	DebugMode       bool
}

func ParseFlags(flags *Flags) *Flags {
	flag.BoolVar(&flags.DebugMode, "debug", enum.DEBUG_MODE, fmt.Sprint("Execution mode. Defaults to ", enum.DEBUG_MODE, "."))
	flag.StringVar(&flags.NodeAddress, "node", enum.NODE_ADDRESS, "Node address of chain. Defaults to Polygon mainnet.")
	flag.IntVar(&flags.TimeoutBatch, "timeout", enum.TIMEOUT_BATCH_SECONDS, fmt.Sprint("Batch job timeout in seconds. Defaults to ", enum.TIMEOUT_BATCH_SECONDS, " seconds."))
	flag.IntVar(&flags.NumberOfWallets, "wallets", enum.NUMBER_OF_WALLETS, fmt.Sprint("Number of wallets to load test. Defaults to ", enum.NUMBER_OF_WALLETS, " wallets."))
	flag.IntVar(&flags.NumberOfWorkers, "workers", enum.NUMBER_OF_WORKERS, "Number of workers. Defaults to system's number of CPUs.")

	flag.Parse()
	return flags
}
