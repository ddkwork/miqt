package qt

import (
	"unsafe"
)

type QSequentialAnimationGroup struct {
	h          uintptr
	isSubclass bool
}

// NewQSequentialAnimationGroup constructs a new QSequentialAnimationGroup object.
func NewQSequentialAnimationGroup() *QSequentialAnimationGroup {
	g := newQSequentialAnimationGroup(QSequentialAnimationGroup_new())
	g.isSubclass = true
	return g
}

// NewQSequentialAnimationGroup2 constructs a new QSequentialAnimationGroup object.
func NewQSequentialAnimationGroup2(parent *QObject) *QSequentialAnimationGroup {
	g := newQSequentialAnimationGroup(QSequentialAnimationGroup_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QSequentialAnimationGroup) MetaObject() *QMetaObject {
	return newQMetaObject(QSequentialAnimationGroup_MetaObject(this.h))
}

func (this *QSequentialAnimationGroup) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSequentialAnimationGroup_Metacast(this.h, param1_Cstring))
}

func QSequentialAnimationGroup_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSequentialAnimationGroup_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSequentialAnimationGroup) AddPause(msecs int) *QPauseAnimation {
	return newQPauseAnimation(QSequentialAnimationGroup_AddPause(this.h, (int)(msecs)))
}

func (this *QSequentialAnimationGroup) InsertPause(index int, msecs int) *QPauseAnimation {
	return newQPauseAnimation(QSequentialAnimationGroup_InsertPause(this.h, (int)(index), (int)(msecs)))
}

func (this *QSequentialAnimationGroup) CurrentAnimation() *QAbstractAnimation {
	return newQAbstractAnimation(QSequentialAnimationGroup_CurrentAnimation(this.h))
}

func (this *QSequentialAnimationGroup) Duration() int {
	return (int)(QSequentialAnimationGroup_Duration(this.h))
}

func (this *QSequentialAnimationGroup) CurrentAnimationChanged(current *QAbstractAnimation) {
	QSequentialAnimationGroup_CurrentAnimationChanged(this.h, current.cPointer())
}

func (this *QSequentialAnimationGroup) OnCurrentAnimationChanged(slot func(current *QAbstractAnimation)) {
	QSequentialAnimationGroup_connect_CurrentAnimationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSequentialAnimationGroup_CurrentAnimationChanged
func miqt_exec_callback_QSequentialAnimationGroup_CurrentAnimationChanged(cb intptr_t, current *QAbstractAnimation) {
	gofunc, ok := cgo.Handle(cb).Value().(func(current *QAbstractAnimation))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractAnimation(current)

	gofunc(slotval1)
}

func QSequentialAnimationGroup_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSequentialAnimationGroup_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSequentialAnimationGroup_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSequentialAnimationGroup_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSequentialAnimationGroup) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSequentialAnimationGroup_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSequentialAnimationGroup) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSequentialAnimationGroup_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSequentialAnimationGroup_MetaObject
func miqt_exec_callback_QSequentialAnimationGroup_MetaObject(self QSequentialAnimationGroup, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSequentialAnimationGroup{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSequentialAnimationGroup) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSequentialAnimationGroup_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSequentialAnimationGroup) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSequentialAnimationGroup_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSequentialAnimationGroup_Metacast
func miqt_exec_callback_QSequentialAnimationGroup_Metacast(self QSequentialAnimationGroup, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QSequentialAnimationGroup{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
