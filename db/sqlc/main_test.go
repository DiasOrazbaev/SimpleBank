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
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatalln("cannot get connection: ", err)
	}

	testQueries = New(testDb)

	// m.Run
	stauts := m.Run()
	err = testDb.Close()
	if err != nil {
		panic("not close:")
	}
	os.Exit(stauts)
}
