package conclusion

import "strings"

func decodeString(s string) string {
	var stk []byte
	for i := range s {
		if s[i] == ']' {
			si := strings.LastIndex(string(stk), "[")
			var n, ni int
			ti := 1
			for ni = si - 1; ni >= 0; ni-- {
				if d := int(stk[ni] - '0'); d >= 0 && d <= 9 {
					n += d * ti
					ti *= 10
				} else {
					break
				}
			}
			ns := genStr(n, stk[si+1:])
			stk = append(stk[:ni+1], ns...)
		} else {
			stk = append(stk, s[i])
		}
	}
	return string(stk)
}

func genStr(n int, s []byte) (out []byte) {
	for i := n; i > 0; i-- {
		out = append(out, s...)
	}
	return
}

func decodeString1(s string) string {
	var iStk []int
	var nStk []int
	var out []byte
	var n int
	for i := range s {
		if s[i] == '[' {
			iStk = append(iStk, len(out))
			nStk = append(nStk, n)
			n = 0
		} else if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int(s[i]-'0')
		} else if s[i] == ']' {
			idx := iStk[len(iStk)-1]
			iStk = iStk[:len(iStk)-1]
			subS := out[idx:]
			rep := nStk[len(nStk)-1]
			nStk = nStk[:len(nStk)-1]
			out = append(out[:idx], genStr(rep, subS)...)
		} else {
			out = append(out, s[i])
		}
	}
	return string(out)
}
