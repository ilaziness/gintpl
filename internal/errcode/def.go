// Copyright (c) 2023 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

package errcode

import (
	"net/http"
)

// 错误码定义

var (
	CodeNil      = NewCode(0, "ok", http.StatusOK)
	CodeNotFound = NewCode(404, "Not Found", http.StatusOK)
)
