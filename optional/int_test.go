package optional

import (
	"testing"
)

func TestIntPresent(t *testing.T) {
	var emptyInt Int
	if actual := emptyInt.Present(); actual != false {
		t.Error("Int{nil}#Present should return false")
	}

	if actual := NewInt(0).Present(); actual != true {
		t.Error("Int{0}#Present should return true")
	}

	if actual := NewInt(-1).Present(); actual != true {
		t.Error("Int{-1}#Present should return true")
	}
}

func TestIntGet(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Int{nil}#Get should cause panic")
			}
		}()
		var emptyInt Int
		emptyInt.Get()
	}()

	if actual := NewInt(0).Get(); actual != 0 {
		t.Error("Int{0}#Get should return ''")
	}

	if actual := NewInt(-1).Get(); actual != -1 {
		t.Error("Int{-1}#Get should return 'null'")
	}
}
