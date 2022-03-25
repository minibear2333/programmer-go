package main

import (
	"flag"
	"fmt"
	"github.com/minibear2333/programmer-go/api/global"
	"github.com/minibear2333/programmer-go/api/initialize"

	"github.com/minibear2333/programmer-go/api/internal/handler"
	"github.com/minibear2333/programmer-go/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/pg-backend.yaml", "the config file")

func main() {
	flag.Parse()

	conf.MustLoad(*configFile, &global.CONFIG) // 初始化配置文件
	global.LOG = initialize.Zap()              // 初始化日志库
	initialize.Redis()                         // 初始化 Redis

	ctx := svc.NewServiceContext(global.CONFIG)
	server := rest.MustNewServer(global.CONFIG.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", global.CONFIG.Host, global.CONFIG.Port)
	server.Start()
}
