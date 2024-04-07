package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task2/contactapp/service/model"
)

func (i *impl) UpdateSpamStatusHandler(c *gin.Context) {
	req := &model.ContactListRequest{}
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if req.PhoneNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "phone number is empty"})
		return
	}
	err = i.authService.UpdateSpamStatus(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully marked phone number as spam."})
}
