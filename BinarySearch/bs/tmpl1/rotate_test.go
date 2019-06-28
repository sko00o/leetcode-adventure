package tmpl1

import (
	"testing"

	"adventure/BinarySearch/bs/test"
)

func TestSearch(t *testing.T) {
	test.RotateTest(t, search)
}
