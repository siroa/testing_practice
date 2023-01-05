package app

import (
	s "main/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service s.UserService
}

func (uh UserHandler) GetUser(c *gin.Context) {
	reUsers, err := uh.service.GetUsers()
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, nil)
	} else {
		c.SecureJSON(http.StatusOK, reUsers)
	}
}
