package app

import (
	"log"
	d "main/domain"
	s "main/service"

	auth "main/utils/auth"
	"main/utils/errs"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	proxies := [...]string{"localhost", "127.0.0.1"}
	r.SetTrustedProxies(proxies[:])
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
		v1.POST("/login", uh.Login)
		v1.GET("/secure", auth.CheckJWT(), uh.SecureTest)
		v1.GET("/users", auth.CheckJWT(), uh.GetUser)
		v1.POST("/orders", auth.CheckJWT(), oh.PostOrders)
		v1.POST("/orders/shipping_fee", auth.CheckJWT(), sh.PostShippingFee)
	}

	log.Println("Webサーバーを開始します。")
	r.Run(":18080")
}

func ValidJwt(c *gin.Context) *errs.AppError {
	claims, ok := c.Request.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)
	if !ok {
		return errs.NewUnauthorizedError(errs.Unauthorized, errs.UnauthorizedMesss)
	}

	customClaims, ok := claims.CustomClaims.(*auth.CustomClaimsExample)
	if !ok {
		return errs.NewUnauthorizedError(errs.Unauthorized, errs.UnauthorizedMesss)
	}

	if customClaims.Name != "admin" {
		return errs.NewUnauthorizedError(errs.Unauthorized, errs.UnauthorizedMesss)
	}

	return nil
}
