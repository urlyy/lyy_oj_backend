package util

import "github.com/jmoiron/sqlx"

var db *sqlx.DB
var cfg *MyConfig

func init() {
	cfg = initConfig()
	db = initDB(&cfg.Database)
}

func GetDB() *sqlx.DB {
	return db
}

func GetProjectConfig() *MyConfig {
	return cfg
}
