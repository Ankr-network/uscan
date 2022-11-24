package response

import "fmt"

type Error struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Msg)
}

const (
	unknown           = 20001
	invalidParameter  = 10001
	recordNotFindErr  = 10002
	contractVerityErr = 10003
	exportNumErr      = 10004
)

func NewUnknownError(err error) *Error {
	return &Error{
		Code: unknown,
		Msg:  err.Error(),
	}
}

var (
	ErrInvalidParameter = &Error{
		Code: invalidParameter,
		Msg:  "invalid parameter",
	}

	ErrRecordNotFind = &Error{
		Code: recordNotFindErr,
		Msg:  "record not find",
	}
)

var (
	ErrVerityContract = &Error{
		Code: contractVerityErr,
		Msg:  "",
	}

	ErrExportNum = &Error{
		Code: exportNumErr,
		Msg:  "the number of exported items cannot exceed 10000.",
	}
)
