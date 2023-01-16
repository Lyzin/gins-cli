### 一、项目目录
- configs：配置文件
- docs：文档集合
- global：全局变量
- internal：
    - api：处理入参层
    - dao：数据访问层，所有与数据相关的操作都会在dao层，比如Mysql，Redis等
    - middleware：HTTP中间件
    - model：模型层，用于存放model对象
    - routers：路由相关的逻辑
    - service：项目核心业务逻辑
- pkg：项目相关的模块包
- storage：项目生成的临时文件
- third_party：第三方的资源工具，如Swagger UI
- Makefile：构建、运行、打包脚本
- .air.toml: 项目热加载配置文件

### 二、项目运行
#### 1、创建项目
> 本项目由如下脚手架命令快速创建
```bash
gins-cli startproject {{.ProjectName}}
```

#### 2、启动项目
> 在{{.ProjectName}}项目的根目录下执行以下命令
```bash
# 启动项目
make run

# 编译项目
make build
```
