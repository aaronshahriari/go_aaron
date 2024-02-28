package main

import (
	"fmt"
	"maps"
	_ "maps"
)

// hashes or dicts lmao
func main() {
    // create and print map
    map1 := make(map[string]int)
    map1["goals"] = 3
    map1["assists"] = 1

    fmt.Println(map1)

    // grab a value based on key
    key1 := map1["goals"]
    fmt.Println("key1", key1)

    // 
    header_key, key_exist := map1["header"]
    fmt.Println("test_key", header_key)
    fmt.Println("test_key exist?", key_exist)

    // length
    fmt.Println("length:", len(map1))

    // delete a record
    delete(map1, "assists")
    fmt.Println(map1)

    // clear the whole dict
    clear(map1)
    fmt.Println(map1)

    // set dict in new format
    n := map[string]int{"a": 1, "b":2}
    fmt.Println(n)

    // comparing dicts
    n2 := map[string]int{"a": 1, "b":2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    } else { fmt.Println("n != n2") }

}
