package box

import (
	"fmt"
	"os"
	"path/filepath"
)




func getBoxDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("home directory topilmadi: %w", err)
	}

	return filepath.Join(homeDir, ".flash", "boxes"), nil
}
 
func getBoxPath(name string) (string, error) {
	boxDir, err := getBoxDir()
    if err != nil {
		return "", err
	}
	return  filepath.Join(boxDir, name+".box"),nil
}