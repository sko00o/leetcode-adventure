package linkedlist

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	. "linkedlist/ll3"
)

func TestLinkedList(t *testing.T) {
	Convey("construct one linked list", t, func() {
		obj := Constructor()
		Convey("get invalid index will return -1", func() {
			got := obj.Get(1)
			So(got, ShouldEqual, -1)
		})

		Convey("add numbers to tail", func() {
			obj.AddAtTail(1)
			obj.AddAtTail(2)
			obj.AddAtTail(3)
			obj.AddAtTail(4)
			obj.AddAtTail(5)
			So(obj.Get(0), ShouldEqual, 1)
			So(obj.Get(1), ShouldEqual, 2)
			So(obj.Get(2), ShouldEqual, 3)
			So(obj.Get(3), ShouldEqual, 4)
			So(obj.Get(4), ShouldEqual, 5)

			Convey("get invalid index will return -1", func() {
				So(obj.Get(-1), ShouldEqual, -1)
				So(obj.Get(5), ShouldEqual, -1)
			})
		})

		Convey("add numbers to head", func() {
			obj.AddAtHead(1)
			obj.AddAtHead(2)
			obj.AddAtHead(3)
			obj.AddAtHead(4)
			obj.AddAtHead(5)
			So(obj.Get(0), ShouldEqual, 5)
			So(obj.Get(1), ShouldEqual, 4)
			So(obj.Get(2), ShouldEqual, 3)
			So(obj.Get(3), ShouldEqual, 2)
			So(obj.Get(4), ShouldEqual, 1)

			Convey("get invalid index will return -1", func() {
				So(obj.Get(-1), ShouldEqual, -1)
				So(obj.Get(5), ShouldEqual, -1)
			})
		})

		Convey("add head tail head tail", func() {
			obj.AddAtHead(1)
			obj.AddAtTail(2)
			obj.AddAtHead(3)
			obj.AddAtTail(4)
			So(obj.Get(0), ShouldEqual, 3)
			So(obj.Get(1), ShouldEqual, 1)
			So(obj.Get(2), ShouldEqual, 2)
			So(obj.Get(3), ShouldEqual, 4)

			Convey("get invalid index will return -1", func() {
				So(obj.Get(-1), ShouldEqual, -1)
				So(obj.Get(4), ShouldEqual, -1)
			})
		})

		Convey("add tail head tail head", func() {
			obj.AddAtTail(1)
			obj.AddAtHead(2)
			obj.AddAtTail(3)
			obj.AddAtHead(4)
			So(obj.Get(0), ShouldEqual, 4)
			So(obj.Get(1), ShouldEqual, 2)
			So(obj.Get(2), ShouldEqual, 1)
			So(obj.Get(3), ShouldEqual, 3)

			Convey("get invalid index will return -1", func() {
				So(obj.Get(-1), ShouldEqual, -1)
				So(obj.Get(4), ShouldEqual, -1)
			})
		})

		Convey("add at index 0", func() {
			obj.AddAtIndex(0, 6)
			So(obj.Get(0), ShouldEqual, 6)

			Convey("get invalid index will return -1", func() {
				So(obj.Get(-1), ShouldEqual, -1)
				So(obj.Get(1), ShouldEqual, -1)
			})

			Convey("add at head", func() {
				obj.AddAtHead(1)
				So(obj.Get(0), ShouldEqual, 1)
				So(obj.Get(1), ShouldEqual, 6)

				Convey("get invalid index will return -1", func() {
					So(obj.Get(-1), ShouldEqual, -1)
					So(obj.Get(2), ShouldEqual, -1)
				})
			})

			Convey("add on index -1", func() {
				obj.AddAtIndex(-1, 8)
				So(obj.Get(0), ShouldEqual, 8)
				So(obj.Get(1), ShouldEqual, 6)
			})
		})

		Convey("add on index 0 then head", func() {
			obj.AddAtIndex(0, 7)
			So(obj.Get(0), ShouldEqual, 7)
			obj.AddAtHead(6)
			So(obj.Get(0), ShouldEqual, 6)
			So(obj.Get(1), ShouldEqual, 7)

			Convey("get invalid index will return -1", func() {
				So(obj.Get(-1), ShouldEqual, -1)
				So(obj.Get(2), ShouldEqual, -1)
			})
		})

		Convey("add on index 0 then tail", func() {
			obj.AddAtIndex(0, 7)
			So(obj.Get(0), ShouldEqual, 7)
			obj.AddAtTail(6)
			So(obj.Get(0), ShouldEqual, 7)
			So(obj.Get(1), ShouldEqual, 6)

			Convey("get invalid index will return -1", func() {
				So(obj.Get(-1), ShouldEqual, -1)
				So(obj.Get(2), ShouldEqual, -1)
			})
		})

		Convey("add three number", func() {
			obj.AddAtTail(1)
			obj.AddAtTail(2)
			obj.AddAtTail(3)

			Convey("then delete last one", func() {
				obj.DeleteAtIndex(2)
				So(obj.Get(0), ShouldEqual, 1)
				So(obj.Get(1), ShouldEqual, 2)

				Convey("get invalid index will return -1", func() {
					So(obj.Get(-1), ShouldEqual, -1)
					So(obj.Get(2), ShouldEqual, -1)
				})
			})

			Convey("then delete head one", func() {
				obj.DeleteAtIndex(0)
				So(obj.Get(0), ShouldEqual, 2)
				So(obj.Get(1), ShouldEqual, 3)

				Convey("get invalid index will return -1", func() {
					So(obj.Get(-1), ShouldEqual, -1)
					So(obj.Get(2), ShouldEqual, -1)
				})
			})

			Convey("delete one by one from head", func() {
				obj.DeleteAtIndex(0)
				So(obj.Get(0), ShouldEqual, 2)
				So(obj.Get(1), ShouldEqual, 3)
				So(obj.Get(2), ShouldEqual, -1)

				obj.DeleteAtIndex(0)
				So(obj.Get(0), ShouldEqual, 3)
				So(obj.Get(1), ShouldEqual, -1)

				obj.DeleteAtIndex(0)
				So(obj.Get(0), ShouldEqual, -1)

				Convey("get invalid index will return -1", func() {
					So(obj.Get(-1), ShouldEqual, -1)
					So(obj.Get(2), ShouldEqual, -1)
				})
			})

			Convey("delete one by one from tail", func() {

				Convey("delete last one", func() {
					obj.DeleteAtIndex(2)
					So(obj.Get(0), ShouldEqual, 1)
					So(obj.Get(1), ShouldEqual, 2)
					So(obj.Get(2), ShouldEqual, -1)

					Convey("delete that index again", func() {
						obj.DeleteAtIndex(2)
						So(obj.Get(0), ShouldEqual, 1)
						So(obj.Get(1), ShouldEqual, 2)
						So(obj.Get(2), ShouldEqual, -1)

						Convey("delete the rest", func() {
							obj.DeleteAtIndex(1)
							So(obj.Get(0), ShouldEqual, 1)
							So(obj.Get(1), ShouldEqual, -1)

							obj.DeleteAtIndex(0)
							So(obj.Get(0), ShouldEqual, -1)

							Convey("get invalid index will return -1", func() {
								So(obj.Get(-1), ShouldEqual, -1)
								So(obj.Get(2), ShouldEqual, -1)
							})
						})
					})
				})

			})
		})
	})
}
