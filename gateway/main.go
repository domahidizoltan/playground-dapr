package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	dapr "github.com/dapr/go-sdk/client"
	"github.com/dapr/go-sdk/service/common"
	"github.com/domahidizoltan/playground-dapr/common/client"
	"github.com/domahidizoltan/playground-dapr/common/dto"
	"github.com/domahidizoltan/playground-dapr/common/helper"
	"github.com/domahidizoltan/playground-dapr/common/model"
	"github.com/oklog/ulid/v2"
)

const (
	service                      = "GATEWAY"
	configStore                  = "configstore"
	minTransferAmountKey         = "min_transfer_amount"
	transferFilesPath            = "transferfiles"
	scheduledTransfersPathFormat = transferFilesPath + "/newtransfers_%s.csv"
	completedTransfersPathFormat = transferFilesPath + "/completedtransfers_%s.csv"

	srcAccQuery = "srcAcc"
	dstAccQuery = "dstAcc"
	amountQuery = "amount"
)

var (
	balanceTopic      = helper.GetEnv(service+"_TOPIC_BALANCE", "topic.balance")
	completedTnxTopic = helper.GetEnv(service+"_TOPIC_COMPLETED_TRANSACTION", "topic.completed_transaction")
	ErrInvalidAccount = errors.New("invalid account")
	ErrInvalidAmount  = errors.New("invalid amount")

	daprClient        dapr.Client
	minTransferAmount = float64(10)
)

var completedTnxSub = &common.Subscription{
	PubsubName: client.PubsubName,
	Topic:      completedTnxTopic,
	Route:      "/completedTransaction",
}

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

	if err := initTransfer(r.Context(), params[srcAccQuery], params[dstAccQuery], params[amountQuery]); err != nil {
		e := errors.Unwrap(err)
		helper.HttpError(w, http.StatusBadRequest, err.Error(), e)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func initTransfer(ctx context.Context, srcAcc, dstAcc, amount string) error {
	amt, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		return fmt.Errorf("failed to parse amount %w", err)
	}

	if amt < minTransferAmount {
		return fmt.Errorf("amount must be min %.f", minTransferAmount)
	}

	if err := validateParams(srcAcc, dstAcc, amt); err != nil {
		return fmt.Errorf("invalid parameters %w", err)
	}

	cmd := dto.TransferCommand{
		Command: dto.LockBalanceCommandType,
		Tnx:     ulid.Make().String(),
		SrcAcc:  srcAcc,
		DstAcc:  dstAcc,
		Amount:  amt,
	}

	if err := daprClient.PublishEvent(ctx, client.PubsubName, balanceTopic, cmd); err != nil {
		return fmt.Errorf("failed to publish command %w", err)
	}

	log.Printf("request transfer %+v:", cmd)
	return nil
}

func checkNewTransfers(w http.ResponseWriter, r *http.Request) {
	go processScheduledTransfers()
	if w != nil {
		w.WriteHeader(http.StatusOK)
	}
}

func processScheduledTransfers() {
	path := fmt.Sprintf(scheduledTransfersPathFormat, time.Now().UTC().Format("06010215"))
	log.Println("process sheduled transfers from", path)

	file, err := os.Open(path)
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("failed to close file %s: %s", path, err.Error())
		}
	}()

	if err != nil {
		log.Printf("failed to read file %s: %s", path, err.Error())
		return
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Println("failed to read records", err.Error())
		return
	}

	for _, record := range records[1:] {
		if err := initTransfer(context.TODO(), record[0], record[1], record[2]); err != nil {
			log.Printf("failed to init transfer for record %+v: %+s", record, err.Error())
		}
	}
}

func completedTransfersHandler(ctx context.Context, event *common.TopicEvent) (bool, error) {
	log.Println("received", event)

	var cmd *dto.TransferCommand
	if err := event.Struct(&cmd); err != nil {
		log.Printf("failed to parse TransferCommand for %s: %s", string(event.RawData), err)
		return true, err
	}

	path := fmt.Sprintf(completedTransfersPathFormat, time.Now().UTC().Format("06010215"))
	file, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.WriteFile(path, []byte("source_acc,dest_acc,amount,tnx,completed_at\n"), 0777); err != nil {
				log.Printf("failed to create file %s: %s", path, err.Error())
				return true, err
			}
		} else {
			log.Printf("failed to open file %s: %s", path, err)
			return true, err
		}
	}

	record := []string{cmd.SrcAcc, cmd.DstAcc, fmt.Sprintf("%.2f", cmd.Amount), cmd.Tnx, time.Now().UTC().Format(time.RFC3339)}
	log.Println("writing record", record)
	writer := csv.NewWriter(file)
	if err := writer.Write(record); err != nil {
		log.Printf("failed to write record %s: %s", record, err.Error())
		return true, err
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Printf("failed to flush data: %s", err.Error())
		return true, err
	}

	return false, nil
}

func main() {
	var err error
	daprClient, err = dapr.NewClient()
	if err != nil {
		log.Fatalf("failed to init DAPR client: %s", err)
	}
	defer daprClient.Close()

	go setConfigs()

	subscriptions := []client.SubscriptionHandler{
		{
			Subscription: completedTnxSub,
			Handler:      completedTransfersHandler,
		},
	}

	serviceHook := func(s common.Service) {
		s.AddBindingInvocationHandler("checknewtransfers", func(ctx context.Context, in *common.BindingEvent) (out []byte, err error) {
			checkNewTransfers(nil, nil)
			return []byte("{\"status\":200}"), nil
		})
	}
	go client.SubscribeTopic(service, subscriptions, serviceHook)

	isDevMode := len(os.Args) > 1 && os.Args[1] == "devMode"
	if isDevMode {
		http.HandleFunc("/inittransfer", initTransferHandler)
	}

	http.HandleFunc("/checknewtransfers", checkNewTransfers)

	address := helper.GetAddress(service, "3000")
	log.Printf("gateway listening on address %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
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
