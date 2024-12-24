package qt

import (
	"unsafe"
)

type QBackingStore struct {
	h          uintptr
	isSubclass bool
}

// NewQBackingStore constructs a new QBackingStore object.
func NewQBackingStore(window *QWindow) *QBackingStore {
	ret := newQBackingStore(QBackingStore_new(window.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QBackingStore) Window() *QWindow {
	return newQWindow(QBackingStore_Window(this.h))
}

func (this *QBackingStore) PaintDevice() *QPaintDevice {
	return newQPaintDevice(QBackingStore_PaintDevice(this.h))
}

func (this *QBackingStore) Flush(region *QRegion) {
	QBackingStore_Flush(this.h, region.cPointer())
}

func (this *QBackingStore) Resize(size *QSize) {
	QBackingStore_Resize(this.h, size.cPointer())
}

func (this *QBackingStore) Size() *QSize {
	_goptr := newQSize(QBackingStore_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBackingStore) Scroll(area *QRegion, dx int, dy int) bool {
	return (bool)(QBackingStore_Scroll(this.h, area.cPointer(), (int)(dx), (int)(dy)))
}

func (this *QBackingStore) BeginPaint(param1 *QRegion) {
	QBackingStore_BeginPaint(this.h, param1.cPointer())
}

func (this *QBackingStore) EndPaint() {
	QBackingStore_EndPaint(this.h)
}

func (this *QBackingStore) SetStaticContents(region *QRegion) {
	QBackingStore_SetStaticContents(this.h, region.cPointer())
}

func (this *QBackingStore) StaticContents() *QRegion {
	_goptr := newQRegion(QBackingStore_StaticContents(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBackingStore) HasStaticContents() bool {
	return (bool)(QBackingStore_HasStaticContents(this.h))
}

func (this *QBackingStore) Flush2(region *QRegion, window *QWindow) {
	QBackingStore_Flush2(this.h, region.cPointer(), window.cPointer())
}

func (this *QBackingStore) Flush3(region *QRegion, window *QWindow, offset *QPoint) {
	QBackingStore_Flush3(this.h, region.cPointer(), window.cPointer(), offset.cPointer())
}
