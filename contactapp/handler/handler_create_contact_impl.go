package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"task2/contactapp/service/model"
)

func (h *impl) CreateContact(c *gin.Context) {
	reqData := &model.ContactListRequest{}
	err := c.Bind(&reqData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.contactAppService.CreateContact(c, reqData)
	if err != nil {
		if errors.Is(err, errors.New("Unauthorized: User Not Found")) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Contact Created Successfully"})
}
