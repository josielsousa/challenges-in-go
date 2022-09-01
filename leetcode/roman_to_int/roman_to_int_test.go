package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_romanToInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		s    string
		want int
	}{
		{s: "MCMXCIV", want: 1994},
		{s: "IV", want: 4},
		{s: "IX", want: 9},
		{s: "LVIII", want: 58},
		{s: "XXVII", want: 27},
	}
	for _, tt := range tests {
		tt := tt // capture range variable

		t.Run(
			fmt.Sprintf("should convert roman %s to int %d", tt.s, tt.want),
			func(t *testing.T) {
				t.Parallel()
				got := romanToInt(tt.s)
				assert.Equal(t, tt.want, got)
			})
	}
}
