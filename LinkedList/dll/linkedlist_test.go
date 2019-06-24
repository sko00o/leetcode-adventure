package dll

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"linkedlist/dll/ll1"
	"linkedlist/dll/ll2"
	"linkedlist/dll/ll3"
	"linkedlist/dll/ll4"
)

func TestLinkedList(t *testing.T) {
	tasks := []func() LinkedList{
		func() LinkedList {
			out := ll1.Constructor()
			return &out
		},
		func() LinkedList {
			out := ll2.Constructor()
			return &out
		},
		func() LinkedList {
			out := ll3.Constructor()
			return &out
		},
		func() LinkedList {
			out := ll4.Constructor()
			return &out
		},
	}
	for idx, Constructor := range tasks {
		Convey(fmt.Sprintf("task #%d", idx), t, func() {
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
					Convey("delete last one", func() {
						obj.DeleteAtIndex(0)
						So(obj.Get(0), ShouldEqual, 2)
						So(obj.Get(1), ShouldEqual, 3)
						So(obj.Get(2), ShouldEqual, -1)

						Convey("delete last one", func() {
							obj.DeleteAtIndex(0)
							So(obj.Get(0), ShouldEqual, 3)
							So(obj.Get(1), ShouldEqual, -1)

							Convey("delete last one", func() {
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

			Convey("complex test", func() {
				obj.AddAtHead(56) // [56]
				So(obj.Get(1), ShouldEqual, -1)
				obj.AddAtHead(41) // [41,56]
				obj.AddAtTail(98) // [41,56,98]
				Convey("now we got [41,56,98]", func() {
					So(obj.Get(0), ShouldEqual, 41)
					So(obj.Get(1), ShouldEqual, 56)
					So(obj.Get(2), ShouldEqual, 98)
				})
				So(obj.Get(3), ShouldEqual, -1)
				obj.AddAtIndex(1, 33) // [41,33,56,98]
				Convey("now we got [41,33,56,98]", func() {
					So(obj.Get(0), ShouldEqual, 41)
					So(obj.Get(1), ShouldEqual, 33)
					So(obj.Get(2), ShouldEqual, 56)
					So(obj.Get(3), ShouldEqual, 98)
				})
				obj.AddAtHead(72) // [72,41,33,56,98]
				obj.AddAtHead(52) // [52,72,41,33,56,98]
				obj.AddAtHead(89) // [89,52,72,41,33,56,98]
				obj.AddAtHead(0)  // [0,89,52,72,41,33,56,98]
				obj.AddAtHead(98) // [98,0,89,52,72,41,33,56,98]
				Convey("now we got [98,0,89,52,72,41,33,56,98]", func() {
					So(obj.Get(0), ShouldEqual, 98)
					So(obj.Get(1), ShouldEqual, 0)
					So(obj.Get(2), ShouldEqual, 89)
					So(obj.Get(3), ShouldEqual, 52)
					So(obj.Get(4), ShouldEqual, 72)
					So(obj.Get(5), ShouldEqual, 41)
					So(obj.Get(6), ShouldEqual, 33)
					So(obj.Get(7), ShouldEqual, 56)
					So(obj.Get(8), ShouldEqual, 98)
				})
				obj.AddAtIndex(7, 97) // [98,0,89,52,72,41,33,97,56,98]
				obj.AddAtIndex(2, 51) // [98,0,51,89,52,72,41,33,97,56,98]
				Convey("now we got [98,0,51,89,52,72,41,33,97,56,98]", func() {
					So(obj.Get(0), ShouldEqual, 98)
					So(obj.Get(1), ShouldEqual, 0)
					So(obj.Get(2), ShouldEqual, 51)
					So(obj.Get(3), ShouldEqual, 89)
					So(obj.Get(4), ShouldEqual, 52)
					So(obj.Get(5), ShouldEqual, 72)
					So(obj.Get(10), ShouldEqual, 98)
					So(obj.Get(9), ShouldEqual, 56)
					So(obj.Get(8), ShouldEqual, 97)
					So(obj.Get(7), ShouldEqual, 33)
					So(obj.Get(6), ShouldEqual, 41)
				})
			})

			Convey("add at head", func() {
				obj.AddAtHead(56) // [56]
				So(obj.Get(0), ShouldEqual, 56)

				Convey("add at head", func() {
					obj.AddAtHead(1) // [1,56]
					So(obj.Get(0), ShouldEqual, 1)
					Convey("get invalid index will return -1", func() {
						So(obj.Get(-1), ShouldEqual, -1)
					})

					Convey("add at head", func() {
						obj.AddAtHead(41) // [41,1,56]
						So(obj.Get(0), ShouldEqual, 41)

						Convey("add at tail", func() {
							obj.AddAtTail(98) // [41,1,56,98]
							So(obj.Get(3), ShouldEqual, 98)

							Convey("add at index", func() {
								obj.AddAtIndex(1, 33) // [41,33,1,56,98]
								So(obj.Get(3), ShouldEqual, 56)

								Convey("add at head", func() {
									obj.AddAtHead(75) // [75,41,33,1,56,98]
									obj.AddAtHead(52) // [52,75,41,33,1,56,98]
									obj.AddAtHead(89) // [89,52,75,41,33,1,56,98]
									obj.AddAtHead(66) // [66,89,52,75,41,33,1,56,98]
									So(obj.Get(3), ShouldEqual, 75)
									So(obj.Get(2), ShouldEqual, 52)
									So(obj.Get(1), ShouldEqual, 89)
									So(obj.Get(0), ShouldEqual, 66)

									Convey("add at index", func() {
										obj.AddAtIndex(6, 22) // [66,89,52,75,41,33,22,1,56,98]
										So(obj.Get(1), ShouldEqual, 89)
										So(obj.Get(3), ShouldEqual, 75)
										So(obj.Get(8), ShouldEqual, 56)
										So(obj.Get(9), ShouldEqual, 98)

										Convey("get invalid index will return -1", func() {
											So(obj.Get(-1), ShouldEqual, -1)
											So(obj.Get(10), ShouldEqual, -1)
										})

										Convey("delete at index", func() {
											obj.DeleteAtIndex(2) // [66,89,75,41,33,22,1,56,98]
											So(obj.Get(1), ShouldEqual, 89)
											So(obj.Get(2), ShouldEqual, 75)

											Convey("delete at index", func() {
												obj.DeleteAtIndex(7) // [66,89,75,41,33,22,1,98]
												So(obj.Get(7), ShouldEqual, 98)
												So(obj.Get(6), ShouldEqual, 1)
											})
										})
									})
								})
							})
						})
					})
				})
			})
		})
	}
}
