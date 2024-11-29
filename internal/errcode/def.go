// Copyright (c) 2023 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

package errcode

import (
	"github.com/ilaziness/gokit/base/errcode"
)

// 错误码定义

var (
	CodeNil      = errcode.NewCode(0, "ok")
	CodeNotFound = errcode.NewCode(404, "Not Found")

	CodeDBCreateFailed = errcode.NewCode(500, "DB Create Failed")
)
