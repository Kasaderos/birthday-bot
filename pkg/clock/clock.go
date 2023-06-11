package clock

import (
	"fmt"
	"strings"
	"time"
)

type TimeInterval struct {
	Start time.Time
	End   time.Time

	nowRun bool
}

func (tl *TimeInterval) WaitNext() <-chan time.Time {
	now := time.Now()
	if !tl.nowRun && now.After(tl.Start) && now.Before(tl.End) {
		timer := time.NewTimer(0)
		tl.nowRun = true
		return timer.C
	}

	// next day
	next := now.Add(time.Hour * 24)
	next = truncateToDay(next)

	h, m, _ := tl.Start.Clock()
	next = next.Add(
		time.Hour*time.Duration(h) +
			time.Minute*time.Duration(m))

	timer := time.NewTimer(next.Sub(now))
	return timer.C
}

func truncateToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func (tl TimeInterval) Timeout() time.Duration {
	return tl.End.Sub(tl.Start)
}

// NewTimeInterval is constructor of TimeInterval.
// It parses intervalConf by format HH24:MM-HH24:MM
func NewTimeInterval(intervalConf string) (*TimeInterval, error) {
	clocks := strings.Split(intervalConf, "-")
	if len(clocks) != 2 {
		return nil, fmt.Errorf("invalid notify time %s (format HH:MM-HH:MM)", intervalConf)
	}
	t1, err := time.Parse("15:04", clocks[0])
	if err != nil {
		return nil, fmt.Errorf("invalid clock %s (format HH:MM)", clocks[0])
	}
	t2, err := time.Parse("15:04", clocks[1])
	if err != nil {
		return nil, fmt.Errorf("invalid clock %s (format HH:MM)", clocks[1])
	}

	if t2.Before(t1) {
		return nil, fmt.Errorf("invalid interval %s", intervalConf)
	}

	return &TimeInterval{
		Start: t1,
		End:   t2,
	}, nil
}
