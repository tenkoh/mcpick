package main

import (
	"reflect"
	"testing"
)

func Test_toRGB(t *testing.T) {
	type args struct {
		h string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"white", args{"FFFFFF"}, []int{255, 255, 255}},
		{"white", args{"ffffff"}, []int{255, 255, 255}},
		{"black", args{"000000"}, []int{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toRGB(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toRGB() = %v, want %v", got, tt.want)
			}
		})
	}
}
