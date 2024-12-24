package qt

import (
	"unsafe"
)

type QLayout__SizeConstraint int

const (
	QLayout__SetDefaultConstraint QLayout__SizeConstraint = 0
	QLayout__SetNoConstraint      QLayout__SizeConstraint = 1
	QLayout__SetMinimumSize       QLayout__SizeConstraint = 2
	QLayout__SetFixedSize         QLayout__SizeConstraint = 3
	QLayout__SetMaximumSize       QLayout__SizeConstraint = 4
	QLayout__SetMinAndMaxSize     QLayout__SizeConstraint = 5
)

type QLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQLayout constructs a new QLayout object.
func NewQLayout(parent *QWidget) *QLayout {
	g := newQLayout(QLayout_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQLayout2 constructs a new QLayout object.
func NewQLayout2() *QLayout {
	g := newQLayout(QLayout_new2())
	g.isSubclass = true
	return g
}

func (this *QLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QLayout_MetaObject(this.h))
}

func (this *QLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QLayout_Metacast(this.h, param1_Cstring))
}

func QLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLayout) Spacing() int {
	return (int)(QLayout_Spacing(this.h))
}

func (this *QLayout) SetSpacing(spacing int) {
	QLayout_SetSpacing(this.h, (int)(spacing))
}

func (this *QLayout) SetContentsMargins(left int, top int, right int, bottom int) {
	QLayout_SetContentsMargins(this.h, (int)(left), (int)(top), (int)(right), (int)(bottom))
}

func (this *QLayout) SetContentsMarginsWithMargins(margins *QMargins) {
	QLayout_SetContentsMarginsWithMargins(this.h, margins.cPointer())
}

func (this *QLayout) UnsetContentsMargins() {
	QLayout_UnsetContentsMargins(this.h)
}

func (this *QLayout) GetContentsMargins(left *int, top *int, right *int, bottom *int) {
	QLayout_GetContentsMargins(this.h, (*int)(unsafe.Pointer(left)), (*int)(unsafe.Pointer(top)), (*int)(unsafe.Pointer(right)), (*int)(unsafe.Pointer(bottom)))
}

func (this *QLayout) ContentsMargins() *QMargins {
	_goptr := newQMargins(QLayout_ContentsMargins(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) ContentsRect() *QRect {
	_goptr := newQRect(QLayout_ContentsRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) SetAlignment(w *QWidget, alignment AlignmentFlag) bool {
	return (bool)(QLayout_SetAlignment(this.h, w.cPointer(), (int)(alignment)))
}

func (this *QLayout) SetAlignment2(l *QLayout, alignment AlignmentFlag) bool {
	return (bool)(QLayout_SetAlignment2(this.h, l.cPointer(), (int)(alignment)))
}

func (this *QLayout) SetSizeConstraint(sizeConstraint SizeConstraint) {
	QLayout_SetSizeConstraint(this.h, sizeConstraint)
}

func (this *QLayout) SizeConstraint() SizeConstraint {
	xxxxxxxxx
}

func (this *QLayout) SetMenuBar(w *QWidget) {
	QLayout_SetMenuBar(this.h, w.cPointer())
}

func (this *QLayout) MenuBar() *QWidget {
	return newQWidget(QLayout_MenuBar(this.h))
}

func (this *QLayout) ParentWidget() *QWidget {
	return newQWidget(QLayout_ParentWidget(this.h))
}

func (this *QLayout) Invalidate() {
	QLayout_Invalidate(this.h)
}

func (this *QLayout) Geometry() *QRect {
	_goptr := newQRect(QLayout_Geometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) Activate() bool {
	return (bool)(QLayout_Activate(this.h))
}

func (this *QLayout) Update() {
	QLayout_Update(this.h)
}

func (this *QLayout) AddWidget(w *QWidget) {
	QLayout_AddWidget(this.h, w.cPointer())
}

func (this *QLayout) AddItem(param1 *QLayoutItem) {
	QLayout_AddItem(this.h, param1.cPointer())
}

func (this *QLayout) RemoveWidget(w *QWidget) {
	QLayout_RemoveWidget(this.h, w.cPointer())
}

func (this *QLayout) RemoveItem(param1 *QLayoutItem) {
	QLayout_RemoveItem(this.h, param1.cPointer())
}

func (this *QLayout) ExpandingDirections() Orientation {
	return (Orientation)(QLayout_ExpandingDirections(this.h))
}

func (this *QLayout) MinimumSize() *QSize {
	_goptr := newQSize(QLayout_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) MaximumSize() *QSize {
	_goptr := newQSize(QLayout_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) SetGeometry(geometry *QRect) {
	QLayout_SetGeometry(this.h, geometry.cPointer())
}

func (this *QLayout) ItemAt(index int) *QLayoutItem {
	return newQLayoutItem(QLayout_ItemAt(this.h, (int)(index)))
}

func (this *QLayout) TakeAt(index int) *QLayoutItem {
	return newQLayoutItem(QLayout_TakeAt(this.h, (int)(index)))
}

func (this *QLayout) IndexOf(param1 *QWidget) int {
	return (int)(QLayout_IndexOf(this.h, param1.cPointer()))
}

func (this *QLayout) IndexOfWithQLayoutItem(param1 *QLayoutItem) int {
	return (int)(QLayout_IndexOfWithQLayoutItem(this.h, param1.cPointer()))
}

func (this *QLayout) Count() int {
	return (int)(QLayout_Count(this.h))
}

func (this *QLayout) IsEmpty() bool {
	return (bool)(QLayout_IsEmpty(this.h))
}

func (this *QLayout) ControlTypes() ControlType {
	return (ControlType)(QLayout_ControlTypes(this.h))
}

func (this *QLayout) ReplaceWidget(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {
	return newQLayoutItem(QLayout_ReplaceWidget(this.h, from.cPointer(), to.cPointer(), (int)(options)))
}

func (this *QLayout) TotalMinimumHeightForWidth(w int) int {
	return (int)(QLayout_TotalMinimumHeightForWidth(this.h, (int)(w)))
}

func (this *QLayout) TotalHeightForWidth(w int) int {
	return (int)(QLayout_TotalHeightForWidth(this.h, (int)(w)))
}

func (this *QLayout) TotalMinimumSize() *QSize {
	_goptr := newQSize(QLayout_TotalMinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) TotalMaximumSize() *QSize {
	_goptr := newQSize(QLayout_TotalMaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) TotalSizeHint() *QSize {
	_goptr := newQSize(QLayout_TotalSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayout) Layout() *QLayout {
	return newQLayout(QLayout_Layout(this.h))
}

func (this *QLayout) SetEnabled(enabled bool) {
	QLayout_SetEnabled(this.h, (bool)(enabled))
}

func (this *QLayout) IsEnabled() bool {
	return (bool)(QLayout_IsEnabled(this.h))
}

func QLayout_ClosestAcceptableSize(w *QWidget, s *QSize) *QSize {
	_goptr := newQSize(QLayout_ClosestAcceptableSize(w.cPointer(), s.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLayout) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QLayout_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QLayout) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_MetaObject
func miqt_exec_callback_QLayout_MetaObject(self QLayout, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QLayout) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QLayout_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QLayout) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLayout_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLayout_Metacast
func miqt_exec_callback_QLayout_Metacast(self QLayout, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QLayout{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
