package problems

import (
	"testing"
)

func Test_maxFreeTime(t *testing.T) {
	type args struct {
		eventTime int
		startTime []int
		endTime   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				eventTime: 5,
				startTime: []int{1, 3},
				endTime:   []int{2, 5},
			},
			want: 2,
			/*
				gap         ev
				1   0,1
				            1,2   1
				1   2,3
				            3,5   2
				0   5,5
			*/
		},
		{
			name: "test2",
			args: args{
				eventTime: 10,
				startTime: []int{0, 7, 9},
				endTime:   []int{1, 8, 10},
			},
			want: 7,
		},
		{
			name: "test3",
			args: args{
				eventTime: 10,
				startTime: []int{0, 3, 7, 9},
				endTime:   []int{1, 4, 8, 10},
			},
			want: 6,
			/*
				gap         ev
				0   0,0
				            0,1   1
				2   1,3
				            3,4   1
				3   4,7
				            7,8   1
				1   8,9
				            9,10  1
				0   10,10
			*/
		},
		{
			name: "test4",
			args: args{
				eventTime: 5,
				startTime: []int{0, 1, 2, 3, 4},
				endTime:   []int{1, 2, 3, 4, 5},
			},
			want: 0,
		},
		{
			name: "test5",
			args: args{
				eventTime: 86,
				startTime: []int{22, 82},
				endTime:   []int{66, 85},
			},
			want: 38,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFreeTime(tt.args.eventTime, tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("maxFreeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
