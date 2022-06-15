//go:build windows

package main

import (
	"time"

	"github.com/dacapoday/sendinput"
)

var build = "unknown"

func main() {
	time.Sleep(5 * time.Second)
	sendinput.SendTextInput("Hello 中文测试 world!")

	time.Sleep(1 * time.Second)

	sendinput.SendMouseWhlInput(0, 1000)
	time.Sleep(1 * time.Second)
	sendinput.SendMouseWhlInput(0, -500)
	time.Sleep(1 * time.Second)
	sendinput.SendKeyboardInput(sendinput.KEY_0, true)
	sendinput.SendKeyboardInput(sendinput.KEY_0, false)

	sendinput.SendMouseBtnInput(sendinput.MOUSE_RIGHTDOWN)
	sendinput.SendMouseBtnInput(sendinput.MOUSE_RIGHTUP)

	sendinput.SendMouseRelInput(10, 0)
	sendinput.SendMouseRelInput(0, 20)

	time.Sleep(1 * time.Second)

	sendinput.SendMouseRelInput(-10, 0)
	sendinput.SendMouseRelInput(0, -10)

}
