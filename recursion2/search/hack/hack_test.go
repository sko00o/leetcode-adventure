package hack

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hack(t *testing.T) {
	tests := []struct {
		input []byte
		out   []byte
	}{
		{
			input: []byte(`[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]
5
[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]
20
[[-5]]
-5
`),
			out: []byte(`true
false
true
`),
		},
		{
			input: []byte(`[[-5]]
-5
`),
			out: []byte(`true
`),
		},
	}

	for idx, tt := range tests {
		t.Run(fmt.Sprintf("case#%d", idx), func(t *testing.T) {
			in := bufio.NewScanner(bytes.NewReader(tt.input))
			got := bytes.NewBuffer(nil)
			out := bufio.NewWriter(got)
			hack(in, out)
			require.Equal(t, string(tt.out), got.String())
		})
	}
}
