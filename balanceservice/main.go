package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/dapr/go-sdk/service/common"
	ent "github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated"
	"github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated/balance"
	"github.com/domahidizoltan/playground-dapr/common/client"
	"github.com/domahidizoltan/playground-dapr/common/dto"
	"github.com/domahidizoltan/playground-dapr/common/helper"
)

var entClient *ent.Client

func listBalances(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		helper.HttpError(w, http.StatusInternalServerError, "failed to parse query params", err)
		return
	}

	accounts := make([]any, len(r.Form["account"]))
	for _, acc := range r.Form["account"] {
		accounts = append(accounts, acc)
	}
	balances, err := entClient.Balance.Query().Where(func(s *sql.Selector) {
		s.Where(sql.In(balance.FieldID, accounts...))
	}).All(r.Context())
	if err != nil {
		helper.HttpError(w, http.StatusInternalServerError, "failed to get account balances", err)
		return
	}

	resp, err := json.Marshal(balances)
	if err != nil {
		helper.HttpError(w, http.StatusInternalServerError, "failed to marshall response", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(resp); err != nil {
		helper.HttpError(w, http.StatusInternalServerError, "failed to write response", err)
		return
	}

}

func updateBalance(ctx context.Context, event *common.TopicEvent) (bool, error) {
	var dto *dto.UpdateBalance
	if err := event.Struct(&dto); err != nil {
		log.Printf("failed to parse UpdateBalance event for %s: %s", string(event.RawData), err)
		return true, err
	}

	amount := float64(dto.Amount)
	update := entClient.Balance.
		UpdateOneID(dto.Account).
		AddBalance(amount)
	if amount < 0 {
		update = update.Where(
			balance.PendingGTE(-amount),
		).AddPending(amount)
	}
	newBalance, err := update.Save(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			log.Printf("account not found or can't update balance: %+v", dto)
			return false, err
		}

		log.Printf("failed to update balance for %s: %s", dto.Account, err)
		return true, err
	}

	log.Printf("updated balance: %v", newBalance)
	return false, nil
}

func router(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/balances"):
		listBalances(w, r)
	}
}

const (
	service         = "BALANCE"
	defaultHttpPort = "3000"
)

func main() {
	entClient = client.GetEntClient(service)
	if entClient == nil {
		panic(errors.New("failed to create database connection"))
	}

	sub := &common.Subscription{
		PubsubName: "updatebalance",
		Topic:      "balance",
		Route:      "/updatebalance",
	}
	go client.SubscribeTopic(service, sub, updateBalance)

	address := helper.GetAddress(service, defaultHttpPort)
	srv := &http.Server{
		Addr:              address,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           http.HandlerFunc(router),
	}

	http.HandleFunc("/balances", listBalances)

	log.Printf("balance service listening on address %s", address)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
