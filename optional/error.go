package optional

type valueIsNotSetError struct {
}

func (err valueIsNotSetError) Error() string {
	return "value is absent"
}
