package main

import (
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		list1 *ListNode
		list2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				list1: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
				list2: &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			},
			want: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
		{
			name: "test2",
			args: args{
				list1: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
				list2: nil,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
