package qt

import (
	"unsafe"
)

type QScopedPointerPodDeleter struct {
	h          uintptr
	isSubclass bool
}

func QScopedPointerPodDeleter_Cleanup(pointer unsafe.Pointer) {
	QScopedPointerPodDeleter_Cleanup(pointer)
}

func (this *QScopedPointerPodDeleter) OperatorCall(pointer unsafe.Pointer) {
	QScopedPointerPodDeleter_OperatorCall(this.h, pointer)
}
