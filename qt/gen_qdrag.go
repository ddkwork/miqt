package qt

import (
	"unsafe"
)

type QDrag struct {
	h          uintptr
	isSubclass bool
}

// NewQDrag constructs a new QDrag object.
func NewQDrag(dragSource *QObject) *QDrag {
	g := newQDrag(QDrag_new(dragSource.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QDrag) MetaObject() *QMetaObject {
	return newQMetaObject(QDrag_MetaObject(this.h))
}

func (this *QDrag) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDrag_Metacast(this.h, param1_Cstring))
}

func QDrag_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDrag_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDrag) SetMimeData(data *QMimeData) {
	QDrag_SetMimeData(this.h, data.cPointer())
}

func (this *QDrag) MimeData() *QMimeData {
	return newQMimeData(QDrag_MimeData(this.h))
}

func (this *QDrag) SetPixmap(pixmap *QPixmap) {
	QDrag_SetPixmap(this.h, pixmap.cPointer())
}

func (this *QDrag) Pixmap() *QPixmap {
	_goptr := newQPixmap(QDrag_Pixmap(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDrag) SetHotSpot(hotspot *QPoint) {
	QDrag_SetHotSpot(this.h, hotspot.cPointer())
}

func (this *QDrag) HotSpot() *QPoint {
	_goptr := newQPoint(QDrag_HotSpot(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDrag) Source() *QObject {
	return newQObject(QDrag_Source(this.h))
}

func (this *QDrag) Target() *QObject {
	return newQObject(QDrag_Target(this.h))
}

func (this *QDrag) Exec() DropAction {
	return (DropAction)(QDrag_Exec(this.h))
}

func (this *QDrag) Exec2(supportedActions DropAction, defaultAction DropAction) DropAction {
	return (DropAction)(QDrag_Exec2(this.h, (int)(supportedActions), (int)(defaultAction)))
}

func (this *QDrag) SetDragCursor(cursor *QPixmap, action DropAction) {
	QDrag_SetDragCursor(this.h, cursor.cPointer(), (int)(action))
}

func (this *QDrag) DragCursor(action DropAction) *QPixmap {
	_goptr := newQPixmap(QDrag_DragCursor(this.h, (int)(action)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDrag) SupportedActions() DropAction {
	return (DropAction)(QDrag_SupportedActions(this.h))
}

func (this *QDrag) DefaultAction() DropAction {
	return (DropAction)(QDrag_DefaultAction(this.h))
}

func QDrag_Cancel() {
	QDrag_Cancel()
}

func (this *QDrag) ActionChanged(action DropAction) {
	QDrag_ActionChanged(this.h, (int)(action))
}

func (this *QDrag) OnActionChanged(slot func(action DropAction)) {
	QDrag_connect_ActionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDrag_ActionChanged
func miqt_exec_callback_QDrag_ActionChanged(cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(action DropAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DropAction)(action)

	gofunc(slotval1)
}

func (this *QDrag) TargetChanged(newTarget *QObject) {
	QDrag_TargetChanged(this.h, newTarget.cPointer())
}

func (this *QDrag) OnTargetChanged(slot func(newTarget *QObject)) {
	QDrag_connect_TargetChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDrag_TargetChanged
func miqt_exec_callback_QDrag_TargetChanged(cb intptr_t, newTarget *QObject) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newTarget *QObject))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(newTarget)

	gofunc(slotval1)
}

func QDrag_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDrag_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDrag_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDrag_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDrag) Exec1(supportedActions DropAction) DropAction {
	return (DropAction)(QDrag_Exec1(this.h, (int)(supportedActions)))
}

func (this *QDrag) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QDrag_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QDrag) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDrag_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDrag_MetaObject
func miqt_exec_callback_QDrag_MetaObject(self QDrag, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDrag{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QDrag) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QDrag_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QDrag) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDrag_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDrag_Metacast
func miqt_exec_callback_QDrag_Metacast(self QDrag, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QDrag{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
