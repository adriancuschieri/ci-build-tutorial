package main

import "testing"

func TestGetMessage(t *testing.T) {
    expected := "CI Build Successful: Go App is running!"
    actual := GetMessage()

    if actual != expected {
        t.Errorf("Expected %s but got %s", expected, actual)
    }
}


func TestAddTableDriven(t *testing.T) {
    // Define a slice of anonymous structs for our test cases
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"Positive numbers", 5, 5, 10},
        {"Negative numbers", -1, -1, -2},
        {"Mixing types", -5, 10, 5},
        {"Zero case", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("got %d, want %d", result, tt.expected)
            }
        })
    }
}