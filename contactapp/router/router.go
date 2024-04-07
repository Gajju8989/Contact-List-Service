package router

import (
	"task2/contactapp/handler"
)

type Router interface {
	MapRoutes()
}
type impl struct {
	handler handler.ContactAppHandler
}

func NewRouter(handler handler.ContactAppHandler) Router {
	return &impl{handler: handler}
}
