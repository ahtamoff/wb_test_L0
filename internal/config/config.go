package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env"`
	Postgres `yaml:postgres`
	HTTPServer `yaml:http_srver`
}

type Postgres struct {
	Host string `yaml:host`
	Port string `yaml:port`
	User string `yaml:user`
	Password string `yaml:password`
	DBName string `yaml:db_name`
}

type HTTPServer struct {
	Address string `yaml:address`
	Timeout time.Duration `yaml:timeout`
	IdleTimeout time.Duration `yaml:idle_timeout`
}

type Nats struct {
	ClusterId string `yaml:cluster_id`
	ClientId string `yaml:client_id`
	Url string `yaml:url`
}


//читаем config.yaml, создаем и заполняем объект Config struct 
//(приставка "Must" означает что функция вместо возврата ошибки будет паниковать)
func MustLoad(configPath string) *Config {
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil{
		log.Fatalf("cannot read config: %s", err)
	}
	return &cfg
}