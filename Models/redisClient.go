package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-gorp/gorp"
	"github.com/go-redis/redis/v7"
	_ "github.com/lib/pq"
)

var db *gorp.DbMap
var _redisClient *redis.Clien

func InitDB() {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

func GetDB() *gorp.DbMap {
	return db
}

func InitRedis(selectDB int) {

	redisHost := os.Getenv("REDIS_HOST")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	_redisClient = redis.NewClient(&redis.Options
		Addr:     redisHost,
		Password: redisPassword,
		DB:       selectDB,
		// DialTimeout:        10 * time.Second,
		// ReadTimeout:        30 * time.Second,
		// WriteTimeout:       30 * time.Second,
		// PoolSize:           10,
		// PoolTimeout:        30 * time.Second,
		// IdleTimeout:        500 * time.Millisecond,
		// IdleCheckFrequency: 500 * time.Millisecond,
		// TLSConfig: &tls.Config{
		// 	 InsecureSkipVerify: true,
		// },
	})
}

func GetRedis() *redis.Client {
	return _redisClien
}
