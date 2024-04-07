package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (r *impl) MapRoutes() {
	r1 := gin.Default()
	r1.POST("/register", r.handler.CreateUserHandler)
	r1.POST("/contacts", r.handler.CreateContact)
	err := r1.Run(":8080")
	if err != nil {
		fmt.Print("Unable To Run The Server")
		return
	}

}
