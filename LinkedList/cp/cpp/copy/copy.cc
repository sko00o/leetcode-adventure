
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
    if (head == nullptr) {
      return nullptr;
    }

    // a list to store copy list
    Node* nHead;
    Node* np = nullptr;  // must initialize np here!

    // a list to store origin random pointer
    Node* rHead;
    Node* rp = nullptr;
    for (auto p = head; p != nullptr; p = p->next) {
      // store origin random for recover
      if (rp == nullptr) {
        rHead = new Node(0, nullptr, p->random);
        rp = rHead;
      } else {
        rp->next = new Node(0, nullptr, p->random);
        rp = rp->next;
      }

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

    // recover origin linked list
    for (auto p1 = head, p2 = rHead; p1 != nullptr && p2 != nullptr;
         p1 = p1->next, p2 = p2->next) {
      p1->random = p2->random;
    }

    return nHead;
  }
};