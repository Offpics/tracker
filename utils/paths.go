package utils

import (
	"os"
	"path/filepath"
)

func GetCsvPath() string {
	homeDir, err := os.UserHomeDir()
	Check(err)
	fileName := ".tracker.csv"
	return filepath.Join(homeDir, fileName)
}
