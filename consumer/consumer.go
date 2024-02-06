package consumer

// Consumer adalah interface untuk fungsi consume ke berbagai third party event driven (RabbitMQ, Redis, dll.)
type Consumer interface {
	ConsumeMessages(callback func([]byte) error) error
	Close()
}
