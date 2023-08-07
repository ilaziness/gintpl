// Copyright (c) 2023 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

package web

import (
	"fish/internal/handler"
	"fish/internal/handler/user"
)

func (a *Web) SetRoute() {
	a.Get("/", handler.Index)
	a.Get("/user/:id", user.Index)
}
