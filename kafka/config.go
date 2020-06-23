package kafka

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Shopify/sarama"
)

var (
	defaultAsyncProducer *asyncProducer
	defaultSyncProducer  *syncProducer
	defaultConsumerGroup *consumerGroup
	defaultConsumer      *consumer
)

var (
	ErrKafkaAddressesNil = errors.New("the addresses of kafka are nil")
	ConsumeCanceledError = errors.New("consume was canceled")
)

type Config struct {
	Addresses              []string
	ConsumerGroupId        string
	ConsumerOffsetsInitial string // oldest or newest
}

func generateSaramaProducerMsg(ctx context.Context, topic string, msg interface{}) (*sarama.ProducerMessage, error) {
	value := ctx.Value("Trace-ID")
	traceID := ""
	if value != nil {
		traceID = value.(string)
	}
	var producerMessage sarama.Encoder
	switch body := msg.(type) {
	case []byte:
		producerMessage = sarama.ByteEncoder(body)
	case string:
		producerMessage = sarama.StringEncoder(body)
	default:
		ms, err := json.Marshal(msg)
		if err != nil {
			return nil, err
		}
		producerMessage = sarama.ByteEncoder(ms)
	}
	pMsg := &sarama.ProducerMessage{
		Topic:   topic,
		Value:   producerMessage,
		Headers: extractHeaders(map[string]string{"Trace-ID": traceID}),
	}
	return pMsg, nil
}

func extractHeaders(headerMap map[string]string) []sarama.RecordHeader {
	headers := make([]sarama.RecordHeader, 0, len(headerMap))
	for key, value := range headerMap {
		headers = append(headers, sarama.RecordHeader{
			Key:   []byte(key),
			Value: []byte(value),
		})
	}
	return headers
}
