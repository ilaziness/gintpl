# 定时任务使用方法

## 1. 实现`Jober`接口，注册任务

例如：

```go
// 注册任务
func init() {
	timer.RegisterJob(&TestTimer{})
}

// 实现Jober接口
type TestTimer struct {
}

func (t *TestTimer) GetName() string {
	return "TestTimer"
}

func (t *TestTimer) GetCron() string {
	return "*/1 * * * *"
}

func (t *TestTimer) Run() {
	log.Println("TestTimer is run")
}
```

## 2. 在入口出引入即可

例: `import _ "gintpl/internal/timer"`