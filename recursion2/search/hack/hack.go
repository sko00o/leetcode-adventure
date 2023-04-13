package hack

import (
	"bufio"
	"fmt"
	// "os"
	"strconv"
)

//// Uncomment following lines on LeetCode
//	func init() {
//		in := bufio.NewScanner(os.Stdin)
//		in.Buffer(nil, 1e9)
//		fout, _ := os.Create("user.out")
//		defer fout.Close()
//		out := bufio.NewWriter(fout)
//		defer out.Flush()
//		hack(in, out)
//		os.Exit(0)
//	}
func hack(in *bufio.Scanner, out *bufio.Writer) {
	defer out.Flush()
next:
	for in.Scan() {
		s := in.Text() // matrix string
		in.Scan()
		target, err := strconv.Atoi(in.Text())
		if err != nil {
			panic(err)
		}
		for i := 1; i < len(s); i++ {
			isNeg := false
			if s[i] == '-' {
				i++
				isNeg = true
			}
			val, oldI := 0, i
			// convert number charators only
			for i < len(s) && int(s[i]&15) < 10 {
				val = val*10 + int(s[i]&15)
				i++
			}
			if isNeg {
				val = -val
			}
			if oldI < i && val == target {
				fmt.Fprintln(out, "true")
				continue next
			}
		}
		fmt.Fprintln(out, "false")
	}
}

// Just do nothing
func searchMatrix(_ [][]int, _ int) (_ bool) { return }
