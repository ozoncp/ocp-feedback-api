package alarmer

import (
	"context"
	"time"
)

// Alarmer is the interface that allows consumer to be notified over a channel
type Alarmer interface {
	Alarm() <-chan struct{}
}

type void struct{}

type alarmer struct {
	duration time.Duration
	alarms   chan void
	done     chan void
}

// New returns a new alarmer object
func New(duration time.Duration) *alarmer {
	return &alarmer{
		duration: duration,
		alarms:   make(chan void),
		done:     make(chan void),
	}
}

// Init starts repeatedly delivering asynchronous
// notifications at regular intervals until context is cancelled
func (a *alarmer) Init(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(a.duration)
		defer close(a.alarms)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				select {
				case a.alarms <- void{}:
				default:
				}
			case <-ctx.Done():
				a.done <- void{}
				return
			}
		}
	}()
}

// Alarm provides an access to notification channel
func (a *alarmer) Alarm() <-chan void {
	return a.alarms
}

// WaitClose waits until alarmer is closed
func (a *alarmer) WaitClosed() {
	<-a.done
}
