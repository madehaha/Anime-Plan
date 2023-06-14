package main

import (
	"backend/cmd"
	"backend/internal/logger"
)

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		logger.Fatal("Failed to Start")
	}
}
