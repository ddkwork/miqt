package qt

import (
	"unsafe"
)

type QMetaMethod__Access int

const (
	QMetaMethod__Private   QMetaMethod__Access = 0
	QMetaMethod__Protected QMetaMethod__Access = 1
	QMetaMethod__Public    QMetaMethod__Access = 2
)

type QMetaMethod__MethodType int

const (
	QMetaMethod__Method      QMetaMethod__MethodType = 0
	QMetaMethod__Signal      QMetaMethod__MethodType = 1
	QMetaMethod__Slot        QMetaMethod__MethodType = 2
	QMetaMethod__Constructor QMetaMethod__MethodType = 3
)

type QMetaMethod__Attributes int

const (
	QMetaMethod__Compatibility QMetaMethod__Attributes = 1
	QMetaMethod__Cloned        QMetaMethod__Attributes = 2
	QMetaMethod__Scriptable    QMetaMethod__Attributes = 4
)

type QMetaMethod struct {
	h          uintptr
	isSubclass bool
}

// NewQMetaMethod constructs a new QMetaMethod object.
func NewQMetaMethod() *QMetaMethod {
	ret := newQMetaMethod(QMetaMethod_new())
	ret.isSubclass = true
	return ret
}

// NewQMetaMethod2 constructs a new QMetaMethod object.
func NewQMetaMethod2(param1 *QMetaMethod) *QMetaMethod {
	ret := newQMetaMethod(QMetaMethod_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QMetaMethod) MethodSignature() []byte {
	var _bytearray struct_miqt_string = QMetaMethod_MethodSignature(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QMetaMethod) Name() []byte {
	var _bytearray struct_miqt_string = QMetaMethod_Name(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QMetaMethod) NameView() *QByteArrayView {
	_goptr := newQByteArrayView(QMetaMethod_NameView(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaMethod) TypeName() string {
	_ret := QMetaMethod_TypeName(this.h)
	return GoString(_ret)
}

func (this *QMetaMethod) ReturnType() int {
	return (int)(QMetaMethod_ReturnType(this.h))
}

func (this *QMetaMethod) ReturnMetaType() *QMetaType {
	_goptr := newQMetaType(QMetaMethod_ReturnMetaType(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaMethod) ParameterCount() int {
	return (int)(QMetaMethod_ParameterCount(this.h))
}

func (this *QMetaMethod) ParameterType(index int) int {
	return (int)(QMetaMethod_ParameterType(this.h, (int)(index)))
}

func (this *QMetaMethod) ParameterMetaType(index int) *QMetaType {
	_goptr := newQMetaType(QMetaMethod_ParameterMetaType(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaMethod) GetParameterTypes(types *int) {
	QMetaMethod_GetParameterTypes(this.h, (*int)(unsafe.Pointer(types)))
}

func (this *QMetaMethod) ParameterTypes() [][]byte {
	var _ma struct_miqt_array = QMetaMethod_ParameterTypes(this.h)
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QMetaMethod) ParameterTypeName(index int) []byte {
	var _bytearray struct_miqt_string = QMetaMethod_ParameterTypeName(this.h, (int)(index))
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QMetaMethod) ParameterNames() [][]byte {
	var _ma struct_miqt_array = QMetaMethod_ParameterNames(this.h)
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QMetaMethod) Tag() string {
	_ret := QMetaMethod_Tag(this.h)
	return GoString(_ret)
}

func (this *QMetaMethod) Access() Access {
	xxxxxxxxx
}

func (this *QMetaMethod) MethodType() MethodType {
	xxxxxxxxx
}

func (this *QMetaMethod) Attributes() int {
	return (int)(QMetaMethod_Attributes(this.h))
}

func (this *QMetaMethod) MethodIndex() int {
	return (int)(QMetaMethod_MethodIndex(this.h))
}

func (this *QMetaMethod) RelativeMethodIndex() int {
	return (int)(QMetaMethod_RelativeMethodIndex(this.h))
}

func (this *QMetaMethod) Revision() int {
	return (int)(QMetaMethod_Revision(this.h))
}

func (this *QMetaMethod) IsConst() bool {
	return (bool)(QMetaMethod_IsConst(this.h))
}

func (this *QMetaMethod) EnclosingMetaObject() *QMetaObject {
	return newQMetaObject(QMetaMethod_EnclosingMetaObject(this.h))
}

func (this *QMetaMethod) Invoke(object *QObject, connectionType ConnectionType, returnValue QGenericReturnArgument) bool {
	return (bool)(QMetaMethod_Invoke(this.h, object.cPointer(), (int)(connectionType), returnValue.cPointer()))
}

func (this *QMetaMethod) Invoke2(object *QObject, returnValue QGenericReturnArgument) bool {
	return (bool)(QMetaMethod_Invoke2(this.h, object.cPointer(), returnValue.cPointer()))
}

func (this *QMetaMethod) Invoke3(object *QObject, connectionType ConnectionType, val0 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke3(this.h, object.cPointer(), (int)(connectionType), val0.cPointer()))
}

func (this *QMetaMethod) Invoke4(object *QObject, val0 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke4(this.h, object.cPointer(), val0.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget(gadget unsafe.Pointer, returnValue QGenericReturnArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget(this.h, gadget, returnValue.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget2(gadget unsafe.Pointer, val0 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget2(this.h, gadget, val0.cPointer()))
}

func (this *QMetaMethod) IsValid() bool {
	return (bool)(QMetaMethod_IsValid(this.h))
}

func (this *QMetaMethod) Invoke42(object *QObject, connectionType ConnectionType, returnValue QGenericReturnArgument, val0 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke42(this.h, object.cPointer(), (int)(connectionType), returnValue.cPointer(), val0.cPointer()))
}

func (this *QMetaMethod) Invoke5(object *QObject, connectionType ConnectionType, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke5(this.h, object.cPointer(), (int)(connectionType), returnValue.cPointer(), val0.cPointer(), val1.cPointer()))
}

func (this *QMetaMethod) Invoke6(object *QObject, connectionType ConnectionType, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke6(this.h, object.cPointer(), (int)(connectionType), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func (this *QMetaMethod) Invoke7(object *QObject, connectionType ConnectionType, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke7(this.h, object.cPointer(), (int)(connectionType), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func (this *QMetaMethod) Invoke8(object *QObject, connectionType ConnectionType, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke8(this.h, object.cPointer(), (int)(connectionType), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func (this *QMetaMethod) Invoke9(object *QObject, connectionType ConnectionType, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke9(this.h, object.cPointer(), (int)(connectionType), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func (this *QMetaMethod) Invoke10(object *QObject, connectionType ConnectionType, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke10(this.h, object.cPointer(), (int)(connectionType), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func (this *QMetaMethod) Invoke11(object *QObject, connectionType ConnectionType, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke11(this.h, object.cPointer(), (int)(connectionType), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func (this *QMetaMethod) Invoke12(object *QObject, connectionType ConnectionType, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke12(this.h, object.cPointer(), (int)(connectionType), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func (this *QMetaMethod) Invoke13(object *QObject, connectionType ConnectionType, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke13(this.h, object.cPointer(), (int)(connectionType), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func (this *QMetaMethod) Invoke32(object *QObject, returnValue QGenericReturnArgument, val0 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke32(this.h, object.cPointer(), returnValue.cPointer(), val0.cPointer()))
}

func (this *QMetaMethod) Invoke43(object *QObject, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke43(this.h, object.cPointer(), returnValue.cPointer(), val0.cPointer(), val1.cPointer()))
}

func (this *QMetaMethod) Invoke52(object *QObject, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke52(this.h, object.cPointer(), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func (this *QMetaMethod) Invoke62(object *QObject, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke62(this.h, object.cPointer(), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func (this *QMetaMethod) Invoke72(object *QObject, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke72(this.h, object.cPointer(), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func (this *QMetaMethod) Invoke82(object *QObject, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke82(this.h, object.cPointer(), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func (this *QMetaMethod) Invoke92(object *QObject, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke92(this.h, object.cPointer(), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func (this *QMetaMethod) Invoke102(object *QObject, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke102(this.h, object.cPointer(), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func (this *QMetaMethod) Invoke112(object *QObject, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke112(this.h, object.cPointer(), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func (this *QMetaMethod) Invoke122(object *QObject, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke122(this.h, object.cPointer(), returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func (this *QMetaMethod) Invoke44(object *QObject, connectionType ConnectionType, val0 QGenericArgument, val1 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke44(this.h, object.cPointer(), (int)(connectionType), val0.cPointer(), val1.cPointer()))
}

func (this *QMetaMethod) Invoke53(object *QObject, connectionType ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke53(this.h, object.cPointer(), (int)(connectionType), val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func (this *QMetaMethod) Invoke63(object *QObject, connectionType ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke63(this.h, object.cPointer(), (int)(connectionType), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func (this *QMetaMethod) Invoke73(object *QObject, connectionType ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke73(this.h, object.cPointer(), (int)(connectionType), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func (this *QMetaMethod) Invoke83(object *QObject, connectionType ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke83(this.h, object.cPointer(), (int)(connectionType), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func (this *QMetaMethod) Invoke93(object *QObject, connectionType ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke93(this.h, object.cPointer(), (int)(connectionType), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func (this *QMetaMethod) Invoke103(object *QObject, connectionType ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke103(this.h, object.cPointer(), (int)(connectionType), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func (this *QMetaMethod) Invoke113(object *QObject, connectionType ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke113(this.h, object.cPointer(), (int)(connectionType), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func (this *QMetaMethod) Invoke123(object *QObject, connectionType ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke123(this.h, object.cPointer(), (int)(connectionType), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func (this *QMetaMethod) Invoke33(object *QObject, val0 QGenericArgument, val1 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke33(this.h, object.cPointer(), val0.cPointer(), val1.cPointer()))
}

func (this *QMetaMethod) Invoke45(object *QObject, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke45(this.h, object.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func (this *QMetaMethod) Invoke54(object *QObject, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke54(this.h, object.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func (this *QMetaMethod) Invoke64(object *QObject, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke64(this.h, object.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func (this *QMetaMethod) Invoke74(object *QObject, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke74(this.h, object.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func (this *QMetaMethod) Invoke84(object *QObject, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke84(this.h, object.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func (this *QMetaMethod) Invoke94(object *QObject, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke94(this.h, object.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func (this *QMetaMethod) Invoke104(object *QObject, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke104(this.h, object.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func (this *QMetaMethod) Invoke114(object *QObject, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	return (bool)(QMetaMethod_Invoke114(this.h, object.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget3(gadget unsafe.Pointer, returnValue QGenericReturnArgument, val0 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget3(this.h, gadget, returnValue.cPointer(), val0.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget4(gadget unsafe.Pointer, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget4(this.h, gadget, returnValue.cPointer(), val0.cPointer(), val1.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget5(gadget unsafe.Pointer, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget5(this.h, gadget, returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget6(gadget unsafe.Pointer, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget6(this.h, gadget, returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget7(gadget unsafe.Pointer, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget7(this.h, gadget, returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget8(gadget unsafe.Pointer, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget8(this.h, gadget, returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget9(gadget unsafe.Pointer, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget9(this.h, gadget, returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget10(gadget unsafe.Pointer, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget10(this.h, gadget, returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget11(gadget unsafe.Pointer, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget11(this.h, gadget, returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget12(gadget unsafe.Pointer, returnValue QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget12(this.h, gadget, returnValue.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget32(gadget unsafe.Pointer, val0 QGenericArgument, val1 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget32(this.h, gadget, val0.cPointer(), val1.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget42(gadget unsafe.Pointer, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget42(this.h, gadget, val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget52(gadget unsafe.Pointer, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget52(this.h, gadget, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget62(gadget unsafe.Pointer, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget62(this.h, gadget, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget72(gadget unsafe.Pointer, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget72(this.h, gadget, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget82(gadget unsafe.Pointer, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget82(this.h, gadget, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget92(gadget unsafe.Pointer, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget92(this.h, gadget, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget102(gadget unsafe.Pointer, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget102(this.h, gadget, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func (this *QMetaMethod) InvokeOnGadget112(gadget unsafe.Pointer, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	return (bool)(QMetaMethod_InvokeOnGadget112(this.h, gadget, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

type QMetaEnum struct {
	h          uintptr
	isSubclass bool
}

// NewQMetaEnum constructs a new QMetaEnum object.
func NewQMetaEnum() *QMetaEnum {
	ret := newQMetaEnum(QMetaEnum_new())
	ret.isSubclass = true
	return ret
}

// NewQMetaEnum2 constructs a new QMetaEnum object.
func NewQMetaEnum2(param1 *QMetaEnum) *QMetaEnum {
	ret := newQMetaEnum(QMetaEnum_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QMetaEnum) Name() string {
	_ret := QMetaEnum_Name(this.h)
	return GoString(_ret)
}

func (this *QMetaEnum) EnumName() string {
	_ret := QMetaEnum_EnumName(this.h)
	return GoString(_ret)
}

func (this *QMetaEnum) MetaType() *QMetaType {
	_goptr := newQMetaType(QMetaEnum_MetaType(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaEnum) IsFlag() bool {
	return (bool)(QMetaEnum_IsFlag(this.h))
}

func (this *QMetaEnum) IsScoped() bool {
	return (bool)(QMetaEnum_IsScoped(this.h))
}

func (this *QMetaEnum) Is64Bit() bool {
	return (bool)(QMetaEnum_Is64Bit(this.h))
}

func (this *QMetaEnum) KeyCount() int {
	return (int)(QMetaEnum_KeyCount(this.h))
}

func (this *QMetaEnum) Key(index int) string {
	_ret := QMetaEnum_Key(this.h, (int)(index))
	return GoString(_ret)
}

func (this *QMetaEnum) Value(index int) int {
	return (int)(QMetaEnum_Value(this.h, (int)(index)))
}

func (this *QMetaEnum) Scope() string {
	_ret := QMetaEnum_Scope(this.h)
	return GoString(_ret)
}

func (this *QMetaEnum) KeyToValue(key string) int {
	key_Cstring := CString(key)
	defer free(unsafe.Pointer(key_Cstring))
	return (int)(QMetaEnum_KeyToValue(this.h, key_Cstring))
}

func (this *QMetaEnum) KeysToValue(keys string) int {
	keys_Cstring := CString(keys)
	defer free(unsafe.Pointer(keys_Cstring))
	return (int)(QMetaEnum_KeysToValue(this.h, keys_Cstring))
}

func (this *QMetaEnum) ValueToKey(value uint64) string {
	_ret := QMetaEnum_ValueToKey(this.h, (ulonglong)(value))
	return GoString(_ret)
}

func (this *QMetaEnum) ValueToKeys(value uint64) []byte {
	var _bytearray struct_miqt_string = QMetaEnum_ValueToKeys(this.h, (ulonglong)(value))
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QMetaEnum) EnclosingMetaObject() *QMetaObject {
	return newQMetaObject(QMetaEnum_EnclosingMetaObject(this.h))
}

func (this *QMetaEnum) IsValid() bool {
	return (bool)(QMetaEnum_IsValid(this.h))
}

func (this *QMetaEnum) KeyToValue2(key string, ok *bool) int {
	key_Cstring := CString(key)
	defer free(unsafe.Pointer(key_Cstring))
	return (int)(QMetaEnum_KeyToValue2(this.h, key_Cstring, (*bool)(unsafe.Pointer(ok))))
}

func (this *QMetaEnum) KeysToValue2(keys string, ok *bool) int {
	keys_Cstring := CString(keys)
	defer free(unsafe.Pointer(keys_Cstring))
	return (int)(QMetaEnum_KeysToValue2(this.h, keys_Cstring, (*bool)(unsafe.Pointer(ok))))
}

type QMetaProperty struct {
	h          uintptr
	isSubclass bool
}

// NewQMetaProperty constructs a new QMetaProperty object.
func NewQMetaProperty() *QMetaProperty {
	ret := newQMetaProperty(QMetaProperty_new())
	ret.isSubclass = true
	return ret
}

func (this *QMetaProperty) Name() string {
	_ret := QMetaProperty_Name(this.h)
	return GoString(_ret)
}

func (this *QMetaProperty) TypeName() string {
	_ret := QMetaProperty_TypeName(this.h)
	return GoString(_ret)
}

func (this *QMetaProperty) Type() QVariant__Type {
	return (QVariant__Type)(QMetaProperty_Type(this.h))
}

func (this *QMetaProperty) UserType() int {
	return (int)(QMetaProperty_UserType(this.h))
}

func (this *QMetaProperty) TypeId() int {
	return (int)(QMetaProperty_TypeId(this.h))
}

func (this *QMetaProperty) MetaType() *QMetaType {
	_goptr := newQMetaType(QMetaProperty_MetaType(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaProperty) PropertyIndex() int {
	return (int)(QMetaProperty_PropertyIndex(this.h))
}

func (this *QMetaProperty) RelativePropertyIndex() int {
	return (int)(QMetaProperty_RelativePropertyIndex(this.h))
}

func (this *QMetaProperty) IsReadable() bool {
	return (bool)(QMetaProperty_IsReadable(this.h))
}

func (this *QMetaProperty) IsWritable() bool {
	return (bool)(QMetaProperty_IsWritable(this.h))
}

func (this *QMetaProperty) IsResettable() bool {
	return (bool)(QMetaProperty_IsResettable(this.h))
}

func (this *QMetaProperty) IsDesignable() bool {
	return (bool)(QMetaProperty_IsDesignable(this.h))
}

func (this *QMetaProperty) IsScriptable() bool {
	return (bool)(QMetaProperty_IsScriptable(this.h))
}

func (this *QMetaProperty) IsStored() bool {
	return (bool)(QMetaProperty_IsStored(this.h))
}

func (this *QMetaProperty) IsUser() bool {
	return (bool)(QMetaProperty_IsUser(this.h))
}

func (this *QMetaProperty) IsConstant() bool {
	return (bool)(QMetaProperty_IsConstant(this.h))
}

func (this *QMetaProperty) IsFinal() bool {
	return (bool)(QMetaProperty_IsFinal(this.h))
}

func (this *QMetaProperty) IsRequired() bool {
	return (bool)(QMetaProperty_IsRequired(this.h))
}

func (this *QMetaProperty) IsBindable() bool {
	return (bool)(QMetaProperty_IsBindable(this.h))
}

func (this *QMetaProperty) IsFlagType() bool {
	return (bool)(QMetaProperty_IsFlagType(this.h))
}

func (this *QMetaProperty) IsEnumType() bool {
	return (bool)(QMetaProperty_IsEnumType(this.h))
}

func (this *QMetaProperty) Enumerator() *QMetaEnum {
	_goptr := newQMetaEnum(QMetaProperty_Enumerator(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaProperty) HasNotifySignal() bool {
	return (bool)(QMetaProperty_HasNotifySignal(this.h))
}

func (this *QMetaProperty) NotifySignal() *QMetaMethod {
	_goptr := newQMetaMethod(QMetaProperty_NotifySignal(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaProperty) NotifySignalIndex() int {
	return (int)(QMetaProperty_NotifySignalIndex(this.h))
}

func (this *QMetaProperty) Revision() int {
	return (int)(QMetaProperty_Revision(this.h))
}

func (this *QMetaProperty) Read(obj *QObject) *QVariant {
	_goptr := newQVariant(QMetaProperty_Read(this.h, obj.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaProperty) Write(obj *QObject, value *QVariant) bool {
	return (bool)(QMetaProperty_Write(this.h, obj.cPointer(), value.cPointer()))
}

func (this *QMetaProperty) Reset(obj *QObject) bool {
	return (bool)(QMetaProperty_Reset(this.h, obj.cPointer()))
}

func (this *QMetaProperty) Bindable(object *QObject) *QUntypedBindable {
	_goptr := newQUntypedBindable(QMetaProperty_Bindable(this.h, object.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaProperty) ReadOnGadget(gadget unsafe.Pointer) *QVariant {
	_goptr := newQVariant(QMetaProperty_ReadOnGadget(this.h, gadget))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaProperty) WriteOnGadget(gadget unsafe.Pointer, value *QVariant) bool {
	return (bool)(QMetaProperty_WriteOnGadget(this.h, gadget, value.cPointer()))
}

func (this *QMetaProperty) ResetOnGadget(gadget unsafe.Pointer) bool {
	return (bool)(QMetaProperty_ResetOnGadget(this.h, gadget))
}

func (this *QMetaProperty) HasStdCppSet() bool {
	return (bool)(QMetaProperty_HasStdCppSet(this.h))
}

func (this *QMetaProperty) IsAlias() bool {
	return (bool)(QMetaProperty_IsAlias(this.h))
}

func (this *QMetaProperty) IsValid() bool {
	return (bool)(QMetaProperty_IsValid(this.h))
}

func (this *QMetaProperty) EnclosingMetaObject() *QMetaObject {
	return newQMetaObject(QMetaProperty_EnclosingMetaObject(this.h))
}

type QMetaClassInfo struct {
	h          uintptr
	isSubclass bool
}

// NewQMetaClassInfo constructs a new QMetaClassInfo object.
func NewQMetaClassInfo() *QMetaClassInfo {
	ret := newQMetaClassInfo(QMetaClassInfo_new())
	ret.isSubclass = true
	return ret
}

// NewQMetaClassInfo2 constructs a new QMetaClassInfo object.
func NewQMetaClassInfo2(param1 *QMetaClassInfo) *QMetaClassInfo {
	ret := newQMetaClassInfo(QMetaClassInfo_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QMetaClassInfo) Name() string {
	_ret := QMetaClassInfo_Name(this.h)
	return GoString(_ret)
}

func (this *QMetaClassInfo) Value() string {
	_ret := QMetaClassInfo_Value(this.h)
	return GoString(_ret)
}

func (this *QMetaClassInfo) EnclosingMetaObject() *QMetaObject {
	return newQMetaObject(QMetaClassInfo_EnclosingMetaObject(this.h))
}
