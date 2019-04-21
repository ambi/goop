package infra

import (
	"github.com/ambi/goop/adapter/mysql"
	"github.com/ambi/goop/domain/db"
)

// ConnectRDB は RDB に接続する。
func ConnectRDB() (db.DAO, error) {
	return mysql.NewDAO("root@tcp/goop?parseTime=true")
}
