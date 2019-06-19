package ll3

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// test in case `find` function is not right
func Test_find(t *testing.T) {
	Convey("make a linked list", t, func() {
		ll := Constructor()

		Convey("add numbers to tail", func() {
			for i := 0; i < 11; i++ {
				ll.AddAtTail(i)
			}

			Convey("then find them", func() {
				for i := 0; i < 11; i++ {
					So(ll.find(i).data, ShouldEqual, i)
				}
			})
		})

		Convey("add numbers to head", func() {
			l := 11

			for i := 0; i < l; i++ {
				ll.AddAtHead(i)
			}

			Convey("then find them", func() {
				for i := 0; i < l; i++ {
					So(ll.find(i).data, ShouldEqual, l-1-i)
				}
			})
		})

		Convey("add numbers at index", func() {
			l := 11

			for i := 0; i < l; i++ {
				ll.AddAtIndex(i, i)
			}

			Convey("then find them", func() {
				for i := 0; i < l; i++ {
					So(ll.find(i).data, ShouldEqual, i)
				}
			})

			for i := 0; i < l; i++ {
				ll.AddAtIndex(l, i)
			}
		})
	})
}
