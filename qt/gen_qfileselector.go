package qt

import (
	"unsafe"
)

type QFileSelector struct {
	h          uintptr
	isSubclass bool
}

// NewQFileSelector constructs a new QFileSelector object.
func NewQFileSelector() *QFileSelector {
	g := newQFileSelector(QFileSelector_new())
	g.isSubclass = true
	return g
}

// NewQFileSelector2 constructs a new QFileSelector object.
func NewQFileSelector2(parent *QObject) *QFileSelector {
	g := newQFileSelector(QFileSelector_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QFileSelector) MetaObject() *QMetaObject {
	return newQMetaObject(QFileSelector_MetaObject(this.h))
}

func (this *QFileSelector) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFileSelector_Metacast(this.h, param1_Cstring))
}

func QFileSelector_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFileSelector_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSelector) Select(filePath string) string {
	filePath_ms := struct_miqt_string{}
	filePath_ms.data = CString(filePath)
	filePath_ms.len = size_t(len(filePath))
	defer free(unsafe.Pointer(filePath_ms.data))
	var _ms struct_miqt_string = QFileSelector_Select(this.h, filePath_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSelector) SelectWithFilePath(filePath *QUrl) *QUrl {
	_goptr := newQUrl(QFileSelector_SelectWithFilePath(this.h, filePath.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFileSelector) ExtraSelectors() []string {
	var _ma struct_miqt_array = QFileSelector_ExtraSelectors(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QFileSelector) SetExtraSelectors(list []string) {
	list_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(list))))
	defer free(unsafe.Pointer(list_CArray))
	for i := range list {
		list_i_ms := struct_miqt_string{}
		list_i_ms.data = CString(list[i])
		list_i_ms.len = size_t(len(list[i]))
		defer free(unsafe.Pointer(list_i_ms.data))
		list_CArray[i] = list_i_ms
	}
	list_ma := struct_miqt_array{len: size_t(len(list)), data: unsafe.Pointer(list_CArray)}
	QFileSelector_SetExtraSelectors(this.h, list_ma)
}

func (this *QFileSelector) AllSelectors() []string {
	var _ma struct_miqt_array = QFileSelector_AllSelectors(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QFileSelector_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFileSelector_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFileSelector_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFileSelector_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFileSelector) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QFileSelector_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QFileSelector) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSelector_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSelector_MetaObject
func miqt_exec_callback_QFileSelector_MetaObject(self QFileSelector, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFileSelector{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QFileSelector) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QFileSelector_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QFileSelector) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFileSelector_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFileSelector_Metacast
func miqt_exec_callback_QFileSelector_Metacast(self QFileSelector, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QFileSelector{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
