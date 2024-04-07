package handler

import (
	"github.com/gin-gonic/gin"
	"task2/contactapp/service"
)

type ContactAppHandler interface {
	CreateUserHandler(c *gin.Context)
	CreateContact(c *gin.Context)
	UpdateSpamStatusHandler(c *gin.Context)
	SearchByNameHandler(c *gin.Context)
	SearchByPhoneNo(c *gin.Context)
}

type impl struct {
	authService service.ContactAppSvc
}

func NewHandler(authService service.ContactAppSvc) ContactAppHandler {
	return &impl{authService: authService}
}
