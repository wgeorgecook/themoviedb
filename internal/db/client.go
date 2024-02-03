package db

import (
	"database/sql"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var client *bun.DB

func Init(uri string, debug bool) {
	log.Printf("inititalizing DB: %s", uri)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(uri)))
	client = bun.NewDB(sqldb, pgdialect.New())
	if debug {
		client.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}
}
