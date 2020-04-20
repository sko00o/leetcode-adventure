package principle

func reverseString(s []byte) {
	m := len(s)
	if m > 1 {
		s[0], s[m-1] = s[m-1], s[0]
		reverseString(s[1 : m-1])
	}
}

// non-tail-recursion
func reverseString0(s []byte) {
	copy(s, rev(s))
}

func rev(s []byte) []byte {
	if s == nil || len(s) <= 1 {
		return s
	}
	return append(rev(s[1:]), s[0])
}

// tail-recursive
func reverseString1(s []byte) {
	copy(s, rev1(s))
}

func rev1(s []byte) []byte {
	if s == nil || len(s) <= 1 {
		return s
	}
	return helper(s, nil)
}

func helper(s, acc []byte) []byte {
	l := len(s)
	if l == 0 {
		return []byte(string(acc))
	}
	return helper(s[1:], append([]byte{s[0]}, acc...))
}

// unicode support
func reverseString2(s []byte) {
	copy(s, rev2(s))
}

func rev2(s []byte) []byte {
	r := []rune(string(s))
	if r == nil || len(r) <= 1 {
		return []byte(string(r))
	}
	return helper2(r, nil)
}

func helper2(s, acc []rune) []byte {
	l := len(s)
	if l == 0 {
		return []byte(string(acc))
	}
	return helper2(s[1:], append([]rune{s[0]}, acc...))
}

func reverseString3(r []byte) {
	s := []rune(string(r))
	rev3(s)
	copy(r, []byte(string(s)))
}

func rev3(s []rune) {
	m := len(s)
	if m > 1 {
		s[0], s[m-1] = s[m-1], s[0]
		rev3(s[1 : m-1])
	}
}
