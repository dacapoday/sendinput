//go:build windows

package sendinput

import "github.com/lxn/win"

type button uint32

const (
	buttonLEFTDOWN   button = win.MOUSEEVENTF_LEFTDOWN
	buttonLEFTUP     button = win.MOUSEEVENTF_LEFTUP
	buttonRIGHTDOWN  button = win.MOUSEEVENTF_RIGHTDOWN
	buttonRIGHTUP    button = win.MOUSEEVENTF_RIGHTUP
	buttonMIDDLEDOWN button = win.MOUSEEVENTF_MIDDLEDOWN
	buttonMIDDLEUP   button = win.MOUSEEVENTF_MIDDLEUP
)
