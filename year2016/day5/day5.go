package day5

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"valbaca.com/advent/elf"
)

/*
TIL: I thought this would be trivially parallelizable, but saw hardly any perf gains.
Fuzting around between strings, []bytes, runes, and ints was a drag.
*/

type Day5 struct{}

func (d Day5) Part1(input string) interface{} {
	pass := strings.Builder{}
	inputBytes := []byte(input)
	zeros := []byte("00000")
	var i int64 = 0
	found := 0
	for found < 8 {
		data := append(inputBytes, []byte(strconv.FormatInt(i, 10))...)
		hexsum := fmt.Sprintf("%x", md5.Sum(data))
		if bytes.Equal(zeros, []byte(hexsum[:5])) {
			pass.WriteRune(rune(hexsum[5]))
			found++
		}
		i++
	}
	return pass.String()
}

/* Attempted to go parallel...but it's actually slower than the single threaded */

const RUNNERS = 4 // Tried between 1-8 runners

func Part1Async(input string) string {
	inputBytes := []byte(input)
	zeros := []byte("00000")

	findHash := func(id int64, in <-chan int64, out chan<- result) {
		for {
			i := <-in
			data := append(inputBytes, []byte(strconv.FormatInt(i, 10))...)
			hexsum := fmt.Sprintf("%x", md5.Sum(data))
			if bytes.Equal(zeros, []byte(hexsum[:5])) {
				out <- result{i, string(hexsum[5])}
			}
		}
	}
	c := make(chan result)
	xs := make(chan int64)
	for i := 0; i < RUNNERS; i++ {
		go findHash(int64(i), xs, c)
	}
	go func() {
		j := int64(0)
		for {
			xs <- j
			j++
		}
	}()
	results := make([]result, 8)
	for i := 0; i < 8; i++ {
		results[i] = <-c
	}
	fmt.Println(results)
	sort.Slice(results, func(i, j int) bool {
		return results[i].i < results[j].i
	})
	fmt.Println(results)
	sb := strings.Builder{}
	for _, r := range results {
		sb.WriteString(r.char)
	}
	return sb.String()
}

type result struct {
	i    int64
	char string
}

func (d Day5) Part2(input string) interface{} {
	pass := [8]rune{}
	inputBytes := []byte(input)
	zeros := []byte("00000")
	var i int64 = 0
	found := 0
	for found < 8 {
		data := append(inputBytes, []byte(strconv.FormatInt(i, 10))...)
		hexsum := fmt.Sprintf("%x", md5.Sum(data))
		if bytes.Equal(zeros, []byte(hexsum[:5])) {
			posRune := rune(hexsum[5])
			if posRune >= '0' && posRune <= '7' {
				pos := elf.UnsafeAtoi(string(posRune))
				if pass[pos] == 0 {
					pass[pos] = rune(hexsum[6])
					found++
				}
			}
		}
		i++
	}
	sb := strings.Builder{}
	for _, r := range pass {
		sb.WriteRune(r)
	}
	return sb.String()
}
