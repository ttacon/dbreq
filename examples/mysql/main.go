package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ttacon/grog"
)

var (
	dsn = flag.String("dsn", "grog:grog@/grog", "the dsn to connect to your db with")
)

func main() {
	flag.Parse()

	// connect to our db
	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		fmt.Println("whoopsy! we couldn't connect to the db with that info!")
		return
	}

	grog.Init(db)
	err = grog.Require(Epic{})
	if err != nil {
		fmt.Println("whoops! couldn't make sure table exists due to err: ", err)
	} else {
		fmt.Println("made sure tables exists as specified!")
	}
}

// These should really be in their own `models` file (or something similar)
type Epic struct {
	HowEpic int
	TooEpic bool
}

func (e Epic) Exist(dble interface{}) (bool, error) {
	// we know it's a sql database connection
	db, _ := dble.(*sql.DB)

	// we just want to know if the table exists at all
	_, err := db.Query("describe epic")
	return err != nil, err
}

func (e Epic) Create(dble interface{}) error {
	// we know it's a sql database connection
	db, _ := dble.(*sql.DB)

	_, err := db.Exec(`
create table Epic (
  HowEpic int,
  TooEpic bool
);
`)
	return err
}
