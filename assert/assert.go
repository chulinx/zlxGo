package assert

import "testing"

func AssertEqualExpect(expect,actual interface{},test *testing.T)  {
	if expect == actual {
		test.Logf("expect:%s==actual%s,Success",expect,actual)
		return
	}
	test.Fatalf("expect:%s!=actual%s,Failed")
}

// 断言assert
func AssertBool(ok bool, t *testing.T) {
	if !ok {
		t.Fatalf("Failed")
	} else {
		t.Logf("Success")
	}
}

func AssertError(ok bool, err error, t *testing.T) {
	if !ok {
		t.Fatalf(err.Error())
	} else {
		t.Logf("Success")
	}
}
