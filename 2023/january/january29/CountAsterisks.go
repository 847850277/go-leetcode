package main

import "fmt"

func main() {
	//var s string = "l|*e*et|c**o|*de|"
	var s string = "yo|uar|e**|b|e***au|tifu|l"
	fmt.Print(countAsterisks(s))
}

func countAsterisks(s string) (res int) {
	valid := true
	for _, c := range s {
		if c == '|' {
			valid = !valid
		} else if c == '*' && valid {
			res++
		}
	}
	return res
}
