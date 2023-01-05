package app

import (
	"io/ioutil"
	s "main/service"
	"main/utils/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShippingHandler struct {
	service s.ShippingService
}

func (sh ShippingHandler) PostShippingFee(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		e := errs.NewBadRequestError(errs.PropErr, errs.PropMess)
		c.SecureJSON(e.Code, e.AsMessage())
		return
	}
	fee, e := sh.service.PostShippingFee(body)
	if e != nil {
		c.SecureJSON(e.Code, e.AsMessage())
		return
	}
	c.SecureJSON(http.StatusOK, fee)
}
