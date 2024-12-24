package network

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QNetworkDiskCache struct {
	h          uintptr
	isSubclass bool
}

// NewQNetworkDiskCache constructs a new QNetworkDiskCache object.
func NewQNetworkDiskCache() *QNetworkDiskCache {

	ret := newQNetworkDiskCache(QNetworkDiskCache_new())
	ret.isSubclass = true
	return ret
}

// NewQNetworkDiskCache2 constructs a new QNetworkDiskCache object.
func NewQNetworkDiskCache2(parent *qt.QObject) *QNetworkDiskCache {

	ret := newQNetworkDiskCache(QNetworkDiskCache_new2((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QNetworkDiskCache) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QNetworkDiskCache_MetaObject(this.h)))
}

func (this *QNetworkDiskCache) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QNetworkDiskCache_Metacast(this.h, param1_Cstring))
}

func QNetworkDiskCache_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QNetworkDiskCache_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkDiskCache) CacheDirectory() string {
	var _ms struct_miqt_string = QNetworkDiskCache_CacheDirectory(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkDiskCache) SetCacheDirectory(cacheDir string) {
	cacheDir_ms := struct_miqt_string{}
	cacheDir_ms.data = CString(cacheDir)
	cacheDir_ms.len = size_t(len(cacheDir))
	defer free(unsafe.Pointer(cacheDir_ms.data))
	QNetworkDiskCache_SetCacheDirectory(this.h, cacheDir_ms)
}

func (this *QNetworkDiskCache) MaximumCacheSize() int64 {
	return (int64)(QNetworkDiskCache_MaximumCacheSize(this.h))
}

func (this *QNetworkDiskCache) SetMaximumCacheSize(size int64) {
	QNetworkDiskCache_SetMaximumCacheSize(this.h, (longlong)(size))
}

func (this *QNetworkDiskCache) CacheSize() int64 {
	return (int64)(QNetworkDiskCache_CacheSize(this.h))
}

func (this *QNetworkDiskCache) MetaData(url *qt.QUrl) *QNetworkCacheMetaData {
	_goptr := newQNetworkCacheMetaData(QNetworkDiskCache_MetaData(this.h, (*QUrl)(url.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkDiskCache) UpdateMetaData(metaData *QNetworkCacheMetaData) {
	QNetworkDiskCache_UpdateMetaData(this.h, metaData.cPointer())
}

func (this *QNetworkDiskCache) Data(url *qt.QUrl) *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QNetworkDiskCache_Data(this.h, (*QUrl)(url.UnsafePointer()))))
}

func (this *QNetworkDiskCache) Remove(url *qt.QUrl) bool {
	return (bool)(QNetworkDiskCache_Remove(this.h, (*QUrl)(url.UnsafePointer())))
}

func (this *QNetworkDiskCache) Prepare(metaData *QNetworkCacheMetaData) *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QNetworkDiskCache_Prepare(this.h, metaData.cPointer())))
}

func (this *QNetworkDiskCache) Insert(device *qt.QIODevice) {
	QNetworkDiskCache_Insert(this.h, (*QIODevice)(device.UnsafePointer()))
}

func (this *QNetworkDiskCache) FileMetaData(fileName string) *QNetworkCacheMetaData {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	_goptr := newQNetworkCacheMetaData(QNetworkDiskCache_FileMetaData(this.h, fileName_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QNetworkDiskCache) Clear() {
	QNetworkDiskCache_Clear(this.h)
}

func QNetworkDiskCache_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QNetworkDiskCache_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QNetworkDiskCache_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QNetworkDiskCache_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QNetworkDiskCache) callVirtualBase_CacheSize() int64 {

	return (int64)(QNetworkDiskCache_virtualbase_CacheSize(unsafe.Pointer(this.h)))

}
func (this *QNetworkDiskCache) OnCacheSize(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkDiskCache_override_virtual_CacheSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkDiskCache_CacheSize
func miqt_exec_callback_QNetworkDiskCache_CacheSize(self QNetworkDiskCache, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QNetworkDiskCache{h: self}).callVirtualBase_CacheSize)

	return (longlong)(virtualReturn)

}

func (this *QNetworkDiskCache) callVirtualBase_MetaData(url *qt.QUrl) *QNetworkCacheMetaData {

	_goptr := newQNetworkCacheMetaData(QNetworkDiskCache_virtualbase_MetaData(unsafe.Pointer(this.h), (*QUrl)(url.UnsafePointer())))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QNetworkDiskCache) OnMetaData(slot func(super func(url *qt.QUrl) *QNetworkCacheMetaData, url *qt.QUrl) *QNetworkCacheMetaData) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkDiskCache_override_virtual_MetaData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkDiskCache_MetaData
func miqt_exec_callback_QNetworkDiskCache_MetaData(self QNetworkDiskCache, cb intptr_t, url *QUrl) *QNetworkCacheMetaData {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(url *qt.QUrl) *QNetworkCacheMetaData, url *qt.QUrl) *QNetworkCacheMetaData)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQUrl(unsafe.Pointer(url))

	virtualReturn := gofunc((&QNetworkDiskCache{h: self}).callVirtualBase_MetaData, slotval1)

	return virtualReturn.cPointer()

}

func (this *QNetworkDiskCache) callVirtualBase_UpdateMetaData(metaData *QNetworkCacheMetaData) {

	QNetworkDiskCache_virtualbase_UpdateMetaData(unsafe.Pointer(this.h), metaData.cPointer())

}
func (this *QNetworkDiskCache) OnUpdateMetaData(slot func(super func(metaData *QNetworkCacheMetaData), metaData *QNetworkCacheMetaData)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkDiskCache_override_virtual_UpdateMetaData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkDiskCache_UpdateMetaData
func miqt_exec_callback_QNetworkDiskCache_UpdateMetaData(self QNetworkDiskCache, cb intptr_t, metaData *QNetworkCacheMetaData) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(metaData *QNetworkCacheMetaData), metaData *QNetworkCacheMetaData))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQNetworkCacheMetaData(metaData)

	gofunc((&QNetworkDiskCache{h: self}).callVirtualBase_UpdateMetaData, slotval1)

}

func (this *QNetworkDiskCache) callVirtualBase_Data(url *qt.QUrl) *qt.QIODevice {

	return qt.UnsafeNewQIODevice(unsafe.Pointer(QNetworkDiskCache_virtualbase_Data(unsafe.Pointer(this.h), (*QUrl)(url.UnsafePointer()))))

}
func (this *QNetworkDiskCache) OnData(slot func(super func(url *qt.QUrl) *qt.QIODevice, url *qt.QUrl) *qt.QIODevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkDiskCache_override_virtual_Data(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkDiskCache_Data
func miqt_exec_callback_QNetworkDiskCache_Data(self QNetworkDiskCache, cb intptr_t, url *QUrl) *QIODevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(url *qt.QUrl) *qt.QIODevice, url *qt.QUrl) *qt.QIODevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQUrl(unsafe.Pointer(url))

	virtualReturn := gofunc((&QNetworkDiskCache{h: self}).callVirtualBase_Data, slotval1)

	return (*QIODevice)(virtualReturn.UnsafePointer())

}

func (this *QNetworkDiskCache) callVirtualBase_Remove(url *qt.QUrl) bool {

	return (bool)(QNetworkDiskCache_virtualbase_Remove(unsafe.Pointer(this.h), (*QUrl)(url.UnsafePointer())))

}
func (this *QNetworkDiskCache) OnRemove(slot func(super func(url *qt.QUrl) bool, url *qt.QUrl) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkDiskCache_override_virtual_Remove(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkDiskCache_Remove
func miqt_exec_callback_QNetworkDiskCache_Remove(self QNetworkDiskCache, cb intptr_t, url *QUrl) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(url *qt.QUrl) bool, url *qt.QUrl) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQUrl(unsafe.Pointer(url))

	virtualReturn := gofunc((&QNetworkDiskCache{h: self}).callVirtualBase_Remove, slotval1)

	return (bool)(virtualReturn)

}

func (this *QNetworkDiskCache) callVirtualBase_Prepare(metaData *QNetworkCacheMetaData) *qt.QIODevice {

	return qt.UnsafeNewQIODevice(unsafe.Pointer(QNetworkDiskCache_virtualbase_Prepare(unsafe.Pointer(this.h), metaData.cPointer())))

}
func (this *QNetworkDiskCache) OnPrepare(slot func(super func(metaData *QNetworkCacheMetaData) *qt.QIODevice, metaData *QNetworkCacheMetaData) *qt.QIODevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkDiskCache_override_virtual_Prepare(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkDiskCache_Prepare
func miqt_exec_callback_QNetworkDiskCache_Prepare(self QNetworkDiskCache, cb intptr_t, metaData *QNetworkCacheMetaData) *QIODevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(metaData *QNetworkCacheMetaData) *qt.QIODevice, metaData *QNetworkCacheMetaData) *qt.QIODevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQNetworkCacheMetaData(metaData)

	virtualReturn := gofunc((&QNetworkDiskCache{h: self}).callVirtualBase_Prepare, slotval1)

	return (*QIODevice)(virtualReturn.UnsafePointer())

}

func (this *QNetworkDiskCache) callVirtualBase_Insert(device *qt.QIODevice) {

	QNetworkDiskCache_virtualbase_Insert(unsafe.Pointer(this.h), (*QIODevice)(device.UnsafePointer()))

}
func (this *QNetworkDiskCache) OnInsert(slot func(super func(device *qt.QIODevice), device *qt.QIODevice)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkDiskCache_override_virtual_Insert(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkDiskCache_Insert
func miqt_exec_callback_QNetworkDiskCache_Insert(self QNetworkDiskCache, cb intptr_t, device *QIODevice) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(device *qt.QIODevice), device *qt.QIODevice))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQIODevice(unsafe.Pointer(device))

	gofunc((&QNetworkDiskCache{h: self}).callVirtualBase_Insert, slotval1)

}

func (this *QNetworkDiskCache) callVirtualBase_Clear() {

	QNetworkDiskCache_virtualbase_Clear(unsafe.Pointer(this.h))

}
func (this *QNetworkDiskCache) OnClear(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkDiskCache_override_virtual_Clear(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkDiskCache_Clear
func miqt_exec_callback_QNetworkDiskCache_Clear(self QNetworkDiskCache, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QNetworkDiskCache{h: self}).callVirtualBase_Clear)

}

func (this *QNetworkDiskCache) callVirtualBase_Expire() int64 {

	return (int64)(QNetworkDiskCache_virtualbase_Expire(unsafe.Pointer(this.h)))

}
func (this *QNetworkDiskCache) OnExpire(slot func(super func() int64) int64) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QNetworkDiskCache_override_virtual_Expire(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QNetworkDiskCache_Expire
func miqt_exec_callback_QNetworkDiskCache_Expire(self QNetworkDiskCache, cb intptr_t) longlong {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int64) int64)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QNetworkDiskCache{h: self}).callVirtualBase_Expire)

	return (longlong)(virtualReturn)

}
