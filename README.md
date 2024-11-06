## <center>GinTpl</center>

基于Gin框架的快速开发脚手架

Rapid development scaffolding based on Gin framework

## 目录结构

```shell
├─cmd     // main入口
├─config  // 配置文件
├─internal // 应用私有包
│  ├─app   // 实际应用目录
│  │  └─web // 具体应用业务实现
│  └─errcode  // 应用公用错误码
└─pkg       // 公共包
    ├─base   // 基础包
    ├─config // 配置
    ├─log    // 日志
    ├─middleware  // 通用中间件
    ├─queue   // 队列
    ├─storage  // 存储
    ├─timer   // 定时器
    ├─utils  // 工具方法
    └─webapp  // 应用启动入口
```
