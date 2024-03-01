package main

import (
    "fmt"
)

func main() {
    // create a list and sum through in range loop
    nums := []int{2, 3, 4, 5}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    fmt.Println()
    // use i when you need iterator count of loop
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    fmt.Println()
    // now do the same using a slice
    kvs := map[string]string{"a": "apple", "b": "banana", "c": "candle"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    fmt.Println()
    // iterate over keys
    for k := range kvs {
        fmt.Println("key ->", k)
    }

    fmt.Println()
    // iterate over values
    for _, v := range kvs {
        fmt.Println("value ->", v)
    }

    fmt.Println()
    // iterate over unicode points
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
