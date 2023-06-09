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
	PubsubName        = "pubsub"
)

type SubscriptionHandler struct {
	Subscription *common.Subscription
	Handler      common.TopicEventHandler
}

func SubscribeTopic(servicePrefix string, subscriptions []SubscriptionHandler, serviceHook func(s common.Service)) {
	port := helper.GetEnv(servicePrefix+subscriberPortKey, "3001")
	log.Printf("starting subscription service %s on port %s", servicePrefix, port)
	s := daprd.NewService(":" + port)
	if serviceHook != nil {
		serviceHook(s)
	}

	for _, sub := range subscriptions {
		if err := s.AddTopicEventHandler(sub.Subscription, sub.Handler); err != nil {
			log.Fatalf("error adding topic subscription for %s: %v", sub.Subscription.Topic, err)
		}
	}

	if err := s.Start(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("%s failed listenning: %v", servicePrefix, err)
	}
}
