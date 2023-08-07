// Copyright (c) 2023 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

package middleware

import (
	"context"
	"errors"
	"fish/internal/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResponseFormat struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type Context struct {
	Ctx *context.Context
	*gin.Context
}

type Handler func(c *Context) (any, error)

func Handle(h Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		myctx := &Context{
			Ctx:     &ctx,
			Context: c,
		}
		result, err := h(myctx)
		if err != nil {
			var errCode errcode.ErrCode
			if errors.As(err, &errCode) {
				c.JSON(errCode.StatusCode, ResponseFormat{
					errCode.Code,
					errCode.Message,
					nil,
				})
				return
			}
			c.JSON(http.StatusInternalServerError, ResponseFormat{
				http.StatusInternalServerError,
				http.StatusText(http.StatusInternalServerError),
				nil,
			})
			return
		}
		c.JSON(http.StatusOK, result)
	}
}
