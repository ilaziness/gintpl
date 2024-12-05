## <center>GinTpl</center>

基于Gin框架的快速开发脚手架，支持多子应用

Rapid development scaffolding based on Gin framework

## 可用组件组件

### 1. MySQL数据库
- gorm
- ent
- sqlx

### 2. Redis
### 3. Redis Cache
### 4. 队列
- rocket mq

### 5. 定时器
### 6. 日志
- zap

### 7. 配置中心
- nacos

### 8. 服务注册和发现
- nacos

### 9. 可观测性
- OTEL链路追踪

## 使用方法

复制项目文件到自己的目录中，修改`go.mod`模块名称。

批量替换导入路径，`github.com/ilaziness/gintpl`替换为自己模块。

修改`cmd/web.go`，按需添加自己需要的组件。

具体应用逻辑在`internal/app`下面，多个子应用新建多个目录即可。


## 应用配置

配置文件默认在工作目录的`config`目录下，所有`toml`文件都会被加载解析。

按环境加载配置：

- 设置环境变了`ENV_CONFIG_ENV`, 比如设置的值为`dev`, 那么会加载后缀是`dev.toml`的文件。

多应用配置：

- 在`config`创建对应应用名称的目录，比如`web`应用，那么会加载`config/web`里的文件。
- `github.com/ilaziness/gokit/config.LoadConfig`，第二个参数传入应用名称，比如`web`。

