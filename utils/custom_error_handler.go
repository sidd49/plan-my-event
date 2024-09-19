package utils

type CustomError struct {
	Err        error
	StatusCode int
}

func NewCustomError(statusCode int, err error) *CustomError {
	return &CustomError{
		Err:        err,
		StatusCode: statusCode,
	}
}
