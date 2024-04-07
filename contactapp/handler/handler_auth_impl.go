package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"task2/contactapp/service/model"
)

func (h *impl) CreateUserHandler(c *gin.Context) {
	reqData := &model.UserRegistrationsRequestModel{}
	err := c.Bind(&reqData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	//Hash The Password
	hashedPassword, err := hashPassword(reqData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	reqData.Password = hashedPassword
	err = h.authService.CreateUser(c, reqData)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Contact created successfully"})

}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
