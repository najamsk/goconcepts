package validator

import (
	"math"
	"sort"
	"strings"

	"github.com/agnivade/levenshtein"
)

// Given an invalid input string and a list of valid options, returns a filtered
// list of valid options sorted based on their similarity with the input.
func SuggestionList(input string, options []string) []string {
	var results []string
	optionsByDistance := map[string]int{}

	for _, option := range options {
		distance := lexicalDistance(input, option)
		threshold := calcThreshold(input)
		if distance <= threshold {
			results = append(results, option)
			optionsByDistance[option] = distance
		}
	}

	sort.Slice(results, func(i, j int) bool {
		return optionsByDistance[results[i]] < optionsByDistance[results[j]]
	})
	return results
}

func calcThreshold(a string) (threshold int) {
	// the logic is copied from here
	// https://github.com/graphql/graphql-js/blob/47bd8c8897c72d3efc17ecb1599a95cee6bac5e8/src/jsutils/suggestionList.ts#L14
	threshold = int(math.Floor(float64(len(a))*0.4) + 1)

	if threshold < 1 {
		threshold = 1
	}
	return
}

// Computes the lexical distance between strings A and B.
//
// The "distance" between two strings is given by counting the minimum number
// of edits needed to transform string A into string B. An edit can be an
// insertion, deletion, or substitution of a single character, or a swap of two
// adjacent characters.
//
// Includes a custom alteration from Damerau-Levenshtein to treat case changes
// as a single edit which helps identify mis-cased values with an edit distance
// of 1.
//
// This distance can be useful for detecting typos in input or sorting
func lexicalDistance(a, b string) int {
	if a == b {
		return 0
	}

	a = strings.ToLower(a)
	b = strings.ToLower(b)

	// Any case change counts as a single edit
	if a == b {
		return 1
	}

	return levenshtein.ComputeDistance(a, b)
}
