package main

import (
	"strings"
	"testing"
)


func Test2(t *testing.T) {
    input := "eighthree"

    res := b(strings.NewReader(input))

    if res != 83 {
        t.Fatalf("Should've gotten 83, got %d\n", res)
    }
}
func Test1(t *testing.T) {
    input := "two1nine"

    res := b(strings.NewReader(input))

    if res != 29 {
        t.Fatalf("Should've gotten 29, got %d\n", res)
    }
}
