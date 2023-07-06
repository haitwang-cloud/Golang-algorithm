package main

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name  string
		args  args
		wants []int
	}{
		{
			name:  "test1",
			args:  args{nums: []int{2, 0, 2, 1, 1, 0}},
			wants: []int{0, 0, 1, 1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.nums)
		})
	}
}
