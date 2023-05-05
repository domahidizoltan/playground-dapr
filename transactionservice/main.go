package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/dapr/go-sdk/service/common"
	"github.com/domahidizoltan/playground-dapr/common/client"
	"github.com/domahidizoltan/playground-dapr/common/dto"
	"github.com/domahidizoltan/playground-dapr/common/model"
	"github.com/domahidizoltan/playground-dapr/transactionservice/ent"
	entGen "github.com/domahidizoltan/playground-dapr/transactionservice/ent/generated"
	"github.com/google/uuid"

	dapr "github.com/dapr/go-sdk/client"
)

const (
	service          = "TRANSACTION"
	debitPubSubName  = "debitsource"
	creditPubSubName = "creditdest"
	topicName        = "transfer"
)

var (
	entClient      *entGen.Client
	daprClient     dapr.Client
	debitSourceSub = &common.Subscription{
		PubsubName: "debit-source",
		Topic:      "transfer",
		Route:      "/debitsource",
	}
	creditDestSub = &common.Subscription{
		PubsubName: "credit-dest",
		Topic:      "transfer",
		Route:      "/creditdest",
	}
)

func main() {
	entClient = ent.GetClient(service)
	if entClient == nil {
		panic(errors.New("failed to create database connection"))
	}

	var err error
	daprClient, err = dapr.NewClient()
	if err != nil {
		log.Fatalf("failed to init DAPR client: %s", err)
	}
	defer daprClient.Close()

	subscriptions := []client.SubscriptionHandler{
		{
			Subscription: debitSourceSub,
			Handler:      doTransaction(true),
		},
		{
			Subscription: creditDestSub,
			Handler:      doTransaction(false),
		},
	}
	client.SubscribeTopic(service, subscriptions)
}

func doTransaction(isDebit bool) func(ctx context.Context, event *common.TopicEvent) (bool, error) {
	return func(ctx context.Context, event *common.TopicEvent) (bool, error) {
		var transfer *dto.Transfer
		if err := event.Struct(&transfer); err != nil {
			log.Printf("failed to read transfer %s: %s", string(event.RawData), err)
			return true, err
		}

		dest := transfer.DstAcc
		src := transfer.SrcAcc
		var pubSubName string
		var updateBalance dto.UpdateBalance
		switch isDebit {
		case true:
			dest = model.BankAccount
			pubSubName = debitPubSubName
			updateBalance = dto.UpdateBalance{
				Account: transfer.DstAcc,
				Amount:  -transfer.Amount,
			}
		case false:
			src = model.BankAccount
			pubSubName = creditPubSubName
			updateBalance = dto.UpdateBalance{
				Account: transfer.SrcAcc,
				Amount:  transfer.Amount,
			}

		}

		newTransaction, err := entClient.Transaction.Create().
			SetID(uuid.New()).
			SetTnx(transfer.Tnx).
			SetDestAcc(dest).
			SetSourceAcc(src).
			SetAmount(float64(transfer.Amount)).
			SetDatetime(time.Now().UTC()).
			Save(ctx)
		if err != nil {
			log.Printf("failed to save transaction for %v: %s", transfer, err)
			return true, err
		}

		log.Printf("saved transaction %+v", newTransaction)

		if err := daprClient.PublishEvent(ctx, pubSubName, topicName, updateBalance); err != nil {
			log.Printf("failed to update balance %v: %s", updateBalance, err)
			return false, err
		}

		return false, nil
	}
}
