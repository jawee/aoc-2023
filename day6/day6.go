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
    // file, err := os.Open(pwd + "/day6/atest.txt")
    file, err := os.Open(pwd + "/day6/input.txt")

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
    sum := 1
    scanner.Scan()
    tmpTimes := strings.Split(strings.Split(scanner.Text(), ":")[1], " ")
    scanner.Scan()
    tmpDists := strings.Split(strings.Trim(strings.Split(scanner.Text(), ":")[1], " "), " ")

    times := []string{}
    for _, v := range tmpTimes {
        if v == "" {
            continue
        }
        times = append(times, v)
    }
    dists := []string{}
    for _, v := range tmpDists {
        if v == "" {
            continue
        }
        dists = append(dists, v)
    }

    time, _ := strconv.Atoi(strings.Join(times, ""))
    dist, _ := strconv.Atoi(strings.Join(dists, ""))

    timesToWin := 0
    for j := 1; j < time; j++ {
        timeToTravel := time - j
        speed := j

        if timeToTravel*speed > dist {
            timesToWin++
        }
    }
    sum = sum * timesToWin

    return sum
}

func a(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    sum := 1
    scanner.Scan()
    tmpTimes := strings.Split(strings.Split(scanner.Text(), ":")[1], " ")
    scanner.Scan()
    tmpDists := strings.Split(strings.Trim(strings.Split(scanner.Text(), ":")[1], " "), " ")

    times := []string{}
    for _, v := range tmpTimes {
        if v == "" {
            continue
        }
        times = append(times, v)
    }
    dists := []string{}
    for _, v := range tmpDists {
        if v == "" {
            continue
        }
        dists = append(dists, v)
    }

    for t, timeStr := range times {
        maxDist, _ := strconv.Atoi(dists[t])
        time, _ := strconv.Atoi(timeStr)

        timesToWin := 0
        for j := 1; j < time; j++ {
            timeToTravel := time - j
            speed := j

            if timeToTravel*speed > maxDist {
                timesToWin++
            }
        }
        sum = sum * timesToWin
    }

    return sum
}
