package main

import (
	"fmt"
	"sort"
	"strings"
)

func sortString(word string) string {
	letters := []rune(word)
	sort.Slice(letters, func(i, j int) bool {
		return letters[i] < letters[j]
	})
	return string(letters)
}

func groupAnagrams(words []string) *map[string][]string {
	noDuplicate := make(map[string]bool, len(words))
	m_ := make(map[string][]string)
	m := make(map[string][]string)

	for _, upperWord := range words {

		word := strings.ToLower(upperWord)
		if len(word) < 2 {
			continue 
		}

		if !noDuplicate[word] {
			m_[sortString(word)] = append(m_[sortString(word)], word)
			noDuplicate[word] = true
		}
	}

	for _, v := range m_ {
		firstWord := v[0]
		sort.Strings(v)
		m[firstWord] = v
	}

	return &m
}

func main() {
	fmt.Println(*groupAnagrams([]string{"eat", "ate", "tea", "bike", "kibe", "cab", "bca", "abc", "abc"}))
}