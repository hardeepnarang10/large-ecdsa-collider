package enum

import (
	"runtime"

	"go.uber.org/zap"
)

var (
	TIMEOUT_BATCH_SECONDS int    = 15
	NODE_ADDRESS          string = "<Insert Node Address>"
	NUMBER_OF_WORKERS     int    = 8 * runtime.NumCPU()
	NUMBER_OF_WALLETS     int    = 100
	DEBUG_MODE            bool   = false
)

var (
	LOGGER *zap.Logger
)
