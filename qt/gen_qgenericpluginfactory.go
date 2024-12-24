package qt

import (
	"unsafe"
)

type QGenericPluginFactory struct {
	h          uintptr
	isSubclass bool
}

func QGenericPluginFactory_Keys() []string {
	var _ma struct_miqt_array = QGenericPluginFactory_Keys()
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QGenericPluginFactory_Create(param1 string, param2 string) *QObject {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	param2_ms := struct_miqt_string{}
	param2_ms.data = CString(param2)
	param2_ms.len = size_t(len(param2))
	defer free(unsafe.Pointer(param2_ms.data))
	return newQObject(QGenericPluginFactory_Create(param1_ms, param2_ms))
}
