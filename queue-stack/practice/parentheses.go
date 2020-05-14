package practice

func isValid(s string) bool {
	var stack []rune

outer:
	for _, c := range s {
		for _, p := range [][2]rune{
			{'(', ')'},
			{'[', ']'},
			{'{', '}'},
		} {
			if c == p[1] {
				if len(stack) == 0 {
					return false
				}

				if stack[len(stack)-1] != p[0] {
					return false
				}

				stack = stack[:len(stack)-1]
				continue outer
			}
		}

		stack = append(stack, c)
	}

	return len(stack) == 0
}
