package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/domahidizoltan/playground-dapr/common/helper"
)

type Balance struct {
	ID      string  `json:"id"`
	Balance float32 `json:"balance"`
}

func listBalances(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		helper.HttpError(w, http.StatusInternalServerError, "failed to parse query params", err)
		return
	}

	response := []Balance{}
	for i, acc := range r.Form["account"] {
		response = append(response, Balance{
			ID:      acc,
			Balance: float32(i),
		})
	}

	resp, err := json.Marshal(response)
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

	var balance *Balance
	if err := json.Unmarshal(body, &balance); err != nil {
		helper.HttpError(w, http.StatusInternalServerError, "failed to parse body", err)
		return
	}

	if id != balance.ID {
		helper.HttpError(w, http.StatusBadRequest, "id must be the same", nil)
	}

	log.Printf("updated balance %+v", balance)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/balances/", updateBalance)
	http.HandleFunc("/balances", listBalances)

	address := helper.GetAddress("BALANCE", "3000")
	log.Printf("balance service listening on address %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		panic(err)
	}
}
