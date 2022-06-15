//go:build windows

package sendinput

import (
	"errors"
	"unicode/utf16"
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

func SendTextInput(text string) error {
	if len(text) == 0 {
		return nil
	}

	chars := utf16.Encode([]rune(text))
	inputs := make([]win.KEYBD_INPUT, 0, len(chars))
	for _, char := range chars {
		inputs = append(inputs, win.KEYBD_INPUT{
			Type: win.INPUT_KEYBOARD,
			Ki: win.KEYBDINPUT{
				WScan:   char,
				DwFlags: win.KEYEVENTF_UNICODE,
			},
		})
	}

	// ret means the number of events success sent.
	ret := win.SendInput(uint32(len(inputs)), unsafe.Pointer(&inputs[0]), int32(unsafe.Sizeof(inputs[0])))
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

func SendMouseWhlInput(x, y int32) error {
	inputs := make([]win.MOUSE_INPUT, 0, 2)
	if y != 0 {
		inputs = append(inputs, win.MOUSE_INPUT{
			Type: win.INPUT_MOUSE,
			Mi: win.MOUSEINPUT{
				MouseData: uint32(y),
				DwFlags:   win.MOUSEEVENTF_WHEEL,
			},
		})
	}
	if x != 0 {
		inputs = append(inputs, win.MOUSE_INPUT{
			Type: win.INPUT_MOUSE,
			Mi: win.MOUSEINPUT{
				MouseData: uint32(x),
				DwFlags:   win.MOUSEEVENTF_HWHEEL,
			},
		})
	}

	// ret means the number of events success sent.
	ret := win.SendInput(uint32(len(inputs)), unsafe.Pointer(&inputs[0]), int32(unsafe.Sizeof(inputs[0])))
	if ret == 0 {
		return errors.New("SendInput failed")
	}
	return nil
}
