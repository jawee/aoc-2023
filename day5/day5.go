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
    // file, err := os.Open(pwd + "/day5/atest.txt")
    // file, err := os.Open(pwd + "/day5/btest.txt")
    file, err := os.Open(pwd + "/day5/input.txt")

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    res := a(file)
    // res := b(file)
    fmt.Printf("%d\n", res)
}

func b(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        // line := scanner.Text()
    }

    return sum
}

type mapLine struct {
    dest int
    src int
    length int
}

func newMapLine(s string) mapLine {
    vals := strings.Split(s, " ")
    dest, _ := strconv.Atoi(vals[0])
    src, _ := strconv.Atoi(vals[1])
    length, _ := strconv.Atoi(vals[2])
    return mapLine {
        dest: dest,
        src: src,
        length: length,
    }
}

func a(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    sum := 0
    scanner.Scan() 
    line := scanner.Text()
    seeds := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
    seedsArr := []int{}
    for _, v := range seeds {
        n, _ := strconv.Atoi(v)
        seedsArr = append(seedsArr, n)
    }

    // fmt.Printf("Seeds: %v\n", seedsArr);
    scanner.Scan(); //header
    for scanner.Scan() { //empty line
        // fmt.Printf("CurrSeedsArr: %v\n", seedsArr)
        nextDest := []int{}
        line = scanner.Text()
        // fmt.Printf("Header: %s\n", line)
        scanner.Scan()
        line = scanner.Text()
        destArr := []mapLine{}
        for line != "" {
            m := newMapLine(line)
            destArr = append(destArr, m)
            // fmt.Printf("%+v\n", m)
            scanner.Scan()
            line = scanner.Text()
        }

        // fmt.Printf("%+v\n", destArr)
        for _, v := range seedsArr {
            found := false
            for _, r := range destArr {
                if r.src <= v && v <= r.src+r.length {
                    dest := (v-r.src) + r.dest
                    // fmt.Printf("%d matches %+v. New dest: %d\n", v, r, dest)
                    nextDest = append(nextDest, dest)
                    found = true
                    break;
                }
            }
            if !found {
                // fmt.Printf("%d matches nothing\n", v)
                nextDest = append(nextDest, v)
            }
        }

        seedsArr = nextDest
    }

    sum = getMin(seedsArr)

    return sum
}

func getMin(values []int) int {
    min := values[0]
    for _, v := range values {
        if (v < min) {
            min = v
        }
    }

    return min
}
