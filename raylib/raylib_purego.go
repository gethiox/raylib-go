//go:build !cgo
// +build !cgo

package rl

import (
	"github.com/gen2brain/raylib-go/raylib/internal/convert"
	"github.com/jupiterrider/ffi"
)

var (
	// raylibDll is the pointer to the shared library
	raylibDll ffi.Lib

	initWindow          ffi.Fun
	closeWindow         ffi.Fun
	setTraceLogCallback ffi.Fun
)

func init() {
	raylibDll = loadLibrary()

	initWindow = must(raylibDll.Prep("InitWindow", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypePointer))
	closeWindow = must(raylibDll.Prep("CloseWindow", &ffi.TypeVoid, &ffi.TypeVoid))
	setTraceLogCallback = must(raylibDll.Prep("SetTraceLogCallback", &ffi.TypeVoid, &ffi.TypePointer))
}

// InitWindow - Initialize window and OpenGL context
func InitWindow(width int32, height int32, title string) {
	ctitle := convert.ToBytePtr(title)
	initWindow.Call(nil, &width, &height, &ctitle)
}

// CloseWindow - Close window and unload OpenGL context
func CloseWindow() {
	closeWindow.Call(nil)
}

// SetTraceLogCallback - Set custom trace log
func SetTraceLogCallback(fn TraceLogCallbackFun) {
	cb := traceLogCallbackWrapper(fn)
	setTraceLogCallback.Call(nil, &cb)
}

// GetMouseDelta - Get mouse delta between frames
func GetMouseDelta() Vector2 {
	return Vector2{}
}

// IsKeyDown - Check if a key is being pressed
func IsKeyDown(key int32) bool {
	return false
}

// IsGamepadAvailable - Check if a gamepad is available
func IsGamepadAvailable(gamepad int32) bool {
	return false
}

// GetGamepadAxisMovement - Get axis movement value for a gamepad axis
func GetGamepadAxisMovement(gamepad int32, axis int32) float32 {
	return 0
}

// IsMouseButtonDown - Check if a mouse button is being pressed
func IsMouseButtonDown(button MouseButton) bool {
	return false
}

// IsKeyPressed - Check if a key has been pressed once
func IsKeyPressed(key int32) bool {
	return false
}

// GetMouseWheelMove - Get mouse wheel movement for X or Y, whichever is larger
func GetMouseWheelMove() float32 {
	return 0
}

// GetFrameTime - Get time in seconds for last frame drawn (delta time)
func GetFrameTime() float32 {
	return 0
}
