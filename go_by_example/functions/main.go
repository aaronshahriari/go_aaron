package main

import "fmt"

// return single values
func two_add(a int, b int) int {
    return a + b
}

func three_add(a int, b int, c int) int {
    return a + b + c
}

// return multiple values
func multiple_return(a int, b int, c string, d int) (int, string) {
    int_res := a + b + d
    string_res := c + " penis!!!"

    return int_res, string_res
}

// takes in a slice and does some calc
func variadic_function(nums...int) {
    fmt.Println(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println("total:", total)
}

// takes in a slice and does some calc
// we want a result here so that we can use it laster
func return_variadic_function(nums...int) int {
    fmt.Println(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func anonymous_closure() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func recursive(n int) int {
    if n == 0 {
        return 1
    }
    return n * recursive(n-1)
}

func main() {
    // single return
    two := two_add(1, 2)
    fmt.Println("two add:", two)
    three := three_add(1, 2, 3)
    fmt.Println("three add:", three)

    // multi return
    fmt.Println()
    multi_int, multi_string := multiple_return(1, 2, "Grace's", 3)
    fmt.Println("string result:", multi_string)
    fmt.Println("int result:", multi_int)

    // variadic functions accepting slices
    fmt.Println()
    variadic_function(1, 2, 3, 4, 5)
    total_result := return_variadic_function(1, 2, 3, 4, 5)
    fmt.Println("returned total:", total_result)

    // anonymous function forming closures
    // start counting using this var call
    fmt.Println()
    nextInt := anonymous_closure()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // start counting using this NEW var call
    fmt.Println()
    nextInt2 := anonymous_closure()

    fmt.Println(nextInt2())
    fmt.Println(nextInt2())
    fmt.Println(nextInt2())

    // recursive function use
    fmt.Println()
    fmt.Println(recursive(7))

    var fib func(n int) int

    fib = func(n int) int {
        if n < 2 {
            return n
        }
        return fib(n-1) + fib(n-2)
    }
    fmt.Println(fib(7))
}
