package qt

import (
	"unsafe"
)

type QBoxLayout__Direction int

const (
	QBoxLayout__LeftToRight QBoxLayout__Direction = 0
	QBoxLayout__RightToLeft QBoxLayout__Direction = 1
	QBoxLayout__TopToBottom QBoxLayout__Direction = 2
	QBoxLayout__BottomToTop QBoxLayout__Direction = 3
	QBoxLayout__Down        QBoxLayout__Direction = 2
	QBoxLayout__Up          QBoxLayout__Direction = 3
)

type QBoxLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQBoxLayout constructs a new QBoxLayout object.
func NewQBoxLayout(param1 Direction) *QBoxLayout {
	g := newQBoxLayout(QBoxLayout_new(param1))
	g.isSubclass = true
	return g
}

// NewQBoxLayout2 constructs a new QBoxLayout object.
func NewQBoxLayout2(param1 Direction, parent *QWidget) *QBoxLayout {
	g := newQBoxLayout(QBoxLayout_new2(param1, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QBoxLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QBoxLayout_MetaObject(this.h))
}

func (this *QBoxLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QBoxLayout_Metacast(this.h, param1_Cstring))
}

func QBoxLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QBoxLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QBoxLayout) Direction() Direction {
	xxxxxxxxx
}

func (this *QBoxLayout) SetDirection(direction Direction) {
	QBoxLayout_SetDirection(this.h, direction)
}

func (this *QBoxLayout) AddSpacing(size int) {
	QBoxLayout_AddSpacing(this.h, (int)(size))
}

func (this *QBoxLayout) AddStretch() {
	QBoxLayout_AddStretch(this.h)
}

func (this *QBoxLayout) AddSpacerItem(spacerItem *QSpacerItem) {
	QBoxLayout_AddSpacerItem(this.h, spacerItem.cPointer())
}

func (this *QBoxLayout) AddWidget(param1 *QWidget) {
	QBoxLayout_AddWidget(this.h, param1.cPointer())
}

func (this *QBoxLayout) AddLayout(layout *QLayout) {
	QBoxLayout_AddLayout(this.h, layout.cPointer())
}

func (this *QBoxLayout) AddStrut(param1 int) {
	QBoxLayout_AddStrut(this.h, (int)(param1))
}

func (this *QBoxLayout) AddItem(param1 *QLayoutItem) {
	QBoxLayout_AddItem(this.h, param1.cPointer())
}

func (this *QBoxLayout) InsertSpacing(index int, size int) {
	QBoxLayout_InsertSpacing(this.h, (int)(index), (int)(size))
}

func (this *QBoxLayout) InsertStretch(index int) {
	QBoxLayout_InsertStretch(this.h, (int)(index))
}

func (this *QBoxLayout) InsertSpacerItem(index int, spacerItem *QSpacerItem) {
	QBoxLayout_InsertSpacerItem(this.h, (int)(index), spacerItem.cPointer())
}

func (this *QBoxLayout) InsertWidget(index int, widget *QWidget) {
	QBoxLayout_InsertWidget(this.h, (int)(index), widget.cPointer())
}

func (this *QBoxLayout) InsertLayout(index int, layout *QLayout) {
	QBoxLayout_InsertLayout(this.h, (int)(index), layout.cPointer())
}

func (this *QBoxLayout) InsertItem(index int, param2 *QLayoutItem) {
	QBoxLayout_InsertItem(this.h, (int)(index), param2.cPointer())
}

func (this *QBoxLayout) Spacing() int {
	return (int)(QBoxLayout_Spacing(this.h))
}

func (this *QBoxLayout) SetSpacing(spacing int) {
	QBoxLayout_SetSpacing(this.h, (int)(spacing))
}

func (this *QBoxLayout) SetStretchFactor(w *QWidget, stretch int) bool {
	return (bool)(QBoxLayout_SetStretchFactor(this.h, w.cPointer(), (int)(stretch)))
}

func (this *QBoxLayout) SetStretchFactor2(l *QLayout, stretch int) bool {
	return (bool)(QBoxLayout_SetStretchFactor2(this.h, l.cPointer(), (int)(stretch)))
}

func (this *QBoxLayout) SetStretch(index int, stretch int) {
	QBoxLayout_SetStretch(this.h, (int)(index), (int)(stretch))
}

func (this *QBoxLayout) Stretch(index int) int {
	return (int)(QBoxLayout_Stretch(this.h, (int)(index)))
}

func (this *QBoxLayout) SizeHint() *QSize {
	_goptr := newQSize(QBoxLayout_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBoxLayout) MinimumSize() *QSize {
	_goptr := newQSize(QBoxLayout_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBoxLayout) MaximumSize() *QSize {
	_goptr := newQSize(QBoxLayout_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBoxLayout) HasHeightForWidth() bool {
	return (bool)(QBoxLayout_HasHeightForWidth(this.h))
}

func (this *QBoxLayout) HeightForWidth(param1 int) int {
	return (int)(QBoxLayout_HeightForWidth(this.h, (int)(param1)))
}

func (this *QBoxLayout) MinimumHeightForWidth(param1 int) int {
	return (int)(QBoxLayout_MinimumHeightForWidth(this.h, (int)(param1)))
}

func (this *QBoxLayout) ExpandingDirections() Orientation {
	return (Orientation)(QBoxLayout_ExpandingDirections(this.h))
}

func (this *QBoxLayout) Invalidate() {
	QBoxLayout_Invalidate(this.h)
}

func (this *QBoxLayout) ItemAt(param1 int) *QLayoutItem {
	return newQLayoutItem(QBoxLayout_ItemAt(this.h, (int)(param1)))
}

func (this *QBoxLayout) TakeAt(param1 int) *QLayoutItem {
	return newQLayoutItem(QBoxLayout_TakeAt(this.h, (int)(param1)))
}

func (this *QBoxLayout) Count() int {
	return (int)(QBoxLayout_Count(this.h))
}

func (this *QBoxLayout) SetGeometry(geometry *QRect) {
	QBoxLayout_SetGeometry(this.h, geometry.cPointer())
}

func QBoxLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QBoxLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QBoxLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QBoxLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QBoxLayout) AddStretch1(stretch int) {
	QBoxLayout_AddStretch1(this.h, (int)(stretch))
}

func (this *QBoxLayout) AddWidget2(param1 *QWidget, stretch int) {
	QBoxLayout_AddWidget2(this.h, param1.cPointer(), (int)(stretch))
}

func (this *QBoxLayout) AddWidget3(param1 *QWidget, stretch int, alignment AlignmentFlag) {
	QBoxLayout_AddWidget3(this.h, param1.cPointer(), (int)(stretch), (int)(alignment))
}

func (this *QBoxLayout) AddLayout2(layout *QLayout, stretch int) {
	QBoxLayout_AddLayout2(this.h, layout.cPointer(), (int)(stretch))
}

func (this *QBoxLayout) InsertStretch2(index int, stretch int) {
	QBoxLayout_InsertStretch2(this.h, (int)(index), (int)(stretch))
}

func (this *QBoxLayout) InsertWidget3(index int, widget *QWidget, stretch int) {
	QBoxLayout_InsertWidget3(this.h, (int)(index), widget.cPointer(), (int)(stretch))
}

func (this *QBoxLayout) InsertWidget4(index int, widget *QWidget, stretch int, alignment AlignmentFlag) {
	QBoxLayout_InsertWidget4(this.h, (int)(index), widget.cPointer(), (int)(stretch), (int)(alignment))
}

func (this *QBoxLayout) InsertLayout3(index int, layout *QLayout, stretch int) {
	QBoxLayout_InsertLayout3(this.h, (int)(index), layout.cPointer(), (int)(stretch))
}

func (this *QBoxLayout) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QBoxLayout_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QBoxLayout) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_MetaObject
func miqt_exec_callback_QBoxLayout_MetaObject(self QBoxLayout, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QBoxLayout) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QBoxLayout_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QBoxLayout) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBoxLayout_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBoxLayout_Metacast
func miqt_exec_callback_QBoxLayout_Metacast(self QBoxLayout, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QBoxLayout{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QHBoxLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQHBoxLayout constructs a new QHBoxLayout object.
func NewQHBoxLayout(parent *QWidget) *QHBoxLayout {
	g := newQHBoxLayout(QHBoxLayout_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQHBoxLayout2 constructs a new QHBoxLayout object.
func NewQHBoxLayout2() *QHBoxLayout {
	g := newQHBoxLayout(QHBoxLayout_new2())
	g.isSubclass = true
	return g
}

func (this *QHBoxLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QHBoxLayout_MetaObject(this.h))
}

func (this *QHBoxLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QHBoxLayout_Metacast(this.h, param1_Cstring))
}

func QHBoxLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QHBoxLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QHBoxLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QHBoxLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QHBoxLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QHBoxLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QHBoxLayout) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QHBoxLayout_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QHBoxLayout) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_MetaObject
func miqt_exec_callback_QHBoxLayout_MetaObject(self QHBoxLayout, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QHBoxLayout) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QHBoxLayout_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QHBoxLayout) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHBoxLayout_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHBoxLayout_Metacast
func miqt_exec_callback_QHBoxLayout_Metacast(self QHBoxLayout, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QHBoxLayout{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QVBoxLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQVBoxLayout constructs a new QVBoxLayout object.
func NewQVBoxLayout(parent *QWidget) *QVBoxLayout {
	g := newQVBoxLayout(QVBoxLayout_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQVBoxLayout2 constructs a new QVBoxLayout object.
func NewQVBoxLayout2() *QVBoxLayout {
	g := newQVBoxLayout(QVBoxLayout_new2())
	g.isSubclass = true
	return g
}

func (this *QVBoxLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QVBoxLayout_MetaObject(this.h))
}

func (this *QVBoxLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QVBoxLayout_Metacast(this.h, param1_Cstring))
}

func QVBoxLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QVBoxLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVBoxLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVBoxLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVBoxLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVBoxLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVBoxLayout) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QVBoxLayout_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QVBoxLayout) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_MetaObject
func miqt_exec_callback_QVBoxLayout_MetaObject(self QVBoxLayout, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QVBoxLayout) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QVBoxLayout_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QVBoxLayout) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVBoxLayout_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVBoxLayout_Metacast
func miqt_exec_callback_QVBoxLayout_Metacast(self QVBoxLayout, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QVBoxLayout{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
