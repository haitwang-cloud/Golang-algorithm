package main

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				head: &ListNode{1, &ListNode{4, &ListNode{2, nil}}},
				x:    3,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		},
		{
			name: "test2",
			args: args{
				head: &ListNode{2, &ListNode{4, &ListNode{1, nil}}},
				x:    2,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
