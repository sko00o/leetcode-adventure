package conclusion

import "testing"

func Test_nextGreatestLetter(t *testing.T) {
	type args struct {
		letters []byte
		target  byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			"case 1",
			args{
				[]byte{'c', 'f', 'j'},
				'c',
			},
			'f',
		},
		{
			"case 2",
			args{
				[]byte{'c', 'f', 'j'},
				'd',
			},
			'f',
		},
		{
			"case 3",
			args{
				[]byte{'c', 'f', 'j'},
				'g',
			},
			'j',
		},
		{
			"case 4",
			args{
				[]byte{'c', 'f', 'j'},
				'j',
			},
			'c',
		},
		{
			"case 5",
			args{
				[]byte{'c', 'f', 'j'},
				'k',
			},
			'c',
		},
		{
			"case 6",
			args{
				[]byte{'c', 'c', 'c', 'c', 'f', 'f', 'g'},
				'c',
			},
			'f',
		},
		{
			"case 7",
			args{
				[]byte{'e', 'e', 'e', 'e', 'e', 'e', 'n', 'n', 'n', 'n'},
				'n',
			},
			'e',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreatestLetter(tt.args.letters, tt.args.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
