package problems

// Runtime: 4 ms
// Memory Usage: 5.1 MB
func deserialize(s string) *NestedInteger {
	var stk []*NestedInteger
	var iStk []int

	var n int
	var hasNum, isNeg bool
	for i := range s {
		if s[i] == '[' {
			iStk = append(iStk, len(stk))
		} else if s[i] == '-' {
			isNeg = true
		} else if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int(s[i]-'0')
			hasNum = true
		} else if s[i] == ',' || s[i] == ']' {
			if hasNum {
				if isNeg {
					n = -n
				}
				var ni NestedInteger
				ni.SetInteger(n)
				stk = append(stk, &ni)

				isNeg = false
				hasNum = false
				n = 0
			}
			if s[i] == ']' {
				if len(iStk) == 0 {
					continue
				}
				// pop last '[' index
				idx := iStk[len(iStk)-1]
				iStk = iStk[:len(iStk)-1]

				var nni NestedInteger
				if idx < len(stk) {
					sstk := stk[idx:]
					for i := range sstk {
						nni.Add(*sstk[i])
					}
				}

				stk = append(stk[:idx], &nni)
			}
		}
	}

	if hasNum {
		if isNeg {
			n = -n
		}
		var ni NestedInteger
		ni.SetInteger(n)
		stk = append(stk, &ni)
	}

	if len(stk) != 1 {
		return nil
	}

	return stk[0]
}

// Runtime: 4 ms
// Memory Usage: 5.2 MB
func deserialize1(s string) *NestedInteger {
	if len(s) == 0 {
		return new(NestedInteger)
	}

	if len(s) > 1 && s[0] == '[' {
		_, ans := parseList([]byte(s), 1)
		return ans
	}

	_, ans := parseInt([]byte(s), 0)
	return ans
}

func parseList(s []byte, i int) (int, *NestedInteger) {
	var root NestedInteger
	var ni *NestedInteger
	for i < len(s) {
		if s[i] == '[' {
			i, ni = parseList(s, i+1)
			root.Add(*ni)
		} else if s[i] == '-' || (s[i] >= '0' && s[i] <= '9') {
			i, ni = parseInt(s, i)
			root.Add(*ni)
		} else if s[i] == ']' {
			i++
			break
		} else {
			i++
		}
	}

	return i, &root
}

func parseInt(s []byte, i int) (int, *NestedInteger) {
	var n int
	var isNeg bool
	for ; i < len(s); i++ {
		if s[i] == '-' {
			isNeg = true
		} else if s[i] >= '0' && s[i] <= '9' {
			n = n*10 + int(s[i]-'0')
		} else {
			break
		}
	}

	if isNeg {
		n = -n
	}
	var ni NestedInteger
	ni.SetInteger(n)
	return i, &ni
}
