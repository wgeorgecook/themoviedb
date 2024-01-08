package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var client *bun.DB

func Init(uri string) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(uri)))
	client = bun.NewDB(sqldb, pgdialect.New())
}
