//go:build windows

package sendinput

import (
	"errors"
	"unsafe"

	"github.com/lxn/win"
)

func sendBtnEvent(btn button) error {
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

func sendRelEvent(x, y int32) error {
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

func sendwhlEvent(horizontal bool, delta int32) error {
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
