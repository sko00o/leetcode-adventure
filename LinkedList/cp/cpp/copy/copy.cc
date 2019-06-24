
class Node {
 public:
  int val;
  Node* next;
  Node* random;

  Node() {}

  Node(int _val, Node* _next, Node* _random) {
    val = _val;
    next = _next;
    random = _random;
  }
};

class Solution {
 public:
  Node* copyRandomList(Node* head) {
    if(head==nullptr) {
      return nullptr;
    }

    Node* nHead;
    // must initialize np here!
    Node* np = nullptr;

    for (auto p = head; p != nullptr; p = p->next) {
      auto nNode = new Node(p->val, p->next, p->random);
      // point the old node random to new node
      p->random = nNode;

      if (np == nullptr) {
        nHead = nNode;
        np = nNode;
      } else {
        np->next = nNode;
        np = np->next;
      }
    }

    for (auto p = nHead; p != nullptr; p = p->next) {
      if (p->random != nullptr) {
        p->random = p->random->random;
      }
    }
    return nHead;
  }
};