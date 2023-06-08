package main

import (
	"fmt"
	"math"
	"sort"
)

type meeting struct {
	startTime int
	endTime   int
}

func hiCal(meetings []meeting) []meeting {
	if len(meetings) == 0 {
		return meetings
	}

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i].startTime < meetings[j].startTime
	})

	for i := 1; i < len(meetings); i++ {
		x, y := meetings[i-1], meetings[i]

		if y.startTime <= x.endTime {
			x = mergeMeetings(x, y)
			meetings = shiftLeft(meetings, i)
		}
	}

	return meetings
}

func mergeMeetings(x meeting, y meeting) meeting {
	start := math.Min(float64(x.startTime), float64(y.startTime))
	end := math.Max(float64(x.endTime), float64(y.endTime))

	return meeting{
		startTime: int(start),
		endTime:   int(end),
	}
}

func shiftLeft(meetings []meeting, i int) []meeting {
	last := len(meetings) - 1
	for j := last; j > i; j-- {
		meetings[j-1] = meetings[j]
	}
	return meetings[:last]
}

func main() {
	meetings := []meeting{
		{startTime: 1, endTime: 10},
		{startTime: 2, endTime: 6},
		{startTime: 3, endTime: 5},
		{startTime: 7, endTime: 9},

		/*
		{startTime: 0, endTime: 1},
		{startTime: 3, endTime: 5},
		{startTime: 4, endTime: 8},
		{startTime: 10, endTime: 12},
		{startTime: 9, endTime: 10},
		*/
	}

	fmt.Println(hiCal(meetings))
}
