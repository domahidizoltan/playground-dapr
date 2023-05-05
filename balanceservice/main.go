package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/domahidizoltan/playground-dapr/balanceservice/client"
	ent "github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated"
	"github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated/balance"
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

func updateBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParams := strings.Split(r.URL.Path, "/")
	if len(pathParams) != 3 {
		helper.HttpError(w, http.StatusBadRequest, "id is missing", nil)
		return
	}
	id := pathParams[2]

	defer func() {
		if err := r.Body.Close(); err != nil {
			helper.HttpError(w, http.StatusInternalServerError, "failed to close body", err)
			return
		}
	}()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		helper.HttpError(w, http.StatusInternalServerError, "failed to read body", err)
		return
	}

	var balance *ent.Balance
	if err := json.Unmarshal(body, &balance); err != nil {
		helper.HttpError(w, http.StatusInternalServerError, "failed to parse body", err)
		return
	}

	if id != balance.ID {
		helper.HttpError(w, http.StatusBadRequest, "id must be the same", nil)
	}

	if _, err := entClient.Balance.
		UpdateOneID(balance.ID).
		SetBalance(balance.Balance).
		Save(r.Context()); err != nil {
		helper.HttpError(w, http.StatusInternalServerError, "failed to update balance", err)
		return
	}

	log.Printf("updated balance %+v", balance)
	w.WriteHeader(http.StatusOK)
}

func router(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(r.URL.Path, "/balances/"):
		updateBalance(w, r)
	case strings.HasPrefix(r.URL.Path, "/balances"):
		listBalances(w, r)
	}
}

func main() {
	entClient = client.GetEntClient()
	if entClient == nil {
		panic(errors.New("failed to create database connection"))
	}

	address := helper.GetAddress("BALANCE", "3000")
	srv := &http.Server{
		Addr:              address,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      2 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		Handler:           http.HandlerFunc(router),
	}

	http.HandleFunc("/balances/", updateBalance)
	http.HandleFunc("/balances", listBalances)

	log.Printf("balance service listening on address %s", address)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
