package producer

import (
	"time"
)

type EventType int

var (
	events = [...]string{"Create", "Update", "Remove"}
)

const (
	Create EventType = iota
	Update
	Remove
)

func (e EventType) String() string {
	return events[e]
}

type Event struct {
	Type EventType
	Body map[string]interface{}
}

func CreateEvent(ev EventType, id uint64) Event {
	return Event{
		Type: ev,
		Body: map[string]interface{}{
			"id":        id,
			"event":     ev.String(),
			"timestamp": time.Now().Unix(),
		},
	}
}
