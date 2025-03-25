package main

import "fmt"

func divide(a, b int) int {
    return a / b // Potential division by zero
}

func main() {
    x := 10
    y := 0
    result := divide(x, y) // This will cause a runtime error
    fmt.Println("Result:", result)
}
