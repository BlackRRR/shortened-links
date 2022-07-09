package config

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	ServicePort   string          `yaml:"service_port"`
	RepositoryCfg *pgxpool.Config `yaml:"repository_cfg"`
}

func InitConfig() (*Config, error) {
	vp := viper.New()

	vp.AddConfigPath("config")
	vp.SetConfigName("config")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config")
	}
	var cfg Config

	connString := fmt.Sprintf("postgres://%s:%s@%s/%s?pool_max_conns=%s",
		vp.Get("config.db_conn_config.user"),
		vp.Get("config.db_conn_config.host"),
		vp.Get("config.db_conn_config.password"),
		vp.Get("config.db_conn_config.db_name"),
		vp.Get("config.db_conn_config.pool_max_conns"))

	repCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, errors.Wrap(err, "parse repository config")
	}

	cfg.ServicePort = vp.GetString("config.service_port")
	cfg.RepositoryCfg = repCfg

	return &cfg, err

}
