package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@localhost:4321/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	con, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln("cannot get connection: ", err)
	}

	testQueries = New(con)

	// m.Run
	os.Exit(m.Run())
}
