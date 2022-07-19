package basic

import (
	"errors"
	"time"

	"github.com/zjkmxy/go-ndn/pkg/ndn"
)

type Timer struct{}

func NewTimer() ndn.Timer {
	return Timer{}
}

func (_ Timer) Sleep(d time.Duration) {
	time.Sleep(d)
}

func (_ Timer) Schedule(d time.Duration, f func()) func() error {
	t := time.AfterFunc(d, f)
	return func() error {
		if t != nil {
			t.Stop()
			t = nil
			return nil
		} else {
			return errors.New("Event has already been canceled")
		}
	}
}

func (_ Timer) Now() time.Time {
	return time.Now()
}
