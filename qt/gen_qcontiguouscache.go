package qt

import (
	"unsafe"
)

type QContiguousCacheData struct {
	h          uintptr
	isSubclass bool
}

func QContiguousCacheData_AllocateData(size int64, alignment int64) *QContiguousCacheData {
	return newQContiguousCacheData(QContiguousCacheData_AllocateData((ptrdiff_t)(size), (ptrdiff_t)(alignment)))
}

func QContiguousCacheData_FreeData(data *QContiguousCacheData) {
	QContiguousCacheData_FreeData(data.cPointer())
}
