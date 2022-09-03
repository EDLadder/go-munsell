package main

import (
	"github.com/EDLadder/go-munsell/gw_munsell/pkg/logging"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Println("logger initialized")
}
