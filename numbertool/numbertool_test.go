package numbertool

import "testing"

func TestIsNumeric(t *testing.T) {
	num1 := 0.99
	boo1 := IsNumeric(num1)
	t.Logf("boo1:%v", boo1) // true
	num2 := 'a'
	boo2 := IsNumeric(num2)
	t.Logf("boo2:%v", boo2) // true
	num3 := "7.823E5"
	boo3 := IsNumeric(num3)
	t.Logf("boo3:%v", boo3) // true
	num4 := "0x0012e"
	boo4 := IsNumeric(num4)
	t.Logf("boo4:%v", boo4) // true
}

func TestIsInclouldArray(t *testing.T) {
	arr1 := []int{1,2,3,4}
	arr2 := []int{1,2,3,4,5}
	boo := IsInclouldArray(arr1, arr2)
	t.Logf("boo:%v", boo) // true
}

