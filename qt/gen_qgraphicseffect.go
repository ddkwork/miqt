package qt

import (
	"unsafe"
)

type QGraphicsEffect__ChangeFlag int

const (
	QGraphicsEffect__SourceAttached            QGraphicsEffect__ChangeFlag = 1
	QGraphicsEffect__SourceDetached            QGraphicsEffect__ChangeFlag = 2
	QGraphicsEffect__SourceBoundingRectChanged QGraphicsEffect__ChangeFlag = 4
	QGraphicsEffect__SourceInvalidated         QGraphicsEffect__ChangeFlag = 8
)

type QGraphicsEffect__PixmapPadMode int

const (
	QGraphicsEffect__NoPad                      QGraphicsEffect__PixmapPadMode = 0
	QGraphicsEffect__PadToTransparentBorder     QGraphicsEffect__PixmapPadMode = 1
	QGraphicsEffect__PadToEffectiveBoundingRect QGraphicsEffect__PixmapPadMode = 2
)

type QGraphicsBlurEffect__BlurHint int

const (
	QGraphicsBlurEffect__PerformanceHint QGraphicsBlurEffect__BlurHint = 0
	QGraphicsBlurEffect__QualityHint     QGraphicsBlurEffect__BlurHint = 1
	QGraphicsBlurEffect__AnimationHint   QGraphicsBlurEffect__BlurHint = 2
)

type QGraphicsEffect struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsEffect constructs a new QGraphicsEffect object.
func NewQGraphicsEffect() *QGraphicsEffect {

	ret := newQGraphicsEffect(QGraphicsEffect_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsEffect2 constructs a new QGraphicsEffect object.
func NewQGraphicsEffect2(parent *QObject) *QGraphicsEffect {

	ret := newQGraphicsEffect(QGraphicsEffect_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsEffect) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsEffect_MetaObject(this.h))
}

func (this *QGraphicsEffect) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsEffect_Metacast(this.h, param1_Cstring))
}

func QGraphicsEffect_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsEffect_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsEffect) BoundingRectFor(sourceRect *QRectF) *QRectF {
	_goptr := newQRectF(QGraphicsEffect_BoundingRectFor(this.h, sourceRect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsEffect) BoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsEffect_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsEffect) IsEnabled() bool {
	return (bool)(QGraphicsEffect_IsEnabled(this.h))
}

func (this *QGraphicsEffect) SetEnabled(enable bool) {
	QGraphicsEffect_SetEnabled(this.h, (bool)(enable))
}

func (this *QGraphicsEffect) Update() {
	QGraphicsEffect_Update(this.h)
}

func (this *QGraphicsEffect) EnabledChanged(enabled bool) {
	QGraphicsEffect_EnabledChanged(this.h, (bool)(enabled))
}
func (this *QGraphicsEffect) OnEnabledChanged(slot func(enabled bool)) {
	QGraphicsEffect_connect_EnabledChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsEffect_EnabledChanged
func miqt_exec_callback_QGraphicsEffect_EnabledChanged(cb intptr_t, enabled bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(enabled bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(enabled)

	gofunc(slotval1)
}

func QGraphicsEffect_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsEffect_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsEffect_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsEffect_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsEffect) callVirtualBase_BoundingRectFor(sourceRect *QRectF) *QRectF {

	_goptr := newQRectF(QGraphicsEffect_virtualbase_BoundingRectFor(unsafe.Pointer(this.h), sourceRect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsEffect) OnBoundingRectFor(slot func(super func(sourceRect *QRectF) *QRectF, sourceRect *QRectF) *QRectF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsEffect_override_virtual_BoundingRectFor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsEffect_BoundingRectFor
func miqt_exec_callback_QGraphicsEffect_BoundingRectFor(self QGraphicsEffect, cb intptr_t, sourceRect *QRectF) *QRectF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sourceRect *QRectF) *QRectF, sourceRect *QRectF) *QRectF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(sourceRect)

	virtualReturn := gofunc((&QGraphicsEffect{h: self}).callVirtualBase_BoundingRectFor, slotval1)

	return virtualReturn.cPointer()

}
func (this *QGraphicsEffect) OnDraw(slot func(painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsEffect_override_virtual_Draw(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsEffect_Draw
func miqt_exec_callback_QGraphicsEffect_Draw(self QGraphicsEffect, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc(slotval1)

}

func (this *QGraphicsEffect) callVirtualBase_SourceChanged(flags ChangeFlags) {

	QGraphicsEffect_virtualbase_SourceChanged(unsafe.Pointer(this.h), flags)

}
func (this *QGraphicsEffect) OnSourceChanged(slot func(super func(flags ChangeFlags), flags ChangeFlags)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsEffect_override_virtual_SourceChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsEffect_SourceChanged
func miqt_exec_callback_QGraphicsEffect_SourceChanged(self QGraphicsEffect, cb intptr_t, flags ChangeFlags) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(flags ChangeFlags), flags ChangeFlags))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QGraphicsEffect{h: self}).callVirtualBase_SourceChanged, slotval1)

}

func (this *QGraphicsEffect) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QGraphicsEffect_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QGraphicsEffect) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsEffect_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsEffect_Event
func miqt_exec_callback_QGraphicsEffect_Event(self QGraphicsEffect, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QGraphicsEffect{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsEffect) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QGraphicsEffect_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QGraphicsEffect) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsEffect_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsEffect_EventFilter
func miqt_exec_callback_QGraphicsEffect_EventFilter(self QGraphicsEffect, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QGraphicsEffect{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QGraphicsEffect) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QGraphicsEffect_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsEffect) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsEffect_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsEffect_TimerEvent
func miqt_exec_callback_QGraphicsEffect_TimerEvent(self QGraphicsEffect, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QGraphicsEffect{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QGraphicsEffect) callVirtualBase_ChildEvent(event *QChildEvent) {

	QGraphicsEffect_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsEffect) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsEffect_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsEffect_ChildEvent
func miqt_exec_callback_QGraphicsEffect_ChildEvent(self QGraphicsEffect, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QGraphicsEffect{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QGraphicsEffect) callVirtualBase_CustomEvent(event *QEvent) {

	QGraphicsEffect_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsEffect) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsEffect_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsEffect_CustomEvent
func miqt_exec_callback_QGraphicsEffect_CustomEvent(self QGraphicsEffect, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGraphicsEffect{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QGraphicsEffect) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QGraphicsEffect_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QGraphicsEffect) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsEffect_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsEffect_ConnectNotify
func miqt_exec_callback_QGraphicsEffect_ConnectNotify(self QGraphicsEffect, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QGraphicsEffect{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QGraphicsEffect) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QGraphicsEffect_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QGraphicsEffect) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsEffect_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsEffect_DisconnectNotify
func miqt_exec_callback_QGraphicsEffect_DisconnectNotify(self QGraphicsEffect, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QGraphicsEffect{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

type QGraphicsColorizeEffect struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsColorizeEffect constructs a new QGraphicsColorizeEffect object.
func NewQGraphicsColorizeEffect() *QGraphicsColorizeEffect {

	ret := newQGraphicsColorizeEffect(QGraphicsColorizeEffect_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsColorizeEffect2 constructs a new QGraphicsColorizeEffect object.
func NewQGraphicsColorizeEffect2(parent *QObject) *QGraphicsColorizeEffect {

	ret := newQGraphicsColorizeEffect(QGraphicsColorizeEffect_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsColorizeEffect) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsColorizeEffect_MetaObject(this.h))
}

func (this *QGraphicsColorizeEffect) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsColorizeEffect_Metacast(this.h, param1_Cstring))
}

func QGraphicsColorizeEffect_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsColorizeEffect_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsColorizeEffect) Color() *QColor {
	_goptr := newQColor(QGraphicsColorizeEffect_Color(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsColorizeEffect) Strength() float64 {
	return (float64)(QGraphicsColorizeEffect_Strength(this.h))
}

func (this *QGraphicsColorizeEffect) SetColor(c *QColor) {
	QGraphicsColorizeEffect_SetColor(this.h, c.cPointer())
}

func (this *QGraphicsColorizeEffect) SetStrength(strength float64) {
	QGraphicsColorizeEffect_SetStrength(this.h, (double)(strength))
}

func (this *QGraphicsColorizeEffect) ColorChanged(color *QColor) {
	QGraphicsColorizeEffect_ColorChanged(this.h, color.cPointer())
}
func (this *QGraphicsColorizeEffect) OnColorChanged(slot func(color *QColor)) {
	QGraphicsColorizeEffect_connect_ColorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsColorizeEffect_ColorChanged
func miqt_exec_callback_QGraphicsColorizeEffect_ColorChanged(cb intptr_t, color *QColor) {
	gofunc, ok := cgo.Handle(cb).Value().(func(color *QColor))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQColor(color)

	gofunc(slotval1)
}

func (this *QGraphicsColorizeEffect) StrengthChanged(strength float64) {
	QGraphicsColorizeEffect_StrengthChanged(this.h, (double)(strength))
}
func (this *QGraphicsColorizeEffect) OnStrengthChanged(slot func(strength float64)) {
	QGraphicsColorizeEffect_connect_StrengthChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsColorizeEffect_StrengthChanged
func miqt_exec_callback_QGraphicsColorizeEffect_StrengthChanged(cb intptr_t, strength double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(strength float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(strength)

	gofunc(slotval1)
}

func QGraphicsColorizeEffect_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsColorizeEffect_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsColorizeEffect_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsColorizeEffect_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsColorizeEffect) callVirtualBase_Draw(painter *QPainter) {

	QGraphicsColorizeEffect_virtualbase_Draw(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QGraphicsColorizeEffect) OnDraw(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsColorizeEffect_override_virtual_Draw(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsColorizeEffect_Draw
func miqt_exec_callback_QGraphicsColorizeEffect_Draw(self QGraphicsColorizeEffect, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QGraphicsColorizeEffect{h: self}).callVirtualBase_Draw, slotval1)

}

func (this *QGraphicsColorizeEffect) callVirtualBase_BoundingRectFor(sourceRect *QRectF) *QRectF {

	_goptr := newQRectF(QGraphicsColorizeEffect_virtualbase_BoundingRectFor(unsafe.Pointer(this.h), sourceRect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsColorizeEffect) OnBoundingRectFor(slot func(super func(sourceRect *QRectF) *QRectF, sourceRect *QRectF) *QRectF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsColorizeEffect_override_virtual_BoundingRectFor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsColorizeEffect_BoundingRectFor
func miqt_exec_callback_QGraphicsColorizeEffect_BoundingRectFor(self QGraphicsColorizeEffect, cb intptr_t, sourceRect *QRectF) *QRectF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sourceRect *QRectF) *QRectF, sourceRect *QRectF) *QRectF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(sourceRect)

	virtualReturn := gofunc((&QGraphicsColorizeEffect{h: self}).callVirtualBase_BoundingRectFor, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGraphicsColorizeEffect) callVirtualBase_SourceChanged(flags ChangeFlags) {

	QGraphicsColorizeEffect_virtualbase_SourceChanged(unsafe.Pointer(this.h), flags)

}
func (this *QGraphicsColorizeEffect) OnSourceChanged(slot func(super func(flags ChangeFlags), flags ChangeFlags)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsColorizeEffect_override_virtual_SourceChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsColorizeEffect_SourceChanged
func miqt_exec_callback_QGraphicsColorizeEffect_SourceChanged(self QGraphicsColorizeEffect, cb intptr_t, flags ChangeFlags) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(flags ChangeFlags), flags ChangeFlags))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QGraphicsColorizeEffect{h: self}).callVirtualBase_SourceChanged, slotval1)

}

type QGraphicsBlurEffect struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsBlurEffect constructs a new QGraphicsBlurEffect object.
func NewQGraphicsBlurEffect() *QGraphicsBlurEffect {

	ret := newQGraphicsBlurEffect(QGraphicsBlurEffect_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsBlurEffect2 constructs a new QGraphicsBlurEffect object.
func NewQGraphicsBlurEffect2(parent *QObject) *QGraphicsBlurEffect {

	ret := newQGraphicsBlurEffect(QGraphicsBlurEffect_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsBlurEffect) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsBlurEffect_MetaObject(this.h))
}

func (this *QGraphicsBlurEffect) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsBlurEffect_Metacast(this.h, param1_Cstring))
}

func QGraphicsBlurEffect_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsBlurEffect_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsBlurEffect) BoundingRectFor(rect *QRectF) *QRectF {
	_goptr := newQRectF(QGraphicsBlurEffect_BoundingRectFor(this.h, rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsBlurEffect) BlurRadius() float64 {
	return (float64)(QGraphicsBlurEffect_BlurRadius(this.h))
}

func (this *QGraphicsBlurEffect) BlurHints() BlurHints {
	xxxxxxxxx
}

func (this *QGraphicsBlurEffect) SetBlurRadius(blurRadius float64) {
	QGraphicsBlurEffect_SetBlurRadius(this.h, (double)(blurRadius))
}

func (this *QGraphicsBlurEffect) SetBlurHints(hints BlurHints) {
	QGraphicsBlurEffect_SetBlurHints(this.h, hints)
}

func (this *QGraphicsBlurEffect) BlurRadiusChanged(blurRadius float64) {
	QGraphicsBlurEffect_BlurRadiusChanged(this.h, (double)(blurRadius))
}
func (this *QGraphicsBlurEffect) OnBlurRadiusChanged(slot func(blurRadius float64)) {
	QGraphicsBlurEffect_connect_BlurRadiusChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsBlurEffect_BlurRadiusChanged
func miqt_exec_callback_QGraphicsBlurEffect_BlurRadiusChanged(cb intptr_t, blurRadius double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(blurRadius float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(blurRadius)

	gofunc(slotval1)
}

func (this *QGraphicsBlurEffect) BlurHintsChanged(hints BlurHints) {
	QGraphicsBlurEffect_BlurHintsChanged(this.h, hints)
}
func (this *QGraphicsBlurEffect) OnBlurHintsChanged(slot func(hints BlurHints)) {
	QGraphicsBlurEffect_connect_BlurHintsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsBlurEffect_BlurHintsChanged
func miqt_exec_callback_QGraphicsBlurEffect_BlurHintsChanged(cb intptr_t, hints BlurHints) {
	gofunc, ok := cgo.Handle(cb).Value().(func(hints BlurHints))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc(slotval1)
}

func QGraphicsBlurEffect_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsBlurEffect_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsBlurEffect_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsBlurEffect_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsBlurEffect) callVirtualBase_BoundingRectFor(rect *QRectF) *QRectF {

	_goptr := newQRectF(QGraphicsBlurEffect_virtualbase_BoundingRectFor(unsafe.Pointer(this.h), rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsBlurEffect) OnBoundingRectFor(slot func(super func(rect *QRectF) *QRectF, rect *QRectF) *QRectF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsBlurEffect_override_virtual_BoundingRectFor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsBlurEffect_BoundingRectFor
func miqt_exec_callback_QGraphicsBlurEffect_BoundingRectFor(self QGraphicsBlurEffect, cb intptr_t, rect *QRectF) *QRectF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRectF) *QRectF, rect *QRectF) *QRectF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(rect)

	virtualReturn := gofunc((&QGraphicsBlurEffect{h: self}).callVirtualBase_BoundingRectFor, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGraphicsBlurEffect) callVirtualBase_Draw(painter *QPainter) {

	QGraphicsBlurEffect_virtualbase_Draw(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QGraphicsBlurEffect) OnDraw(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsBlurEffect_override_virtual_Draw(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsBlurEffect_Draw
func miqt_exec_callback_QGraphicsBlurEffect_Draw(self QGraphicsBlurEffect, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QGraphicsBlurEffect{h: self}).callVirtualBase_Draw, slotval1)

}

func (this *QGraphicsBlurEffect) callVirtualBase_SourceChanged(flags ChangeFlags) {

	QGraphicsBlurEffect_virtualbase_SourceChanged(unsafe.Pointer(this.h), flags)

}
func (this *QGraphicsBlurEffect) OnSourceChanged(slot func(super func(flags ChangeFlags), flags ChangeFlags)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsBlurEffect_override_virtual_SourceChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsBlurEffect_SourceChanged
func miqt_exec_callback_QGraphicsBlurEffect_SourceChanged(self QGraphicsBlurEffect, cb intptr_t, flags ChangeFlags) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(flags ChangeFlags), flags ChangeFlags))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QGraphicsBlurEffect{h: self}).callVirtualBase_SourceChanged, slotval1)

}

type QGraphicsDropShadowEffect struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsDropShadowEffect constructs a new QGraphicsDropShadowEffect object.
func NewQGraphicsDropShadowEffect() *QGraphicsDropShadowEffect {

	ret := newQGraphicsDropShadowEffect(QGraphicsDropShadowEffect_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsDropShadowEffect2 constructs a new QGraphicsDropShadowEffect object.
func NewQGraphicsDropShadowEffect2(parent *QObject) *QGraphicsDropShadowEffect {

	ret := newQGraphicsDropShadowEffect(QGraphicsDropShadowEffect_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsDropShadowEffect) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsDropShadowEffect_MetaObject(this.h))
}

func (this *QGraphicsDropShadowEffect) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsDropShadowEffect_Metacast(this.h, param1_Cstring))
}

func QGraphicsDropShadowEffect_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsDropShadowEffect_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsDropShadowEffect) BoundingRectFor(rect *QRectF) *QRectF {
	_goptr := newQRectF(QGraphicsDropShadowEffect_BoundingRectFor(this.h, rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsDropShadowEffect) Offset() *QPointF {
	_goptr := newQPointF(QGraphicsDropShadowEffect_Offset(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsDropShadowEffect) XOffset() float64 {
	return (float64)(QGraphicsDropShadowEffect_XOffset(this.h))
}

func (this *QGraphicsDropShadowEffect) YOffset() float64 {
	return (float64)(QGraphicsDropShadowEffect_YOffset(this.h))
}

func (this *QGraphicsDropShadowEffect) BlurRadius() float64 {
	return (float64)(QGraphicsDropShadowEffect_BlurRadius(this.h))
}

func (this *QGraphicsDropShadowEffect) Color() *QColor {
	_goptr := newQColor(QGraphicsDropShadowEffect_Color(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsDropShadowEffect) SetOffset(ofs *QPointF) {
	QGraphicsDropShadowEffect_SetOffset(this.h, ofs.cPointer())
}

func (this *QGraphicsDropShadowEffect) SetOffset2(dx float64, dy float64) {
	QGraphicsDropShadowEffect_SetOffset2(this.h, (double)(dx), (double)(dy))
}

func (this *QGraphicsDropShadowEffect) SetOffsetWithQreal(d float64) {
	QGraphicsDropShadowEffect_SetOffsetWithQreal(this.h, (double)(d))
}

func (this *QGraphicsDropShadowEffect) SetXOffset(dx float64) {
	QGraphicsDropShadowEffect_SetXOffset(this.h, (double)(dx))
}

func (this *QGraphicsDropShadowEffect) SetYOffset(dy float64) {
	QGraphicsDropShadowEffect_SetYOffset(this.h, (double)(dy))
}

func (this *QGraphicsDropShadowEffect) SetBlurRadius(blurRadius float64) {
	QGraphicsDropShadowEffect_SetBlurRadius(this.h, (double)(blurRadius))
}

func (this *QGraphicsDropShadowEffect) SetColor(color *QColor) {
	QGraphicsDropShadowEffect_SetColor(this.h, color.cPointer())
}

func (this *QGraphicsDropShadowEffect) OffsetChanged(offset *QPointF) {
	QGraphicsDropShadowEffect_OffsetChanged(this.h, offset.cPointer())
}
func (this *QGraphicsDropShadowEffect) OnOffsetChanged(slot func(offset *QPointF)) {
	QGraphicsDropShadowEffect_connect_OffsetChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsDropShadowEffect_OffsetChanged
func miqt_exec_callback_QGraphicsDropShadowEffect_OffsetChanged(cb intptr_t, offset *QPointF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(offset *QPointF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPointF(offset)

	gofunc(slotval1)
}

func (this *QGraphicsDropShadowEffect) BlurRadiusChanged(blurRadius float64) {
	QGraphicsDropShadowEffect_BlurRadiusChanged(this.h, (double)(blurRadius))
}
func (this *QGraphicsDropShadowEffect) OnBlurRadiusChanged(slot func(blurRadius float64)) {
	QGraphicsDropShadowEffect_connect_BlurRadiusChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsDropShadowEffect_BlurRadiusChanged
func miqt_exec_callback_QGraphicsDropShadowEffect_BlurRadiusChanged(cb intptr_t, blurRadius double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(blurRadius float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(blurRadius)

	gofunc(slotval1)
}

func (this *QGraphicsDropShadowEffect) ColorChanged(color *QColor) {
	QGraphicsDropShadowEffect_ColorChanged(this.h, color.cPointer())
}
func (this *QGraphicsDropShadowEffect) OnColorChanged(slot func(color *QColor)) {
	QGraphicsDropShadowEffect_connect_ColorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsDropShadowEffect_ColorChanged
func miqt_exec_callback_QGraphicsDropShadowEffect_ColorChanged(cb intptr_t, color *QColor) {
	gofunc, ok := cgo.Handle(cb).Value().(func(color *QColor))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQColor(color)

	gofunc(slotval1)
}

func QGraphicsDropShadowEffect_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsDropShadowEffect_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsDropShadowEffect_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsDropShadowEffect_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsDropShadowEffect) callVirtualBase_BoundingRectFor(rect *QRectF) *QRectF {

	_goptr := newQRectF(QGraphicsDropShadowEffect_virtualbase_BoundingRectFor(unsafe.Pointer(this.h), rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsDropShadowEffect) OnBoundingRectFor(slot func(super func(rect *QRectF) *QRectF, rect *QRectF) *QRectF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsDropShadowEffect_override_virtual_BoundingRectFor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsDropShadowEffect_BoundingRectFor
func miqt_exec_callback_QGraphicsDropShadowEffect_BoundingRectFor(self QGraphicsDropShadowEffect, cb intptr_t, rect *QRectF) *QRectF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRectF) *QRectF, rect *QRectF) *QRectF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(rect)

	virtualReturn := gofunc((&QGraphicsDropShadowEffect{h: self}).callVirtualBase_BoundingRectFor, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGraphicsDropShadowEffect) callVirtualBase_Draw(painter *QPainter) {

	QGraphicsDropShadowEffect_virtualbase_Draw(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QGraphicsDropShadowEffect) OnDraw(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsDropShadowEffect_override_virtual_Draw(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsDropShadowEffect_Draw
func miqt_exec_callback_QGraphicsDropShadowEffect_Draw(self QGraphicsDropShadowEffect, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QGraphicsDropShadowEffect{h: self}).callVirtualBase_Draw, slotval1)

}

func (this *QGraphicsDropShadowEffect) callVirtualBase_SourceChanged(flags ChangeFlags) {

	QGraphicsDropShadowEffect_virtualbase_SourceChanged(unsafe.Pointer(this.h), flags)

}
func (this *QGraphicsDropShadowEffect) OnSourceChanged(slot func(super func(flags ChangeFlags), flags ChangeFlags)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsDropShadowEffect_override_virtual_SourceChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsDropShadowEffect_SourceChanged
func miqt_exec_callback_QGraphicsDropShadowEffect_SourceChanged(self QGraphicsDropShadowEffect, cb intptr_t, flags ChangeFlags) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(flags ChangeFlags), flags ChangeFlags))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QGraphicsDropShadowEffect{h: self}).callVirtualBase_SourceChanged, slotval1)

}

type QGraphicsOpacityEffect struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsOpacityEffect constructs a new QGraphicsOpacityEffect object.
func NewQGraphicsOpacityEffect() *QGraphicsOpacityEffect {

	ret := newQGraphicsOpacityEffect(QGraphicsOpacityEffect_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsOpacityEffect2 constructs a new QGraphicsOpacityEffect object.
func NewQGraphicsOpacityEffect2(parent *QObject) *QGraphicsOpacityEffect {

	ret := newQGraphicsOpacityEffect(QGraphicsOpacityEffect_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsOpacityEffect) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsOpacityEffect_MetaObject(this.h))
}

func (this *QGraphicsOpacityEffect) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsOpacityEffect_Metacast(this.h, param1_Cstring))
}

func QGraphicsOpacityEffect_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsOpacityEffect_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsOpacityEffect) Opacity() float64 {
	return (float64)(QGraphicsOpacityEffect_Opacity(this.h))
}

func (this *QGraphicsOpacityEffect) OpacityMask() *QBrush {
	_goptr := newQBrush(QGraphicsOpacityEffect_OpacityMask(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsOpacityEffect) SetOpacity(opacity float64) {
	QGraphicsOpacityEffect_SetOpacity(this.h, (double)(opacity))
}

func (this *QGraphicsOpacityEffect) SetOpacityMask(mask *QBrush) {
	QGraphicsOpacityEffect_SetOpacityMask(this.h, mask.cPointer())
}

func (this *QGraphicsOpacityEffect) OpacityChanged(opacity float64) {
	QGraphicsOpacityEffect_OpacityChanged(this.h, (double)(opacity))
}
func (this *QGraphicsOpacityEffect) OnOpacityChanged(slot func(opacity float64)) {
	QGraphicsOpacityEffect_connect_OpacityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsOpacityEffect_OpacityChanged
func miqt_exec_callback_QGraphicsOpacityEffect_OpacityChanged(cb intptr_t, opacity double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(opacity float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(opacity)

	gofunc(slotval1)
}

func (this *QGraphicsOpacityEffect) OpacityMaskChanged(mask *QBrush) {
	QGraphicsOpacityEffect_OpacityMaskChanged(this.h, mask.cPointer())
}
func (this *QGraphicsOpacityEffect) OnOpacityMaskChanged(slot func(mask *QBrush)) {
	QGraphicsOpacityEffect_connect_OpacityMaskChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsOpacityEffect_OpacityMaskChanged
func miqt_exec_callback_QGraphicsOpacityEffect_OpacityMaskChanged(cb intptr_t, mask *QBrush) {
	gofunc, ok := cgo.Handle(cb).Value().(func(mask *QBrush))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQBrush(mask)

	gofunc(slotval1)
}

func QGraphicsOpacityEffect_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsOpacityEffect_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsOpacityEffect_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsOpacityEffect_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsOpacityEffect) callVirtualBase_Draw(painter *QPainter) {

	QGraphicsOpacityEffect_virtualbase_Draw(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QGraphicsOpacityEffect) OnDraw(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsOpacityEffect_override_virtual_Draw(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsOpacityEffect_Draw
func miqt_exec_callback_QGraphicsOpacityEffect_Draw(self QGraphicsOpacityEffect, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QGraphicsOpacityEffect{h: self}).callVirtualBase_Draw, slotval1)

}

func (this *QGraphicsOpacityEffect) callVirtualBase_BoundingRectFor(sourceRect *QRectF) *QRectF {

	_goptr := newQRectF(QGraphicsOpacityEffect_virtualbase_BoundingRectFor(unsafe.Pointer(this.h), sourceRect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsOpacityEffect) OnBoundingRectFor(slot func(super func(sourceRect *QRectF) *QRectF, sourceRect *QRectF) *QRectF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsOpacityEffect_override_virtual_BoundingRectFor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsOpacityEffect_BoundingRectFor
func miqt_exec_callback_QGraphicsOpacityEffect_BoundingRectFor(self QGraphicsOpacityEffect, cb intptr_t, sourceRect *QRectF) *QRectF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(sourceRect *QRectF) *QRectF, sourceRect *QRectF) *QRectF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(sourceRect)

	virtualReturn := gofunc((&QGraphicsOpacityEffect{h: self}).callVirtualBase_BoundingRectFor, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGraphicsOpacityEffect) callVirtualBase_SourceChanged(flags ChangeFlags) {

	QGraphicsOpacityEffect_virtualbase_SourceChanged(unsafe.Pointer(this.h), flags)

}
func (this *QGraphicsOpacityEffect) OnSourceChanged(slot func(super func(flags ChangeFlags), flags ChangeFlags)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsOpacityEffect_override_virtual_SourceChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsOpacityEffect_SourceChanged
func miqt_exec_callback_QGraphicsOpacityEffect_SourceChanged(self QGraphicsOpacityEffect, cb intptr_t, flags ChangeFlags) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(flags ChangeFlags), flags ChangeFlags))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QGraphicsOpacityEffect{h: self}).callVirtualBase_SourceChanged, slotval1)

}
