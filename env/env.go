package env // delete me

import ( // delete me
	"fmt"           // delete me
	"os"            // delete me
	"path/filepath" // delete me
) // delete me

func GetAppDir() (string, error) { // delete me
	homeDir, err := os.UserHomeDir() // delete me
	if err != nil {                  // delete me
		return "", fmt.Errorf("home dir error %v", err) // delete me
	} // delete me
	appDir := filepath.Join(homeDir, ".local", "share", "todo") // delete me
	if err = os.MkdirAll(appDir, os.ModePerm); err != nil {     // delete me
		return "", fmt.Errorf("error creating dir %v", err) // delete me
	} // delete me
	return appDir, nil // delete me
} // delete me
