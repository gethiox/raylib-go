//go:build !cgo
// +build !cgo

package rl

import (
	"unsafe"

	"github.com/jupiterrider/ffi"
)

var (
	// Matrix operations

	rlMatrixMode          = dll.MustPrep("rlMatrixMode", &ffi.TypeVoid, &ffi.TypeSint32)
	rlPushMatrix          = dll.MustPrep("rlPushMatrix", &ffi.TypeVoid)
	rlPopMatrix           = dll.MustPrep("rlPopMatrix", &ffi.TypeVoid)
	rlLoadIdentity        = dll.MustPrep("rlLoadIdentity", &ffi.TypeVoid)
	rlTranslatef          = dll.MustPrep("rlTranslatef", &ffi.TypeVoid, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat)
	rlRotatef             = dll.MustPrep("rlRotatef", &ffi.TypeVoid, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat)
	rlScalef              = dll.MustPrep("rlScalef", &ffi.TypeVoid, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat)
	rlMultMatrixf         = dll.MustPrep("rlMultMatrixf", &ffi.TypeVoid, &ffi.TypePointer)
	rlFrustum             = dll.MustPrep("rlFrustum", &ffi.TypeVoid, &ffi.TypeDouble, &ffi.TypeDouble, &ffi.TypeDouble, &ffi.TypeDouble, &ffi.TypeDouble, &ffi.TypeDouble)
	rlOrtho               = dll.MustPrep("rlOrtho", &ffi.TypeVoid, &ffi.TypeDouble, &ffi.TypeDouble, &ffi.TypeDouble, &ffi.TypeDouble, &ffi.TypeDouble, &ffi.TypeDouble)
	rlViewport            = dll.MustPrep("rlViewport", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32)
	rlSetClipPlanes       = dll.MustPrep("rlSetClipPlanes", &ffi.TypeVoid, &ffi.TypeDouble, &ffi.TypeDouble)
	rlGetCullDistanceNear = dll.MustPrep("rlGetCullDistanceNear", &ffi.TypeDouble)
	rlGetCullDistanceFar  = dll.MustPrep("rlGetCullDistanceFar", &ffi.TypeDouble)

	// Vertex level operations

	rlBegin      = dll.MustPrep("rlBegin", &ffi.TypeVoid, &ffi.TypeSint32)
	rlEnd        = dll.MustPrep("rlEnd", &ffi.TypeVoid)
	rlVertex2i   = dll.MustPrep("rlVertex2i", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32)
	rlVertex2f   = dll.MustPrep("rlVertex2f", &ffi.TypeVoid, &ffi.TypeFloat, &ffi.TypeFloat)
	rlVertex3f   = dll.MustPrep("rlVertex3f", &ffi.TypeVoid, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat)
	rlTexCoord2f = dll.MustPrep("rlTexCoord2f", &ffi.TypeVoid, &ffi.TypeFloat, &ffi.TypeFloat)
	rlNormal3f   = dll.MustPrep("rlNormal3f", &ffi.TypeVoid, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat)
	rlColor4ub   = dll.MustPrep("rlColor4ub", &ffi.TypeVoid, &ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypeUint8)
	rlColor3f    = dll.MustPrep("rlColor3f", &ffi.TypeVoid, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat)
	rlColor4f    = dll.MustPrep("rlColor4f", &ffi.TypeVoid, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat, &ffi.TypeFloat)

	// Vertex buffers state

	rlEnableVertexArray          = dll.MustPrep("rlEnableVertexArray", &ffi.TypeUint8, &ffi.TypeUint32)
	rlDisableVertexArray         = dll.MustPrep("rlDisableVertexArray", &ffi.TypeVoid)
	rlEnableVertexBuffer         = dll.MustPrep("rlEnableVertexBuffer", &ffi.TypeVoid, &ffi.TypeUint32)
	rlDisableVertexBuffer        = dll.MustPrep("rlDisableVertexBuffer", &ffi.TypeVoid)
	rlEnableVertexBufferElement  = dll.MustPrep("rlEnableVertexBufferElement", &ffi.TypeVoid, &ffi.TypeUint32)
	rlDisableVertexBufferElement = dll.MustPrep("rlDisableVertexBufferElement", &ffi.TypeVoid)
	rlEnableVertexAttribute      = dll.MustPrep("rlEnableVertexAttribute", &ffi.TypeVoid, &ffi.TypeUint32)
	rlDisableVertexAttribute     = dll.MustPrep("rlDisableVertexAttribute", &ffi.TypeVoid, &ffi.TypeUint32)
	rlEnableStatePointer         = dll.MustPrep("rlEnableStatePointer", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypePointer)
	rlDisableStatePointer        = dll.MustPrep("rlDisableStatePointer", &ffi.TypeVoid, &ffi.TypeSint32)

	// Textures state

	rlActiveTextureSlot     = dll.MustPrep("rlActiveTextureSlot", &ffi.TypeVoid, &ffi.TypeSint32)
	rlEnableTexture         = dll.MustPrep("rlEnableTexture", &ffi.TypeVoid, &ffi.TypeUint32)
	rlDisableTexture        = dll.MustPrep("rlDisableTexture", &ffi.TypeVoid)
	rlEnableTextureCubemap  = dll.MustPrep("rlEnableTextureCubemap", &ffi.TypeVoid, &ffi.TypeUint32)
	rlDisableTextureCubemap = dll.MustPrep("rlDisableTextureCubemap", &ffi.TypeVoid)
	rlTextureParameters     = dll.MustPrep("rlTextureParameters", &ffi.TypeVoid, &ffi.TypeUint32, &ffi.TypeSint32, &ffi.TypeSint32)
	rlCubemapParameters     = dll.MustPrep("rlCubemapParameters", &ffi.TypeVoid, &ffi.TypeUint32, &ffi.TypeSint32, &ffi.TypeSint32)

	// Shader state

	rlEnableShader  = dll.MustPrep("rlEnableShader", &ffi.TypeVoid, &ffi.TypeUint32)
	rlDisableShader = dll.MustPrep("rlDisableShader", &ffi.TypeVoid)

	// Framebuffer state

	rlEnableFramebuffer    = dll.MustPrep("rlEnableFramebuffer", &ffi.TypeVoid, &ffi.TypeUint32)
	rlDisableFramebuffer   = dll.MustPrep("rlDisableFramebuffer", &ffi.TypeVoid)
	rlGetActiveFramebuffer = dll.MustPrep("rlGetActiveFramebuffer", &ffi.TypeUint32)
	rlActiveDrawBuffers    = dll.MustPrep("rlActiveDrawBuffers", &ffi.TypeVoid, &ffi.TypeSint32)
	rlBlitFramebuffer      = dll.MustPrep("rlBlitFramebuffer", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32)
	rlBindFramebuffer      = dll.MustPrep("rlBindFramebuffer", &ffi.TypeVoid, &ffi.TypeUint32, &ffi.TypeUint32)
)

// SetVertexAttribute - Set vertex attribute data configuration
func SetVertexAttribute(index uint32, compSize int32, attrType int32, normalized bool, stride int32, offset int32) {
}

// MatrixMode - Choose the current matrix to be transformed
func MatrixMode(mode int32) {
	rlMatrixMode.Call(nil, &mode)
}

// PushMatrix - Push the current matrix to stack
func PushMatrix() {
	rlPushMatrix.Call(nil)
}

// PopMatrix - Pop lattest inserted matrix from stack
func PopMatrix() {
	rlPopMatrix.Call(nil)
}

// LoadIdentity - Reset current matrix to identity matrix
func LoadIdentity() {
	rlLoadIdentity.Call(nil)
}

// Translatef - Multiply the current matrix by a translation matrix
func Translatef(x float32, y float32, z float32) {
	rlTranslatef.Call(nil, &x, &y, &z)
}

// Rotatef - Multiply the current matrix by a rotation matrix
func Rotatef(angle float32, x float32, y float32, z float32) {
	rlRotatef.Call(nil, &angle, &x, &y, &z)
}

// Scalef - Multiply the current matrix by a scaling matrix
func Scalef(x float32, y float32, z float32) {
	rlScalef.Call(nil, &x, &y, &z)
}

// MultMatrix - Multiply the current matrix by another matrix
func MultMatrix(m Matrix) {
	f := unsafe.SliceData(MatrixToFloat(m))
	rlMultMatrixf.Call(nil, &f)
}

// Frustum .
func Frustum(left float64, right float64, bottom float64, top float64, znear float64, zfar float64) {
	rlFrustum.Call(nil, &left, &right, &bottom, &top, &znear, &zfar)
}

// Ortho .
func Ortho(left float64, right float64, bottom float64, top float64, znear float64, zfar float64) {
	rlOrtho.Call(nil, &left, &right, &bottom, &top, &znear, &zfar)
}

// Viewport - Set the viewport area
func Viewport(x int32, y int32, width int32, height int32) {
	rlViewport.Call(nil, &x, &y, &width, &height)
}

// SetClipPlanes - Set clip planes distances
func SetClipPlanes(nearPlane, farPlane float64) {
	rlSetClipPlanes.Call(nil, &nearPlane, &farPlane)
}

// GetCullDistanceNear - Get cull plane distance near
func GetCullDistanceNear() float64 {
	var ret float64
	rlGetCullDistanceNear.Call(&ret)
	return ret
}

// GetCullDistanceFar - Get cull plane distance far
func GetCullDistanceFar() float64 {
	var ret float64
	rlGetCullDistanceFar.Call(&ret)
	return ret
}

// Begin - Initialize drawing mode (how to organize vertex)
func Begin(mode int32) {
	rlBegin.Call(nil, &mode)
}

// End - Finish vertex providing
func End() {
	rlEnd.Call(nil)
}

// Vertex2i - Define one vertex (position) - 2 int
func Vertex2i(x int32, y int32) {
	rlVertex2i.Call(nil, &x, &y)
}

// Vertex2f - Define one vertex (position) - 2 float
func Vertex2f(x float32, y float32) {
	rlVertex2f.Call(nil, &x, &y)
}

// Vertex3f - Define one vertex (position) - 3 float
func Vertex3f(x float32, y float32, z float32) {
	rlVertex3f.Call(nil, &x, &y, &z)
}

// TexCoord2f - Define one vertex (texture coordinate) - 2 float
func TexCoord2f(x float32, y float32) {
	rlTexCoord2f.Call(nil, &x, &y)
}

// Normal3f - Define one vertex (normal) - 3 float
func Normal3f(x float32, y float32, z float32) {
	rlNormal3f.Call(nil, &x, &y, &z)
}

// Color4ub - Define one vertex (color) - 4 byte
func Color4ub(r uint8, g uint8, b uint8, a uint8) {
	rlColor4ub.Call(nil, &r, &g, &b, &a)
}

// Color3f - Define one vertex (color) - 3 float
func Color3f(x float32, y float32, z float32) {
	rlColor3f.Call(nil, &x, &y, &z)
}

// Color4f - Define one vertex (color) - 4 float
func Color4f(x float32, y float32, z float32, w float32) {
	rlColor4f.Call(nil, &x, &y, &z, &w)
}

// EnableVertexArray - Enable vertex array (VAO, if supported)
func EnableVertexArray(vaoId uint32) bool {
	var ret ffi.Arg
	rlEnableVertexArray.Call(&ret, &vaoId)
	return ret.Bool()
}

// DisableVertexArray - Disable vertex array (VAO, if supported)
func DisableVertexArray() {
	rlDisableVertexArray.Call(nil)
}

// EnableVertexBuffer - Enable vertex buffer (VBO)
func EnableVertexBuffer(id uint32) {
	rlEnableVertexBuffer.Call(nil, &id)
}

// DisableVertexBuffer - Disable vertex buffer (VBO)
func DisableVertexBuffer() {
	rlDisableVertexBuffer.Call(nil)
}

// EnableVertexBufferElement - Enable vertex buffer element (VBO element)
func EnableVertexBufferElement(id uint32) {
	rlEnableVertexBufferElement.Call(nil, &id)
}

// DisableVertexBufferElement - Disable vertex buffer element (VBO element)
func DisableVertexBufferElement() {
	rlDisableVertexBufferElement.Call(nil)
}

// EnableVertexAttribute - Enable vertex attribute index
func EnableVertexAttribute(index uint32) {
	rlEnableVertexAttribute.Call(nil, &index)
}

// DisableVertexAttribute - Disable vertex attribute index
func DisableVertexAttribute(index uint32) {
	rlDisableVertexAttribute.Call(nil, &index)
}

// EnableStatePointer - Enable attribute state pointer
func EnableStatePointer(vertexAttribType int32, buffer unsafe.Pointer) {
	rlEnableStatePointer.Call(nil, &vertexAttribType, &buffer)
}

// DisableStatePointer - Disable attribute state pointer
func DisableStatePointer(vertexAttribType int32) {
	rlDisableStatePointer.Call(nil, &vertexAttribType)
}

// ActiveTextureSlot - Select and active a texture slot
func ActiveTextureSlot(slot int32) {
	rlActiveTextureSlot.Call(nil, &slot)
}

// EnableTexture - Enable texture
func EnableTexture(id uint32) {
	rlEnableTexture.Call(nil, &id)
}

// DisableTexture - Disable texture
func DisableTexture() {
	rlDisableTexture.Call(nil)
}

// EnableTextureCubemap - Enable texture cubemap
func EnableTextureCubemap(id uint32) {
	rlEnableTextureCubemap.Call(nil, &id)
}

// DisableTextureCubemap - Disable texture cubemap
func DisableTextureCubemap() {
	rlDisableTextureCubemap.Call(nil)
}

// TextureParameters - Set texture parameters (filter, wrap)
func TextureParameters(id uint32, param int32, value int32) {
	rlTextureParameters.Call(nil, &id, &param, &value)
}

// CubemapParameters - Set cubemap parameters (filter, wrap)
func CubemapParameters(id uint32, param int32, value int32) {
	rlCubemapParameters.Call(nil, &id, &param, &value)
}

// EnableShader - Enable shader program
func EnableShader(id uint32) {
	rlEnableShader.Call(nil, &id)
}

// DisableShader - Disable shader program
func DisableShader() {
	rlDisableShader.Call(nil)
}

// EnableFramebuffer - Enable render texture (fbo)
func EnableFramebuffer(id uint32) {
	rlEnableFramebuffer.Call(nil, &id)
}

// DisableFramebuffer - Disable render texture (fbo), return to default framebuffer
func DisableFramebuffer() {
	rlDisableFramebuffer.Call(nil)
}

// GetActiveFramebuffer - Get the currently active render texture (fbo), 0 for default framebuffer
func GetActiveFramebuffer() uint32 {
	var ret ffi.Arg
	rlGetActiveFramebuffer.Call(&ret)
	return uint32(ret)
}

// ActiveDrawBuffers - Activate multiple draw color buffers
func ActiveDrawBuffers(count int32) {
	rlActiveDrawBuffers.Call(nil, &count)
}

// BlitFramebuffer - Blit active framebuffer to main framebuffer
func BlitFramebuffer(srcX, srcY, srcWidth, srcHeight, dstX, dstY, dstWidth, dstHeight, bufferMask int32) {
	rlBlitFramebuffer.Call(nil, &srcX, &srcY, &srcWidth, &srcHeight, &dstX, &dstY, &dstWidth, &dstHeight, &bufferMask)
}

// BindFramebuffer - Bind framebuffer (FBO)
func BindFramebuffer(target, framebuffer uint32) {
	rlBindFramebuffer.Call(nil, &target, &framebuffer)
}
