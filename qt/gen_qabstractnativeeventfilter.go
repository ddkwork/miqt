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
