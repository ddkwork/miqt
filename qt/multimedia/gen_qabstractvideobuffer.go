package multimedia

import (
	"unsafe"
)

type QAbstractVideoBuffer struct {
	h          uintptr
	isSubclass bool
}

func (this *QAbstractVideoBuffer) Map(mode QVideoFrame__MapMode) MapData {
	xxxxxxxxx
}

func (this *QAbstractVideoBuffer) Unmap() {
	QAbstractVideoBuffer_Unmap(this.h)
}

func (this *QAbstractVideoBuffer) Format() *QVideoFrameFormat {
	_goptr := newQVideoFrameFormat(QAbstractVideoBuffer_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QAbstractVideoBuffer__MapData struct {
	h          uintptr
	isSubclass bool
}
