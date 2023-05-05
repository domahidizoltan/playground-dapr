package client

import (
	"log"

	ent "github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated"
	"github.com/domahidizoltan/playground-dapr/common/helper"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriverKey     = "BALANCE_DB_DRIVER"
	dbDataSourceKey = "BALANCE_DB_DATASOURCE"
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
