package optional

type Int struct {
	val *int
}

func (opt Int) Present() bool {
	return opt.val != nil
}

func (opt Int) GetE() (val int, err error) {
	if opt.Present() {
		val = *opt.val
	} else {
		err = valueIsNotSetError{}
	}
	return
}

func (opt Int) Get() int {
	if val, err := opt.GetE(); err == nil {
		return val
	} else {
		return 0
	}
}

func NewInt(val int) Int {
	return Int{&val}
}
