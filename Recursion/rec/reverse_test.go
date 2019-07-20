package rec

import (
	"bytes"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name   string
		args   args
		expect []byte
	}{
		{
			"example1",
			args{
				[]byte("hello"),
			},
			[]byte("olleh"),
		},
		{
			"example2",
			args{
				[]byte("Hannah"),
			},
			[]byte("hannaH"),
		},
		{
			"not change",
			args{
				[]byte("a"),
			},
			[]byte("a"),
		},
		{
			"enpty",
			args{
				[]byte{},
			},
			[]byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !bytes.Equal(tt.args.s, tt.expect) {
				t.Errorf("got: %s, expect: %s", tt.args.s, tt.expect)
			}
		})
	}
}
