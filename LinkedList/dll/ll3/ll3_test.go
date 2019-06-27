package ll3

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"adventure/LinkedList/dll"
	"adventure/LinkedList/dll/test"
)

func TestLinkedList(t *testing.T) {
	f := func() dll.LinkedList {
		c := Constructor()
		return &c
	}
	test.CommonTest(t, f)
}

// test in case `find` function is not right
func TestExtra(t *testing.T) {
	Convey("make a linked list", t, func() {
		ll := Constructor()

		Convey("add 4 numbers to tail", func() {
			for i := 0; i < 4; i++ {
				ll.AddAtTail(i)
			}

			Convey("got node at index 2", func() {
				p := ll.head.next.next
				So(p.data, ShouldEqual, 2)

				Convey("insert 9 previous of index 2", func() {
					p.insertPrev(9)
					So(p.prev.data, ShouldEqual, 9)
					So(p.prev.next.data, ShouldEqual, 2)
				})

				Convey("insert 9 next of index 2", func() {
					p.insertNext(9)
					So(p.next.data, ShouldEqual, 9)
					So(p.next.prev.data, ShouldEqual, 2)
				})
			})
		})

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
