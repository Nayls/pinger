package main

import (
	"log"
	"os"

	"github.com/Nayls/pinger/cmd/pinger"
)

func init() {
	if err := pinger.InitCobraConfig().Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func main() {
}
