// Package datastore for store and load email hashes information in json file
package datastore

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type DataStore struct {
	sync.RWMutex
	Items    []Item
	filename string
}

type Item struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func NewDataStore(filename string) *DataStore {
	ds := &DataStore{
		filename: filename,
	}

	err := ds.loadFromFile()
	if err != nil {
		ds.Items = make([]Item, 0)
	}
	return ds
}

func (ds *DataStore) AddItem(item Item) {
	ds.Lock()
	ds.Items = append(ds.Items, item)
	ds.Unlock()
	err := ds.saveToFile()
	if err != nil {
		log.Println("Ошибка при записи данных в файл:", ds.filename)
	}
}

func (ds *DataStore) RemoveItem(item Item) {
	ds.Lock()
	defer ds.Unlock()

	for i, v := range ds.Items {
		if v.Email == item.Email {
			ds.Items = append(ds.Items[:i], ds.Items[i+1:]...)
		}
	}
	err := ds.saveToFile()
	if err != nil {
		log.Println("Ошибка при записи данных в файл:", ds.filename)
	}
}

func (ds *DataStore) FindByEmail(email string) *Item {
	for _, v := range ds.Items {
		if v.Email == email {
			return &v
		}
	}

	return nil
}

func (ds *DataStore) FindByHash(hash string) *Item {
	for _, v := range ds.Items {
		if v.Hash == hash {
			return &v
		}
	}

	return nil
}

func (ds *DataStore) saveToFile() error {
	data, err := json.MarshalIndent(ds.Items, "", " ")
	if err != nil {
		return err
	}

	tempFile := ds.filename + ".tmp"
	err = os.WriteFile(tempFile, data, 0o644)
	if err != nil {
		return err
	}

	return os.Rename(tempFile, ds.filename)
}

func (ds *DataStore) loadFromFile() error {
	data, err := os.ReadFile(ds.filename)
	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(data, &ds.Items)
	if err != nil {
		return err
	}

	return nil
}

func DropFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}
