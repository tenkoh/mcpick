/*
mcpick:
Light weight Terminal User Interface (TUI) to pick material colors.

copyright:
tenkoh

license:
MIT

GitHub:
https://github.com/tenkoh/mcpick
*/

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
	"github.com/rivo/tview"
)

const manual = `Move up/down = (k)/(j) or arrow keys. Press Enter/Tab to switch color and light.`

var version = "v0.1.0"

var (
	app     *tview.Application
	clView  *RadioButtons
	liView  *RadioButtons
	output  *tview.List
	setting *tview.List
	footer  *tview.TextView
	stdout  string
)

// prevent locale problem, set rune width before launching app.
// with the problem, the first letter in each row would be lacked.
func init() {
	runewidth.DefaultCondition = &runewidth.Condition{EastAsianWidth: false}
}

func next(current, length int) int {
	current++
	if current >= length {
		current = 0
	}
	return current
}

func prev(current, length int) int {
	current--
	if current < 0 {
		current = length - 1
	}
	return current
}

func commonCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyRune:
		switch event.Rune() {
		case 'w':
			stdout = currentCode()
			app.Stop()
		case 'r':
			stdout = toTextRGB(currentCode())
			app.Stop()
		case 'b':
			current := clView.GetBackgroundColor()
			switchToDark := current == tcell.ColorWhite
			clView.dark(switchToDark)
			liView.dark(switchToDark)
		}
	}
	return event
}

func clCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEnter, tcell.KeyTAB:
		app.SetFocus(liView)
	}
	return commonCapture(event)
}

func clAfterDraw(screen tcell.Screen) {
	n := colorMajors[clView.currentOption].name
	c := colorMap[n]
	liView.options = c
	// to refresh
	liView.Draw(screen)
}

func liCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEnter, tcell.KeyTAB:
		// avoid nil pointer access
		liView.currentOption = 0
		app.SetFocus(clView)
	}
	return commonCapture(event)
}

func liAfterDraw(screen tcell.Screen) {
	w := currentCode()
	r := toTextRGB(w)
	output.SetItemText(0, "Web color:", w)
	output.SetItemText(1, "RGB color:", r)
	// need to refresh screen
	output.Draw(screen)
}

func currentCode() string {
	n := colorMajors[clView.currentOption].name
	c := colorMap[n]
	label := c[liView.currentOption]
	return label[1:8]
}

func main() {
	// versioning
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Printf("mcpick: version = %s\n", version)
		return
	}

	// main
	app = tview.NewApplication()

	colors := colorMap[colorMajors[0].name]
	liView = NewRadioButtons(colors)
	liView.SetBorder(true).
		SetTitle("light").
		SetInputCapture(liCapture)
	liView.SetAfterDraw(liAfterDraw)

	cmajors := []string{}
	for _, c := range colorMajors {
		cmajors = append(cmajors, c.String())
	}
	clView = NewRadioButtons(cmajors)
	clView.SetBorder(true).
		SetTitle("color").
		SetInputCapture(clCapture)
	clView.SetAfterDraw(clAfterDraw)

	output = tview.NewList().
		AddItem("Web color:", currentCode(), 'w', nil).
		AddItem("RGB color:", toTextRGB(currentCode()), 'r', nil)
	output.SetBorder(true).
		SetTitle("output operation")
	output.SetSelectedFocusOnly(true)
	output.SetSecondaryTextColor(tcell.ColorWhite)

	setting = tview.NewList().
		AddItem("Turn over background", "black <-> white", 'b', nil)
	setting.SetBorder(true).
		SetTitle("setting")
	setting.SetSelectedFocusOnly(true)
	setting.SetSecondaryTextColor(tcell.ColorWhite)

	footer = tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(manual)

	rflex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(output, 0, 1, false).
		AddItem(setting, 0, 1, false)

	flex := tview.NewFlex().
		AddItem(clView, 0, 1, true).
		AddItem(liView, 0, 1, false).
		AddItem(rflex, 0, 1, false)

	wrapper := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(flex, 0, 1, true).
		AddItem(footer, 1, 0, false)

	if err := app.SetRoot(wrapper, true).Run(); err != nil {
		panic(err)
	}

	if stdout != "" {
		fmt.Fprintf(os.Stdout, "%s", stdout)
	}
}
