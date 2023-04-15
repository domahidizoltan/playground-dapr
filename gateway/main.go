package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	defaultPort   = 3000
	defaultHost   = "0.0.0.0"
	httpPortEnv   = "HTTP_PORT"
	httpHostEnv   = "HTTP_HOST"
	pubsubNameEnv = "PUBSUB_NAME"

	srcAccQuery = "srcAcc"
	dstAccQuery = "dstAcc"
	amountQuery = "amount"
)

var transferSvc *TransferSvc

func getQueryParam(r *http.Request, key string) (string, error) {
	value := r.URL.Query().Get(key)
	if len(value) == 0 {
		return "", fmt.Errorf("%s is required", key)
	}

	return value, nil
}

func initTransferHandler(w http.ResponseWriter, r *http.Request) {
	params := map[string]string{}
	for _, key := range []string{srcAccQuery, dstAccQuery, amountQuery} {
		value, err := getQueryParam(r, key)
		if err != nil {
			log.Printf("failed to get query param: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		params[key] = value
	}

	amount, err := strconv.ParseFloat(params[amountQuery], 32)
	if err != nil {
		log.Printf("failed to parse amount: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	if err := transferSvc.InitTransfer(context.Background(), params[srcAccQuery], params[dstAccQuery], float32(amount)); err != nil {
		log.Printf("failed to init transfer: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func getAddress() string {
	host := getEnv(httpHostEnv, defaultHost)
	p := getEnv(httpPortEnv, strconv.Itoa(defaultPort))
	port, err := strconv.Atoi(p)
	if err != nil {
		log.Printf("failed to parse port: %s", err)
	}

	return fmt.Sprintf("%s:%d", host, port)
}

func getEnv(key, defaultValue string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultValue
	}
	return val
}

func main() {
	pubsubName := getEnv(pubsubNameEnv, DefaultPubsubName)
	svc, err := NewTransferSvc(pubsubName)
	if err != nil {
		panic(err)
	}
	transferSvc = svc

	http.HandleFunc("/inittransfer", initTransferHandler)

	address := getAddress()
	log.Printf("gateway listening on address %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
	}
}
