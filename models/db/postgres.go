package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	"github.com/lib/pq"
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

// 用以執行複數 Query
// 第 returnIndex 筆 Query 回傳 Data
func ExecuteTransaction(queries []string, params map[int][]any, returnIndex int, holder interface{}, singleRow bool) error {

	txn, err := db.Begin()

	if err != nil {
		return err
	}

	defer func() {
		// Rollback the transaction after the function returns.
		// If the transaction was already commited, this will do nothing.
		_ = txn.Rollback()
	}()

	// 依序執行所有 Query
	for i, q := range queries {
		rows, err := txn.Query(q, params[i]...)
		defer rows.Close()

		if err != nil {
			return err
		}
		// 取得資料
		if i == returnIndex {
			if singleRow {
				if rows.Next() {

				}
			} else {
				// var tmpHolder interface{}
				// for rows.Next() {
				// 	if err := rows.Scan(&tmpHolder); err != nil {
				// 		return err
				// 	}
				// 	holder = append(holder, tmpHolder)
				// }
				// if err = rows.Err(); err != nil {
				// 	return err
				// }
			}
		}
	}

	// Commit the transaction.
	return txn.Commit()
}

func PrintDbError(_err error) {

	err, ok := _err.(*pq.Error)
	if !ok {
		return
	}

	fmt.Println("Database error:\n-----------------------")
	fmt.Println("Severity: ", err.Severity)
	fmt.Println("Code: ", err.Code)
	fmt.Println("Message: ", err.Message)
	fmt.Println("Detail: ", err.Detail)
	fmt.Println("Hint: ", err.Hint)
	fmt.Println("Position: ", err.Position)
	fmt.Println("InternalPosition: ", err.InternalPosition)
	fmt.Println("Where: ", err.Where)
	fmt.Println("Schema: ", err.Schema)
	fmt.Println("Table: ", err.Table)
	fmt.Println("Column: ", err.Column)
	fmt.Println("DataTypeName: ", err.DataTypeName)
	fmt.Println("Constraint: ", err.Constraint)
	fmt.Println("File: ", err.File)
	fmt.Println("Line: ", err.Line)
	fmt.Println("Routine: ", err.Routine)
}
