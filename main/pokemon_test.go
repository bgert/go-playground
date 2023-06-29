package main

//import (
//	"testing"
//)
//
//func Test_filterWeight(t *testing.T) {
//	type args struct {
//		weight int
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		{
//			name: "Less than 51 returns false",
//			args: args{weight: 50},
//			want: false,
//		},
//		{
//			name: "51 or greater returns true",
//			args: args{weight: 50},
//			want: true,
//		},
//	}
//		for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := FilterWeight(tt.args.weight); got != tt.want {
//				t.Errorf("filterWeight() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}