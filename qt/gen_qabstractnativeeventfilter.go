package qt

import (
	"unsafe"
)

type QAbstractNativeEventFilter struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractNativeEventFilter constructs a new QAbstractNativeEventFilter object.
func NewQAbstractNativeEventFilter() *QAbstractNativeEventFilter {

	ret := newQAbstractNativeEventFilter(QAbstractNativeEventFilter_new())
	ret.isSubclass = true
	return ret
}

func (this *QAbstractNativeEventFilter) NativeEventFilter(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))
	return (bool)(QAbstractNativeEventFilter_NativeEventFilter(this.h, eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))
}
func (this *QAbstractNativeEventFilter) OnNativeEventFilter(slot func(eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractNativeEventFilter_override_virtual_NativeEventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractNativeEventFilter_NativeEventFilter
func miqt_exec_callback_QAbstractNativeEventFilter_NativeEventFilter(self QAbstractNativeEventFilter, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc(slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}
