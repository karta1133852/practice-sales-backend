package models

type CustomError struct {
	statusCode int // JSON 忽略小寫
	Title      string
	Message    string
}

func (err *CustomError) Error() string {
	return err.Message
}
