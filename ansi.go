package ansi

import "strconv"

// en.wikipedia.org/wiki/ANSI_escape_code
// man 4 console_codes

// superuser.com/a/301355
var Readline bool

const (
	readlineStart = "\x01"
	readlineStop  = "\x02"
)

const (
	CSI = "\033[" // control sequence introducer
)

// SGR (Select Graphic Rendition) parameters
// control sequence CSI n m
type Sgr uint32

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

/*
func Sgr(i int) (sgr, bool) {
	if i < 0 || i >= 128 || sgr(i) == fgColor || sgr(i) == bgColor || sgr(i) == ulColor {
		return 0, false
	}
	return sgr(i), true
}
*/

func (s Sgr) String() string {
	return Style{s}.String()
}

func (s Sgr) Paint(str string) string {
	return s.String() + str + Reset.String()
}

func Color8(i uint8) Sgr {
	return fgColor | Sgr(uint32(i)<<8)
}

func BgColor8(i uint8) Sgr {
	return bgColor | Sgr(uint32(i)<<8)
}

func UlColor8(i uint8) Sgr {
	return ulColor | Sgr(uint32(i)<<8)
}

func ColorRGB(r, g, b uint8) Sgr {
	return fgColor | colorRGB | Sgr(uint32(r)<<8|uint32(g)<<16|uint32(b)<<24)
}

func BgColorRGB(r, g, b uint8) Sgr {
	return bgColor | colorRGB | Sgr(uint32(r)<<8|uint32(g)<<16|uint32(b)<<24)
}

func UlColorRGB(r, g, b uint8) Sgr {
	return ulColor | colorRGB | Sgr(uint32(r)<<8|uint32(g)<<16|uint32(b)<<24)
}

type Style []Sgr

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

func (s Style) Paint(str string) string {
	return s.String() + str + Reset.String()
}
