package rec

import (
	"bytes"
	"testing"
)

type args struct {
	s []byte
}

type testCase struct {
	name   string
	args   args
	expect []byte
}

var tests = []testCase{
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
		"empty",
		args{
			[]byte{},
		},
		[]byte{},
	},
}

func Test_reverseString(t *testing.T) {
	for fIdx, f := range []func([]byte){
		reverseString,
		reverseString0,
		reverseString1,
	} {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				ts := make([]byte, len(tt.args.s))
				copy(ts, tt.args.s)
				f(ts)
				if !bytes.Equal(ts, tt.expect) {
					t.Errorf("func #%d, got: %s, expect: %s", fIdx, ts, tt.expect)
				}
			})
		}
	}
}

var unicodeTests = []testCase{
	{
		"chinese",
		args{
			[]byte("大家好"),
		},
		[]byte("好家大"),
	},
	{
		"emoji",
		args{
			[]byte("🚭😀🙌🙊"),
		},
		[]byte("🙊🙌😀🚭"),
	},
}

func Test_reverseString_unicode(t *testing.T) {
	for fIdx, f := range []func([]byte){
		reverseString2,
		reverseString3,
	} {
		for _, tt := range append(tests, unicodeTests...) {
			t.Run(tt.name, func(t *testing.T) {
				ts := make([]byte, len(tt.args.s))
				copy(ts, tt.args.s)
				f(ts)
				if !bytes.Equal(ts, tt.expect) {
					t.Errorf("func #%d, got: %s, expect: %s", fIdx, ts, tt.expect)
				}
			})
		}
	}
}
