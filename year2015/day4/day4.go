// Package day4
// TIL: go has md5 sum and others built-in
// TIL: hacking checksums is slow :P
package day4

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func Part1(in string) string {
	return findPrefix(in, "00000")
}

func Part2(in string) string {
	return findPrefix(in, "000000")
}

type Soln struct {
	n int
	s string
}

const CONC = 3

func work(jobs chan int, found chan Soln, key, prefix string) {
	for n := range jobs {
		hash := getHash(key, n)
		if strings.HasPrefix(hash, prefix) {
			found <- Soln{n, strconv.Itoa(n)}
			return
		} else {
			select {
			case ans := <-found:
				found <- ans
				return
			default:
				jobs <- n + CONC
			}
		}
	}
}

func findPrefix(in string, prefix string) string {
	key := in
	found := make(chan Soln)
	jobs := make(chan int, CONC)
	for i := 0; i < CONC; i++ {
		go work(jobs, found, key, prefix)
	}
	for i := 0; i < CONC; i++ {
		jobs <- i
	}
	return (<-found).s
}

func getHash(key string, n int) string {
	data := []byte(fmt.Sprintf("%s%d", key, n))
	return fmt.Sprintf("%x", md5.Sum(data))
}
