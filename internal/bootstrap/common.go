// Copyright (c) 2023 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

package bootstrap

type StartupFun func()
type DestroyFun func()

var startupList = make([]StartupFun, 0)
var destroyList = make([]DestroyFun, 0)
