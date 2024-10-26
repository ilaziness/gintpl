// Copyright (c) 2024 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

package errcode

import "fmt"

type Code struct {
	Code    int
	Message string
	Data    any
}

func NewCode(code int, msg string) *Code {
	return &Code{
		Code:    code,
		Message: msg,
		Data:    struct{}{},
	}
}

func (ec *Code) Error() string {
	return ec.Message
}

func (ec *Code) SetData(data any) {
	ec.Data = data
}

func (ec *Code) SetMessage(msg string) {
	ec.Message = msg
}

func (ec *Code) Param(a ...any) {
	ec.Message = fmt.Sprintf(ec.Message, a...)
}
