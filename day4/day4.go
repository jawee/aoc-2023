package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day4/atest.txt")
    // file, err := os.Open(pwd + "/day4/btest.txt")
    file, err := os.Open(pwd + "/day4/input.txt")

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
    lines := []string{}
    for scanner.Scan() {
        line := scanner.Text()
        lines = append(lines, line)
    }

    sum := calculateBSum(lines)
    return sum
}

func calculateBSum(lines []string) int {
    cards := map[string]int{}
    for i := range lines {
        cards[fmt.Sprintf("%d", i)] = 1;
    }

    for i, v := range lines {
        str := fmt.Sprintf("%d", i)
        for j := 0; j < cards[str]; j++ {
            sum := calculateMatchingNumbers(v)
            for k := 0; k < sum; k++ {
                if i+1+k == len(lines) {
                    fmt.Printf("%d why?\n", i+1+k)
                    os.Exit(1);
                }
                key := fmt.Sprintf("%d", i+1+k)
                cards[key]++
            }
        }
    }

    sum := 0
    for _, v := range cards {
        sum += v
    }

    return sum
}

func a(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        sum += calculateASum(line)
    }

    return sum
}

func calculateMatchingNumbers(line string) int {
    line = strings.Split(line, ":")[1]
    numbers := strings.Split(line, "|")
    correctNumbers := strings.Split(strings.Trim(numbers[0], " "), " ")
    myNumbers := strings.Split(strings.Trim(numbers[1], " "), " ")

    sum := 0
    for _, vm := range myNumbers {
        if vm == " " || vm == "" {
            continue
        }
        for _, vc := range correctNumbers {
            if vm == vc {
                sum++
            }
        }
    }
    return sum
}

func calculateASum(line string) int {
    line = strings.Split(line, ":")[1]
    numbers := strings.Split(line, "|")
    correctNumbers := strings.Split(strings.Trim(numbers[0], " "), " ")
    myNumbers := strings.Split(strings.Trim(numbers[1], " "), " ")

    sum := 0
    for _, vm := range myNumbers {
        if vm == " " || vm == "" {
            continue
        }
        for _, vc := range correctNumbers {
            if vm == vc {
                if sum == 0 {
                    sum = 1
                } else {
                    sum *= 2
                }
            }
        }
    }
    return sum
}

