package models

type CustomError struct {
	statusCode int // 小寫讓 JSON 忽略
	Title      string
	Message    string
}

func (err *CustomError) Error() string {
	return err.Message
}
