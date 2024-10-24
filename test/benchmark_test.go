package main

import "testing"

// benchmark 主要用来测试函数性能，包括执行效率、内存分配情况
// https://www.cnblogs.com/yahuian/p/go-benchmark.html

// 菲波那切数
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

/**
 * 基准测试 命令：go test -bench .
 * b.N 从 1 开始，如果该用例能够在 1s 内完成，b.N 的值便会增加，再次执行，递增顺序：1, 2, 3, 5, 10, 20, 30, 50, 100
 *
 */
func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = fib(30) // run fib(30) b.N times
	}
}
