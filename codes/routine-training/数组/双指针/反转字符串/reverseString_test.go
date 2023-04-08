package main

import "testing"

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test1",
			args: args{
				s: []byte{'h', 'e', 'l', 'l', 'o'},
			},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "test2",
			args: args{
				s: []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			for i := 0; i < len(tt.want); i++ {
				if tt.args.s[i] != tt.want[i] {
					t.Errorf("reverseString() = %v, want %v", tt.args.s, tt.want)
				}
			}
		})
	}
}
