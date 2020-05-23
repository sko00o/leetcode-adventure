package conclusion

import "testing"

func Test_kthGrammar(t *testing.T) {
	type T struct {
		N, K, want int
	}

	var tests = []T{
		{1, 1, 0},
		{2, 1, 0},
		{2, 2, 1},
		{3, 1, 0},
		{3, 2, 1},
		{3, 3, 1},
		{3, 4, 0},
		{4, 1, 0},
		{4, 2, 1},
		{4, 3, 1},
		{4, 4, 0},
		{4, 5, 1},
		{4, 6, 0},
		{4, 7, 0},
		{4, 8, 1},
	}
	for idx, tt := range tests {
		if got := kthGrammar(tt.N, tt.K); got != tt.want {
			t.Errorf("case#%d: kthGrammar() = %v, want %v", idx, got, tt.want)
		}
	}
}
