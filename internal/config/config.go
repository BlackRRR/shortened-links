package config

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"go.uber.org/config"
	"os"
)

type Config struct {
	ServicePort string `yaml:"service_port"`

	RepositoryCfg *pgxpool.Config `yaml:"repository_cfg"`
	DBConnConfig  DBConnConfig    `yaml:"db_conn_config"`
}

type DBConnConfig struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Host         string `yaml:"host"`
	DBName       string `yaml:"db_name"`
	PoolMaxConns string `yaml:"pool_max_conns"`
}

func InitConfig() (*Config, error) {
	//Get path to config
	cfgPath := os.Getenv("CONFIG_PATH")

	//Open config file
	configFile, err := os.Open(cfgPath + "config.yaml")
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config")
	}

	//Decode yaml config
	cfgYaml, err := config.NewYAML(config.Source(configFile))
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode config")
	}

	var cfg Config
	//Encode yaml config into Config struct
	if err = cfgYaml.Get("config").Populate(&cfg); err != nil {
		return nil, errors.Wrap(err, "marshal config")
	}

	//Fill connection to postgres from config
	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?pool_max_conns=%s",
		cfg.DBConnConfig.User,
		cfg.DBConnConfig.Password,
		cfg.DBConnConfig.Host,
		cfg.DBConnConfig.DBName,
		cfg.DBConnConfig.PoolMaxConns)

	//Builds a Config from connString.
	repCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "parse repository config")
	}

	cfg.RepositoryCfg = repCfg

	return &cfg, err

}
