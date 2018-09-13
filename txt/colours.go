package txt

import "github.com/fatih/color"

var Info = color.New(color.FgHiYellow).SprintfFunc()

var Error = color.New(color.FgRed).SprintfFunc()

var Success = color.New(color.FgGreen).SprintfFunc()

var Broadcast = color.New(color.FgCyan).SprintfFunc()
