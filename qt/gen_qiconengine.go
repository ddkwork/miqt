package qt

import (
	"unsafe"
)

type QIconEngine__IconEngineHook int

const (
	QIconEngine__IsNullHook       QIconEngine__IconEngineHook = 3
	QIconEngine__ScaledPixmapHook QIconEngine__IconEngineHook = 4
)

type QIconEngine struct {
	h          uintptr
	isSubclass bool
}

// NewQIconEngine constructs a new QIconEngine object.
func NewQIconEngine() *QIconEngine {

	ret := newQIconEngine(QIconEngine_new())
	ret.isSubclass = true
	return ret
}

func (this *QIconEngine) Paint(painter *QPainter, rect *QRect, mode QIcon__Mode, state QIcon__State) {
	QIconEngine_Paint(this.h, painter.cPointer(), rect.cPointer(), (int)(mode), (int)(state))
}

func (this *QIconEngine) ActualSize(size *QSize, mode QIcon__Mode, state QIcon__State) *QSize {
	_goptr := newQSize(QIconEngine_ActualSize(this.h, size.cPointer(), (int)(mode), (int)(state)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIconEngine) Pixmap(size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap {
	_goptr := newQPixmap(QIconEngine_Pixmap(this.h, size.cPointer(), (int)(mode), (int)(state)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIconEngine) AddPixmap(pixmap *QPixmap, mode QIcon__Mode, state QIcon__State) {
	QIconEngine_AddPixmap(this.h, pixmap.cPointer(), (int)(mode), (int)(state))
}

func (this *QIconEngine) AddFile(fileName string, size *QSize, mode QIcon__Mode, state QIcon__State) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QIconEngine_AddFile(this.h, fileName_ms, size.cPointer(), (int)(mode), (int)(state))
}

func (this *QIconEngine) Key() string {
	var _ms struct_miqt_string = QIconEngine_Key(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIconEngine) Clone() *QIconEngine {
	return newQIconEngine(QIconEngine_Clone(this.h))
}

func (this *QIconEngine) Read(in *QDataStream) bool {
	return (bool)(QIconEngine_Read(this.h, in.cPointer()))
}

func (this *QIconEngine) Write(out *QDataStream) bool {
	return (bool)(QIconEngine_Write(this.h, out.cPointer()))
}

func (this *QIconEngine) AvailableSizes(mode QIcon__Mode, state QIcon__State) []QSize {
	var _ma struct_miqt_array = QIconEngine_AvailableSizes(this.h, (int)(mode), (int)(state))
	_ret := make([]QSize, int(_ma.len))
	_outCast := (*[0xffff]*QSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSize(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QIconEngine) IconName() string {
	var _ms struct_miqt_string = QIconEngine_IconName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIconEngine) IsNull() bool {
	return (bool)(QIconEngine_IsNull(this.h))
}

func (this *QIconEngine) ScaledPixmap(size *QSize, mode QIcon__Mode, state QIcon__State, scale float64) *QPixmap {
	_goptr := newQPixmap(QIconEngine_ScaledPixmap(this.h, size.cPointer(), (int)(mode), (int)(state), (double)(scale)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QIconEngine) VirtualHook(id int, data unsafe.Pointer) {
	QIconEngine_VirtualHook(this.h, (int)(id), data)
}
func (this *QIconEngine) OnPaint(slot func(painter *QPainter, rect *QRect, mode QIcon__Mode, state QIcon__State)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_Paint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_Paint
func miqt_exec_callback_QIconEngine_Paint(self QIconEngine, cb intptr_t, painter *QPainter, rect *QRect, mode int, state int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(painter *QPainter, rect *QRect, mode QIcon__Mode, state QIcon__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQRect(rect)

	slotval3 := (QIcon__Mode)(mode)

	slotval4 := (QIcon__State)(state)

	gofunc(slotval1, slotval2, slotval3, slotval4)

}

func (this *QIconEngine) callVirtualBase_ActualSize(size *QSize, mode QIcon__Mode, state QIcon__State) *QSize {

	_goptr := newQSize(QIconEngine_virtualbase_ActualSize(unsafe.Pointer(this.h), size.cPointer(), (int)(mode), (int)(state)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QIconEngine) OnActualSize(slot func(super func(size *QSize, mode QIcon__Mode, state QIcon__State) *QSize, size *QSize, mode QIcon__Mode, state QIcon__State) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_ActualSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_ActualSize
func miqt_exec_callback_QIconEngine_ActualSize(self QIconEngine, cb intptr_t, size *QSize, mode int, state int) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(size *QSize, mode QIcon__Mode, state QIcon__State) *QSize, size *QSize, mode QIcon__Mode, state QIcon__State) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSize(size)

	slotval2 := (QIcon__Mode)(mode)

	slotval3 := (QIcon__State)(state)

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_ActualSize, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QIconEngine) callVirtualBase_Pixmap(size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap {

	_goptr := newQPixmap(QIconEngine_virtualbase_Pixmap(unsafe.Pointer(this.h), size.cPointer(), (int)(mode), (int)(state)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QIconEngine) OnPixmap(slot func(super func(size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap, size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_Pixmap(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_Pixmap
func miqt_exec_callback_QIconEngine_Pixmap(self QIconEngine, cb intptr_t, size *QSize, mode int, state int) *QPixmap {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap, size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSize(size)

	slotval2 := (QIcon__Mode)(mode)

	slotval3 := (QIcon__State)(state)

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_Pixmap, slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}

func (this *QIconEngine) callVirtualBase_AddPixmap(pixmap *QPixmap, mode QIcon__Mode, state QIcon__State) {

	QIconEngine_virtualbase_AddPixmap(unsafe.Pointer(this.h), pixmap.cPointer(), (int)(mode), (int)(state))

}
func (this *QIconEngine) OnAddPixmap(slot func(super func(pixmap *QPixmap, mode QIcon__Mode, state QIcon__State), pixmap *QPixmap, mode QIcon__Mode, state QIcon__State)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_AddPixmap(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_AddPixmap
func miqt_exec_callback_QIconEngine_AddPixmap(self QIconEngine, cb intptr_t, pixmap *QPixmap, mode int, state int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pixmap *QPixmap, mode QIcon__Mode, state QIcon__State), pixmap *QPixmap, mode QIcon__Mode, state QIcon__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPixmap(pixmap)

	slotval2 := (QIcon__Mode)(mode)

	slotval3 := (QIcon__State)(state)

	gofunc((&QIconEngine{h: self}).callVirtualBase_AddPixmap, slotval1, slotval2, slotval3)

}

func (this *QIconEngine) callVirtualBase_AddFile(fileName string, size *QSize, mode QIcon__Mode, state QIcon__State) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	QIconEngine_virtualbase_AddFile(unsafe.Pointer(this.h), fileName_ms, size.cPointer(), (int)(mode), (int)(state))

}
func (this *QIconEngine) OnAddFile(slot func(super func(fileName string, size *QSize, mode QIcon__Mode, state QIcon__State), fileName string, size *QSize, mode QIcon__Mode, state QIcon__State)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_AddFile(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_AddFile
func miqt_exec_callback_QIconEngine_AddFile(self QIconEngine, cb intptr_t, fileName struct_miqt_string, size *QSize, mode int, state int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(fileName string, size *QSize, mode QIcon__Mode, state QIcon__State), fileName string, size *QSize, mode QIcon__Mode, state QIcon__State))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var fileName_ms struct_miqt_string = fileName
	fileName_ret := GoStringN(fileName_ms.data, int(int64(fileName_ms.len)))
	free(unsafe.Pointer(fileName_ms.data))
	slotval1 := fileName_ret
	slotval2 := newQSize(size)

	slotval3 := (QIcon__Mode)(mode)

	slotval4 := (QIcon__State)(state)

	gofunc((&QIconEngine{h: self}).callVirtualBase_AddFile, slotval1, slotval2, slotval3, slotval4)

}

func (this *QIconEngine) callVirtualBase_Key() string {

	var _ms struct_miqt_string = QIconEngine_virtualbase_Key(unsafe.Pointer(this.h))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QIconEngine) OnKey(slot func(super func() string) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_Key(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_Key
func miqt_exec_callback_QIconEngine_Key(self QIconEngine, cb intptr_t) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() string) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_Key)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}
func (this *QIconEngine) OnClone(slot func() *QIconEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_Clone(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_Clone
func miqt_exec_callback_QIconEngine_Clone(self QIconEngine, cb intptr_t) *QIconEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func() *QIconEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return virtualReturn.cPointer()

}

func (this *QIconEngine) callVirtualBase_Read(in *QDataStream) bool {

	return (bool)(QIconEngine_virtualbase_Read(unsafe.Pointer(this.h), in.cPointer()))

}
func (this *QIconEngine) OnRead(slot func(super func(in *QDataStream) bool, in *QDataStream) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_Read(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_Read
func miqt_exec_callback_QIconEngine_Read(self QIconEngine, cb intptr_t, in *QDataStream) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(in *QDataStream) bool, in *QDataStream) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDataStream(in)

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_Read, slotval1)

	return (bool)(virtualReturn)

}

func (this *QIconEngine) callVirtualBase_Write(out *QDataStream) bool {

	return (bool)(QIconEngine_virtualbase_Write(unsafe.Pointer(this.h), out.cPointer()))

}
func (this *QIconEngine) OnWrite(slot func(super func(out *QDataStream) bool, out *QDataStream) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_Write(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_Write
func miqt_exec_callback_QIconEngine_Write(self QIconEngine, cb intptr_t, out *QDataStream) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(out *QDataStream) bool, out *QDataStream) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDataStream(out)

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_Write, slotval1)

	return (bool)(virtualReturn)

}

func (this *QIconEngine) callVirtualBase_AvailableSizes(mode QIcon__Mode, state QIcon__State) []QSize {

	var _ma struct_miqt_array = QIconEngine_virtualbase_AvailableSizes(unsafe.Pointer(this.h), (int)(mode), (int)(state))
	_ret := make([]QSize, int(_ma.len))
	_outCast := (*[0xffff]*QSize)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQSize(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QIconEngine) OnAvailableSizes(slot func(super func(mode QIcon__Mode, state QIcon__State) []QSize, mode QIcon__Mode, state QIcon__State) []QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_AvailableSizes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_AvailableSizes
func miqt_exec_callback_QIconEngine_AvailableSizes(self QIconEngine, cb intptr_t, mode int, state int) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(mode QIcon__Mode, state QIcon__State) []QSize, mode QIcon__Mode, state QIcon__State) []QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QIcon__Mode)(mode)

	slotval2 := (QIcon__State)(state)

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_AvailableSizes, slotval1, slotval2)
	virtualReturn_CArray := (*[0xffff]*QSize)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QIconEngine) callVirtualBase_IconName() string {

	var _ms struct_miqt_string = QIconEngine_virtualbase_IconName(unsafe.Pointer(this.h))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QIconEngine) OnIconName(slot func(super func() string) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_IconName(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_IconName
func miqt_exec_callback_QIconEngine_IconName(self QIconEngine, cb intptr_t) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() string) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_IconName)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QIconEngine) callVirtualBase_IsNull() bool {

	return (bool)(QIconEngine_virtualbase_IsNull(unsafe.Pointer(this.h)))

}
func (this *QIconEngine) OnIsNull(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_IsNull(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_IsNull
func miqt_exec_callback_QIconEngine_IsNull(self QIconEngine, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_IsNull)

	return (bool)(virtualReturn)

}

func (this *QIconEngine) callVirtualBase_ScaledPixmap(size *QSize, mode QIcon__Mode, state QIcon__State, scale float64) *QPixmap {

	_goptr := newQPixmap(QIconEngine_virtualbase_ScaledPixmap(unsafe.Pointer(this.h), size.cPointer(), (int)(mode), (int)(state), (double)(scale)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QIconEngine) OnScaledPixmap(slot func(super func(size *QSize, mode QIcon__Mode, state QIcon__State, scale float64) *QPixmap, size *QSize, mode QIcon__Mode, state QIcon__State, scale float64) *QPixmap) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_ScaledPixmap(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_ScaledPixmap
func miqt_exec_callback_QIconEngine_ScaledPixmap(self QIconEngine, cb intptr_t, size *QSize, mode int, state int, scale double) *QPixmap {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(size *QSize, mode QIcon__Mode, state QIcon__State, scale float64) *QPixmap, size *QSize, mode QIcon__Mode, state QIcon__State, scale float64) *QPixmap)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSize(size)

	slotval2 := (QIcon__Mode)(mode)

	slotval3 := (QIcon__State)(state)

	slotval4 := (float64)(scale)

	virtualReturn := gofunc((&QIconEngine{h: self}).callVirtualBase_ScaledPixmap, slotval1, slotval2, slotval3, slotval4)

	return virtualReturn.cPointer()

}

func (this *QIconEngine) callVirtualBase_VirtualHook(id int, data unsafe.Pointer) {

	QIconEngine_virtualbase_VirtualHook(unsafe.Pointer(this.h), (int)(id), data)

}
func (this *QIconEngine) OnVirtualHook(slot func(super func(id int, data unsafe.Pointer), id int, data unsafe.Pointer)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEngine_override_virtual_VirtualHook(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEngine_VirtualHook
func miqt_exec_callback_QIconEngine_VirtualHook(self QIconEngine, cb intptr_t, id int, data unsafe.Pointer) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(id int, data unsafe.Pointer), id int, data unsafe.Pointer))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	slotval2 := (unsafe.Pointer)(data)

	gofunc((&QIconEngine{h: self}).callVirtualBase_VirtualHook, slotval1, slotval2)

}

type QIconEngine__ScaledPixmapArgument struct {
	h          uintptr
	isSubclass bool
}

// NewQIconEngine__ScaledPixmapArgument constructs a new QIconEngine::ScaledPixmapArgument object.
func NewQIconEngine__ScaledPixmapArgument(param1 *ScaledPixmapArgument) *QIconEngine__ScaledPixmapArgument {

	ret := newQIconEngine__ScaledPixmapArgument(QIconEngine__ScaledPixmapArgument_new(param1))
	ret.isSubclass = true
	return ret
}

func (this *QIconEngine__ScaledPixmapArgument) OperatorAssign(param1 *ScaledPixmapArgument) {
	QIconEngine__ScaledPixmapArgument_OperatorAssign(this.h, param1)
}
