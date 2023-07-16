package errcode

import (
	"fmt"
	"net/http"
)

type ErrCode int64

const (
	ERR_CODE_OK              ErrCode = 200 //OK
	ERR_CODE_INVALID_PARAMS  ErrCode = 400 // 无效参数
	ERR_CODE_INTERNAL_ERROR  ErrCode = 500 // 超时
	ERR_CODE_NOT_FOUND_TOKEN ErrCode = 405 //找不到token
)

type Error struct {
	code    int
	msg     string
	details []string
}

type ErrorV2 struct {
	code    ErrCode
	msg     string
	details []string
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

func NewErrorV2(code ErrCode, str ...string) *ErrorV2 {
	err := &ErrorV2{
		code: code,
		msg:  code.String(),
	}
	err.details = str
	return err
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)
	return &newError
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码:%d,错误信息: %s", e.code, e.msg)
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	case NotFound.Code():
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
