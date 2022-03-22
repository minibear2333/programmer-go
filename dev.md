### 环境准备

* 运行 [env_ready.sh](env_ready.sh) 参考讲解 https://go-zero.dev/cn/
* 环境参考 https://go-zero.dev/cn/concept-introduction.html

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

### mongo model 生成

mongo的json文件放在 mongo 目录下

* 首先使用[json2struct](http://json2struct.mervine.net/)工具，把单个json转换为struct
* 增加到 [api/model/types.go](api/model/types.go) 里面，并修改增加`bson`的命名
* 在 types.go 文件头部更新 `//go:generate goctl model mongo -t Comments -t xxx`
* 执行脚本生成model

参考：[mongo生成model](https://pkg.go.dev/git.i2edu.net/i2/go-zero/tools/goctl/model/mongo#section-readme)

### 其他

开发文档参考 [go-zero项目开发](https://go-zero.dev/cn/project-dev.html)