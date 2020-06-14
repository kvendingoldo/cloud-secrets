package controller

import (
	"context"
	"github.com/kvendingoldo/cloud-secrets/provider"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// Controller is responsible for orchestrating the different components.
type Controller struct {
	Provider provider.Provider

	// The interval between individual synchronizations
	Interval time.Duration

	SecretName string

	// The nextRunAt used for throttling and batching reconciliation
	nextRunAt time.Time
	// The nextRunAtMux is for atomic updating of nextRunAt
	nextRunAtMux sync.Mutex
}

// RunOnce runs a single iteration of a reconciliation loop.
func (c *Controller) RunOnce(ctx context.Context) error {

	c.Provider.GetSecret(c.SecretName)

	return nil
}

// MinInterval is used as window for batching events
const MinInterval = 5 * time.Second

// RunOnceThrottled makes sure execution happens at most once per interval.
func (c *Controller) ScheduleRunOnce(now time.Time) {
	c.nextRunAtMux.Lock()
	defer c.nextRunAtMux.Unlock()
	c.nextRunAt = now.Add(MinInterval)
}

func (c *Controller) ShouldRunOnce(now time.Time) bool {
	c.nextRunAtMux.Lock()
	defer c.nextRunAtMux.Unlock()
	if now.Before(c.nextRunAt) {
		return false
	}
	c.nextRunAt = now.Add(c.Interval)
	return true
}

// Run runs RunOnce in a loop with a delay until context is canceled
func (c *Controller) Run(ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		if c.ShouldRunOnce(time.Now()) {
			if err := c.RunOnce(ctx); err != nil {
				log.Error(err)
			}
		}
		select {
		case <-ticker.C:
		case <-ctx.Done():
			log.Info("Terminating main controller loop")
			return
		}
	}
}
