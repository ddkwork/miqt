package qt

import (
	"unsafe"
)

type QFontVariableAxis struct {
	h          uintptr
	isSubclass bool
}

// NewQFontVariableAxis constructs a new QFontVariableAxis object.
func NewQFontVariableAxis() *QFontVariableAxis {

	ret := newQFontVariableAxis(QFontVariableAxis_new())
	ret.isSubclass = true
	return ret
}

// NewQFontVariableAxis2 constructs a new QFontVariableAxis object.
func NewQFontVariableAxis2(axis *QFontVariableAxis) *QFontVariableAxis {

	ret := newQFontVariableAxis(QFontVariableAxis_new2(axis.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QFontVariableAxis) Swap(other *QFontVariableAxis) {
	QFontVariableAxis_Swap(this.h, other.cPointer())
}

func (this *QFontVariableAxis) OperatorAssign(axis *QFontVariableAxis) {
	QFontVariableAxis_OperatorAssign(this.h, axis.cPointer())
}

func (this *QFontVariableAxis) Tag() *QFont__Tag {
	_goptr := newQFont__Tag(QFontVariableAxis_Tag(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontVariableAxis) SetTag(tag QFont__Tag) {
	QFontVariableAxis_SetTag(this.h, tag.cPointer())
}

func (this *QFontVariableAxis) Name() string {
	var _ms struct_miqt_string = QFontVariableAxis_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontVariableAxis) SetName(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QFontVariableAxis_SetName(this.h, name_ms)
}

func (this *QFontVariableAxis) MinimumValue() float64 {
	return (float64)(QFontVariableAxis_MinimumValue(this.h))
}

func (this *QFontVariableAxis) SetMinimumValue(minimumValue float64) {
	QFontVariableAxis_SetMinimumValue(this.h, (double)(minimumValue))
}

func (this *QFontVariableAxis) MaximumValue() float64 {
	return (float64)(QFontVariableAxis_MaximumValue(this.h))
}

func (this *QFontVariableAxis) SetMaximumValue(maximumValue float64) {
	QFontVariableAxis_SetMaximumValue(this.h, (double)(maximumValue))
}

func (this *QFontVariableAxis) DefaultValue() float64 {
	return (float64)(QFontVariableAxis_DefaultValue(this.h))
}

func (this *QFontVariableAxis) SetDefaultValue(defaultValue float64) {
	QFontVariableAxis_SetDefaultValue(this.h, (double)(defaultValue))
}
