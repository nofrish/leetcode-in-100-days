package _229

import (
	"reflect"
	"testing"
)

func Test_minAvailableDuration(t *testing.T) {
	type args struct {
		slots1   [][]int
		slots2   [][]int
		duration int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAvailableDuration(tt.args.slots1, tt.args.slots2, tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minAvailableDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
