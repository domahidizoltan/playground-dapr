package client

import (
	"log"
	"net/http"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/http"
	"github.com/domahidizoltan/playground-dapr/common/helper"
)

const (
	subscriberPortKey = "_SUBSCRIBER_PORT"
)

func SubscribeTopic(servicePrefix string, sub *common.Subscription, handler common.TopicEventHandler) {
	port := helper.GetEnv(servicePrefix+subscriberPortKey, "3001")
	log.Printf("starting subscriber %s on port %s", sub.PubsubName, port)
	s := daprd.NewService(":" + port)
	if err := s.AddTopicEventHandler(sub, handler); err != nil {
		log.Fatalf("error adding topic subscription for %s: %v", sub.Topic, err)
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("error listenning %s: %v", sub.PubsubName, err)
	}
}
