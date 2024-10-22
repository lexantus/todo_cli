package env

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAppDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home dir error %v", err)
	}
	appDir := filepath.Join(homeDir, ".local", "share", "todo")
	if err = os.MkdirAll(appDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("error creating dir %v", err)
	}
	return appDir, nil
}
