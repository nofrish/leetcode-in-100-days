package lintcode_391

import (
	"sort"
)

type Interval struct {
	Start, End int
}

func CountOfAirplanes(airplanes []*Interval) int {
	type Point struct {
		Time int
		Add  int
	}
	points := make([]Point, 0)
	for _, airplane := range airplanes {
		points = append(points, Point{airplane.Start, 1})
		points = append(points, Point{airplane.End, -1})
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].Time == points[j].Time {
			return points[i].Add < points[j].Add
		}
		return points[i].Time < points[j].Time
	})

	count, max := 0, 0
	for _, point := range points {
		count += point.Add
		if max < count {
			max = count
		}
	}
	return max
}
