package qt

import (
	"unsafe"
)

type QUntypedPropertyData struct {
	h          uintptr
	isSubclass bool
}

type QPropertyProxyBindingData struct {
	h          uintptr
	isSubclass bool
}
