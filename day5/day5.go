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

type seedRange struct {
    initial int
    length int
}

func b(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    scanner.Scan() 
    line := scanner.Text()
    seeds := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
    seedsArr := []seedRange{}

    i := 0
    for i < len(seeds) {
        initialSeed, _ := strconv.Atoi(seeds[i])
        length, _ := strconv.Atoi(seeds[i+1])

        s := seedRange {
            initial: initialSeed,
            length: length,
        }

        seedsArr = append(seedsArr, s)

        i = i+2
    }

    i = 0
    maps := [][]mapLine{}
    scanner.Scan(); //empty line
    for scanner.Scan() { //header
        scanner.Text()
        scanner.Scan()
        line = scanner.Text()
        
        maps = append(maps, []mapLine{})
        for line != "" {
            m := newMapLine(line)
            maps[i] = append(maps[i], m)
            if !scanner.Scan() { break; }
            line = scanner.Text()
        }
        i++
    }


    j := 0;
    for true {
        currentSeekValue := j
        for i := len(maps)-1; 0 <= i; i-- {
            for _, v := range maps[i] {
                if isBetween(currentSeekValue, v.dest, v.dest+v.length-1) {
                    currentSeekValue = (v.src-v.dest)+currentSeekValue
                    break
                }
            }
        }
        if containsValue(currentSeekValue, seedsArr) {
            return j
        }
        j++
    }

    return -1
}

func containsValue(currentSeekValue int, seeds []seedRange) bool {
    for _, v := range seeds {
        if currentSeekValue < v.initial {
            continue
        }
        if isBetween(currentSeekValue, v.initial, v.initial+v.length-1) {
            return true
        }
    }

    return false
}

func isBetween(val, min, max int) bool {
    if min <= val && val <= max {
        return true
    }

    return false
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

    scanner.Scan(); //header
    for scanner.Scan() { //empty line
        nextDest := []int{}
        line = scanner.Text()
        scanner.Scan()
        line = scanner.Text()
        destArr := []mapLine{}
        for line != "" {
            m := newMapLine(line)
            destArr = append(destArr, m)
            scanner.Scan()
            line = scanner.Text()
        }

        for _, v := range seedsArr {
            found := false
            for _, r := range destArr {
                if r.src <= v && v <= r.src+r.length {
                    dest := (v-r.src) + r.dest
                    nextDest = append(nextDest, dest)
                    found = true
                    break;
                }
            }
            if !found {
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
