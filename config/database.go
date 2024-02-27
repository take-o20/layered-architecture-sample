package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Fatalf("Fatal Error in connect_tcp.go: %s environment variable not set.", k)
		}
		return v
	}

	var (
		dbUser    = mustGetenv("DB_USER")       // e.g. 'my-db-user'
		dbPwd     = mustGetenv("DB_PASS")       // e.g. 'my-db-password'
		dbName    = mustGetenv("DB_NAME")       // e.g. 'my-database'
		dbPort    = mustGetenv("DB_PORT")       // e.g. '3306'
		dbTCPHost = mustGetenv("INSTANCE_HOST") // e.g. '127.0.0.1' ('172.17.0.1' if deployed to GAE Flex)
	)

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPwd, dbTCPHost, dbPort, dbName)

	var dbOpenErr error
	DB, dbOpenErr = sql.Open("mysql", dbURI)
	if dbOpenErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", dbOpenErr)
		os.Exit(1)
	}
	// エラーハンドリング省略

	if err := DB.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	if err := migrateDB(DB); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func migrateDB(db *sql.DB) error {
	createVotes := `CREATE TABLE IF NOT EXISTS users (
		user_id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL
	);`
	_, err := db.Exec(createVotes)
	return err
}
