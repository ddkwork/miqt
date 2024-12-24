package svg

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QSvgWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQSvgWidget constructs a new QSvgWidget object.
func NewQSvgWidget(parent *qt.QWidget) *QSvgWidget {
	g := newQSvgWidget(QSvgWidget_new((*QWidget)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQSvgWidget2 constructs a new QSvgWidget object.
func NewQSvgWidget2() *QSvgWidget {
	g := newQSvgWidget(QSvgWidget_new2())
	g.isSubclass = true
	return g
}

// NewQSvgWidget3 constructs a new QSvgWidget object.
func NewQSvgWidget3(file string) *QSvgWidget {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	g := newQSvgWidget(QSvgWidget_new3(file_ms))
	g.isSubclass = true
	return g
}

// NewQSvgWidget4 constructs a new QSvgWidget object.
func NewQSvgWidget4(file string, parent *qt.QWidget) *QSvgWidget {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	g := newQSvgWidget(QSvgWidget_new4(file_ms, (*QWidget)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QSvgWidget) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QSvgWidget_MetaObject(this.h)))
}

func (this *QSvgWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSvgWidget_Metacast(this.h, param1_Cstring))
}

func QSvgWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSvgWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSvgWidget) Renderer() *QSvgRenderer {
	return newQSvgRenderer(QSvgWidget_Renderer(this.h))
}

func (this *QSvgWidget) SizeHint() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QSvgWidget_SizeHint(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSvgWidget) Options() Option {
	return (Option)(QSvgWidget_Options(this.h))
}

func (this *QSvgWidget) SetOptions(options Option) {
	QSvgWidget_SetOptions(this.h, (int)(options))
}

func (this *QSvgWidget) Load(file string) {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	QSvgWidget_Load(this.h, file_ms)
}

func (this *QSvgWidget) LoadWithContents(contents []byte) {
	contents_alias := struct_miqt_string{}
	contents_alias.data = (char)(unsafe.Pointer(&contents[0]))
	contents_alias.len = size_t(len(contents))
	QSvgWidget_LoadWithContents(this.h, contents_alias)
}

func QSvgWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSvgWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSvgWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSvgWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSvgWidget) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QSvgWidget_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QSvgWidget) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_MetaObject
func miqt_exec_callback_QSvgWidget_MetaObject(self QSvgWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QSvgWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSvgWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSvgWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_Metacast
func miqt_exec_callback_QSvgWidget_Metacast(self QSvgWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
