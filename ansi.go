package ansi

const (
	Esc   = "\x1b"
	Reset = Esc + "[0m" // reset; clears all colors and styles (to white on black)

	Bold = esc + "[1m" // bold on
	Italics = esc + "[3m" // italics on
	Underline = esc + "[4m" // underline on
	Inverse = esc + "[7m" // inverse on; reverses foreground & background colors
	Strike = esc + "[9m" // strikethrough on

	BoldOff      = esc + "[22m" // bold off
	ItalicsOff   = esc + "[23m" // italics off
	UnderlineOff = esc + "[24m" // underline off
	InverseOff   = esc + "[27m" // inverse off
	StrikeOff    = esc + "[29m" // strikethrough off

	Black   = esc + "[30m" // set foreground color to black
	Red     = esc + "[31m" // set foreground color to red
	Green   = esc + "[32m" // set foreground color to green
	Yellow  = esc + "[33m" // set foreground color to yellow
	Blue    = esc + "[34m" // set foreground color to blue
	Magenta = esc + "[35m" // set foreground color to magenta
	Cyan    = esc + "[36m" // set foreground color to cyan
	White   = esc + "[37m" // set foreground color to white
	Default = esc + "[39m" // set foreground color to default

	BgBlack   = esc + "[40m" // set background color to black
	BgRed     = esc + "[41m" // set background color to red
	BgGreen   = esc + "[42m" // set background color to green
	BgYellow  = esc + "[43m" // set background color to yellow
	BgBlue    = esc + "[44m" // set background color to blue
	BgMagenta = esc + "[45m" // set background color to magenta
	BgCyan    = esc + "[46m" // set background color to cyan
	BgWhite   = esc + "[47m" // set background color to white
	BgDefault = esc + "[49m" // set background color to default
)

