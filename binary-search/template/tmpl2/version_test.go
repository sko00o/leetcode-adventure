package tmpl2

import (
	"fmt"
	"testing"
)

func TestIsBadVersion(t *testing.T) {
	tcs := []struct {
		n      int
		ver    int
		expect int
	}{
		{5, 5, 5},
		{5, 4, 4},
		{5, 3, 3},
		{5, 2, 2},
		{5, 1, 1},
		{6, 6, 6},
		{6, 5, 5},
		{6, 4, 4},
		{6, 3, 3},
		{6, 2, 2},
		{6, 1, 1},

		{6, -1, 1},
		{6, 7, -1},
	}

	for i, tc := range tcs {
		t.Run(fmt.Sprintf("task #%d", i), func(t *testing.T) {
			badVersion = tc.ver
			if got := firstBadVersion(tc.n); got != tc.expect {
				t.Errorf("got: %d, expect; %d", got, tc.expect)
			}
		})
	}
}
