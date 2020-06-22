package kafka

var (
	ErrKafkaAddressesNil = errors.New("the addresses of kafka are nil")
	ConsumeCanceledError = errors.New("consume was canceled")
)

type Config struct {
	Addresses              []string
	ConsumerGroupId        string
	ConsumerOffsetsInitial string // oldest or newest
}
