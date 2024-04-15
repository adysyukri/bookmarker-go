package rosedb

import "github.com/rosedblabs/rosedb/v2"

type Client interface {
	Add(key string, data any) error
	Update(key string, data any) error
	Delete(key string) error
	Get(key string, data any) error
}

type ClientImpl struct {
	svc *rosedb.DB
}

func NewClient(dbPath string) (*ClientImpl, error) {
	optDB := rosedb.DefaultOptions
	optDB.DirPath = dbPath

	db, err := rosedb.Open(optDB)
	if err != nil {
		return nil, err
	}

	return &ClientImpl{svc: db}, nil
}
