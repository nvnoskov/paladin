package controllers

import (
	// "encoding/json"
	"fmt"

	// "kurs.kz/paladin/cache"

	"github.com/dgraph-io/badger/v2"
	"github.com/savsgio/atreugo/v11"
	"kurs.kz/paladin/db"
	"kurs.kz/paladin/models"
)

var _cachedPunkts []models.Punkt
var _cachedPunktsString []byte

/*
GetPunkt save punkt object to the budger store
*/
func GetPunkt(ctx *atreugo.RequestCtx) error {
	id := ctx.UserValue("id")
	var punkt models.Punkt

	if id == nil {
		return ctx.JSONResponse(map[string]interface{}{
			"message": "id must be defined",
			"status":  400,
		}, 400)
	}
	var key []byte

	key = []byte(fmt.Sprintf("punkt-%s", id))
	var value []byte
	err := db.DB.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		value, err = item.ValueCopy(nil)
		value, err = punkt.UnmarshalMsg(value)

		return err
	})

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"message": fmt.Sprintf("%s", err),
			"status":  400,
		}, 400)
	}

	return ctx.JSONResponse(punkt)
}

func GetPunkts(ctx *atreugo.RequestCtx) error {

	var punkts []models.Punkt
	var err error

	var punkt models.Punkt
	err = db.DB.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("punkt-")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				punkt.UnmarshalMsg(v)
				punkts = append(punkts, punkt)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"message": fmt.Sprintf("%s", err),
			"status":  400,
		}, 400)
	}

	return ctx.JSONResponse(punkts)
}
