package ent

import (
	"log"

	"github.com/domahidizoltan/playground-dapr/common/client"
	entGen "github.com/domahidizoltan/playground-dapr/transactionservice/ent/generated"

	_ "github.com/mattn/go-sqlite3"
)

func GetClient(prefix string) *entGen.Client {
	dbDriver, dbDataSource := client.GetDBOpenParams(prefix)

	client, err := entGen.Open(dbDriver, dbDataSource)
	if err != nil {
		log.Fatalf("failed opening database connection: %v", err)
	}
	return client
}
