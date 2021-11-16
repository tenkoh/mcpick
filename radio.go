package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// RadioButtons implements a simple primitive for radio button selections.
type RadioButtons struct {
	*tview.Box
	options       []string
	currentOption int
	afterDraw     func(tcell.Screen)
	buttonColor   tcell.Color
}

// NewRadioButtons returns a new radio button primitive.
func NewRadioButtons(options []string) *RadioButtons {
	return &RadioButtons{
		Box:         tview.NewBox(),
		options:     options,
		buttonColor: tcell.ColorYellow,
	}
}

// SetOption sets options to this primitive
func (r *RadioButtons) SetOption(options []string) {
	r.options = options
}

func (r *RadioButtons) SetAfterDraw(f func(tcell.Screen)) {
	r.afterDraw = f
}

// Draw draws this primitive onto the screen.
func (r *RadioButtons) Draw(screen tcell.Screen) {
	r.Box.DrawForSubclass(screen, r)
	x, y, width, height := r.GetInnerRect()

	for index, option := range r.options {
		if index >= height {
			break
		}
		radioButton := " \u25ef" // Unchecked.
		if index == r.currentOption {
			radioButton = " \u25c9" // Checked.
		}
		line := fmt.Sprintf(`%s[white]  %s`, radioButton, option)
		tview.Print(screen, line, x, y+index, width, tview.AlignLeft, r.buttonColor)
	}

	if r.afterDraw != nil {
		r.afterDraw(screen)
	}
}

// InputHandler returns the handler for this primitive.
func (r *RadioButtons) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return r.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		switch event.Key() {
		case tcell.KeyUp:
			r.prev()
		case tcell.KeyDown:
			r.next()
		case tcell.KeyRune:
			switch event.Rune() {
			case 'j':
				r.next()
			case 'k':
				r.prev()
			}
		}
	})
}

func (r *RadioButtons) next() {
	r.currentOption = next(r.currentOption, len(r.options))
}

func (r *RadioButtons) prev() {
	r.currentOption = prev(r.currentOption, len(r.options))
}

func (r *RadioButtons) dark(b bool) {
	if b {
		r.SetBackgroundColor(tcell.ColorBlack)
		r.SetTitleColor(tcell.ColorWhite)
		r.SetBorderColor(tcell.ColorWhite)
		r.buttonColor = tcell.ColorYellow
	} else {
		r.SetBackgroundColor(tcell.ColorWhite)
		r.SetTitleColor(tcell.ColorBlack)
		r.SetBorderColor(tcell.ColorBlack)
		r.buttonColor = tcell.ColorBlue
	}
}
