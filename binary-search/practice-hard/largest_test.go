package practice_hard

import "testing"

func Test_splitArray(t *testing.T) {
	type args struct {
		nums []int
		m    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"ex0",
			args{
				[]int{7, 2, 5, 10, 8},
				1,
			},
			32,
		},
		{
			"ex1",
			args{
				[]int{7, 2, 5, 10, 8},
				2,
			},
			18,
		},
		{
			"ex2",
			args{
				[]int{7, 2, 5, 10, 8},
				5,
			},
			10,
		},
		{
			"ex3",
			args{
				[]int{1, 2147483647},
				2,
			},
			2147483647,
		},
		{
			"ex4",
			args{
				[]int{2, 3, 1, 2, 4, 3},
				5,
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitArray(tt.args.nums, tt.args.m); got != tt.want {
				t.Errorf("splitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitIndexes(t *testing.T) {
	tasks := []struct {
		n, m      int
		expectLen int
	}{

		{
			5, 1,
			0, // C(4,0)
		},
		{
			5, 2,
			4, // C(4,1)
		},
		{
			5, 3,
			6, // C(4,2)
		},
		{
			5, 4,
			4, // C(4,3)
		},
		{
			5, 5,
			1, // C(4,4)
		},
		{
			5, 0,
			0, // invalid
		},
		{
			5, 6,
			0, // invalid
		},
		/*
			{
				10, 3,
				36, // C(9,2) = 9*8/2=36
			},
			{
				1000, 3,
				498501, // C(999,2) = 999*998/2
			},
		*/
	}

	for _, tt := range tasks {
		if got := splitIndexes(tt.n, tt.m); len(got) != tt.expectLen {
			t.Errorf("got: %v, expectLen: %d", got, tt.expectLen)
		} else {
			t.Log(got)
		}
	}
}
