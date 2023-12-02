package main

import (
	"testing"
)

func Test2A1(t *testing.T) {
    input := "Game 1: 3 blue, 4 red"

    res := getGameIdIfPossible(input)

    if res != 1 {
        t.Fatalf("Expected 1, got %d\n", res)
    }
}
func Test2A2(t *testing.T) {
    input := "Game 1: 3 blue, 13 red"

    res := getGameIdIfPossible(input)

    if res != 0 {
        t.Fatalf("Expected 0, got %d\n", res)
    }
}

func Test2B1(t *testing.T) {
    input := "Game 1: 3 blue, 4 red, 2 green"

    res := getPower(input)

    if res != 24 {
        t.Fatalf("Expected 12, got %d\n", res)
    }
}
