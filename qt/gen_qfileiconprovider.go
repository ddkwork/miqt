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

func (this *QFileIconProvider) callVirtualBase_Icon(typeVal IconType) *QIcon {

	_goptr := newQIcon(QFileIconProvider_virtualbase_Icon(unsafe.Pointer(this.h), typeVal))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFileIconProvider) OnIcon(slot func(super func(typeVal IconType) *QIcon, typeVal IconType) *QIcon) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileIconProvider_override_virtual_Icon(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileIconProvider_Icon
func miqt_exec_callback_QFileIconProvider_Icon(self QFileIconProvider, cb intptr_t, typeVal IconType) *QIcon {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(typeVal IconType) *QIcon, typeVal IconType) *QIcon)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QFileIconProvider{h: self}).callVirtualBase_Icon, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFileIconProvider) callVirtualBase_IconWithInfo(info *QFileInfo) *QIcon {

	_goptr := newQIcon(QFileIconProvider_virtualbase_IconWithInfo(unsafe.Pointer(this.h), info.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFileIconProvider) OnIconWithInfo(slot func(super func(info *QFileInfo) *QIcon, info *QFileInfo) *QIcon) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileIconProvider_override_virtual_IconWithInfo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileIconProvider_IconWithInfo
func miqt_exec_callback_QFileIconProvider_IconWithInfo(self QFileIconProvider, cb intptr_t, info *QFileInfo) *QIcon {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(info *QFileInfo) *QIcon, info *QFileInfo) *QIcon)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFileInfo(info)

	virtualReturn := gofunc((&QFileIconProvider{h: self}).callVirtualBase_IconWithInfo, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFileIconProvider) callVirtualBase_Type(param1 *QFileInfo) string {

	var _ms struct_miqt_string = QFileIconProvider_virtualbase_Type(unsafe.Pointer(this.h), param1.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QFileIconProvider) OnType(slot func(super func(param1 *QFileInfo) string, param1 *QFileInfo) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileIconProvider_override_virtual_Type(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileIconProvider_Type
func miqt_exec_callback_QFileIconProvider_Type(self QFileIconProvider, cb intptr_t, param1 *QFileInfo) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QFileInfo) string, param1 *QFileInfo) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFileInfo(param1)

	virtualReturn := gofunc((&QFileIconProvider{h: self}).callVirtualBase_Type, slotval1)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QFileIconProvider) callVirtualBase_SetOptions(options Options) {

	QFileIconProvider_virtualbase_SetOptions(unsafe.Pointer(this.h), options)

}
func (this *QFileIconProvider) OnSetOptions(slot func(super func(options Options), options Options)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileIconProvider_override_virtual_SetOptions(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileIconProvider_SetOptions
func miqt_exec_callback_QFileIconProvider_SetOptions(self QFileIconProvider, cb intptr_t, options Options) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(options Options), options Options))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QFileIconProvider{h: self}).callVirtualBase_SetOptions, slotval1)

}

func (this *QFileIconProvider) callVirtualBase_Options() Options {

	xxxxxxxxx
}
func (this *QFileIconProvider) OnOptions(slot func(super func() Options) Options) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileIconProvider_override_virtual_Options(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileIconProvider_Options
func miqt_exec_callback_QFileIconProvider_Options(self QFileIconProvider, cb intptr_t) Options {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() Options) Options)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFileIconProvider{h: self}).callVirtualBase_Options)

	return virtualReturn

}
