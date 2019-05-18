package infra

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySQL
)

// OpenRDB creates a new RDB client.
func OpenRDB() (*sql.DB, error) {
	return sql.Open("mysql", "root@tcp/goop?parseTime=true")
}
