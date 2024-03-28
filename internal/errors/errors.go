package errors

type CustomErrors struct {
	Err  error
	Code int
}

func (e *CustomErrors) Error() string {
	return e.Err.Error()
}

func ErrDatabaseInternal(err error) error {
	return &CustomErrors{Err: err, Code: 500}
}

func ErrAlreadyExists(err error) error {
	return &CustomErrors{Err: err, Code: 409}
}