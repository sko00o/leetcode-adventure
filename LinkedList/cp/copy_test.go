package cp

import (
	"encoding/json"
	"errors"
	"testing"
)

func makeLinkedListByJSON(raw []byte) (*ListNode, error) {
	list := new(ListNode)
	if err := json.Unmarshal(raw, list); err != nil {
		return nil, err
	}

	var head, np *ListNode
	M := make(map[string]*ListNode)
	for p := list; p != nil; p = p.Next {
		M[p.ID] = p
		if np == nil {
			head, np = p, p
		} else {
			np.Next = p
			np = np.Next
		}
	}

	for p := head; p != nil; p = p.Next {
		if p.Random != nil {
			ref, ok := M[p.Random.Ref]
			if !ok {
				return nil, errors.New("random lost ref")
			}
			p.Random = ref
		}
	}

	return head, nil
}

func Test_makeLinkedListByJSON(t *testing.T) {
	tasks := []struct {
		listRaw []byte
	}{
		{
			listRaw: []byte(`{"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}`),
		},
	}

	for _, task := range tasks {
		list, err := makeLinkedListByJSON(task.listRaw)
		if err != nil {
			t.Fatal(err)
		}
		for p := list; p != nil; p = p.Next {
			t.Logf("Val: %d, Address: %p", p.Val, p)
			if p.Next != nil {
				t.Logf("Next: %+v", p.Next)
			}
			if p.Random != nil {
				t.Logf("Random: %+v", p.Random)
			}
		}
	}
}

func Test_copyRandomList(t *testing.T) {
	tasks := []struct {
		listRaw []byte
	}{
		{
			listRaw: []byte(`{"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}`),
		},
	}

	for fIdx, f := range []func(*ListNode) *ListNode{copyRandomList} {
		for i, task := range tasks {
			h1, err := makeLinkedListByJSON(task.listRaw)
			if err != nil {
				t.Fatal(err)
			}
			h2 := f(h1)

			// check
			M1, M2 := make(map[*ListNode]int), make(map[*ListNode]int)
			for idx, p1, p2 := -0, h1, h2; p1 != nil && p2 != nil; idx, p1, p2 = idx+1, p1.Next, p2.Next {
				M1[p1], M2[p2] = idx, idx
			}
			for p1, p2 := h1, h2; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
				if p1 == p2 {
					t.Errorf("func #%d, task #%d failed, two lists have sharded node", fIdx, i)
				}

				if p1.Val != p2.Val {
					t.Errorf("func #%d, task #%d failed, value not match, got: %d, expect: %d", fIdx, i, p2.Val, p1.Val)
				}

				if p1.Random == nil {
					if p2.Random != nil {
						t.Errorf("func #%d, task #%d failed, wrong random, got: %+v, expect: %+v", fIdx, i, p2.Random, p1.Random)
					}
				} else {
					if idx1, idx2 := M1[p1.Random], M2[p2.Random]; idx1 != idx2 {
						t.Errorf("func #%d, task #%d failed, wrong random index, got: %+v, expect: %+v", fIdx, i, idx2, idx1)
					}
				}
			}
		}
	}
}
