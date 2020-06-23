package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	config2 "github.com/joyllee/blocks/config"
)

type syncProducer struct {
	exec sarama.SyncProducer
}

func SyncProducer() *syncProducer {
	return defaultSyncProducer
}

func InitDefaultSyncProducer(config Config) {
	if len(config.Addresses) <= 0 {
		config2.Fatal(ErrKafkaAddressesNil)
	}
	defaultSyncProducer = &syncProducer{newSyncProducer(config)}
}

func (p *syncProducer) Send(ctx context.Context, topic string, msg interface{}) (partition int32, offset int64, err error) {
	pMsg, err := generateSaramaProducerMsg(ctx, topic, msg)
	if err != nil {
		return 0, 0, err
	}
	return p.exec.SendMessage(pMsg)
}

func (p *syncProducer) Close() error {
	return p.exec.Close()
}

func newSyncProducer(config Config) sarama.SyncProducer {
	saramaConf := sarama.NewConfig()
	saramaConf.Version = sarama.V0_11_0_0
	saramaConf.Producer.Retry.Max = 3
	saramaConf.Producer.RequiredAcks = sarama.WaitForAll
	saramaConf.Producer.Return.Successes = true
	saramaConf.Producer.Compression = sarama.CompressionSnappy
	//tlsConfig := createTlsConfiguration()
	//if tlsConfig != nil {
	//	saramaConf.Net.TLS.Config = tlsConfig
	//	saramaConf.Net.TLS.Enable = true
	//}
	producer, err := sarama.NewSyncProducer(config.Addresses, saramaConf)
	if err != nil {
		config2.Fatal("Failed to start Sarama producer:", err)
	}
	return producer
}
