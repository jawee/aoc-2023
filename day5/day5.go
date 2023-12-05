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
    // res := a(file)
    res := b(file)
    fmt.Printf("%d\n", res)
}

type seedRange struct {
    initial int
    length int
}
type numberRange struct {
    start int
    end int
}

func b(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    sum := 0
    scanner.Scan() 
    line := scanner.Text()
    seeds := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
    seedsArr := []seedRange{}

    //numbers are too big
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

    // fmt.Printf("Seeds: %v\n", seedsArr);
    scanner.Scan(); //header
    for scanner.Scan() { //empty line
        // fmt.Printf("loop2\n")
        // fmt.Printf("CurrSeedsArr: %v\n", seedsArr)
        nextDest := []seedRange{}
        line = scanner.Text()
        // fmt.Printf("Header: %s\n", line)
        scanner.Scan()
        line = scanner.Text()
        destArr := []mapLine{}
        for line != "" {
            // fmt.Printf("loop1\n")
            m := newMapLine(line)
            destArr = append(destArr, m)
            // fmt.Printf("%+v\n", m)
            scanner.Scan()
            line = scanner.Text()
        }

        // fmt.Printf("seedsArrlen %d\n", len(seedsArr))
        for _, seed := range seedsArr {
            usedRanges := []numberRange{}
            for _, dest := range destArr {
                // fmt.Printf("loop\n")
                seedMin := seed.initial
                seedMax := seed.initial+seed.length-1
                destMin := dest.src
                destMax := dest.src+dest.length-1

                // ignore dest range below and above
                if destMax < seedMin || seedMax < destMin {
                    // fmt.Printf("continue\n")
                    continue
                }

                // will use whole seed
                // if destMin < seedMin && seedMax < destMax {
                //     newSeed, usedRange := calculateRange(seed, dest)
                //     nextDest = append(nextDest, newSeed)
                //
                //     usedRanges = append(usedRanges, usedRange)
                // }

                if isBetween(seedMin, destMin, destMax) || isBetween(seedMax, destMin, destMax) || isBetween(destMin, seedMin, seedMax) || isBetween(destMax, seedMin, seedMax) {
                    newSeed, usedRange := calculateRange(seed, dest)
                    // fmt.Printf("After calculate range\n")
                    nextDest = append(nextDest, newSeed)
                    usedRanges = append(usedRanges, usedRange)

                    // fmt.Printf("After appends\n")
                }
            }

            // add to nextDest
            unusedDests := getUnusedRanges2(seed, usedRanges)
            nextDest = append(nextDest, unusedDests[:]...)
        }

        seedsArr = nextDest
    }

    sum = getMinSeedRange(seedsArr)

    return sum
}


func getUnusedRanges2(seed seedRange, usedRanges []numberRange) []seedRange  {
    res := []seedRange{seed}
    for _, v := range usedRanges {
        tmpRes := []seedRange{}
        for _, s := range res {
            if s.initial == v.start {
                // fmt.Printf("a\n")
                s1 := seedRange {
                    initial: v.end+1,
                    length: s.initial+s.length-1-v.end,
                }
                tmpRes = append(tmpRes, s1)
            } else if s.initial < v.start && s.initial+s.length-1 < v.end {
                // fmt.Printf("b\n")
                s1 := seedRange {
                    initial: s.initial,
                    length: s.initial+s.length-1-v.start,
                }
                s2 := seedRange {
                    initial: v.end+1,
                    length: s.initial+s.length-1-(v.end+1),
                }

                tmpRes = append(tmpRes, s1, s2)
            } else if s.initial < v.start && s.initial+s.length-1 == v.end {
                // fmt.Printf("c\n")
                s1 := seedRange {
                    initial: s.initial,
                    length: s.initial+s.length-v.start,
                }
                tmpRes = append(tmpRes, s1)
            } else if v.start < s.initial+s.length-1 && v.end < s.initial+s.length-1 {
                // fmt.Printf("d\n")
                s1 := seedRange {
                    initial: s.initial,
                    length: s.initial+s.length-v.end,
                }
                s2 := seedRange {
                    initial: v.end+1,
                    length: s.initial+s.length-1-v.end,
                }
                tmpRes = append(tmpRes, s1, s2)
            } else {
                // fmt.Printf("e\n")
                tmpRes = append(tmpRes, s)
            }
        }
        // fmt.Printf("%+v\n", tmpRes)
        res = tmpRes
    }

    return res
}
func getUnusedRanges(seed seedRange, usedRanges []numberRange) []seedRange  {
    fmt.Printf("in getUnusedRanges. %+v\n", seed)
    // for _, u := range usedRanges {
    //     fmt.Printf("numberRange: %+v\n", u)
    //     fmt.Printf("%+v\n", unused)
    //     for _, v := range unused {
    //     }
    //
    // }

    mapz := map[int]bool{}
    for _, v := range usedRanges {
        for i := v.start; i <= v.end; i++ {
            mapz[i] = true
        }
    }

    fmt.Printf("Mapz: %+v\n", mapz)

    unused := []seedRange{}
    currSeed := seedRange {}
    currentSeed := false

    for i := seed.initial; i < seed.initial+seed.length-1; i++ {
        // fmt.Printf("i: '%d' maps[i]: '%v'\n", i, mapz[i])
        // fmt.Printf("currSeed: %+v\n", currSeed)
        if !mapz[i] && !currentSeed {
            // fmt.Printf("Initializing seed\n")
            currSeed = seedRange {
                initial: i,
                length: 1,
            }
            currentSeed = true
        } else if !mapz[i] {
            // fmt.Printf("Incrementing length\n")
            currSeed.length += 1
        } else {
            // fmt.Printf("Clear\n")
            if currentSeed {
                unused = append(unused, currSeed)
                currentSeed = false
            }
        }
    }

    if currentSeed {
        unused = append(unused, currSeed)
    }

    return unused
}

func isBetween(val, min, max int) bool {
    if min <= val && val <= max {
        return true
    }

    return false
}

func calculateRange(in seedRange, dest mapLine) (seedRange, numberRange) {
    // fmt.Printf("Calculate range\n")
    delta := (-1*dest.src)+dest.dest

    newLen := in.length
    if in.initial+in.length-1 < dest.src+dest.length-1 {
        //using all
    }

    if in.initial+in.length-1 > dest.src+dest.length-1 {
        newLen = in.length - ((in.initial+in.length-1) - (dest.src+dest.length-1))
    }

    // fmt.Printf("newLen: '%d'\n", newLen)
    seed := seedRange {
        initial: in.initial + delta,
        length: newLen,
    }

    used := numberRange {
        start: in.initial,
        end: in.initial+newLen-1,
    }
    return seed, used
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

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}

func getMinSeedRange(values []seedRange) int {
    min := values[0].initial
    for _, v := range values {
        if v.initial < min {
            min = v.initial
        }
    }
    return min
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
