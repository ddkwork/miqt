package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QVideoSink struct {
	h          uintptr
	isSubclass bool
}

// NewQVideoSink constructs a new QVideoSink object.
func NewQVideoSink() *QVideoSink {
	ret := newQVideoSink(QVideoSink_new())
	ret.isSubclass = true
	return ret
}

// NewQVideoSink2 constructs a new QVideoSink object.
func NewQVideoSink2(parent *qt.QObject) *QVideoSink {
	ret := newQVideoSink(QVideoSink_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QVideoSink) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QVideoSink_MetaObject(this.h)))
}

func (this *QVideoSink) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QVideoSink_Metacast(this.h, param1_Cstring))
}

func QVideoSink_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QVideoSink_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoSink) VideoSize() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QVideoSink_VideoSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoSink) SubtitleText() string {
	var _ms struct_miqt_string = QVideoSink_SubtitleText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoSink) SetSubtitleText(subtitle string) {
	subtitle_ms := struct_miqt_string{}
	subtitle_ms.data = CString(subtitle)
	subtitle_ms.len = size_t(len(subtitle))
	defer free(unsafe.Pointer(subtitle_ms.data))
	QVideoSink_SetSubtitleText(this.h, subtitle_ms)
}

func (this *QVideoSink) SetVideoFrame(frame *QVideoFrame) {
	QVideoSink_SetVideoFrame(this.h, frame.cPointer())
}

func (this *QVideoSink) VideoFrame() *QVideoFrame {
	_goptr := newQVideoFrame(QVideoSink_VideoFrame(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoSink) VideoFrameChanged(frame *QVideoFrame) {
	QVideoSink_VideoFrameChanged(this.h, frame.cPointer())
}

func (this *QVideoSink) OnVideoFrameChanged(slot func(frame *QVideoFrame)) {
	QVideoSink_connect_VideoFrameChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoSink_VideoFrameChanged
func miqt_exec_callback_QVideoSink_VideoFrameChanged(cb intptr_t, frame *QVideoFrame) {
	gofunc, ok := cgo.Handle(cb).Value().(func(frame *QVideoFrame))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQVideoFrame(frame)

	gofunc(slotval1)
}

func (this *QVideoSink) SubtitleTextChanged(subtitleText string) {
	subtitleText_ms := struct_miqt_string{}
	subtitleText_ms.data = CString(subtitleText)
	subtitleText_ms.len = size_t(len(subtitleText))
	defer free(unsafe.Pointer(subtitleText_ms.data))
	QVideoSink_SubtitleTextChanged(this.h, subtitleText_ms)
}

func (this *QVideoSink) OnSubtitleTextChanged(slot func(subtitleText string)) {
	QVideoSink_connect_SubtitleTextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoSink_SubtitleTextChanged
func miqt_exec_callback_QVideoSink_SubtitleTextChanged(cb intptr_t, subtitleText struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(subtitleText string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var subtitleText_ms struct_miqt_string = subtitleText
	subtitleText_ret := GoStringN(subtitleText_ms.data, int(int64(subtitleText_ms.len)))
	free(unsafe.Pointer(subtitleText_ms.data))
	slotval1 := subtitleText_ret

	gofunc(slotval1)
}

func (this *QVideoSink) VideoSizeChanged() {
	QVideoSink_VideoSizeChanged(this.h)
}

func (this *QVideoSink) OnVideoSizeChanged(slot func()) {
	QVideoSink_connect_VideoSizeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoSink_VideoSizeChanged
func miqt_exec_callback_QVideoSink_VideoSizeChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QVideoSink_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVideoSink_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVideoSink_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVideoSink_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoSink) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QVideoSink_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QVideoSink) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVideoSink_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoSink_MetaObject
func miqt_exec_callback_QVideoSink_MetaObject(self QVideoSink, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVideoSink{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QVideoSink) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QVideoSink_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QVideoSink) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVideoSink_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoSink_Metacast
func miqt_exec_callback_QVideoSink_Metacast(self QVideoSink, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QVideoSink{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
