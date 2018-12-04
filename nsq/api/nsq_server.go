package api

import (
	"github.com/nsqio/go-nsq"

	"github.com/ifchange/botKit/nsq/message"
)

type PublisherServerI interface {
	Publish(msg *message.Message) error
	StopPublisher()
}

type ConsumerServerI interface {
	StartConcurrentHandlers(handler nsq.Handler)
	StopConsumer()
}
