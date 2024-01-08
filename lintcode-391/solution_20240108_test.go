package lintcode_391

import "testing"

func TestSolution(t *testing.T) {
	result := CountOfAirplanes([]*Interval{
		{10, 14},
		{10, 15},
		{10, 16},
		{1, 10},
		{2, 10},
		{3, 10},
		{4, 10}})
	t.Error(result)
}
