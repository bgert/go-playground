package main

import (
	"reflect"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{ s: "barfoothefoobarman", words: []string{"foo","bar"}},
			want: []int{0,9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
