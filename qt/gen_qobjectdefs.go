package qt

import (
	"unsafe"
)

type QMetaObject__Call int

const (
	QMetaObject__InvokeMetaMethod               QMetaObject__Call = 0
	QMetaObject__ReadProperty                   QMetaObject__Call = 1
	QMetaObject__WriteProperty                  QMetaObject__Call = 2
	QMetaObject__ResetProperty                  QMetaObject__Call = 3
	QMetaObject__CreateInstance                 QMetaObject__Call = 4
	QMetaObject__IndexOfMethod                  QMetaObject__Call = 5
	QMetaObject__RegisterPropertyMetaType       QMetaObject__Call = 6
	QMetaObject__RegisterMethodArgumentMetaType QMetaObject__Call = 7
	QMetaObject__BindableProperty               QMetaObject__Call = 8
	QMetaObject__CustomCall                     QMetaObject__Call = 9
	QMetaObject__ConstructInPlace               QMetaObject__Call = 10
)

type QMethodRawArguments struct {
	h          uintptr
	isSubclass bool
}
type QGenericArgument struct {
	h          uintptr
	isSubclass bool
}

// NewQGenericArgument constructs a new QGenericArgument object.
func NewQGenericArgument() *QGenericArgument {
	g := newQGenericArgument(QGenericArgument_new())
	g.isSubclass = true
	return g
}

// NewQGenericArgument2 constructs a new QGenericArgument object.
func NewQGenericArgument2(param1 *QGenericArgument) *QGenericArgument {
	g := newQGenericArgument(QGenericArgument_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGenericArgument3 constructs a new QGenericArgument object.
func NewQGenericArgument3(aName string) *QGenericArgument {
	aName_Cstring := CString(aName)
	defer free(unsafe.Pointer(aName_Cstring))
	g := newQGenericArgument(QGenericArgument_new3(aName_Cstring))
	g.isSubclass = true
	return g
}

// NewQGenericArgument4 constructs a new QGenericArgument object.
func NewQGenericArgument4(aName string, aData unsafe.Pointer) *QGenericArgument {
	aName_Cstring := CString(aName)
	defer free(unsafe.Pointer(aName_Cstring))
	g := newQGenericArgument(QGenericArgument_new4(aName_Cstring, aData))
	g.isSubclass = true
	return g
}

func (this *QGenericArgument) Data() unsafe.Pointer {
	return (unsafe.Pointer)(QGenericArgument_Data(this.h))
}

func (this *QGenericArgument) Name() string {
	_ret := QGenericArgument_Name(this.h)
	return GoString(_ret)
}

type QGenericReturnArgument struct {
	h          uintptr
	isSubclass bool
}

// NewQGenericReturnArgument constructs a new QGenericReturnArgument object.
func NewQGenericReturnArgument() *QGenericReturnArgument {
	g := newQGenericReturnArgument(QGenericReturnArgument_new())
	g.isSubclass = true
	return g
}

// NewQGenericReturnArgument2 constructs a new QGenericReturnArgument object.
func NewQGenericReturnArgument2(param1 *QGenericReturnArgument) *QGenericReturnArgument {
	g := newQGenericReturnArgument(QGenericReturnArgument_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGenericReturnArgument3 constructs a new QGenericReturnArgument object.
func NewQGenericReturnArgument3(aName string) *QGenericReturnArgument {
	aName_Cstring := CString(aName)
	defer free(unsafe.Pointer(aName_Cstring))
	g := newQGenericReturnArgument(QGenericReturnArgument_new3(aName_Cstring))
	g.isSubclass = true
	return g
}

// NewQGenericReturnArgument4 constructs a new QGenericReturnArgument object.
func NewQGenericReturnArgument4(aName string, aData unsafe.Pointer) *QGenericReturnArgument {
	aName_Cstring := CString(aName)
	defer free(unsafe.Pointer(aName_Cstring))
	g := newQGenericReturnArgument(QGenericReturnArgument_new4(aName_Cstring, aData))
	g.isSubclass = true
	return g
}

type QMetaMethodArgument struct {
	h          uintptr
	isSubclass bool
}
type QMetaMethodReturnArgument struct {
	h          uintptr
	isSubclass bool
}
type QMetaObject struct {
	h          uintptr
	isSubclass bool
}

// NewQMetaObject constructs a new QMetaObject object.
func NewQMetaObject() *QMetaObject {
	g := newQMetaObject(QMetaObject_new())
	g.isSubclass = true
	return g
}

// NewQMetaObject2 constructs a new QMetaObject object.
func NewQMetaObject2(param1 *QMetaObject) *QMetaObject {
	g := newQMetaObject(QMetaObject_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QMetaObject) ClassName() string {
	_ret := QMetaObject_ClassName(this.h)
	return GoString(_ret)
}

func (this *QMetaObject) SuperClass() *QMetaObject {
	return newQMetaObject(QMetaObject_SuperClass(this.h))
}

func (this *QMetaObject) Inherits(metaObject *QMetaObject) bool {
	return (bool)(QMetaObject_Inherits(this.h, metaObject.cPointer()))
}

func (this *QMetaObject) Cast(obj *QObject) *QObject {
	return newQObject(QMetaObject_Cast(this.h, obj.cPointer()))
}

func (this *QMetaObject) CastWithObj(obj *QObject) *QObject {
	return newQObject(QMetaObject_CastWithObj(this.h, obj.cPointer()))
}

func (this *QMetaObject) Tr(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMetaObject_Tr(this.h, s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMetaObject) MetaType() *QMetaType {
	_goptr := newQMetaType(QMetaObject_MetaType(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaObject) MethodOffset() int {
	return (int)(QMetaObject_MethodOffset(this.h))
}

func (this *QMetaObject) EnumeratorOffset() int {
	return (int)(QMetaObject_EnumeratorOffset(this.h))
}

func (this *QMetaObject) PropertyOffset() int {
	return (int)(QMetaObject_PropertyOffset(this.h))
}

func (this *QMetaObject) ClassInfoOffset() int {
	return (int)(QMetaObject_ClassInfoOffset(this.h))
}

func (this *QMetaObject) ConstructorCount() int {
	return (int)(QMetaObject_ConstructorCount(this.h))
}

func (this *QMetaObject) MethodCount() int {
	return (int)(QMetaObject_MethodCount(this.h))
}

func (this *QMetaObject) EnumeratorCount() int {
	return (int)(QMetaObject_EnumeratorCount(this.h))
}

func (this *QMetaObject) PropertyCount() int {
	return (int)(QMetaObject_PropertyCount(this.h))
}

func (this *QMetaObject) ClassInfoCount() int {
	return (int)(QMetaObject_ClassInfoCount(this.h))
}

func (this *QMetaObject) IndexOfConstructor(constructor string) int {
	constructor_Cstring := CString(constructor)
	defer free(unsafe.Pointer(constructor_Cstring))
	return (int)(QMetaObject_IndexOfConstructor(this.h, constructor_Cstring))
}

func (this *QMetaObject) IndexOfMethod(method string) int {
	method_Cstring := CString(method)
	defer free(unsafe.Pointer(method_Cstring))
	return (int)(QMetaObject_IndexOfMethod(this.h, method_Cstring))
}

func (this *QMetaObject) IndexOfSignal(signal string) int {
	signal_Cstring := CString(signal)
	defer free(unsafe.Pointer(signal_Cstring))
	return (int)(QMetaObject_IndexOfSignal(this.h, signal_Cstring))
}

func (this *QMetaObject) IndexOfSlot(slot string) int {
	slot_Cstring := CString(slot)
	defer free(unsafe.Pointer(slot_Cstring))
	return (int)(QMetaObject_IndexOfSlot(this.h, slot_Cstring))
}

func (this *QMetaObject) IndexOfEnumerator(name string) int {
	name_Cstring := CString(name)
	defer free(unsafe.Pointer(name_Cstring))
	return (int)(QMetaObject_IndexOfEnumerator(this.h, name_Cstring))
}

func (this *QMetaObject) IndexOfProperty(name string) int {
	name_Cstring := CString(name)
	defer free(unsafe.Pointer(name_Cstring))
	return (int)(QMetaObject_IndexOfProperty(this.h, name_Cstring))
}

func (this *QMetaObject) IndexOfClassInfo(name string) int {
	name_Cstring := CString(name)
	defer free(unsafe.Pointer(name_Cstring))
	return (int)(QMetaObject_IndexOfClassInfo(this.h, name_Cstring))
}

func (this *QMetaObject) Constructor(index int) *QMetaMethod {
	_goptr := newQMetaMethod(QMetaObject_Constructor(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaObject) Method(index int) *QMetaMethod {
	_goptr := newQMetaMethod(QMetaObject_Method(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaObject) Enumerator(index int) *QMetaEnum {
	_goptr := newQMetaEnum(QMetaObject_Enumerator(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaObject) Property(index int) *QMetaProperty {
	_goptr := newQMetaProperty(QMetaObject_Property(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaObject) ClassInfo(index int) *QMetaClassInfo {
	_goptr := newQMetaClassInfo(QMetaObject_ClassInfo(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMetaObject) UserProperty() *QMetaProperty {
	_goptr := newQMetaProperty(QMetaObject_UserProperty(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QMetaObject_CheckConnectArgs(signal string, method string) bool {
	signal_Cstring := CString(signal)
	defer free(unsafe.Pointer(signal_Cstring))
	method_Cstring := CString(method)
	defer free(unsafe.Pointer(method_Cstring))
	return (bool)(QMetaObject_CheckConnectArgs(signal_Cstring, method_Cstring))
}

func QMetaObject_CheckConnectArgs2(signal *QMetaMethod, method *QMetaMethod) bool {
	return (bool)(QMetaObject_CheckConnectArgs2(signal.cPointer(), method.cPointer()))
}

func QMetaObject_NormalizedSignature(method string) []byte {
	method_Cstring := CString(method)
	defer free(unsafe.Pointer(method_Cstring))
	var _bytearray struct_miqt_string = QMetaObject_NormalizedSignature(method_Cstring)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QMetaObject_NormalizedType(typeVal string) []byte {
	typeVal_Cstring := CString(typeVal)
	defer free(unsafe.Pointer(typeVal_Cstring))
	var _bytearray struct_miqt_string = QMetaObject_NormalizedType(typeVal_Cstring)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QMetaObject_Connect(sender *QObject, signal_index int, receiver *QObject, method_index int) Connection {
	xxxxxxxxx
}

func QMetaObject_Disconnect(sender *QObject, signal_index int, receiver *QObject, method_index int) bool {
	return (bool)(QMetaObject_Disconnect(sender.cPointer(), (int)(signal_index), receiver.cPointer(), (int)(method_index)))
}

func QMetaObject_DisconnectOne(sender *QObject, signal_index int, receiver *QObject, method_index int) bool {
	return (bool)(QMetaObject_DisconnectOne(sender.cPointer(), (int)(signal_index), receiver.cPointer(), (int)(method_index)))
}

func QMetaObject_ConnectSlotsByName(o *QObject) {
	QMetaObject_ConnectSlotsByName(o.cPointer())
}

func QMetaObject_InvokeMethod(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod(obj.cPointer(), member_Cstring, (int)(param3), retVal.cPointer()))
}

func QMetaObject_InvokeMethod2(obj *QObject, member string, retVal QGenericReturnArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod2(obj.cPointer(), member_Cstring, retVal.cPointer()))
}

func QMetaObject_InvokeMethod3(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod3(obj.cPointer(), member_Cstring, (int)(typeVal), val0.cPointer()))
}

func QMetaObject_InvokeMethod4(obj *QObject, member string, val0 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod4(obj.cPointer(), member_Cstring, val0.cPointer()))
}

func (this *QMetaObject) NewInstance(val0 QGenericArgument) *QObject {
	return newQObject(QMetaObject_NewInstance(this.h, val0.cPointer()))
}

func (this *QMetaObject) Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMetaObject_Tr3(this.h, s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMetaObject_Connect5(sender *QObject, signal_index int, receiver *QObject, method_index int, typeVal int) Connection {
	xxxxxxxxx
}

func QMetaObject_Connect6(sender *QObject, signal_index int, receiver *QObject, method_index int, typeVal int, types *int) Connection {
	xxxxxxxxx
}

func QMetaObject_InvokeMethod5(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod5(obj.cPointer(), member_Cstring, (int)(param3), retVal.cPointer(), val0.cPointer()))
}

func QMetaObject_InvokeMethod6(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod6(obj.cPointer(), member_Cstring, (int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer()))
}

func QMetaObject_InvokeMethod7(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod7(obj.cPointer(), member_Cstring, (int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func QMetaObject_InvokeMethod8(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod8(obj.cPointer(), member_Cstring, (int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func QMetaObject_InvokeMethod9(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod9(obj.cPointer(), member_Cstring, (int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func QMetaObject_InvokeMethod10(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod10(obj.cPointer(), member_Cstring, (int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func QMetaObject_InvokeMethod11(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod11(obj.cPointer(), member_Cstring, (int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func QMetaObject_InvokeMethod12(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod12(obj.cPointer(), member_Cstring, (int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func QMetaObject_InvokeMethod13(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod13(obj.cPointer(), member_Cstring, (int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func QMetaObject_InvokeMethod14(obj *QObject, member string, param3 ConnectionType, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod14(obj.cPointer(), member_Cstring, (int)(param3), retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func QMetaObject_InvokeMethod42(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod42(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer()))
}

func QMetaObject_InvokeMethod52(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod52(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer()))
}

func QMetaObject_InvokeMethod62(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod62(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func QMetaObject_InvokeMethod72(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod72(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func QMetaObject_InvokeMethod82(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod82(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func QMetaObject_InvokeMethod92(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod92(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func QMetaObject_InvokeMethod102(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod102(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func QMetaObject_InvokeMethod112(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod112(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func QMetaObject_InvokeMethod122(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod122(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func QMetaObject_InvokeMethod132(obj *QObject, member string, retVal QGenericReturnArgument, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod132(obj.cPointer(), member_Cstring, retVal.cPointer(), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func QMetaObject_InvokeMethod53(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod53(obj.cPointer(), member_Cstring, (int)(typeVal), val0.cPointer(), val1.cPointer()))
}

func QMetaObject_InvokeMethod63(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod63(obj.cPointer(), member_Cstring, (int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func QMetaObject_InvokeMethod73(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod73(obj.cPointer(), member_Cstring, (int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func QMetaObject_InvokeMethod83(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod83(obj.cPointer(), member_Cstring, (int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func QMetaObject_InvokeMethod93(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod93(obj.cPointer(), member_Cstring, (int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func QMetaObject_InvokeMethod103(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod103(obj.cPointer(), member_Cstring, (int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func QMetaObject_InvokeMethod113(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod113(obj.cPointer(), member_Cstring, (int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func QMetaObject_InvokeMethod123(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod123(obj.cPointer(), member_Cstring, (int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func QMetaObject_InvokeMethod133(obj *QObject, member string, typeVal ConnectionType, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod133(obj.cPointer(), member_Cstring, (int)(typeVal), val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func QMetaObject_InvokeMethod43(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod43(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer()))
}

func QMetaObject_InvokeMethod54(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod54(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func QMetaObject_InvokeMethod64(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod64(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func QMetaObject_InvokeMethod74(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod74(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func QMetaObject_InvokeMethod84(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod84(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func QMetaObject_InvokeMethod94(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod94(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func QMetaObject_InvokeMethod104(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod104(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func QMetaObject_InvokeMethod114(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod114(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func QMetaObject_InvokeMethod124(obj *QObject, member string, val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) bool {
	member_Cstring := CString(member)
	defer free(unsafe.Pointer(member_Cstring))
	return (bool)(QMetaObject_InvokeMethod124(obj.cPointer(), member_Cstring, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

func (this *QMetaObject) NewInstance2(val0 QGenericArgument, val1 QGenericArgument) *QObject {
	return newQObject(QMetaObject_NewInstance2(this.h, val0.cPointer(), val1.cPointer()))
}

func (this *QMetaObject) NewInstance3(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument) *QObject {
	return newQObject(QMetaObject_NewInstance3(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer()))
}

func (this *QMetaObject) NewInstance4(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument) *QObject {
	return newQObject(QMetaObject_NewInstance4(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer()))
}

func (this *QMetaObject) NewInstance5(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument) *QObject {
	return newQObject(QMetaObject_NewInstance5(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer()))
}

func (this *QMetaObject) NewInstance6(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument) *QObject {
	return newQObject(QMetaObject_NewInstance6(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer()))
}

func (this *QMetaObject) NewInstance7(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument) *QObject {
	return newQObject(QMetaObject_NewInstance7(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer()))
}

func (this *QMetaObject) NewInstance8(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument) *QObject {
	return newQObject(QMetaObject_NewInstance8(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer()))
}

func (this *QMetaObject) NewInstance9(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument) *QObject {
	return newQObject(QMetaObject_NewInstance9(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer()))
}

func (this *QMetaObject) NewInstance10(val0 QGenericArgument, val1 QGenericArgument, val2 QGenericArgument, val3 QGenericArgument, val4 QGenericArgument, val5 QGenericArgument, val6 QGenericArgument, val7 QGenericArgument, val8 QGenericArgument, val9 QGenericArgument) *QObject {
	return newQObject(QMetaObject_NewInstance10(this.h, val0.cPointer(), val1.cPointer(), val2.cPointer(), val3.cPointer(), val4.cPointer(), val5.cPointer(), val6.cPointer(), val7.cPointer(), val8.cPointer(), val9.cPointer()))
}

type QMetaObject__Connection struct {
	h          uintptr
	isSubclass bool
}

// NewQMetaObject__Connection constructs a new QMetaObject::Connection object.
func NewQMetaObject__Connection() *QMetaObject__Connection {
	g := newQMetaObject__Connection(QMetaObject__Connection_new())
	g.isSubclass = true
	return g
}

// NewQMetaObject__Connection2 constructs a new QMetaObject::Connection object.
func NewQMetaObject__Connection2(other *Connection) *QMetaObject__Connection {
	g := newQMetaObject__Connection(QMetaObject__Connection_new2(other))
	g.isSubclass = true
	return g
}

func (this *QMetaObject__Connection) OperatorAssign(other *Connection) {
	QMetaObject__Connection_OperatorAssign(this.h, other)
}

func (this *QMetaObject__Connection) Swap(other *Connection) {
	QMetaObject__Connection_Swap(this.h, other)
}

type QMetaObject__SuperData struct {
	h          uintptr
	isSubclass bool
}

// NewQMetaObject__SuperData constructs a new QMetaObject::SuperData object.
func NewQMetaObject__SuperData() *QMetaObject__SuperData {
	g := newQMetaObject__SuperData(QMetaObject__SuperData_new())
	g.isSubclass = true
	return g
}

// NewQMetaObject__SuperData2 constructs a new QMetaObject::SuperData object.
func NewQMetaObject__SuperData2(mo *QMetaObject) *QMetaObject__SuperData {
	g := newQMetaObject__SuperData(QMetaObject__SuperData_new2(mo.cPointer()))
	g.isSubclass = true
	return g
}

// NewQMetaObject__SuperData3 constructs a new QMetaObject::SuperData object.
func NewQMetaObject__SuperData3(g Getter) *QMetaObject__SuperData {
	g := newQMetaObject__SuperData(QMetaObject__SuperData_new3(g))
	g.isSubclass = true
	return g
}

// NewQMetaObject__SuperData4 constructs a new QMetaObject::SuperData object.
func NewQMetaObject__SuperData4(param1 *SuperData) *QMetaObject__SuperData {
	g := newQMetaObject__SuperData(QMetaObject__SuperData_new4(param1))
	g.isSubclass = true
	return g
}

func (this *QMetaObject__SuperData) OperatorMinusGreater() *QMetaObject {
	return newQMetaObject(QMetaObject__SuperData_OperatorMinusGreater(this.h))
}

func (this *QMetaObject__SuperData) OperatorAssign(param1 *SuperData) {
	QMetaObject__SuperData_OperatorAssign(this.h, param1)
}

type QMetaObject__Data struct {
	h          uintptr
	isSubclass bool
}

// NewQMetaObject__Data constructs a new QMetaObject::Data object.
func NewQMetaObject__Data() *QMetaObject__Data {
	g := newQMetaObject__Data(QMetaObject__Data_new())
	g.isSubclass = true
	return g
}

// NewQMetaObject__Data2 constructs a new QMetaObject::Data object.
func NewQMetaObject__Data2(param1 *Data) *QMetaObject__Data {
	g := newQMetaObject__Data(QMetaObject__Data_new2(param1))
	g.isSubclass = true
	return g
}

func (this *QMetaObject__Data) OperatorAssign(param1 *Data) {
	QMetaObject__Data_OperatorAssign(this.h, param1)
}
