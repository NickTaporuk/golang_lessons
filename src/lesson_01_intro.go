package main

import (
	"fmt"
	"runtime"
	"lesson_02_intro"

)

func main() {
	fmt.Println("Hello World \n")
	fmt.Println("My OS TYPE" + runtime.GOOS)
	fmt.Println("My OS TYPE" + lesson_02_intro.GetOsType())
}
