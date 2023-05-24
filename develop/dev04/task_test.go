package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type sortStringTest struct {
	str      string
	expected string
}

type groupAnagramsTest struct {
	words    []string
	expected *map[string][]string
}

var sortStringTests = []sortStringTest{
	{"dabc", "abcd"},
	{"caaa", "aaac"},
	{"bike", "beik"},
}

var groupAnagramsTests = []groupAnagramsTest{
	{words: []string{"eat", "ate", "teA", "bike", "kibe", "cab", "bca", "Abc", "aBc", "a", "b"}, expected: &map[string][]string{"bike": {"bike", "kibe"}, "cab": {"abc", "bca", "cab"}, "eat": {"ate", "eat", "tea"}}},
	{words: []string{"eat", "eat", "eat", "ate", "teA", "bike", "kibe", "cab", "bca", "Abc", "aBc"}, expected: &map[string][]string{"bike": {"bike", "kibe"}, "cab": {"abc", "bca", "cab"}, "eat": {"ate", "eat", "tea"}}},
}

func TestSortString(t *testing.T) {
	for _, test := range sortStringTests {
		assert.Equal(t, test.expected, sortString(test.str))
	}
}

func TestGroupAnagrams(t *testing.T) {
	for _, test := range groupAnagramsTests {
		assert.Equal(t, test.expected, groupAnagrams(test.words))
	}
}