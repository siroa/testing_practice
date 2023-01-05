package app

import (
	"log"
	d "main/domain"
	s "main/service"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	uStub := d.NewUserRepositoryStub()
	us := s.NewUserService(uStub)
	uh := UserHandler{us}

	oStub := d.NewOrderRepositoryStub()
	os := s.NewOrderService(oStub)
	oh := OrderHandler{os}

	sStub := d.NewShippingRepositoryStub()
	ss := s.NewShippingService(sStub)
	sh := ShippingHandler{ss}

	v1 := r.Group("/v1")
	{
		v1.GET("/users", uh.GetUser)
		v1.POST("/orders", oh.PostOrders)
		v1.POST("/orders/shipping_fee", sh.PostShippingFee)
	}

	log.Println("Webサーバーを開始します。")
	r.Run(":18080")
}
