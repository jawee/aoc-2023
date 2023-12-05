package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCalculateRange(t *testing.T) {
    testCases := []struct {
        seed seedRange
        dest mapLine
        initial int
        length int
        start int
        end int
    }{
        {
            seed: seedRange {
                initial: 79,
                length: 14,
            },
            dest: mapLine {
                dest: 52,
                src: 50,
                length: 42,
            },
            initial: 81,
            length: 13,
            start: 79,
            end: 91,
        },
        {
            seed: seedRange {
                initial: 79,
                length: 14,
            },
            dest: mapLine {
                dest: 52,
                src: 50,
                length: 42,
            },
            initial: 81,
            length: 13,
            start: 79,
            end: 91,
        },
        {
            seed: seedRange {
                initial: 79,
                length: 15,
            },
            dest: mapLine {
                dest: 52,
                src: 50,
                length: 42,
            },
            initial: 81,
            length: 13,
            start: 79,
            end: 91,
        },
        {
            seed: seedRange {
                initial: 79,
                length: 30,
            },
            dest: mapLine {
                dest: 52,
                src: 50,
                length: 42,
            },
            initial: 81,
            length: 13,
            start: 79,
            end: 91,
        },
    }
        
    for _, v := range testCases {
        a, b := calculateRange(v.seed, v.dest)

        if a.initial != v.initial {
            t.Fatalf("Expected seed.initial '%d', got '%d'\n", v.initial, a.initial)
        }

        if a.length != v.length {
            t.Fatalf("Expected seed.length '%d', got '%d'\n", v.length,  a.length)
        }

        if b.start != v.start {
            t.Fatalf("Expected range.start '%d', got '%d'\n", v.start, b.start)
        }
        if b.end != v.end {
            t.Fatalf("Expected range.end '%d', got '%d'\n", v.end, b.end)
        }
    }
}

func TestGetUnusedRanges(t *testing.T) {
    testCases := []struct {
        seed seedRange
        used []numberRange
        expected []seedRange
    }{
        {
            seed: seedRange {
                initial: 1,
                length: 9,
            },
            used: []numberRange {
                {
                    start: 1,
                    end: 3,
                },
            },
            expected: []seedRange {
                {
                    initial: 4,
                    length: 6,
                },
            },
        },
        {
            seed: seedRange {
                initial: 1,
                length: 9,
            },
            used: []numberRange {
                {
                    start: 1,
                    end: 3,
                },
                {
                    start: 7,
                    end: 9,
                },
            },
            expected: []seedRange {
                {
                    initial: 4,
                    length: 3,
                },
            },
        },
        {
            seed: seedRange {
                initial: 1,
                length: 9,
            },
            used: []numberRange {
            },
            expected: []seedRange {
                {
                    initial: 1,
                    length: 9,
                },
            },
        },
        {
            seed: seedRange {
                initial: 1,
                length: 11,
            },
            used: []numberRange {
                {
                    start: 1,
                    end: 3,
                },
                {
                    start: 7,
                    end: 9,
                },
            },
            expected: []seedRange {
                {
                    initial: 4,
                    length: 3,
                },
                {
                    initial: 10,
                    length: 2,
                },
            },
        },
    }

    for i, v := range testCases {
        fmt.Printf("=========TESTCASE %d==========\n", i+1)

        res := getUnusedRanges2(v.seed, v.used)

        if !reflect.DeepEqual(res, v.expected) {
            t.Fatalf("Failed: got %+v, expected %+v\n", res, v.expected)
        }

    }
}

