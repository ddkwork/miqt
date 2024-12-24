package qt

import (
	"unsafe"
)

type QLayoutItem struct {
	h          uintptr
	isSubclass bool
}

// NewQLayoutItem constructs a new QLayoutItem object.
func NewQLayoutItem() *QLayoutItem {
	g := newQLayoutItem(QLayoutItem_new())
	g.isSubclass = true
	return g
}

// NewQLayoutItem2 constructs a new QLayoutItem object.
func NewQLayoutItem2(param1 *QLayoutItem) *QLayoutItem {
	g := newQLayoutItem(QLayoutItem_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQLayoutItem3 constructs a new QLayoutItem object.
func NewQLayoutItem3(alignment AlignmentFlag) *QLayoutItem {
	g := newQLayoutItem(QLayoutItem_new3((int)(alignment)))
	g.isSubclass = true
	return g
}

func (this *QLayoutItem) SizeHint() *QSize {
	_goptr := newQSize(QLayoutItem_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayoutItem) MinimumSize() *QSize {
	_goptr := newQSize(QLayoutItem_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayoutItem) MaximumSize() *QSize {
	_goptr := newQSize(QLayoutItem_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayoutItem) ExpandingDirections() Orientation {
	return (Orientation)(QLayoutItem_ExpandingDirections(this.h))
}

func (this *QLayoutItem) SetGeometry(geometry *QRect) {
	QLayoutItem_SetGeometry(this.h, geometry.cPointer())
}

func (this *QLayoutItem) Geometry() *QRect {
	_goptr := newQRect(QLayoutItem_Geometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLayoutItem) IsEmpty() bool {
	return (bool)(QLayoutItem_IsEmpty(this.h))
}

func (this *QLayoutItem) HasHeightForWidth() bool {
	return (bool)(QLayoutItem_HasHeightForWidth(this.h))
}

func (this *QLayoutItem) HeightForWidth(param1 int) int {
	return (int)(QLayoutItem_HeightForWidth(this.h, (int)(param1)))
}

func (this *QLayoutItem) MinimumHeightForWidth(param1 int) int {
	return (int)(QLayoutItem_MinimumHeightForWidth(this.h, (int)(param1)))
}

func (this *QLayoutItem) Invalidate() {
	QLayoutItem_Invalidate(this.h)
}

func (this *QLayoutItem) Widget() *QWidget {
	return newQWidget(QLayoutItem_Widget(this.h))
}

func (this *QLayoutItem) Layout() *QLayout {
	return newQLayout(QLayoutItem_Layout(this.h))
}

func (this *QLayoutItem) SpacerItem() *QSpacerItem {
	return newQSpacerItem(QLayoutItem_SpacerItem(this.h))
}

func (this *QLayoutItem) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QLayoutItem_Alignment(this.h))
}

func (this *QLayoutItem) SetAlignment(a AlignmentFlag) {
	QLayoutItem_SetAlignment(this.h, (int)(a))
}

func (this *QLayoutItem) ControlTypes() ControlType {
	return (ControlType)(QLayoutItem_ControlTypes(this.h))
}

type QSpacerItem struct {
	h          uintptr
	isSubclass bool
}

// NewQSpacerItem constructs a new QSpacerItem object.
func NewQSpacerItem(w int, h int) *QSpacerItem {
	g := newQSpacerItem(QSpacerItem_new((int)(w), (int)(h)))
	g.isSubclass = true
	return g
}

// NewQSpacerItem2 constructs a new QSpacerItem object.
func NewQSpacerItem2(param1 *QSpacerItem) *QSpacerItem {
	g := newQSpacerItem(QSpacerItem_new2(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSpacerItem3 constructs a new QSpacerItem object.
func NewQSpacerItem3(w int, h int, hData QSizePolicy__Policy) *QSpacerItem {
	g := newQSpacerItem(QSpacerItem_new3((int)(w), (int)(h), (int)(hData)))
	g.isSubclass = true
	return g
}

// NewQSpacerItem4 constructs a new QSpacerItem object.
func NewQSpacerItem4(w int, h int, hData QSizePolicy__Policy, vData QSizePolicy__Policy) *QSpacerItem {
	g := newQSpacerItem(QSpacerItem_new4((int)(w), (int)(h), (int)(hData), (int)(vData)))
	g.isSubclass = true
	return g
}

func (this *QSpacerItem) ChangeSize(w int, h int) {
	QSpacerItem_ChangeSize(this.h, (int)(w), (int)(h))
}

func (this *QSpacerItem) SizeHint() *QSize {
	_goptr := newQSize(QSpacerItem_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpacerItem) MinimumSize() *QSize {
	_goptr := newQSize(QSpacerItem_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpacerItem) MaximumSize() *QSize {
	_goptr := newQSize(QSpacerItem_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpacerItem) ExpandingDirections() Orientation {
	return (Orientation)(QSpacerItem_ExpandingDirections(this.h))
}

func (this *QSpacerItem) IsEmpty() bool {
	return (bool)(QSpacerItem_IsEmpty(this.h))
}

func (this *QSpacerItem) SetGeometry(geometry *QRect) {
	QSpacerItem_SetGeometry(this.h, geometry.cPointer())
}

func (this *QSpacerItem) Geometry() *QRect {
	_goptr := newQRect(QSpacerItem_Geometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpacerItem) SpacerItem() *QSpacerItem {
	return newQSpacerItem(QSpacerItem_SpacerItem(this.h))
}

func (this *QSpacerItem) SizePolicy() *QSizePolicy {
	_goptr := newQSizePolicy(QSpacerItem_SizePolicy(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSpacerItem) ChangeSize3(w int, h int, hData QSizePolicy__Policy) {
	QSpacerItem_ChangeSize3(this.h, (int)(w), (int)(h), (int)(hData))
}

func (this *QSpacerItem) ChangeSize4(w int, h int, hData QSizePolicy__Policy, vData QSizePolicy__Policy) {
	QSpacerItem_ChangeSize4(this.h, (int)(w), (int)(h), (int)(hData), (int)(vData))
}

type QWidgetItem struct {
	h          uintptr
	isSubclass bool
}

// NewQWidgetItem constructs a new QWidgetItem object.
func NewQWidgetItem(w *QWidget) *QWidgetItem {
	g := newQWidgetItem(QWidgetItem_new(w.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QWidgetItem) SizeHint() *QSize {
	_goptr := newQSize(QWidgetItem_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItem) MinimumSize() *QSize {
	_goptr := newQSize(QWidgetItem_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItem) MaximumSize() *QSize {
	_goptr := newQSize(QWidgetItem_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItem) ExpandingDirections() Orientation {
	return (Orientation)(QWidgetItem_ExpandingDirections(this.h))
}

func (this *QWidgetItem) IsEmpty() bool {
	return (bool)(QWidgetItem_IsEmpty(this.h))
}

func (this *QWidgetItem) SetGeometry(geometry *QRect) {
	QWidgetItem_SetGeometry(this.h, geometry.cPointer())
}

func (this *QWidgetItem) Geometry() *QRect {
	_goptr := newQRect(QWidgetItem_Geometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItem) Widget() *QWidget {
	return newQWidget(QWidgetItem_Widget(this.h))
}

func (this *QWidgetItem) HasHeightForWidth() bool {
	return (bool)(QWidgetItem_HasHeightForWidth(this.h))
}

func (this *QWidgetItem) HeightForWidth(param1 int) int {
	return (int)(QWidgetItem_HeightForWidth(this.h, (int)(param1)))
}

func (this *QWidgetItem) MinimumHeightForWidth(param1 int) int {
	return (int)(QWidgetItem_MinimumHeightForWidth(this.h, (int)(param1)))
}

func (this *QWidgetItem) ControlTypes() ControlType {
	return (ControlType)(QWidgetItem_ControlTypes(this.h))
}

type QWidgetItemV2 struct {
	h          uintptr
	isSubclass bool
}

// NewQWidgetItemV2 constructs a new QWidgetItemV2 object.
func NewQWidgetItemV2(widget *QWidget) *QWidgetItemV2 {
	g := newQWidgetItemV2(QWidgetItemV2_new(widget.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QWidgetItemV2) SizeHint() *QSize {
	_goptr := newQSize(QWidgetItemV2_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItemV2) MinimumSize() *QSize {
	_goptr := newQSize(QWidgetItemV2_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItemV2) MaximumSize() *QSize {
	_goptr := newQSize(QWidgetItemV2_MaximumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWidgetItemV2) HeightForWidth(width int) int {
	return (int)(QWidgetItemV2_HeightForWidth(this.h, (int)(width)))
}
