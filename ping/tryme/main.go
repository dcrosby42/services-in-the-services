package main

import (
	"fmt"
	"os"
	"time"

	"github.com/dcrosby42/status/ping"
	"github.com/dcrosby42/status/ping/osx_strategy"
)

func main() {
	host := "google.com"
	if len(os.Args) > 1 {
		host = os.Args[1]
	}

	// DoFivePings(host)

	pinger := ping.CreatePinger(osx_strategy.Ping, host, 1)
	pings, stop := ping.StartPinging(pinger, 500*time.Millisecond)

	go func() {
		time.Sleep(5 * time.Second)
		close(stop)
	}()

	for r := range pings {
		fmt.Printf("%s\n", r)
	}
}

func DoFivePings(host string) {
	pinger := ping.CreatePinger(osx_strategy.Ping, host, 1)
	for i := 0; i < 5; i++ {
		fmt.Printf("%s\n", pinger())
	}
}
