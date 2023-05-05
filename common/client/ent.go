package client

import (
	"github.com/domahidizoltan/playground-dapr/common/helper"
)

const (
	dbDriverKey     = "_DB_DRIVER"
	dbDataSourceKey = "_DB_DATASOURCE"
)

func GetDBOpenParams(prefix string) (string, string) {
	dbDriver := helper.GetEnv(prefix+dbDriverKey, "sqlite3")
	dbDataSource := helper.GetEnv(prefix+dbDataSourceKey, "file:localdev/sqlitedata/main.db")
	return dbDriver, dbDataSource
}
