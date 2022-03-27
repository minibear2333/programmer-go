package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	Auth  JWT	`json:"Auth" yaml:"Auth"`
	Zap   Zap   `json:"Zap" yaml:"Zap"`
	Redis Redis `json:"Redis" yaml:"Redis"`
}
