package main

import (
	"gitlab.com/nayls.cloud/ping.nayls.cloud/pinger/internal/config"
)

func main() {
	config.CobraServerConfiguration()
	config.ViperServerConfiguration()
}
