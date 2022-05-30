package uinput

type Mouse struct{}

// MoveLeft will move the mouse cursor left by the given number of pixel.
func (m Mouse) MoveLeft(pixel int32) error

// MoveRight will move the mouse cursor right by the given number of pixel.
func (m Mouse) MoveRight(pixel int32) error

// MoveUp will move the mouse cursor up by the given number of pixel.
func (m Mouse) MoveUp(pixel int32) error

// MoveDown will move the mouse cursor down by the given number of pixel.
func (m Mouse) MoveDown(pixel int32) error

// Move will perform a move of the mouse pointer along the x and y axes relative to the current position as requested.
// Note that the upper left corner is (0, 0), so positive x and y means moving right (x) and down (y), whereas negative values will cause a move towards the upper left corner.
func (m Mouse) Move(x, y int32) error

// LeftClick will issue a single left click.
func (m Mouse) LeftClick() error

// RightClick will issue a right click.
func (m Mouse) RightClick() error

// MiddleClick will issue a middle click.
func (m Mouse) MiddleClick() error

// LeftPress will simulate a press of the left mouse button. Note that the button will not be released until LeftRelease is invoked.
func (m Mouse) LeftPress() error

// LeftRelease will simulate the release of the left mouse button.
func (m Mouse) LeftRelease() error

// RightPress will simulate the press of the right mouse button. Note that the button will not be released until RightRelease is invoked.
func (m Mouse) RightPress() error

// RightRelease will simulate the release of the right mouse button.
func (m Mouse) RightRelease() error

// MiddlePress will simulate the press of the middle mouse button. Note that the button will not be released until MiddleRelease is invoked.
func (m Mouse) MiddlePress() error

// MiddleRelease will simulate the release of the middle mouse button.
func (m Mouse) MiddleRelease() error

// Wheel will simulate a wheel movement.
func (m Mouse) Wheel(horizontal bool, delta int32) error

func (m Mouse) Close() error
