package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func NewLoggerToVolume() {
	dir := "./logs"
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Directory does not exist:", dir)
			err := os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}

			fmt.Println("Directory created successfully:", dir)
		} else {
			fmt.Println("Error checking directory:", err)
		}
	}
	absPath, err := filepath.Abs(dir)
	if err != nil {
		fmt.Println("Error reading given path:", err)
	}

	vollogger, err := os.OpenFile(absPath+"/new_relic.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	VolLogger = log.New(vollogger, "", 0)
}
