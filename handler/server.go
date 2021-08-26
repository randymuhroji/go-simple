package handler

import (
	"context"
	"fmt"
	"go-simple/config/env"
	"go-simple/utl/middleware/logging"
	"go-simple/utl/middleware/secure"
	"net/http"

	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

const DefaultPort = 8080

// HTTPServerMain main function for serving services over http

func (s *Service) HTTPServerMain() *echo.Echo {

	e := echo.New()

	// e.Use(middleware.ServerHeader, middleware.Logger)

	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())
	e.Use(secure.Headers())
	e.Use(logging.Logging())

	// administrator group
	adm := e.Group("/api/v1")

	// module system module
	ModuleSystem := adm.Group("/system")
	s.SystemModule.HandleRest(ModuleSystem)

	// auth module
	ModuleAuth := adm.Group("/auth")
	s.AuthModule.HandleRest(ModuleAuth)

	// module user
	ModuleUser := adm.Group("/user", s.MiddlewareAuth.BearerVerify())
	s.UserModule.HandleRest(ModuleUser)

	// module product
	ModuleProduct := adm.Group("/product", s.MiddlewareAuth.BearerVerify())
	s.ProductModule.HandleRest(ModuleProduct)

	// module order
	ModuleOrder := adm.Group("/order", s.MiddlewareAuth.BearerVerify())
	s.OrderModule.HandleRest(ModuleOrder)

	// module payment
	ModulePayment := adm.Group("/payment", s.MiddlewareAuth.BearerVerify())
	s.PaymentModule.HandleRest(ModulePayment)

	return e
}

// start handle service
func (s *Service) StartServer() {
	server := s.HTTPServerMain()
	listenerPort := fmt.Sprintf(":%v", env.Conf.HttpPort)
	if err := server.StartServer(&http.Server{
		Addr:         listenerPort,
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}); err != nil {
		server.Logger.Fatal(err.Error())
	}
}

// shutdown handler service
func (s *Service) ShutdownServer() {
	server := s.HTTPServerMain()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		server.Logger.Fatal(err.Error())
	}
}
