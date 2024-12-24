package qt

import (
	"unsafe"
)

type QCommandLineOption__Flag int

const (
	QCommandLineOption__HiddenFromHelp     QCommandLineOption__Flag = 1
	QCommandLineOption__ShortOptionStyle   QCommandLineOption__Flag = 2
	QCommandLineOption__IgnoreOptionsAfter QCommandLineOption__Flag = 4
)

type QCommandLineOption struct {
	h          uintptr
	isSubclass bool
}

// NewQCommandLineOption constructs a new QCommandLineOption object.
func NewQCommandLineOption(name string) *QCommandLineOption {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))

	ret := newQCommandLineOption(QCommandLineOption_new(name_ms))
	ret.isSubclass = true
	return ret
}

// NewQCommandLineOption2 constructs a new QCommandLineOption object.
func NewQCommandLineOption2(names []string) *QCommandLineOption {
	names_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(names))))
	defer free(unsafe.Pointer(names_CArray))
	for i := range names {
		names_i_ms := struct_miqt_string{}
		names_i_ms.data = CString(names[i])
		names_i_ms.len = size_t(len(names[i]))
		defer free(unsafe.Pointer(names_i_ms.data))
		names_CArray[i] = names_i_ms
	}
	names_ma := struct_miqt_array{len: size_t(len(names)), data: unsafe.Pointer(names_CArray)}

	ret := newQCommandLineOption(QCommandLineOption_new2(names_ma))
	ret.isSubclass = true
	return ret
}

// NewQCommandLineOption3 constructs a new QCommandLineOption object.
func NewQCommandLineOption3(name string, description string) *QCommandLineOption {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))

	ret := newQCommandLineOption(QCommandLineOption_new3(name_ms, description_ms))
	ret.isSubclass = true
	return ret
}

// NewQCommandLineOption4 constructs a new QCommandLineOption object.
func NewQCommandLineOption4(names []string, description string) *QCommandLineOption {
	names_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(names))))
	defer free(unsafe.Pointer(names_CArray))
	for i := range names {
		names_i_ms := struct_miqt_string{}
		names_i_ms.data = CString(names[i])
		names_i_ms.len = size_t(len(names[i]))
		defer free(unsafe.Pointer(names_i_ms.data))
		names_CArray[i] = names_i_ms
	}
	names_ma := struct_miqt_array{len: size_t(len(names)), data: unsafe.Pointer(names_CArray)}
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))

	ret := newQCommandLineOption(QCommandLineOption_new4(names_ma, description_ms))
	ret.isSubclass = true
	return ret
}

// NewQCommandLineOption5 constructs a new QCommandLineOption object.
func NewQCommandLineOption5(other *QCommandLineOption) *QCommandLineOption {

	ret := newQCommandLineOption(QCommandLineOption_new5(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQCommandLineOption6 constructs a new QCommandLineOption object.
func NewQCommandLineOption6(name string, description string, valueName string) *QCommandLineOption {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	valueName_ms := struct_miqt_string{}
	valueName_ms.data = CString(valueName)
	valueName_ms.len = size_t(len(valueName))
	defer free(unsafe.Pointer(valueName_ms.data))

	ret := newQCommandLineOption(QCommandLineOption_new6(name_ms, description_ms, valueName_ms))
	ret.isSubclass = true
	return ret
}

// NewQCommandLineOption7 constructs a new QCommandLineOption object.
func NewQCommandLineOption7(name string, description string, valueName string, defaultValue string) *QCommandLineOption {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	valueName_ms := struct_miqt_string{}
	valueName_ms.data = CString(valueName)
	valueName_ms.len = size_t(len(valueName))
	defer free(unsafe.Pointer(valueName_ms.data))
	defaultValue_ms := struct_miqt_string{}
	defaultValue_ms.data = CString(defaultValue)
	defaultValue_ms.len = size_t(len(defaultValue))
	defer free(unsafe.Pointer(defaultValue_ms.data))

	ret := newQCommandLineOption(QCommandLineOption_new7(name_ms, description_ms, valueName_ms, defaultValue_ms))
	ret.isSubclass = true
	return ret
}

// NewQCommandLineOption8 constructs a new QCommandLineOption object.
func NewQCommandLineOption8(names []string, description string, valueName string) *QCommandLineOption {
	names_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(names))))
	defer free(unsafe.Pointer(names_CArray))
	for i := range names {
		names_i_ms := struct_miqt_string{}
		names_i_ms.data = CString(names[i])
		names_i_ms.len = size_t(len(names[i]))
		defer free(unsafe.Pointer(names_i_ms.data))
		names_CArray[i] = names_i_ms
	}
	names_ma := struct_miqt_array{len: size_t(len(names)), data: unsafe.Pointer(names_CArray)}
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	valueName_ms := struct_miqt_string{}
	valueName_ms.data = CString(valueName)
	valueName_ms.len = size_t(len(valueName))
	defer free(unsafe.Pointer(valueName_ms.data))

	ret := newQCommandLineOption(QCommandLineOption_new8(names_ma, description_ms, valueName_ms))
	ret.isSubclass = true
	return ret
}

// NewQCommandLineOption9 constructs a new QCommandLineOption object.
func NewQCommandLineOption9(names []string, description string, valueName string, defaultValue string) *QCommandLineOption {
	names_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(names))))
	defer free(unsafe.Pointer(names_CArray))
	for i := range names {
		names_i_ms := struct_miqt_string{}
		names_i_ms.data = CString(names[i])
		names_i_ms.len = size_t(len(names[i]))
		defer free(unsafe.Pointer(names_i_ms.data))
		names_CArray[i] = names_i_ms
	}
	names_ma := struct_miqt_array{len: size_t(len(names)), data: unsafe.Pointer(names_CArray)}
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	valueName_ms := struct_miqt_string{}
	valueName_ms.data = CString(valueName)
	valueName_ms.len = size_t(len(valueName))
	defer free(unsafe.Pointer(valueName_ms.data))
	defaultValue_ms := struct_miqt_string{}
	defaultValue_ms.data = CString(defaultValue)
	defaultValue_ms.len = size_t(len(defaultValue))
	defer free(unsafe.Pointer(defaultValue_ms.data))

	ret := newQCommandLineOption(QCommandLineOption_new9(names_ma, description_ms, valueName_ms, defaultValue_ms))
	ret.isSubclass = true
	return ret
}

func (this *QCommandLineOption) OperatorAssign(other *QCommandLineOption) {
	QCommandLineOption_OperatorAssign(this.h, other.cPointer())
}

func (this *QCommandLineOption) Swap(other *QCommandLineOption) {
	QCommandLineOption_Swap(this.h, other.cPointer())
}

func (this *QCommandLineOption) Names() []string {
	var _ma struct_miqt_array = QCommandLineOption_Names(this.h)
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

func (this *QCommandLineOption) SetValueName(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QCommandLineOption_SetValueName(this.h, name_ms)
}

func (this *QCommandLineOption) ValueName() string {
	var _ms struct_miqt_string = QCommandLineOption_ValueName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineOption) SetDescription(description string) {
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	QCommandLineOption_SetDescription(this.h, description_ms)
}

func (this *QCommandLineOption) Description() string {
	var _ms struct_miqt_string = QCommandLineOption_Description(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCommandLineOption) SetDefaultValue(defaultValue string) {
	defaultValue_ms := struct_miqt_string{}
	defaultValue_ms.data = CString(defaultValue)
	defaultValue_ms.len = size_t(len(defaultValue))
	defer free(unsafe.Pointer(defaultValue_ms.data))
	QCommandLineOption_SetDefaultValue(this.h, defaultValue_ms)
}

func (this *QCommandLineOption) SetDefaultValues(defaultValues []string) {
	defaultValues_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(defaultValues))))
	defer free(unsafe.Pointer(defaultValues_CArray))
	for i := range defaultValues {
		defaultValues_i_ms := struct_miqt_string{}
		defaultValues_i_ms.data = CString(defaultValues[i])
		defaultValues_i_ms.len = size_t(len(defaultValues[i]))
		defer free(unsafe.Pointer(defaultValues_i_ms.data))
		defaultValues_CArray[i] = defaultValues_i_ms
	}
	defaultValues_ma := struct_miqt_array{len: size_t(len(defaultValues)), data: unsafe.Pointer(defaultValues_CArray)}
	QCommandLineOption_SetDefaultValues(this.h, defaultValues_ma)
}

func (this *QCommandLineOption) DefaultValues() []string {
	var _ma struct_miqt_array = QCommandLineOption_DefaultValues(this.h)
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

func (this *QCommandLineOption) Flags() Flags {
	xxxxxxxxx
}

func (this *QCommandLineOption) SetFlags(aflags Flags) {
	QCommandLineOption_SetFlags(this.h, aflags)
}
