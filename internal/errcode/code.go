// Copyright (c) 2023 ilaziness. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
//
// Author: ilaziness  https://github.com/ilaziness

package errcode

type ErrCode struct {
	Code       int
	Message    string
	StatusCode int
}

func (ec ErrCode) Error() string {
	return ec.Message
}

func NewCode(code int, msg string, statusCode int) ErrCode {
	return ErrCode{
		code,
		msg,
		statusCode,
	}
}
