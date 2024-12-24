package qt

import (
	"unsafe"
)

type QMimeType struct {
	h          uintptr
	isSubclass bool
}

// NewQMimeType constructs a new QMimeType object.
func NewQMimeType() *QMimeType {
	ret := newQMimeType(QMimeType_new())
	ret.isSubclass = true
	return ret
}

// NewQMimeType2 constructs a new QMimeType object.
func NewQMimeType2(other *QMimeType) *QMimeType {
	ret := newQMimeType(QMimeType_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QMimeType) OperatorAssign(other *QMimeType) {
	QMimeType_OperatorAssign(this.h, other.cPointer())
}

func (this *QMimeType) Swap(other *QMimeType) {
	QMimeType_Swap(this.h, other.cPointer())
}

func (this *QMimeType) IsValid() bool {
	return (bool)(QMimeType_IsValid(this.h))
}

func (this *QMimeType) IsDefault() bool {
	return (bool)(QMimeType_IsDefault(this.h))
}

func (this *QMimeType) Name() string {
	var _ms struct_miqt_string = QMimeType_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMimeType) Comment() string {
	var _ms struct_miqt_string = QMimeType_Comment(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMimeType) GenericIconName() string {
	var _ms struct_miqt_string = QMimeType_GenericIconName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMimeType) IconName() string {
	var _ms struct_miqt_string = QMimeType_IconName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMimeType) GlobPatterns() []string {
	var _ma struct_miqt_array = QMimeType_GlobPatterns(this.h)
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

func (this *QMimeType) ParentMimeTypes() []string {
	var _ma struct_miqt_array = QMimeType_ParentMimeTypes(this.h)
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

func (this *QMimeType) AllAncestors() []string {
	var _ma struct_miqt_array = QMimeType_AllAncestors(this.h)
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

func (this *QMimeType) Aliases() []string {
	var _ma struct_miqt_array = QMimeType_Aliases(this.h)
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

func (this *QMimeType) Suffixes() []string {
	var _ma struct_miqt_array = QMimeType_Suffixes(this.h)
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

func (this *QMimeType) PreferredSuffix() string {
	var _ms struct_miqt_string = QMimeType_PreferredSuffix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMimeType) Inherits(mimeTypeName string) bool {
	mimeTypeName_ms := struct_miqt_string{}
	mimeTypeName_ms.data = CString(mimeTypeName)
	mimeTypeName_ms.len = size_t(len(mimeTypeName))
	defer free(unsafe.Pointer(mimeTypeName_ms.data))
	return (bool)(QMimeType_Inherits(this.h, mimeTypeName_ms))
}

func (this *QMimeType) FilterString() string {
	var _ms struct_miqt_string = QMimeType_FilterString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
