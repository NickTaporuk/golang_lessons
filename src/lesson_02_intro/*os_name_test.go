package lesson_02_intro

import (
	"runtime"
	"testing"
	"fmt"
)
func TestGetOsType(t *testing.T) {
	excepted := runtime.GOOS
	actual := GetOsType()

	fmt.Printf("%q != %q", actual, excepted)
	if excepted != actual {
		t.Errorf("%q != %q", actual, excepted)
	}
}
