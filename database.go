package main

import (
	"github.com/dgraph-io/badger/v4"
)

type InMemDB struct {
	db *badger.DB
}

func InitDB() (*InMemDB, error) {
	opts := badger.DefaultOptions("").WithInMemory(true)
	opts.Logger = nil
	instance, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &InMemDB{db: instance}, nil
}

func (db *InMemDB) Close() error {
	return db.db.Close()
}

func (db *InMemDB) Get(key string) (string, error) {
	var value string
	return value, db.db.View(
		func(tx *badger.Txn) error {
			item, err := tx.Get([]byte(key))
			if err != nil {
				return err
			}
			valCopy, err := item.ValueCopy(nil)
			if err != nil {
				return err
			}
			value = string(valCopy)
			return nil
		})
}

func (db *InMemDB) Set(key, value string) error {
	return db.db.Update(
		func(txn *badger.Txn) error {
			return txn.Set([]byte(key), []byte(value))
		})
}
