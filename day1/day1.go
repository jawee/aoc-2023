package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day1/atest.txt")
    // file, err := os.Open(pwd + "/day1/btest.txt")
    file, err := os.Open(pwd + "/day1/input.txt")

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    // res := a(file)
    res := b(file)
    fmt.Printf("%d\n", res)
}

func replaceWithNumber(s string) string {
    result := ""
    for len(s) > 0 {
        if strings.HasPrefix(s, "one") {
            result += "1"
            s = s[len("one")-1:]
            continue;
        }
        if strings.HasPrefix(s, "two") {
            result += "2"
            s = s[len("two")-1:]
            continue;
        }
        if strings.HasPrefix(s, "three") {
            result += "3"
            s = s[len("three")-1:]
            continue;
        }
        if strings.HasPrefix(s, "four") {
            result += "4"
            s = s[len("four")-1:]
            continue;
        }
        if strings.HasPrefix(s, "five") {
            result += "5"
            s = s[len("five")-1:]
            continue;
        }
        if strings.HasPrefix(s, "six") {
            result += "6"
            s = s[len("six")-1:]
            continue;
        }
        if strings.HasPrefix(s, "seven") {
            result += "7"
            s = s[len("seven")-1:]
            continue;
        }
        if strings.HasPrefix(s, "eight") {
            result += "8"
            s = s[len("eight")-1:]
            continue;
        }
        if strings.HasPrefix(s, "nine") {
            result += "9"
            s = s[len("nine")-1:]
            continue;
        }
        result += string(s[0])
        s = s[1:]
    }
    return result
}

func b(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        str := replaceWithNumber(line)
        arr := []string{}
        for _, v := range(str) {
            if unicode.IsDigit(v) {
                arr = append(arr, string(v))
            }
        }

        number, err := strconv.Atoi(fmt.Sprintf("%s%s", arr[0], arr[len(arr)-1]))
        if err != nil {
            log.Fatalf("Err: %s\n", err)
            os.Exit(1)
        }
        sum += number
    }

    return sum
}

func a(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        arr := []string{}
        for _, v := range(line) {
            if unicode.IsDigit(v) {
                arr = append(arr, string(v))
            }
        }

        number, err := strconv.Atoi(fmt.Sprintf("%s%s", arr[0], arr[len(arr)-1]))
        if err != nil {
            log.Fatalf("Err: %s\n", err)
            os.Exit(1)
        }
        sum += number
    }

    return sum
}
