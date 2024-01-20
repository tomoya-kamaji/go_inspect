package recursive

// (3, 1, 3, 2)
func move(n, from, to, via int) {
	if n > 0 {
		// 最上部のn-1枚の円盤を、棒fromから棒viaへ移動させる。
		move(n-1, from, via, to)
		// 残りの1枚の円盤を、棒fromから棒toへ移動させる。
		println("Move disk", n, "from peg", from, "to peg", to)
		// 棒viaに移動させたn-1枚の円盤を、棒toへ移動させる。
		move(n-1, via, to, from)
	}
}
