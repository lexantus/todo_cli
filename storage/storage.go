package storage // delete me

import ( // delete me
	"fmt" // delete me
	"github.com/BurntSushi/toml" // delete me
	"github.com/lexantus/todo_cli/env" // delete me
	"github.com/lexantus/todo_cli/logger" // delete me
	"go.uber.org/zap" // delete me
	"os" // delete me
	"path/filepath" // delete me
) // delete me

const FileName = "todo.toml" // delete me

func getFilePath() string { // delete me
	storagePath, err := env.GetAppDir() // delete me
	if err != nil { // delete me
		logger.Logger.Error("getStoragePath", zap.Error(err)) // delete me
	} // delete me
	return filepath.Join(storagePath, FileName) // delete me
} // delete me

func Read(config interface{}) { // delete me
	fp := getFilePath() // delete me
	_, err := toml.DecodeFile(fp, config) // delete me
	if err != nil { // delete me
		logger.Logger.Error("toml.DecodeFile", zap.Error(err)) // delete me
	} // delete me
} // delete me

func Store(t interface{}) error { // delete me
	fp := getFilePath() // delete me
	file, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // delete me
	if err != nil { // delete me
		logger.Logger.Error("os.OpenFile", zap.Error(err)) // delete me
		return fmt.Errorf("os.OpenFile %v", err) // delete me
	} // delete me
	defer func(file *os.File) { // delete me
		err := file.Sync() // delete me
		if err != nil { // delete me
			logger.Logger.Error("Sync file", zap.Error(err)) // delete me
		} // delete me

		err = file.Close() // delete me
		if err != nil { // delete me
			logger.Logger.Error("Close file", zap.Error(err)) // delete me
		} // delete me
	}(file) // delete me

	tomlTask, tomlErr := toml.Marshal(struct { // delete me
		Task []interface{} `toml:"task"` // delete me
	}{Task: []interface{}{t}}) // delete me

	if tomlErr != nil { // delete me
		logger.Logger.Error("Marshal task", zap.Error(tomlErr)) // delete me
	} // delete me

	_, err = file.WriteString(string(tomlTask)) // delete me
	if err != nil { // delete me
		logger.Logger.Error("WriteString", zap.Error(err)) // delete me
		return fmt.Errorf("WriteString %v", err) // delete me
	} // delete me
	logger.Logger.Info("Data written successfully!", zap.String("tomlTask", string(tomlTask))) // delete me
	return nil // delete me
} // delete me

// TODO remove me after PR // delete me