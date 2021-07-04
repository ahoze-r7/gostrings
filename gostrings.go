// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package gostrings

import (
	"github.com/ahoze-r7/gostrings2"
)

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	gostrings2.ReverseRunes2("sd")

	return string(r)
}
