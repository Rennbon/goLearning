package config

import (
	"time"
)

type MongoDBConfig struct {
	Addr      string
	Timeout   time.Duration
	PoolLimit int
	DbName1   string
}
