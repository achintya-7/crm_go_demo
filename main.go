package main

import (
	"bufio"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logConf := zap.NewDevelopmentConfig()
	logConf.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	logger, err := logConf.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // flushes buffer, if any

	confLogPath := "test"

	if confLogPath != "" {
		logPath := fmt.Sprintf("%s.log", confLogPath)
		logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			logger.Error("failed to create log file", zap.Error(err))
		}
		defer logFile.Close()

		fileSyncer := zapcore.AddSync(logFile)

		core := zapcore.NewCore(
			zapcore.NewConsoleEncoder(logConf.EncoderConfig),
			fileSyncer,
			logConf.Level,
		)

		logger = zap.New(core)

		seprator := "---------"

		scanner := bufio.NewScanner(logFile)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		sepIndex := -1
		for i, line := range lines {
			if line == seprator {
				sepIndex = i
			}
		}

		var newLines []string
		if sepIndex != -1 && sepIndex < len(lines)-1 {
			newLines = lines[sepIndex+1:]
		}

		if err = os.Truncate(logPath, 0); err != nil {
			panic(err)
		}

		writer := bufio.NewWriter(logFile)
		for _, line := range newLines {
			_, err := writer.WriteString(line + "\n")
			if err != nil {
				fmt.Println(err)
			}
		}

		writer.WriteString(seprator + "\n")

		writer.Flush()

	}

	logger.Info("Hello")
	logger.Info("Bye")
}
