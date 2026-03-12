package main

import "testing"

func TestGetMessage(t *testing.T) {
    expected := "CI Build Successful: Go App is running!"
    actual := GetMessage()

    if actual != expected {
        t.Errorf("Expected %s but got %s", expected, actual)
    }
}