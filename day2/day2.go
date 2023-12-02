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
    // file, err := os.Open(pwd + "/day2/atest.txt")
    file, err := os.Open(pwd + "/day2/input.txt")

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()
    // res := a(file)
    res := b(file)
    fmt.Printf("%d\n", res)
}

func getGameIdIfPossible(input string) int {
    maxRed := 12
    maxGreen := 13
    maxBlue := 14

    split := strings.Split(input, ":")

    gameIdSplit := strings.Split(split[0], " ")
    gameId, _ := strconv.Atoi(gameIdSplit[1])

    for _, frames := range strings.Split(strings.Trim(split[1], " "), ";") {

        frameValues := strings.Split(frames, ",") 
        for _, frameValue := range frameValues {
            frameSplit := strings.Split(strings.Trim(frameValue, " "), " ")
            no,_ := strconv.Atoi(frameSplit[0])
            if frameSplit[1] == "blue" && no > maxBlue {
                return 0
            }
            if frameSplit[1] == "red" && no > maxRed {
                return 0
            }
            if frameSplit[1] == "green" && no > maxGreen {
                return 0
            }
        }
    }

    return gameId;
}

func getPower(input string) int {
    maxRed := 0 
    maxGreen := 0
    maxBlue := 0

    split := strings.Split(input, ":")
    for _, frames := range strings.Split(strings.Trim(split[1], " "), ";") {

        frameValues := strings.Split(frames, ",") 
        for _, frameValue := range frameValues {
            frameSplit := strings.Split(strings.Trim(frameValue, " "), " ")
            no,_ := strconv.Atoi(frameSplit[0])
            if frameSplit[1] == "blue" && no > maxBlue {
                maxBlue = no
            }
            if frameSplit[1] == "red" && no > maxRed {
                maxRed = no
            }
            if frameSplit[1] == "green" && no > maxGreen {
                maxGreen = no
            }
        }
    }

    return maxRed * maxGreen * maxBlue;
}

func b(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        sum += getPower(line)
    }

    return sum
}

func a(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() {
        line := scanner.Text()
        sum += getGameIdIfPossible(line)
    }

    return sum
}
