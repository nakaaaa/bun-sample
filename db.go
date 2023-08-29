package bunsample

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*bun.DB
}

func (db *DB) Open() error {
	dsn := os.Getenv("DSN")
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	db.DB = bun.NewDB(sqldb, mysqldialect.New())

	return nil
}
