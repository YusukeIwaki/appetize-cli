package optional

type Bool struct {
	val *bool
}

func (opt Bool) Present() bool {
	return opt.val != nil
}

func (opt Bool) GetE() (val bool, err error) {
	if opt.Present() {
		val = *opt.val
	} else {
		err = valueIsNotSetError{}
	}
	return
}

func (opt Bool) Get() bool {
	if val, err := opt.GetE(); err == nil {
		return val
	} else {
		panic(err.Error())
	}
}

func NewBool(val bool) Bool {
	return Bool{&val}
}
