#include <bits/stdc++.h>
using namespace std;

/**
 *
 * Definition for a binary tree node.
 *
 */
struct TreeNode {
  int val;
  TreeNode* left;
  TreeNode* right;
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

class Solution {
 public:
  Solution() { this->num = 0; }

  vector<TreeNode*> generateTrees(int n) {
    if (n <= 0) {
      return vector<TreeNode*>();
    }
    this->num = n;

    vector<TreeNode*> G[n + 1];

    G[0].push_back(nullptr);
    G[1].push_back(new TreeNode(1));
    for (int i = 2; i <= n; i++) {
      for (int j = 1; j <= i; j++) {
        for (auto l : G[j - 1]) {
          for (auto r : G[i - j]) {
            auto nd = new TreeNode(j);
            nd->left = l;
            // nd->right = r;
            nd->right = newTree(r, j);
            G[i].push_back(nd);
          }
        }
      }
    }
    return G[n];
  }

  // print trees
  void printTrees(vector<TreeNode*> t, bool needFix) {
    bool first = true;
    cout << "[";
    int n = num;
    for (auto i : t) {
      if (first) {
        cout << endl;
        first = false;
      } else {
        cout << "," << endl;
      }
      if (needFix) {
        cout << "  " << treeBFSStrFix(i, n);
      } else {
        cout << "  " << treeBFSStr(i, n);
      }
    }
    if (!first) {
      cout << endl;
    }
    cout << "]" << endl;
  }

 private:
  int num;

  // copy and add offset in val
  TreeNode* newTree(TreeNode* a, int offset) {
    if (!a) {
      return nullptr;
    }
    auto r = new TreeNode(a->val + offset);
    r->left = newTree(a->left, offset);
    r->right = newTree(a->right, offset);
    return r;
  }

  // BFS the tree and return the result string
  string treeBFSStr(TreeNode* a, int n) {
    queue<TreeNode*> q;
    q.push(a);
    bool first = true;
    string out;
    out += "[";
    while (!q.empty()) {
      auto b = q.front();
      q.pop();
      if (n) {
        if (first) {
          first = false;
        } else {
          out += ",";
        }
        if (b) {
          out += to_string(b->val);
          n--;
          if (b->left || b->right) {
            q.push(b->left);
            q.push(b->right);
          }
        } else {
          out += "null";
        }
      }
    }
    out += "]";
    return out;
  }

  // fix val while printing
  string treeBFSStrFix(TreeNode* a, int n) {
    struct tmp {
      TreeNode* t;
      bool r;
    };
    queue<tmp> q;
    map<TreeNode*, int> mp;
    q.push({a, false});
    bool first = true;
    string out;
    out += "[";
    while (!q.empty()) {
      auto qb = q.front();
      auto b = qb.t;
      q.pop();
      if (n) {
        if (first) {
          first = false;
        } else {
          out += ",";
        }
        if (b) {
          auto val = b->val;
          auto offset = 0;
          if (qb.r) {
            offset = mp[b];
          }
          val += offset;
          out += to_string(val);
          n--;
          if (b->left || b->right) {
            q.push({b->left, qb.r});
            if (qb.r) {
              // 左节点记录，父节点为祖父节点右孩子时的增量
              mp[b->left] = offset;
            }
            q.push({b->right, true});
            // 右节点记录，父节点的值
            mp[b->right] = val;
          }
        } else {
          out += "null";
        }
      }
    }
    out += "]";
    return out;
  }
};

void stdTest() {
  Solution s;
  auto g = s.generateTrees(3);
  s.printTrees(g, false);
}

int main() {
  stdTest();

  return 0;
}