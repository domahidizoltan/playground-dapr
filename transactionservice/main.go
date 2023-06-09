package main

import (
	"context"
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dapr/go-sdk/service/common"
	"github.com/domahidizoltan/playground-dapr/common/client"
	"github.com/domahidizoltan/playground-dapr/common/dto"
	"github.com/domahidizoltan/playground-dapr/common/helper"
	"github.com/domahidizoltan/playground-dapr/common/model"
	"github.com/domahidizoltan/playground-dapr/transactionservice/ent"
	entGen "github.com/domahidizoltan/playground-dapr/transactionservice/ent/generated"
	"github.com/google/uuid"

	dapr "github.com/dapr/go-sdk/client"
)

type appModeType string

const (
	creditAppModeType appModeType = "credit"
	debitAppModeType  appModeType = "debit"
)

var (
	appMode                                           appModeType
	debitTopicName, creditTopicName, balanceTopicName string
	service                                           string

	entClient     *entGen.Client
	daprClient    dapr.Client
	subscriptions = map[appModeType]*common.Subscription{
		debitAppModeType: {
			PubsubName: client.PubsubName,
			Topic:      "",
			Route:      "/debitSourceCommand",
		},
		creditAppModeType: {
			PubsubName: client.PubsubName,
			Topic:      "",
			Route:      "/creditDestCommand",
		},
	}
)

func main() {
	if len(os.Args) < 2 {
		panic(errors.New("appMode is not defined"))
	}

	switch os.Args[1] {
	case string(debitAppModeType), string(creditAppModeType):
		appMode = appModeType(os.Args[1])
	default:
		panic(errors.New("unknown appMode"))
	}

	service = strings.ToUpper(string(appMode)) + "TNX"

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

	debitTopicName = helper.GetEnv(strings.ToUpper(service)+"TNX_TOPIC_DEBIT_TRANSACTION", "topic.debit_transaction")
	creditTopicName = helper.GetEnv(strings.ToUpper(service)+"TNX_TOPIC_CREDIT_TRANSACTION", "topic.credit_transaction")
	balanceTopicName = helper.GetEnv(strings.ToUpper(service)+"TNX_TOPIC_BALANCE", "topic.balance")
	subscriptions[creditAppModeType].Topic = creditTopicName
	subscriptions[debitAppModeType].Topic = debitTopicName

	subscriptions := []client.SubscriptionHandler{
		{
			Subscription: subscriptions[appMode],
			Handler:      commandHandler,
		},
	}
	client.SubscribeTopic(service, subscriptions, nil)
}

func commandHandler(ctx context.Context, event *common.TopicEvent) (bool, error) {
	log.Println("received", event)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var cmd *dto.TransferCommand
	if err := event.Struct(&cmd); err != nil {
		log.Printf("failed to parse TransferCommand for %s: %s", string(event.RawData), err)
		return true, err
	}

	switch cmd.Command {
	case dto.DebitSourceCommandType:
		if err := doTransaction(ctx, *cmd); err != nil {
			if errors.Is(err, dto.BusinessError) {
				return false, err
			}
			return true, err
		}

		return publishCommand(ctx, *cmd, dto.UpdateSourceBalanceCommandType)

	case dto.CreditDestCommandType:
		if err := doTransaction(ctx, *cmd); err != nil {
			if errors.Is(err, dto.BusinessError) {
				return false, err
			}
			return true, err
		}

		return publishCommand(ctx, *cmd, dto.UpdateDestBalanceCommandType)
	}

	log.Printf("unknown command type: %s", cmd.Command)
	return false, nil
}

func publishCommand(ctx context.Context, cmd dto.TransferCommand, newCommandType dto.CommandType) (bool, error) {
	newCmd := cmd
	newCmd.Command = newCommandType
	topicName := ""
	switch newCommandType {
	case dto.UpdateSourceBalanceCommandType, dto.UpdateDestBalanceCommandType:
		topicName = balanceTopicName
	}

	log.Printf("publish to %s command %+v", topicName, newCmd)
	if err := daprClient.PublishEvent(ctx, client.PubsubName, topicName, newCmd); err != nil {
		log.Printf("failed to publish command %v: %s", newCmd, err)
		return false, err
	}

	return false, nil
}

func doTransaction(ctx context.Context, transfer dto.TransferCommand) error {
	dest := transfer.DstAcc
	src := transfer.SrcAcc
	switch transfer.Command {
	case dto.DebitSourceCommandType:
		dest = model.BankAccount
	case dto.CreditDestCommandType:
		src = model.BankAccount
	}

	newTransaction, err := entClient.Transaction.Create().
		SetID(uuid.New()).
		SetTnx(transfer.Tnx).
		SetDestAcc(dest).
		SetSourceAcc(src).
		SetAmount(transfer.Amount).
		SetDatetime(time.Now().UTC()).
		Save(ctx)
	if err != nil {
		log.Printf("failed to save transaction for %v: %s", transfer, err)
		return err
	}

	log.Printf("saved transaction %+v", newTransaction)

	return nil
}
