package ping

import "time"

func CreatePinger(pingFn PingExecFunc, host string, timeout int) Pinger {
	return func() PingRecord {
		rec := PingRecord{
			Host: host,
			Sent: time.Now(),
		}
		err := pingFn(host, timeout)
		rec.Elapsed = rec.Sent.Sub(time.Now())
		rec.Error = err
		return rec
	}
}
