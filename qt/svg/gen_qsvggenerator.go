package svg

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QSvgGenerator__SvgVersion int

const (
	QSvgGenerator__SvgTiny12 QSvgGenerator__SvgVersion = 0
	QSvgGenerator__Svg11     QSvgGenerator__SvgVersion = 1
)

type QSvgGenerator struct {
	h          uintptr
	isSubclass bool
}

// NewQSvgGenerator constructs a new QSvgGenerator object.
func NewQSvgGenerator() *QSvgGenerator {
	ret := newQSvgGenerator(QSvgGenerator_new())
	ret.isSubclass = true
	return ret
}

// NewQSvgGenerator2 constructs a new QSvgGenerator object.
func NewQSvgGenerator2(version SvgVersion) *QSvgGenerator {
	ret := newQSvgGenerator(QSvgGenerator_new2(version))
	ret.isSubclass = true
	return ret
}

func (this *QSvgGenerator) Title() string {
	var _ms struct_miqt_string = QSvgGenerator_Title(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSvgGenerator) SetTitle(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QSvgGenerator_SetTitle(this.h, title_ms)
}

func (this *QSvgGenerator) Description() string {
	var _ms struct_miqt_string = QSvgGenerator_Description(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSvgGenerator) SetDescription(description string) {
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	QSvgGenerator_SetDescription(this.h, description_ms)
}

func (this *QSvgGenerator) Size() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QSvgGenerator_Size(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSvgGenerator) SetSize(size *qt.QSize) {
	QSvgGenerator_SetSize(this.h, (*QSize)(size.UnsafePointer()))
}

func (this *QSvgGenerator) ViewBox() *qt.QRect {
	_goptr := qt.UnsafeNewQRect(unsafe.Pointer(QSvgGenerator_ViewBox(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSvgGenerator) ViewBoxF() *qt.QRectF {
	_goptr := qt.UnsafeNewQRectF(unsafe.Pointer(QSvgGenerator_ViewBoxF(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSvgGenerator) SetViewBox(viewBox *qt.QRect) {
	QSvgGenerator_SetViewBox(this.h, (*QRect)(viewBox.UnsafePointer()))
}

func (this *QSvgGenerator) SetViewBoxWithViewBox(viewBox *qt.QRectF) {
	QSvgGenerator_SetViewBoxWithViewBox(this.h, (*QRectF)(viewBox.UnsafePointer()))
}

func (this *QSvgGenerator) FileName() string {
	var _ms struct_miqt_string = QSvgGenerator_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSvgGenerator) SetFileName(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QSvgGenerator_SetFileName(this.h, fileName_ms)
}

func (this *QSvgGenerator) OutputDevice() *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QSvgGenerator_OutputDevice(this.h)))
}

func (this *QSvgGenerator) SetOutputDevice(outputDevice *qt.QIODevice) {
	QSvgGenerator_SetOutputDevice(this.h, (*QIODevice)(outputDevice.UnsafePointer()))
}

func (this *QSvgGenerator) SetResolution(dpi int) {
	QSvgGenerator_SetResolution(this.h, (int)(dpi))
}

func (this *QSvgGenerator) Resolution() int {
	return (int)(QSvgGenerator_Resolution(this.h))
}

func (this *QSvgGenerator) SvgVersion() SvgVersion {
	xxxxxxxxx
}
