package april01

import (
	"strings"
	"unicode"
)

func maskPII(s string) string {
	at := strings.Index(s, "@")
	if at > 0 {
		s = strings.ToLower(s)
		return strings.ToLower(string(s[0])) + "*****" + s[at-1:]
	}
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if unicode.IsDigit(rune(c)) {
			sb.WriteByte(c)
		}
	}
	s = sb.String()
	country := []string{"", "+*-", "+**-", "+***-"}
	return country[len(s)-10] + "***-***-" + s[len(s)-4:]
}
