package core

import (
	"birthday-bot/internal/adapters/logger"
	"birthday-bot/internal/adapters/notifier"
	"birthday-bot/internal/adapters/repo"
	"birthday-bot/internal/adapters/repo/pg"
	"birthday-bot/pkg/clock"
	"context"
	"sync"
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

func (c *St) Start(ctx context.Context, interval *clock.TimeInterval) {
	for {

		select {
		case <-interval.WaitNext():
			notifyCtx, _ := context.WithTimeout(ctx, interval.Timeout())
			c.User.NotifyBirthday(notifyCtx)

		case <-ctx.Done():
			return
		}
	}
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
