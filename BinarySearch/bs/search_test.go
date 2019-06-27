package bs

import (
	"testing"

	"adventure/BinarySearch/bs/test"
)

func TestSearch(t *testing.T) {
	test.CommonTest(t, search)
}
