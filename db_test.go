package main

import (
	"fmt"
	"testing"

	ddb "github.com/peterhellberg/wiki/db"
)

func TestDB(t *testing.T) {
	db := &ddb.DB{}

	if err := db.Open("./test.db", 0600); err != nil {
		t.Error("open error")
	}
	db.Update(func(tx *ddb.Tx) error {
		err := tx.Bucket([]byte("pages")).Put([]byte("k1"), []byte("v1"))
		err = tx.Bucket([]byte("pages")).Put([]byte("k2"), []byte("v2"))
		return err
	})

	db.View(func(tx *ddb.Tx) error {
		b := tx.Bucket([]byte("pages"))
		v := b.Get([]byte("k1"))
		t.Log("get value: ", v, string(v))
		return nil
	})
	db.View(func(tx *ddb.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("pages"))

		c := b.Cursor()

		keys := make([]string, 0)
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("key=%s, value=%s\n", k, v)
			keys = append(keys, string(k))
		}
		fmt.Print("keys: ", keys, len(keys))
		return nil
	})

}
