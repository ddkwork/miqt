package qt

import (
	"unsafe"
)

type QAbstractScrollArea__SizeAdjustPolicy int

const (
	QAbstractScrollArea__AdjustIgnored               QAbstractScrollArea__SizeAdjustPolicy = 0
	QAbstractScrollArea__AdjustToContentsOnFirstShow QAbstractScrollArea__SizeAdjustPolicy = 1
	QAbstractScrollArea__AdjustToContents            QAbstractScrollArea__SizeAdjustPolicy = 2
)

type QAbstractScrollArea struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractScrollArea constructs a new QAbstractScrollArea object.
func NewQAbstractScrollArea(parent *QWidget) *QAbstractScrollArea {
	g := newQAbstractScrollArea(QAbstractScrollArea_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQAbstractScrollArea2 constructs a new QAbstractScrollArea object.
func NewQAbstractScrollArea2() *QAbstractScrollArea {
	g := newQAbstractScrollArea(QAbstractScrollArea_new2())
	g.isSubclass = true
	return g
}

func (this *QAbstractScrollArea) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractScrollArea_MetaObject(this.h))
}

func (this *QAbstractScrollArea) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractScrollArea_Metacast(this.h, param1_Cstring))
}

func QAbstractScrollArea_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractScrollArea_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractScrollArea) VerticalScrollBarPolicy() ScrollBarPolicy {
	return (ScrollBarPolicy)(QAbstractScrollArea_VerticalScrollBarPolicy(this.h))
}

func (this *QAbstractScrollArea) SetVerticalScrollBarPolicy(verticalScrollBarPolicy ScrollBarPolicy) {
	QAbstractScrollArea_SetVerticalScrollBarPolicy(this.h, (int)(verticalScrollBarPolicy))
}

func (this *QAbstractScrollArea) VerticalScrollBar() *QScrollBar {
	return newQScrollBar(QAbstractScrollArea_VerticalScrollBar(this.h))
}

func (this *QAbstractScrollArea) SetVerticalScrollBar(scrollbar *QScrollBar) {
	QAbstractScrollArea_SetVerticalScrollBar(this.h, scrollbar.cPointer())
}

func (this *QAbstractScrollArea) HorizontalScrollBarPolicy() ScrollBarPolicy {
	return (ScrollBarPolicy)(QAbstractScrollArea_HorizontalScrollBarPolicy(this.h))
}

func (this *QAbstractScrollArea) SetHorizontalScrollBarPolicy(horizontalScrollBarPolicy ScrollBarPolicy) {
	QAbstractScrollArea_SetHorizontalScrollBarPolicy(this.h, (int)(horizontalScrollBarPolicy))
}

func (this *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar {
	return newQScrollBar(QAbstractScrollArea_HorizontalScrollBar(this.h))
}

func (this *QAbstractScrollArea) SetHorizontalScrollBar(scrollbar *QScrollBar) {
	QAbstractScrollArea_SetHorizontalScrollBar(this.h, scrollbar.cPointer())
}

func (this *QAbstractScrollArea) CornerWidget() *QWidget {
	return newQWidget(QAbstractScrollArea_CornerWidget(this.h))
}

func (this *QAbstractScrollArea) SetCornerWidget(widget *QWidget) {
	QAbstractScrollArea_SetCornerWidget(this.h, widget.cPointer())
}

func (this *QAbstractScrollArea) AddScrollBarWidget(widget *QWidget, alignment AlignmentFlag) {
	QAbstractScrollArea_AddScrollBarWidget(this.h, widget.cPointer(), (int)(alignment))
}

func (this *QAbstractScrollArea) ScrollBarWidgets(alignment AlignmentFlag) []*QWidget {
	var _ma struct_miqt_array = QAbstractScrollArea_ScrollBarWidgets(this.h, (int)(alignment))
	_ret := make([]*QWidget, int(_ma.len))
	_outCast := (*[0xffff]*QWidget)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQWidget(_outCast[i])
	}
	return _ret
}

func (this *QAbstractScrollArea) Viewport() *QWidget {
	return newQWidget(QAbstractScrollArea_Viewport(this.h))
}

func (this *QAbstractScrollArea) SetViewport(widget *QWidget) {
	QAbstractScrollArea_SetViewport(this.h, widget.cPointer())
}

func (this *QAbstractScrollArea) MaximumViewportSize() *QSize {
	_goptr := newQSize(QAbstractScrollArea_MaximumViewportSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractScrollArea) MinimumSizeHint() *QSize {
	_goptr := newQSize(QAbstractScrollArea_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractScrollArea) SizeHint() *QSize {
	_goptr := newQSize(QAbstractScrollArea_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractScrollArea) SetupViewport(viewport *QWidget) {
	QAbstractScrollArea_SetupViewport(this.h, viewport.cPointer())
}

func (this *QAbstractScrollArea) SizeAdjustPolicy() SizeAdjustPolicy {
	xxxxxxxxx
}

func (this *QAbstractScrollArea) SetSizeAdjustPolicy(policy SizeAdjustPolicy) {
	QAbstractScrollArea_SetSizeAdjustPolicy(this.h, policy)
}

func QAbstractScrollArea_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractScrollArea_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractScrollArea_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractScrollArea_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractScrollArea) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractScrollArea_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAbstractScrollArea) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_MetaObject
func miqt_exec_callback_QAbstractScrollArea_MetaObject(self QAbstractScrollArea, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAbstractScrollArea) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAbstractScrollArea_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAbstractScrollArea) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractScrollArea_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractScrollArea_Metacast
func miqt_exec_callback_QAbstractScrollArea_Metacast(self QAbstractScrollArea, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QAbstractScrollArea{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
