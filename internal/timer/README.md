自定义定时任务

用法：

先实现定时任务功能，必须实现接口`pkg/time/Jober`，再在入口导入`_ "gintpl/internal/timer"`

定时任务时间写法参考：https://pkg.go.dev/github.com/robfig/cron/v3