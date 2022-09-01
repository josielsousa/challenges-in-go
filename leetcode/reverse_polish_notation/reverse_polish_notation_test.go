package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evalRPN(t *testing.T) {
	t.Parallel()

	tests := []struct {
		tokens []string
		want   int
	}{
		{
			tokens: []string{"2", "1", "+", "3", "*", "2", "-"},
			want:   7,
		},
		{
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
		{
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable

		t.Run(fmt.Sprintf("should evaluate the value %d", tt.want), func(t *testing.T) {
			t.Parallel()
			got := evalRPN(tt.tokens)
			assert.Equal(t, tt.want, got)
		})
	}
}
