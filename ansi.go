package ansi

import "strconv"

// Package ansi simplifies work with ANSI escape codes (specifically SGR
// sequences), used to control the display of terminals.

// see:
// http://en.wikipedia.org/wiki/ANSI_escape_code
// man 4 console_codes

// set Readline to true, to provide readline delimiters of non-printable
// sequences, which will enable it to calculate character and cursor positions
// properly. See superuser.com/a/301355
var Readline bool

const (
	readlineStart = "\x01"
	readlineStop  = "\x02"
)

const (
	CSI = "\033[" // Control Sequence Introducer
)

// Sgr type encodes an SGR (Select Graphic Rendition) sequence, which
// sets display attributes of terminal emulators.
//   - SGR sequence CSI + n + "m", is encoded as n in the low 7 bits of Sgr
//   - SGR sequence for 8-bit color CSI + n + ;5 + i + "m" is encoded in
//     the low two bytes of Sgr. Where:
//   - bits 0-6 encode n
//   - bit 7 (0-indexed) is unset
//   - bits 8-15 (0-indexed) encode the 8-bit color
//   - SGR sequence for 24-bit RGB color CSI + n + ;2r;g;b + "m" is encoded in
//     the four bytes of s. Where:
//   - bits 0-6 encode n
//   - bit 7 (0-indexed) is set
//   - bits 8-15 (0-indexed) hold Red value
//   - bits 16-23 (0-indexed) hold Green value
//   - bits 24-31 (0-indexed) hold Blue value
type Sgr uint32

const (
	sgrMask  = Sgr(0b0111_1111)
	colorRGB = Sgr(0b1000_0000)
)

// parameter decodes SGR parameter encoded in s, and returns it as a string.
// - non color parameters are decoded trivially: 5 -> "5"
// - 8-bit colors are decoded by:
// - 24-bit colors are decoded by:
func (s Sgr) parameter() string {
	n := s & sgrMask
	str := strconv.Itoa(int(n))

	// normal, easy parameter
	if n != fgColor && n != bgColor && n != ulColor {
		return str
	}

	if s&colorRGB != 0 { // RGB
		r := int((s & (0xff << 8)) >> 8)
		g := int((s & (0xff << 16)) >> 16)
		b := int((s & (0xff << 24)) >> 24)

		str += ";2;"
		str += strconv.Itoa(r) + ";"
		str += strconv.Itoa(g) + ";"
		str += strconv.Itoa(b)
	} else { // 8-bit color
		c := int(0xffff & (s >> 8))

		str += ";5;"
		str += strconv.Itoa(c)
	}
	return str
}

// String decodes s, and returns the corresponding ANSI-escape string. If
// Readline is set, it prepends and appends readline delimiters.
func (s Sgr) String() (str string) {
	if Readline {
		str += readlineStart
	}
	str += CSI + s.parameter() + "m"
	if Readline {
		str += readlineStop
	}
	return str
}

// Paint returns a string, stylized by the SGR parameter encoded in s.
// It prepends str with ANSI-escape sequence encoded in s, and appends
// ANSI Reset. Effectively painting str with display attribute in s.
func (s Sgr) Paint(str string) string {
	return s.String() + str + Reset.String()
}

// Color8 returns an Sgr parameter that encodes an 8-bit foreground color i.
func Color8(i uint8) Sgr {
	return fgColor | Sgr(uint32(i)<<8)
}

// BgColor8 returns an Sgr parameter that encodes an 8-bit background color i.
func BgColor8(i uint8) Sgr {
	return bgColor | Sgr(uint32(i)<<8)
}

// ColorRGB returns an Sgr parameter that encodes a RGB foreground color.
func ColorRGB(r, g, b uint8) Sgr {
	return fgColor | colorRGB | Sgr(uint32(r)<<8|uint32(g)<<16|uint32(b)<<24)
}

// BgColorRGB returns an Sgr parameter that encodes a RGB background color.
func BgColorRGB(r, g, b uint8) Sgr {
	return bgColor | colorRGB | Sgr(uint32(r)<<8|uint32(g)<<16|uint32(b)<<24)
}

// type Style is set of different SGR-parameters
type Style []Sgr

// String() constructs a string by chaining together
// SGR-sequences (of type CSI + ? + "m") of each
// element of s. If Readline is set, it prepends and appends
// to the chain readline's delimiters.
//
// While specification allows for chaing of parameters after a single CSI:
//
//	CSI + n0 + ";" + n1 + ";" + n2 + "m",
//
// This is not used, rather the chain is consturcted as so:
//
//	CSI + n0 + "m" + CSI + n1 + "m" + CSI + n2 + "m"
func (s Style) String() string {
	str := ""
	if Readline {
		str += readlineStart
	}

	for i := range s {
		str += CSI + s[i].parameter() + "m"
	}

	if Readline {
		str += readlineStop
	}
	return str
}

// Paint returns a string, stylized by all SGR parameters in slice s.
// It prepends str with the chain of SGR-sequences in slice s, and appends
// ANSI Reset to the end of str. In effect painting the string str with
// display attributes in s, and reseting them at the end.
func (s Style) Paint(str string) string {
	return s.String() + str + Reset.String()
}
