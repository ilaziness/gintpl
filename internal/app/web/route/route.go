// Copyright (c) 2023 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

package route

import (
	"gintpl/internal/app/web/handler"
	"gintpl/internal/app/web/handler/user"

	"github.com/gin-gonic/gin"
)

func InitRoute(e *gin.Engine) {
	e.GET("/", handler.Index)
	e.GET("/p", handler.Painc)
	e.GET("/user/:id", user.Index)

	e.GET("/send_mq", handler.SendMq)
	e.GET("/trace_test", handler.Trace)
	e.GET("/ser_dis", handler.ServiceDis)
}
