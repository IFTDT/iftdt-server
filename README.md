## Golang 学习
---
#### 相关文档 + 视频
+ [gin官方文档](https://gin-gonic.com/zh-cn/)
+ [go论坛](https://learnku.com/go)
+ [gorm](https://gorm.io/zh_CN/docs)
+ [大佬博客](https://liwenzhou.com/)
+ [《The Go Programming Language》](http://gopl.io/)
+ [《The Go Programming Language（中文版）》](https://books.studygolang.com/gopl-zh/)


#### 现阶段完成内容
##### 基础设施
- [x] gin框架集成
- [x] mysql数据库集成
- [x] Docker容器化部署
  - [x] ~~环境变量目前是硬编码进去，之后更新~~(已完成)
- [x] github action 集成
  - [X] 集成了最基本的东西，需要后续完善
  - [X] CI流程完成
  - [ ] CD流程完成

##### 功能模块
- [x] 设备信息相关API(可查看http文件夹下相关接口)
- [x] 用户系统
- [x] 鉴权系统
- [x] 往APNS发送通知信息
- [x] CI单元测试
- [ ] 相关模块单元测试



#### 跑起来看看？

+ 首先拷贝环境变量配置文件 `copy .env.example app.env`

+ 修改相关配置信息

+ 如果需要推送测试，请在项目根目录下添加 `cert.p12` 文件

##### 本地环境
```shell
go run .

curl http://localhost:8080/api/v1/devices
```


##### Dockekr环境
```shell

docker-compose up -d

curl http://localhost:8080/api/v1/devices
```
