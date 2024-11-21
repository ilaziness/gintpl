## <center>GinTpl</center>

基于Gin框架的快速开发脚手架

Rapid development scaffolding based on Gin framework

## 目录结构

```shell
├─cmd                   // main入口
├─config                // 配置文件
├─internal              // 应用私有包
│  ├─app                // 实际应用目录
│  │  └─web             // web应用示例
│  ├─errcode            // 应用公用错误码示例
│  ├─queue              // 队列使用示例
│  └─timer              // 定时器使用示例
└─pkg                   // 公共包
    ├─base              // 基础包
    │  ├─constant       // 常量定义
    │  ├─errcode        // 错误码
    │  └─response       // 返回响应方法
    ├─config            // 配置
    ├─log               // 日志
    ├─middleware        // 通用中间件
    ├─queue             // 队列
    │  └─rocketmq       // rocket mq
    ├─storage           // 存储
    ├─timer             // 定时器
    ├─utils             // 工具方法
    └─webapp            // 应用启动入口
```
