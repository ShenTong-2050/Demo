package gotest

import "testing"

func fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func optimizeFibonacci(x int) int {
	const mod = 1e9

	if x < 2 {
		return x
	}

	p,q,r := 0,0,1

	for i:=2; i< x; i++ {
		p = q
		q = r
		r = (p+q) % mod
	}

	return r
}

// 最优解
const mod int = 1e9 + 7

type matrix [2][2]int

func multiply(a, b matrix) (c matrix) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = (a[i][0]*b[0][j] + a[i][1]*b[1][j]) % mod
		}
	}
	return
}

func pow(a matrix, n int) matrix {
	ret := matrix{{1, 0}, {0, 1}}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ret = multiply(ret, a)
		}
		a = multiply(a, a)
	}
	return ret
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	res := pow(matrix{{1, 1}, {1, 0}}, n-1)
	return res[0][0]
}

func BenchmarkFibonacci20(t *testing.B) {
	for i:=0; i<t.N; i++ {
		fibonacci(20)
	}
}

func BenchmarkFibonacciOptimize(t *testing.B) {
	for i:=0; i<t.N; i++ {
		fib(20)
	}
}

func BenchmarkFibonacciMostOptimize(t *testing.B) {
	for i:=0; i<t.N; i++ {
		optimizeFibonacci(20)
	}
}


