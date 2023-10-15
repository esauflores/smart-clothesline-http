package helpers

type HTTPError struct {
	StatusCode int
	Message    string
}

func (ce HTTPError) Error() string {
	return ce.Message
}

func Fatal(code int, err error) {
	panic(HTTPError{code, err.Error()})
}

func CheckFatal(checked error, code int, err error) {
	if checked != nil {
		panic(HTTPError{code, err.Error()})
	}
}
