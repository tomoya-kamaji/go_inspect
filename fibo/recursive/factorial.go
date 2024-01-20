package recursive

// 階差数列を求める: 1, 2, 6, 24, 120, 720, 5040, 40320, 362880
func factorial(n int) int {
	// ベースケース
	if n == 0 {
		return 1
	}

	// 再帰呼び出し
	return n * factorial(n-1)
}
