package errors

type Error struct {
	Message []string `json:"message"`
}

func NewError(message string) Error {
	var error Error
	error.Message = append(error.Message, message)
	return error
}
