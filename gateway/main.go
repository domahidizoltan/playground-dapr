package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/domahidizoltan/playground-dapr/common/client"
	"github.com/domahidizoltan/playground-dapr/common/dto"
	"github.com/domahidizoltan/playground-dapr/common/helper"
	"github.com/domahidizoltan/playground-dapr/common/model"
	"github.com/oklog/ulid/v2"
)

const (
	service              = "GATEWAY"
	configStore          = "configstore"
	minTransferAmountKey = "min_transfer_amount"

	srcAccQuery = "srcAcc"
	dstAccQuery = "dstAcc"
	amountQuery = "amount"
)

var (
	balanceTopic      = helper.GetEnv("GATEWAY_TOPIC_BALANCE", "topic.balance")
	ErrInvalidAccount = errors.New("invalid account")
	ErrInvalidAmount  = errors.New("invalid amount")

	daprClient        dapr.Client
	minTransferAmount = float64(10)
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

	if amount < minTransferAmount {
		helper.HttpError(w, http.StatusBadRequest, "", fmt.Errorf("amount must be min %.f", minTransferAmount))
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

	if err := daprClient.PublishEvent(r.Context(), client.PubsubName, balanceTopic, cmd); err != nil {
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

	go setConfigs()

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

func setConfigs() {
	updateMinTransferAmount := func(cfg *dapr.ConfigurationItem) {
		var err error
		minTransferAmount, err = strconv.ParseFloat(cfg.Value, 64)
		if err != nil {
			log.Printf("failed to update config %s: %s", minTransferAmountKey, err)
		}

		log.Printf("updated config %s to %f", minTransferAmountKey, minTransferAmount)
	}

	cfg, err := daprClient.GetConfigurationItem(context.TODO(), configStore, minTransferAmountKey)
	if err != nil {
		log.Println("failed to get config", err)
	}
	updateMinTransferAmount(cfg)

	if err := daprClient.SubscribeConfigurationItems(context.TODO(), configStore, []string{minTransferAmountKey}, func(s string, m map[string]*dapr.ConfigurationItem) {
		cfg, ok := m[minTransferAmountKey]
		if !ok {
			log.Printf("%s config not found, using the current value %.f", minTransferAmountKey, minTransferAmount)
			return
		}
		updateMinTransferAmount(cfg)
	}); err != nil {
		log.Println("failed to get config", err)
	}

}
