package ping

import "time"

func CreatePinger(pingFn PingExecFunc, host string, timeout int) Pinger {
	return func() PingRecord {
		rec := PingRecord{
			Host: host,
			Sent: time.Now(),
		}
		err := pingFn(host, timeout)
		rec.Elapsed = time.Now().Sub(rec.Sent)
		rec.Error = err
		return rec
	}
}
