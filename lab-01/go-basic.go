package main

import "fmt"

func main() {
    // Variables
    a := 10
    b := 20

    // Addition
    sum := a + b
    fmt.Println("Sum =", sum)

    // Even/Odd
    if sum%2 == 0 {
        fmt.Println("Even")
    } else {
        fmt.Println("Odd")
    }

    // Loop
    for i := 1; i <= 5; i++ {
        fmt.Println("i =", i)
    }
}