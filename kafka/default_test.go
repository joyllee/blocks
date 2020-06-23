package kafka

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"testing"
	"time"
)

func TestConsumer(t *testing.T) {
	InitDefaultConsumer(Config{Addresses: []string{"127.0.0.1:9092"}})
	defer Consumer().Close()
	topics, err := Consumer().Topics()
	if err != nil {
		t.Error(err)
	}

	for _, topic := range topics {
		t.Log(topic)
		partitions, err := Consumer().Partitions(topic)
		if err != nil {
			t.Error(err)
		}
		for _, partition := range partitions {
			t.Log(partition)
		}
	}

	partitionConsumer, err := Consumer().ConsumePartition("person", 0, 65557)
	if err != nil {
		t.Error(err)
	} else {
		for message := range partitionConsumer.Messages() {
			t.Log(message.Offset)
		}
	}
}

func TestConsumerGroup(t *testing.T) {
	InitDefaultConsumerGroup(Config{
		Addresses:              []string{"127.0.0.1:9092"},
		ConsumerGroupId:        "group-1",
		ConsumerOffsetsInitial: "newest",
	})

	go func() {
		for err := range ConsumerGroup().Errors() {
			t.Error(err)
		}
	}()

	ws := make(chan struct{})
	go func() {
		<-time.After(5 * time.Second)
		ws <- struct{}{}
	}()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := ConsumerGroup().Consume(ctx, []string{"test"}, func(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
		for {
			select {
			case <-ctx.Done():
				return ConsumeCanceledError
			case msg := <-claim.Messages():
				fmt.Println("receive:", msg.Offset)
				sess.MarkMessage(msg, "")
			}
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}

}

type Person struct {
	Name string
	Age  uint8
}

func TestAsyncProducer(t *testing.T) {
	InitDefaultAsyncProducer(Config{
		Addresses: []string{"127.0.0.1:9092"},
	})

	go func() {
		for {
			e := <-AsyncProducer().Errors()
			fmt.Printf("error: %v \n", e)
		}
	}()

	for {
		t := time.Now()
		person := &Person{
			Name: fmt.Sprintf("%d", t.UnixNano()),
			Age:  uint8(t.Nanosecond()),
		}
		err := AsyncProducer().Send(context.WithValue(context.Background(), "traceID", "asdasd"), "person", person)
		fmt.Println(err)
	}
}

func TestSyncProducer(t *testing.T) {
	InitDefaultSyncProducer(Config{
		Addresses: []string{"127.0.0.1:9092"},
	})

	for {
		t := <-time.After(200 * time.Millisecond)
		person := &Person{
			Name: fmt.Sprintf("%d", t.UnixNano()),
			Age:  uint8(t.Nanosecond()),
		}
		fmt.Println(SyncProducer().Send(context.WithValue(context.Background(), "Trace-ID", "asdasd"), "person", person))
	}
}
