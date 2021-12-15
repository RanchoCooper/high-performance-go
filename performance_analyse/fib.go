package main

/**
 * @author Rancho
 * @date 2021/12/15
 */

func fib(n int) int {
    if n == 0 || n == 1 {
        return n
    }

    return fib(n-1) + fib(n-2)
}
