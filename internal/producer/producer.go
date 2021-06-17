package producer

import (
	"context"
	"encoding/json"

	"github.com/rs/zerolog/log"

	"github.com/Shopify/sarama"
)

type void struct{}

type Producer interface {
	SendEvent(ev Event)
	Close()
}

type producer struct {
	prod   sarama.AsyncProducer
	topic  string
	events chan Event
	done   chan void
}

func New(topic string, prod sarama.AsyncProducer) *producer {

	return &producer{
		prod:   prod,
		topic:  topic,
		events: make(chan Event),
		done:   make(chan void),
	}
}

func (p *producer) Init(ctx context.Context) {

	go func() {
		for err := range p.prod.Errors() {
			//TODO: write message to disk to read and re-transmit later
			log.Err(err).Msg("Failed to write entry")
		}
	}()
	go func() {
		defer p.prod.Close()
		for {
			select {
			case event := <-p.events:
				bytes, err := json.Marshal(event)
				if err != nil {
					log.Err(err).Msg("Failed to marshal event:")
					return
				}
				p.prod.Input() <- &sarama.ProducerMessage{Topic: p.topic, Value: sarama.StringEncoder(bytes)}
			case <-ctx.Done():
				p.done <- void{}
				return
			}
		}
	}()
}

func (p *producer) SendEvent(ev Event) {
	p.events <- ev
}

func (p *producer) Close() {
	<-p.done
}
