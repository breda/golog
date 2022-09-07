package server

import (
	"sync"
	"errors"
	"math/rand"
	"time"
)

type Log struct {
	mtx sync.Mutex
	records map[string]Record
}

type Record struct {
	Key string `json:"key"`
	Data string `json:"data"`
}

func NewLog() *Log {
	return &Log{
		records: make(map[string]Record),
	}
}

func (log *Log) Append(data string) (string, error) {
	log.mtx.Lock();
	defer log.mtx.Unlock();

	key := randomKey()
	record := Record{
		Key: key,
		Data: data,
	}

	log.records[key] = record
	return key, nil
}

func (log *Log) Get(key string) (*Record, error) {
	log.mtx.Lock();
	defer log.mtx.Unlock()

	record, present := log.records[key]
	if present == false {
		return nil, errors.New("Cannot find record with given key")
	}

	return &record, nil
}

func (log *Log) GetAll() []Record {
	log.mtx.Lock();
	defer log.mtx.Unlock()

	items := make([]Record, 0, len(log.records))
	for _, value := range log.records {
		items = append(items, value)
	}

	return items
}


const letters = "abcdefghijklmnopqrstuvwxyz"
func randomKey() string {
	tmp := make([]byte, 8)

	rand.Seed(time.Now().UnixNano())
	for i := range tmp {
		tmp[i] = letters[rand.Intn(len(letters))]
	}

	return string(tmp)
}