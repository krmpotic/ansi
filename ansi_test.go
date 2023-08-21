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

func TestColor8(t *testing.T) {
	in := []uint8{
		3,
		100,
		0,
		255,
	}
	out := []string{
		"\033[38;5;3m",
		"\033[38;5;100m",
		"\033[38;5;0m",
		"\033[38;5;255m",
	}

	for i := range in {
		if str := Color8(in[i]).String(); str != out[i] {
			t.Fatalf("Color8(%d) = %q != %q\n", in[i], str, out[i])
		}
	}
}

func TestBgColor8(t *testing.T) {
	in := []uint8{
		3,
		113,
		0,
		255,
	}
	out := []string{
		"\033[48;5;3m",
		"\033[48;5;113m",
		"\033[48;5;0m",
		"\033[48;5;255m",
	}

	for i := range in {
		if str := BgColor8(in[i]).String(); str != out[i] {
			t.Fatalf("BgColor8(%d) = %q != %q\n", in[i], str, out[i])
		}
	}
}
func TestColorRGB(t *testing.T) {
	in := [][3]uint8{
		{3, 15, 200},
		{113, 200, 10},
		{0, 255, 13},
		{255, 13, 16},
	}
	out := []string{
		"\033[38;2;3;15;200m",
		"\033[38;2;113;200;10m",
		"\033[38;2;0;255;13m",
		"\033[38;2;255;13;16m",
	}

	for i := range in {
		r, g, b := in[i][0], in[i][1], in[i][2]
		if str := ColorRGB(r, g, b).String(); str != out[i] {
			t.Fatalf("ColorRGB(%d, %d, %d) = %q != %q\n", r, g, b, str, out[i])
		}
	}
}
func TestBgColorRGB(t *testing.T) {
	in := [][3]uint8{
		{3, 15, 200},
		{113, 200, 10},
		{0, 255, 13},
		{255, 13, 16},
	}
	out := []string{
		"\033[48;2;3;15;200m",
		"\033[48;2;113;200;10m",
		"\033[48;2;0;255;13m",
		"\033[48;2;255;13;16m",
	}

	for i := range in {
		r, g, b := in[i][0], in[i][1], in[i][2]
		if str := BgColorRGB(r, g, b).String(); str != out[i] {
			t.Fatalf("BgColorRGB(%d, %d, %d) = %q != %q\n", r, g, b, str, out[i])
		}
	}
}
func TestParameter(t *testing.T) {
	in := []Sgr{
		Magenta,
		Blink,
		ColorRGB(1, 2, 3),
		Color8(122),
	}
	out := []string{
		"35",
		"5",
		"38;2;1;2;3",
		"38;5;122",
	}
	for i := range in {
		if p := in[i].parameter(); p != out[i] {
			t.Fatalf("in[%d].parameter() = %q != %q", i, p, out[i])
		}
	}
}
func TestPaint(t *testing.T) {
	in := []struct {
		a Sgr
		b string
	}{
		{Magenta, "YoYo"},
	}
	out := []string{
		"\033[35mYoYo\033[0m",
	}

	for i := range in {
		if p := in[i].a.Paint(in[i].b); p != out[i] {
			t.Fatalf("in[%d].Paint(%s) = %q != %q", i, in[i].b, p, out[i])
		}
	}
}
func TestSgrStringer(t *testing.T) {
	Readline = true

	in := []Sgr{
		Magenta,
		Blink,
	}
	out := []string{
		"\001\033[35m\002",
		"\x01\x1b[5m\x02",
	}
	for i := range in {
		if p := in[i].String(); p != out[i] {
			t.Fatalf("in[%d].parameter() = %q != %q", i, p, out[i])
		}
	}
}

func TestStyleStringer(t *testing.T) {
	Readline = true

	in := []Style{
		{Magenta, Bold},
		{Blink, Green},
	}
	out := []string{
		"\001\033[35m\033[1m\002",
		"\x01\x1b[5m\x1b[32m\x02",
	}
	for i := range in {
		if p := in[i].String(); p != out[i] {
			t.Fatalf("in[%d].parameter() = %q != %q", i, p, out[i])
		}
	}
}
