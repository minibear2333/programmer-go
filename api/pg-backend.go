package main

import (
	"flag"
	"fmt"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/initialize"
	"github.com/minibear2333/programmer-go/api/utils"

	"github.com/minibear2333/programmer-go/api/internal/handler"
	"github.com/minibear2333/programmer-go/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	configFileStr := "etc/pg-backend.yaml"
	if utils.Exists("etc/.pg-backend.yaml") {
		configFileStr = "etc/.pg-backend.yaml"
	}
	configFile := flag.String("f", configFileStr, "the config file")

	flag.Parse()

	conf.MustLoad(*configFile, &global.CONFIG) // 初始化配置文件
	global.LOG = initialize.Zap()              // 初始化日志库
	initialize.Redis()                         // 初始化 Redis
	initialize.Mongo()                         // 初始化 mongo

	ctx := svc.NewServiceContext(global.CONFIG)
	server := rest.MustNewServer(global.CONFIG.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", global.CONFIG.Host, global.CONFIG.Port)
	server.Start()
}
