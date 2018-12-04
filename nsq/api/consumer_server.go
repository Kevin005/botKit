package api

import (
	"fmt"
	"log"

	"github.com/nsqio/go-nsq"

	"github.com/ifchange/botKit/nsq/consumer"
)

type ConsumerServer struct {
	ci          consumer.ConsumerI
	quitCh      chan string
	lookupHosts []string
	concurrency int
}

type ConsumerP struct {
	topicName   string
	channelName string
	lookupHosts []string
	concurrency int
}

func NewConsumer(cp *ConsumerP) (ConsumerServerI) {
	cinew, err := consumer.NewConsumer(cp.topicName, cp.channelName)
	if err != nil {
		panic(fmt.Errorf("NewConsumer error, %v", err))
	}
	cs := &ConsumerServer{
		ci:          cinew,
		lookupHosts: cp.lookupHosts,
		concurrency: cp.concurrency,
	}
	return cs
}

func (cs *ConsumerServer) StartConcurrentHandlers(handler nsq.Handler) {
	cs.ci.AddConcurrentHandlers(handler, cs.concurrency)
	err := cs.ci.ConnectToNSQLookupd(cs.lookupHosts)
	fmt.Println("ch start")
	if err != nil {
		log.Fatal("connect to nsq lookup err ", err.Error())
	}
	cs.quitCh = make(chan string)
	<-cs.quitCh //todo 退出机制
}

func (cs *ConsumerServer) StopConsumer() {
	if cs.quitCh != nil {
		cs.ci.Stop()
		cs.quitCh <- "quit"
	}
}
