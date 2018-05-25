package optional

import (
	"testing"
)

func TestBoolPresent(t *testing.T) {
	var emptyBool Bool
	if actual := emptyBool.Present(); actual != false {
		t.Error("Bool{nil}#Present should return false")
	}

	if actual := NewBool(false).Present(); actual != true {
		t.Error("Bool{false}#Present should return true")
	}

	if actual := NewBool(true).Present(); actual != true {
		t.Error("Bool{true}#Present should return true")
	}
}

func TestBoolGet(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Bool{nil}#Get should cause panic")
			}
		}()
		var emptyBool Bool
		emptyBool.Get()
	}()

	if actual := NewBool(false).Get(); actual != false {
		t.Error("Bool{false}#Get should return false")
	}

	if actual := NewBool(true).Get(); actual != true {
		t.Error("Bool{true}#Get should return true")
	}
}
