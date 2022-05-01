package models

type CustomError struct {
	StatusCode int // 小寫讓 JSON 忽略
	Title      string
	Message    string
}

func (err *CustomError) Error() string {
	return err.Message
}

func (err *CustomError) ToJSON() (obj map[string]interface{}) {

	// 若無自定義 Title 則根據 StatusCode 設定
	if err.Title == "" {
		err.fetchTitle()
	}

	return map[string]interface{}{
		"title":   err.Title,
		"message": err.Message,
	}
}

func (err *CustomError) fetchTitle() {

	switch err.StatusCode {
	case 400:
		err.Title = "Bad Request"
	case 401:
		err.Title = "Unauthorized"
	case 403:
		err.Title = "Forbidden"
	case 404:
		err.Title = "Not Found"
	case 422:
		err.Title = "Unprocessable Entity"
	case 500:
		err.Title = "Internal Server Error"
	default:
		err.Title = "Unknown Error"
	}
}
