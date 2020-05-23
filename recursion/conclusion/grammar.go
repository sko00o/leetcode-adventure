package conclusion

func kthGrammar(N int, K int) int {
	return gen(N-1, K-1)
}

func gen(N, K int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return K
	}

	return (gen(N-1, K/2) + K%2) % 2
}
