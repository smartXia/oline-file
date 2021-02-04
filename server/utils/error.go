package utils

type CustomError struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

func NewCustomError(msg string, code int) *CustomError {
	return &CustomError{
		Msg:  msg,
		Code: code,
	}
}

func (err *CustomError) Error() string {
	return err.Msg
}
