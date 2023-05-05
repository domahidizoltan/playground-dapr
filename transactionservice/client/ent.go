package client

import (
	"log"

	"github.com/domahidizoltan/playground-dapr/common/helper"
	ent "github.com/domahidizoltan/playground-dapr/transactionservice/ent/generated"

	_ "github.com/mattn/go-sqlite3"
)

const (
	envPrefix       = "TRANSACTION"
	dbDriverKey     = envPrefix + "_DB_DRIVER"
	dbDataSourceKey = envPrefix + "_DB_DATASOURCE"
)

func GetEntClient() *ent.Client {
	dbDriver := helper.GetEnv(dbDriverKey, "sqlite3")
	dbDataSource := helper.GetEnv(dbDataSourceKey, "file:localdev/sqlitedata/main.db")

	client, err := ent.Open(dbDriver, dbDataSource)
	if err != nil {
		log.Fatalf("failed opening database connection: %v", err)
	}
	return client
}
