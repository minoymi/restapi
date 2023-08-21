package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

const url string = "postgres://postgres:1@localhost:5432/restapiDB"
const GET_statsQuery string = "SELECT json_object_agg(code, visitors) FROM countries"
const POST_incrementStatsQuery string = "INSERT INTO countries (code, visitors) VALUES ('%s', 1) ON CONFLICT (code) DO UPDATE SET visitors = countries.visitors + 1"

var dbpool *pgxpool.Pool

func DB_connect() {
	var err error
	dbpool, err = pgxpool.New(context.Background(), url)
	if err != nil {
		fmt.Println("Unable to create connection pool:", err)
		os.Exit(1)
	}
}

func DB_disconnect() {
	dbpool.Close()
}

func DatabaseRead() []byte {
	var jsonStats []byte
	row := dbpool.QueryRow(context.Background(), GET_statsQuery)
	row.Scan(&jsonStats)
	return jsonStats
}

func DatabaseWrite(country_code []byte) []byte {
	var dbReply string
	row := dbpool.QueryRow(context.Background(), fmt.Sprintf(POST_incrementStatsQuery, string(country_code)))
	err := row.Scan(&dbReply)
	if err != nil {
		return []byte("Added one visitor to " + string(country_code))
	} else {
		return []byte("Error")
	}
}
