## 一、介绍

- 本项目是通过参考[golang标准工程结构布局](https://github.com/golang-standards/project-layout/blob/master/README_zh.md)而来的一个快速创建go语言web项目的工具
- 本项目的工具功能实现使用了[cobra](https://github.com/spf13/cobra)库
### 1、功能

- 快速创建golang标准工程结构的web项目，并且底层是[gin框架](https://gin-gonic.com/)
- 自定义的统一响应输出结构
- 支持swag生成接口文档
- 支持air包热加载
- 支持zap日志库进行全量日志、错误日志记录
- 支持makefile管理项目的启动和构建

### 2、本工具创建的项目目录结构

> 主要包含如下目录
>
> - configs：配置文件
> - docs：文档集合
> - global：全局变量
> - internal：
>     - api：处理入参层
>     - dao：数据访问层，所有与数据相关的操作都会在dao层，比如Mysql，Redis等
>     - middleware：HTTP中间件
>     - model：模型层，用于存放model对象
>     - routers：路由相关的逻辑
>     - service：项目核心业务逻辑
> - pkg：项目相关的模块包
> - storage：项目生成的临时文件
> - third_party：第三方的资源工具，如Swagger UI
> - Makefile：构建、运行、打包脚本
> - .air.toml: 项目热加载配置文件

## 二、工具使用

### 1、安装前置依赖

#### 1.1 安装swag

> 使用的是[swag库](https://github.com/swaggo/swag)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

#### 1.2 安装air

> 使用的[air库](https://github.com/cosmtrek/air)

```bash
go install github.com/cosmtrek/air@latests
```

### 2、安装本工具

```bash
go install github.com/Lyzin/gins-cli@latest
```
#### 2.1 查看本工具帮助信息

```bash
gins-cli -h
```

### 3、创建项目

```bash
# 进去需要创建项目的目录
cd your-project-dir

# 创建项目，项目名就是项目的目录名以及go mod文件中的mod名，所以建议创建的时候规范一些，比如这样的：example.com/xx/xxx
gins-cli startproject you-project-name
```

### 4、运行项目

```bash
# 进入新创建的项目
cd your-new-project-dir

# 启动项目
make run
```

### 5、查看项目效果

> 项目启动以后，浏览器打开[http://127.0.0.1:8090/swagger/index.html](http://127.0.0.1:8090/swagger/index.html)，可看到有一个的示例user-info接口，可以直接发起请求



