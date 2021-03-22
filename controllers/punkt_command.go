package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgraph-io/badger/v2"
	"github.com/savsgio/atreugo/v11"
	"kurs.kz/paladin/cache"
	"kurs.kz/paladin/db"
	"kurs.kz/paladin/models"
)

/*
UpdatePunkt save punkt object to the budger store
Uses POST
*/

func UpdatePunktData(ctx *atreugo.RequestCtx) error {
	id := ctx.UserValue("id")
	var punkt models.Punkt
	var data models.Data
	err := data.Unmarshal(ctx.PostBody())
	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"message": fmt.Sprintf("%s", err),
			"status":  400,
		}, 400)
	}

	if id == nil {
		return ctx.JSONResponse(map[string]interface{}{
			"message": "ID not set",
			"status":  400,
		}, 400)
	}
	var key []byte

	key = []byte(fmt.Sprintf("punkt-%s", id))
	var value []byte
	err = db.DB.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		value, err = item.ValueCopy(nil)
		value, err = punkt.UnmarshalMsg(value)
		punkt.Data = data
		err = db.DB.Update(func(txn *badger.Txn) error {
			dataToSave, err := punkt.MarshalMsg(nil)
			if err != nil {
				return err
			}

			e := badger.NewEntry(key, dataToSave).WithTTL(time.Hour * 24)
			return txn.SetEntry(e)
		})
		return err
	})

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"message": "error in encoding to MSGP",
			"status":  400,
		}, 400)
	}
	// cache.PaladinCache.Remove("punkts")
	return ctx.JSONResponse(punkt)
}

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

		e := badger.NewEntry(key, data).WithTTL(time.Hour * 24)

		return txn.SetEntry(e)
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

		key = []byte(fmt.Sprintf("punkt-%d", punkt.ID))

		data, _ := punkt.MarshalMsg(nil)
		e := badger.NewEntry(key, data).WithTTL(time.Hour * 24)

		_ = wb.SetEntry(e)
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
