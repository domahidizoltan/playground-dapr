package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dapr/go-sdk/service/common"
	"github.com/domahidizoltan/playground-dapr/balanceservice/ent"
	entGen "github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated"
	"github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated/balance"
	"github.com/domahidizoltan/playground-dapr/common/client"
	"github.com/domahidizoltan/playground-dapr/common/dto"
	"github.com/domahidizoltan/playground-dapr/common/helper"

	dapr "github.com/dapr/go-sdk/client"
)

const (
	service    = "BALANCE"
	pubsubName = service + "_PUBSUB"
)

var (
	entClient  entGen.Client
	daprClient dapr.Client

	balanceTopicName   = helper.GetEnv(service+"_TOPIC_BALANCE", "topic.balance")
	creditTnxTopicName = helper.GetEnv(service+"_TOPIC_CREDIT_TRANSACTION", "topic.credit_transaction")
	debitTnxTopicName  = helper.GetEnv(service+"_TOPIC_DEBIT_TRANSACTION", "topic.debit_transaction")
)

func lockBalance(ctx context.Context, account string, amount float64) error {
	acc, err := entClient.Balance.Get(ctx, account)
	if err != nil {
		if entGen.IsNotFound(err) {
			log.Printf("account %s not found: %s", account, err)
			return fmt.Errorf("%w: account not found: %s", dto.BusinessError, err)
		}
		return err
	}

	if acc.Balance < amount {
		log.Printf("insufficient balance for account %s: %s", account, err)
		return fmt.Errorf("%w: insufficient balance", dto.BusinessError)
	}

	if _, err := entClient.Balance.
		UpdateOneID(acc.ID).
		AddLocked(amount).
		Save(ctx); err != nil {
		log.Printf("failed to lock balance for account %s: %s", account, err)
		return err
	}

	return nil
}

func updateBalance(ctx context.Context, account string, amount float64) error {
	update := entClient.Balance.
		UpdateOneID(account).
		AddBalance(amount)
	if amount < 0 {
		update = update.Where(
			balance.LockedGTE(-amount),
		).AddLocked(amount)
	}
	newBalance, err := update.Save(ctx)
	if err != nil {
		if entGen.IsNotFound(err) {
			log.Printf("account not found or can't update balance %s: %s", account, err)
			return fmt.Errorf("%w: account not found or can't update balance: %s", dto.BusinessError, err)
		}

		log.Printf("failed to update balance for %s: %s", account, err)
		return err
	}

	log.Printf("updated balance: %v", newBalance)
	return nil
}

func publishCommand(ctx context.Context, cmd dto.TransferCommand, newCommandType dto.CommandType) (bool, error) {
	newCmd := cmd
	newCmd.Command = newCommandType
	topicName := ""
	switch newCommandType {
	case dto.DebitSourceCommandType:
		topicName = creditTnxTopicName
	case dto.CreditDestCommandType:
		topicName = debitTnxTopicName
	}
	if err := daprClient.PublishEvent(ctx, pubsubName, topicName, newCmd); err != nil {
		log.Printf("failed to publish command %v: %s", newCmd, err)
		return false, err
	}

	return false, nil
}

func commandHandler(ctx context.Context, event *common.TopicEvent) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	var cmd *dto.TransferCommand
	if err := event.Struct(&cmd); err != nil {
		log.Printf("failed to parse TransferCommand for %s: %s", string(event.RawData), err)
		return true, err
	}

	switch cmd.Command {
	case dto.LockBalanceCommandType:
		if err := lockBalance(ctx, cmd.SrcAcc, cmd.Amount); err != nil {
			if errors.Is(err, dto.BusinessError) {
				return false, err
			}
			return true, err
		}

		return publishCommand(ctx, *cmd, dto.DebitSourceCommandType)

	case dto.UpdateSourceBalanceCommandType:
		if err := updateBalance(ctx, cmd.SrcAcc, -cmd.Amount); err != nil {
			if errors.Is(err, dto.BusinessError) {
				return false, err
			}
			return true, err
		}

		return publishCommand(ctx, *cmd, dto.CreditDestCommandType)

	case dto.UpdateDestBalanceCommandType:
		if err := updateBalance(ctx, cmd.DstAcc, cmd.Amount); err != nil {
			if errors.Is(err, dto.BusinessError) {
				return false, err
			}
			return true, err
		}

	}

	log.Printf("unknown command type: %s", cmd.Command)
	return false, nil
}

var balanceSub = &common.Subscription{
	PubsubName: "balance",
	Topic:      balanceTopicName,
	Route:      "/balanceCommand",
}

func main() {
	ec := ent.GetClient(service)
	if ec == nil {
		panic(errors.New("failed to create database connection"))
	}
	entClient = *ec

	daprClient, err := dapr.NewClient()
	if err != nil {
		log.Fatalf("failed to init DAPR client: %s", err)
	}
	defer daprClient.Close()

	subscriptions := []client.SubscriptionHandler{
		{
			Subscription: balanceSub,
			Handler:      commandHandler,
		},
	}
	client.SubscribeTopic(service, subscriptions)
}
