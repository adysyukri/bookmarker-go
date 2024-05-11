package concurrent

import (
	"fmt"
	"time"

	"github.com/a-h/templ"
)

type service struct {
}

type Service interface {
	Page() (templ.Component, error)
}

func NewService() Service {
	return &service{}
}

func (s *service) Page() (templ.Component, error) {
	go logBackground()
	return ConcurrentPage(), nil
}

func logBackground() {
	fmt.Println("start..")

	time.Sleep(3 * time.Second)

	fmt.Println("..1")
	time.Sleep(1 * time.Second)
	fmt.Println("..2")
	time.Sleep(1 * time.Second)
	fmt.Println("..3")
	time.Sleep(1 * time.Second)
	fmt.Println("..4")
	time.Sleep(1 * time.Second)

	fmt.Println("..end")
}
