package main

import "fmt"

// GetMessage returns the success string
func GetMessage() string {
    return "CI Build Successful: Go App is running!"
}


// Add takes two integers and returns their sum
func Add(a int, b int) int {
    return a + b
}

func main() {
    fmt.Println(GetMessage())
    result := Add(5, 10)
    fmt.Printf("The sum is: %d\n", result)
}
