package main

import (
        "fmt"
        "time"
)

func main() {

        fmt.Println("Problem n. ")
        start := time.Now()

        answer := 0

        fmt.Printf("-->Answer: The answer is %d\n", answer)
        fmt.Printf("took=%s\n\n", time.Since(start))

}
