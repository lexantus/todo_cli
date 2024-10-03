package tasks

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

var Logger *zap.Logger

func Init() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"log/todo.log"}

	Logger, _ = config.Build()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}(Logger)

	Logger.Info("Init", zap.String("outputPath", "log/todo.log"))
}

func getStoragePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home dir error %v", err)
	}
	storageDir := filepath.Join(homeDir, ".local", "share", "todo")
	if err = os.MkdirAll(storageDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("error creating dir %v", err)
	}
	return storageDir, nil
}

func Store(t Task) error {
	storagePath, err := getStoragePath()
	if err != nil {
		Logger.Error("getStoragePath", zap.Error(err))
		return fmt.Errorf("getStoragePath %v", err)
	}
	fp := filepath.Join(storagePath, "todo.toml")
	file, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		Logger.Error("os.OpenFile", zap.Error(err))
		return fmt.Errorf("os.OpenFile %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			Logger.Error("Close file", zap.Error(err))
		}
	}(file)

	tomlTask, tomlErr := toml.Marshal(struct {
		Task Task `toml:"[task]"`
	}{t})

	if tomlErr != nil {
		Logger.Error("Marshal task", zap.Error(tomlErr))
	}

	_, err = file.WriteString(string(tomlTask))
	if err != nil {
		Logger.Error("WriteString", zap.Error(err))
		return fmt.Errorf("WriteString %v", err)
	}
	fmt.Printf("Data written successfully! [%s]\n", tomlTask)
	Logger.Info("Data written successfully!", zap.String("tomlTask", string(tomlTask)))
	return nil
}
