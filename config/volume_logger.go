package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func NewLoggerToVolume() {
	dir := "./logs"
	if err := ensureDir(dir); err != nil {
		fmt.Println("Error:", err)
		return
	}

	logFile, err := setupLogFile(dir, "new_relic.log")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	VolLogger = log.New(logFile, "", 0)
}

func ensureDir(dir string) error {
	_, err := os.Stat(dir)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		fmt.Println("Directory does not exist:", dir)
		err := os.MkdirAll(dir, 0755) // Changed to 0755 for better security
		if err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}

		fmt.Println("Directory created successfully:", dir)
		return nil
	}

	return fmt.Errorf("error checking directory %s: %v", dir, err)
}

func setupLogFile(dir, fileName string) (*os.File, error) {
	// Ensure the directory exists
	if err := ensureDir(dir); err != nil {
		return nil, err
	}

	logFilePath := filepath.Join(dir, fileName) // No need for Abs since ensureDir checks

	// Adjusted file permissions to 0666
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %v", logFilePath, err)
	}

	return logFile, nil
}
