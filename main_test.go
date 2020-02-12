package GoUnitTest

import "testing"

func Test_Add(t *testing.T){
	var a uint
	var b uint
	a = 3
	b = 3

	if add(a,b) != 7{
		T.Error("Add failed")
	}
}
