package webengine

/*

#include "gen_qwebengineurlscheme.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QWebEngineUrlScheme__Syntax int

const (
	QWebEngineUrlScheme__HostPortAndUserInformation QWebEngineUrlScheme__Syntax = 0
	QWebEngineUrlScheme__HostAndPort                QWebEngineUrlScheme__Syntax = 1
	QWebEngineUrlScheme__Host                       QWebEngineUrlScheme__Syntax = 2
	QWebEngineUrlScheme__Path                       QWebEngineUrlScheme__Syntax = 3
)

type QWebEngineUrlScheme__SpecialPort int

const (
	QWebEngineUrlScheme__PortUnspecified QWebEngineUrlScheme__SpecialPort = -1
)

type QWebEngineUrlScheme__Flag int

const (
	QWebEngineUrlScheme__SecureScheme                 QWebEngineUrlScheme__Flag = 1
	QWebEngineUrlScheme__LocalScheme                  QWebEngineUrlScheme__Flag = 2
	QWebEngineUrlScheme__LocalAccessAllowed           QWebEngineUrlScheme__Flag = 4
	QWebEngineUrlScheme__NoAccessAllowed              QWebEngineUrlScheme__Flag = 8
	QWebEngineUrlScheme__ServiceWorkersAllowed        QWebEngineUrlScheme__Flag = 16
	QWebEngineUrlScheme__ViewSourceAllowed            QWebEngineUrlScheme__Flag = 32
	QWebEngineUrlScheme__ContentSecurityPolicyIgnored QWebEngineUrlScheme__Flag = 64
	QWebEngineUrlScheme__CorsEnabled                  QWebEngineUrlScheme__Flag = 128
)

type QWebEngineUrlScheme struct {
	h *C.QWebEngineUrlScheme
}

func (this *QWebEngineUrlScheme) cPointer() *C.QWebEngineUrlScheme {
	if this == nil {
		return nil
	}
	return this.h
}

func (this *QWebEngineUrlScheme) UnsafePointer() unsafe.Pointer {
	if this == nil {
		return nil
	}
	return unsafe.Pointer(this.h)
}

// newQWebEngineUrlScheme constructs the type using only CGO pointers.
func newQWebEngineUrlScheme(h *C.QWebEngineUrlScheme) *QWebEngineUrlScheme {
	if h == nil {
		return nil
	}

	return &QWebEngineUrlScheme{h: h}
}

// UnsafeNewQWebEngineUrlScheme constructs the type using only unsafe pointers.
func UnsafeNewQWebEngineUrlScheme(h unsafe.Pointer) *QWebEngineUrlScheme {
	return newQWebEngineUrlScheme((*C.QWebEngineUrlScheme)(h))
}

// NewQWebEngineUrlScheme constructs a new QWebEngineUrlScheme object.
func NewQWebEngineUrlScheme() *QWebEngineUrlScheme {

	return newQWebEngineUrlScheme(C.QWebEngineUrlScheme_new())
}

// NewQWebEngineUrlScheme2 constructs a new QWebEngineUrlScheme object.
func NewQWebEngineUrlScheme2(name []byte) *QWebEngineUrlScheme {
	name_alias := C.struct_miqt_string{}
	if len(name) > 0 {
		name_alias.data = (*C.char)(unsafe.Pointer(&name[0]))
	} else {
		name_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	name_alias.len = C.size_t(len(name))

	return newQWebEngineUrlScheme(C.QWebEngineUrlScheme_new2(name_alias))
}

// NewQWebEngineUrlScheme3 constructs a new QWebEngineUrlScheme object.
func NewQWebEngineUrlScheme3(that *QWebEngineUrlScheme) *QWebEngineUrlScheme {

	return newQWebEngineUrlScheme(C.QWebEngineUrlScheme_new3(that.cPointer()))
}

func (this *QWebEngineUrlScheme) OperatorAssign(that *QWebEngineUrlScheme) {
	C.QWebEngineUrlScheme_operatorAssign(this.h, that.cPointer())
}

func (this *QWebEngineUrlScheme) OperatorEqual(that *QWebEngineUrlScheme) bool {
	return (bool)(C.QWebEngineUrlScheme_operatorEqual(this.h, that.cPointer()))
}

func (this *QWebEngineUrlScheme) OperatorNotEqual(that *QWebEngineUrlScheme) bool {
	return (bool)(C.QWebEngineUrlScheme_operatorNotEqual(this.h, that.cPointer()))
}

func (this *QWebEngineUrlScheme) Name() []byte {
	var _bytearray C.struct_miqt_string = C.QWebEngineUrlScheme_name(this.h)
	_ret := C.GoBytes(unsafe.Pointer(_bytearray.data), C.int(int64(_bytearray.len)))
	C.free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QWebEngineUrlScheme) SetName(newValue []byte) {
	newValue_alias := C.struct_miqt_string{}
	if len(newValue) > 0 {
		newValue_alias.data = (*C.char)(unsafe.Pointer(&newValue[0]))
	} else {
		newValue_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	newValue_alias.len = C.size_t(len(newValue))
	C.QWebEngineUrlScheme_setName(this.h, newValue_alias)
}

func (this *QWebEngineUrlScheme) Syntax() QWebEngineUrlScheme__Syntax {
	return (QWebEngineUrlScheme__Syntax)(C.QWebEngineUrlScheme_syntax(this.h))
}

func (this *QWebEngineUrlScheme) SetSyntax(newValue QWebEngineUrlScheme__Syntax) {
	C.QWebEngineUrlScheme_setSyntax(this.h, (C.int)(newValue))
}

func (this *QWebEngineUrlScheme) DefaultPort() int {
	return (int)(C.QWebEngineUrlScheme_defaultPort(this.h))
}

func (this *QWebEngineUrlScheme) SetDefaultPort(newValue int) {
	C.QWebEngineUrlScheme_setDefaultPort(this.h, (C.int)(newValue))
}

func (this *QWebEngineUrlScheme) Flags() QWebEngineUrlScheme__Flag {
	return (QWebEngineUrlScheme__Flag)(C.QWebEngineUrlScheme_flags(this.h))
}

func (this *QWebEngineUrlScheme) SetFlags(newValue QWebEngineUrlScheme__Flag) {
	C.QWebEngineUrlScheme_setFlags(this.h, (C.int)(newValue))
}

func QWebEngineUrlScheme_RegisterScheme(scheme *QWebEngineUrlScheme) {
	C.QWebEngineUrlScheme_registerScheme(scheme.cPointer())
}

func QWebEngineUrlScheme_SchemeByName(name []byte) *QWebEngineUrlScheme {
	name_alias := C.struct_miqt_string{}
	if len(name) > 0 {
		name_alias.data = (*C.char)(unsafe.Pointer(&name[0]))
	} else {
		name_alias.data = (*C.char)(unsafe.Pointer(nil))
	}
	name_alias.len = C.size_t(len(name))
	_goptr := newQWebEngineUrlScheme(C.QWebEngineUrlScheme_schemeByName(name_alias))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

// Delete this object from C++ memory.
func (this *QWebEngineUrlScheme) Delete() {
	C.QWebEngineUrlScheme_delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QWebEngineUrlScheme) GoGC() {
	runtime.SetFinalizer(this, func(this *QWebEngineUrlScheme) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
