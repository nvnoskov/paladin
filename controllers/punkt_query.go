package controllers

import (
	"encoding/json"
	"fmt"

	"kurs.kz/paladin/cache"

	"github.com/dgraph-io/badger"
	"github.com/savsgio/atreugo"
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

	// var cache models.Cache
	// cacheKey := []byte("cache-punkts")
	var punkts []models.Punkt
	var err error
	// err := db.DB.View(func(txn *badger.Txn) error {
	// 	item, err := txn.Get(cacheKey)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	value, err := item.ValueCopy(nil)
	// 	value, err = cache.UnmarshalMsg(value)

	// 	return err
	// })
	// if err == nil {
	if cache.PaladinCache.Has("punkts") == true {
		// json.Unmarshal([]byte(cache.Value), &punkts)
		// for k := range _cachedPunkts {
		// 	_cachedPunkts[k].ActualTime = 0
		// }

		// return ctx.JSONResponse(_cachedPunkts)
		//[]byte("{ \"punkts\": ")[:], , []byte("\", \"delta\":1}")
		// test := []byte(", \"delta\":1}")
		// start := []byte("{ \"punkts\": ")
		// val := append(start, _cachedPunktsString...)
		// val = append(val, test...)
		ctx.SetContentType("application/json")
		return ctx.TextResponseBytes(cache.PaladinCache.Get("punkts"))
	}
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

	// jsonString, _ := json.Marshal(punkts)
	// cache.Value = string(jsonString)
	// err = db.DB.Update(func(txn *badger.Txn) error {

	// 	data, err := cache.MarshalMsg(nil)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	return txn.Set(cacheKey, data)
	// })
	// var value []byte
	// err := db.DB.View(func(txn *badger.Txn) error {

	// 	item, err := txn.Get([]byte(key))
	// 	if err != nil {
	// 		return err
	// 	}

	// 	value, err = item.ValueCopy(nil)
	// 	value, err = punkt.UnmarshalMsg(value)

	// 	fmt.Printf("value: %+v err: %+v", value, err)

	// 	return err
	// })

	if err != nil {
		return ctx.JSONResponse(map[string]interface{}{
			"message": fmt.Sprintf("%s", err),
			"status":  400,
		}, 400)
	}
	_cachedPunkts = punkts
	_cachedPunktsString, _ = json.Marshal(punkts)
	cache.PaladinCache.Set("punkts", _cachedPunktsString)
	return ctx.JSONResponse(punkts)
}
