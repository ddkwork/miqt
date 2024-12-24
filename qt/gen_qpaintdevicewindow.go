package qt

import (
	"unsafe"
)

type QPaintDeviceWindow struct {
	h          uintptr
	isSubclass bool
}

func (this *QPaintDeviceWindow) MetaObject() *QMetaObject {
	return newQMetaObject(QPaintDeviceWindow_MetaObject(this.h))
}

func (this *QPaintDeviceWindow) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPaintDeviceWindow_Metacast(this.h, param1_Cstring))
}

func QPaintDeviceWindow_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPaintDeviceWindow_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPaintDeviceWindow) Update(rect *QRect) {
	QPaintDeviceWindow_Update(this.h, rect.cPointer())
}

func (this *QPaintDeviceWindow) UpdateWithRegion(region *QRegion) {
	QPaintDeviceWindow_UpdateWithRegion(this.h, region.cPointer())
}

func (this *QPaintDeviceWindow) Update2() {
	QPaintDeviceWindow_Update2(this.h)
}

func QPaintDeviceWindow_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPaintDeviceWindow_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPaintDeviceWindow_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPaintDeviceWindow_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
