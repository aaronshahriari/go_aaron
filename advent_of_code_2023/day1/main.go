package main

import (
	"fmt"
	"log"
	"os"
)

// CONSTRAINTS:
// 1. no digit? -> ???
// 2. one digit? -> **USE THAT DIGIT TWICE**
// 3. multiple digits? -> **FIRST AND LAST**
// 4. ???

func main() {
    content, err := os.ReadFile("input.txt")
    if err != nil {
        log.Fatal(err)
    }

    // string output of file
    string_output := string(content)
    fmt.Println(string_output)


    // fmt.Println("DOC LENGTH:", len(doc))
    // fmt.Println()
}
