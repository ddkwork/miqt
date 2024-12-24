package spatialaudio

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QAudioListener struct {
	h          uintptr
	isSubclass bool
}

// NewQAudioListener constructs a new QAudioListener object.
func NewQAudioListener(engine *QAudioEngine) *QAudioListener {
	g := newQAudioListener(QAudioListener_new(engine.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QAudioListener) SetPosition(pos qt.QVector3D) {
	QAudioListener_SetPosition(this.h, (*QVector3D)(pos.UnsafePointer()))
}

func (this *QAudioListener) Position() *qt.QVector3D {
	_goptr := qt.UnsafeNewQVector3D(unsafe.Pointer(QAudioListener_Position(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioListener) SetRotation(q *qt.QQuaternion) {
	QAudioListener_SetRotation(this.h, (*QQuaternion)(q.UnsafePointer()))
}

func (this *QAudioListener) Rotation() *qt.QQuaternion {
	_goptr := qt.UnsafeNewQQuaternion(unsafe.Pointer(QAudioListener_Rotation(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAudioListener) Engine() *QAudioEngine {
	return newQAudioEngine(QAudioListener_Engine(this.h))
}

func (this *QAudioListener) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAudioListener_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QAudioListener) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioListener_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioListener_MetaObject
func miqt_exec_callback_QAudioListener_MetaObject(self QAudioListener, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAudioListener{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QAudioListener) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAudioListener_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAudioListener) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAudioListener_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAudioListener_Metacast
func miqt_exec_callback_QAudioListener_Metacast(self QAudioListener, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QAudioListener{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
