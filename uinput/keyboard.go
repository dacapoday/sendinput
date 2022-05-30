package uinput

type Keyboard struct{}

// KeyPress will cause the key to be pressed and immediately released.
func (k Keyboard) KeyPress(key int) error

// KeyDown will send a keypress event to an existing keyboard device.
// The key can be any of the predefined keycodes from keycodes.go.
// Note that the key will be "held down" until "KeyUp" is called.
func (k Keyboard) KeyDown(key int) error

// KeyUp will send a keyrelease event to an existing keyboard device.
// The key can be any of the predefined keycodes from keycodes.go.
func (k Keyboard) KeyUp(key int) error

func (k Keyboard) Close() error
