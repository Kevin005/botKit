package api

import (
	"log"
	"testing"
	"time"

	"github.com/ifchange/botKit/nsq/mock"
	"github.com/nsqio/go-nsq"
)

func TestConsumerServer_StartHandler(t *testing.T) {
	nc := NewConsumer(&ConsumerP{
		topicName:   mock.TopicName,
		channelName: mock.ChannelName,
		lookupHosts: []string{mock.LookupdHost},
		concurrency: 10,
	})

	go func() {
		time.Sleep(10 * time.Second)
		nc.StopConsumer()
	}()

	nc.StartConcurrentHandlers(nsq.HandlerFunc(func(message *nsq.Message) error {
		message.Finish()
		//do something
		log.Printf("Got a message: %v", string(message.Body))
		return nil
	}))

	t.Log("test consumer exit")
}
