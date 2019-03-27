// utils provides short, unsafe versions of functions for use in hacking
package utils

import "strconv"

// AtoI is an unsafe verison of strconv.Atoi
// it panics if there's any error
func AtoI(s string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return out
}
