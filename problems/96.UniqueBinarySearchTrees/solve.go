package problems

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
func numTrees(n int) int {
	var dp = make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

/**
 *
 *  Approach 2: Mathematical Deduction
 *  G(n) is known as Catalan number (Cn).
 *
 *  C[0] = 1, C[n+1] = 2*(2*n+1)/(n+2)*C[n]
 *
 **/
func numTrees1(n int) int {
	var ans = 1
	for i := 0; i < n; i++ {
		ans = ans * 2 * (2*i + 1) / (i + 2)
	}
	return ans
}
