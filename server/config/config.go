package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
