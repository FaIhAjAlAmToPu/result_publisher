package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	serviceURI := "postgres://avnadmin:AVNS_ro9OcBGvaBclv4BGw-V@pg-2500cc00-faihajalamtopu64-3ff9.e.aivencloud.com:15245/defaultdb?sslmode=require"

	conn, _ := url.Parse(serviceURI)
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"

	DB, err := sql.Open("postgres", conn.String())

	if err != nil {
		log.Fatal(err)
	}
	// defer DB.Close()

	rows, err := DB.Query("SELECT version()")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var result string
		err = rows.Scan(&result)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Version: %s\n", result)
	}
	return DB
}
