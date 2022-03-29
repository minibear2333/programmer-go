### 环境准备

* 运行 [env_ready.sh](env_ready.sh) 参考讲解 https://go-zero.dev/cn/
* 环境参考 https://go-zero.dev/cn/concept-introduction.html

初始化本地redis mongo环境

```shell
cd api
docker-compose up -d redis mongo
```

写入hosts文件，如果是`win`环境请手动写入

```shell
echo pg-redis 127.0.0.1 >> /etc/hosts
echo pg-mongo 127.0.0.1 >> /etc/hosts
```

运行项目，可以用任何`ide`运行，使用的`go mod`的模式，本地下载包

```shell
cd api
go mod tidy
```

### 开发流程

* goctl环境准备
* 数据库设计
* 新建工程
* 创建服务目录
* 创建服务类型（api/rpc/rmq/job/script）
* 编写api~~、proto文件~~
* 代码生成
* 生成数据库访问层代码model
* 配置config，yaml变更
* 资源依赖填充（ServiceContext）
* 添加中间件
* 业务代码填充
* 错误处理

### 目录参考

https://github.com/zeromicro/go-zero-template

```
├── LICENSE 证书
├── README.md 自述
├── api api文档
│   ├── Dockerfile api构建dock
│   ├── api.api go-zero的api主文件
│   ├── build.sh 构建容器命令
│   ├── etc 配置文件所在目录
│   ├── global 全局变量
│   ├── go.mod
│   ├── go.sum
│   ├── goctl-apis go-zero的api子文件
│   ├── initialize 初始化引用组件
│   ├── internal 内部代码
│   │   ├── config 配置文件构建代码
│   │   ├── handler api的controller层
│   │   ├── logic 逻辑代码
│   │   ├── svc 请忽略
│   │   └── types api所使用的model，如果有更新请同步修改，建议删除后再生成一次
│   ├── logs 日志
│   ├── model mongo的model文件
│   ├── pg-backend.go 主文件
│   ├── run.sh 运行docker的脚本
│   └── utils 工具类
├── dev.md 开发自述
├── env_ready.sh 环境配置脚本
└── mongo model的数据库脚本
```

### mongo model 生成

mongo的json文件放在 mongo 目录下

* 首先使用[json2struct](http://json2struct.mervine.net/)工具，把单个json转换为struct
* 增加到 [api/model/types.go](api/model/types.go) 里面，并修改增加`bson`的命名
* 在 types.go 文件头部更新 `//go:generate goctl model mongo -t Comments -t xxx`
* 执行脚本生成model

参考：[mongo生成model](https://pkg.go.dev/git.i2edu.net/i2/go-zero/tools/goctl/model/mongo#section-readme)

### 加密环境

这里的内容请非内部开发者忽略

含有密码文件放到 `etc/.pg-backend.yaml` 文件中了，并加了忽略提交，正式上线时，将会切换为环境变量的形式。

### 贡献须知

在issue中会经常发布任务，欢迎尝试挑战

再次感谢您的贡献，一起学习共建开源项目

### 其他

开发文档参考 [go-zero项目开发](https://go-zero.dev/cn/project-dev.html)