package response

import (
	"github.com/sirupsen/logrus"
)

func Err(err error) *Error {
	switch err.(type) {
	case Error:
		e := err.(Error)
		return &e
	case *Error:
		return err.(*Error)
	default:
		logrus.Errorf("unknown error: %v", err)
		return NewUnknownError(err)
	}
}

func Ok(data interface{}) *Error {
	return &Error{
		Code: 200,
		Msg:  "success",
		Data: data,
	}
}
