package core

import (
	"birthday-bot/internal/adapters/logger"
	"birthday-bot/internal/adapters/notifier"
	"birthday-bot/internal/adapters/repo"
	"birthday-bot/internal/adapters/repo/pg"
	"context"
	"sync"
	"time"
)

type St struct {
	lg       logger.Lite
	repo     repo.Repo
	notifier notifier.Notifier

	stopped   bool
	stoppedMu sync.RWMutex

	User *User
}

func New(
	lg logger.Lite,
	repo *pg.St,
	notifier notifier.Notifier,
) *St {
	c := &St{
		lg:       lg,
		repo:     repo,
		notifier: notifier,
	}
	c.User = NewUser(c)

	return c
}

func (c *St) Start(ctx context.Context, notifyHour int) {
	// todo notify time
	timer := time.NewTimer(getSleepDuration(notifyHour))
	for {
		select {
		case <-timer.C:
			c.User.NotifyBirthday()
			timer.Reset(getSleepDuration(notifyHour))
		case <-ctx.Done():
			timer.Stop()
			return
		}
	}
}

func getSleepDuration(hour int) time.Duration {
	now := time.Now()
	next := now.Add(time.Hour * 24)
	h, m, s := next.Clock()
	next.Add(-time.Hour * time.Duration(h-hour))
	next.Add(-time.Minute * time.Duration(m))
	next.Add(-time.Second * time.Duration(s))
	return next.Sub(now)
}

func (c *St) IsStopped() bool {
	c.stoppedMu.RLock()
	defer c.stoppedMu.RUnlock()
	return c.stopped
}

func (c *St) StopAndWaitJobs() {
	c.stoppedMu.Lock()
	c.stopped = true
	c.stoppedMu.Unlock()

}
