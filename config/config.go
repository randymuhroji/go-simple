package config

import (
	"github.com/jmoiron/sqlx"
	"github.com/olivere/elastic/v7"
)

type Configuration struct {
	MysqlDB      *sqlx.DB
	ElasticeConn *elastic.Client
}
