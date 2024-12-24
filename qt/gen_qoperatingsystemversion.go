package qt

import (
	"unsafe"
)

type QOperatingSystemVersionBase__OSType int

const (
	QOperatingSystemVersionBase__Unknown  QOperatingSystemVersionBase__OSType = 0
	QOperatingSystemVersionBase__Windows  QOperatingSystemVersionBase__OSType = 1
	QOperatingSystemVersionBase__MacOS    QOperatingSystemVersionBase__OSType = 2
	QOperatingSystemVersionBase__IOS      QOperatingSystemVersionBase__OSType = 3
	QOperatingSystemVersionBase__TvOS     QOperatingSystemVersionBase__OSType = 4
	QOperatingSystemVersionBase__WatchOS  QOperatingSystemVersionBase__OSType = 5
	QOperatingSystemVersionBase__Android  QOperatingSystemVersionBase__OSType = 6
	QOperatingSystemVersionBase__VisionOS QOperatingSystemVersionBase__OSType = 7
)

type QOperatingSystemVersion__OSType int

const (
	QOperatingSystemVersion__Unknown  QOperatingSystemVersion__OSType = 0
	QOperatingSystemVersion__Windows  QOperatingSystemVersion__OSType = 1
	QOperatingSystemVersion__MacOS    QOperatingSystemVersion__OSType = 2
	QOperatingSystemVersion__IOS      QOperatingSystemVersion__OSType = 3
	QOperatingSystemVersion__TvOS     QOperatingSystemVersion__OSType = 4
	QOperatingSystemVersion__WatchOS  QOperatingSystemVersion__OSType = 5
	QOperatingSystemVersion__Android  QOperatingSystemVersion__OSType = 6
	QOperatingSystemVersion__VisionOS QOperatingSystemVersion__OSType = 7
)

type QOperatingSystemVersionBase struct {
	h          uintptr
	isSubclass bool
}

// NewQOperatingSystemVersionBase constructs a new QOperatingSystemVersionBase object.
func NewQOperatingSystemVersionBase(osType OSType, vmajor int) *QOperatingSystemVersionBase {
	g := newQOperatingSystemVersionBase(QOperatingSystemVersionBase_new(osType, (int)(vmajor)))
	g.isSubclass = true
	return g
}

// NewQOperatingSystemVersionBase2 constructs a new QOperatingSystemVersionBase object.
func NewQOperatingSystemVersionBase2(param1 *QOperatingSystemVersionBase) *QOperatingSystemVersionBase {
	g := newQOperatingSystemVersionBase(QOperatingSystemVersionBase_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQOperatingSystemVersionBase3 constructs a new QOperatingSystemVersionBase object.
func NewQOperatingSystemVersionBase3(osType OSType, vmajor int, vminor int) *QOperatingSystemVersionBase {
	g := newQOperatingSystemVersionBase(QOperatingSystemVersionBase_new3(osType, (int)(vmajor), (int)(vminor)))
	g.isSubclass = true
	return g
}

// NewQOperatingSystemVersionBase4 constructs a new QOperatingSystemVersionBase object.
func NewQOperatingSystemVersionBase4(osType OSType, vmajor int, vminor int, vmicro int) *QOperatingSystemVersionBase {
	g := newQOperatingSystemVersionBase(QOperatingSystemVersionBase_new4(osType, (int)(vmajor), (int)(vminor), (int)(vmicro)))
	g.isSubclass = true
	return g
}

func QOperatingSystemVersionBase_Current() *QOperatingSystemVersionBase {
	_goptr := newQOperatingSystemVersionBase(QOperatingSystemVersionBase_Current())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QOperatingSystemVersionBase_Name(osversion QOperatingSystemVersionBase) string {
	var _ms struct_miqt_string = QOperatingSystemVersionBase_Name(osversion.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QOperatingSystemVersionBase_CurrentType() OSType {
	xxxxxxxxx
}

func (this *QOperatingSystemVersionBase) Version() *QVersionNumber {
	_goptr := newQVersionNumber(QOperatingSystemVersionBase_Version(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QOperatingSystemVersionBase) MajorVersion() int {
	return (int)(QOperatingSystemVersionBase_MajorVersion(this.h))
}

func (this *QOperatingSystemVersionBase) MinorVersion() int {
	return (int)(QOperatingSystemVersionBase_MinorVersion(this.h))
}

func (this *QOperatingSystemVersionBase) MicroVersion() int {
	return (int)(QOperatingSystemVersionBase_MicroVersion(this.h))
}

func (this *QOperatingSystemVersionBase) SegmentCount() int {
	return (int)(QOperatingSystemVersionBase_SegmentCount(this.h))
}

func (this *QOperatingSystemVersionBase) Type() OSType {
	xxxxxxxxx
}

func (this *QOperatingSystemVersionBase) Name2() string {
	var _ms struct_miqt_string = QOperatingSystemVersionBase_Name2(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QOperatingSystemVersionUnexported struct {
	h          uintptr
	isSubclass bool
}

// NewQOperatingSystemVersionUnexported constructs a new QOperatingSystemVersionUnexported object.
func NewQOperatingSystemVersionUnexported(other QOperatingSystemVersionBase) *QOperatingSystemVersionUnexported {
	g := newQOperatingSystemVersionUnexported(QOperatingSystemVersionUnexported_new(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQOperatingSystemVersionUnexported2 constructs a new QOperatingSystemVersionUnexported object.
func NewQOperatingSystemVersionUnexported2() *QOperatingSystemVersionUnexported {
	g := newQOperatingSystemVersionUnexported(QOperatingSystemVersionUnexported_new2())
	g.isSubclass = true
	return g
}

// NewQOperatingSystemVersionUnexported3 constructs a new QOperatingSystemVersionUnexported object.
func NewQOperatingSystemVersionUnexported3(param1 *QOperatingSystemVersionUnexported) *QOperatingSystemVersionUnexported {
	g := newQOperatingSystemVersionUnexported(QOperatingSystemVersionUnexported_new3(param1.cPointer()))
	g.isSubclass = true
	return g
}

type QOperatingSystemVersion struct {
	h          uintptr
	isSubclass bool
}

// NewQOperatingSystemVersion constructs a new QOperatingSystemVersion object.
func NewQOperatingSystemVersion(osversion *QOperatingSystemVersionBase) *QOperatingSystemVersion {
	g := newQOperatingSystemVersion(QOperatingSystemVersion_new(osversion.cPointer()))
	g.isSubclass = true
	return g
}

// NewQOperatingSystemVersion2 constructs a new QOperatingSystemVersion object.
func NewQOperatingSystemVersion2(osType OSType, vmajor int) *QOperatingSystemVersion {
	g := newQOperatingSystemVersion(QOperatingSystemVersion_new2(osType, (int)(vmajor)))
	g.isSubclass = true
	return g
}

// NewQOperatingSystemVersion3 constructs a new QOperatingSystemVersion object.
func NewQOperatingSystemVersion3(param1 *QOperatingSystemVersion) *QOperatingSystemVersion {
	g := newQOperatingSystemVersion(QOperatingSystemVersion_new3(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQOperatingSystemVersion4 constructs a new QOperatingSystemVersion object.
func NewQOperatingSystemVersion4(osType OSType, vmajor int, vminor int) *QOperatingSystemVersion {
	g := newQOperatingSystemVersion(QOperatingSystemVersion_new4(osType, (int)(vmajor), (int)(vminor)))
	g.isSubclass = true
	return g
}

// NewQOperatingSystemVersion5 constructs a new QOperatingSystemVersion object.
func NewQOperatingSystemVersion5(osType OSType, vmajor int, vminor int, vmicro int) *QOperatingSystemVersion {
	g := newQOperatingSystemVersion(QOperatingSystemVersion_new5(osType, (int)(vmajor), (int)(vminor), (int)(vmicro)))
	g.isSubclass = true
	return g
}

func QOperatingSystemVersion_CurrentType() OSType {
	xxxxxxxxx
}

func (this *QOperatingSystemVersion) Type() OSType {
	xxxxxxxxx
}
