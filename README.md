
# 程序员向前冲

面向社招程序员的一站式网站-后端

前端位于 [programmer-go-front](https://github.com/Sakura-echos/programmer-go-front)

## 如何使用

如果只是想在本地启动本程序便于前端调试或者体验`api`的话，可以继续；如果是需要参与开发，请跳到 [如何贡献](#如何贡献)

先安装 `docker` 和 `docker-compose` 安装完毕后直接在该项目的根目录运行

```bash
docker-compose up -d
```

运行成功后，插入一跳数据

```bash
docker exec -it pg-redis redis-cli -p 16379 -a "1234567890"  HMSET "c7cec5" "openid" "aROD9s357bclJK9eLwXrGfMsuvZm"
```

这样子前端同学就可以在本地调试了！可以把 [接口文档](https://documenter.getpostman.com/view/18714614/UVyoXyKn) 导入到 `postman` 请求本地接口就可以了。

注意配置两个环境变量

* `pg-backend` 对应 `http://localhost`
* `token` 对应调用 login 接口后获得的 `access_token` 的值

## 如何贡献

非常感谢您的贡献，开发环境与运行方式、代码目录结构请参考 [dev.md](dev.md)

建议安装`git commit`提交规范辅助工具

```shell
# git 提交辅助工具 可选 参考 https://coding3min.com/1617.html
npm install -g conventional-changelog
npm install -g conventional-changelog-cli
```

使用 `git cz` 进行提交

## Authors&Contributors

感谢技术大拿们的倾力贡献！

- [@minibear2333](https://github.com/minibear2333)
- [@gong-guowei](https://github.com/guowei-gong)
- [@kurumii](https://github.com/kurumii)
- 其余贡献者请自行补充


## Support

For support, please create [issue](https://github.com/minibear2333/programmer-go/issues/new)


## License

可闭源，请注明出处！

[GNU General Public License v3.0](LICENSE)

