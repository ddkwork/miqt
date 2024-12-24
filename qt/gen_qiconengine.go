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
	g := newQIconEngine(QIconEngine_new())
	g.isSubclass = true
	return g
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

type QIconEngine__ScaledPixmapArgument struct {
	h          uintptr
	isSubclass bool
}

// NewQIconEngine__ScaledPixmapArgument constructs a new QIconEngine::ScaledPixmapArgument object.
func NewQIconEngine__ScaledPixmapArgument(param1 *ScaledPixmapArgument) *QIconEngine__ScaledPixmapArgument {
	g := newQIconEngine__ScaledPixmapArgument(QIconEngine__ScaledPixmapArgument_new(param1))
	g.isSubclass = true
	return g
}

func (this *QIconEngine__ScaledPixmapArgument) OperatorAssign(param1 *ScaledPixmapArgument) {
	QIconEngine__ScaledPixmapArgument_OperatorAssign(this.h, param1)
}
