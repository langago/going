package strings

import "strings"

func HasSuffix(s string, suffixes ...string) bool {
	var b bool
	for _, suffix := range suffixes {
		b = strings.HasSuffix(s, suffix)
		if b {
			return b
		}
	}
	return b
}

func LowerFirstChar(s string) string {
	chs := []byte(s)
	if chs[0] > 64 && chs[0] < 91 {
		chs[0] = chs[0] + 32
		return string(chs)
	}
	return s
}