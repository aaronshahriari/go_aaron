package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// CONSTRAINTS:
// 1. no digit? -> ???
// 2. one digit? -> **USE THAT DIGIT TWICE**
// 3. multiple digits? -> **FIRST AND LAST**

// Aaron's Logic
func readtxt(file_path string) []string {
    // read txt in to content (string) variable
    content, err := os.ReadFile(file_path)
    if err != nil {
        log.Fatal(err)
    }

    // get string output and convert to slice
    slice_output := strings.Split(string(content), "\n")

    return slice_output
}

func isDigit(slice_output []string) int {
    total := 0
    for _, line := range slice_output {
        if len(line) == 0 {
            continue
        }
        digit_slice := []string{}
        for _, rune := range line {
            is_digit := unicode.IsDigit(rune)
            if is_digit == true {
                digit_slice = append(digit_slice, string(rune))
            }
        }
        concat_digit := digit_slice[0] + digit_slice[len(digit_slice)-1]
        digit_total, err := strconv.Atoi(concat_digit)
        if err != nil {
            log.Fatal(err)
        }
        total += digit_total
    }
    return total
}

// Optimized Logic


func main() {
    // starting time to check optimization
    start := time.Now()

    // read in txt; convert to string; create slice
    file_path := "input.txt"
    slice_output := readtxt(file_path)

    // find total based on aaron logic
    total := isDigit(slice_output)
    fmt.Println(total)

    // elapsed time for optimization
    stop := time.Now()
    elapsed := stop.Sub(start).Seconds()
    fmt.Println()
    fmt.Println(elapsed, "seconds")
}
