package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (r *impl) MapRoutes() {
	r1 := gin.Default()
	r1.POST("/register", r.handler.CreateUserHandler)
	r1.POST("/contacts", r.handler.CreateContact)
	r1.PUT("/contacts", r.handler.UpdateSpamStatusHandler)
	r1.GET("/contacts/search/name", r.handler.SearchByNameHandler)
	r1.GET("/contacts/search/phoneno", r.handler.SearchByPhoneNo)
	err := r1.Run(":8080")
	if err != nil {
		fmt.Print("Unable To Run The Server")
		return
	}

}
