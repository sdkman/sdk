package txt

import "github.com/fatih/color"

var yellow = color.New(color.FgHiYellow)
var Info = yellow.SprintFunc()
var InfoF = yellow.SprintfFunc()

var red = color.New(color.FgRed)
var Error = red.SprintFunc()
var ErrorF = red.SprintfFunc()

var green = color.New(color.FgGreen)
var Success = green.SprintFunc()
var SuccessF = green.SprintfFunc()

var BroadcastF = color.New(color.FgCyan).SprintfFunc()
