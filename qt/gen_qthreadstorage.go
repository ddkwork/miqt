package qt

import (
	"unsafe"
)

type QThreadStorageData struct {
	h          uintptr
	isSubclass bool
}
