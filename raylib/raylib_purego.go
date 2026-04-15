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

	initWindow               ffi.Fun
	closeWindow              ffi.Fun
	setTraceLogCallback      ffi.Fun
	windowShouldClose        ffi.Fun
	isWindowReady            ffi.Fun
	isWindowFullscreen       ffi.Fun
	isWindowHidden           ffi.Fun
	isWindowMinimized        ffi.Fun
	isWindowMaximized        ffi.Fun
	isWindowFocused          ffi.Fun
	isWindowResized          ffi.Fun
	isWindowState            ffi.Fun
	setWindowState           ffi.Fun
	clearWindowState         ffi.Fun
	toggleFullscreen         ffi.Fun
	toggleBorderlessWindowed ffi.Fun
	maximizeWindow           ffi.Fun
	minimizeWindow           ffi.Fun
	restoreWindow            ffi.Fun
	setWindowIcon            ffi.Fun
	setWindowIcons           ffi.Fun
)

func init() {
	raylibDll = loadLibrary()

	initWindow = must(raylibDll.Prep("InitWindow", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypePointer))
	closeWindow = must(raylibDll.Prep("CloseWindow", &ffi.TypeVoid))
	setTraceLogCallback = must(raylibDll.Prep("SetTraceLogCallback", &ffi.TypeVoid, &ffi.TypePointer))
	windowShouldClose = must(raylibDll.Prep("WindowShouldClose", &ffi.TypeUint8))
	isWindowReady = must(raylibDll.Prep("IsWindowReady", &ffi.TypeUint8))
	isWindowFullscreen = must(raylibDll.Prep("IsWindowFullscreen", &ffi.TypeUint8))
	isWindowHidden = must(raylibDll.Prep("IsWindowHidden", &ffi.TypeUint8))
	isWindowMinimized = must(raylibDll.Prep("IsWindowMinimized", &ffi.TypeUint8))
	isWindowMaximized = must(raylibDll.Prep("IsWindowMaximized", &ffi.TypeUint8))
	isWindowFocused = must(raylibDll.Prep("IsWindowFocused", &ffi.TypeUint8))
	isWindowResized = must(raylibDll.Prep("IsWindowResized", &ffi.TypeUint8))
	isWindowState = must(raylibDll.Prep("IsWindowState", &ffi.TypeUint8, &ffi.TypeUint32))
	setWindowState = must(raylibDll.Prep("SetWindowState", &ffi.TypeVoid, &ffi.TypeUint32))
	clearWindowState = must(raylibDll.Prep("ClearWindowState", &ffi.TypeVoid, &ffi.TypeUint32))
	toggleFullscreen = must(raylibDll.Prep("ToggleFullscreen", &ffi.TypeVoid))
	toggleBorderlessWindowed = must(raylibDll.Prep("ToggleBorderlessWindowed", &ffi.TypeVoid))
	maximizeWindow = must(raylibDll.Prep("MaximizeWindow", &ffi.TypeVoid))
	minimizeWindow = must(raylibDll.Prep("MinimizeWindow", &ffi.TypeVoid))
	restoreWindow = must(raylibDll.Prep("RestoreWindow", &ffi.TypeVoid))
	setWindowIcon = must(raylibDll.Prep("SetWindowIcon", &ffi.TypeVoid, &typeImage))
	setWindowIcons = must(raylibDll.Prep("SetWindowIcons", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypeSint32))
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

// WindowShouldClose - Check if application should close (KEY_ESCAPE pressed or windows close icon clicked)
func WindowShouldClose() bool {
	var ret ffi.Arg
	windowShouldClose.Call(&ret)
	return ret.Bool()
}

// IsWindowReady - Check if window has been initialized successfully
func IsWindowReady() bool {
	var ret ffi.Arg
	isWindowReady.Call(&ret)
	return ret.Bool()
}

// IsWindowFullscreen - Check if window is currently fullscreen
func IsWindowFullscreen() bool {
	var ret ffi.Arg
	isWindowFullscreen.Call(&ret)
	return ret.Bool()
}

// IsWindowHidden - Check if window is currently hidden (only PLATFORM_DESKTOP)
func IsWindowHidden() bool {
	var ret ffi.Arg
	isWindowHidden.Call(&ret)
	return ret.Bool()
}

// IsWindowMinimized - Check if window is currently minimized (only PLATFORM_DESKTOP)
func IsWindowMinimized() bool {
	var ret ffi.Arg
	isWindowMinimized.Call(&ret)
	return ret.Bool()
}

// IsWindowMaximized - Check if window is currently maximized (only PLATFORM_DESKTOP)
func IsWindowMaximized() bool {
	var ret ffi.Arg
	isWindowMaximized.Call(&ret)
	return ret.Bool()
}

// IsWindowFocused - Check if window is currently focused (only PLATFORM_DESKTOP)
func IsWindowFocused() bool {
	var ret ffi.Arg
	isWindowFocused.Call(&ret)
	return ret.Bool()
}

// IsWindowResized - Check if window has been resized last frame
func IsWindowResized() bool {
	var ret ffi.Arg
	isWindowResized.Call(&ret)
	return ret.Bool()
}

// IsWindowState - Check if one specific window flag is enabled
func IsWindowState(flag uint32) bool {
	var ret ffi.Arg
	isWindowState.Call(&ret, &flag)
	return ret.Bool()
}

// SetWindowState - Set window configuration state using flags (only PLATFORM_DESKTOP)
func SetWindowState(flags uint32) {
	setWindowState.Call(nil, &flags)
}

// ClearWindowState - Clear window configuration state flags
func ClearWindowState(flags uint32) {
	clearWindowState.Call(nil, &flags)
}

// ToggleFullscreen - Toggle window state: fullscreen/windowed (only PLATFORM_DESKTOP)
func ToggleFullscreen() {
	toggleFullscreen.Call(nil)
}

// ToggleBorderlessWindowed - Toggle window state: borderless windowed (only PLATFORM_DESKTOP)
func ToggleBorderlessWindowed() {
	toggleBorderlessWindowed.Call(nil)
}

// MaximizeWindow - Set window state: maximized, if resizable (only PLATFORM_DESKTOP)
func MaximizeWindow() {
	maximizeWindow.Call(nil)
}

// MinimizeWindow - Set window state: minimized, if resizable (only PLATFORM_DESKTOP)
func MinimizeWindow() {
	minimizeWindow.Call(nil)
}

// RestoreWindow - Set window state: not minimized/maximized (only PLATFORM_DESKTOP)
func RestoreWindow() {
	restoreWindow.Call(nil)
}

// SetWindowIcon - Set icon for window (single image, RGBA 32bit, only PLATFORM_DESKTOP)
func SetWindowIcon(image Image) {
	setWindowIcon.Call(nil, &image)
}

// SetWindowIcons - Set icon for window (multiple images, RGBA 32bit, only PLATFORM_DESKTOP)
func SetWindowIcons(images []Image, count int32) {
	imagesPtr := &images[0]
	setWindowIcons.Call(nil, &imagesPtr, &count)
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
