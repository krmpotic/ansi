package ansi

const (
	Reset = sgr(iota)
	Bold
	Faint
	Italic
	Underline
	Blink
	BlinkFast
	Invert
	Hide
	Strike

	Font0 = sgr(10 + iota)
	Font1
	Font3
	Font4
	Font5
	Font6
	Font7
	Font8
	Font9
	Font10

	Fraktur = sgr(20 + iota)
	DoubleUnderline
	NormalIntensity
	NoItalic
	NoUnderline
	NoBlink
	_
	NoInvert
	NoHide
	NoStrike

	// set foreground color
	Black = sgr(30 + iota)
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	fgColor // use Color8 / ColorRGB functions instead
	FgDefault

	// set background color
	BgBlack = sgr(40 + iota)
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	bgColor // use BgColor8 / BgColorRGB functions instead
	BgDefault

	_ = sgr(50 + iota)
	Frame
	Encircle
	Overline
	NoFrame
	NoOverline
	ulColor // underline color (not in standard)
	UlDefault

	IdeoRight = sgr(60 + iota)
	IdeoDoubleRight
	IdeoLeft
	IdeoDoubleLeft
	IdeoStress
	NoIdeo

	Superscript = sgr(73 + iota)
	Subscript
	NoScript

	FgBright0 = sgr(90 + iota)
	FgBright1
	FgBright2
	FgBright3
	FgBright4
	FgBright5
	FgBright6
	FgBright7

	BgBright0 = sgr(100 + iota)
	BgBright1
	BgBright2
	BgBright3
	BgBright4
	BgBright5
	BgBright6
	BgBright7

	sgrMask  = sgr(0b0111_1111)
	colorRGB = sgr(0b1000_0000)
)
