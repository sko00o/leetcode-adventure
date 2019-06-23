
class Node {
public:
    int val;
    Node* prev;
    Node* next;
    Node* child;

    Node() {}

    Node(int _val, Node* _prev, Node* _next, Node* _child) {
        val = _val;
        prev = _prev;
        next = _next;
        child = _child;
    }
};

class Solution {
 public:
  Node* flatten(Node* head) {
    auto res = dfs(head, true);
    return res.out;
  }

 private:
  typedef struct {
    Node* out;
    Node* np;
  } result;

  result dfs(Node* head, bool topLevel) {
    auto res = result{};
    if (head == nullptr) {
      return res;
    }

    auto p = head;
    for (res.np = head; p != nullptr; res.np = p, p = p->next) {
      if (topLevel) {
        res.out = head;
      }
      if (p->child != nullptr) {
        auto next = p->next;
        auto cHead = p->child;
        p->child = nullptr;
        auto child = dfs(cHead, false);
        auto cTail = child.np;

        p->next = cHead;
        cHead->prev = p;
        if (cTail != nullptr) {
          cTail->next = next;
          if (next != nullptr) {
            next->prev = cTail;
          }
          p = cTail;
        }
      }
    }
    return res;
  }
};
