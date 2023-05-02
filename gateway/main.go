package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/domahidizoltan/playground-dapr/common/helper"
)

const (
	pubsubNameEnv = "PUBSUB_NAME"

	srcAccQuery = "srcAcc"
	dstAccQuery = "dstAcc"
	amountQuery = "amount"
)

var transferSvc *TransferSvc

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
	if err := transferSvc.InitTransfer(context.Background(), params[srcAccQuery], params[dstAccQuery], float32(amount)); err != nil {
		helper.HttpError(w, http.StatusInternalServerError, "failed to init transfer", err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func main() {
	pubsubName := helper.GetEnv(pubsubNameEnv, DefaultPubsubName)
	svc, err := NewTransferSvc(pubsubName)
	if err != nil {
		panic(err)
	}
	transferSvc = svc

	http.HandleFunc("/inittransfer", initTransferHandler)

	address := helper.GetAddress("GATEWAY", "3000")
	log.Printf("gateway listening on address %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
	}
}
