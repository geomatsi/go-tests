//
// mix of simple array examples
//

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPivotIndex(t *testing.T) {
	assert.Equal(t, "0", addBinary("0", "0"))
	assert.Equal(t, "1", addBinary("1", "0"))
	assert.Equal(t, "10", addBinary("1", "1"))
	assert.Equal(t, "110", addBinary("11", "11"))
	assert.Equal(t, "100", addBinary("11", "1"))
	assert.Equal(t, "10101", addBinary("1010", "1011"))
	assert.Equal(t, "10110", addBinary("111", "1111"))
}

func TestStrStr(t *testing.T) {
	assert.Equal(t, strStr("", "a"), -1)
	assert.Equal(t, strStr("hello", ""), 0)
	assert.Equal(t, strStr("short", "tooooolong"), -1)
	assert.Equal(t, strStr("hello", "ll"), 2)
	assert.Equal(t, strStr("aaaaa", "bba"), -1)
	assert.Equal(t, strStr("hello", "hello"), 0)
	assert.Equal(t, strStr("ababab", "ab"), 0)
	assert.Equal(t, strStr("ababab", "ba"), 1)
}

func TestLongestCommonPrefix(t *testing.T) {
	assert.Equal(t, longestCommonPrefix([]string{}), "")
	assert.Equal(t, longestCommonPrefix([]string{"aaa", "aa"}), "aa")
	assert.Equal(t, longestCommonPrefix([]string{"flower", "flow", "flight"}), "fl")
	assert.Equal(t, longestCommonPrefix([]string{"dog", "racecar", "car"}), "")
	assert.Equal(t, longestCommonPrefix([]string{"aaaaa", "bbbbbbb", "ccccccccccc"}), "")
	assert.Equal(t, longestCommonPrefix([]string{"abcd", "abcdabcd", "abcdabcdabcd"}), "abcd")
	assert.Equal(t, longestCommonPrefix([]string{"axxxxxx", "ayyyyyyy", "abcd"}), "a")
}
