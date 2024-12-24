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

	ret := newQAbstractFileIconProvider(QAbstractFileIconProvider_new())
	ret.isSubclass = true
	return ret
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

func (this *QAbstractFileIconProvider) callVirtualBase_Icon(param1 IconType) *QIcon {

	_goptr := newQIcon(QAbstractFileIconProvider_virtualbase_Icon(unsafe.Pointer(this.h), param1))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractFileIconProvider) OnIcon(slot func(super func(param1 IconType) *QIcon, param1 IconType) *QIcon) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractFileIconProvider_override_virtual_Icon(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractFileIconProvider_Icon
func miqt_exec_callback_QAbstractFileIconProvider_Icon(self QAbstractFileIconProvider, cb intptr_t, param1 IconType) *QIcon {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 IconType) *QIcon, param1 IconType) *QIcon)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QAbstractFileIconProvider{h: self}).callVirtualBase_Icon, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractFileIconProvider) callVirtualBase_IconWithQFileInfo(param1 *QFileInfo) *QIcon {

	_goptr := newQIcon(QAbstractFileIconProvider_virtualbase_IconWithQFileInfo(unsafe.Pointer(this.h), param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractFileIconProvider) OnIconWithQFileInfo(slot func(super func(param1 *QFileInfo) *QIcon, param1 *QFileInfo) *QIcon) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractFileIconProvider_override_virtual_IconWithQFileInfo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractFileIconProvider_IconWithQFileInfo
func miqt_exec_callback_QAbstractFileIconProvider_IconWithQFileInfo(self QAbstractFileIconProvider, cb intptr_t, param1 *QFileInfo) *QIcon {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QFileInfo) *QIcon, param1 *QFileInfo) *QIcon)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFileInfo(param1)

	virtualReturn := gofunc((&QAbstractFileIconProvider{h: self}).callVirtualBase_IconWithQFileInfo, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractFileIconProvider) callVirtualBase_Type(param1 *QFileInfo) string {

	var _ms struct_miqt_string = QAbstractFileIconProvider_virtualbase_Type(unsafe.Pointer(this.h), param1.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QAbstractFileIconProvider) OnType(slot func(super func(param1 *QFileInfo) string, param1 *QFileInfo) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractFileIconProvider_override_virtual_Type(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractFileIconProvider_Type
func miqt_exec_callback_QAbstractFileIconProvider_Type(self QAbstractFileIconProvider, cb intptr_t, param1 *QFileInfo) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QFileInfo) string, param1 *QFileInfo) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFileInfo(param1)

	virtualReturn := gofunc((&QAbstractFileIconProvider{h: self}).callVirtualBase_Type, slotval1)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QAbstractFileIconProvider) callVirtualBase_SetOptions(options Options) {

	QAbstractFileIconProvider_virtualbase_SetOptions(unsafe.Pointer(this.h), options)

}
func (this *QAbstractFileIconProvider) OnSetOptions(slot func(super func(options Options), options Options)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractFileIconProvider_override_virtual_SetOptions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractFileIconProvider_SetOptions
func miqt_exec_callback_QAbstractFileIconProvider_SetOptions(self QAbstractFileIconProvider, cb intptr_t, options Options) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(options Options), options Options))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QAbstractFileIconProvider{h: self}).callVirtualBase_SetOptions, slotval1)

}

func (this *QAbstractFileIconProvider) callVirtualBase_Options() Options {

	xxxxxxxxx
}
func (this *QAbstractFileIconProvider) OnOptions(slot func(super func() Options) Options) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractFileIconProvider_override_virtual_Options(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractFileIconProvider_Options
func miqt_exec_callback_QAbstractFileIconProvider_Options(self QAbstractFileIconProvider, cb intptr_t) Options {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Options) Options)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractFileIconProvider{h: self}).callVirtualBase_Options)

	return virtualReturn

}
