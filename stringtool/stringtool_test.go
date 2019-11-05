package stringtool

import "testing"

func TestGetRandomString(t *testing.T) {
	size := 10
	str := GetRandomString(size)
	t.Logf("str:%v", str) // a string str:Y7FFBkWB7h
}
