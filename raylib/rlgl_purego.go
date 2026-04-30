//go:build !cgo
// +build !cgo

package rl

import (
	"unsafe"

	"github.com/gen2brain/raylib-go/raylib/internal/convert"
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

	// General render state

	rlEnableColorBlend        = dll.MustPrep("rlEnableColorBlend", &ffi.TypeVoid)
	rlDisableColorBlend       = dll.MustPrep("rlDisableColorBlend", &ffi.TypeVoid)
	rlEnableDepthTest         = dll.MustPrep("rlEnableDepthTest", &ffi.TypeVoid)
	rlDisableDepthTest        = dll.MustPrep("rlDisableDepthTest", &ffi.TypeVoid)
	rlEnableDepthMask         = dll.MustPrep("rlEnableDepthMask", &ffi.TypeVoid)
	rlDisableDepthMask        = dll.MustPrep("rlDisableDepthMask", &ffi.TypeVoid)
	rlEnableBackfaceCulling   = dll.MustPrep("rlEnableBackfaceCulling", &ffi.TypeVoid)
	rlDisableBackfaceCulling  = dll.MustPrep("rlDisableBackfaceCulling", &ffi.TypeVoid)
	rlColorMask               = dll.MustPrep("rlColorMask", &ffi.TypeVoid, &ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypeUint8)
	rlSetCullFace             = dll.MustPrep("rlSetCullFace", &ffi.TypeVoid, &ffi.TypeSint32)
	rlEnableScissorTest       = dll.MustPrep("rlEnableScissorTest", &ffi.TypeVoid)
	rlDisableScissorTest      = dll.MustPrep("rlDisableScissorTest", &ffi.TypeVoid)
	rlScissor                 = dll.MustPrep("rlScissor", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32)
	rlEnablePointMode         = dll.MustPrep("rlEnablePointMode", &ffi.TypeVoid)
	rlDisablePointMode        = dll.MustPrep("rlDisablePointMode", &ffi.TypeVoid)
	rlSetPointSize            = dll.MustPrep("rlSetPointSize", &ffi.TypeVoid, &ffi.TypeFloat)
	rlGetPointSize            = dll.MustPrep("rlGetPointSize", &ffi.TypeFloat)
	rlEnableWireMode          = dll.MustPrep("rlEnableWireMode", &ffi.TypeVoid)
	rlDisableWireMode         = dll.MustPrep("rlDisableWireMode", &ffi.TypeVoid)
	rlSetLineWidth            = dll.MustPrep("rlSetLineWidth", &ffi.TypeVoid, &ffi.TypeFloat)
	rlGetLineWidth            = dll.MustPrep("rlGetLineWidth", &ffi.TypeFloat)
	rlEnableSmoothLines       = dll.MustPrep("rlEnableSmoothLines", &ffi.TypeVoid)
	rlDisableSmoothLines      = dll.MustPrep("rlDisableSmoothLines", &ffi.TypeVoid)
	rlEnableStereoRender      = dll.MustPrep("rlEnableStereoRender", &ffi.TypeVoid)
	rlDisableStereoRender     = dll.MustPrep("rlDisableStereoRender", &ffi.TypeVoid)
	rlIsStereoRenderEnabled   = dll.MustPrep("rlIsStereoRenderEnabled", &ffi.TypeUint8)
	rlClearColor              = dll.MustPrep("rlClearColor", &ffi.TypeVoid, &ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypeUint8, &ffi.TypeUint8)
	rlClearScreenBuffers      = dll.MustPrep("rlClearScreenBuffers", &ffi.TypeVoid)
	rlCheckErrors             = dll.MustPrep("rlCheckErrors", &ffi.TypeVoid)
	rlSetBlendMode            = dll.MustPrep("rlSetBlendMode", &ffi.TypeVoid, &ffi.TypeSint32)
	rlSetBlendFactors         = dll.MustPrep("rlSetBlendFactors", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32)
	rlSetBlendFactorsSeparate = dll.MustPrep("rlSetBlendFactorsSeparate", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32, &ffi.TypeSint32)

	// rlgl initialization functions

	rlglInit               = dll.MustPrep("rlglInit", &ffi.TypeVoid, &ffi.TypeSint32, &ffi.TypeSint32)
	rlglClose              = dll.MustPrep("rlglClose", &ffi.TypeVoid)
	rlGetProcAddress       = dll.MustPrep("rlGetProcAddress", &ffi.TypePointer, &ffi.TypePointer)
	rlGetVersion           = dll.MustPrep("rlGetVersion", &ffi.TypeSint32)
	rlSetFramebufferWidth  = dll.MustPrep("rlSetFramebufferWidth", &ffi.TypeVoid, &ffi.TypeSint32)
	rlGetFramebufferWidth  = dll.MustPrep("rlGetFramebufferWidth", &ffi.TypeSint32)
	rlSetFramebufferHeight = dll.MustPrep("rlSetFramebufferHeight", &ffi.TypeVoid, &ffi.TypeSint32)
	rlGetFramebufferHeight = dll.MustPrep("rlGetFramebufferHeight", &ffi.TypeSint32)
	rlGetTextureIdDefault  = dll.MustPrep("rlGetTextureIdDefault", &ffi.TypeUint32)
	rlGetShaderIdDefault   = dll.MustPrep("rlGetShaderIdDefault", &ffi.TypeUint32)
	rlGetShaderLocsDefault = dll.MustPrep("rlGetShaderLocsDefault", &ffi.TypePointer)

	// Render batch management

	rlLoadRenderBatch       = dll.MustPrep("rlLoadRenderBatch", &typeRenderBatch, &ffi.TypeSint32, &ffi.TypeSint32)
	rlUnloadRenderBatch     = dll.MustPrep("rlUnloadRenderBatch", &ffi.TypeVoid, &typeRenderBatch)
	rlDrawRenderBatch       = dll.MustPrep("rlDrawRenderBatch", &ffi.TypeVoid, &ffi.TypePointer)
	rlSetRenderBatchActive  = dll.MustPrep("rlSetRenderBatchActive", &ffi.TypeVoid, &ffi.TypePointer)
	rlDrawRenderBatchActive = dll.MustPrep("rlDrawRenderBatchActive", &ffi.TypeVoid)
	rlCheckRenderBatchLimit = dll.MustPrep("rlCheckRenderBatchLimit", &ffi.TypeUint8, &ffi.TypeSint32)
	rlSetTexture            = dll.MustPrep("rlSetTexture", &ffi.TypeVoid, &ffi.TypeUint32)
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

// EnableColorBlend - Enable color blending
func EnableColorBlend() {
	rlEnableColorBlend.Call(nil)
}

// DisableColorBlend - Disable color blending
func DisableColorBlend() {
	rlDisableColorBlend.Call(nil)
}

// EnableDepthTest - Enable depth test
func EnableDepthTest() {
	rlEnableDepthTest.Call(nil)
}

// DisableDepthTest - Disable depth test
func DisableDepthTest() {
	rlDisableDepthTest.Call(nil)
}

// EnableDepthMask - Enable depth write
func EnableDepthMask() {
	rlEnableDepthMask.Call(nil)
}

// DisableDepthMask - Disable depth write
func DisableDepthMask() {
	rlDisableDepthMask.Call(nil)
}

// EnableBackfaceCulling - Enable backface culling
func EnableBackfaceCulling() {
	rlEnableBackfaceCulling.Call(nil)
}

// DisableBackfaceCulling - Disable backface culling
func DisableBackfaceCulling() {
	rlDisableBackfaceCulling.Call(nil)
}

// ColorMask - Color mask control
func ColorMask(r, g, b, a bool) {
	rlColorMask.Call(nil, &r, &g, &b, &a)
}

// SetCullFace - Set face culling mode
func SetCullFace(mode int32) {
	rlSetCullFace.Call(nil, &mode)
}

// EnableScissorTest - Enable scissor test
func EnableScissorTest() {
	rlEnableScissorTest.Call(nil)
}

// DisableScissorTest - Disable scissor test
func DisableScissorTest() {
	rlDisableScissorTest.Call(nil)
}

// Scissor - Scissor test
func Scissor(x int32, y int32, width int32, height int32) {
	rlScissor.Call(nil, &x, &y, &width, &height)
}

// EnablePointMode - Enable point mode
func EnablePointMode() {
	rlEnablePointMode.Call(nil)
}

// DisablePointMode - Disable point mode
func DisablePointMode() {
	rlDisablePointMode.Call(nil)
}

// SetPointSize - Set the point drawing size
func SetPointSize(size float32) {
	rlSetPointSize.Call(nil, &size)
}

// GetPointSize - Get the point drawing size
func GetPointSize() float32 {
	var ret float32
	rlGetPointSize.Call(&ret)
	return ret
}

// EnableWireMode - Enable wire mode
func EnableWireMode() {
	rlEnableWireMode.Call(nil)
}

// DisableWireMode - Disable wire mode
func DisableWireMode() {
	rlDisableWireMode.Call(nil)
}

// SetLineWidth - Set the line drawing width
func SetLineWidth(width float32) {
	rlSetLineWidth.Call(nil, &width)
}

// GetLineWidth - Get the line drawing width
func GetLineWidth() float32 {
	var ret float32
	rlGetLineWidth.Call(&ret)
	return ret
}

// EnableSmoothLines - Enable line aliasing
func EnableSmoothLines() {
	rlEnableSmoothLines.Call(nil)
}

// DisableSmoothLines - Disable line aliasing
func DisableSmoothLines() {
	rlDisableSmoothLines.Call(nil)
}

// EnableStereoRender - Enable stereo rendering
func EnableStereoRender() {
	rlEnableStereoRender.Call(nil)
}

// DisableStereoRender - Disable stereo rendering
func DisableStereoRender() {
	rlDisableStereoRender.Call(nil)
}

// IsStereoRenderEnabled - Check if stereo render is enabled
func IsStereoRenderEnabled() bool {
	var ret ffi.Arg
	rlIsStereoRenderEnabled.Call(&ret)
	return ret.Bool()
}

// ClearColor - Clear color buffer with color
func ClearColor(r, g, b, a uint8) {
	rlClearColor.Call(nil, &r, &g, &b, &a)
}

// ClearScreenBuffers - Clear used screen buffers (color and depth)
func ClearScreenBuffers() {
	rlClearScreenBuffers.Call(nil)
}

// CheckErrors - Check and log OpenGL error codes
func CheckErrors() {
	rlCheckErrors.Call(nil)
}

// SetBlendMode - Set blending mode
func SetBlendMode(mode BlendMode) {
	rlSetBlendMode.Call(nil, &mode)
}

// SetBlendFactors - Set blending mode factor and equation (using OpenGL factors)
func SetBlendFactors(glSrcFactor, glDstFactor, glEquation int32) {
	rlSetBlendFactors.Call(nil, &glSrcFactor, &glDstFactor, &glEquation)
}

// SetBlendFactorsSeparate - Set blending mode factors and equations separately (using OpenGL factors)
func SetBlendFactorsSeparate(glSrcRGB, glDstRGB, glSrcAlpha, glDstAlpha, glEqRGB, glEqAlpha int32) {
	rlSetBlendFactorsSeparate.Call(nil, &glSrcRGB, &glDstRGB, &glSrcAlpha, &glDstAlpha, &glEqRGB, &glEqAlpha)
}

// GlInit - Initialize rlgl (buffers, shaders, textures, states)
func GlInit(width int32, height int32) {
	rlglInit.Call(nil, &width, &height)
}

// GlClose - De-inititialize rlgl (buffers, shaders, textures)
func GlClose() {
	rlglClose.Call(nil)
}

// GetProcAddress - Get OpenGL procedure address
func GetProcAddress(procName string) unsafe.Pointer {
	procNamePtr := convert.ToBytePtr(procName)
	var ret unsafe.Pointer
	rlGetProcAddress.Call(&ret, &procNamePtr)
	return ret
}

// GetVersion - Get current OpenGL version
func GetVersion() int32 {
	var ret ffi.Arg
	rlGetVersion.Call(&ret)
	return int32(ret)
}

// SetFramebufferWidth - Set current framebuffer width
func SetFramebufferWidth(width int32) {
	rlSetFramebufferWidth.Call(nil, &width)
}

// GetFramebufferWidth - Get default framebuffer width
func GetFramebufferWidth() int32 {
	var ret ffi.Arg
	rlGetFramebufferWidth.Call(&ret)
	return int32(ret)
}

// SetFramebufferHeight - Set current framebuffer height
func SetFramebufferHeight(height int32) {
	rlSetFramebufferHeight.Call(nil, &height)
}

// GetFramebufferHeight - Get default framebuffer height
func GetFramebufferHeight() int32 {
	var ret ffi.Arg
	rlGetFramebufferHeight.Call(&ret)
	return int32(ret)
}

// GetTextureIdDefault - Get default texture id
func GetTextureIdDefault() uint32 {
	var ret ffi.Arg
	rlGetTextureIdDefault.Call(&ret)
	return uint32(ret)
}

// GetShaderIdDefault - Get default shader id
func GetShaderIdDefault() uint32 {
	var ret ffi.Arg
	rlGetShaderIdDefault.Call(&ret)
	return uint32(ret)
}

// GetShaderLocsDefault - Get default shader locations
func GetShaderLocsDefault() []int32 {
	var ret *int32
	rlGetShaderLocsDefault.Call(&ret)
	// the default value of RL_MAX_SHADER_LOCATIONS is 32
	return unsafe.Slice(ret, 32)
}

// LoadRenderBatch - Load a render batch system
func LoadRenderBatch(numBuffers int32, bufferElements int32) RenderBatch {
	var ret RenderBatch
	rlLoadRenderBatch.Call(&ret, &numBuffers, &bufferElements)
	return ret
}

// UnloadRenderBatch - Unload render batch system
func UnloadRenderBatch(batch RenderBatch) {
	rlUnloadRenderBatch.Call(nil, &batch)
}

// DrawRenderBatch - Draw render batch data (Update->Draw->Reset)
func DrawRenderBatch(batch *RenderBatch) {
	rlDrawRenderBatch.Call(nil, &batch)
}

// rlSetRenderBatchActive - Set the active render batch for rlgl (NULL for default internal)
func SetRenderBatchActive(batch *RenderBatch) {
	rlSetRenderBatchActive.Call(nil, &batch)
}

// DrawRenderBatchActive - Update and draw internal render batch
func DrawRenderBatchActive() {
	rlDrawRenderBatchActive.Call(nil)
}

// CheckRenderBatchLimit - Check internal buffer overflow for a given number of vertex
func CheckRenderBatchLimit(vCount int32) bool {
	var ret ffi.Arg
	rlCheckRenderBatchLimit.Call(&ret, &vCount)
	return ret.Bool()
}

// SetTexture - Set current texture for render batch and check buffers limits
func SetTexture(id uint32) {
	rlSetTexture.Call(nil, &id)
}
