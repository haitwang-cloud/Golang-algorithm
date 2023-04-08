package main

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	commonPart := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "test1",
			args: args{
				headA: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  1,
						Next: commonPart,
					},
				},
				headB: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val:  1,
							Next: commonPart,
						},
					},
				},
			},
			want: commonPart,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
