package hook

type Closure func()

type exitType struct {
	hooks []Closure
}

// Exit 应用退出钩子
var Exit *exitType

func init() {
	Exit = &exitType{}
}

func (e *exitType) Register(hook Closure) {
	e.hooks = append(e.hooks, hook)
}

func (e *exitType) Trigger() {
	for _, hook := range e.hooks {
		hook()
	}
}
