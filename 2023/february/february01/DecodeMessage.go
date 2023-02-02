package main

import "fmt"

func main() {
	key := "the quick brown fox jumps over the lazy dog"
	message := "vkbs bs t suepuv"
	s := decodeMessage(key, message)
	fmt.Println(s)
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
