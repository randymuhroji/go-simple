package env

import "github.com/caarlos0/env/v6"

type ConfigurationEnvironment struct {
	ApplicationEnvironment
	DatabaseEnvironment
	NATSEnvironment
	ELASTICEnvironment
}

type ApplicationEnvironment struct {
	AppEnv   string `env:"APP_ENV" envDefault:"development"`
	HttpPort string `env:"PORT" envDefault:"8080"`
}

type DatabaseEnvironment struct {
	MaxIdle     int    `env:"DB_MAX_IDLE_CONN" envDefault:"10"`
	MaxOpenConn int    `env:"DB_MAX_OPEN_CONN" envDefault:"100"`
	DBName      string `env:"DB_NAME_MYSQL"`
	DBUser      string `env:"DB_USER_MYSQL"`
	DBPass      string `env:"DB_PASSWORD_MYSQL"`
	DBHost      string `env:"DB_HOST_MYSQL"`
	DBPort      string `env:"DB_PORT_MYSQL"`
}

type NATSEnvironment struct {
	NatsHost string `env:"NATS_HOST" envDefault:"0.0.0.0:4222"`
}

type ELASTICEnvironment struct {
	ElasticHost string `env:"ELASTIC_HOST" envDefault:"http://0.0.0.0:9200"`
}

var Conf = ConfigurationEnvironment{}

func LoadEnv() {
	if err := env.Parse(&Conf); err != nil {
		panic(err.Error())
	}
}
