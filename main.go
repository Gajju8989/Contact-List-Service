package main

import (
	_ "github.com/go-sql-driver/mysql"
	"task2/contactapp/handler"
	"task2/contactapp/router"
	"task2/contactapp/service"
	"task2/database/config"
	"task2/database/migrate"
	"task2/database/mysql"
)

func main() {

	config.InitDB()
	err := migrate.MigrateAll(config.GetDB())
	if err != nil {
		return
	}
	// Initialize MySQL store
	mySqlStore := mysql.NewMySqlStore(config.GetDB())

	// Initialize service
	userService := service.NewUserSvc(mySqlStore)

	// Initialize handler
	userHandler := handler.NewHandler(userService)

	// Initialize router
	myRouter := router.NewRouter(userHandler)

	// Map the routes
	myRouter.MapRoutes()

}
