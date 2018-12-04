package consumer

import "github.com/nsqio/go-nsq"

type ConsumerI interface {
	AddConcurrentHandlers(handler nsq.Handler, concurrency int)
	ConnectToNSQLookupd(lookupHosts []string) error
	Stop()
}
