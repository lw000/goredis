package config

import (
	"encoding/json"
	"github.com/lw000/gocommon/db/rdsex"
	"io/ioutil"
)

type cfgConfig struct {
	Debug int64 `json:"debug"`
	Redis struct {
		Host         string `json:"host"`
		Psd          string `json:"psd"`
		Db           int64  `json:"db"`
		PoolSize     int64  `json:"poolSize"`
		MinIdleConns int64  `json:"minIdleConns"`
	} `json:"redis"`
}

// JSONConfig json配置
type JSONConfig struct {
	Debug  int64
	RdsCfg tyrdsex.JsonConfig
}

// NewJSONConfig ...
func NewJSONConfig() *JSONConfig {
	return &JSONConfig{}
}

// LoadJSONConfig ...
func LoadJSONConfig(file string) (*JSONConfig, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var ccf cfgConfig
	if err = json.Unmarshal(data, &ccf); err != nil {
		return nil, err
	}

	cfg := NewJSONConfig()
	cfg.Debug = ccf.Debug
	// 缓存配置
	cfg.RdsCfg.Host = ccf.Redis.Host
	cfg.RdsCfg.Db = ccf.Redis.Db
	cfg.RdsCfg.Psd = ccf.Redis.Psd
	cfg.RdsCfg.PoolSize = ccf.Redis.PoolSize
	cfg.RdsCfg.MinIdleConns = ccf.Redis.MinIdleConns

	return cfg, err
}
