package main

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{0, 1, 0, 3, 12},
			},
			want: []int{1, 3, 12, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			for i := 0; i < len(tt.want); i++ {
				if tt.args.nums[i] != tt.want[i] {
					t.Errorf("moveZeroes() = %v, want %v", tt.args.nums, tt.want)
				}
			}
		})
	}
}
