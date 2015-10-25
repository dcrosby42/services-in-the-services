package ping

import (
	"fmt"
	"time"
)

type Pinger func() PingRecord

type PingExecFunc func(host string, timeout int) error

type PingRecord struct {
	Host    string
	Sent    time.Time
	Elapsed time.Duration
	Error   error
}

func (me PingRecord) Ok() bool {
	return me.Error == nil
}

func (me PingRecord) String() string {
	ok := "Ok"
	if !me.Ok() {
		ok = "FAIL"
	}
	return fmt.Sprintf("Ping %s: %s", me.Host, ok)
}
