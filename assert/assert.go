package assert

import (
	"reflect"
	"testing"
)

func AssertEqualExpect(expect, actual interface{}, test *testing.T) {
	if reflect.DeepEqual(expect, actual) {
		test.Logf("expect:%s==actual:%s,Success", expect, actual)
		return
	}
	test.Fatalf("expect:%s!=actual:%s,Failed", expect, actual)
}

// 断言assert
func AssertBool(ok bool, t *testing.T) {
	if !ok {
		t.Fatalf("Failed")
	} else {
		t.Logf("Success")
	}
}

func AssertError(err error, t *testing.T) {
	if err != nil {
		t.Fatalf(err.Error())
	} else {
		t.Logf("Success")
	}
}
