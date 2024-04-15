package bookmark

import "github.com/adysyukri/bookemarker-go/pkg/sqlite"

type service struct {
	db sqlite.Client
}

type Service interface{}

func NewService(db sqlite.Client) Service {
	return &service{db}
}
