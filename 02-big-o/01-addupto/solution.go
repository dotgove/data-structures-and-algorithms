package addupto

// tc O(n)
func addUptoA(n int) (total int) {
	for i := 0; i <= n; i++ {
		total += i
	}
	return
}

// tc O(1)
func addUptoB(n int) int {
	return n * (n + 1) / 2
}
