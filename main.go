package main

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/zserge/lorca"
)

func main() {
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"
	// start chrome

	// var ui lorca.UI
	// currentDir, _ := os.Getwd()
	// dir := filepath.Join(currentDir, ".cache")
	// ui, _ = lorca.New("https://baidu.com", dir, 800, 600)
	chSignal := make(chan os.Signal, 1)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	select {
	// case <-ui.Done():
	case <-chSignal:
	}
	// ui.Close()
}
