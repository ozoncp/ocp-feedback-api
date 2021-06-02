package alarmer

import (
	"time"
)

// Alarmer is the interface that allows consumer to be notified over a channel
type Alarmer interface {
	Alarm() <-chan void
}

type void struct{}

type alarm struct {
	duration time.Duration
	alarms   chan void
	done     chan void
}

// New returns a new alarmer object
func New(duration time.Duration) *alarm {
	return &alarm{
		duration: duration,
		alarms:   make(chan void),
		done:     make(chan void),
	}
}

// Init starts repeatedly delivering asynchronous
// notifications at regular intervals until Close is called
func (a *alarm) Init() {
	go func() {
		ticker := time.NewTicker(a.duration)
		defer close(a.alarms)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				a.alarms <- void{}
			case <-a.done:
				return
			}
		}
	}()
}

// Alarm provides an access to notification channel
func (a *alarm) Alarm() <-chan void {
	return a.alarms
}

// Close notifies alarmer that no more alarms should be delivered
func (a *alarm) Close() {
	close(a.done)
}
