package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QNetworkCookieJar struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkCookieJar constructs a new QNetworkCookieJar object.
func NewQNetworkCookieJar() *QNetworkCookieJar {
	g := newQNetworkCookieJar(QNetworkCookieJar_new())
	g.isSubclass = true
	return g
}

// NewQNetworkCookieJar2 constructs a new QNetworkCookieJar object.
func NewQNetworkCookieJar2(parent *qt.QObject) *QNetworkCookieJar {
	g := newQNetworkCookieJar(QNetworkCookieJar_new2((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QNetworkCookieJar) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QNetworkCookieJar_MetaObject(this.h)))
}

func (this *QNetworkCookieJar) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QNetworkCookieJar_Metacast(this.h, param1_Cstring))
}

func QNetworkCookieJar_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QNetworkCookieJar_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkCookieJar) CookiesForUrl(url *qt.QUrl) []QNetworkCookie {
	var _ma struct_miqt_array = QNetworkCookieJar_CookiesForUrl(this.h, (*QUrl)(url.UnsafePointer()))
	_ret := make([]QNetworkCookie, int(_ma.len))
	_outCast := (*[0xffff]*QNetworkCookie)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQNetworkCookie(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QNetworkCookieJar) SetCookiesFromUrl(cookieList []QNetworkCookie, url *qt.QUrl) bool {
	cookieList_CArray := (*[0xffff]*QNetworkCookie)(malloc(size_t(8 * len(cookieList))))
	defer free(unsafe.Pointer(cookieList_CArray))
	for i := range cookieList {
		cookieList_CArray[i] = cookieList[i].cPointer()
	}
	cookieList_ma := struct_miqt_array{len: size_t(len(cookieList)), data: unsafe.Pointer(cookieList_CArray)}
	return (bool)(QNetworkCookieJar_SetCookiesFromUrl(this.h, cookieList_ma, (*QUrl)(url.UnsafePointer())))
}

func (this *QNetworkCookieJar) InsertCookie(cookie *QNetworkCookie) bool {
	return (bool)(QNetworkCookieJar_InsertCookie(this.h, cookie.cPointer()))
}

func (this *QNetworkCookieJar) UpdateCookie(cookie *QNetworkCookie) bool {
	return (bool)(QNetworkCookieJar_UpdateCookie(this.h, cookie.cPointer()))
}

func (this *QNetworkCookieJar) DeleteCookie(cookie *QNetworkCookie) bool {
	return (bool)(QNetworkCookieJar_DeleteCookie(this.h, cookie.cPointer()))
}

func QNetworkCookieJar_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QNetworkCookieJar_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkCookieJar_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QNetworkCookieJar_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkCookieJar) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QNetworkCookieJar_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QNetworkCookieJar) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkCookieJar_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkCookieJar_MetaObject
func miqt_exec_callback_QNetworkCookieJar_MetaObject(self QNetworkCookieJar, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QNetworkCookieJar{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QNetworkCookieJar) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QNetworkCookieJar_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QNetworkCookieJar) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkCookieJar_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkCookieJar_Metacast
func miqt_exec_callback_QNetworkCookieJar_Metacast(self QNetworkCookieJar, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QNetworkCookieJar{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
