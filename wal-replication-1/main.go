package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dropQuery   = `DROP TABLE IF EXISTS mytable`
	createQuery = `
CREATE TABLE IF NOT EXISTS mytable (
	name string not null,
	output text not null
)
`
	insertQuery  = `INSERT INTO mytable (name, output) values (?, ?)`
	selectQuery  = `SELECT COUNT(*) FROM mytable WHERE name LIKE '%2%'`
	producerTick = time.Second / 1000
	consumerTick = time.Second / 1000
)

var pathMap = map[string]string{
	"producer": "file:./dbs-primary/state.db?_journal=WAL",
	"consumer": "file:./dbs-replica/state.db?_journal=WAL",
}

var (
	dropTable = flag.Bool("z", false, "Recreate tables")
	blobSize  = flag.Int("s", 2000000, "Recreate tables")
)

func main() {
	flag.Parse()
	var err error
	switch flag.Arg(0) {
	case "producer":
		err = runPrimary()
	case "consumer":
		err = runReplica()
	default:
		fmt.Printf("Usage:\n go run . [-z] <producer|consumer>\n")
		return
	}
	log.Fatal(err)
}

func openDb(role string) (*sql.DB, error) {
	path, ok := pathMap[role]
	if !ok {
		return nil, fmt.Errorf("Role %s not found", role)
	}
	return sql.Open("sqlite3", path)
}

func runPrimary() error {
	db, err := openDb("producer")
	if err != nil {
		return err
	}
	db.Ping()

	if *dropTable {
		if _, err := db.Exec(dropQuery); err != nil {
			panic(err)
		}
	}

	if _, err := db.Exec(createQuery); err != nil {
		panic(err)
	}

	for t := range time.Tick(producerTick) {
		size := rand.Intn(*blobSize)
		blob := randString(size)
		_, err := db.Exec(insertQuery, t.String(), blob)
		if err != nil {
			return err
		}
	}
	return nil
}

func runReplica() error {
	db, err := openDb("consumer")
	if err != nil {
		return err
	}
	db.Ping()

	var count int
	for t := range time.Tick(consumerTick) {
		if err := db.QueryRow(selectQuery).Scan(&count); err != nil {
			return err
		}
		fmt.Println(t.String(), count)
	}
	return nil
}
