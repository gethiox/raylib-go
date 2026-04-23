//go:build !cgo
// +build !cgo

package rl

import "github.com/jupiterrider/ffi"

var typeTexture2D = ffi.NewType(&ffi.TypeUint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32)
var typeImage = ffi.NewType(&ffi.TypePointer, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32)
var typeVector2 = ffi.NewType(&ffi.TypeFloat, &ffi.TypeFloat)
var typeVector3 = ffi.NewType(&ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat)
var typeColor = ffi.NewType(&ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypeUint8)
var typeCamera2D = ffi.NewType(&typeVector2, &typeVector2, &ffi.TypeFloat, &ffi.TypeFloat)
var typeCamera3D = ffi.NewType(&typeVector3, &typeVector3, &typeVector3, &ffi.TypeFloat, &ffi.TypeSint32)
var typeRenderTexture2D = ffi.NewType(&ffi.TypeUint32, &typeTexture2D, &typeTexture2D)
var typeShader = ffi.NewType(&ffi.TypeUint32, &ffi.TypePointer)
var typeMatrix = ffi.NewType(
	&ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat,
	&ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat,
	&ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat,
	&ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat,
)
var typeVrStereoConfig = ffi.NewType(
	&typeMatrix, &typeMatrix,
	&typeMatrix, &typeMatrix,
	&ffi.TypeFloat, &ffi.TypeFloat,
	&ffi.TypeFloat, &ffi.TypeFloat,
	&ffi.TypeFloat, &ffi.TypeFloat,
	&ffi.TypeFloat, &ffi.TypeFloat,
	&ffi.TypeFloat, &ffi.TypeFloat,
	&ffi.TypeFloat, &ffi.TypeFloat,
)
var typeVrDeviceInfo = ffi.NewType(
	&ffi.TypeSint32, &ffi.TypeSint32,
	&ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat,
	&ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat,
	&ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat,
)
var typeRay = ffi.NewType(&typeVector3, &typeVector3)
var typeFilePathList = ffi.NewType(&ffi.TypeUint32, &ffi.TypePointer)
var typeAutomationEvent = ffi.NewType(
	&ffi.TypeUint32, &ffi.TypeUint32,
	&ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32,
)
var typeAutomationEventList = ffi.NewType(&ffi.TypeUint32, &ffi.TypeUint32, &ffi.TypePointer)
var typeRectangle = ffi.NewType(&ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat)
