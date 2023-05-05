package ent

import (
	"log"

	entGen "github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated"
	"github.com/domahidizoltan/playground-dapr/common/client"

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
