package ping

import "time"

func StartPinging(pinger Pinger, delay time.Duration) (chan PingRecord, chan int) {
	pings := make(chan PingRecord)
	stop := make(chan int)
	go func() {
		ever := true
		for ever {
			select {
			case pings <- pinger():
				time.Sleep(delay)
			case <-stop:
				ever = false
			}
		}
		close(pings)
	}()
	return pings, stop
}
