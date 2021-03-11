package arrsandstrs

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	tests := []struct {
		name string
		arg  matrix
		want matrix
	}{
		{"zero", matrix{}, matrix{}},
		{"empty", matrix{row{}}, matrix{row{}}},

		{"1x1 no zeros", matrix{row{123}}, matrix{row{123}}},
		{"1x1 one zero", matrix{row{0}}, matrix{row{0}}},

		{"2x2 no zeros", matrix{{5, 6}, {7, 8}}, matrix{{5, 6}, {7, 8}}},
		{"2x2 one zero", matrix{{5, 6}, {7, 0}}, matrix{{0, 0}, {0, 0}}},

		{"3x3 no zeros", matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		{"3x3 one zero", matrix{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}}, matrix{{1, 0, 3}, {0, 0, 0}, {7, 0, 9}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroMatrix(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZeroMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
