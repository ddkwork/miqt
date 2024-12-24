package qt

import (
	"unsafe"
)

type QGraphicsAnchor struct {
	h          uintptr
	isSubclass bool
}

func (this *QGraphicsAnchor) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsAnchor_MetaObject(this.h))
}

func (this *QGraphicsAnchor) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsAnchor_Metacast(this.h, param1_Cstring))
}

func QGraphicsAnchor_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsAnchor_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsAnchor) SetSpacing(spacing float64) {
	QGraphicsAnchor_SetSpacing(this.h, (double)(spacing))
}

func (this *QGraphicsAnchor) UnsetSpacing() {
	QGraphicsAnchor_UnsetSpacing(this.h)
}

func (this *QGraphicsAnchor) Spacing() float64 {
	return (float64)(QGraphicsAnchor_Spacing(this.h))
}

func (this *QGraphicsAnchor) SetSizePolicy(policy QSizePolicy__Policy) {
	QGraphicsAnchor_SetSizePolicy(this.h, (int)(policy))
}

func (this *QGraphicsAnchor) SizePolicy() QSizePolicy__Policy {
	return (QSizePolicy__Policy)(QGraphicsAnchor_SizePolicy(this.h))
}

func QGraphicsAnchor_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsAnchor_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsAnchor_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsAnchor_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QGraphicsAnchorLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsAnchorLayout constructs a new QGraphicsAnchorLayout object.
func NewQGraphicsAnchorLayout() *QGraphicsAnchorLayout {
	ret := newQGraphicsAnchorLayout(QGraphicsAnchorLayout_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsAnchorLayout2 constructs a new QGraphicsAnchorLayout object.
func NewQGraphicsAnchorLayout2(parent *QGraphicsLayoutItem) *QGraphicsAnchorLayout {
	ret := newQGraphicsAnchorLayout(QGraphicsAnchorLayout_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsAnchorLayout) AddAnchor(firstItem *QGraphicsLayoutItem, firstEdge AnchorPoint, secondItem *QGraphicsLayoutItem, secondEdge AnchorPoint) *QGraphicsAnchor {
	return newQGraphicsAnchor(QGraphicsAnchorLayout_AddAnchor(this.h, firstItem.cPointer(), (int)(firstEdge), secondItem.cPointer(), (int)(secondEdge)))
}

func (this *QGraphicsAnchorLayout) Anchor(firstItem *QGraphicsLayoutItem, firstEdge AnchorPoint, secondItem *QGraphicsLayoutItem, secondEdge AnchorPoint) *QGraphicsAnchor {
	return newQGraphicsAnchor(QGraphicsAnchorLayout_Anchor(this.h, firstItem.cPointer(), (int)(firstEdge), secondItem.cPointer(), (int)(secondEdge)))
}

func (this *QGraphicsAnchorLayout) AddCornerAnchors(firstItem *QGraphicsLayoutItem, firstCorner Corner, secondItem *QGraphicsLayoutItem, secondCorner Corner) {
	QGraphicsAnchorLayout_AddCornerAnchors(this.h, firstItem.cPointer(), (int)(firstCorner), secondItem.cPointer(), (int)(secondCorner))
}

func (this *QGraphicsAnchorLayout) AddAnchors(firstItem *QGraphicsLayoutItem, secondItem *QGraphicsLayoutItem) {
	QGraphicsAnchorLayout_AddAnchors(this.h, firstItem.cPointer(), secondItem.cPointer())
}

func (this *QGraphicsAnchorLayout) SetHorizontalSpacing(spacing float64) {
	QGraphicsAnchorLayout_SetHorizontalSpacing(this.h, (double)(spacing))
}

func (this *QGraphicsAnchorLayout) SetVerticalSpacing(spacing float64) {
	QGraphicsAnchorLayout_SetVerticalSpacing(this.h, (double)(spacing))
}

func (this *QGraphicsAnchorLayout) SetSpacing(spacing float64) {
	QGraphicsAnchorLayout_SetSpacing(this.h, (double)(spacing))
}

func (this *QGraphicsAnchorLayout) HorizontalSpacing() float64 {
	return (float64)(QGraphicsAnchorLayout_HorizontalSpacing(this.h))
}

func (this *QGraphicsAnchorLayout) VerticalSpacing() float64 {
	return (float64)(QGraphicsAnchorLayout_VerticalSpacing(this.h))
}

func (this *QGraphicsAnchorLayout) RemoveAt(index int) {
	QGraphicsAnchorLayout_RemoveAt(this.h, (int)(index))
}

func (this *QGraphicsAnchorLayout) SetGeometry(rect *QRectF) {
	QGraphicsAnchorLayout_SetGeometry(this.h, rect.cPointer())
}

func (this *QGraphicsAnchorLayout) Count() int {
	return (int)(QGraphicsAnchorLayout_Count(this.h))
}

func (this *QGraphicsAnchorLayout) ItemAt(index int) *QGraphicsLayoutItem {
	return newQGraphicsLayoutItem(QGraphicsAnchorLayout_ItemAt(this.h, (int)(index)))
}

func (this *QGraphicsAnchorLayout) Invalidate() {
	QGraphicsAnchorLayout_Invalidate(this.h)
}

func (this *QGraphicsAnchorLayout) AddAnchors3(firstItem *QGraphicsLayoutItem, secondItem *QGraphicsLayoutItem, orientations Orientation) {
	QGraphicsAnchorLayout_AddAnchors3(this.h, firstItem.cPointer(), secondItem.cPointer(), (int)(orientations))
}
