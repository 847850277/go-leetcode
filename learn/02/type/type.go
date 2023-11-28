package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a byte
	a = 1
	fmt.Println(a)

	var istr = "12"
	myint, error := strconv.Atoi(istr)
	if error != nil {
		fmt.Println("convert error")
	}
	fmt.Println(myint)

	c := myint + 1
	fmt.Println(c)

}
