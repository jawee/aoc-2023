package main

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)


func TestSort(t *testing.T) {
    a:= hand {
        cards: strings.Split("32T3K", ""),
        bet: 765,
    }

    b := hand {
        cards: strings.Split("T55J5", ""),
        bet: 684,
    }
    c := hand {
        cards: strings.Split("KK677", ""),
        bet: 28,
    }
    d := hand {
        cards: strings.Split("KTJJT", ""),
        bet: 220,
    }

    e := hand {
        cards: strings.Split("QQQJA", ""),
        bet: 483,
    }
    hands := []hand{a, b, c, d, e}

    slices.SortFunc(hands, func(a, b hand) int {
        if isHigherRank(a, b) {
            return 1
        }
        return -1
    })

    fmt.Printf("%+v\n", hands)
    if hands[0].bet != 765 {
        t.Fatalf("Failed on 0\n")
    }
    if hands[1].bet != 220 {
        t.Fatalf("Failed on 1\n")
    }
    if hands[2].bet != 28 {
        t.Fatalf("Failed on 2\n")
    }
    if hands[3].bet != 684 {
        t.Fatalf("Failed on 3\n")
    }
    if hands[4].bet != 483 {
        t.Fatalf("Failed on 4\n")
    }
}

func TestIsHigherHand(t *testing.T) {
    testCases := []struct {
        a string
        b string
        exp bool
    }{
        {
            a: "32T3K",
            b: "KK677",
            exp: false,
        },
        {
            a: "KK677",
            b: "KTJJT",
            exp: true,
        },
        {
            a: "T55J5",
            b: "QQQJA",
            exp: false,
        },
        {
            a: "T55J5", //684
            b: "KTJJT", //220
            exp: true,  // unsure?
        },
        {
            a: "32T3K",
            b: "KK677",
            exp: false,
        },
        {
            a: "32T3K",
            b: "KTJJT",
            exp: false,
        },
        {
            a: "KK677",
            b: "KTJJT",
            exp: true,
        },
        {
            a: "QQQJA",
            b: "T55J5",
            exp: true,
        },
    }

    for i, v := range testCases {
        a := getHand(v.a)
        b := getHand(v.b)
        res := isHigherRank(a, b)

        if res != v.exp {
            t.Fatalf("Case %d: Got '%v' expected '%v'\n", i, res, v.exp)
        }
    }
}

func TestGetHandType(t *testing.T) {
    testCases := []struct {
        h string    
        exp handType
    }{
        {
            h: "AAAAA",
            exp: Five,
        },
        {
            h: "AA8AA",
            exp: Four,
        },
        {
            h: "23332",
            exp: FullHouse,
        },
        {
            h: "TTT98",
            exp: Three,
        },
        {
            h: "23432",
            exp: TwoPair,
        },
        {
            h: "A23A4",
            exp: OnePair,
        },
        {
            h: "23456",
            exp: HighCard,
        },
        {
            h: "KTJJT",
            exp: TwoPair,
        },
        {
            h: "KK677",
            exp: TwoPair,
        },
    }


    for i, v := range testCases {
        h := getHand(v.h)
        res := getType(h)

        if res != v.exp {
            t.Fatalf("Case %d '%s': Got '%s' expected '%s'\n", i, v.h, res, v.exp)
        }
    }
}

func getHand(a string) hand {
    hand := hand {
        cards: strings.Split(a, ""),
        bet: 0,
    }

    return hand
}
