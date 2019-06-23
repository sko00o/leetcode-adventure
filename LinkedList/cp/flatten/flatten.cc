#include "node.h"

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
    if (head == NULL) {
      return res;
    }

    auto p = head;
    for (res.np = head; p != NULL; res.np = p, p = p->next) {
      if (topLevel) {
        res.out = head;
      }
      if (p->child != NULL) {
        auto next = p->next;
        auto cHead = p->child;
        p->child = NULL;
        auto child = dfs(cHead, false);
        auto cTail = child.np;

        p->next = cHead;
        cHead->prev = p;
        if (cTail != NULL) {
          cTail->next = next;
          if (next != NULL) {
            next->prev = cTail;
          }
          p = cTail;
        }
      }
    }
    return res;
  }
};
