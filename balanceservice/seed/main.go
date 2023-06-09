package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/domahidizoltan/playground-dapr/balanceservice/ent"
	entGen "github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated"

	_ "github.com/mattn/go-sqlite3"
)

const (
	defaultBalance = 1_000
)

func main() {
	log.Println("start balance seed")
	client := ent.GetClient("BALANCE")
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()

	c := client.Balance
	deleteAll(c, ctx)
	createAccounts(c, ctx)
}

func deleteAll(c *entGen.BalanceClient, ctx context.Context) {
	rows, err := c.Delete().Exec(ctx)
	if err != nil {
		log.Fatalf("failed to delete balances: %s", err.Error())
	}
	log.Printf("deleted %d balance records", rows)
}

func createAccounts(c *entGen.BalanceClient, ctx context.Context) {
	const balanceCount = 1000
	errorCount := 0
	for i := range [balanceCount]struct{}{} {
		id := fmt.Sprintf("%s%03d", "ACC", i)
		_, err := c.Create().
			SetID(id).
			SetBalance(defaultBalance).
			Save(ctx)
		if err != nil {
			errorCount++
			log.Fatalf("failed to create balance for %s: %s", id, err.Error())
		}
	}

	log.Printf("added %d balance records", balanceCount-errorCount)
}
