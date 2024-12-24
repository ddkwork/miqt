package qt

import (
	"unsafe"
)

type QVariantAnimation struct {
	h          uintptr
	isSubclass bool
}

// NewQVariantAnimation constructs a new QVariantAnimation object.
func NewQVariantAnimation() *QVariantAnimation {
	g := newQVariantAnimation(QVariantAnimation_new())
	g.isSubclass = true
	return g
}

// NewQVariantAnimation2 constructs a new QVariantAnimation object.
func NewQVariantAnimation2(parent *QObject) *QVariantAnimation {
	g := newQVariantAnimation(QVariantAnimation_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QVariantAnimation) MetaObject() *QMetaObject {
	return newQMetaObject(QVariantAnimation_MetaObject(this.h))
}

func (this *QVariantAnimation) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QVariantAnimation_Metacast(this.h, param1_Cstring))
}

func QVariantAnimation_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QVariantAnimation_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVariantAnimation) StartValue() *QVariant {
	_goptr := newQVariant(QVariantAnimation_StartValue(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariantAnimation) SetStartValue(value *QVariant) {
	QVariantAnimation_SetStartValue(this.h, value.cPointer())
}

func (this *QVariantAnimation) EndValue() *QVariant {
	_goptr := newQVariant(QVariantAnimation_EndValue(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariantAnimation) SetEndValue(value *QVariant) {
	QVariantAnimation_SetEndValue(this.h, value.cPointer())
}

func (this *QVariantAnimation) KeyValueAt(step float64) *QVariant {
	_goptr := newQVariant(QVariantAnimation_KeyValueAt(this.h, (double)(step)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariantAnimation) SetKeyValueAt(step float64, value *QVariant) {
	QVariantAnimation_SetKeyValueAt(this.h, (double)(step), value.cPointer())
}

func (this *QVariantAnimation) KeyValues() KeyValues {
	xxxxxxxxx
}

func (this *QVariantAnimation) SetKeyValues(values *KeyValues) {
	QVariantAnimation_SetKeyValues(this.h, values)
}

func (this *QVariantAnimation) CurrentValue() *QVariant {
	_goptr := newQVariant(QVariantAnimation_CurrentValue(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariantAnimation) Duration() int {
	return (int)(QVariantAnimation_Duration(this.h))
}

func (this *QVariantAnimation) SetDuration(msecs int) {
	QVariantAnimation_SetDuration(this.h, (int)(msecs))
}

func (this *QVariantAnimation) EasingCurve() *QEasingCurve {
	_goptr := newQEasingCurve(QVariantAnimation_EasingCurve(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVariantAnimation) SetEasingCurve(easing *QEasingCurve) {
	QVariantAnimation_SetEasingCurve(this.h, easing.cPointer())
}

func (this *QVariantAnimation) ValueChanged(value *QVariant) {
	QVariantAnimation_ValueChanged(this.h, value.cPointer())
}

func (this *QVariantAnimation) OnValueChanged(slot func(value *QVariant)) {
	QVariantAnimation_connect_ValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVariantAnimation_ValueChanged
func miqt_exec_callback_QVariantAnimation_ValueChanged(cb intptr_t, value *QVariant) {
	gofunc, ok := cgo.Handle(cb).Value().(func(value *QVariant))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQVariant(value)

	gofunc(slotval1)
}

func QVariantAnimation_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVariantAnimation_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVariantAnimation_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVariantAnimation_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVariantAnimation) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QVariantAnimation_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QVariantAnimation) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVariantAnimation_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVariantAnimation_MetaObject
func miqt_exec_callback_QVariantAnimation_MetaObject(self QVariantAnimation, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVariantAnimation{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QVariantAnimation) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QVariantAnimation_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QVariantAnimation) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVariantAnimation_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVariantAnimation_Metacast
func miqt_exec_callback_QVariantAnimation_Metacast(self QVariantAnimation, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QVariantAnimation{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
