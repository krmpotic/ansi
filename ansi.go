package ansi

const (
	Esc   = "\x1b"
	Reset = Esc + "[0m" // reset; clears all colors and styles (to white on black)

	Bold = Esc + "[1m" // bold on
	Italics = Esc + "[3m" // italics on
	Underline = Esc + "[4m" // underline on
	Inverse = Esc + "[7m" // inverse on; reverses foreground & background colors
	Strike = Esc + "[9m" // strikethrough on

	BoldOff      = Esc + "[22m" // bold off
	ItalicsOff   = Esc + "[23m" // italics off
	UnderlineOff = Esc + "[24m" // underline off
	InverseOff   = Esc + "[27m" // inverse off
	StrikeOff    = Esc + "[29m" // strikethrough off

	Black   = Esc + "[30m" // set foreground color to black
	Red     = Esc + "[31m" // set foreground color to red
	Green   = Esc + "[32m" // set foreground color to green
	Yellow  = Esc + "[33m" // set foreground color to yellow
	Blue    = Esc + "[34m" // set foreground color to blue
	Magenta = Esc + "[35m" // set foreground color to magenta
	Cyan    = Esc + "[36m" // set foreground color to cyan
	White   = Esc + "[37m" // set foreground color to white
	Default = Esc + "[39m" // set foreground color to default

	BgBlack   = Esc + "[40m" // set background color to black
	BgRed     = Esc + "[41m" // set background color to red
	BgGreen   = Esc + "[42m" // set background color to green
	BgYellow  = Esc + "[43m" // set background color to yellow
	BgBlue    = Esc + "[44m" // set background color to blue
	BgMagenta = Esc + "[45m" // set background color to magenta
	BgCyan    = Esc + "[46m" // set background color to cyan
	BgWhite   = Esc + "[47m" // set background color to white
	BgDefault = Esc + "[49m" // set background color to default
)

