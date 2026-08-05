package ginc

import "fmt"

type BizError struct {
	Code    string
	Message string
}

func (e *BizError) Error() string {
	return fmt.Sprintf(`{"code":"%s","message":"%s"}`, e.Code, e.Message)
}

func NewBizError(code, message string) *BizError {
	return &BizError{
		Code:    code,
		Message: message,
	}
}

func NewCommBizErr(message string) *BizError {
	return NewBizError("1", message)
}
