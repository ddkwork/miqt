package qt

import (
	"unsafe"
)

type QSettings__Status int

const (
	QSettings__NoError     QSettings__Status = 0
	QSettings__AccessError QSettings__Status = 1
	QSettings__FormatError QSettings__Status = 2
)

type QSettings__Format int

const (
	QSettings__NativeFormat     QSettings__Format = 0
	QSettings__IniFormat        QSettings__Format = 1
	QSettings__Registry32Format QSettings__Format = 2
	QSettings__Registry64Format QSettings__Format = 3
	QSettings__InvalidFormat    QSettings__Format = 16
	QSettings__CustomFormat1    QSettings__Format = 17
	QSettings__CustomFormat2    QSettings__Format = 18
	QSettings__CustomFormat3    QSettings__Format = 19
	QSettings__CustomFormat4    QSettings__Format = 20
	QSettings__CustomFormat5    QSettings__Format = 21
	QSettings__CustomFormat6    QSettings__Format = 22
	QSettings__CustomFormat7    QSettings__Format = 23
	QSettings__CustomFormat8    QSettings__Format = 24
	QSettings__CustomFormat9    QSettings__Format = 25
	QSettings__CustomFormat10   QSettings__Format = 26
	QSettings__CustomFormat11   QSettings__Format = 27
	QSettings__CustomFormat12   QSettings__Format = 28
	QSettings__CustomFormat13   QSettings__Format = 29
	QSettings__CustomFormat14   QSettings__Format = 30
	QSettings__CustomFormat15   QSettings__Format = 31
	QSettings__CustomFormat16   QSettings__Format = 32
)

type QSettings__Scope int

const (
	QSettings__UserScope   QSettings__Scope = 0
	QSettings__SystemScope QSettings__Scope = 1
)

type QSettings struct {
	h          uintptr
	isSubclass bool
}

// NewQSettings constructs a new QSettings object.
func NewQSettings(organization string) *QSettings {
	organization_ms := struct_miqt_string{}
	organization_ms.data = CString(organization)
	organization_ms.len = size_t(len(organization))
	defer free(unsafe.Pointer(organization_ms.data))
	g := newQSettings(QSettings_new(organization_ms))
	g.isSubclass = true
	return g
}

// NewQSettings2 constructs a new QSettings object.
func NewQSettings2(scope Scope, organization string) *QSettings {
	organization_ms := struct_miqt_string{}
	organization_ms.data = CString(organization)
	organization_ms.len = size_t(len(organization))
	defer free(unsafe.Pointer(organization_ms.data))
	g := newQSettings(QSettings_new2(scope, organization_ms))
	g.isSubclass = true
	return g
}

// NewQSettings3 constructs a new QSettings object.
func NewQSettings3(format Format, scope Scope, organization string) *QSettings {
	organization_ms := struct_miqt_string{}
	organization_ms.data = CString(organization)
	organization_ms.len = size_t(len(organization))
	defer free(unsafe.Pointer(organization_ms.data))
	g := newQSettings(QSettings_new3(format, scope, organization_ms))
	g.isSubclass = true
	return g
}

// NewQSettings4 constructs a new QSettings object.
func NewQSettings4(fileName string, format Format) *QSettings {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	g := newQSettings(QSettings_new4(fileName_ms, format))
	g.isSubclass = true
	return g
}

// NewQSettings5 constructs a new QSettings object.
func NewQSettings5() *QSettings {
	g := newQSettings(QSettings_new5())
	g.isSubclass = true
	return g
}

// NewQSettings6 constructs a new QSettings object.
func NewQSettings6(scope Scope) *QSettings {
	g := newQSettings(QSettings_new6(scope))
	g.isSubclass = true
	return g
}

// NewQSettings7 constructs a new QSettings object.
func NewQSettings7(organization string, application string) *QSettings {
	organization_ms := struct_miqt_string{}
	organization_ms.data = CString(organization)
	organization_ms.len = size_t(len(organization))
	defer free(unsafe.Pointer(organization_ms.data))
	application_ms := struct_miqt_string{}
	application_ms.data = CString(application)
	application_ms.len = size_t(len(application))
	defer free(unsafe.Pointer(application_ms.data))
	g := newQSettings(QSettings_new7(organization_ms, application_ms))
	g.isSubclass = true
	return g
}

// NewQSettings8 constructs a new QSettings object.
func NewQSettings8(organization string, application string, parent *QObject) *QSettings {
	organization_ms := struct_miqt_string{}
	organization_ms.data = CString(organization)
	organization_ms.len = size_t(len(organization))
	defer free(unsafe.Pointer(organization_ms.data))
	application_ms := struct_miqt_string{}
	application_ms.data = CString(application)
	application_ms.len = size_t(len(application))
	defer free(unsafe.Pointer(application_ms.data))
	g := newQSettings(QSettings_new8(organization_ms, application_ms, parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSettings9 constructs a new QSettings object.
func NewQSettings9(scope Scope, organization string, application string) *QSettings {
	organization_ms := struct_miqt_string{}
	organization_ms.data = CString(organization)
	organization_ms.len = size_t(len(organization))
	defer free(unsafe.Pointer(organization_ms.data))
	application_ms := struct_miqt_string{}
	application_ms.data = CString(application)
	application_ms.len = size_t(len(application))
	defer free(unsafe.Pointer(application_ms.data))
	g := newQSettings(QSettings_new9(scope, organization_ms, application_ms))
	g.isSubclass = true
	return g
}

// NewQSettings10 constructs a new QSettings object.
func NewQSettings10(scope Scope, organization string, application string, parent *QObject) *QSettings {
	organization_ms := struct_miqt_string{}
	organization_ms.data = CString(organization)
	organization_ms.len = size_t(len(organization))
	defer free(unsafe.Pointer(organization_ms.data))
	application_ms := struct_miqt_string{}
	application_ms.data = CString(application)
	application_ms.len = size_t(len(application))
	defer free(unsafe.Pointer(application_ms.data))
	g := newQSettings(QSettings_new10(scope, organization_ms, application_ms, parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSettings11 constructs a new QSettings object.
func NewQSettings11(format Format, scope Scope, organization string, application string) *QSettings {
	organization_ms := struct_miqt_string{}
	organization_ms.data = CString(organization)
	organization_ms.len = size_t(len(organization))
	defer free(unsafe.Pointer(organization_ms.data))
	application_ms := struct_miqt_string{}
	application_ms.data = CString(application)
	application_ms.len = size_t(len(application))
	defer free(unsafe.Pointer(application_ms.data))
	g := newQSettings(QSettings_new11(format, scope, organization_ms, application_ms))
	g.isSubclass = true
	return g
}

// NewQSettings12 constructs a new QSettings object.
func NewQSettings12(format Format, scope Scope, organization string, application string, parent *QObject) *QSettings {
	organization_ms := struct_miqt_string{}
	organization_ms.data = CString(organization)
	organization_ms.len = size_t(len(organization))
	defer free(unsafe.Pointer(organization_ms.data))
	application_ms := struct_miqt_string{}
	application_ms.data = CString(application)
	application_ms.len = size_t(len(application))
	defer free(unsafe.Pointer(application_ms.data))
	g := newQSettings(QSettings_new12(format, scope, organization_ms, application_ms, parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSettings13 constructs a new QSettings object.
func NewQSettings13(fileName string, format Format, parent *QObject) *QSettings {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	g := newQSettings(QSettings_new13(fileName_ms, format, parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSettings14 constructs a new QSettings object.
func NewQSettings14(parent *QObject) *QSettings {
	g := newQSettings(QSettings_new14(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSettings15 constructs a new QSettings object.
func NewQSettings15(scope Scope, parent *QObject) *QSettings {
	g := newQSettings(QSettings_new15(scope, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QSettings) MetaObject() *QMetaObject {
	return newQMetaObject(QSettings_MetaObject(this.h))
}

func (this *QSettings) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSettings_Metacast(this.h, param1_Cstring))
}

func QSettings_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSettings_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSettings) Clear() {
	QSettings_Clear(this.h)
}

func (this *QSettings) Sync() {
	QSettings_Sync(this.h)
}

func (this *QSettings) Status() Status {
	xxxxxxxxx
}

func (this *QSettings) IsAtomicSyncRequired() bool {
	return (bool)(QSettings_IsAtomicSyncRequired(this.h))
}

func (this *QSettings) SetAtomicSyncRequired(enable bool) {
	QSettings_SetAtomicSyncRequired(this.h, (bool)(enable))
}

func (this *QSettings) BeginGroup(prefix QAnyStringView) {
	QSettings_BeginGroup(this.h, prefix.cPointer())
}

func (this *QSettings) EndGroup() {
	QSettings_EndGroup(this.h)
}

func (this *QSettings) Group() string {
	var _ms struct_miqt_string = QSettings_Group(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSettings) BeginReadArray(prefix QAnyStringView) int {
	return (int)(QSettings_BeginReadArray(this.h, prefix.cPointer()))
}

func (this *QSettings) BeginWriteArray(prefix QAnyStringView) {
	QSettings_BeginWriteArray(this.h, prefix.cPointer())
}

func (this *QSettings) EndArray() {
	QSettings_EndArray(this.h)
}

func (this *QSettings) SetArrayIndex(i int) {
	QSettings_SetArrayIndex(this.h, (int)(i))
}

func (this *QSettings) AllKeys() []string {
	var _ma struct_miqt_array = QSettings_AllKeys(this.h)
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

func (this *QSettings) ChildKeys() []string {
	var _ma struct_miqt_array = QSettings_ChildKeys(this.h)
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

func (this *QSettings) ChildGroups() []string {
	var _ma struct_miqt_array = QSettings_ChildGroups(this.h)
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

func (this *QSettings) IsWritable() bool {
	return (bool)(QSettings_IsWritable(this.h))
}

func (this *QSettings) SetValue(key QAnyStringView, value *QVariant) {
	QSettings_SetValue(this.h, key.cPointer(), value.cPointer())
}

func (this *QSettings) Value(key QAnyStringView, defaultValue *QVariant) *QVariant {
	_goptr := newQVariant(QSettings_Value(this.h, key.cPointer(), defaultValue.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSettings) ValueWithKey(key QAnyStringView) *QVariant {
	_goptr := newQVariant(QSettings_ValueWithKey(this.h, key.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSettings) Remove(key QAnyStringView) {
	QSettings_Remove(this.h, key.cPointer())
}

func (this *QSettings) Contains(key QAnyStringView) bool {
	return (bool)(QSettings_Contains(this.h, key.cPointer()))
}

func (this *QSettings) SetFallbacksEnabled(b bool) {
	QSettings_SetFallbacksEnabled(this.h, (bool)(b))
}

func (this *QSettings) FallbacksEnabled() bool {
	return (bool)(QSettings_FallbacksEnabled(this.h))
}

func (this *QSettings) FileName() string {
	var _ms struct_miqt_string = QSettings_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSettings) Format() Format {
	xxxxxxxxx
}

func (this *QSettings) Scope() Scope {
	xxxxxxxxx
}

func (this *QSettings) OrganizationName() string {
	var _ms struct_miqt_string = QSettings_OrganizationName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSettings) ApplicationName() string {
	var _ms struct_miqt_string = QSettings_ApplicationName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSettings_SetDefaultFormat(format Format) {
	QSettings_SetDefaultFormat(format)
}

func QSettings_DefaultFormat() Format {
	xxxxxxxxx
}

func QSettings_SetPath(format Format, scope Scope, path string) {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	QSettings_SetPath(format, scope, path_ms)
}

func QSettings_RegisterFormat(extension string, readFunc ReadFunc, writeFunc WriteFunc) Format {
	extension_ms := struct_miqt_string{}
	extension_ms.data = CString(extension)
	extension_ms.len = size_t(len(extension))
	defer free(unsafe.Pointer(extension_ms.data))
	xxxxxxxxx
}

func QSettings_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSettings_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSettings_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSettings_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSettings) BeginWriteArray2(prefix QAnyStringView, size int) {
	QSettings_BeginWriteArray2(this.h, prefix.cPointer(), (int)(size))
}

func QSettings_RegisterFormat4(extension string, readFunc ReadFunc, writeFunc WriteFunc, caseSensitivity CaseSensitivity) Format {
	extension_ms := struct_miqt_string{}
	extension_ms.data = CString(extension)
	extension_ms.len = size_t(len(extension))
	defer free(unsafe.Pointer(extension_ms.data))
	xxxxxxxxx
}

func (this *QSettings) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSettings_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSettings) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSettings_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSettings_MetaObject
func miqt_exec_callback_QSettings_MetaObject(self QSettings, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSettings{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSettings) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSettings_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSettings) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSettings_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSettings_Metacast
func miqt_exec_callback_QSettings_Metacast(self QSettings, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QSettings{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
