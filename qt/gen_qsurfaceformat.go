package qt

import (
	"unsafe"
)

type QSurfaceFormat__FormatOption int

const (
	QSurfaceFormat__StereoBuffers       QSurfaceFormat__FormatOption = 1
	QSurfaceFormat__DebugContext        QSurfaceFormat__FormatOption = 2
	QSurfaceFormat__DeprecatedFunctions QSurfaceFormat__FormatOption = 4
	QSurfaceFormat__ResetNotification   QSurfaceFormat__FormatOption = 8
	QSurfaceFormat__ProtectedContent    QSurfaceFormat__FormatOption = 16
)

type QSurfaceFormat__SwapBehavior int

const (
	QSurfaceFormat__DefaultSwapBehavior QSurfaceFormat__SwapBehavior = 0
	QSurfaceFormat__SingleBuffer        QSurfaceFormat__SwapBehavior = 1
	QSurfaceFormat__DoubleBuffer        QSurfaceFormat__SwapBehavior = 2
	QSurfaceFormat__TripleBuffer        QSurfaceFormat__SwapBehavior = 3
)

type QSurfaceFormat__RenderableType int

const (
	QSurfaceFormat__DefaultRenderableType QSurfaceFormat__RenderableType = 0
	QSurfaceFormat__OpenGL                QSurfaceFormat__RenderableType = 1
	QSurfaceFormat__OpenGLES              QSurfaceFormat__RenderableType = 2
	QSurfaceFormat__OpenVG                QSurfaceFormat__RenderableType = 4
)

type QSurfaceFormat__OpenGLContextProfile int

const (
	QSurfaceFormat__NoProfile            QSurfaceFormat__OpenGLContextProfile = 0
	QSurfaceFormat__CoreProfile          QSurfaceFormat__OpenGLContextProfile = 1
	QSurfaceFormat__CompatibilityProfile QSurfaceFormat__OpenGLContextProfile = 2
)

type QSurfaceFormat__ColorSpace int

const (
	QSurfaceFormat__DefaultColorSpace QSurfaceFormat__ColorSpace = 0
	QSurfaceFormat__sRGBColorSpace    QSurfaceFormat__ColorSpace = 1
)

type QSurfaceFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQSurfaceFormat constructs a new QSurfaceFormat object.
func NewQSurfaceFormat() *QSurfaceFormat {
	ret := newQSurfaceFormat(QSurfaceFormat_new())
	ret.isSubclass = true
	return ret
}

// NewQSurfaceFormat2 constructs a new QSurfaceFormat object.
func NewQSurfaceFormat2(options FormatOptions) *QSurfaceFormat {
	ret := newQSurfaceFormat(QSurfaceFormat_new2(options))
	ret.isSubclass = true
	return ret
}

// NewQSurfaceFormat3 constructs a new QSurfaceFormat object.
func NewQSurfaceFormat3(other *QSurfaceFormat) *QSurfaceFormat {
	ret := newQSurfaceFormat(QSurfaceFormat_new3(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSurfaceFormat) OperatorAssign(other *QSurfaceFormat) {
	QSurfaceFormat_OperatorAssign(this.h, other.cPointer())
}

func (this *QSurfaceFormat) SetDepthBufferSize(size int) {
	QSurfaceFormat_SetDepthBufferSize(this.h, (int)(size))
}

func (this *QSurfaceFormat) DepthBufferSize() int {
	return (int)(QSurfaceFormat_DepthBufferSize(this.h))
}

func (this *QSurfaceFormat) SetStencilBufferSize(size int) {
	QSurfaceFormat_SetStencilBufferSize(this.h, (int)(size))
}

func (this *QSurfaceFormat) StencilBufferSize() int {
	return (int)(QSurfaceFormat_StencilBufferSize(this.h))
}

func (this *QSurfaceFormat) SetRedBufferSize(size int) {
	QSurfaceFormat_SetRedBufferSize(this.h, (int)(size))
}

func (this *QSurfaceFormat) RedBufferSize() int {
	return (int)(QSurfaceFormat_RedBufferSize(this.h))
}

func (this *QSurfaceFormat) SetGreenBufferSize(size int) {
	QSurfaceFormat_SetGreenBufferSize(this.h, (int)(size))
}

func (this *QSurfaceFormat) GreenBufferSize() int {
	return (int)(QSurfaceFormat_GreenBufferSize(this.h))
}

func (this *QSurfaceFormat) SetBlueBufferSize(size int) {
	QSurfaceFormat_SetBlueBufferSize(this.h, (int)(size))
}

func (this *QSurfaceFormat) BlueBufferSize() int {
	return (int)(QSurfaceFormat_BlueBufferSize(this.h))
}

func (this *QSurfaceFormat) SetAlphaBufferSize(size int) {
	QSurfaceFormat_SetAlphaBufferSize(this.h, (int)(size))
}

func (this *QSurfaceFormat) AlphaBufferSize() int {
	return (int)(QSurfaceFormat_AlphaBufferSize(this.h))
}

func (this *QSurfaceFormat) SetSamples(numSamples int) {
	QSurfaceFormat_SetSamples(this.h, (int)(numSamples))
}

func (this *QSurfaceFormat) Samples() int {
	return (int)(QSurfaceFormat_Samples(this.h))
}

func (this *QSurfaceFormat) SetSwapBehavior(behavior SwapBehavior) {
	QSurfaceFormat_SetSwapBehavior(this.h, behavior)
}

func (this *QSurfaceFormat) SwapBehavior() SwapBehavior {
	xxxxxxxxx
}

func (this *QSurfaceFormat) HasAlpha() bool {
	return (bool)(QSurfaceFormat_HasAlpha(this.h))
}

func (this *QSurfaceFormat) SetProfile(profile OpenGLContextProfile) {
	QSurfaceFormat_SetProfile(this.h, profile)
}

func (this *QSurfaceFormat) Profile() OpenGLContextProfile {
	xxxxxxxxx
}

func (this *QSurfaceFormat) SetRenderableType(typeVal RenderableType) {
	QSurfaceFormat_SetRenderableType(this.h, typeVal)
}

func (this *QSurfaceFormat) RenderableType() RenderableType {
	xxxxxxxxx
}

func (this *QSurfaceFormat) SetMajorVersion(majorVersion int) {
	QSurfaceFormat_SetMajorVersion(this.h, (int)(majorVersion))
}

func (this *QSurfaceFormat) MajorVersion() int {
	return (int)(QSurfaceFormat_MajorVersion(this.h))
}

func (this *QSurfaceFormat) SetMinorVersion(minorVersion int) {
	QSurfaceFormat_SetMinorVersion(this.h, (int)(minorVersion))
}

func (this *QSurfaceFormat) MinorVersion() int {
	return (int)(QSurfaceFormat_MinorVersion(this.h))
}

func (this *QSurfaceFormat) Version() struct {
	First  int
	Second int
} {
	var _mm struct_miqt_map = QSurfaceFormat_Version(this.h)
	_First_CArray := (*[0xffff]int)(unsafe.Pointer(_mm.keys))
	_Second_CArray := (*[0xffff]int)(unsafe.Pointer(_mm.values))
	_entry_First := (int)(_First_CArray[0])

	_entry_Second := (int)(_Second_CArray[0])

	return struct {
		First  int
		Second int
	}{First: _entry_First, Second: _entry_Second}
}

func (this *QSurfaceFormat) SetVersion(major int, minor int) {
	QSurfaceFormat_SetVersion(this.h, (int)(major), (int)(minor))
}

func (this *QSurfaceFormat) Stereo() bool {
	return (bool)(QSurfaceFormat_Stereo(this.h))
}

func (this *QSurfaceFormat) SetStereo(enable bool) {
	QSurfaceFormat_SetStereo(this.h, (bool)(enable))
}

func (this *QSurfaceFormat) SetOptions(options FormatOption) {
	QSurfaceFormat_SetOptions(this.h, (int)(options))
}

func (this *QSurfaceFormat) SetOption(option FormatOption) {
	QSurfaceFormat_SetOption(this.h, option)
}

func (this *QSurfaceFormat) TestOption(option FormatOption) bool {
	return (bool)(QSurfaceFormat_TestOption(this.h, option))
}

func (this *QSurfaceFormat) Options() FormatOption {
	return (FormatOption)(QSurfaceFormat_Options(this.h))
}

func (this *QSurfaceFormat) SwapInterval() int {
	return (int)(QSurfaceFormat_SwapInterval(this.h))
}

func (this *QSurfaceFormat) SetSwapInterval(interval int) {
	QSurfaceFormat_SetSwapInterval(this.h, (int)(interval))
}

func (this *QSurfaceFormat) ColorSpace() *QColorSpace {
	return newQColorSpace(QSurfaceFormat_ColorSpace(this.h))
}

func (this *QSurfaceFormat) SetColorSpace(colorSpace *QColorSpace) {
	QSurfaceFormat_SetColorSpace(this.h, colorSpace.cPointer())
}

func (this *QSurfaceFormat) SetColorSpaceWithColorSpace(colorSpace ColorSpace) {
	QSurfaceFormat_SetColorSpaceWithColorSpace(this.h, colorSpace)
}

func QSurfaceFormat_SetDefaultFormat(format *QSurfaceFormat) {
	QSurfaceFormat_SetDefaultFormat(format.cPointer())
}

func QSurfaceFormat_DefaultFormat() *QSurfaceFormat {
	_goptr := newQSurfaceFormat(QSurfaceFormat_DefaultFormat())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSurfaceFormat) SetOption2(option FormatOption, on bool) {
	QSurfaceFormat_SetOption2(this.h, option, (bool)(on))
}
