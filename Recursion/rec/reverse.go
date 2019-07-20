package rec

func reverseString(s []byte) {
	m := len(s)
	if m > 1 {
		s[0], s[m-1] = s[m-1], s[0]
		reverseString(s[1 : m-1])
	}
}
