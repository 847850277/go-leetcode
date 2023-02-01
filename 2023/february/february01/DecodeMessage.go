package main

func main() {
	key := "123"
	message := "456"
	s := decodeMessage(key, message)
	println(s)
}

func decodeMessage(key string, message string) string {
	cur := byte('a')
	rules := map[rune]byte{}

	for _, c := range key {
		if c != ' ' && rules[c] == 0 {
			rules[c] = cur
			cur++
		}
	}

	m := []byte(message)
	for i, c := range message {
		if c != ' ' {
			m[i] = rules[c]
		}
	}

	return string(m)
}
