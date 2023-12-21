package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"sweetbot/conf/config"
	"sweetbot/internal/handler"
)

func loadConfig(relativePath string) error {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("runtime error")
	}

	rootDir := filepath.Join(filepath.Dir(filename), "../..")

	configPath := filepath.Join(rootDir, relativePath)
	config.Conf.LoadConfig(configPath)
	return nil
}

func main() {
	loadConfig("conf/env.yaml")

	http.HandleFunc("/webhook", handler.LineBotHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
