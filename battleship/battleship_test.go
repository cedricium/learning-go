package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCoordinates(t *testing.T) {
	tests := []struct {
		shot   string
		coords *Coordinate
		errMsg string
	}{
		{"A0", &Coordinate{0, 0}, ""},
		{"J9", &Coordinate{9, 9}, ""},
		{"D3", &Coordinate{3, 3}, ""},
		{"b7", &Coordinate{1, 7}, ""},
		{"C40", nil, errShotFormat},
		{"C-4", nil, errShotFormat},
		{"12", nil, errShotRow},
		{"AA", nil, errShotCol},
		{"💤4", nil, errShotFormat},
		{"A💯", nil, errShotFormat},
	}

	for _, tt := range tests {
		t.Run(tt.shot, func(t *testing.T) {
			c, err := coordinates(tt.shot)

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != tt.errMsg {
				t.Errorf("Expected error message `%s`, got `%s`", tt.errMsg, errMsg)
			}

			if !cmp.Equal(c, tt.coords) {
				t.Errorf("Expected %v, got %v", tt.coords, c)
			}
		})
	}
}
