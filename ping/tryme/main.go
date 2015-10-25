package main

import (
	"fmt"
	"os"

	"github.com/dcrosby42/status/ping"
	"github.com/dcrosby42/status/ping/osx_strategy"
)

func main() {
	host := "google.com"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}

	pinger := ping.CreatePinger(osx_strategy.Ping, host, 1)
	for i := 0; i < 5; i++ {
		fmt.Printf("%s\n", pinger())
	}
}
