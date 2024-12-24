package qt

import (
	"unsafe"
)

type QSizePolicy__PolicyFlag int

const (
	QSizePolicy__GrowFlag   QSizePolicy__PolicyFlag = 1
	QSizePolicy__ExpandFlag QSizePolicy__PolicyFlag = 2
	QSizePolicy__ShrinkFlag QSizePolicy__PolicyFlag = 4
	QSizePolicy__IgnoreFlag QSizePolicy__PolicyFlag = 8
)

type QSizePolicy__Policy int

const (
	QSizePolicy__Fixed            QSizePolicy__Policy = 0
	QSizePolicy__Minimum          QSizePolicy__Policy = 1
	QSizePolicy__Maximum          QSizePolicy__Policy = 4
	QSizePolicy__Preferred        QSizePolicy__Policy = 5
	QSizePolicy__MinimumExpanding QSizePolicy__Policy = 3
	QSizePolicy__Expanding        QSizePolicy__Policy = 7
	QSizePolicy__Ignored          QSizePolicy__Policy = 13
)

type QSizePolicy__ControlType int

const (
	QSizePolicy__DefaultType QSizePolicy__ControlType = 1
	QSizePolicy__ButtonBox   QSizePolicy__ControlType = 2
	QSizePolicy__CheckBox    QSizePolicy__ControlType = 4
	QSizePolicy__ComboBox    QSizePolicy__ControlType = 8
	QSizePolicy__Frame       QSizePolicy__ControlType = 16
	QSizePolicy__GroupBox    QSizePolicy__ControlType = 32
	QSizePolicy__Label       QSizePolicy__ControlType = 64
	QSizePolicy__Line        QSizePolicy__ControlType = 128
	QSizePolicy__LineEdit    QSizePolicy__ControlType = 256
	QSizePolicy__PushButton  QSizePolicy__ControlType = 512
	QSizePolicy__RadioButton QSizePolicy__ControlType = 1024
	QSizePolicy__Slider      QSizePolicy__ControlType = 2048
	QSizePolicy__SpinBox     QSizePolicy__ControlType = 4096
	QSizePolicy__TabWidget   QSizePolicy__ControlType = 8192
	QSizePolicy__ToolButton  QSizePolicy__ControlType = 16384
)

type QSizePolicy struct {
	h          uintptr
	isSubclass bool
}

// NewQSizePolicy constructs a new QSizePolicy object.
func NewQSizePolicy() *QSizePolicy {

	ret := newQSizePolicy(QSizePolicy_new())
	ret.isSubclass = true
	return ret
}

// NewQSizePolicy2 constructs a new QSizePolicy object.
func NewQSizePolicy2(horizontal Policy, vertical Policy) *QSizePolicy {

	ret := newQSizePolicy(QSizePolicy_new2(horizontal, vertical))
	ret.isSubclass = true
	return ret
}

// NewQSizePolicy3 constructs a new QSizePolicy object.
func NewQSizePolicy3(param1 *QSizePolicy) *QSizePolicy {

	ret := newQSizePolicy(QSizePolicy_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSizePolicy4 constructs a new QSizePolicy object.
func NewQSizePolicy4(horizontal Policy, vertical Policy, typeVal ControlType) *QSizePolicy {

	ret := newQSizePolicy(QSizePolicy_new4(horizontal, vertical, typeVal))
	ret.isSubclass = true
	return ret
}

func (this *QSizePolicy) HorizontalPolicy() Policy {
	xxxxxxxxx
}

func (this *QSizePolicy) VerticalPolicy() Policy {
	xxxxxxxxx
}

func (this *QSizePolicy) ControlType() ControlType {
	xxxxxxxxx
}

func (this *QSizePolicy) SetHorizontalPolicy(d Policy) {
	QSizePolicy_SetHorizontalPolicy(this.h, d)
}

func (this *QSizePolicy) SetVerticalPolicy(d Policy) {
	QSizePolicy_SetVerticalPolicy(this.h, d)
}

func (this *QSizePolicy) SetControlType(typeVal ControlType) {
	QSizePolicy_SetControlType(this.h, typeVal)
}

func (this *QSizePolicy) ExpandingDirections() Orientation {
	return (Orientation)(QSizePolicy_ExpandingDirections(this.h))
}

func (this *QSizePolicy) SetHeightForWidth(b bool) {
	QSizePolicy_SetHeightForWidth(this.h, (bool)(b))
}

func (this *QSizePolicy) HasHeightForWidth() bool {
	return (bool)(QSizePolicy_HasHeightForWidth(this.h))
}

func (this *QSizePolicy) SetWidthForHeight(b bool) {
	QSizePolicy_SetWidthForHeight(this.h, (bool)(b))
}

func (this *QSizePolicy) HasWidthForHeight() bool {
	return (bool)(QSizePolicy_HasWidthForHeight(this.h))
}

func (this *QSizePolicy) OperatorEqual(s *QSizePolicy) bool {
	return (bool)(QSizePolicy_OperatorEqual(this.h, s.cPointer()))
}

func (this *QSizePolicy) OperatorNotEqual(s *QSizePolicy) bool {
	return (bool)(QSizePolicy_OperatorNotEqual(this.h, s.cPointer()))
}

func (this *QSizePolicy) HorizontalStretch() int {
	return (int)(QSizePolicy_HorizontalStretch(this.h))
}

func (this *QSizePolicy) VerticalStretch() int {
	return (int)(QSizePolicy_VerticalStretch(this.h))
}

func (this *QSizePolicy) SetHorizontalStretch(stretchFactor int) {
	QSizePolicy_SetHorizontalStretch(this.h, (int)(stretchFactor))
}

func (this *QSizePolicy) SetVerticalStretch(stretchFactor int) {
	QSizePolicy_SetVerticalStretch(this.h, (int)(stretchFactor))
}

func (this *QSizePolicy) RetainSizeWhenHidden() bool {
	return (bool)(QSizePolicy_RetainSizeWhenHidden(this.h))
}

func (this *QSizePolicy) SetRetainSizeWhenHidden(retainSize bool) {
	QSizePolicy_SetRetainSizeWhenHidden(this.h, (bool)(retainSize))
}

func (this *QSizePolicy) Transpose() {
	QSizePolicy_Transpose(this.h)
}

func (this *QSizePolicy) Transposed() *QSizePolicy {
	_goptr := newQSizePolicy(QSizePolicy_Transposed(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
