package config

import (
	"fmt"

	"github.com/boltdb/bolt"
)

type Config struct {
	HttpPort   int
	HttpHost   string
	DbFilename string
	Database   *bolt.DB
}

func (cfg *Config) GetHttpHost() string {
	return fmt.Sprintf("%s:%v", cfg.HttpHost, cfg.HttpPort)
}

func New() *Config {
	return &Config{
		DbFilename: "exchange.bolt.db",
		HttpPort:   3000,
		HttpHost:   "127.0.0.1",
	}
}
