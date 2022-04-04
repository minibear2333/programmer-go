package initialize

import (
	"github.com/minibear2333/programmer-go/api/global"
	"log"
	"os"
	"strconv"
)

func CheckAndReplaceEnv() {
	if value, ok := os.LookupEnv("MongoAddr"); ok {
		global.CONFIG.Mongo.Addr = value
	}
	if value, ok := os.LookupEnv("RedisDB"); ok {
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("请检查环境变量 err:%s", err)
		}
		global.CONFIG.Redis.DB = v
	}
	if value, ok := os.LookupEnv("RedisPass"); ok {
		global.CONFIG.Redis.Password = value
	}
	if value, ok := os.LookupEnv("RedisAddr"); ok {
		global.CONFIG.Redis.Addr = value
	}
	if value, ok := os.LookupEnv("AccessExpire"); ok {
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			log.Fatalf("请检查环境变量 err:%s", err)
		}
		global.CONFIG.Auth.AccessExpire = v
	}
	if value, ok := os.LookupEnv("AccessSecret"); ok {
		global.CONFIG.Auth.AccessSecret = value
	}
}
