package handler

import (
	"kumparan/config"
	"kumparan/config/database"
	article "kumparan/module/v1/article"
	authMid "kumparan/utl/middleware/auth"

	cnfElastic "kumparan/config/elastic"
)

type Service struct {
	MiddlewareAuth *authMid.Handle
	ArticleModule  *article.Module
}

func InitHandler() *Service {

	// mysql init
	MySQLConnection, err := database.MysqlDB()
	if err != nil {
		panic(err)
	}

	// elastic init
	ElasticConnection, err := cnfElastic.NewElastic()
	if err != nil {
		panic(err)
	}

	config := config.Configuration{
		MysqlDB:      MySQLConnection,
		ElasticeConn: ElasticConnection,
	}

	// set service modular
	middlewareAuth := authMid.InitAuthMiddleware(config)
	moduleArticle := article.InitModule(config)

	return &Service{
		ArticleModule:  moduleArticle,
		MiddlewareAuth: middlewareAuth,
	}
}
