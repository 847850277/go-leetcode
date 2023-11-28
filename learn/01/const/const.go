package main

import "fmt"

const (
	a         int    = 1
	b         string = "vvv"
	ERROR_1          = iota
	ERROR_2          = iota
	ERROR_3          = iota
	ERROR_4          = iota
	ERROR_5          = iota
	ERROR_6          = iota
	ERROR_MES        = "ERROR"
	ERROR            = iota
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(ERROR_1, ERROR_2, ERROR_3, ERROR_4, ERROR_5, ERROR_6, ERROR_MES, ERROR)

}
