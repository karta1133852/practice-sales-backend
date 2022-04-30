package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
)

var dbMap *gorp.DbMap
var db *sql.DB

func InitDB() {

	dbinfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL"),
	)

	var err error
	db, err = sql.Open("postgres", dbinfo)
	//db, err := sql.Open("postgres", os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	dbMap = &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
}

func GetDB() *gorp.DbMap {
	return dbMap
}

func CloseDB() error {
	return dbMap.Db.Close()
}
