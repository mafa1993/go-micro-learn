# go-micro-learn

用golang实现微服务架构的学习案例。

## 目录结构

1. consul下放了consul的源码包
2. go\src下放了代码

### gin目录结构

```
// gin 
Frame 
    - app
    - conf
    - middleward
    - models
    - routes
    - services
    - uitl
    - main.go
// app 目录
app 
    - api
        - v1
            - goods.go
        - v2
        - route.go
    - app.go

//conf目录
conf
    - config.ini
models
    - models.go
    - goods.go
services
    - goodsService.go
util
    - config.go
```

### 使用到的包

1. ini： go get -u github.com/go-ini/ini
2. gorm  
    - go get -u gorm.io/gorm
    - go get -u gorm.io/driver/sqlite
3. gin
    - go get -u github.com/gin-gonic/gin

### 开发流程

1. 我们使用go mod方式初始化gin目录 (不在go path 目录下开发，使用mod模式)
    - go mod init module name
2. 