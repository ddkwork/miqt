package qt

import (
	"unsafe"
)

type QGraphicsItemAnimation struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsItemAnimation constructs a new QGraphicsItemAnimation object.
func NewQGraphicsItemAnimation() *QGraphicsItemAnimation {
	ret := newQGraphicsItemAnimation(QGraphicsItemAnimation_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsItemAnimation2 constructs a new QGraphicsItemAnimation object.
func NewQGraphicsItemAnimation2(parent *QObject) *QGraphicsItemAnimation {
	ret := newQGraphicsItemAnimation(QGraphicsItemAnimation_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsItemAnimation) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsItemAnimation_MetaObject(this.h))
}

func (this *QGraphicsItemAnimation) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsItemAnimation_Metacast(this.h, param1_Cstring))
}

func QGraphicsItemAnimation_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsItemAnimation_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsItemAnimation) Item() *QGraphicsItem {
	return newQGraphicsItem(QGraphicsItemAnimation_Item(this.h))
}

func (this *QGraphicsItemAnimation) SetItem(item *QGraphicsItem) {
	QGraphicsItemAnimation_SetItem(this.h, item.cPointer())
}

func (this *QGraphicsItemAnimation) TimeLine() *QTimeLine {
	return newQTimeLine(QGraphicsItemAnimation_TimeLine(this.h))
}

func (this *QGraphicsItemAnimation) SetTimeLine(timeLine *QTimeLine) {
	QGraphicsItemAnimation_SetTimeLine(this.h, timeLine.cPointer())
}

func (this *QGraphicsItemAnimation) PosAt(step float64) *QPointF {
	_goptr := newQPointF(QGraphicsItemAnimation_PosAt(this.h, (double)(step)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItemAnimation) PosList() []struct {
	First  float64
	Second QPointF
} {
	var _ma struct_miqt_array = QGraphicsItemAnimation_PosList(this.h)
	_ret := make([]struct {
		First  float64
		Second QPointF
	}, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_map)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_mm struct_miqt_map = _outCast[i]
		_lv_First_CArray := (*[0xffff]double)(unsafe.Pointer(_lv_mm.keys))
		_lv_Second_CArray := (*[0xffff]*QPointF)(unsafe.Pointer(_lv_mm.values))
		_lv_entry_First := (float64)(_lv_First_CArray[0])

		_lv_second_goptr := newQPointF(_lv_Second_CArray[0])
		_lv_second_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_lv_entry_Second := *_lv_second_goptr

		_ret[i] = struct {
			First  float64
			Second QPointF
		}{First: _lv_entry_First, Second: _lv_entry_Second}
	}
	return _ret
}

func (this *QGraphicsItemAnimation) SetPosAt(step float64, pos *QPointF) {
	QGraphicsItemAnimation_SetPosAt(this.h, (double)(step), pos.cPointer())
}

func (this *QGraphicsItemAnimation) TransformAt(step float64) *QTransform {
	_goptr := newQTransform(QGraphicsItemAnimation_TransformAt(this.h, (double)(step)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItemAnimation) RotationAt(step float64) float64 {
	return (float64)(QGraphicsItemAnimation_RotationAt(this.h, (double)(step)))
}

func (this *QGraphicsItemAnimation) RotationList() []struct {
	First  float64
	Second float64
} {
	var _ma struct_miqt_array = QGraphicsItemAnimation_RotationList(this.h)
	_ret := make([]struct {
		First  float64
		Second float64
	}, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_map)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_mm struct_miqt_map = _outCast[i]
		_lv_First_CArray := (*[0xffff]double)(unsafe.Pointer(_lv_mm.keys))
		_lv_Second_CArray := (*[0xffff]double)(unsafe.Pointer(_lv_mm.values))
		_lv_entry_First := (float64)(_lv_First_CArray[0])

		_lv_entry_Second := (float64)(_lv_Second_CArray[0])

		_ret[i] = struct {
			First  float64
			Second float64
		}{First: _lv_entry_First, Second: _lv_entry_Second}
	}
	return _ret
}

func (this *QGraphicsItemAnimation) SetRotationAt(step float64, angle float64) {
	QGraphicsItemAnimation_SetRotationAt(this.h, (double)(step), (double)(angle))
}

func (this *QGraphicsItemAnimation) XTranslationAt(step float64) float64 {
	return (float64)(QGraphicsItemAnimation_XTranslationAt(this.h, (double)(step)))
}

func (this *QGraphicsItemAnimation) YTranslationAt(step float64) float64 {
	return (float64)(QGraphicsItemAnimation_YTranslationAt(this.h, (double)(step)))
}

func (this *QGraphicsItemAnimation) TranslationList() []struct {
	First  float64
	Second QPointF
} {
	var _ma struct_miqt_array = QGraphicsItemAnimation_TranslationList(this.h)
	_ret := make([]struct {
		First  float64
		Second QPointF
	}, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_map)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_mm struct_miqt_map = _outCast[i]
		_lv_First_CArray := (*[0xffff]double)(unsafe.Pointer(_lv_mm.keys))
		_lv_Second_CArray := (*[0xffff]*QPointF)(unsafe.Pointer(_lv_mm.values))
		_lv_entry_First := (float64)(_lv_First_CArray[0])

		_lv_second_goptr := newQPointF(_lv_Second_CArray[0])
		_lv_second_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_lv_entry_Second := *_lv_second_goptr

		_ret[i] = struct {
			First  float64
			Second QPointF
		}{First: _lv_entry_First, Second: _lv_entry_Second}
	}
	return _ret
}

func (this *QGraphicsItemAnimation) SetTranslationAt(step float64, dx float64, dy float64) {
	QGraphicsItemAnimation_SetTranslationAt(this.h, (double)(step), (double)(dx), (double)(dy))
}

func (this *QGraphicsItemAnimation) VerticalScaleAt(step float64) float64 {
	return (float64)(QGraphicsItemAnimation_VerticalScaleAt(this.h, (double)(step)))
}

func (this *QGraphicsItemAnimation) HorizontalScaleAt(step float64) float64 {
	return (float64)(QGraphicsItemAnimation_HorizontalScaleAt(this.h, (double)(step)))
}

func (this *QGraphicsItemAnimation) ScaleList() []struct {
	First  float64
	Second QPointF
} {
	var _ma struct_miqt_array = QGraphicsItemAnimation_ScaleList(this.h)
	_ret := make([]struct {
		First  float64
		Second QPointF
	}, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_map)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_mm struct_miqt_map = _outCast[i]
		_lv_First_CArray := (*[0xffff]double)(unsafe.Pointer(_lv_mm.keys))
		_lv_Second_CArray := (*[0xffff]*QPointF)(unsafe.Pointer(_lv_mm.values))
		_lv_entry_First := (float64)(_lv_First_CArray[0])

		_lv_second_goptr := newQPointF(_lv_Second_CArray[0])
		_lv_second_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_lv_entry_Second := *_lv_second_goptr

		_ret[i] = struct {
			First  float64
			Second QPointF
		}{First: _lv_entry_First, Second: _lv_entry_Second}
	}
	return _ret
}

func (this *QGraphicsItemAnimation) SetScaleAt(step float64, sx float64, sy float64) {
	QGraphicsItemAnimation_SetScaleAt(this.h, (double)(step), (double)(sx), (double)(sy))
}

func (this *QGraphicsItemAnimation) VerticalShearAt(step float64) float64 {
	return (float64)(QGraphicsItemAnimation_VerticalShearAt(this.h, (double)(step)))
}

func (this *QGraphicsItemAnimation) HorizontalShearAt(step float64) float64 {
	return (float64)(QGraphicsItemAnimation_HorizontalShearAt(this.h, (double)(step)))
}

func (this *QGraphicsItemAnimation) ShearList() []struct {
	First  float64
	Second QPointF
} {
	var _ma struct_miqt_array = QGraphicsItemAnimation_ShearList(this.h)
	_ret := make([]struct {
		First  float64
		Second QPointF
	}, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_map)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_mm struct_miqt_map = _outCast[i]
		_lv_First_CArray := (*[0xffff]double)(unsafe.Pointer(_lv_mm.keys))
		_lv_Second_CArray := (*[0xffff]*QPointF)(unsafe.Pointer(_lv_mm.values))
		_lv_entry_First := (float64)(_lv_First_CArray[0])

		_lv_second_goptr := newQPointF(_lv_Second_CArray[0])
		_lv_second_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_lv_entry_Second := *_lv_second_goptr

		_ret[i] = struct {
			First  float64
			Second QPointF
		}{First: _lv_entry_First, Second: _lv_entry_Second}
	}
	return _ret
}

func (this *QGraphicsItemAnimation) SetShearAt(step float64, sh float64, sv float64) {
	QGraphicsItemAnimation_SetShearAt(this.h, (double)(step), (double)(sh), (double)(sv))
}

func (this *QGraphicsItemAnimation) Clear() {
	QGraphicsItemAnimation_Clear(this.h)
}

func (this *QGraphicsItemAnimation) SetStep(x float64) {
	QGraphicsItemAnimation_SetStep(this.h, (double)(x))
}

func QGraphicsItemAnimation_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsItemAnimation_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsItemAnimation_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsItemAnimation_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsItemAnimation) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsItemAnimation_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QGraphicsItemAnimation) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsItemAnimation_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsItemAnimation_MetaObject
func miqt_exec_callback_QGraphicsItemAnimation_MetaObject(self QGraphicsItemAnimation, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsItemAnimation{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QGraphicsItemAnimation) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGraphicsItemAnimation_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGraphicsItemAnimation) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsItemAnimation_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsItemAnimation_Metacast
func miqt_exec_callback_QGraphicsItemAnimation_Metacast(self QGraphicsItemAnimation, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QGraphicsItemAnimation{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
