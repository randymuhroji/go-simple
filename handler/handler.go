package handler

import (
	"go-simple/config"
	"go-simple/config/database"
	authModule "go-simple/module/v1/auth"
	orderModule "go-simple/module/v1/order"
	paymentModule "go-simple/module/v1/payment"
	productModule "go-simple/module/v1/product"
	system "go-simple/module/v1/system"
	authMid "go-simple/utl/middleware/auth"

	userModule "go-simple/module/v1/user"
)

type Service struct {
	MiddlewareAuth *authMid.Handle
	SystemModule   *system.Module
	UserModule     *userModule.Module
	ProductModule  *productModule.Module
	OrderModule    *orderModule.Module
	PaymentModule  *paymentModule.Module
	AuthModule     *authModule.Module
}

func InitHandler() *Service {

	// mysql init
	MySQLConnection, err := database.MysqlDB()
	if err != nil {
		panic(err)
	}

	config := config.Configuration{
		MysqlDB: MySQLConnection,
	}

	// set service modular
	middlewareAuth := authMid.InitAuthMiddleware(config)

	moduleAuth := authModule.InitModule(config)
	moduleSystem := system.InitModule(config)
	moduleUser := userModule.InitModule(config)
	moduleProduct := productModule.InitModule(config)
	moduleOrder := orderModule.InitModule(config)
	modulePayment := paymentModule.InitModule(config)

	return &Service{
		SystemModule:   moduleSystem,
		AuthModule:     moduleAuth,
		UserModule:     moduleUser,
		ProductModule:  moduleProduct,
		OrderModule:    moduleOrder,
		PaymentModule:  modulePayment,
		MiddlewareAuth: middlewareAuth,
	}
}
