package qt

import (
	"unsafe"
)

type QDesktopServices struct {
	h          uintptr
	isSubclass bool
}

func QDesktopServices_OpenUrl(url *QUrl) bool {
	return (bool)(QDesktopServices_OpenUrl(url.cPointer()))
}

func QDesktopServices_SetUrlHandler(scheme string, receiver *QObject, method string) {
	scheme_ms := struct_miqt_string{}
	scheme_ms.data = CString(scheme)
	scheme_ms.len = size_t(len(scheme))
	defer free(unsafe.Pointer(scheme_ms.data))
	method_Cstring := CString(method)
	defer free(unsafe.Pointer(method_Cstring))
	QDesktopServices_SetUrlHandler(scheme_ms, receiver.cPointer(), method_Cstring)
}

func QDesktopServices_UnsetUrlHandler(scheme string) {
	scheme_ms := struct_miqt_string{}
	scheme_ms.data = CString(scheme)
	scheme_ms.len = size_t(len(scheme))
	defer free(unsafe.Pointer(scheme_ms.data))
	QDesktopServices_UnsetUrlHandler(scheme_ms)
}
