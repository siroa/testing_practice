package app

import (
	"io/ioutil"
	s "main/service"
	"main/utils/errs"
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service s.UserService
}

func (uh UserHandler) Login(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		e := errs.NewBadRequestError(errs.PropErr, errs.PropMess)
		c.SecureJSON(e.Code, e.AsMessage())
		return
	}
	key, e := uh.service.ValidUser(body)
	if e != nil {
		c.SecureJSON(e.Code, e.AsMessage())
		return
	}
	c.SecureJSON(http.StatusOK, map[string]string{"token": key})
}

func (uh UserHandler) SecureTest(c *gin.Context) {
	claims, ok := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			map[string]string{"message": "Failed to get validated JWT claims."},
		)
		return
	}

	//customClaims, ok := claims.CustomClaims.(*auth.CustomClaimsExample)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			map[string]string{"message": "Failed to cast custom JWT claims to specific type."},
		)
		return
	}

	c.JSON(http.StatusOK, claims)
}

func (uh UserHandler) GetUser(c *gin.Context) {
	if jwtErr := ValidJwt(c); jwtErr != nil {
		c.SecureJSON(jwtErr.Code, jwtErr.AsMessage())
	}
	reUsers, err := uh.service.GetUsers()
	if err != nil {
		c.SecureJSON(http.StatusBadRequest, nil)
	} else {
		c.SecureJSON(http.StatusOK, reUsers)
	}
}
