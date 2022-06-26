package entity

type Error struct {
	Err  error
	Code int
}

func (e *Error) Error() string {
	return e.Err.Error()
}
