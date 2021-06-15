package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Shopify/sarama"
)

var brokerList = []string{"127.0.0.1:29092"}

type Producer interface {
	SendEvent(ev Event)
}

type producer struct {
	prod   sarama.AsyncProducer
	topic  string
	events chan Event
}

func New(topic string) (*producer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

	prod, err := sarama.NewAsyncProducer(brokerList, config)
	if err != nil {
		return nil, fmt.Errorf("failed to start Sarama producer:%v", err)
	}
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
				message := &sarama.ProducerMessage{Topic: p.topic, Value: sarama.StringEncoder(bytes)}
				p.prod.Input() <- message
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (p *producer) SendEvent(ev Event) {
	p.events <- ev
}
