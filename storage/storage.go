package storage

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/lexantus/todo_cli/env"
	"github.com/lexantus/todo_cli/logger"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

const FileName = "todo.toml"

func getFilePath() string {
	storagePath, err := env.GetAppDir()
	if err != nil {
		logger.Logger.Error("getStoragePath", zap.Error(err))
	}
	return filepath.Join(storagePath, FileName)
}

func Read(config interface{}) {
	fp := getFilePath()
	_, err := toml.DecodeFile(fp, config)
	if err != nil {
		logger.Logger.Error("toml.DecodeFile", zap.Error(err))
	}
}

func Store(t interface{}) error {
	fp := getFilePath()
	file, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Logger.Error("os.OpenFile", zap.Error(err))
		return fmt.Errorf("os.OpenFile %v", err)
	}
	defer func(file *os.File) {
		err := file.Sync()
		if err != nil {
			logger.Logger.Error("Sync file", zap.Error(err))
		}

		err = file.Close()
		if err != nil {
			logger.Logger.Error("Close file", zap.Error(err))
		}
	}(file)

	tomlTask, tomlErr := toml.Marshal(struct {
		Task []interface{} `toml:"task"`
	}{Task: []interface{}{t}})

	if tomlErr != nil {
		logger.Logger.Error("Marshal task", zap.Error(tomlErr))
	}

	_, err = file.WriteString(string(tomlTask))
	if err != nil {
		logger.Logger.Error("WriteString", zap.Error(err))
		return fmt.Errorf("WriteString %v", err)
	}
	logger.Logger.Info("Data written successfully!", zap.String("tomlTask", string(tomlTask)))
	return nil
}

// TODO remove me after PR
