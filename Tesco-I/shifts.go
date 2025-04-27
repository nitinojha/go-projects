package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
)

type Shift struct {
	Start int
	End   int
}

func main() {
	// intialise shift inputs
	shifts := []Shift{
		{Start: 8, End: 9},
		{Start: 9, End: 10},
		{Start: 12, End: 14},
		{Start: 14, End: 19},
	}

	mergedShifts := mergeShifts(shifts)
	fmt.Println(mergedShifts)
}

func mergeShifts(shifts []Shift) []Shift {
	n := len(shifts)
	if n == 0 {
		return []Shift{}
	}

	// Sort Shifts By start time
	sort.Slice(shifts,
		func(i, j int) bool {
			return shifts[i].Start < shifts[j].Start
		})

	merged := []Shift{shifts[0]}

	for _, shift := range shifts[1:n] {
		merged_last := &merged[len(merged)-1]

		if shift.Start <= merged_last.End { // // To be checked
			if shift.End > merged_last.End {
				merged_last.End = shift.End
			}
		} else {
			merged = append(merged, shift)
		}
	}
	return merged
}

func TestMergeShifts(t *testing.T) {
	tests := []struct {
		inputTest      []Shift
		expectedOutput []Shift
	}{
		{
			inputTest: []Shift{
				{Start: 8, End: 10},
				{Start: 10, End: 11},
				{Start: 12, End: 14},
				{Start: 14, End: 19},
			},
			expectedOutput: []Shift{
				{Start: 8, End: 11},
				{Start: 12, End: 19},
			},
		},
		{
			inputTest: []Shift{
				{Start: 9, End: 10},
				{Start: 8, End: 9},
				{Start: 12, End: 14},
				{Start: 14, End: 19},
			},
			expectedOutput: []Shift{
				{Start: 8, End: 10},
				{Start: 12, End: 19},
			},
		},
		{
			inputTest: []Shift{
				{Start: 8, End: 10},
				{Start: 10, End: 11},
				{Start: 12, End: 14},
				{Start: 15, End: 19},
			},
			expectedOutput: []Shift{
				{Start: 8, End: 11},
				{Start: 12, End: 14},
				{Start: 15, End: 19},
			},
		},
	}

	for _, test := range tests {
		mergedShift := mergeShifts(test.inputTest)
		inputResult, _ := json.Marshal(mergedShift)
		outputResult, _ := json.Marshal(test.expectedOutput)

		if string(inputResult) != string(outputResult) {
			t.Errorf("for testInput %v, epxected inputResult %v but got outputResult %v", test.inputTest, inputResult, outputResult)
		} else {
			t.Log("for testInput %v, epxected inputResult %v and we get equal outputResult %v", test.inputTest, inputResult, outputResult)
		}
	}

	//mergedShift json string

}
