package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"log"
)

type (
	consumerGroup struct {
		exec sarama.ConsumerGroup
	}

	PublicationHandlerFunc func(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error
	consumerGroupHandler   struct {
		handler PublicationHandlerFunc
	}
)

func (*consumerGroupHandler) Setup(sess sarama.ConsumerGroupSession) error {
	return nil
}
func (*consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}
func (h *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	return h.handler(sess, claim)
}

func ConsumerGroup() *consumerGroup {
	return defaultConsumerGroup
}

func InitDefaultConsumerGroup(config Config) {
	if len(config.Addresses) <= 0 {
		log.Fatal(ErrKafkaAddressesNil)
	}
	defaultConsumerGroup = &consumerGroup{newConsumerGroup(config)}
}

func (c *consumerGroup) Consume(ctx context.Context, topics []string, handlerFunc PublicationHandlerFunc) error {
	handler := &consumerGroupHandler{handler: handlerFunc}
	return c.exec.Consume(ctx, topics, handler)
}

func (c *consumerGroup) Errors() <-chan error {
	return c.exec.Errors()
}
func (c *consumerGroup) Close() error {
	return c.exec.Close()
}

func newConsumerGroup(config Config) sarama.ConsumerGroup {
	saramaConf := sarama.NewConfig()
	saramaConf.Version = sarama.V0_11_0_0
	saramaConf.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	saramaConf.Consumer.Offsets.Initial = sarama.OffsetNewest
	switch config.ConsumerOffsetsInitial {
	case "oldest":
		saramaConf.Consumer.Offsets.Initial = sarama.OffsetOldest
	case "newest":
		saramaConf.Consumer.Offsets.Initial = sarama.OffsetNewest
	}
	client, err := sarama.NewConsumerGroup(config.Addresses, config.ConsumerGroupId, saramaConf)
	if err != nil {
		log.Fatal("Failed to creating consumer group:", err)
	}
	return client
}
