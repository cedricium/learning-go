package main

import (
	"fmt"
	"testing"
)

func TestAddInts(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{-2, 2, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v + %v", tt.a, tt.b), func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}

func TestAddFloats(t *testing.T) {
	tests := []struct {
		a    float64
		b    float64
		want float64
	}{
		{1.0, 1.5, 2.5},
		{2.3, 2.7, 5.0},
		{-2.3, -2.7, -5.0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v + %v", tt.a, tt.b), func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
