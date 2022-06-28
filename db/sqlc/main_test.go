package db

import (
	"database/sql"
	"fmt"
	"github.com/golang/simplebank/util"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	fmt.Println("testQueries:", testDB)
	fmt.Println("dasdasd", testQueries.db)
	os.Exit(m.Run())
}
