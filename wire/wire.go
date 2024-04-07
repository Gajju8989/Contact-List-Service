package wire

import (
	"github.com/google/wire"
	"task2/contactapp/handler"
	"task2/contactapp/router"
	"task2/contactapp/service"
	"task2/database/config"
	"task2/database/mysql"
)

func InitializeApp() (router.Router, error) {
	wire.Build(
		config.GetDB,
		mysql.NewMySqlStore,
		service.NewUserSvc,
		handler.NewHandler,
		router.NewRouter,
	)
	return nil, nil
}
