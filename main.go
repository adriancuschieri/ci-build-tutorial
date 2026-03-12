package main

import "fmt"

// GetMessage returns the success string
func GetMessage() string {
    return "CI Build Successful: Go App is running!"
}

func main() {
    fmt.Println(GetMessage())
}
