package ping

import (
	"time"

	"github.com/dcrosby42/status/timer"
)

func StartPinging(pinger Pinger, timer timer.Timer, delay time.Duration) (chan PingRecord, chan int) {
	pings := make(chan PingRecord)
	stop := make(chan int)
	go func() {
		ever := true
		for ever {
			timer.In(time.Now())
			select {
			case pings <- pinger():
				timer.WaitFromIn(delay)
			case <-stop:
				ever = false
			}
		}
		close(pings)
	}()
	return pings, stop
}
