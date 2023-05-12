package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/domahidizoltan/playground-dapr/common/dto"
	"github.com/domahidizoltan/playground-dapr/common/helper"
	"github.com/domahidizoltan/playground-dapr/common/model"
	"github.com/oklog/ulid/v2"
)

const (
	service    = "GATEWAY"
	pubsubName = service + "_PUBSUB"

	srcAccQuery = "srcAcc"
	dstAccQuery = "dstAcc"
	amountQuery = "amount"
)

var (
	balanceTopic      = helper.GetEnv("GATEWAY_TOPIC_BALANCE", "topic.balance")
	ErrInvalidAccount = errors.New("invalid account")
	ErrInvalidAmount  = errors.New("invalid amount")

	daprClient dapr.Client
)

func validateParams(srcAcc, dstAcc string, amount float64) error {
	if !isValidAccount(srcAcc) {
		return fmt.Errorf("%w: %s", ErrInvalidAccount, srcAcc)
	}

	if !isValidAccount(dstAcc) {
		return fmt.Errorf("%w: %s", ErrInvalidAccount, dstAcc)
	}

	if amount <= 0 {
		return ErrInvalidAmount
	}

	return nil
}

func isValidAccount(acc string) bool {
	isMatch, err := regexp.MatchString(model.AccPattern, acc)
	if err != nil {
		log.Printf("failed to validate account: %s", err)
		return false
	}
	return isMatch
}

func initTransferHandler(w http.ResponseWriter, r *http.Request) {
	params := map[string]string{}
	for _, key := range []string{srcAccQuery, dstAccQuery, amountQuery} {
		value, err := helper.GetQueryParam(r, key)
		if err != nil {
			helper.HttpError(w, http.StatusBadRequest, "failed to get query param", err)
			return
		}
		params[key] = value
	}

	amount, err := strconv.ParseFloat(params[amountQuery], 32)
	if err != nil {
		helper.HttpError(w, http.StatusBadRequest, "failed to parse amount", err)
		return
	}

	if err := validateParams(params[srcAccQuery], params[dstAccQuery], amount); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "invalid parameters", err)
		return
	}

	cmd := dto.TransferCommand{
		Command: dto.LockBalanceCommandType,
		Tnx:     ulid.Make().String(),
		SrcAcc:  params[srcAccQuery],
		DstAcc:  params[dstAccQuery],
		Amount:  amount,
	}

	if err := daprClient.PublishEvent(r.Context(), "balance", balanceTopic, cmd); err != nil {
		helper.HttpError(w, http.StatusBadRequest, "failed to publish command", err)
		return
	}

	log.Printf("request transfer %+v:", cmd)
	w.WriteHeader(http.StatusAccepted)
}

func initIncomingTransfers() {
	log.Println("init incoming transfers")
	//TODO
}

func finalizeCompletedTransfers() {
	log.Println("finalize completed transfers")
	//TODO
}

func main() {
	var err error
	daprClient, err = dapr.NewClient()
	if err != nil {
		log.Fatalf("failed to init DAPR client: %s", err)
	}
	defer daprClient.Close()

	isDevMode := len(os.Args) > 1 && os.Args[1] == "devMode"

	if isDevMode {
		http.HandleFunc("/inittransfer", initTransferHandler)

		address := helper.GetAddress(service, "3000")
		log.Printf("gateway listening on address %s", address)
		if err := http.ListenAndServe(address, nil); err != nil {
			panic(err)
		}
	} else {
		finalizeCompletedTransfers()
		initIncomingTransfers()
	}
}
