package qt

import (
	"unsafe"
)

type QPicture struct {
	h          uintptr
	isSubclass bool
}

// NewQPicture constructs a new QPicture object.
func NewQPicture() *QPicture {
	ret := newQPicture(QPicture_new())
	ret.isSubclass = true
	return ret
}

// NewQPicture2 constructs a new QPicture object.
func NewQPicture2(param1 *QPicture) *QPicture {
	ret := newQPicture(QPicture_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPicture3 constructs a new QPicture object.
func NewQPicture3(formatVersion int) *QPicture {
	ret := newQPicture(QPicture_new3((int)(formatVersion)))
	ret.isSubclass = true
	return ret
}

func (this *QPicture) IsNull() bool {
	return (bool)(QPicture_IsNull(this.h))
}

func (this *QPicture) DevType() int {
	return (int)(QPicture_DevType(this.h))
}

func (this *QPicture) Size() uint {
	return (uint)(QPicture_Size(this.h))
}

func (this *QPicture) Data() string {
	_ret := QPicture_Data(this.h)
	return GoString(_ret)
}

func (this *QPicture) SetData(data string, size uint) {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	QPicture_SetData(this.h, data_Cstring, (uint)(size))
}

func (this *QPicture) Play(p *QPainter) bool {
	return (bool)(QPicture_Play(this.h, p.cPointer()))
}

func (this *QPicture) Load(dev *QIODevice) bool {
	return (bool)(QPicture_Load(this.h, dev.cPointer()))
}

func (this *QPicture) LoadWithFileName(fileName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QPicture_LoadWithFileName(this.h, fileName_ms))
}

func (this *QPicture) Save(dev *QIODevice) bool {
	return (bool)(QPicture_Save(this.h, dev.cPointer()))
}

func (this *QPicture) SaveWithFileName(fileName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QPicture_SaveWithFileName(this.h, fileName_ms))
}

func (this *QPicture) BoundingRect() *QRect {
	_goptr := newQRect(QPicture_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPicture) SetBoundingRect(r *QRect) {
	QPicture_SetBoundingRect(this.h, r.cPointer())
}

func (this *QPicture) OperatorAssign(p *QPicture) {
	QPicture_OperatorAssign(this.h, p.cPointer())
}

func (this *QPicture) Swap(other *QPicture) {
	QPicture_Swap(this.h, other.cPointer())
}

func (this *QPicture) Detach() {
	QPicture_Detach(this.h)
}

func (this *QPicture) IsDetached() bool {
	return (bool)(QPicture_IsDetached(this.h))
}

func (this *QPicture) PaintEngine() *QPaintEngine {
	return newQPaintEngine(QPicture_PaintEngine(this.h))
}

func (this *QPicture) DataPtr() *DataPtr {
	xxxxxxxxx
}
