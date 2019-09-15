package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/badger"
	"github.com/savsgio/atreugo"
	"kurs.kz/paladin/cache"
	"kurs.kz/paladin/db"
	"kurs.kz/paladin/models"
)

/*
UpdatePunkt save punkt object to the budger store
Uses POST
*/
func UpdatePunkt(ctx *atreugo.RequestCtx) error {
	id := ctx.UserValue("id")

	var punkt models.Punkt
	err := punkt.Unmarshal(ctx.PostBody())
	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"message": fmt.Sprintf("%s", err),
			"status":  400,
		}, 400)
	}

	if id == nil {
		id = punkt.ID
	}
	var key []byte

	key = []byte(fmt.Sprintf("punkt-%s", id))

	err = db.DB.Update(func(txn *badger.Txn) error {
		// item, err := txn.Get([]byte(key))
		// if err != nil {
		// 	return err
		// }
		data, err := punkt.MarshalMsg(nil)
		if err != nil {
			return err
		}

		return txn.Set(key, data)
	})

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"message": "error in encoding to MSGP",
			"status":  400,
		}, 400)
	}
	cache.PaladinCache.Remove("punkts")
	return ctx.JSONResponse(punkt)
}

/*
UpdatePunkt save punkt object to the budger store
Uses POST
*/
func SyncPunkts(ctx *atreugo.RequestCtx) error {

	var punkts []models.Punkt
	err := json.Unmarshal(ctx.PostBody(), &punkts)
	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"message": fmt.Sprintf("%s", err),
			"status":  400,
		}, 400)
	}

	var key []byte

	wb := db.DB.NewWriteBatch()
	defer wb.Cancel()

	for _, punkt := range punkts {

		// fmt.Printf("key[%d] value[%v]\n", k, punkt)
		key = []byte(fmt.Sprintf("punkt-%d", punkt.ID))

		data, _ := punkt.MarshalMsg(nil)

		_ = wb.Set(key, data) // Will create txns as needed.
	}

	wb.Flush() // Wait for all txns to finish.
	cache.PaladinCache.Remove("punkts")
	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"message": "error in encoding to MSGP",
			"status":  400,
		}, 400)
	}
	return ctx.JSONResponse(map[string]interface{}{
		"message": "saved",
		"status":  200,
	})
}
