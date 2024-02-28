package main

import (
	"fmt"
	"slices"
)

func main() {
    var s []string
    // make a slice with length of 3
    s = make([]string, 3)
    fmt.Println("len:", len(s), "cap:", cap(s))

    s[0] = "h"
    s[1] = "e"
    s[2] = "l"

    s = append(s, "l")
    s = append(s, "o")

    fmt.Println("new len:", len(s), "new cap:", cap(s))
    fmt.Println("full:", s)

    // make a copy of s to c
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("copy of s:", c)

    // cutting slices into smaller slices
    l := c[2:5]
    fmt.Println("sl1", l)

    l = c[:5]
    fmt.Println("sl2", l)

    l = c[2:]
    fmt.Println("sl1", l)

    // comparing slices
    t := []string{"g", "h", "i"}
    t2 := []string{"g", "h", "i"}

    if slices.Equal(t, t2) {
        fmt.Println("t == t2")
    } else { fmt.Println("t != t2") }

    // slices can be multidimensional
    // different than arrays, because the inner can have set len
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLength := i + 1
        twoD[i] = make([]int, innerLength)
        for j := 0; j < innerLength; j++ {
            twoD[i][j] = i + j
        }
    }
    // created incrementing slice per slice in main
    fmt.Println("2d:", twoD)
}
