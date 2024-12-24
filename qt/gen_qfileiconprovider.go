package qt

import (
	"unsafe"
)

type QFileIconProvider struct {
	h          uintptr
	isSubclass bool
}

// NewQFileIconProvider constructs a new QFileIconProvider object.
func NewQFileIconProvider() *QFileIconProvider {
	ret := newQFileIconProvider(QFileIconProvider_new())
	ret.isSubclass = true
	return ret
}

func (this *QFileIconProvider) Icon(typeVal IconType) *QIcon {
	_goptr := newQIcon(QFileIconProvider_Icon(this.h, typeVal))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileIconProvider) IconWithInfo(info *QFileInfo) *QIcon {
	_goptr := newQIcon(QFileIconProvider_IconWithInfo(this.h, info.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
