package db

import (
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/sirupsen/logrus"
)

/*
DB link to badger database
*/
var DB *badger.DB

/*
OpenDatabase open connection to Badger Database
*/
func OpenDatabase(path string) error {
	opts := badger.DefaultOptions(path)
	var err error
	DB, err = badger.Open(opts)

	return err
}

/*
BadgerCleanupProc starts GC collector
*/
func BadgerCleanupProc() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for range ticker.C {
	again:
		logrus.Infof("calling DB.RunValueLogGC...")
		err := DB.RunValueLogGC(0.7)
		if err == nil {
			goto again
		}
	}
}
