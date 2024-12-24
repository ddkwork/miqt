package qt

import (
	"unsafe"
)

type QStorageInfo struct {
	h          uintptr
	isSubclass bool
}

// NewQStorageInfo constructs a new QStorageInfo object.
func NewQStorageInfo() *QStorageInfo {

	ret := newQStorageInfo(QStorageInfo_new())
	ret.isSubclass = true
	return ret
}

// NewQStorageInfo2 constructs a new QStorageInfo object.
func NewQStorageInfo2(path string) *QStorageInfo {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))

	ret := newQStorageInfo(QStorageInfo_new2(path_ms))
	ret.isSubclass = true
	return ret
}

// NewQStorageInfo3 constructs a new QStorageInfo object.
func NewQStorageInfo3(dir *QDir) *QStorageInfo {

	ret := newQStorageInfo(QStorageInfo_new3(dir.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQStorageInfo4 constructs a new QStorageInfo object.
func NewQStorageInfo4(other *QStorageInfo) *QStorageInfo {

	ret := newQStorageInfo(QStorageInfo_new4(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QStorageInfo) OperatorAssign(other *QStorageInfo) {
	QStorageInfo_OperatorAssign(this.h, other.cPointer())
}

func (this *QStorageInfo) Swap(other *QStorageInfo) {
	QStorageInfo_Swap(this.h, other.cPointer())
}

func (this *QStorageInfo) SetPath(path string) {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	QStorageInfo_SetPath(this.h, path_ms)
}

func (this *QStorageInfo) RootPath() string {
	var _ms struct_miqt_string = QStorageInfo_RootPath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStorageInfo) Device() []byte {
	var _bytearray struct_miqt_string = QStorageInfo_Device(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QStorageInfo) Subvolume() []byte {
	var _bytearray struct_miqt_string = QStorageInfo_Subvolume(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QStorageInfo) FileSystemType() []byte {
	var _bytearray struct_miqt_string = QStorageInfo_FileSystemType(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QStorageInfo) Name() string {
	var _ms struct_miqt_string = QStorageInfo_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStorageInfo) DisplayName() string {
	var _ms struct_miqt_string = QStorageInfo_DisplayName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStorageInfo) BytesTotal() int64 {
	return (int64)(QStorageInfo_BytesTotal(this.h))
}

func (this *QStorageInfo) BytesFree() int64 {
	return (int64)(QStorageInfo_BytesFree(this.h))
}

func (this *QStorageInfo) BytesAvailable() int64 {
	return (int64)(QStorageInfo_BytesAvailable(this.h))
}

func (this *QStorageInfo) BlockSize() int {
	return (int)(QStorageInfo_BlockSize(this.h))
}

func (this *QStorageInfo) IsRoot() bool {
	return (bool)(QStorageInfo_IsRoot(this.h))
}

func (this *QStorageInfo) IsReadOnly() bool {
	return (bool)(QStorageInfo_IsReadOnly(this.h))
}

func (this *QStorageInfo) IsReady() bool {
	return (bool)(QStorageInfo_IsReady(this.h))
}

func (this *QStorageInfo) IsValid() bool {
	return (bool)(QStorageInfo_IsValid(this.h))
}

func (this *QStorageInfo) Refresh() {
	QStorageInfo_Refresh(this.h)
}

func QStorageInfo_MountedVolumes() []QStorageInfo {
	var _ma struct_miqt_array = QStorageInfo_MountedVolumes()
	_ret := make([]QStorageInfo, int(_ma.len))
	_outCast := (*[0xffff]*QStorageInfo)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQStorageInfo(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func QStorageInfo_Root() *QStorageInfo {
	_goptr := newQStorageInfo(QStorageInfo_Root())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
