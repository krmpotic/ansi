package ansi

import (
	"testing"
)

func TestConstants(t *testing.T) {
	input := []Sgr{
		Red, Yellow, Encircle,
		BgGreen,
	}
	out := []int{
		31, 33, 52,
		42,
	}

	for i := range input {
		if int(input[i]) != out[i] {
			t.Fatalf("sgr(%d) != %d\n", input[i], out[i])
		}
	}
}
