package arrsandstrs

import (
	"reflect"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	tests := []struct {
		name string
		arg  matrix
		want matrix
	}{
		{"zero", matrix{}, matrix{}},
		{"empty", matrix{row{}}, matrix{row{}}},
		{"1x1", matrix{row{123}}, matrix{row{123}}},

		{"2x2", matrix{{12, 34}, {56, 78}}, matrix{{56, 12}, {78, 34}}},

		{"3x3", matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, matrix{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RotateMatrix(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RotateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
