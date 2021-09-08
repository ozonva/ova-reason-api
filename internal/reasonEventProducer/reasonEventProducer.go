package reasonEventProducer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog"
)

type Event struct {
	Id        string
	Operation string
	Body      string
}

const topic string = "reasonLog"

type Producer interface {
	Publish(event Event)
}

type ProducerConfig struct {
	Host string
	Port string
}

type ReasonEventProducer struct {
	ctx      context.Context
	logger   *zerolog.Logger
	config   ProducerConfig
	messages chan *sarama.ProducerMessage
	producer sarama.SyncProducer
}

func NewProducer(ctx context.Context, config ProducerConfig) (Producer, error) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	zLogger := zerolog.New(output).With().Timestamp().Logger()
	producer := ReasonEventProducer{
		ctx:      ctx,
		logger:   &zLogger,
		config:   config,
		messages: make(chan *sarama.ProducerMessage, 5),
	}

	err := producer.init()
	if err != nil {
		return nil, err
	}

	return &producer, nil
}

func (p *ReasonEventProducer) Publish(event Event) {
	bytes, _ := json.Marshal(event)
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(event.Id),
		Value: sarama.ByteEncoder(bytes),
	}

	p.messages <- msg
}

func (p *ReasonEventProducer) init() error {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	brokers := []string{fmt.Sprintf("%s:%s", p.config.Host, p.config.Port)}

	var err error
	p.producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return err
	}

	go func() {
		defer p.producer.Close()
		defer close(p.messages)

		for {
			select {
			case msg := <-p.messages:
				partition, offset, err := p.producer.SendMessage(msg)
				if err != nil {
					p.logger.Err(err).Msgf("Failed to produce message: %s", err.Error())
				}
				fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", msg.Topic, partition, offset)
			case <-p.ctx.Done():
				return
			}
		}
	}()

	return nil
}
