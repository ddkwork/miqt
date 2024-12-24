package qt

import (
	"unsafe"
)

type QTextStream struct {
	h          uintptr
	isSubclass bool
}

type QSplitter struct {
	h          uintptr
	isSubclass bool
}

// NewQSplitter constructs a new QSplitter object.
func NewQSplitter(parent *QWidget) *QSplitter {
	ret := newQSplitter(QSplitter_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSplitter2 constructs a new QSplitter object.
func NewQSplitter2() *QSplitter {
	ret := newQSplitter(QSplitter_new2())
	ret.isSubclass = true
	return ret
}

// NewQSplitter3 constructs a new QSplitter object.
func NewQSplitter3(param1 Orientation) *QSplitter {
	ret := newQSplitter(QSplitter_new3((int)(param1)))
	ret.isSubclass = true
	return ret
}

// NewQSplitter4 constructs a new QSplitter object.
func NewQSplitter4(param1 Orientation, parent *QWidget) *QSplitter {
	ret := newQSplitter(QSplitter_new4((int)(param1), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSplitter) MetaObject() *QMetaObject {
	return newQMetaObject(QSplitter_MetaObject(this.h))
}

func (this *QSplitter) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSplitter_Metacast(this.h, param1_Cstring))
}

func QSplitter_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSplitter_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSplitter) AddWidget(widget *QWidget) {
	QSplitter_AddWidget(this.h, widget.cPointer())
}

func (this *QSplitter) InsertWidget(index int, widget *QWidget) {
	QSplitter_InsertWidget(this.h, (int)(index), widget.cPointer())
}

func (this *QSplitter) ReplaceWidget(index int, widget *QWidget) *QWidget {
	return newQWidget(QSplitter_ReplaceWidget(this.h, (int)(index), widget.cPointer()))
}

func (this *QSplitter) SetOrientation(orientation Orientation) {
	QSplitter_SetOrientation(this.h, (int)(orientation))
}

func (this *QSplitter) Orientation() Orientation {
	return (Orientation)(QSplitter_Orientation(this.h))
}

func (this *QSplitter) SetChildrenCollapsible(childrenCollapsible bool) {
	QSplitter_SetChildrenCollapsible(this.h, (bool)(childrenCollapsible))
}

func (this *QSplitter) ChildrenCollapsible() bool {
	return (bool)(QSplitter_ChildrenCollapsible(this.h))
}

func (this *QSplitter) SetCollapsible(index int, param2 bool) {
	QSplitter_SetCollapsible(this.h, (int)(index), (bool)(param2))
}

func (this *QSplitter) IsCollapsible(index int) bool {
	return (bool)(QSplitter_IsCollapsible(this.h, (int)(index)))
}

func (this *QSplitter) SetOpaqueResize() {
	QSplitter_SetOpaqueResize(this.h)
}

func (this *QSplitter) OpaqueResize() bool {
	return (bool)(QSplitter_OpaqueResize(this.h))
}

func (this *QSplitter) Refresh() {
	QSplitter_Refresh(this.h)
}

func (this *QSplitter) SizeHint() *QSize {
	_goptr := newQSize(QSplitter_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSplitter) MinimumSizeHint() *QSize {
	_goptr := newQSize(QSplitter_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSplitter) Sizes() []int {
	var _ma struct_miqt_array = QSplitter_Sizes(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func (this *QSplitter) SetSizes(list []int) {
	list_CArray := (*[0xffff]int)(malloc(size_t(8 * len(list))))
	defer free(unsafe.Pointer(list_CArray))
	for i := range list {
		list_CArray[i] = (int)(list[i])
	}
	list_ma := struct_miqt_array{len: size_t(len(list)), data: unsafe.Pointer(list_CArray)}
	QSplitter_SetSizes(this.h, list_ma)
}

func (this *QSplitter) SaveState() []byte {
	var _bytearray struct_miqt_string = QSplitter_SaveState(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QSplitter) RestoreState(state []byte) bool {
	state_alias := struct_miqt_string{}
	state_alias.data = (char)(unsafe.Pointer(&state[0]))
	state_alias.len = size_t(len(state))
	return (bool)(QSplitter_RestoreState(this.h, state_alias))
}

func (this *QSplitter) HandleWidth() int {
	return (int)(QSplitter_HandleWidth(this.h))
}

func (this *QSplitter) SetHandleWidth(handleWidth int) {
	QSplitter_SetHandleWidth(this.h, (int)(handleWidth))
}

func (this *QSplitter) IndexOf(w *QWidget) int {
	return (int)(QSplitter_IndexOf(this.h, w.cPointer()))
}

func (this *QSplitter) Widget(index int) *QWidget {
	return newQWidget(QSplitter_Widget(this.h, (int)(index)))
}

func (this *QSplitter) Count() int {
	return (int)(QSplitter_Count(this.h))
}

func (this *QSplitter) GetRange(index int, param2 *int, param3 *int) {
	QSplitter_GetRange(this.h, (int)(index), (*int)(unsafe.Pointer(param2)), (*int)(unsafe.Pointer(param3)))
}

func (this *QSplitter) Handle(index int) *QSplitterHandle {
	return newQSplitterHandle(QSplitter_Handle(this.h, (int)(index)))
}

func (this *QSplitter) SetStretchFactor(index int, stretch int) {
	QSplitter_SetStretchFactor(this.h, (int)(index), (int)(stretch))
}

func (this *QSplitter) SplitterMoved(pos int, index int) {
	QSplitter_SplitterMoved(this.h, (int)(pos), (int)(index))
}

func (this *QSplitter) OnSplitterMoved(slot func(pos int, index int)) {
	QSplitter_connect_SplitterMoved(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_SplitterMoved
func miqt_exec_callback_QSplitter_SplitterMoved(cb intptr_t, pos int, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(pos int, index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(pos)

	slotval2 := (int)(index)

	gofunc(slotval1, slotval2)
}

func QSplitter_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSplitter_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSplitter_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSplitter_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSplitter) SetOpaqueResize1(opaque bool) {
	QSplitter_SetOpaqueResize1(this.h, (bool)(opaque))
}

func (this *QSplitter) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSplitter_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSplitter) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitter_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_MetaObject
func miqt_exec_callback_QSplitter_MetaObject(self QSplitter, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSplitter{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSplitter) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSplitter_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSplitter) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitter_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitter_Metacast
func miqt_exec_callback_QSplitter_Metacast(self QSplitter, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QSplitter{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QSplitterHandle struct {
	h          uintptr
	isSubclass bool
}

// NewQSplitterHandle constructs a new QSplitterHandle object.
func NewQSplitterHandle(o Orientation, parent *QSplitter) *QSplitterHandle {
	ret := newQSplitterHandle(QSplitterHandle_new((int)(o), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSplitterHandle) MetaObject() *QMetaObject {
	return newQMetaObject(QSplitterHandle_MetaObject(this.h))
}

func (this *QSplitterHandle) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSplitterHandle_Metacast(this.h, param1_Cstring))
}

func QSplitterHandle_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSplitterHandle_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSplitterHandle) SetOrientation(o Orientation) {
	QSplitterHandle_SetOrientation(this.h, (int)(o))
}

func (this *QSplitterHandle) Orientation() Orientation {
	return (Orientation)(QSplitterHandle_Orientation(this.h))
}

func (this *QSplitterHandle) OpaqueResize() bool {
	return (bool)(QSplitterHandle_OpaqueResize(this.h))
}

func (this *QSplitterHandle) Splitter() *QSplitter {
	return newQSplitter(QSplitterHandle_Splitter(this.h))
}

func (this *QSplitterHandle) SizeHint() *QSize {
	_goptr := newQSize(QSplitterHandle_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSplitterHandle_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSplitterHandle_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSplitterHandle_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSplitterHandle_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSplitterHandle) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSplitterHandle_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSplitterHandle) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_MetaObject
func miqt_exec_callback_QSplitterHandle_MetaObject(self QSplitterHandle, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSplitterHandle) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSplitterHandle_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSplitterHandle) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplitterHandle_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplitterHandle_Metacast
func miqt_exec_callback_QSplitterHandle_Metacast(self QSplitterHandle, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QSplitterHandle{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
