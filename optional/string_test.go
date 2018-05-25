package optional

import (
	"testing"
)

func TestStringPresent(t *testing.T) {
	var emptyString String
	if actual := emptyString.Present(); actual != false {
		t.Error("String{nil}#Present should return false")
	}

	if actual := NewString("").Present(); actual != true {
		t.Error("String{''}#Present should return true")
	}

	if actual := NewString("null").Present(); actual != true {
		t.Error("String{'null'}#Present should return true")
	}

	if actual := NewString("0").Present(); actual != true {
		t.Error("String{'0'}#Present should return true")
	}

	if actual := NewString("false").Present(); actual != true {
		t.Error("String{'false'}#Present should return true")
	}
}

func TestStringGet(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("String{nil}#Get should cause panic")
			}
		}()
		var emptyString String
		emptyString.Get()
	}()

	if actual := NewString("").Get(); actual != "" {
		t.Error("String{''}#Get should return ''")
	}

	if actual := NewString("null").Get(); actual != "null" {
		t.Error("String{'null'}#Get should return 'null'")
	}

	if actual := NewString("0").Get(); actual != "0" {
		t.Error("String{'0'}#Get should return '0'")
	}

	if actual := NewString("false").Get(); actual != "false" {
		t.Error("String{'false'}#Get should return 'false'")
	}
}
