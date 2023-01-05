package app

import (
	"io/ioutil"
	s "main/service"
	"main/utils/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service s.OrderService
}

func (oh OrderHandler) PostOrders(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		e := errs.NewBadRequestError(errs.PropErr, errs.PropMess)
		c.SecureJSON(e.Code, e.AsMessage())
		return
	}
	receipt, e := oh.service.PostOrders(body)
	if e != nil {
		c.SecureJSON(e.Code, e.AsMessage())
		return
	}
	c.SecureJSON(http.StatusOK, receipt)
}
