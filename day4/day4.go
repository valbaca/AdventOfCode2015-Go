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

func findPrefix(in string, prefix string) string {
	key := in
	n := 0
	for {
		hash := GetHash(key, n)
		if strings.HasPrefix(hash, prefix) {
			return strconv.Itoa(n)
		}
		n++
	}
	return "exit"
}

func GetHash(key string, n int) string {
	data := []byte(fmt.Sprintf("%s%d", key, n))
	return fmt.Sprintf("%x", md5.Sum(data))
}
