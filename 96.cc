#include <bits/stdc++.h>
using namespace std;

/**
 *
 *  Unique Binary Search Trees
 *  Ref:
 *  https://leetcode.com/problems/unique-binary-search-trees/description/
 *  https://leetcode.com/problems/unique-binary-search-trees/solution/
 *
 **/

class Solution {
 public:
  /**
   *
   *  Approach 1: Dynamic Programming
   *  G(n): the number of unique BST for a sequence of length n.
   *  F(i,n): the number of unique BST where the number i is served as the root
   *of BST (1<=i<=n).
   *
   *  F(i,n) = G(i-1) * G(n-i)
   *  G(n) = SUM(i:1->n,F(i,n)) = SUM(i:1->n, G(i-1)*G(n-i))
   *  G(0) = 1, G(1) = 1
   *
   **/
  int numTrees(int n) {
    int* G = new int[n + 1]{1, 1};
    for (int i = 2; i <= n; i++) {
      for (int j = 1; j <= i; j++) {
        G[i] += G[j - 1] * G[i - j];
      }
    }
    return G[n];
  }

  /**
   *
   *  Approach 2: Mathematical Deduction
   *  G(n) is known as Catalan number (Cn).
   *
   *  C[0] = 1, C[n+1] = 2*(2*n+1)/(n+2)*C[n]
   *
   **/
  int numTrees2(int n) {
    long C = 1;
    for (int i = 0; i < n; i++) {
      C = C * 2 * (2 * i + 1) / (i + 2);
    }
    return (int)C;
  }
};

int main() {
  int q = 19;
  Solution s;
  cout << s.numTrees(q) << endl;
  cout << s.numTrees2(q) << endl;
}