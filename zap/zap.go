package zap

import (
	"log"
	"os"
	"path/filepath"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetWorkingDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalln("Error getting CWD! Exiting...")
	}
	return path
}

// Create logging directory if it doesn't exist
func createDirectoryIfNotExists(path string) {
	if _, err := os.Stat(filepath.Join(path, LOG_DIRECTORY)); os.IsNotExist(err) {
		log.Println("Creating logger directory " + LOG_DIRECTORY)
		errCreatingDir := os.Mkdir(LOG_DIRECTORY, os.ModePerm)
		if errCreatingDir != nil {
			log.Fatalln("Error creating directory " + LOG_DIRECTORY)
		}
	}
}

// Configure logger output
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// Implement Lumberjack for log rotation
func getLogWriter(path string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(path, LOG_DIRECTORY, LOG_FILE),
		MaxSize:    LOG_FILE_MAX_SIZE,
		MaxBackups: LOG_FILE_MAX_NUMBER,
		MaxAge:     LOG_FILE_MAX_AGE,
		Compress:   LOG_FILE_ENABLE_COMPRESSION,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// Initialize Zap Logger
func InitZapLogger(logLevel zapcore.Level) *zap.Logger {
	path := GetWorkingDirectory()
	createDirectoryIfNotExists(path)
	writerSyncer := getLogWriter(path)
	encoder := getEncoder()
	return zap.New(
		zapcore.NewCore(
			encoder,
			writerSyncer,
			logLevel,
		),
		zap.AddCaller(),
	)
}
