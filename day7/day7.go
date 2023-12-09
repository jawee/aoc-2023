package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
    pwd, _ := os.Getwd()
    // file, err := os.Open(pwd + "/day7/atest.txt")
    file, err := os.Open(pwd + "/day7/input.txt")
    // file, err := os.Open(pwd + "/day7/btest.txt")
    
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
    sum := 0
    hands := []hand{}
    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, " ")

        bet, _ := strconv.Atoi(split[1])
        cards := strings.Split(split[0], "")

        hand := hand {
            cards: cards,
            bet: bet,
        }

        hands = append(hands, hand)
    }

    slices.SortStableFunc(hands, func(a, b hand) int {
        if isHigherRankB(a, b) {
            return 1
        }
        return -1
    })

    for i := range hands {
        fmt.Printf("%d * %+v\n", (i+1), hands[i])
        sum += (i+1) * hands[i].bet
    }
    return sum
}

func isHigherRankB(higher, lower hand) bool {
    hType := getTypeB(higher)
    lType := getTypeB(lower)
    if hType < lType {
        // fmt.Printf("%s < %s\n", hType, lType)
        return true
    }

    if hType == lType {
        for i := range higher.cards {
            h := higher.cards[i]
            l := lower.cards[i]
            if ranksB[h] < ranksB[l] {
                // fmt.Printf("%d < %d\n", ranksB[h], ranksB[l])
                return true
            }
            if ranksB[h] > ranksB[l] {
                // fmt.Printf("%d > %d\n", ranksB[h], ranksB[l])
                return false
            }
        }
    }

    // fmt.Printf("%s > %s\n", hType, lType)
    return false
}

func getTypeB(a hand) handType {
    counts := map[string]int{}
    for _, v := range a.cards {
        counts[v]++
    }

    jokers := counts["J"]

    keys := []string{}
    for k := range counts {
        keys = append(keys, k)
    }

    if len(keys) == 1 {
        return Five
    }

    if len(keys) == 2 {
        for _, k := range keys {
            if k == "J" {
                continue
            }
            if counts[k] == 4 {
                if jokers == 1 {
                    return Five
                }
                return Four
            }
            if counts[k] == 3 {
                if jokers == 2 {
                    return Five
                }
                if jokers == 1 {
                    return Four
                }
                return FullHouse
            }
        }
    }

    if len(keys) == 3 {
        for _, k := range keys {
            if k == "J" {
                continue
            }
            if counts[k] == 3 {
                if jokers == 2 {
                    return Five
                }
                if jokers == 1 {
                    return Four
                }
                return Three
            }
            if counts[k] == 2 {
                if jokers == 2 {
                    return Four
                }
                if jokers == 1 {
                    return FullHouse
                }
                return TwoPair
            }
        }
    }
    if len(keys) == 4 {
        for _, k := range keys {
            if k == "J" {
                continue
            }
            if counts[k] == 2 {
                if jokers == 3 {
                    return Five
                }
                if jokers == 2 {
                    return Four
                }
                if jokers == 1 {
                    return Three
                }
                return OnePair
            }
        }
    }

    if jokers == 5 {
        return Five
    }
    if jokers == 4 {
        return Five
    }
    if jokers == 3 {
        return Four
    }
    if jokers == 2 {
        return Three
    }
    if jokers == 1 {
        return OnePair
    }

    return HighCard
}

type hand struct {
    cards []string
    bet int
}

var ranksB = map[string]int {
    "A": 1,
    "K": 2,
    "Q": 3,
    "T": 5,
    "9": 6,
    "8": 7,
    "7": 8,
    "6": 9,
    "5": 10,
    "4": 11,
    "3": 12,
    "2": 13,
    "J": 14,
}

var ranksA = map[string]int {
    "A": 1,
    "K": 2,
    "Q": 3,
    "J": 4,
    "T": 5,
    "9": 6,
    "8": 7,
    "7": 8,
    "6": 9,
    "5": 10,
    "4": 11,
    "3": 12,
    "2": 13,
}

func a(file io.Reader) int {
    scanner := bufio.NewScanner(file)
    sum := 0
    hands := []hand{}
    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, " ")

        bet, _ := strconv.Atoi(split[1])
        cards := strings.Split(split[0], "")

        hand := hand {
            cards: cards,
            bet: bet,
        }

        hands = append(hands, hand)
    }

    slices.SortStableFunc(hands, func(a, b hand) int {
        if isHigherRank(a, b) {
            return 1
        }
        return -1
    })

    for i := range hands {
        sum += (i+1) * hands[i].bet
    }
    return sum
}

func isHigherRank(higher, lower hand) bool {
    hType := getType(higher)
    lType := getType(lower)
    if hType < lType {
        // fmt.Printf("%s < %s\n", hType, lType)
        return true
    }

    if hType == lType {
        for i := range higher.cards {
            h := higher.cards[i]
            l := lower.cards[i]
            if ranksA[h] < ranksA[l] {
                // fmt.Printf("%d < %d\n", ranks[h], ranks[l])
                return true
            }
            if ranksA[h] > ranksA[l] {
                // fmt.Printf("%d > %d\n", ranks[h], ranks[l])
                return false
            }
        }
    }

    // fmt.Printf("%s > %s\n", hType, lType)
    return false
}

func sortHands(a, b hand) bool {
    return true
}

func getType(a hand) handType {
    counts := map[string]int{}
    for _, v := range a.cards {
        counts[v]++
    }

    keys := []string{}
    for k := range counts {
        keys = append(keys, k)
    }

    if len(keys) == 1 {
        return Five
    }

    if len(keys) == 2 {
        for _, k := range keys {
            if counts[k] == 4 {
                return Four
            }
            if counts[k] == 3 {
                return FullHouse
            }
        }
    }

    if len(keys) == 3 {
        for _, k := range keys {
            if counts[k] == 3 {
                return Three
            }
            if counts[k] == 2 {
                return TwoPair
            }
        }
    }
    if len(keys) == 4 {
        for _, k := range keys {
            if counts[k] == 2 {
                return OnePair
            }
        }
    }

    return HighCard
}

type handType int
const (
    Five handType = iota
    Four 
    FullHouse 
    Three 
    TwoPair 
    OnePair
    HighCard
)
func (s handType) String() string {
	switch s {
	case Five:
		return "Five"
	case Four:
		return "Four"
	case FullHouse:
		return "Full House"
	case Three:
		return "Three"
	case TwoPair:
		return "Two Pairs"
	case OnePair:
		return "One Pair"
	case HighCard:
		return "High Card"
	}
	return "unknown"
}
