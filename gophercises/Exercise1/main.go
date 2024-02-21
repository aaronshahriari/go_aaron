package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// allow or disallow decimals
// dont allow spaces
// dont allow enter

func main() {
    records := read_csv()

    timerExpired := make(chan bool)
    quiz_complete := make(chan bool)

    quiz_length := 30

    fmt.Println("DIRECTIONS:")
    fmt.Println("1. READ QUESTION")
    fmt.Println("2. TYPE ANSWER")
    fmt.Println("3. IF THE ANSWER IS A WHOLE NUMBER IT MUST BE IN THIS FORMAT (EX. 5 NOT 5.0 OR 05)")
    fmt.Println("4. WHEN QUESTION ANSWER IS TYPED PRESS ENTER TO SUBMIT AND GO TO THE NEXT QUESTION")
    fmt.Println("5. YOU HAVE 30 SECONDS AFTER PRESSING THE INITIAL ENTER")
    fmt.Println()

    fmt.Println("PRESS ENTER TO START THE QUIZ")
    var input string
    fmt.Scanln(&input)

    go func() {
        timer := time.NewTimer(time.Duration(quiz_length) * time.Second)
        <-timer.C
        timerExpired <- true
    }()

    go func() {
        game(records)
        quiz_complete <- true
    }()

    select {
    case <- timerExpired:
        fmt.Println("\nTIME IS UP!")
    case <- quiz_complete:
        fmt.Println("\nQUIZ COMPLETE!")
    }
}

func read_csv() [][]string {
    file, err := os.Open("problems.csv")
    if err != nil {
        log.Fatal(err)
    }

    defer file.Close()

    reader := csv.NewReader(file)

    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    return records
}

func game(records [][]string) {
    var correct_count int = 0
    reader := bufio.NewReader(os.Stdin)
    for _, fullrecord := range records {
        fmt.Printf("Answer this question: %v\n", fullrecord[0])

        var user_answer string
        // _, err := fmt.Scanln(&user_answer)
        user_answer, err := reader.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        user_answer = strings.TrimSpace(user_answer)

        if user_answer == fullrecord[1] {
            correct_count++
        }
    }

    full_correct := strconv.Itoa(correct_count)
    csv_length := len(records)
    fmt.Printf("%v/%v\n", full_correct, csv_length)
}
