//go:build !cgo
// +build !cgo

package rl

import (
	"fmt"
	"image/color"
	"unsafe"

	"github.com/gen2brain/raylib-go/raylib/internal/convert"
	"github.com/jupiterrider/ffi"
)

var (
	// dll is the pointer to the shared library
	dll ffi.Lib = loadLibrary()

	// Window-related functions

	initWindow               = dll.MustPrep("InitWindow", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypePointer)
	closeWindow              = dll.MustPrep("CloseWindow", &ffi.TypeVoid)
	windowShouldClose        = dll.MustPrep("WindowShouldClose", &ffi.TypeUint8)
	isWindowReady            = dll.MustPrep("IsWindowReady", &ffi.TypeUint8)
	isWindowFullscreen       = dll.MustPrep("IsWindowFullscreen", &ffi.TypeUint8)
	isWindowHidden           = dll.MustPrep("IsWindowHidden", &ffi.TypeUint8)
	isWindowMinimized        = dll.MustPrep("IsWindowMinimized", &ffi.TypeUint8)
	isWindowMaximized        = dll.MustPrep("IsWindowMaximized", &ffi.TypeUint8)
	isWindowFocused          = dll.MustPrep("IsWindowFocused", &ffi.TypeUint8)
	isWindowResized          = dll.MustPrep("IsWindowResized", &ffi.TypeUint8)
	isWindowState            = dll.MustPrep("IsWindowState", &ffi.TypeUint8, &ffi.TypeUint32)
	setWindowState           = dll.MustPrep("SetWindowState", &ffi.TypeVoid, &ffi.TypeUint32)
	clearWindowState         = dll.MustPrep("ClearWindowState", &ffi.TypeVoid, &ffi.TypeUint32)
	toggleFullscreen         = dll.MustPrep("ToggleFullscreen", &ffi.TypeVoid)
	toggleBorderlessWindowed = dll.MustPrep("ToggleBorderlessWindowed", &ffi.TypeVoid)
	maximizeWindow           = dll.MustPrep("MaximizeWindow", &ffi.TypeVoid)
	minimizeWindow           = dll.MustPrep("MinimizeWindow", &ffi.TypeVoid)
	restoreWindow            = dll.MustPrep("RestoreWindow", &ffi.TypeVoid)
	setWindowIcon            = dll.MustPrep("SetWindowIcon", &ffi.TypeVoid, &typeImage)
	setWindowIcons           = dll.MustPrep("SetWindowIcons", &ffi.TypeVoid, &ffi.TypePointer, &ffi.TypeSint32)
	setWindowTitle           = dll.MustPrep("SetWindowTitle", &ffi.TypeVoid, &ffi.TypePointer)
	setWindowPosition        = dll.MustPrep("SetWindowPosition", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32)
	setWindowMonitor         = dll.MustPrep("SetWindowMonitor", &ffi.TypeVoid, &ffi.TypeSint32)
	setWindowMinSize         = dll.MustPrep("SetWindowMinSize", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32)
	setWindowMaxSize         = dll.MustPrep("SetWindowMaxSize", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32)
	setWindowSize            = dll.MustPrep("SetWindowSize", &ffi.TypeSint32, &ffi.TypeSint32)
	setWindowOpacity         = dll.MustPrep("SetWindowOpacity", &ffi.TypeVoid, &ffi.TypeFloat)
	setWindowFocused         = dll.MustPrep("SetWindowFocused", &ffi.TypeVoid)
	getWindowHandle          = dll.MustPrep("GetWindowHandle", &ffi.TypePointer)
	getScreenWidth           = dll.MustPrep("GetWindowHandle", &ffi.TypeSint32)
	getScreenHeight          = dll.MustPrep("GetScreenHeight", &ffi.TypeSint32)
	getRenderWidth           = dll.MustPrep("GetRenderWidth", &ffi.TypeSint32)
	getRenderHeight          = dll.MustPrep("GetRenderHeight", &ffi.TypeSint32)
	getMonitorCount          = dll.MustPrep("GetMonitorCount", &ffi.TypeSint32)
	getCurrentMonitor        = dll.MustPrep("GetCurrentMonitor", &ffi.TypeSint32)
	getMonitorPosition       = dll.MustPrep("GetMonitorPosition", &typeVector2, &ffi.TypeSint32)
	getMonitorWidth          = dll.MustPrep("GetMonitorWidth", &ffi.TypeSint32, &ffi.TypeSint32)
	getMonitorHeight         = dll.MustPrep("GetMonitorHeight", &ffi.TypeSint32, &ffi.TypeSint32)
	getMonitorPhysicalWidth  = dll.MustPrep("GetMonitorPhysicalWidth", &ffi.TypeSint32, &ffi.TypeSint32)
	getMonitorPhysicalHeight = dll.MustPrep("GetMonitorPhysicalHeight", &ffi.TypeSint32, &ffi.TypeSint32)
	getMonitorRefreshRate    = dll.MustPrep("GetMonitorRefreshRate", &ffi.TypeSint32, &ffi.TypeSint32)
	getWindowPosition        = dll.MustPrep("GetWindowPosition", &typeVector2)
	getWindowScaleDPI        = dll.MustPrep("GetWindowScaleDPI", &typeVector2)
	getMonitorName           = dll.MustPrep("GetMonitorName", &ffi.TypePointer, &ffi.TypeSint32)
	setClipboardText         = dll.MustPrep("SetClipboardText", &ffi.TypeVoid, &ffi.TypePointer)
	getClipboardText         = dll.MustPrep("GetClipboardText", &ffi.TypePointer)
	getClipboardImage        = dll.MustPrep("GetClipboardImage", &typeImage)
	enableEventWaiting       = dll.MustPrep("EnableEventWaiting", &ffi.TypeVoid)
	disableEventWaiting      = dll.MustPrep("DisableEventWaiting", &ffi.TypeVoid)

	// Cursor-related functions

	showCursor       = dll.MustPrep("ShowCursor", &ffi.TypeVoid)
	hideCursor       = dll.MustPrep("HideCursor", &ffi.TypeVoid)
	isCursorHidden   = dll.MustPrep("IsCursorHidden", &ffi.TypeUint8)
	enableCursor     = dll.MustPrep("EnableCursor", &ffi.TypeVoid)
	disableCursor    = dll.MustPrep("DisableCursor", &ffi.TypeVoid)
	isCursorOnScreen = dll.MustPrep("IsCursorOnScreen", &ffi.TypeUint8)

	// Drawing-related functions

	clearBackground   = dll.MustPrep("ClearBackground", &ffi.TypeVoid, &typeColor)
	beginDrawing      = dll.MustPrep("BeginDrawing", &ffi.TypeVoid)
	endDrawing        = dll.MustPrep("EndDrawing", &ffi.TypeVoid)
	beginMode2D       = dll.MustPrep("BeginMode2D", &ffi.TypeVoid, &typeCamera2D)
	endMode2D         = dll.MustPrep("EndMode2D", &ffi.TypeVoid)
	beginMode3D       = dll.MustPrep("BeginMode3D", &ffi.TypeVoid, &typeCamera3D)
	endMode3D         = dll.MustPrep("EndMode3D", &ffi.TypeVoid)
	beginTextureMode  = dll.MustPrep("BeginTextureMode", &ffi.TypeVoid, &typeRenderTexture2D)
	endTextureMode    = dll.MustPrep("EndTextureMode", &ffi.TypeVoid)
	beginShaderMode   = dll.MustPrep("BeginShaderMode", &ffi.TypeVoid, &typeShader)
	endShaderMode     = dll.MustPrep("EndShaderMode", &ffi.TypeVoid)
	beginBlendMode    = dll.MustPrep("BeginBlendMode", &ffi.TypeVoid, &ffi.TypeSint32)
	endBlendMode      = dll.MustPrep("EndBlendMode", &ffi.TypeVoid)
	beginScissorMode  = dll.MustPrep("BeginScissorMode", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32)
	endScissorMode    = dll.MustPrep("EndScissorMode", &ffi.TypeVoid)
	beginVrStereoMode = dll.MustPrep("BeginVrStereoMode", &ffi.TypeVoid, &typeVrStereoConfig)
	endVrStereoMode   = dll.MustPrep("EndVrStereoMode", &ffi.TypeVoid)

	// VR stereo config functions for VR simulator

	loadVrStereoConfig   = dll.MustPrep("LoadVrStereoConfig", &typeVrStereoConfig, &typeVrDeviceInfo)
	unloadVrStereoConfig = dll.MustPrep("UnloadVrStereoConfig", &ffi.TypeVoid, &typeVrStereoConfig)

	// Shader management functions

	loadShader              = dll.MustPrep("LoadShader", &typeShader, &ffi.TypePointer, &ffi.TypePointer)
	loadShaderFromMemory    = dll.MustPrep("LoadShaderFromMemory", &typeShader, &ffi.TypePointer, &ffi.TypePointer)
	isShaderValid           = dll.MustPrep("IsShaderValid", &ffi.TypeUint8, &typeShader)
	getShaderLocation       = dll.MustPrep("GetShaderLocation", &ffi.TypeSint32, &typeShader, &ffi.TypePointer)
	getShaderLocationAttrib = dll.MustPrep("GetShaderLocationAttrib", &ffi.TypeSint32, &typeShader, &ffi.TypePointer)
	setShaderValue          = dll.MustPrep("SetShaderValue", &ffi.TypeVoid, &typeShader, &ffi.TypeSint32, &ffi.TypePointer, &ffi.TypeSint32)
	setShaderValueV         = dll.MustPrep("SetShaderValueV", &ffi.TypeVoid, &typeShader, &ffi.TypeSint32, &ffi.TypePointer, &ffi.TypeSint32, &ffi.TypeSint32)
	setShaderValueMatrix    = dll.MustPrep("SetShaderValueMatrix", &ffi.TypeVoid, &typeShader, &ffi.TypeSint32, &typeMatrix)
	setShaderValueTexture   = dll.MustPrep("SetShaderValueTexture", &ffi.TypeVoid, &typeShader, &ffi.TypeSint32, &typeTexture2D)
	unloadShader            = dll.MustPrep("UnloadShader", &ffi.TypeVoid, &typeShader)

	// Screen-space-related functions

	getScreenToWorldRay   = dll.MustPrep("GetScreenToWorldRay", &typeRay, &typeVector2, &typeCamera3D)
	getScreenToWorldRayEx = dll.MustPrep("GetScreenToWorldRayEx", &typeRay, &typeVector2, &typeCamera3D, &ffi.TypeSint32, &ffi.TypeSint32)
	getWorldToScreen      = dll.MustPrep("GetWorldToScreen", &typeVector2, &typeVector3, &typeCamera3D)
	getWorldToScreenEx    = dll.MustPrep("GetWorldToScreenEx", &typeVector2, &typeVector3, &typeCamera3D, &ffi.TypeSint32, &ffi.TypeSint32)
	getWorldToScreen2D    = dll.MustPrep("GetWorldToScreen2D", &typeVector2, &typeVector2, &typeCamera2D)
	getScreenToWorld2D    = dll.MustPrep("GetScreenToWorld2D", &typeVector2, &typeVector2, &typeCamera2D)
	getCameraMatrix       = dll.MustPrep("GetCameraMatrix", &typeMatrix, &typeCamera3D)
	getCameraMatrix2D     = dll.MustPrep("GetCameraMatrix2D", &typeMatrix, &typeCamera2D)

	// Timing-related functions

	setTargetFPS = dll.MustPrep("SetTargetFPS", &ffi.TypeVoid, &ffi.TypeSint32)
	getFrameTime = dll.MustPrep("GetFrameTime", &ffi.TypeFloat)
	getTime      = dll.MustPrep("GetTime", &ffi.TypeDouble)
	getFPS       = dll.MustPrep("GetFPS", &ffi.TypeSint32)

	// Custom frame control functions

	swapScreenBuffer = dll.MustPrep("SwapScreenBuffer", &ffi.TypeVoid)
	pollInputEvents  = dll.MustPrep("PollInputEvents", &ffi.TypeVoid)
	waitTime         = dll.MustPrep("WaitTime", &ffi.TypeVoid, &ffi.TypeDouble)

	// Misc. functions

	takeScreenshot = dll.MustPrep("TakeScreenshot", &ffi.TypeVoid, &ffi.TypePointer)
	setConfigFlags = dll.MustPrep("SetConfigFlags", &ffi.TypeVoid, &ffi.TypeUint32)
	openURL        = dll.MustPrep("OpenURL", &ffi.TypeVoid, &ffi.TypePointer)

	// Logging system

	setTraceLogLevel    = dll.MustPrep("SetTraceLogLevel", &ffi.TypeVoid, &ffi.TypeSint32)
	traceLog            = dll.MustPrepVar("TraceLog", 2, &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypePointer)
	setTraceLogCallback = dll.MustPrep("SetTraceLogCallback", &ffi.TypeVoid, &ffi.TypePointer)

	// File system management functions

	isFileDropped      = dll.MustPrep("IsFileDropped", &ffi.TypeUint8)
	loadDroppedFiles   = dll.MustPrep("LoadDroppedFiles", &typeFilePathList)
	unloadDroppedFiles = dll.MustPrep("UnloadDroppedFiles", &ffi.TypeVoid, &typeFilePathList)

	// Automation events functionality

	loadAutomationEventList       = dll.MustPrep("LoadAutomationEventList", &typeAutomationEventList, &ffi.TypePointer)
	unloadAutomationEventList     = dll.MustPrep("UnloadAutomationEventList", &ffi.TypeVoid, &typeAutomationEventList)
	exportAutomationEventList     = dll.MustPrep("ExportAutomationEventList", &ffi.TypeUint8, &typeAutomationEventList, &ffi.TypePointer)
	setAutomationEventList        = dll.MustPrep("SetAutomationEventList", &ffi.TypeVoid, &ffi.TypePointer)
	setAutomationEventBaseFrame   = dll.MustPrep("SetAutomationEventBaseFrame", &ffi.TypeVoid, &ffi.TypeSint32)
	startAutomationEventRecording = dll.MustPrep("StartAutomationEventRecording", &ffi.TypeVoid)
	stopAutomationEventRecording  = dll.MustPrep("StopAutomationEventRecording", &ffi.TypeVoid)
	playAutomationEvent           = dll.MustPrep("PlayAutomationEvent", &ffi.TypeVoid, &typeAutomationEvent)

	// Input-related functions: keyboard

	isKeyPressed       = dll.MustPrep("IsKeyPressed", &ffi.TypeUint8, &ffi.TypeSint32)
	isKeyPressedRepeat = dll.MustPrep("IsKeyPressedRepeat", &ffi.TypeUint8, &ffi.TypeSint32)
	isKeyDown          = dll.MustPrep("IsKeyDown", &ffi.TypeUint8, &ffi.TypeSint32)
	isKeyReleased      = dll.MustPrep("IsKeyReleased", &ffi.TypeUint8, &ffi.TypeSint32)
	isKeyUp            = dll.MustPrep("IsKeyUp", &ffi.TypeUint8, &ffi.TypeSint32)
	getKeyPressed      = dll.MustPrep("GetKeyPressed", &ffi.TypeSint32)
	getCharPressed     = dll.MustPrep("GetCharPressed", &ffi.TypeSint32)
	getKeyName         = dll.MustPrep("GetKeyName", &ffi.TypePointer, &ffi.TypeSint32)
	setExitKey         = dll.MustPrep("SetExitKey", &ffi.TypeVoid, &ffi.TypeSint32)

	// Input-related functions: gamepads

	isGamepadAvailable      = dll.MustPrep("IsGamepadAvailable", &ffi.TypeUint8, &ffi.TypeSint32)
	getGamepadName          = dll.MustPrep("GetGamepadName", &ffi.TypePointer, &ffi.TypeSint32)
	isGamepadButtonPressed  = dll.MustPrep("IsGamepadButtonPressed", &ffi.TypeUint8, &ffi.TypeSint32, &ffi.TypeSint32)
	isGamepadButtonDown     = dll.MustPrep("IsGamepadButtonDown", &ffi.TypeUint8, &ffi.TypeSint32, &ffi.TypeSint32)
	isGamepadButtonReleased = dll.MustPrep("IsGamepadButtonReleased", &ffi.TypeUint8, &ffi.TypeSint32, &ffi.TypeSint32)
	isGamepadButtonUp       = dll.MustPrep("IsGamepadButtonUp", &ffi.TypeUint8, &ffi.TypeSint32, &ffi.TypeSint32)
	getGamepadButtonPressed = dll.MustPrep("GetGamepadButtonPressed", &ffi.TypeSint32)
	getGamepadAxisCount     = dll.MustPrep("GetGamepadAxisCount", &ffi.TypeSint32, &ffi.TypeSint32)
	getGamepadAxisMovement  = dll.MustPrep("GetGamepadAxisMovement", &ffi.TypeFloat, &ffi.TypeSint32, &ffi.TypeSint32)
	setGamepadMappings      = dll.MustPrep("SetGamepadMappings", &ffi.TypeSint32, &ffi.TypePointer)
	setGamepadVibration     = dll.MustPrep("SetGamepadVibration", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat)

	// Input-related functions: mouse

	isMouseButtonPressed  = dll.MustPrep("IsMouseButtonPressed", &ffi.TypeUint8, &ffi.TypeSint32)
	isMouseButtonDown     = dll.MustPrep("IsMouseButtonDown", &ffi.TypeUint8, &ffi.TypeSint32)
	isMouseButtonReleased = dll.MustPrep("IsMouseButtonReleased", &ffi.TypeUint8, &ffi.TypeSint32)
	isMouseButtonUp       = dll.MustPrep("IsMouseButtonUp", &ffi.TypeUint8, &ffi.TypeSint32)
	getMouseX             = dll.MustPrep("GetMouseX", &ffi.TypeSint32)
	getMouseY             = dll.MustPrep("GetMouseY", &ffi.TypeSint32)
	getMousePosition      = dll.MustPrep("GetMousePosition", &typeVector2)
	getMouseDelta         = dll.MustPrep("GetMouseDelta", &typeVector2)
	setMousePosition      = dll.MustPrep("SetMousePosition", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32)
	setMouseOffset        = dll.MustPrep("SetMouseOffset", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32)
	setMouseScale         = dll.MustPrep("SetMouseScale", &ffi.TypeVoid, &ffi.TypeFloat, &ffi.TypeFloat)
	getMouseWheelMove     = dll.MustPrep("GetMouseWheelMove", &ffi.TypeFloat)
	getMouseWheelMoveV    = dll.MustPrep("GetMouseWheelMoveV", &typeVector2)
	setMouseCursor        = dll.MustPrep("SetMouseCursor", &ffi.TypeVoid, &ffi.TypeSint32)

	// Input-related functions: touch

	getTouchX          = dll.MustPrep("GetTouchX", &ffi.TypeSint32)
	getTouchY          = dll.MustPrep("GetTouchY", &ffi.TypeSint32)
	getTouchPosition   = dll.MustPrep("GetTouchPosition", &typeVector2, &ffi.TypeSint32)
	getTouchPointId    = dll.MustPrep("GetTouchPointId", &ffi.TypeSint32, &ffi.TypeSint32)
	getTouchPointCount = dll.MustPrep("GetTouchPointCount", &ffi.TypeSint32)

	// Gestures and Touch Handling Functions (Module: rgestures)

	setGesturesEnabled     = dll.MustPrep("SetGesturesEnabled", &ffi.TypeVoid, &ffi.TypeUint32)
	isGestureDetected      = dll.MustPrep("IsGestureDetected", &ffi.TypeUint8, &ffi.TypeUint32)
	getGestureDetected     = dll.MustPrep("GetGestureDetected", &ffi.TypeSint32)
	getGestureHoldDuration = dll.MustPrep("GetGestureHoldDuration", &ffi.TypeFloat)
	getGestureDragVector   = dll.MustPrep("GetGestureDragVector", &typeVector2)
	getGestureDragAngle    = dll.MustPrep("GetGestureDragAngle", &ffi.TypeFloat)
	getGesturePinchVector  = dll.MustPrep("GetGesturePinchVector", &typeVector2)
	getGesturePinchAngle   = dll.MustPrep("GetGesturePinchAngle", &ffi.TypeFloat)
)

// InitWindow - Initialize window and OpenGL context
func InitWindow(width int32, height int32, title string) {
	titlePtr := convert.ToBytePtr(title)
	initWindow.Call(nil, &width, &height, &titlePtr)
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

// SetWindowTitle - Set title for window (only PLATFORM_DESKTOP and PLATFORM_WEB)
func SetWindowTitle(title string) {
	titlePtr := convert.ToBytePtr(title)
	setWindowTitle.Call(nil, &titlePtr)
}

// SetWindowPosition - Set window position on screen (only PLATFORM_DESKTOP)
func SetWindowPosition(x int, y int) {
	posX, posY := int32(x), int32(y)
	setWindowPosition.Call(nil, &posX, &posY)
}

// SetWindowMonitor - Set monitor for the current window
func SetWindowMonitor(monitor int) {
	m := int32(monitor)
	setWindowMonitor.Call(nil, &m)
}

// SetWindowMinSize - Set window minimum dimensions (for FLAG_WINDOW_RESIZABLE)
func SetWindowMinSize(width int, height int) {
	w, h := int32(width), int32(height)
	setWindowMinSize.Call(nil, &w, &h)
}

// SetWindowMaxSize - Set window maximum dimensions (for FLAG_WINDOW_RESIZABLE)
func SetWindowMaxSize(width int, height int) {
	w, h := int32(width), int32(height)
	setWindowMaxSize.Call(nil, &w, &h)
}

// SetWindowSize - Set window dimensions
func SetWindowSize(width int, height int) {
	w, h := int32(width), int32(height)
	setWindowSize.Call(nil, &w, &h)
}

// SetWindowOpacity - Set window opacity [0.0f..1.0f] (only PLATFORM_DESKTOP)
func SetWindowOpacity(opacity float32) {
	setWindowOpacity.Call(nil, &opacity)
}

// SetWindowFocused - Set window focused (only PLATFORM_DESKTOP)
func SetWindowFocused() {
	setWindowFocused.Call(nil)
}

// GetWindowHandle - Get native window handle
func GetWindowHandle() unsafe.Pointer {
	var ret unsafe.Pointer
	getWindowHandle.Call(&ret)
	return ret
}

// GetScreenWidth - Get current screen width
func GetScreenWidth() int {
	var ret ffi.Arg
	getScreenWidth.Call(&ret)
	return int(ret)
}

// GetScreenHeight - Get current screen height
func GetScreenHeight() int {
	var ret ffi.Arg
	getScreenHeight.Call(&ret)
	return int(ret)
}

// GetRenderWidth - Get current render width (it considers HiDPI)
func GetRenderWidth() int {
	var ret ffi.Arg
	getRenderWidth.Call(&ret)
	return int(ret)
}

// GetRenderHeight - Get current render height (it considers HiDPI)
func GetRenderHeight() int {
	var ret ffi.Arg
	getRenderHeight.Call(&ret)
	return int(ret)
}

// GetMonitorCount - Get number of connected monitors
func GetMonitorCount() int {
	var ret ffi.Arg
	getMonitorCount.Call(&ret)
	return int(ret)
}

// GetCurrentMonitor - Get current monitor where window is placed
func GetCurrentMonitor() int {
	var ret ffi.Arg
	getCurrentMonitor.Call(&ret)
	return int(ret)
}

// GetMonitorPosition - Get specified monitor position
func GetMonitorPosition(monitor int) Vector2 {
	var ret Vector2
	m := int32(monitor)
	getMonitorPosition.Call(&ret, &m)
	return ret
}

// GetMonitorWidth - Get specified monitor width (current video mode used by monitor)
func GetMonitorWidth(monitor int) int {
	var ret ffi.Arg
	m := int32(monitor)
	getMonitorWidth.Call(&ret, &m)
	return int(ret)
}

// GetMonitorHeight - Get specified monitor height (current video mode used by monitor)
func GetMonitorHeight(monitor int) int {
	var ret ffi.Arg
	m := int32(monitor)
	getMonitorHeight.Call(&ret, &m)
	return int(ret)
}

// GetMonitorPhysicalWidth - Get specified monitor physical width in millimetres
func GetMonitorPhysicalWidth(monitor int) int {
	var ret ffi.Arg
	m := int32(monitor)
	getMonitorPhysicalWidth.Call(&ret, &m)
	return int(ret)
}

// GetMonitorPhysicalHeight - Get specified monitor physical height in millimetres
func GetMonitorPhysicalHeight(monitor int) int {
	var ret ffi.Arg
	m := int32(monitor)
	getMonitorPhysicalHeight.Call(&ret, &m)
	return int(ret)
}

// GetMonitorRefreshRate - Get specified monitor refresh rate
func GetMonitorRefreshRate(monitor int) int {
	var ret ffi.Arg
	m := int32(monitor)
	getMonitorRefreshRate.Call(&ret, &m)
	return int(ret)
}

// GetWindowPosition - Get window position XY on monitor
func GetWindowPosition() Vector2 {
	var ret Vector2
	getWindowPosition.Call(&ret)
	return ret
}

// GetWindowScaleDPI - Get window scale DPI factor
func GetWindowScaleDPI() Vector2 {
	var ret Vector2
	getWindowScaleDPI.Call(&ret)
	return ret
}

// GetMonitorName - Get the human-readable, UTF-8 encoded name of the specified monitor
func GetMonitorName(monitor int) string {
	var ret *byte
	m := int32(monitor)
	getMonitorName.Call(&ret, &m)
	return convert.ToString(ret)
}

// SetClipboardText - Set clipboard text content
func SetClipboardText(text string) {
	textPtr := convert.ToBytePtr(text)
	setClipboardText.Call(nil, &textPtr)
}

// GetClipboardText - Get clipboard text content
func GetClipboardText() string {
	var ret *byte
	getClipboardText.Call(&ret)
	return convert.ToString(ret)
}

// GetClipboardImage - Get clipboard image content
//
// Only works with SDL3 backend or Windows with RGFW/GLFW
func GetClipboardImage() Image {
	var ret Image
	getClipboardImage.Call(&ret)
	return ret
}

// EnableEventWaiting - Enable waiting for events on EndDrawing(), no automatic event polling
func EnableEventWaiting() {
	enableEventWaiting.Call(nil)
}

// DisableEventWaiting - Disable waiting for events on EndDrawing(), automatic events polling
func DisableEventWaiting() {
	disableEventWaiting.Call(nil)
}

// ShowCursor - Shows cursor
func ShowCursor() {
	showCursor.Call(nil)
}

// HideCursor - Hides cursor
func HideCursor() {
	hideCursor.Call(nil)
}

// IsCursorHidden - Check if cursor is not visible
func IsCursorHidden() bool {
	var ret ffi.Arg
	isCursorHidden.Call(&ret)
	return ret.Bool()
}

// EnableCursor - Enables cursor (unlock cursor)
func EnableCursor() {
	enableCursor.Call(nil)
}

// DisableCursor - Disables cursor (lock cursor)
func DisableCursor() {
	disableCursor.Call(nil)
}

// IsCursorOnScreen - Check if cursor is on the screen
func IsCursorOnScreen() bool {
	var ret ffi.Arg
	isCursorOnScreen.Call(&ret)
	return ret.Bool()
}

// ClearBackground - Set background color (framebuffer clear color)
func ClearBackground(col color.RGBA) {
	clearBackground.Call(nil, col)
}

// BeginDrawing - Setup canvas (framebuffer) to start drawing
func BeginDrawing() {
	beginDrawing.Call(nil)
}

// EndDrawing - End canvas drawing and swap buffers (double buffering)
func EndDrawing() {
	endDrawing.Call(nil)
}

// BeginMode2D - Begin 2D mode with custom camera (2D)
func BeginMode2D(camera Camera2D) {
	beginMode2D.Call(nil, &camera)
}

// EndMode2D - Ends 2D mode with custom camera
func EndMode2D() {
	endMode2D.Call(nil)
}

// BeginMode3D - Begin 3D mode with custom camera (3D)
func BeginMode3D(camera Camera3D) {
	beginMode3D.Call(nil, &camera)
}

// EndMode3D - Ends 3D mode and returns to default 2D orthographic mode
func EndMode3D() {
	endMode3D.Call(nil)
}

// BeginTextureMode - Begin drawing to render texture
func BeginTextureMode(target RenderTexture2D) {
	beginTextureMode.Call(nil, &target)
}

// EndTextureMode - Ends drawing to render texture
func EndTextureMode() {
	endTextureMode.Call(nil)
}

// BeginShaderMode - Begin custom shader drawing
func BeginShaderMode(shader Shader) {
	beginShaderMode.Call(nil, &shader)
}

// EndShaderMode - End custom shader drawing (use default shader)
func EndShaderMode() {
	endShaderMode.Call(nil)
}

// BeginBlendMode - Begin blending mode (alpha, additive, multiplied, subtract, custom)
func BeginBlendMode(mode BlendMode) {
	beginBlendMode.Call(nil, &mode)
}

// EndBlendMode - End blending mode (reset to default: alpha blending)
func EndBlendMode() {
	endBlendMode.Call(nil)
}

// BeginScissorMode - Begin scissor mode (define screen area for following drawing)
func BeginScissorMode(x int32, y int32, width int32, height int32) {
	beginScissorMode.Call(nil, &x, &y, &width, &height)
}

// EndScissorMode - End scissor mode
func EndScissorMode() {
	endScissorMode.Call(nil)
}

// BeginVrStereoMode - Begin stereo rendering (requires VR simulator)
func BeginVrStereoMode(config VrStereoConfig) {
	beginVrStereoMode.Call(nil, &config)
}

// EndVrStereoMode - End stereo rendering (requires VR simulator)
func EndVrStereoMode() {
	endVrStereoMode.Call(nil)
}

// LoadVrStereoConfig - Load VR stereo config for VR simulator device parameters
func LoadVrStereoConfig(device VrDeviceInfo) VrStereoConfig {
	var ret VrStereoConfig
	loadVrStereoConfig.Call(&ret, &device)
	return ret
}

// UnloadVrStereoConfig - Unload VR stereo config
func UnloadVrStereoConfig(config VrStereoConfig) {
	unloadVrStereoConfig.Call(nil, &config)
}

// LoadShader - Load shader from files and bind default locations
func LoadShader(vsFileName string, fsFileName string) Shader {
	var ret Shader
	// "" becomes NULL to either load a vertex or fragment shader: https://github.com/gen2brain/raylib-go/issues/172
	vsFileNamePtr, fsFileNamePtr := convert.ToBytePtrNullable(vsFileName), convert.ToBytePtrNullable(fsFileName)
	loadShader.Call(&ret, &vsFileNamePtr, &fsFileNamePtr)
	return ret
}

// LoadShaderFromMemory - Load shader from code strings and bind default locations
func LoadShaderFromMemory(vsCode string, fsCode string) Shader {
	var ret Shader
	// "" becomes NULL to either load a vertex or fragment shader: https://github.com/gen2brain/raylib-go/issues/172
	vsCodePtr, fsCodePtr := convert.ToBytePtrNullable(vsCode), convert.ToBytePtrNullable(fsCode)
	loadShaderFromMemory.Call(&ret, &vsCodePtr, &fsCodePtr)
	return ret
}

// IsShaderValid - Check if a shader is valid (loaded on GPU)
func IsShaderValid(shader Shader) bool {
	var ret ffi.Arg
	isShaderValid.Call(&ret, &shader)
	return ret.Bool()
}

// GetShaderLocation - Get shader uniform location
func GetShaderLocation(shader Shader, uniformName string) int32 {
	var ret ffi.Arg
	uniformNamePtr := convert.ToBytePtr(uniformName)
	getShaderLocation.Call(&ret, &shader, &uniformNamePtr)
	return int32(ret)
}

// GetShaderLocationAttrib - Get shader attribute location
func GetShaderLocationAttrib(shader Shader, attribName string) int32 {
	var ret ffi.Arg
	attribNamePtr := convert.ToBytePtr(attribName)
	getShaderLocationAttrib.Call(&ret, &shader, &attribNamePtr)
	return int32(ret)
}

// SetShaderValue - Set shader uniform value
func SetShaderValue(shader Shader, locIndex int32, value []float32, uniformType ShaderUniformDataType) {
	valuePtr := &value[0]
	setShaderValue.Call(nil, &shader, &locIndex, &valuePtr, &uniformType)
}

// SetShaderValueV - Set shader uniform value vector
func SetShaderValueV(shader Shader, locIndex int32, value []float32, uniformType ShaderUniformDataType, count int32) {
	valuePtr := &value[0]
	setShaderValueV.Call(nil, &shader, &locIndex, &valuePtr, &uniformType, &count)
}

// SetShaderValueMatrix - Set shader uniform value (matrix 4x4)
func SetShaderValueMatrix(shader Shader, locIndex int32, mat Matrix) {
	setShaderValueMatrix.Call(nil, &shader, &locIndex, &mat)
}

// SetShaderValueTexture - Set shader uniform value for texture (sampler2d)
func SetShaderValueTexture(shader Shader, locIndex int32, texture Texture2D) {
	setShaderValueTexture.Call(nil, &shader, &locIndex, &texture)
}

// UnloadShader - Unload shader from GPU memory (VRAM)
func UnloadShader(shader Shader) {
	unloadShader.Call(nil, &shader)
}

// GetMouseRay - Get a ray trace from mouse position
//
// Deprecated: Use [GetScreenToWorldRay] instead.
func GetMouseRay(mousePosition Vector2, camera Camera) Ray {
	return GetScreenToWorldRay(mousePosition, camera)
}

// GetScreenToWorldRay - Get a ray trace from screen position (i.e mouse)
func GetScreenToWorldRay(position Vector2, camera Camera) Ray {
	var ret Ray
	getScreenToWorldRay.Call(&ret, &position, &camera)
	return ret
}

// GetScreenToWorldRayEx - Get a ray trace from screen position (i.e mouse) in a viewport
func GetScreenToWorldRayEx(position Vector2, camera Camera, width, height int32) Ray {
	var ret Ray
	getScreenToWorldRayEx.Call(&ret, &position, &camera, &width, &height)
	return ret
}

// GetWorldToScreen - Get the screen space position for a 3d world space position
func GetWorldToScreen(position Vector3, camera Camera) Vector2 {
	var ret Vector2
	getWorldToScreen.Call(&ret, &position, &camera)
	return ret
}

// GetWorldToScreenEx - Get size position for a 3d world space position
func GetWorldToScreenEx(position Vector3, camera Camera, width int32, height int32) Vector2 {
	var ret Vector2
	getWorldToScreenEx.Call(&ret, &position, &camera, &width, &height)
	return ret
}

// GetWorldToScreen2D - Get the screen space position for a 2d camera world space position
func GetWorldToScreen2D(position Vector2, camera Camera2D) Vector2 {
	var ret Vector2
	getWorldToScreen2D.Call(&ret, &position, &camera)
	return ret
}

// GetScreenToWorld2D - Get the world space position for a 2d camera screen space position
func GetScreenToWorld2D(position Vector2, camera Camera2D) Vector2 {
	var ret Vector2
	getScreenToWorld2D.Call(&ret, &position, &camera)
	return ret
}

// GetCameraMatrix - Get camera transform matrix (view matrix)
func GetCameraMatrix(camera Camera) Matrix {
	var ret Matrix
	getCameraMatrix.Call(&ret, &camera)
	return ret
}

// GetCameraMatrix2D - Get camera 2d transform matrix
func GetCameraMatrix2D(camera Camera2D) Matrix {
	var ret Matrix
	getCameraMatrix2D.Call(&ret, &camera)
	return ret
}

// SetTargetFPS - Set target FPS (maximum)
func SetTargetFPS(fps int32) {
	setTargetFPS.Call(nil, &fps)
}

// GetFrameTime - Get time in seconds for last frame drawn (delta time)
func GetFrameTime() float32 {
	var ret float32
	getFrameTime.Call(&ret)
	return ret
}

// GetTime - Get elapsed time in seconds since InitWindow()
func GetTime() float64 {
	var ret float64
	getTime.Call(&ret)
	return ret
}

// GetFPS - Get current FPS
func GetFPS() int32 {
	var ret ffi.Arg
	getFPS.Call(&ret)
	return int32(ret)
}

// Custom frame control functions
// NOTE: SwapScreenBuffer and PollInputEvents are intended for advanced users that want full control over the frame processing
// By default EndDrawing() does this job: draws everything + SwapScreenBuffer() + manage frame timing + PollInputEvents()
// To avoid that behaviour and control frame processes manually you must recompile raylib with SUPPORT_CUSTOM_FRAME_CONTROL enabled in config.h
//
// See: https://github.com/gen2brain/raylib-go/issues/378

// SwapScreenBuffer - Swap back buffer with front buffer (screen drawing)
func SwapScreenBuffer() {
	swapScreenBuffer.Call(nil)
}

// PollInputEvents - Register all input events
func PollInputEvents() {
	pollInputEvents.Call(nil)
}

// WaitTime - Wait for some time (halt program execution)
func WaitTime(seconds float64) {
	waitTime.Call(nil, &seconds)
}

// TakeScreenshot - Takes a screenshot of current screen (filename extension defines format)
func TakeScreenshot(fileName string) {
	fileNamePtr := convert.ToBytePtr(fileName)
	takeScreenshot.Call(nil, &fileNamePtr)
}

// SetConfigFlags - Setup init configuration flags (view FLAGS)
func SetConfigFlags(flags uint32) {
	setConfigFlags.Call(nil, &flags)
}

// OpenURL - Open URL with default system browser (if available)
func OpenURL(url string) {
	urlPtr := convert.ToBytePtr(url)
	openURL.Call(nil, &urlPtr)
}

// TraceLog - Show trace log messages (LOG_DEBUG, LOG_INFO, LOG_WARNING, LOG_ERROR...)
func TraceLog(logLevel TraceLogLevel, text string, args ...any) {
	level := int32(logLevel)
	textPtr := convert.ToBytePtr(fmt.Sprintf(text, args...))
	traceLog.Call(nil, &level, &textPtr)
}

// SetTraceLogLevel - Set the current threshold (minimum) log level
func SetTraceLogLevel(logLevel TraceLogLevel) {
	level := int32(logLevel)
	setTraceLogLevel.Call(nil, &level)
}

// SetTraceLogCallback - Set custom trace log
func SetTraceLogCallback(fn TraceLogCallbackFun) {
	cb := traceLogCallbackWrapper(fn)
	setTraceLogCallback.Call(nil, &cb)
}

// IsFileDropped - Check if a file has been dropped into window
func IsFileDropped() bool {
	var ret ffi.Arg
	isFileDropped.Call(&ret)
	return ret.Bool()
}

// LoadDroppedFiles - Load dropped filepaths
func LoadDroppedFiles() []string {
	var filePathList = struct {
		count uint32
		paths **byte
	}{}
	loadDroppedFiles.Call(&filePathList)
	defer unloadDroppedFiles.Call(nil, &filePathList)

	paths := unsafe.Slice(filePathList.paths, filePathList.count)
	result := make([]string, filePathList.count)

	for i := range paths {
		result[i] = convert.ToString(paths[i])
	}

	return result
}

// UnloadDroppedFiles - Unload dropped filepaths
func UnloadDroppedFiles() {}

// LoadAutomationEventList - Load automation events list from file, NULL for empty list, capacity = MAX_AUTOMATION_EVENTS
func LoadAutomationEventList(fileName string) AutomationEventList {
	var ret AutomationEventList
	fileNamePtr := convert.ToBytePtr(fileName)
	loadAutomationEventList.Call(&ret, &fileNamePtr)
	return ret
}

// UnloadAutomationEventList - Unload automation events list from file
func UnloadAutomationEventList(list *AutomationEventList) {
	unloadAutomationEventList.Call(nil, list)
}

// ExportAutomationEventList - Export automation events list as text file
func ExportAutomationEventList(list AutomationEventList, fileName string) bool {
	var ret ffi.Arg
	fileNamePtr := convert.ToBytePtr(fileName)
	exportAutomationEventList.Call(&ret, &list, &fileNamePtr)
	return ret.Bool()
}

// SetAutomationEventList - Set automation event list to record to
func SetAutomationEventList(list *AutomationEventList) {
	setAutomationEventList.Call(nil, &list)
}

// SetAutomationEventBaseFrame - Set automation event internal base frame to start recording
func SetAutomationEventBaseFrame(frame int) {
	f := int32(frame)
	setAutomationEventBaseFrame.Call(nil, &f)
}

// StartAutomationEventRecording - Start recording automation events (AutomationEventList must be set)
func StartAutomationEventRecording() {
	startAutomationEventRecording.Call(nil)
}

// StopAutomationEventRecording - Stop recording automation events
func StopAutomationEventRecording() {
	stopAutomationEventRecording.Call(nil)
}

// PlayAutomationEvent - Play a recorded automation event
func PlayAutomationEvent(event AutomationEvent) {
	playAutomationEvent.Call(nil, &event)
}

// IsKeyPressed - Check if a key has been pressed once
func IsKeyPressed(key int32) bool {
	var ret ffi.Arg
	isKeyPressed.Call(&ret, &key)
	return ret.Bool()
}

// IsKeyPressedRepeat - Check if a key has been pressed again (Only PLATFORM_DESKTOP)
func IsKeyPressedRepeat(key int32) bool {
	var ret ffi.Arg
	isKeyPressedRepeat.Call(&ret, &key)
	return ret.Bool()
}

// IsKeyDown - Check if a key is being pressed
func IsKeyDown(key int32) bool {
	var ret ffi.Arg
	isKeyDown.Call(&ret, &key)
	return ret.Bool()
}

// IsKeyReleased - Check if a key has been released once
func IsKeyReleased(key int32) bool {
	var ret ffi.Arg
	isKeyReleased.Call(&ret, &key)
	return ret.Bool()
}

// IsKeyUp - Check if a key is NOT being pressed
func IsKeyUp(key int32) bool {
	var ret ffi.Arg
	isKeyUp.Call(&ret, &key)
	return ret.Bool()
}

// GetKeyPressed - Get key pressed (keycode), call it multiple times for keys queued, returns 0 when the queue is empty
func GetKeyPressed() int32 {
	var ret ffi.Arg
	getKeyPressed.Call(&ret)
	return int32(ret)
}

// GetCharPressed - Get char pressed (unicode), call it multiple times for chars queued, returns 0 when the queue is empty
func GetCharPressed() int32 {
	var ret ffi.Arg
	getCharPressed.Call(&ret)
	return int32(ret)
}

// GetKeyName - Get name of a QWERTY key on the current keyboard layout
// (eg returns string 'q' for KEY_A on an AZERTY keyboard)
func GetKeyName(key int32) string {
	var ret *byte
	getKeyName.Call(&ret, &key)
	return convert.ToString(ret)
}

// SetExitKey - Set a custom key to exit program (default is ESC)
func SetExitKey(key int32) {
	setExitKey.Call(nil, &key)
}

// IsGamepadAvailable - Check if a gamepad is available
func IsGamepadAvailable(gamepad int32) bool {
	var ret ffi.Arg
	isGamepadAvailable.Call(&ret, &gamepad)
	return ret.Bool()
}

// GetGamepadName - Get gamepad internal name id
func GetGamepadName(gamepad int32) string {
	var ret *byte
	getGamepadName.Call(&ret, &gamepad)
	return convert.ToString(ret)
}

// IsGamepadButtonPressed - Check if a gamepad button has been pressed once
func IsGamepadButtonPressed(gamepad int32, button int32) bool {
	var ret ffi.Arg
	isGamepadButtonPressed.Call(&ret, &gamepad, &button)
	return ret.Bool()
}

// IsGamepadButtonDown - Check if a gamepad button is being pressed
func IsGamepadButtonDown(gamepad int32, button int32) bool {
	var ret ffi.Arg
	isGamepadButtonDown.Call(&ret, &gamepad, &button)
	return ret.Bool()
}

// IsGamepadButtonReleased - Check if a gamepad button has been released once
func IsGamepadButtonReleased(gamepad int32, button int32) bool {
	var ret ffi.Arg
	isGamepadButtonReleased.Call(&ret, &gamepad, &button)
	return ret.Bool()
}

// IsGamepadButtonUp - Check if a gamepad button is NOT being pressed
func IsGamepadButtonUp(gamepad int32, button int32) bool {
	var ret ffi.Arg
	isGamepadButtonUp.Call(&ret, &gamepad, &button)
	return ret.Bool()
}

// GetGamepadButtonPressed - Get the last gamepad button pressed
func GetGamepadButtonPressed() int32 {
	var ret ffi.Arg
	getGamepadButtonPressed.Call(&ret)
	return int32(ret)
}

// GetGamepadAxisCount - Get gamepad axis count for a gamepad
func GetGamepadAxisCount(gamepad int32) int32 {
	var ret ffi.Arg
	getGamepadAxisCount.Call(&ret, &gamepad)
	return int32(ret)
}

// GetGamepadAxisMovement - Get axis movement value for a gamepad axis
func GetGamepadAxisMovement(gamepad int32, axis int32) float32 {
	var ret float32
	getGamepadAxisMovement.Call(&ret, &gamepad, &axis)
	return ret
}

// SetGamepadMappings - Set internal gamepad mappings (SDL_GameControllerDB)
func SetGamepadMappings(mappings string) int32 {
	var ret ffi.Arg
	mappingsPtr := convert.ToBytePtr(mappings)
	setGamepadMappings.Call(&ret, &mappingsPtr)
	return int32(ret)
}

// SetGamepadVibration - Set gamepad vibration for both motors (duration in seconds)
func SetGamepadVibration(gamepad int32, leftMotor, rightMotor, duration float32) {
	setGamepadVibration.Call(nil, &gamepad, &leftMotor, &rightMotor, &duration)
}

// IsMouseButtonPressed - Check if a mouse button has been pressed once
func IsMouseButtonPressed(button MouseButton) bool {
	var ret ffi.Arg
	isMouseButtonPressed.Call(&ret, &button)
	return ret.Bool()
}

// IsMouseButtonDown - Check if a mouse button is being pressed
func IsMouseButtonDown(button MouseButton) bool {
	var ret ffi.Arg
	isMouseButtonDown.Call(&ret, &button)
	return ret.Bool()
}

// IsMouseButtonReleased - Check if a mouse button has been released once
func IsMouseButtonReleased(button MouseButton) bool {
	var ret ffi.Arg
	isMouseButtonReleased.Call(&ret, &button)
	return ret.Bool()
}

// IsMouseButtonUp - Check if a mouse button is NOT being pressed
func IsMouseButtonUp(button MouseButton) bool {
	var ret ffi.Arg
	isMouseButtonUp.Call(&ret, &button)
	return ret.Bool()
}

// GetMouseX - Get mouse position X
func GetMouseX() int32 {
	var ret ffi.Arg
	getMouseX.Call(&ret)
	return int32(ret)
}

// GetMouseY - Get mouse position Y
func GetMouseY() int32 {
	var ret ffi.Arg
	getMouseY.Call(&ret)
	return int32(ret)
}

// GetMousePosition - Get mouse position XY
func GetMousePosition() Vector2 {
	var ret Vector2
	getMousePosition.Call(&ret)
	return ret
}

// GetMouseDelta - Get mouse delta between frames
func GetMouseDelta() Vector2 {
	var ret Vector2
	getMouseDelta.Call(&ret)
	return ret
}

// SetMousePosition - Set mouse position XY
func SetMousePosition(x int32, y int32) {
	setMousePosition.Call(nil, &x, &y)
}

// SetMouseOffset - Set mouse offset
func SetMouseOffset(offsetX int32, offsetY int32) {
	setMouseOffset.Call(nil, &offsetX, &offsetY)
}

// SetMouseScale - Set mouse scaling
func SetMouseScale(scaleX float32, scaleY float32) {
	setMouseScale.Call(nil, &scaleX, &scaleY)
}

// GetMouseWheelMove - Get mouse wheel movement for X or Y, whichever is larger
func GetMouseWheelMove() float32 {
	var ret float32
	getMouseWheelMove.Call(&ret)
	return ret
}

// GetMouseWheelMoveV - Get mouse wheel movement for both X and Y
func GetMouseWheelMoveV() Vector2 {
	var ret Vector2
	getMouseWheelMoveV.Call(&ret)
	return ret
}

// SetMouseCursor - Set mouse cursor
func SetMouseCursor(cursor int32) {
	setMouseCursor.Call(nil, &cursor)
}

// GetTouchX - Get touch position X for touch point 0 (relative to screen size)
func GetTouchX() int32 {
	var ret ffi.Arg
	getTouchX.Call(&ret)
	return int32(ret)
}

// GetTouchY - Get touch position Y for touch point 0 (relative to screen size)
func GetTouchY() int32 {
	var ret ffi.Arg
	getTouchY.Call(&ret)
	return int32(ret)
}

// GetTouchPosition - Get touch position XY for a touch point index (relative to screen size)
func GetTouchPosition(index int32) Vector2 {
	var ret Vector2
	getTouchPosition.Call(&ret, &index)
	return ret
}

// GetTouchPointId - Get touch point identifier for given index
func GetTouchPointId(index int32) int32 {
	var ret ffi.Arg
	getTouchPointId.Call(&ret, &index)
	return int32(ret)
}

// GetTouchPointCount - Get number of touch points
func GetTouchPointCount() int32 {
	var ret ffi.Arg
	getTouchPointCount.Call(&ret)
	return int32(ret)
}

// SetGesturesEnabled - Enable a set of gestures using flags
func SetGesturesEnabled(flags uint32) {
	setGesturesEnabled.Call(nil, &flags)
}

// IsGestureDetected - Check if a gesture have been detected
func IsGestureDetected(gesture Gestures) bool {
	var ret ffi.Arg
	isGestureDetected.Call(&ret, &gesture)
	return ret.Bool()
}

// GetGestureDetected - Get latest detected gesture
func GetGestureDetected() Gestures {
	var ret ffi.Arg
	getGestureDetected.Call(&ret)
	return Gestures(ret)
}

// GetGestureHoldDuration - Get gesture hold time in milliseconds
func GetGestureHoldDuration() float32 {
	var ret float32
	getGestureHoldDuration.Call(&ret)
	return ret
}

// GetGestureDragVector - Get gesture drag vector
func GetGestureDragVector() Vector2 {
	var ret Vector2
	getGestureDragVector.Call(&ret)
	return ret
}

// GetGestureDragAngle - Get gesture drag angle
func GetGestureDragAngle() float32 {
	var ret float32
	getGestureDragAngle.Call(&ret)
	return ret
}

// GetGesturePinchVector - Get gesture pinch delta
func GetGesturePinchVector() Vector2 {
	var ret Vector2
	getGesturePinchVector.Call(&ret)
	return ret
}

// GetGesturePinchAngle - Get gesture pinch angle
func GetGesturePinchAngle() float32 {
	var ret float32
	getGesturePinchAngle.Call(&ret)
	return ret
}
