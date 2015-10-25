package timer

import "time"

type Timer interface {
	In(time.Time)
	WaitFromIn(time.Duration)
}

type SimpleTimer struct {
	mark time.Time
}

func NewSimpleTimer() *SimpleTimer {
	return &SimpleTimer{time.Now()}
}

func (me *SimpleTimer) In(t time.Time) {
	me.mark = t
}

func (me *SimpleTimer) WaitFromIn(span time.Duration) {
	now := time.Now()
	remain := me.mark.Add(span).Sub(now)
	if remain > 0 {
		time.Sleep(remain)
	}
}
