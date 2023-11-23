package main

import "strings"

func main() {
	s := "&amp; is an HTML entity but &ambassador; is not."
	print(entityParser(s))
}

func entityParser(text string) string {
	replaceMap := map[string]string{
		"&quot;":  "\"",
		"&apos;":  "'",
		"&amp;":   "&",
		"&gt;":    ">",
		"&lt;":    "<",
		"&frasl;": "/",
	}
	var ans = strings.Builder{}
	length := len(text)
	i := 0
	for i < length {
		found := false
		for l := 1; l < 8; l++ {
			j := i + l
			if j <= length {
				t := text[i:j]
				if val, ok := replaceMap[t]; ok {
					ans.WriteString(val)
					i = j
					found = true
					break
				}
			}
		}
		if !found {
			ans.WriteByte(text[i])
			i++
		}
	}
	return ans.String()
}
