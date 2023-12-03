package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day3/atest.txt")
    // file, err := os.Open(pwd + "/day3/btest.txt")
    file, err := os.Open(pwd + "/day3/input.txt")

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    // res := a(file)
    res := b(file)
    fmt.Printf("%d\n", res)
}

func b(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    matrix := [][]string{}
    for scanner.Scan() {
        line := scanner.Text()
        arr := strings.Split(line, "")
        matrix = append(matrix, arr)
    }

    sum := calculateSumB(matrix)

    return sum
}

func a(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    matrix := [][]string{}
    for scanner.Scan() {
        line := scanner.Text()
        arr := strings.Split(line, "")
        matrix = append(matrix, arr)
    }

    sum := calculateSumA(matrix)

    return sum
}

func isDigit(s string) bool {
    if s== "0" || s == "1" || s == "2" || s == "3"|| s == "4"|| s == "5"|| s == "6"|| s == "7" || s == "8" || s == "9" {
        return true
    }
    return false
}

type gear struct {
    numberA int
    numberB int
}

func calculateSumB(matrix [][]string) int {
    sum := 0 

    for y := 0; y < len(matrix); y++ {
        for x := 0; x < len(matrix[y]); x++ {
            xv := matrix[y][x]
            if xv == "*" {
                if gear, res := isGear(y,x,matrix); res {
                    sum += (gear.numberA * gear.numberB)
                }
                
            }
        }
    }

    return sum
}

func readNumberInPos(y, x int, matrix [][]string) int {
    line := matrix[y]
    beginX := x

    if x != 0 {
        for i := x-1; i >= 0; i-- {
            if !isDigit(line[i]) {
                beginX = i+1
                break
            }
            beginX = i
        }
    }

    str := ""
    idx := beginX;
    digit := line[idx]

    for isDigit(digit) {
        str += digit
        idx++
        if idx == len(line) {
            break
        }
        digit = line[idx]
    }

    val, _ := strconv.Atoi(str)
    return val
}

func isGear(y, x int, matrix [][]string) (gear, bool)  {
    //top left
    dict := map[int]bool{}
    if y-1 >= 0 && x-1 >= 0 {
        if isDigit(matrix[y-1][x-1]) {
            number := readNumberInPos(y-1,x-1,matrix)
            dict[number] = true
        }
    }
    //top
    if y-1 >= 0 {
        if isDigit(matrix[y-1][x]) {
            number := readNumberInPos(y-1,x,matrix)
            dict[number] = true
        }
    }
    //top right
    if y-1 >= 0 && x+1 < len(matrix[y]) {
        if isDigit(matrix[y-1][x+1]) {
            number := readNumberInPos(y-1,x+1,matrix)
            dict[number] = true
        }
    }
    //right
    if x+1 < len(matrix[y]) {
        if isDigit(matrix[y][x+1]) {
            number := readNumberInPos(y,x+1,matrix)
            dict[number] = true
        }
    }
    //bottom right
    if y+1 < len(matrix) && x+1 < len(matrix[y]) {
        if isDigit(matrix[y+1][x+1]) {
            number := readNumberInPos(y+1,x+1,matrix)
            dict[number] = true
        }
    }
    //bottom
    if y+1 < len(matrix) {
        if isDigit(matrix[y+1][x]) {
            number := readNumberInPos(y+1,x,matrix)
            dict[number] = true
        }
    }
    //bottom left
    if y+1 < len(matrix) && x-1 >= 0 {
        if isDigit(matrix[y+1][x-1]) {
            number := readNumberInPos(y+1,x-1,matrix)
            dict[number] = true
        }
    }
    //left
    if x-1 >= 0 {
        if isDigit(matrix[y][x-1]) {
            number := readNumberInPos(y,x-1,matrix)
            dict[number] = true
        }
    }

    keys := []int{}
    for k := range dict {
        keys = append(keys, k)
    }
    if len(dict) == 2 {
        g := gear {
            numberA: keys[0],
            numberB: keys[1],
        }

        return g, true
    }
    
    return gear{}, false
}

func calculateSumA(matrix [][]string) int {
    sum := 0 

    for y := 0; y < len(matrix); y++ {
        for x := 0; x < len(matrix[y]); x++ {
            xv := matrix[y][x]
            number := ""
            foundPartNo := false
            for isDigit(xv) {
                number += xv
                if isPartNumber(y,x,matrix) {
                    foundPartNo = true
                }
                x++
                if x >= len(matrix[y]) {
                    break
                }
                xv = matrix[y][x]
            }
            if foundPartNo {
                val, _ := strconv.Atoi(number)
                sum += val
            }
        }
    }

    return sum
}

func isPartNumber(y, x int, matrix [][]string) bool {
    //top left
    if y-1 >= 0 && x-1 >= 0 {
        if matrix[y-1][x-1] != "." && !isDigit(matrix[y-1][x-1]) {
            return true
        }
    }
    //top
    if y-1 >= 0 {
        if matrix[y-1][x] != "." && !isDigit(matrix[y-1][x]) {
            return true
        }
    }
    //top right
    if y-1 >= 0 && x+1 < len(matrix[y]) {
        if matrix[y-1][x+1] != "." && !isDigit(matrix[y-1][x+1]) {
            return true
        }
    }
    //right
    if x+1 < len(matrix[y]) {
        if matrix[y][x+1] != "." && !isDigit(matrix[y][x+1]) {
            return true
        }
    }
    //bottom right
    if y+1 < len(matrix) && x+1 < len(matrix[y]) {
        if matrix[y+1][x+1] != "." && !isDigit(matrix[y+1][x+1]) {
            return true
        }
    }
    //bottom
    if y+1 < len(matrix) {
        if matrix[y+1][x] != "." && !isDigit(matrix[y+1][x]) {
            return true
        }
    }
    //bottom left
    if y+1 < len(matrix) && x-1 >= 0 {
        if matrix[y+1][x-1] != "." && !isDigit(matrix[y+1][x-1]) {
            return true
        }
    }
    //left
    if x-1 >= 0 {
        if matrix[y][x-1] != "." && !isDigit(matrix[y][x-1]) {
            return true
        }
    }
    return false
}
