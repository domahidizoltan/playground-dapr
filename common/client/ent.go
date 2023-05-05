package client

import (
	"log"

	ent "github.com/domahidizoltan/playground-dapr/balanceservice/ent/generated"
	"github.com/domahidizoltan/playground-dapr/common/helper"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbDriverKey     = "_DB_DRIVER"
	dbDataSourceKey = "_DB_DATASOURCE"
)

func GetEntClient(prefix string) *ent.Client {
	dbDriver := helper.GetEnv(prefix+dbDriverKey, "sqlite3")
	dbDataSource := helper.GetEnv(prefix+dbDataSourceKey, "file:localdev/sqlitedata/main.db")

	client, err := ent.Open(dbDriver, dbDataSource)
	if err != nil {
		log.Fatalf("failed opening database connection: %v", err)
	}
	return client
}
