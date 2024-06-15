package order

import (
	"context"

	"github.com/a-h/templ"
	"github.com/adysyukri/bookemarker-go/pkg/sqlite"
)

type service struct {
	db sqlite.Client
}

type Service interface {
	Page(ctx context.Context) (templ.Component, error)
}

func NewService(db sqlite.Client) Service {
	return &service{db}
}

func (s *service) Page(ctx context.Context) (templ.Component, error) {

	return Page(), nil
}
