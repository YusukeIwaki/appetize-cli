package optional

type String struct {
	val *string
}

func (opt String) Present() bool {
	return opt.val != nil
}

func (opt String) GetE() (val string, err error) {
	if opt.Present() {
		val = *opt.val
	} else {
		err = valueIsNotSetError{}
	}
	return
}

func (opt String) Get() string {
	if val, err := opt.GetE(); err == nil {
		return val
	} else {
		panic(err.Error())
	}
}

func NewString(val string) String {
	return String{&val}
}
