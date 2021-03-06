//go:build windows

package sendinput

import (
	"strings"

	"github.com/lxn/win"
)

// KeyCode https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes
type KeyCode uint16

const (
	KEY_LBUTTON             KeyCode = win.VK_LBUTTON
	KEY_RBUTTON             KeyCode = win.VK_RBUTTON
	KEY_CANCEL              KeyCode = win.VK_CANCEL
	KEY_MBUTTON             KeyCode = win.VK_MBUTTON
	KEY_XBUTTON1            KeyCode = win.VK_XBUTTON1
	KEY_XBUTTON2            KeyCode = win.VK_XBUTTON2
	KEY_BACK                KeyCode = win.VK_BACK
	KEY_TAB                 KeyCode = win.VK_TAB
	KEY_CLEAR               KeyCode = win.VK_CLEAR
	KEY_RETURN              KeyCode = win.VK_RETURN
	KEY_SHIFT               KeyCode = win.VK_SHIFT
	KEY_CONTROL             KeyCode = win.VK_CONTROL
	KEY_MENU                KeyCode = win.VK_MENU
	KEY_PAUSE               KeyCode = win.VK_PAUSE
	KEY_CAPITAL             KeyCode = win.VK_CAPITAL
	KEY_KANA                KeyCode = win.VK_KANA
	KEY_HANGEUL             KeyCode = win.VK_HANGEUL
	KEY_HANGUL              KeyCode = win.VK_HANGUL
	KEY_JUNJA               KeyCode = win.VK_JUNJA
	KEY_FINAL               KeyCode = win.VK_FINAL
	KEY_HANJA               KeyCode = win.VK_HANJA
	KEY_KANJI               KeyCode = win.VK_KANJI
	KEY_ESCAPE              KeyCode = win.VK_ESCAPE
	KEY_CONVERT             KeyCode = win.VK_CONVERT
	KEY_NONCONVERT          KeyCode = win.VK_NONCONVERT
	KEY_ACCEPT              KeyCode = win.VK_ACCEPT
	KEY_MODECHANGE          KeyCode = win.VK_MODECHANGE
	KEY_SPACE               KeyCode = win.VK_SPACE
	KEY_PRIOR               KeyCode = win.VK_PRIOR
	KEY_NEXT                KeyCode = win.VK_NEXT
	KEY_END                 KeyCode = win.VK_END
	KEY_HOME                KeyCode = win.VK_HOME
	KEY_LEFT                KeyCode = win.VK_LEFT
	KEY_UP                  KeyCode = win.VK_UP
	KEY_RIGHT               KeyCode = win.VK_RIGHT
	KEY_DOWN                KeyCode = win.VK_DOWN
	KEY_SELECT              KeyCode = win.VK_SELECT
	KEY_PRINT               KeyCode = win.VK_PRINT
	KEY_EXECUTE             KeyCode = win.VK_EXECUTE
	KEY_SNAPSHOT            KeyCode = win.VK_SNAPSHOT
	KEY_INSERT              KeyCode = win.VK_INSERT
	KEY_DELETE              KeyCode = win.VK_DELETE
	KEY_HELP                KeyCode = win.VK_HELP
	KEY_LWIN                KeyCode = win.VK_LWIN
	KEY_RWIN                KeyCode = win.VK_RWIN
	KEY_APPS                KeyCode = win.VK_APPS
	KEY_SLEEP               KeyCode = win.VK_SLEEP
	KEY_NUMPAD0             KeyCode = win.VK_NUMPAD0
	KEY_NUMPAD1             KeyCode = win.VK_NUMPAD1
	KEY_NUMPAD2             KeyCode = win.VK_NUMPAD2
	KEY_NUMPAD3             KeyCode = win.VK_NUMPAD3
	KEY_NUMPAD4             KeyCode = win.VK_NUMPAD4
	KEY_NUMPAD5             KeyCode = win.VK_NUMPAD5
	KEY_NUMPAD6             KeyCode = win.VK_NUMPAD6
	KEY_NUMPAD7             KeyCode = win.VK_NUMPAD7
	KEY_NUMPAD8             KeyCode = win.VK_NUMPAD8
	KEY_NUMPAD9             KeyCode = win.VK_NUMPAD9
	KEY_MULTIPLY            KeyCode = win.VK_MULTIPLY
	KEY_ADD                 KeyCode = win.VK_ADD
	KEY_SEPARATOR           KeyCode = win.VK_SEPARATOR
	KEY_SUBTRACT            KeyCode = win.VK_SUBTRACT
	KEY_DECIMAL             KeyCode = win.VK_DECIMAL
	KEY_DIVIDE              KeyCode = win.VK_DIVIDE
	KEY_F1                  KeyCode = win.VK_F1
	KEY_F2                  KeyCode = win.VK_F2
	KEY_F3                  KeyCode = win.VK_F3
	KEY_F4                  KeyCode = win.VK_F4
	KEY_F5                  KeyCode = win.VK_F5
	KEY_F6                  KeyCode = win.VK_F6
	KEY_F7                  KeyCode = win.VK_F7
	KEY_F8                  KeyCode = win.VK_F8
	KEY_F9                  KeyCode = win.VK_F9
	KEY_F10                 KeyCode = win.VK_F10
	KEY_F11                 KeyCode = win.VK_F11
	KEY_F12                 KeyCode = win.VK_F12
	KEY_F13                 KeyCode = win.VK_F13
	KEY_F14                 KeyCode = win.VK_F14
	KEY_F15                 KeyCode = win.VK_F15
	KEY_F16                 KeyCode = win.VK_F16
	KEY_F17                 KeyCode = win.VK_F17
	KEY_F18                 KeyCode = win.VK_F18
	KEY_F19                 KeyCode = win.VK_F19
	KEY_F20                 KeyCode = win.VK_F20
	KEY_F21                 KeyCode = win.VK_F21
	KEY_F22                 KeyCode = win.VK_F22
	KEY_F23                 KeyCode = win.VK_F23
	KEY_F24                 KeyCode = win.VK_F24
	KEY_NUMLOCK             KeyCode = win.VK_NUMLOCK
	KEY_SCROLL              KeyCode = win.VK_SCROLL
	KEY_LSHIFT              KeyCode = win.VK_LSHIFT
	KEY_RSHIFT              KeyCode = win.VK_RSHIFT
	KEY_LCONTROL            KeyCode = win.VK_LCONTROL
	KEY_RCONTROL            KeyCode = win.VK_RCONTROL
	KEY_LMENU               KeyCode = win.VK_LMENU
	KEY_RMENU               KeyCode = win.VK_RMENU
	KEY_BROWSER_BACK        KeyCode = win.VK_BROWSER_BACK
	KEY_BROWSER_FORWARD     KeyCode = win.VK_BROWSER_FORWARD
	KEY_BROWSER_REFRESH     KeyCode = win.VK_BROWSER_REFRESH
	KEY_BROWSER_STOP        KeyCode = win.VK_BROWSER_STOP
	KEY_BROWSER_SEARCH      KeyCode = win.VK_BROWSER_SEARCH
	KEY_BROWSER_FAVORITES   KeyCode = win.VK_BROWSER_FAVORITES
	KEY_BROWSER_HOME        KeyCode = win.VK_BROWSER_HOME
	KEY_VOLUME_MUTE         KeyCode = win.VK_VOLUME_MUTE
	KEY_VOLUME_DOWN         KeyCode = win.VK_VOLUME_DOWN
	KEY_VOLUME_UP           KeyCode = win.VK_VOLUME_UP
	KEY_MEDIA_NEXT_TRACK    KeyCode = win.VK_MEDIA_NEXT_TRACK
	KEY_MEDIA_PREV_TRACK    KeyCode = win.VK_MEDIA_PREV_TRACK
	KEY_MEDIA_STOP          KeyCode = win.VK_MEDIA_STOP
	KEY_MEDIA_PLAY_PAUSE    KeyCode = win.VK_MEDIA_PLAY_PAUSE
	KEY_LAUNCH_MAIL         KeyCode = win.VK_LAUNCH_MAIL
	KEY_LAUNCH_MEDIA_SELECT KeyCode = win.VK_LAUNCH_MEDIA_SELECT
	KEY_LAUNCH_APP1         KeyCode = win.VK_LAUNCH_APP1
	KEY_LAUNCH_APP2         KeyCode = win.VK_LAUNCH_APP2
	KEY_OEM_1               KeyCode = win.VK_OEM_1
	KEY_OEM_PLUS            KeyCode = win.VK_OEM_PLUS
	KEY_OEM_COMMA           KeyCode = win.VK_OEM_COMMA
	KEY_OEM_MINUS           KeyCode = win.VK_OEM_MINUS
	KEY_OEM_PERIOD          KeyCode = win.VK_OEM_PERIOD
	KEY_OEM_2               KeyCode = win.VK_OEM_2
	KEY_OEM_3               KeyCode = win.VK_OEM_3
	KEY_OEM_4               KeyCode = win.VK_OEM_4
	KEY_OEM_5               KeyCode = win.VK_OEM_5
	KEY_OEM_6               KeyCode = win.VK_OEM_6
	KEY_OEM_7               KeyCode = win.VK_OEM_7
	KEY_OEM_8               KeyCode = win.VK_OEM_8
	KEY_OEM_102             KeyCode = win.VK_OEM_102
	KEY_PROCESSKEY          KeyCode = win.VK_PROCESSKEY
	KEY_PACKET              KeyCode = win.VK_PACKET
	KEY_ATTN                KeyCode = win.VK_ATTN
	KEY_CRSEL               KeyCode = win.VK_CRSEL
	KEY_EXSEL               KeyCode = win.VK_EXSEL
	KEY_EREOF               KeyCode = win.VK_EREOF
	KEY_PLAY                KeyCode = win.VK_PLAY
	KEY_ZOOM                KeyCode = win.VK_ZOOM
	KEY_NONAME              KeyCode = win.VK_NONAME
	KEY_PA1                 KeyCode = win.VK_PA1
	KEY_OEM_CLEAR           KeyCode = win.VK_OEM_CLEAR

	KEY_0 KeyCode = 0x30
	KEY_1 KeyCode = 0x31
	KEY_2 KeyCode = 0x32
	KEY_3 KeyCode = 0x33
	KEY_4 KeyCode = 0x34
	KEY_5 KeyCode = 0x35
	KEY_6 KeyCode = 0x36
	KEY_7 KeyCode = 0x37
	KEY_8 KeyCode = 0x38
	KEY_9 KeyCode = 0x39
	KEY_A KeyCode = 0x41
	KEY_B KeyCode = 0x42
	KEY_C KeyCode = 0x43
	KEY_D KeyCode = 0x44
	KEY_E KeyCode = 0x45
	KEY_F KeyCode = 0x46
	KEY_G KeyCode = 0x47
	KEY_H KeyCode = 0x48
	KEY_I KeyCode = 0x49
	KEY_J KeyCode = 0x4A
	KEY_K KeyCode = 0x4B
	KEY_L KeyCode = 0x4C
	KEY_M KeyCode = 0x4D
	KEY_N KeyCode = 0x4E
	KEY_O KeyCode = 0x4F
	KEY_P KeyCode = 0x50
	KEY_Q KeyCode = 0x51
	KEY_R KeyCode = 0x52
	KEY_S KeyCode = 0x53
	KEY_T KeyCode = 0x54
	KEY_U KeyCode = 0x55
	KEY_V KeyCode = 0x56
	KEY_W KeyCode = 0x57
	KEY_X KeyCode = 0x58
	KEY_Y KeyCode = 0x59
	KEY_Z KeyCode = 0x5A
)

type MouseBtn uint32

const (
	MOUSE_LEFTDOWN   MouseBtn = win.MOUSEEVENTF_LEFTDOWN
	MOUSE_LEFTUP     MouseBtn = win.MOUSEEVENTF_LEFTUP
	MOUSE_RIGHTDOWN  MouseBtn = win.MOUSEEVENTF_RIGHTDOWN
	MOUSE_RIGHTUP    MouseBtn = win.MOUSEEVENTF_RIGHTUP
	MOUSE_MIDDLEDOWN MouseBtn = win.MOUSEEVENTF_MIDDLEDOWN
	MOUSE_MIDDLEUP   MouseBtn = win.MOUSEEVENTF_MIDDLEUP
)

func Key(name string) KeyCode {
	name = strings.ToUpper(name)
	switch name {
	case "CAPSLOCK":
		return KEY_CAPITAL
	case "ALT":
		return KEY_MENU
	case "ALTLEFT":
		return KEY_LMENU
	case "ALTRIGHT":
		return KEY_RMENU
	case "CONTROL":
		return KEY_CONTROL
	case "CONTROLLEFT":
		return KEY_LCONTROL
	case "CONTROLRIGHT":
		return KEY_RCONTROL
	case "OSLEFT":
		return KEY_LWIN
	case "OSRIGHT":
		return KEY_RWIN
	case "SHIFT":
		return KEY_SHIFT
	case "SHIFTLEFT":
		return KEY_LSHIFT
	case "SHIFTRIGHT":
		return KEY_RSHIFT
	case "BACKSPACE":
		return KEY_BACK
	case "ENTER":
		return KEY_RETURN
	case "SPACE":
		return KEY_SPACE
	case "TAB":
		return KEY_TAB
	case "ESCCAPE":
		return KEY_ESCAPE
	case "INSERT":
		return KEY_INSERT
	case "DELETE":
		return KEY_DELETE
	case "HOME":
		return KEY_HOME
	case "END":
		return KEY_END

	case "PAGEDOWN":
		return KEY_NEXT
	case "PAGEUP":
		return KEY_PRIOR
	case "ARROWUP":
		return KEY_UP
	case "ARROWDOWN":
		return KEY_DOWN
	case "ARROWLEFT":
		return KEY_LEFT
	case "ARROWRIGHT":
		return KEY_RIGHT

	case "F1":
		return KEY_F1
	case "F2":
		return KEY_F2
	case "F3":
		return KEY_F3
	case "F4":
		return KEY_F4
	case "F5":
		return KEY_F5
	case "F6":
		return KEY_F6
	case "F7":
		return KEY_F7
	case "F8":
		return KEY_F8
	case "F9":
		return KEY_F9
	case "F10":
		return KEY_F10
	case "F11":
		return KEY_F11
	case "F12":
		return KEY_F12
	case "F13":
		return KEY_F13
	case "F14":
		return KEY_F14
	case "F15":
		return KEY_F15
	case "F16":
		return KEY_F16
	case "F17":
		return KEY_F17
	case "F18":
		return KEY_F18
	case "F19":
		return KEY_F19
	case "F20":
		return KEY_F20
	case "F21":
		return KEY_F21
	case "F22":
		return KEY_F22
	case "F23":
		return KEY_F23
	case "F24":
		return KEY_F24

	case "0", "DIGIT0":
		return KEY_0
	case "1", "DIGIT1":
		return KEY_1
	case "2", "DIGIT2":
		return KEY_2
	case "3", "DIGIT3":
		return KEY_3
	case "4", "DIGIT4":
		return KEY_4
	case "5", "DIGIT5":
		return KEY_5
	case "6", "DIGIT6":
		return KEY_6
	case "7", "DIGIT7":
		return KEY_7
	case "8", "DIGIT8":
		return KEY_8
	case "9", "DIGIT9":
		return KEY_9

	case "A", "KEYA":
		return KEY_A
	case "B", "KEYB":
		return KEY_B
	case "C", "KEYC":
		return KEY_C
	case "D", "KEYD":
		return KEY_D
	case "E", "KEYE":
		return KEY_E
	case "F", "KEYF":
		return KEY_F
	case "G", "KEYG":
		return KEY_G
	case "H", "KEYH":
		return KEY_H
	case "I", "KEYI":
		return KEY_I
	case "J", "KEYJ":
		return KEY_J
	case "K", "KEYK":
		return KEY_K
	case "L", "KEYL":
		return KEY_L
	case "M", "KEYM":
		return KEY_M
	case "N", "KEYN":
		return KEY_N
	case "O", "KEYO":
		return KEY_O
	case "P", "KEYP":
		return KEY_P
	case "Q", "KEYQ":
		return KEY_Q
	case "R", "KEYR":
		return KEY_R
	case "S", "KEYS":
		return KEY_S
	case "T", "KEYT":
		return KEY_T
	case "U", "KEYU":
		return KEY_U
	case "V", "KEYV":
		return KEY_V
	case "W", "KEYW":
		return KEY_W
	case "X", "KEYX":
		return KEY_X
	case "Y", "KEYY":
		return KEY_Y
	case "Z", "KEYZ":
		return KEY_Z

	case "[", "BRACKETLEFT":
	case "]", "BRACKETRIGHT":
	case "<":
	case ">":
	case "(":
	case ")":
	case "{":
	case "}":
	case "~":
	case "`", "BACKQUOTE":
	case "!":
	case "@":
	case "#":
	case "$":
	case "%":
	case "^":
	case "&":
	case "*":
	case "+":
	case "=", "EQUAL":
	case "_":
	case "-", "MINUS":
	case ":":
	case ";", "SEMICOLON":
	case "|":
	case `\`, "BACKSLASH":
	case `"`, "QUOTE":
	case "'":
	case ",", "COMMA":
	case ".", "PERIOD":
	case "/", "SLASH":
	case "?":
	}

	return 0
}
