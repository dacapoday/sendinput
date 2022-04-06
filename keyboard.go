package sendinput

// A Keyboard is an key event output device. It is used to
// enable a program to simulate HID keyboard input events.
type Keyboard interface {
	// KeyPress will cause the key to be pressed and immediately released.
	KeyPress(key int) error

	// KeyDown will send a keypress event to an existing keyboard device.
	// The key can be any of the predefined keycodes from keycodes.go.
	// Note that the key will be "held down" until "KeyUp" is called.
	KeyDown(key int) error

	// KeyUp will send a keyrelease event to an existing keyboard device.
	// The key can be any of the predefined keycodes from keycodes.go.
	KeyUp(key int) error

	Close() error
}
