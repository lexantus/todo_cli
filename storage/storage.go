package storage

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/lexantus/todo_cli/env"
	"github.com/lexantus/todo_cli/logger"
	"github.com/lexantus/todo_cli/tasks"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type Storage struct {
	filepath string
}

type tomlValues struct {
	Tasks []tasks.Task `toml:"task"`
}

func newTomlValues(task tasks.Task) tomlValues {
	return tomlValues{Tasks: []tasks.Task{task}}
}

func NewStorage() *Storage {
	storagePath, err := env.GetAppDir()
	if err != nil {
		logger.Error("New storage", zap.Error(err))
	}
	return &Storage{
		filepath: filepath.Join(storagePath, "todo.toml"),
	}
}

func (s *Storage) Read() []tasks.Task {
	fp := s.filepath
	var v tomlValues
	_, err := toml.DecodeFile(fp, &v)
	if err != nil {
		logger.Error("toml.DecodeFile", zap.Error(err))
	}
	return v.Tasks
}

func (s *Storage) Store(t tasks.Task) error {
	fp := s.filepath
	file, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("os.OpenFile", zap.Error(err))
		return fmt.Errorf("os.OpenFile %v", err)
	}
	defer func(file *os.File) {
		err := file.Sync()
		if err != nil {
			logger.Error("Sync file", zap.Error(err))
		}

		err = file.Close()
		if err != nil {
			logger.Error("Close file", zap.Error(err))
		}
	}(file)

	tomlTask, tomlErr := toml.Marshal(newTomlValues(t))

	if tomlErr != nil {
		logger.Error("Marshal task", zap.Error(tomlErr))
	}

	_, err = file.Write(tomlTask)
	if err != nil {
		logger.Error("file.Write", zap.Error(err))
		return fmt.Errorf("file.Write %v", err)
	}
	logger.Info("Data written successfully!", zap.String("tomlTask", string(tomlTask)))
	return nil
}
