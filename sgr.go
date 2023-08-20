package ansi

const (
	Reset = Sgr(iota)
	Bold
	Faint
	Italic
	Underline
	Blink
	BlinkFast
	Invert
	Hide
	Strike

	Font0
	Font1
	Font3
	Font4
	Font5
	Font6
	Font7
	Font8
	Font9
	Font10

	Fraktur
	DoubleUnderline
	NormalIntensity
	NoItalic
	NoUnderline
	NoBlink
	_
	NoInvert
	NoHide
	NoStrike

	Black
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	fgColor
	FgDefault

	BgBlack
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	bgColor
	BgDefault
)

const (
	_ = Sgr(50 + iota)
	Frame
	Encircle
	Overline
	NoFrame
	NoOverline
	ulColor // underline color (not in standard)
	UlDefault
)

const (
	IdeoRight = Sgr(60 + iota)
	IdeoDoubleRight
	IdeoLeft
	IdeoDoubleLeft
	IdeoStress
	NoIdeo
)

const (
	Superscript = Sgr(73 + iota)
	Subscript
	NoScript
)

const (
	FgBright0 = Sgr(90 + iota)
	FgBright1
	FgBright2
	FgBright3
	FgBright4
	FgBright5
	FgBright6
	FgBright7
)

const (
	BgBright0 = Sgr(100 + iota)
	BgBright1
	BgBright2
	BgBright3
	BgBright4
	BgBright5
	BgBright6
	BgBright7
)

const (
	sgrMask  = Sgr(0b0111_1111)
	colorRGB = Sgr(0b1000_0000)
)
