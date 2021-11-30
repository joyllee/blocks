package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/joyllee/blocks/logger"
)

type (
	consumer struct {
		exec sarama.Consumer
	}
)

func Consumer() *consumer {
	return defaultConsumer
}

func InitDefaultConsumer(config Config) {
	if len(config.Addresses) <= 0 {
		logger.Fatal(ErrKafkaAddressesNil)
	}
	defaultConsumer = &consumer{newConsumer(config)}
}

func (c *consumer) ConsumePartition(topic string, partition int32, offset int64) (sarama.PartitionConsumer, error) {
	return c.exec.ConsumePartition(topic, partition, offset)
}

func (c *consumer) Partitions(topic string) ([]int32, error) {
	return c.exec.Partitions(topic)
}

func (c *consumer) Topics() ([]string, error) {
	return c.exec.Topics()
}

func (c *consumer) HighWaterMarks() map[string]map[int32]int64 {
	return c.exec.HighWaterMarks()
}

func (c *consumer) Close() error {
	return c.exec.Close()
}

func newConsumer(config Config) sarama.Consumer {
	saramaConf := sarama.NewConfig()
	saramaConf.Version = sarama.V0_11_0_0
	/*	tlsConfig, err := createTLSByCertAndKey(tlsClientCert, tlsClientKey)
		if err == nil {
			saramaConf.Net.TLS.Enable = true
			saramaConf.Net.TLS.Config = tlsConfig
			saramaConf.Net.TLS.Config.InsecureSkipVerify = *tlsSkipVerify
		}*/

	client, err := sarama.NewConsumer(config.Addresses, saramaConf)
	if err != nil {
		logger.Fatal("Failed to creating consumer:", err)
	}
	return client
}
