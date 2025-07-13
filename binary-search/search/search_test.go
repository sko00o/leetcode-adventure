package search

import (
	"testing"

	"github.com/sko00o/leetcode-adventure/binary-search/test"
)

func TestSearch(t *testing.T) {
	test.CommonTest(t, search, search0, search1)
}
