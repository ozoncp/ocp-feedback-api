package producer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
)

type Producer interface {
	SendEvent(ev Event)
}

type producer struct {
	prod   sarama.AsyncProducer
	topic  string
	events chan Event
}

func New(topic string, prod sarama.AsyncProducer) (*producer, error) {

	return &producer{
		prod:   prod,
		topic:  topic,
		events: make(chan Event),
	}, nil
}

func (p *producer) Init(ctx context.Context) {
	go func() {
		for err := range p.prod.Errors() {
			log.Println("Failed to write entry:", err)
			//TODO: write message to disk to read and re-transmit later
		}
	}()
	go func() {
		defer p.prod.Close()
		for {
			select {
			case event := <-p.events:
				bytes, err := json.Marshal(event)
				if err != nil {
					log.Println("Failed to marshal event:", err)
					return
				}
				p.prod.Input() <- &sarama.ProducerMessage{Topic: p.topic, Value: sarama.StringEncoder(bytes)}
			case <-ctx.Done():
				return
			default:
				log.Println("Failed to send message, broker is down")
				//TODO: write message to disk to read and re-transmit later
			}
		}
	}()
}

func (p *producer) SendEvent(ev Event) {
	p.events <- ev
}
