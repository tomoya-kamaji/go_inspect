package recursive

// 要件:
// フィボナッチ公式
// f(n) = f(n-1) + f(n-2)
// f(0) = 0
// f(1) = 1
// f(2) = 1
// f(3) = 2
// f(4) = 3
// f(5) = 5
// f(6) = 8
// f(7) = 13
// f(8) = 21
// f(9) = 34
// f(10) = 55
func fibonacci(n int) int {
	// ベースケース
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	// 再帰呼び出し
	return fibonacci(n-1) + fibonacci(n-2)
}
