package main

import "os"

// ANSI color constants for styling
var (
	reset  = "\x1b[0m"
	bold   = "\x1b[1m"
	gray   = "\x1b[90m"
	red    = "\x1b[31m"
	green  = "\x1b[32m"
	yellow = "\x1b[33m"
	blue   = "\x1b[34m"

	barChar = "🬋"
)

func init() {
	// Respect NO_COLOR convention (https://no-color.org/)
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		reset = ""
		bold = ""
		gray = ""
		red = ""
		green = ""
		yellow = ""
		blue = ""
	}
}
