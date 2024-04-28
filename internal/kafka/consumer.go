package kafka

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

type consumer struct {
	ready   chan bool
	handler func(ctx context.Context, content []byte) error
}

func (c *consumer) Setup(_ sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				return nil
			}
			err := c.handler(session.Context(), msg.Value)
			if err != nil {
				log.Printf("failed to process message: %v\n", err)
				continue
			}
			session.MarkMessage(msg, "")
		case <-session.Context().Done():
			return nil
		}
	}
}

func Run(ctx context.Context, handler func(ctx context.Context, content []byte) error, brokers []string, group string, topics []string) error {
	cfg := sarama.NewConfig()
	cfg.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	c := &consumer{}
	c.handler = handler

	client, err := sarama.NewConsumerGroup(brokers, group, cfg)
	if err != nil {
		return fmt.Errorf("consumer group: %w", err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			c.ready = make(chan bool)
			err := client.Consume(ctx, topics, c)
			if err != nil {
				return fmt.Errorf("consumer group: %w", err)
			}
		}
	}
}
