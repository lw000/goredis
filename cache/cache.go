package cache

import (
	"github.com/lw000/gocommon/db/rdsex"
	"log"
)

var (
	rds *tyrdsex.RdsServer
)

func init() {
	rds = &tyrdsex.RdsServer{}

}

func Instance() *tyrdsex.RdsServer {
	return rds
}

func OpenRedis(cfg *tyrdsex.JsonConfig) error {
	err := rds.OpenWithJsonConfig(cfg)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func CloseRedis() {
	_ = rds.Close()
}
