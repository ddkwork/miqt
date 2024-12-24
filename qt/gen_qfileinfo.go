package qt

import (
	"unsafe"
)

type QFileInfo struct {
	h          uintptr
	isSubclass bool
}

// NewQFileInfo constructs a new QFileInfo object.
func NewQFileInfo() *QFileInfo {

	ret := newQFileInfo(QFileInfo_new())
	ret.isSubclass = true
	return ret
}

// NewQFileInfo2 constructs a new QFileInfo object.
func NewQFileInfo2(file string) *QFileInfo {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))

	ret := newQFileInfo(QFileInfo_new2(file_ms))
	ret.isSubclass = true
	return ret
}

// NewQFileInfo3 constructs a new QFileInfo object.
func NewQFileInfo3(file *QFileDevice) *QFileInfo {

	ret := newQFileInfo(QFileInfo_new3(file.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFileInfo4 constructs a new QFileInfo object.
func NewQFileInfo4(dir *QDir, file string) *QFileInfo {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))

	ret := newQFileInfo(QFileInfo_new4(dir.cPointer(), file_ms))
	ret.isSubclass = true
	return ret
}

// NewQFileInfo5 constructs a new QFileInfo object.
func NewQFileInfo5(fileinfo *QFileInfo) *QFileInfo {

	ret := newQFileInfo(QFileInfo_new5(fileinfo.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QFileInfo) OperatorAssign(fileinfo *QFileInfo) {
	QFileInfo_OperatorAssign(this.h, fileinfo.cPointer())
}

func (this *QFileInfo) Swap(other *QFileInfo) {
	QFileInfo_Swap(this.h, other.cPointer())
}

func (this *QFileInfo) SetFile(file string) {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	QFileInfo_SetFile(this.h, file_ms)
}

func (this *QFileInfo) SetFileWithFile(file *QFileDevice) {
	QFileInfo_SetFileWithFile(this.h, file.cPointer())
}

func (this *QFileInfo) SetFile2(dir *QDir, file string) {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	QFileInfo_SetFile2(this.h, dir.cPointer(), file_ms)
}

func (this *QFileInfo) Exists() bool {
	return (bool)(QFileInfo_Exists(this.h))
}

func QFileInfo_ExistsWithFile(file string) bool {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	return (bool)(QFileInfo_ExistsWithFile(file_ms))
}

func (this *QFileInfo) Refresh() {
	QFileInfo_Refresh(this.h)
}

func (this *QFileInfo) FilePath() string {
	var _ms struct_miqt_string = QFileInfo_FilePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) AbsoluteFilePath() string {
	var _ms struct_miqt_string = QFileInfo_AbsoluteFilePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) CanonicalFilePath() string {
	var _ms struct_miqt_string = QFileInfo_CanonicalFilePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) FileName() string {
	var _ms struct_miqt_string = QFileInfo_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) BaseName() string {
	var _ms struct_miqt_string = QFileInfo_BaseName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) CompleteBaseName() string {
	var _ms struct_miqt_string = QFileInfo_CompleteBaseName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) Suffix() string {
	var _ms struct_miqt_string = QFileInfo_Suffix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) BundleName() string {
	var _ms struct_miqt_string = QFileInfo_BundleName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) CompleteSuffix() string {
	var _ms struct_miqt_string = QFileInfo_CompleteSuffix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) Path() string {
	var _ms struct_miqt_string = QFileInfo_Path(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) AbsolutePath() string {
	var _ms struct_miqt_string = QFileInfo_AbsolutePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) CanonicalPath() string {
	var _ms struct_miqt_string = QFileInfo_CanonicalPath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) Dir() *QDir {
	_goptr := newQDir(QFileInfo_Dir(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileInfo) AbsoluteDir() *QDir {
	_goptr := newQDir(QFileInfo_AbsoluteDir(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileInfo) IsReadable() bool {
	return (bool)(QFileInfo_IsReadable(this.h))
}

func (this *QFileInfo) IsWritable() bool {
	return (bool)(QFileInfo_IsWritable(this.h))
}

func (this *QFileInfo) IsExecutable() bool {
	return (bool)(QFileInfo_IsExecutable(this.h))
}

func (this *QFileInfo) IsHidden() bool {
	return (bool)(QFileInfo_IsHidden(this.h))
}

func (this *QFileInfo) IsNativePath() bool {
	return (bool)(QFileInfo_IsNativePath(this.h))
}

func (this *QFileInfo) IsRelative() bool {
	return (bool)(QFileInfo_IsRelative(this.h))
}

func (this *QFileInfo) IsAbsolute() bool {
	return (bool)(QFileInfo_IsAbsolute(this.h))
}

func (this *QFileInfo) MakeAbsolute() bool {
	return (bool)(QFileInfo_MakeAbsolute(this.h))
}

func (this *QFileInfo) IsFile() bool {
	return (bool)(QFileInfo_IsFile(this.h))
}

func (this *QFileInfo) IsDir() bool {
	return (bool)(QFileInfo_IsDir(this.h))
}

func (this *QFileInfo) IsSymLink() bool {
	return (bool)(QFileInfo_IsSymLink(this.h))
}

func (this *QFileInfo) IsSymbolicLink() bool {
	return (bool)(QFileInfo_IsSymbolicLink(this.h))
}

func (this *QFileInfo) IsShortcut() bool {
	return (bool)(QFileInfo_IsShortcut(this.h))
}

func (this *QFileInfo) IsAlias() bool {
	return (bool)(QFileInfo_IsAlias(this.h))
}

func (this *QFileInfo) IsJunction() bool {
	return (bool)(QFileInfo_IsJunction(this.h))
}

func (this *QFileInfo) IsRoot() bool {
	return (bool)(QFileInfo_IsRoot(this.h))
}

func (this *QFileInfo) IsBundle() bool {
	return (bool)(QFileInfo_IsBundle(this.h))
}

func (this *QFileInfo) SymLinkTarget() string {
	var _ms struct_miqt_string = QFileInfo_SymLinkTarget(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) ReadSymLink() string {
	var _ms struct_miqt_string = QFileInfo_ReadSymLink(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) JunctionTarget() string {
	var _ms struct_miqt_string = QFileInfo_JunctionTarget(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) Owner() string {
	var _ms struct_miqt_string = QFileInfo_Owner(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) OwnerId() uint {
	return (uint)(QFileInfo_OwnerId(this.h))
}

func (this *QFileInfo) Group() string {
	var _ms struct_miqt_string = QFileInfo_Group(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileInfo) GroupId() uint {
	return (uint)(QFileInfo_GroupId(this.h))
}

func (this *QFileInfo) Permission(permissions Permission) bool {
	return (bool)(QFileInfo_Permission(this.h, (int)(permissions)))
}

func (this *QFileInfo) Permissions() Permission {
	return (Permission)(QFileInfo_Permissions(this.h))
}

func (this *QFileInfo) Size() int64 {
	return (int64)(QFileInfo_Size(this.h))
}

func (this *QFileInfo) BirthTime() *QDateTime {
	_goptr := newQDateTime(QFileInfo_BirthTime(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileInfo) MetadataChangeTime() *QDateTime {
	_goptr := newQDateTime(QFileInfo_MetadataChangeTime(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileInfo) LastModified() *QDateTime {
	_goptr := newQDateTime(QFileInfo_LastModified(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileInfo) LastRead() *QDateTime {
	_goptr := newQDateTime(QFileInfo_LastRead(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileInfo) FileTime(time QFileDevice__FileTime) *QDateTime {
	_goptr := newQDateTime(QFileInfo_FileTime(this.h, (int)(time)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileInfo) BirthTimeWithTz(tz *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QFileInfo_BirthTimeWithTz(this.h, tz.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileInfo) MetadataChangeTimeWithTz(tz *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QFileInfo_MetadataChangeTimeWithTz(this.h, tz.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileInfo) LastModifiedWithTz(tz *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QFileInfo_LastModifiedWithTz(this.h, tz.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileInfo) LastReadWithTz(tz *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QFileInfo_LastReadWithTz(this.h, tz.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileInfo) FileTime2(time QFileDevice__FileTime, tz *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QFileInfo_FileTime2(this.h, (int)(time), tz.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileInfo) Caching() bool {
	return (bool)(QFileInfo_Caching(this.h))
}

func (this *QFileInfo) SetCaching(on bool) {
	QFileInfo_SetCaching(this.h, (bool)(on))
}

func (this *QFileInfo) Stat() {
	QFileInfo_Stat(this.h)
}
