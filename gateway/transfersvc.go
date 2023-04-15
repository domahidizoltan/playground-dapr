package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"

	"github.com/dapr/go-sdk/client"
	"github.com/domahidizoltan/playground-dapr/model"
	"github.com/google/uuid"
)

const (
	DefaultPubsubName     = "transferevents"
	pendingTransfersTopic = "pending-transfers"
)

type TransferSvc struct {
	dapr       client.Client
	pubsubName string
}

var (
	ErrInvalidPubsubName = errors.New("invalid pubsub name")
	ErrInvalidAccount    = errors.New("invalid account")
	ErrInvalidAmount     = errors.New("invalid amount")
)

func NewTransferSvc(pubsubName string) (*TransferSvc, error) {
	if len(pubsubName) == 0 {
		return nil, ErrInvalidPubsubName
	}

	client, err := client.NewClient()
	if err != nil {
		return nil, err
	}

	return &TransferSvc{
		dapr:       client,
		pubsubName: pubsubName,
	}, nil
}

func (s TransferSvc) InitTransfer(ctx context.Context, srcAcc, dstAcc string, amount float32) error {
	if !isValidateAccount(srcAcc) {
		return fmt.Errorf("%w: %s", ErrInvalidAccount, srcAcc)
	}
	if !isValidateAccount(dstAcc) {
		return fmt.Errorf("%w: %s", ErrInvalidAccount, dstAcc)
	}
	if amount <= 0 {
		return ErrInvalidAmount
	}

	tnx := uuid.New()
	t := model.Transfer{
		TNX:    tnx.String(),
		SrcAcc: srcAcc,
		DstAcc: dstAcc,
		Amount: amount,
	}
	log.Printf("request transfer: tnx=%s srcAcc=%s dstAcc=%s amount=%f", tnx, srcAcc, dstAcc, amount)
	return s.dapr.PublishEvent(ctx, s.pubsubName, pendingTransfersTopic, t)
}

func isValidateAccount(acc string) bool {
	isMatch, err := regexp.MatchString(model.AccPattern, acc)
	if err != nil {
		log.Printf("failed to validate account: %s", err)
		return false
	}
	return isMatch
}
