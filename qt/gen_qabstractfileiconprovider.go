package qt

import (
	"unsafe"
)

type QAbstractFileIconProvider__IconType int

const (
	QAbstractFileIconProvider__Computer QAbstractFileIconProvider__IconType = 0
	QAbstractFileIconProvider__Desktop  QAbstractFileIconProvider__IconType = 1
	QAbstractFileIconProvider__Trashcan QAbstractFileIconProvider__IconType = 2
	QAbstractFileIconProvider__Network  QAbstractFileIconProvider__IconType = 3
	QAbstractFileIconProvider__Drive    QAbstractFileIconProvider__IconType = 4
	QAbstractFileIconProvider__Folder   QAbstractFileIconProvider__IconType = 5
	QAbstractFileIconProvider__File     QAbstractFileIconProvider__IconType = 6
)

type QAbstractFileIconProvider__Option int

const (
	QAbstractFileIconProvider__DontUseCustomDirectoryIcons QAbstractFileIconProvider__Option = 1
)

type QAbstractFileIconProvider struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractFileIconProvider constructs a new QAbstractFileIconProvider object.
func NewQAbstractFileIconProvider() *QAbstractFileIconProvider {
	g := newQAbstractFileIconProvider(QAbstractFileIconProvider_new())
	g.isSubclass = true
	return g
}

func (this *QAbstractFileIconProvider) Icon(param1 IconType) *QIcon {
	_goptr := newQIcon(QAbstractFileIconProvider_Icon(this.h, param1))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractFileIconProvider) IconWithQFileInfo(param1 *QFileInfo) *QIcon {
	_goptr := newQIcon(QAbstractFileIconProvider_IconWithQFileInfo(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractFileIconProvider) Type(param1 *QFileInfo) string {
	var _ms struct_miqt_string = QAbstractFileIconProvider_Type(this.h, param1.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractFileIconProvider) SetOptions(options Options) {
	QAbstractFileIconProvider_SetOptions(this.h, options)
}

func (this *QAbstractFileIconProvider) Options() Options {
	xxxxxxxxx
}
