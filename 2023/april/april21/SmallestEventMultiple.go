package main

func main() {
	n := 5
	println(smallestEvenMultiple(n))
}

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return n * 2
}
