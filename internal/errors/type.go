package errors

type MyError struct {
	Code   int
	Reason string
}

func (e MyError) Error() string {
	return e.Reason
}

func New(code int, reason string) error {
	return MyError{
		Code:   code,
		Reason: reason,
	}
}
