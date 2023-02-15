package february15

func main() {

}

func isGoodArray(array []int) bool {
	divisor := array[0]
	for i := 0; i < len(array); i++ {
		num := array[i]
		divisor = gcd(divisor, num)
		if divisor == 1 {
			break
		}
	}
	return divisor == 1
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
