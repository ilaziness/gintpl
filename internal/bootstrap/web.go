// Copyright (c) 2023 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

package bootstrap

// 注册应用启动需要初始化的资源
// 注册应用退出需要的销毁或收尾

func RegisterStartup(upfunc StartupFun) {
	startupList = append(startupList, upfunc)
}

func RegisterDestructor(destructor DestroyFun) {
	destroyList = append(destroyList, destructor)
}

// RunInit 运行注册的启动函数
// 会在整个应用初始化完成，运行应用之前执行本函数
func RunInit() {
	for _, f := range startupList {
		f()
	}
}

// RunDestructor 执行销毁
// 会在应用退出之前运行
func RunDestructor() {
	for _, f := range destroyList {
		f()
	}
}
