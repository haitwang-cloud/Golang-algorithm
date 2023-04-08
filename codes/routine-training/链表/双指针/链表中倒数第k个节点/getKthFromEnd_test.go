package main

import (
	"reflect"
	"testing"
)

func Test_getKthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
				k:    2,
			},
			want: &ListNode{4, &ListNode{5, nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEnd(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getKthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
