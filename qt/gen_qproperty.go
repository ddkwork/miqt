package qt

import (
	"unsafe"
)

type QtPrivate__BindableWarnings__Reason int

const (
	QtPrivate__BindableWarnings__InvalidInterface     QtPrivate__BindableWarnings__Reason = 0
	QtPrivate__BindableWarnings__NonBindableInterface QtPrivate__BindableWarnings__Reason = 1
	QtPrivate__BindableWarnings__ReadOnlyInterface    QtPrivate__BindableWarnings__Reason = 2
)

type QPropertyBindingError__Type int

const (
	QPropertyBindingError__NoError         QPropertyBindingError__Type = 0
	QPropertyBindingError__BindingLoop     QPropertyBindingError__Type = 1
	QPropertyBindingError__EvaluationError QPropertyBindingError__Type = 2
	QPropertyBindingError__UnknownError    QPropertyBindingError__Type = 3
)

type QPropertyObserverBase__ObserverTag int

const (
	QPropertyObserverBase__ObserverNotifiesBinding       QPropertyObserverBase__ObserverTag = 0
	QPropertyObserverBase__ObserverNotifiesChangeHandler QPropertyObserverBase__ObserverTag = 1
	QPropertyObserverBase__ObserverIsPlaceholder         QPropertyObserverBase__ObserverTag = 2
)

type QScopedPropertyUpdateGroup struct {
	h          uintptr
	isSubclass bool
}

// NewQScopedPropertyUpdateGroup constructs a new QScopedPropertyUpdateGroup object.
func NewQScopedPropertyUpdateGroup() *QScopedPropertyUpdateGroup {
	g := newQScopedPropertyUpdateGroup(QScopedPropertyUpdateGroup_new())
	g.isSubclass = true
	return g
}

type QPropertyBindingSourceLocation struct {
	h          uintptr
	isSubclass bool
}

// NewQPropertyBindingSourceLocation constructs a new QPropertyBindingSourceLocation object.
func NewQPropertyBindingSourceLocation() *QPropertyBindingSourceLocation {
	g := newQPropertyBindingSourceLocation(QPropertyBindingSourceLocation_new())
	g.isSubclass = true
	return g
}

// NewQPropertyBindingSourceLocation2 constructs a new QPropertyBindingSourceLocation object.
func NewQPropertyBindingSourceLocation2(param1 *QPropertyBindingSourceLocation) *QPropertyBindingSourceLocation {
	g := newQPropertyBindingSourceLocation(QPropertyBindingSourceLocation_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

type QPropertyBindingError struct {
	h          uintptr
	isSubclass bool
}

// NewQPropertyBindingError constructs a new QPropertyBindingError object.
func NewQPropertyBindingError() *QPropertyBindingError {
	g := newQPropertyBindingError(QPropertyBindingError_new())
	g.isSubclass = true
	return g
}

// NewQPropertyBindingError2 constructs a new QPropertyBindingError object.
func NewQPropertyBindingError2(typeVal Type) *QPropertyBindingError {
	g := newQPropertyBindingError(QPropertyBindingError_new2(typeVal))
	g.isSubclass = true
	return g
}

// NewQPropertyBindingError3 constructs a new QPropertyBindingError object.
func NewQPropertyBindingError3(other *QPropertyBindingError) *QPropertyBindingError {
	g := newQPropertyBindingError(QPropertyBindingError_new3(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQPropertyBindingError4 constructs a new QPropertyBindingError object.
func NewQPropertyBindingError4(typeVal Type, description string) *QPropertyBindingError {
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	g := newQPropertyBindingError(QPropertyBindingError_new4(typeVal, description_ms))
	g.isSubclass = true
	return g
}

func (this *QPropertyBindingError) OperatorAssign(other *QPropertyBindingError) {
	QPropertyBindingError_OperatorAssign(this.h, other.cPointer())
}

func (this *QPropertyBindingError) HasError() bool {
	return (bool)(QPropertyBindingError_HasError(this.h))
}

func (this *QPropertyBindingError) Type() Type {
	xxxxxxxxx
}

func (this *QPropertyBindingError) Description() string {
	var _ms struct_miqt_string = QPropertyBindingError_Description(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QUntypedPropertyBinding struct {
	h          uintptr
	isSubclass bool
}

// NewQUntypedPropertyBinding constructs a new QUntypedPropertyBinding object.
func NewQUntypedPropertyBinding() *QUntypedPropertyBinding {
	g := newQUntypedPropertyBinding(QUntypedPropertyBinding_new())
	g.isSubclass = true
	return g
}

// NewQUntypedPropertyBinding2 constructs a new QUntypedPropertyBinding object.
func NewQUntypedPropertyBinding2(metaType QMetaType, vtable *BindingFunctionVTable, function unsafe.Pointer, location *QPropertyBindingSourceLocation) *QUntypedPropertyBinding {
	g := newQUntypedPropertyBinding(QUntypedPropertyBinding_new2(metaType.cPointer(), vtable, function, location.cPointer()))
	g.isSubclass = true
	return g
}

// NewQUntypedPropertyBinding3 constructs a new QUntypedPropertyBinding object.
func NewQUntypedPropertyBinding3(other *QUntypedPropertyBinding) *QUntypedPropertyBinding {
	g := newQUntypedPropertyBinding(QUntypedPropertyBinding_new3(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QUntypedPropertyBinding) OperatorAssign(other *QUntypedPropertyBinding) {
	QUntypedPropertyBinding_OperatorAssign(this.h, other.cPointer())
}

func (this *QUntypedPropertyBinding) IsNull() bool {
	return (bool)(QUntypedPropertyBinding_IsNull(this.h))
}

func (this *QUntypedPropertyBinding) Error() *QPropertyBindingError {
	_goptr := newQPropertyBindingError(QUntypedPropertyBinding_Error(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUntypedPropertyBinding) ValueMetaType() *QMetaType {
	_goptr := newQMetaType(QUntypedPropertyBinding_ValueMetaType(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QPropertyObserverBase struct {
	h          uintptr
	isSubclass bool
}

// NewQPropertyObserverBase constructs a new QPropertyObserverBase object.
func NewQPropertyObserverBase() *QPropertyObserverBase {
	g := newQPropertyObserverBase(QPropertyObserverBase_new())
	g.isSubclass = true
	return g
}

// NewQPropertyObserverBase2 constructs a new QPropertyObserverBase object.
func NewQPropertyObserverBase2(param1 *QPropertyObserverBase) *QPropertyObserverBase {
	g := newQPropertyObserverBase(QPropertyObserverBase_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

type QPropertyObserver struct {
	h          uintptr
	isSubclass bool
}

// NewQPropertyObserver constructs a new QPropertyObserver object.
func NewQPropertyObserver() *QPropertyObserver {
	g := newQPropertyObserver(QPropertyObserver_new())
	g.isSubclass = true
	return g
}

type QPropertyNotifier struct {
	h          uintptr
	isSubclass bool
}

// NewQPropertyNotifier constructs a new QPropertyNotifier object.
func NewQPropertyNotifier() *QPropertyNotifier {
	g := newQPropertyNotifier(QPropertyNotifier_new())
	g.isSubclass = true
	return g
}

type QUntypedBindable struct {
	h          uintptr
	isSubclass bool
}

// NewQUntypedBindable constructs a new QUntypedBindable object.
func NewQUntypedBindable() *QUntypedBindable {
	g := newQUntypedBindable(QUntypedBindable_new())
	g.isSubclass = true
	return g
}

// NewQUntypedBindable2 constructs a new QUntypedBindable object.
func NewQUntypedBindable2(param1 *QUntypedBindable) *QUntypedBindable {
	g := newQUntypedBindable(QUntypedBindable_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QUntypedBindable) IsValid() bool {
	return (bool)(QUntypedBindable_IsValid(this.h))
}

func (this *QUntypedBindable) IsBindable() bool {
	return (bool)(QUntypedBindable_IsBindable(this.h))
}

func (this *QUntypedBindable) IsReadOnly() bool {
	return (bool)(QUntypedBindable_IsReadOnly(this.h))
}

func (this *QUntypedBindable) MakeBinding() *QUntypedPropertyBinding {
	_goptr := newQUntypedPropertyBinding(QUntypedBindable_MakeBinding(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUntypedBindable) TakeBinding() *QUntypedPropertyBinding {
	_goptr := newQUntypedPropertyBinding(QUntypedBindable_TakeBinding(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUntypedBindable) Observe(observer *QPropertyObserver) {
	QUntypedBindable_Observe(this.h, observer.cPointer())
}

func (this *QUntypedBindable) Binding() *QUntypedPropertyBinding {
	_goptr := newQUntypedPropertyBinding(QUntypedBindable_Binding(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUntypedBindable) SetBinding(binding *QUntypedPropertyBinding) bool {
	return (bool)(QUntypedBindable_SetBinding(this.h, binding.cPointer()))
}

func (this *QUntypedBindable) HasBinding() bool {
	return (bool)(QUntypedBindable_HasBinding(this.h))
}

func (this *QUntypedBindable) MetaType() *QMetaType {
	_goptr := newQMetaType(QUntypedBindable_MetaType(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUntypedBindable) MakeBinding1(location *QPropertyBindingSourceLocation) *QUntypedPropertyBinding {
	_goptr := newQUntypedPropertyBinding(QUntypedBindable_MakeBinding1(this.h, location.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
