package util

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func initDB(dbConfig *Database) *sqlx.DB {
	fmt.Println("db")
	params := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBname)
	db := sqlx.MustConnect("postgres", params)
	// 设置最大连接数，而且它内置了连接池
	// db.SetMaxOpenConns(dbConfig.MaxConnections)
	// db.SetMaxIdleConns(dbConfig.MaxConnections)
	return db
}
