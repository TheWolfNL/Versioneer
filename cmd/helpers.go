package cmd

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func red(input string) string     { return color.RedString(input) }
func green(input string) string   { return color.GreenString(input) }
func yellow(input string) string  { return color.YellowString(input) }
func blue(input string) string    { return color.BlueString(input) }
func magenta(input string) string { return color.MagentaString(input) }
func cyan(input string) string    { return color.CyanString(input) }
func white(input string) string   { return color.WhiteString(input) }

var labelColor = color.FgBlue
var label = color.New(labelColor).Add(color.Bold)

func output(message string) {
	fmt.Printf("%v %v\n", label.Sprintf("[ %s ]", applicationName), message)
}

func status(title string, status string) {
	output(fmt.Sprintf("%v: %v", title, status))
}

func handleError(err error) {
	if err != nil {
		output(red(err.Error()))
	}
}

func handleFatal(err error) {
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
