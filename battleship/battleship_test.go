package main

import (
	"testing"
)

func TestCoordinates(t *testing.T) {
	tests := []struct {
		shot     string
		row, col int8
		errMsg   string
	}{
		{"A:0", 0, 0, ""},
		{"J:9", 9, 9, ""},
		{"D:3", 3, 3, ""},
		{"b:7", 1, 7, ""},
		{"C:40", -1, -1, "invalid shot length"},
		{"C-4", -1, -1, "invalid shot format"},
		{":12", -1, -1, "invalid shot format"},
		{"AA:", -1, -1, "invalid shot format"},
		{"Z:4", -1, -1, "shot row out of range"},
		{"C:Z", -1, -1, "strconv.Atoi: parsing \"Z\": invalid syntax"},
	}

	for _, tt := range tests {
		t.Run(tt.shot, func(t *testing.T) {
			r, c, err := coordinates(tt.shot)
			if r != tt.row || c != tt.col {
				t.Errorf("Expected (%d, %d), got (%d, %d)", tt.row, tt.col, r, c)
			}

			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != tt.errMsg {
				t.Errorf("Expected error message `%s`, got `%s`", tt.errMsg, errMsg)
			}
		})
	}
}
