package core

import (
	"birthday-bot/internal/adapters/db/pg"
	"birthday-bot/internal/adapters/logger"
	"birthday-bot/internal/adapters/repo"
	"sync"
)

type St struct {
	lg       logger.Lite
	repo     repo.Repo
	notifier notifier.Notifier

	wg        sync.WaitGroup
	stopped   bool
	stoppedMu sync.RWMutex

	City *City
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

	c.City = NewCity(c)

	return c
}

func (c *St) Start() {
	go c.User.NotifyUserBirthday()
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

	c.wg.Wait()
}
