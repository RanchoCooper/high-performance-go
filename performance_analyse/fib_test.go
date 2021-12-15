package main

import (
    "testing"
)

/**
 * @author Rancho
 * @date 2021/12/15
 */

func Benchmark_fib(b *testing.B) {
    for n := 0; n < b.N; n++ {
        fib(10)
    }
}
