package main

import (
	"fmt"
	"strings"
	"testing"
)

func createMatrix(lines []string) [][]string {
    matrix := [][]string{}
    for _, v := range lines {
        matrix = append(matrix, strings.Split(v, ""))
    }

    return matrix
} 



func Test1(t *testing.T) {
    testCases := []struct {
        lines    []string
        expected int
    }{
        {
            []string{
                "..1..",
                ".#...",
            },
            1,
        },
        {
            []string{
                "..1..",
                ".....",
            },
            0,
        },
        {
            []string{
                "..1..",
                "..#..",
            },
            1,
        },
        {
            []string{
                "..1..",
                "...#.",
            },
            1,
        },
        {
            []string{
                "..1#.",
                ".....",
            },
            1,
        },
        {
            []string{
                ".#1..",
                ".....",
            },
            1,
        },
        {
            []string{
                ".#...",
                "..1..",
            },
            1,
        },
        {
            []string{
                "..#..",
                "..1..",
            },
            1,
        },
        {
            []string{
                "...#.",
                "..1..",
            },
            1,
        },
        {
            []string{
                "..1#.",
                "..1..",
            },
            2,
        },
        {
            lines: []string{
                "..123",
                ".#...",
            },
            expected: 123,
        },
        {
            lines: []string{
                "1....",
                ".#...",
            },
            expected: 1,
        },
        {
            lines: []string{
                "123..",
                "...#.",
            },
            expected: 123,
        },
        {
            lines: []string{
                "1...1",
                "1#.#1",
            },
            expected: 4,
        },
        {
            lines: []string {
                "206#13..",
                "........",
            },
            expected: 219,
        },
    }
    for i, v := range testCases {
        fmt.Printf("TestCase %d\n", i)
        matrix := createMatrix(v.lines)
        sum := calculateSumA(matrix)

        if sum != v.expected {
            t.Errorf("Expected '%d', got '%d'\n", v.expected, sum)
        }
    }
}
