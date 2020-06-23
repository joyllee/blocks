package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/joyllee/blocks/logger"
	"time"
)

type asyncProducer struct {
	exec sarama.AsyncProducer
}

func AsyncProducer() *asyncProducer {
	return defaultAsyncProducer
}

func InitDefaultAsyncProducer(config Config) {
	if len(config.Addresses) <= 0 {
		logger.Fatal(ErrKafkaAddressesNil)
	}
	defaultAsyncProducer = &asyncProducer{newAsyncProducer(config)}
}

func (p *asyncProducer) Send(ctx context.Context, topic string, msg interface{}) error {
	pMsg, err := generateSaramaProducerMsg(ctx, topic, msg)
	if err != nil {
		return err
	}
	p.exec.Input() <- pMsg
	return nil
}

func (p *asyncProducer) AsyncClose() {
	p.exec.AsyncClose()
}

func (p *asyncProducer) Close() error {
	return p.exec.Close()
}

func (p *asyncProducer) Successes() <-chan *sarama.ProducerMessage {
	return p.exec.Successes()
}

func (p *asyncProducer) Errors() <-chan *sarama.ProducerError {
	return p.exec.Errors()
}

func newAsyncProducer(config Config) sarama.AsyncProducer {
	saramaConf := sarama.NewConfig()
	//tlsConfig := createTlsConfiguration()
	//if tlsConfig != nil {
	//	saramaConf.Net.TLS.Enable = true
	//	saramaConf.Net.TLS.Config = tlsConfig
	//}
	saramaConf.Version = sarama.V0_11_0_0
	saramaConf.Producer.RequiredAcks = sarama.WaitForLocal // Only wait for the leader to ack
	saramaConf.Producer.Compression = sarama.CompressionSnappy
	saramaConf.Producer.Flush.Frequency = 1 * time.Second
	saramaConf.Producer.Flush.Messages = 10000

	producer, err := sarama.NewAsyncProducer(config.Addresses, saramaConf)
	if err != nil {
		logger.Fatal("Failed to start Sarama producer:", err)
	}
	return producer
}
