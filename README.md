## Golang 学习
---
#### 相关文档 + 视频
+ [gin官方文档](https://gin-gonic.com/zh-cn/)
+ [go论坛](https://learnku.com/go)
+ [gorm](https://gorm.io/zh_CN/docs)
+ [大佬博客](https://liwenzhou.com/)


#### 现阶段完成内容
##### 基础设施
- [x] gin框架集成
- [x] mysql数据库集成
- [x] Docker容器化部署
  - [ ] 环境变量目前是硬编码进去，之后更新
- [ ] git action 集成

##### 功能模块
- [x] 设备信息相关API(可查看http文件夹下相关接口)
- [ ] 用户系统
- [ ] 鉴权系统
- [ ] 往APNS发送通知信息
- [ ] 相关模块单元测试



#### 跑起来看看？

##### 本地环境
***本地环境需要安装GO环境和mysql环境, 修改`dao/mysql.go`文件中相关信息***
```shell
go run .
curl http://localhost:8080/api/v1/devices
```


##### Dockekr环境
```shell

docker-compose up -d

curl http://localhost:8080/api/v1/devices
```