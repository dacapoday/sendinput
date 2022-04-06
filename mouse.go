//go:build windows

package sendinput

import (
	"fmt"
)

// A Mouse is a device that will trigger an absolute change event.
type Mouse interface {
	// MoveLeft will move the mouse cursor left by the given number of pixel.
	MoveLeft(pixel int32) error

	// MoveRight will move the mouse cursor right by the given number of pixel.
	MoveRight(pixel int32) error

	// MoveUp will move the mouse cursor up by the given number of pixel.
	MoveUp(pixel int32) error

	// MoveDown will move the mouse cursor down by the given number of pixel.
	MoveDown(pixel int32) error

	// Move will perform a move of the mouse pointer along the x and y axes relative to the current position as requested.
	// Note that the upper left corner is (0, 0), so positive x and y means moving right (x) and down (y), whereas negative values will cause a move towards the upper left corner.
	Move(x, y int32) error

	// LeftClick will issue a single left click.
	LeftClick() error

	// RightClick will issue a right click.
	RightClick() error

	// MiddleClick will issue a middle click.
	MiddleClick() error

	// LeftPress will simulate a press of the left mouse button. Note that the button will not be released until LeftRelease is invoked.
	LeftPress() error

	// LeftRelease will simulate the release of the left mouse button.
	LeftRelease() error

	// RightPress will simulate the press of the right mouse button. Note that the button will not be released until RightRelease is invoked.
	RightPress() error

	// RightRelease will simulate the release of the right mouse button.
	RightRelease() error

	// MiddlePress will simulate the press of the middle mouse button. Note that the button will not be released until MiddleRelease is invoked.
	MiddlePress() error

	// MiddleRelease will simulate the release of the middle mouse button.
	MiddleRelease() error

	// Wheel will simulate a wheel movement.
	Wheel(horizontal bool, delta int32) error

	Close() error
}

type vMouse struct{}

// CreateMouse will create a new mouse input device. A mouse is a device that allows relative input.
// Relative input means that all changes to the x and y coordinates of the mouse pointer will be relative to the previous position.
// CreateMouse function signature for compatibility with https://github.com/bendahl/uinput, path and name is useless under windows.
func CreateMouse(path string, name []byte) (Mouse, error) {
	return vMouse{}, nil
}

// MoveLeft will move the cursor left by the number of pixel specified.
func (mouse vMouse) MoveLeft(pixel int32) (err error) {
	if err := assertNotNegative(pixel); err != nil {
		return err
	}
	return sendRelEvent(-pixel, 0)
}

// MoveRight will move the cursor right by the number of pixel specified.
func (mouse vMouse) MoveRight(pixel int32) (err error) {
	if err := assertNotNegative(pixel); err != nil {
		return err
	}
	return sendRelEvent(pixel, 0)
}

// MoveUp will move the cursor up by the number of pixel specified.
func (mouse vMouse) MoveUp(pixel int32) (err error) {
	if err := assertNotNegative(pixel); err != nil {
		return err
	}
	return sendRelEvent(0, -pixel)
}

// MoveDown will move the cursor down by the number of pixel specified.
func (mouse vMouse) MoveDown(pixel int32) (err error) {
	if err := assertNotNegative(pixel); err != nil {
		return err
	}
	return sendRelEvent(0, pixel)
}

// Move will perform a move of the mouse pointer along the x and y axes relative to the current position as requested.
// Note that the upper left corner is (0, 0), so positive x and y means moving right (x) and down (y), whereas negative
// values will cause a move towards the upper left corner.
func (mouse vMouse) Move(x, y int32) error {
	return sendRelEvent(x, y)
}

// LeftClick will issue a LeftClick.
func (mouse vMouse) LeftClick() error {
	err := sendBtnEvent(buttonLEFTDOWN)
	if err != nil {
		return fmt.Errorf("Failed to issue the LeftClick event: %v", err)
	}
	return sendBtnEvent(buttonLEFTUP)
}

// RightClick will issue a RightClick.
func (mouse vMouse) RightClick() error {
	err := sendBtnEvent(buttonRIGHTDOWN)
	if err != nil {
		return fmt.Errorf("Failed to issue the RightClick event: %v", err)
	}
	return sendBtnEvent(buttonRIGHTUP)
}

// MiddleClick will issue a MiddleClick.
func (mouse vMouse) MiddleClick() error {
	err := sendBtnEvent(buttonMIDDLEDOWN)
	if err != nil {
		return fmt.Errorf("Failed to issue the MiddleClick event: %v", err)
	}
	return sendBtnEvent(buttonMIDDLEUP)
}

// LeftPress will simulate a press of the left mouse button. Note that the button will not be released until LeftRelease is invoked.
func (mouse vMouse) LeftPress() error {
	return sendBtnEvent(buttonLEFTDOWN)
}

// LeftRelease will simulate the release of the left mouse button.
func (mouse vMouse) LeftRelease() error {
	return sendBtnEvent(buttonLEFTUP)
}

// RightPress will simulate the press of the right mouse button. Note that the button will not be released until RightRelease is invoked.
func (mouse vMouse) RightPress() error {
	return sendBtnEvent(buttonRIGHTDOWN)
}

// RightRelease will simulate the release of the right mouse button.
func (mouse vMouse) RightRelease() error {
	return sendBtnEvent(buttonRIGHTUP)
}

// MiddlePress will simulate the press of the middle mouse button. Note that the button will not be released until MiddleRelease is invoked.
func (mouse vMouse) MiddlePress() error {
	return sendBtnEvent(buttonMIDDLEDOWN)
}

// MiddleRelease will simulate the release of the middle mouse button.
func (mouse vMouse) MiddleRelease() (err error) {
	return sendBtnEvent(buttonMIDDLEUP)
}

// Wheel will simulate a wheel movement.
func (mouse vMouse) Wheel(horizontal bool, delta int32) error {
	return sendwhlEvent(horizontal, delta)
}

// Close closes the device and releases the device.
func (mouse vMouse) Close() (err error) { return }

func assertNotNegative(val int32) error {
	if val < 0 {
		return fmt.Errorf("%v is out of range. Expected a positive or zero value", val)
	}
	return nil
}
