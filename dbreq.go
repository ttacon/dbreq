package dbreq

import "strings"

// TODO(ttacon): ✔

var db interface{}

func Init(dbConn interface{}) {
	db = dbConn
}

func Require(i Existing) error {
	ok, err := i.Exist(db)
	if err != nil && !strings.Contains(err.Error(), "doesn't exist") {
		return err
	}

	if ok {
		return nil
	}

	err = i.Create(db)
	return err
}

type Existing interface {
	Exist(db interface{}) (bool, error)
	Create(db interface{}) error
}
