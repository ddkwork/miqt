package network

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QHstsPolicy__PolicyFlag int

const (
	QHstsPolicy__IncludeSubDomains QHstsPolicy__PolicyFlag = 1
)

type QHstsPolicy struct {
	h          uintptr
	isSubclass bool
}

// NewQHstsPolicy constructs a new QHstsPolicy object.
func NewQHstsPolicy() *QHstsPolicy {
	g := newQHstsPolicy(QHstsPolicy_new())
	g.isSubclass = true
	return g
}

// NewQHstsPolicy2 constructs a new QHstsPolicy object.
func NewQHstsPolicy2(expiry *qt.QDateTime, flags PolicyFlags, host string) *QHstsPolicy {
	host_ms := struct_miqt_string{}
	host_ms.data = CString(host)
	host_ms.len = size_t(len(host))
	defer free(unsafe.Pointer(host_ms.data))
	g := newQHstsPolicy(QHstsPolicy_new2((*QDateTime)(expiry.UnsafePointer()), flags, host_ms))
	g.isSubclass = true
	return g
}

// NewQHstsPolicy3 constructs a new QHstsPolicy object.
func NewQHstsPolicy3(rhs *QHstsPolicy) *QHstsPolicy {
	g := newQHstsPolicy(QHstsPolicy_new3(rhs.cPointer()))
	g.isSubclass = true
	return g
}

// NewQHstsPolicy4 constructs a new QHstsPolicy object.
func NewQHstsPolicy4(expiry *qt.QDateTime, flags PolicyFlags, host string, mode qt.QUrl__ParsingMode) *QHstsPolicy {
	host_ms := struct_miqt_string{}
	host_ms.data = CString(host)
	host_ms.len = size_t(len(host))
	defer free(unsafe.Pointer(host_ms.data))
	g := newQHstsPolicy(QHstsPolicy_new4((*QDateTime)(expiry.UnsafePointer()), flags, host_ms, (int)(mode)))
	g.isSubclass = true
	return g
}

func (this *QHstsPolicy) OperatorAssign(rhs *QHstsPolicy) {
	QHstsPolicy_OperatorAssign(this.h, rhs.cPointer())
}

func (this *QHstsPolicy) Swap(other *QHstsPolicy) {
	QHstsPolicy_Swap(this.h, other.cPointer())
}

func (this *QHstsPolicy) SetHost(host string) {
	host_ms := struct_miqt_string{}
	host_ms.data = CString(host)
	host_ms.len = size_t(len(host))
	defer free(unsafe.Pointer(host_ms.data))
	QHstsPolicy_SetHost(this.h, host_ms)
}

func (this *QHstsPolicy) Host() string {
	var _ms struct_miqt_string = QHstsPolicy_Host(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QHstsPolicy) SetExpiry(expiry *qt.QDateTime) {
	QHstsPolicy_SetExpiry(this.h, (*QDateTime)(expiry.UnsafePointer()))
}

func (this *QHstsPolicy) Expiry() *qt.QDateTime {
	_goptr := qt.UnsafeNewQDateTime(unsafe.Pointer(QHstsPolicy_Expiry(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QHstsPolicy) SetIncludesSubDomains(include bool) {
	QHstsPolicy_SetIncludesSubDomains(this.h, (bool)(include))
}

func (this *QHstsPolicy) IncludesSubDomains() bool {
	return (bool)(QHstsPolicy_IncludesSubDomains(this.h))
}

func (this *QHstsPolicy) IsExpired() bool {
	return (bool)(QHstsPolicy_IsExpired(this.h))
}

func (this *QHstsPolicy) SetHost2(host string, mode qt.QUrl__ParsingMode) {
	host_ms := struct_miqt_string{}
	host_ms.data = CString(host)
	host_ms.len = size_t(len(host))
	defer free(unsafe.Pointer(host_ms.data))
	QHstsPolicy_SetHost2(this.h, host_ms, (int)(mode))
}

func (this *QHstsPolicy) Host1(options ComponentFormattingOption) string {
	var _ms struct_miqt_string = QHstsPolicy_Host1(this.h, (int)(options))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
