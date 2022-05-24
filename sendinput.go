//go:build windows

package sendinput

import (
	"errors"
	"unsafe"

	"github.com/lxn/win"
)

func SendKeyboardInput(code KeyCode, press bool) error {
	var direction uint32
	if !press {
		direction = 2
	}

	i := win.KEYBD_INPUT{
		Type: win.INPUT_KEYBOARD,
		Ki: win.KEYBDINPUT{
			WVk:     uint16(code),
			DwFlags: direction,
		},
	}

	// ret means the number of events success sent.
	ret := win.SendInput(1, unsafe.Pointer(&i), int32(unsafe.Sizeof(i)))
	if ret == 0 {
		return errors.New("SendInput failed")
	}
	return nil
}

func SendMouseBtnInput(btn MouseBtn) error {
	i := win.MOUSE_INPUT{
		Type: win.INPUT_MOUSE,
		Mi: win.MOUSEINPUT{
			DwFlags: uint32(btn),
		},
	}

	// ret means the number of events success sent.
	ret := win.SendInput(1, unsafe.Pointer(&i), int32(unsafe.Sizeof(i)))
	if ret == 0 {
		return errors.New("SendInput failed")
	}
	return nil
}

func SendMouseRelInput(x, y int32) error {
	i := win.MOUSE_INPUT{
		Type: win.INPUT_MOUSE,
		Mi: win.MOUSEINPUT{
			Dx:      x,
			Dy:      y,
			DwFlags: win.MOUSEEVENTF_MOVE,
		},
	}

	// ret means the number of events success sent.
	ret := win.SendInput(1, unsafe.Pointer(&i), int32(unsafe.Sizeof(i)))
	if ret == 0 {
		return errors.New("SendInput failed")
	}
	return nil
}

func SendMouseWhlInput(delta int32, horizontal bool) error {
	var flags uint32
	if horizontal {
		flags = win.MOUSEEVENTF_HWHEEL
	} else {
		flags = win.MOUSEEVENTF_WHEEL
	}

	i := win.MOUSE_INPUT{
		Type: win.INPUT_MOUSE,
		Mi: win.MOUSEINPUT{
			MouseData: uint32(delta),
			DwFlags:   flags,
		},
	}

	// ret means the number of events success sent.
	ret := win.SendInput(1, unsafe.Pointer(&i), int32(unsafe.Sizeof(i)))
	if ret == 0 {
		return errors.New("SendInput failed")
	}
	return nil
}
