package utils

import (
	"fmt"
	"runtime"
)

// SafeGo 启动一个协程并处理panic
func SafeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				// 这里可以记录日志或者做其他的错误处理
				fmt.Printf("Recovered in SafeGo: %v\n", r)
				// 打印堆栈信息
				debugStack := make([]byte, 1024)
				n := runtime.Stack(debugStack, false)
				fmt.Printf("%s\n", debugStack[:n])
			}
		}()
		fn()
	}()
}
