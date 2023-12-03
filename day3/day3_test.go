package main

import (
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



func TestGetNumber(t *testing.T) {
    strArr := []string { 
        "467.",
    }
    matrix := createMatrix(strArr)
    val := readNumberInPos(0,2,matrix)

    if val != 467 {
        t.Fatalf("Expected '467', got '%d'\n", val)
    }
}

func TestGetNumber2(t *testing.T) {
    strArr := []string { 
        "..35",
    }
    matrix := createMatrix(strArr)
    val := readNumberInPos(0,3,matrix)

    if val != 35 {
        t.Fatalf("Expected '35', got '%d'\n", val)
    }
}

func TestB(t *testing.T) {
    testCases := []struct {
        lines    []string
        expected int
    }{
        {
            []string{
                "...*4",
                ".....",
            },
            0,
        },
        {
            []string{
                "...*4.",
                "......",
            },
            0,
        },
        {
            []string{
                "2*4.",
                "....",
            },
            8,
        },
        {
            []string{
                "..2*4",
                ".....",
            },
            8,
        },
        {
            []string{
                "..10*",
                "....9",
            },
            90,
        },
        {
            []string{
                "467.",
                "...*",
                "..35",
            },
            16345,
        },
        {
            []string{
                "467..114..",
                "...*......",
                "..35..633.",
                "......#...",
                "617*......",
                ".....+.58.",
                "..592.....",
                "......755.",
                "...$.*....",
                ".664.598..",
            },
            467835,
        },
    }

    for _, v := range testCases {
        matrix := createMatrix(v.lines)
        sum := calculateSumB(matrix)

        if sum != v.expected {
            t.Errorf("Expected '%d', got '%d'\n", v.expected, sum)
        } 
    }
}

func TestA(t *testing.T) {
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
    for _, v := range testCases {
        matrix := createMatrix(v.lines)
        sum := calculateSumA(matrix)

        if sum != v.expected {
            t.Errorf("Expected '%d', got '%d'\n", v.expected, sum)
        }
    }
}
