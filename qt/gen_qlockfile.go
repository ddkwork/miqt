package qt

import (
	"unsafe"
)

type QLockFile__LockError int

const (
	QLockFile__NoError         QLockFile__LockError = 0
	QLockFile__LockFailedError QLockFile__LockError = 1
	QLockFile__PermissionError QLockFile__LockError = 2
	QLockFile__UnknownError    QLockFile__LockError = 3
)

type QLockFile struct {
	h          uintptr
	isSubclass bool
}

// NewQLockFile constructs a new QLockFile object.
func NewQLockFile(fileName string) *QLockFile {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	g := newQLockFile(QLockFile_new(fileName_ms))
	g.isSubclass = true
	return g
}

func (this *QLockFile) FileName() string {
	var _ms struct_miqt_string = QLockFile_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLockFile) Lock() bool {
	return (bool)(QLockFile_Lock(this.h))
}

func (this *QLockFile) TryLock(timeout int) bool {
	return (bool)(QLockFile_TryLock(this.h, (int)(timeout)))
}

func (this *QLockFile) Unlock() {
	QLockFile_Unlock(this.h)
}

func (this *QLockFile) SetStaleLockTime(staleLockTime int) {
	QLockFile_SetStaleLockTime(this.h, (int)(staleLockTime))
}

func (this *QLockFile) StaleLockTime() int {
	return (int)(QLockFile_StaleLockTime(this.h))
}

func (this *QLockFile) TryLock2() bool {
	return (bool)(QLockFile_TryLock2(this.h))
}

func (this *QLockFile) IsLocked() bool {
	return (bool)(QLockFile_IsLocked(this.h))
}

func (this *QLockFile) RemoveStaleLockFile() bool {
	return (bool)(QLockFile_RemoveStaleLockFile(this.h))
}

func (this *QLockFile) Error() LockError {
	xxxxxxxxx
}
