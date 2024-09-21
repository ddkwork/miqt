package qt

/*

#include "gen_qdockwidget.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type QDockWidget__DockWidgetFeature int

const (
	QDockWidget__DockWidgetClosable         QDockWidget__DockWidgetFeature = 1
	QDockWidget__DockWidgetMovable          QDockWidget__DockWidgetFeature = 2
	QDockWidget__DockWidgetFloatable        QDockWidget__DockWidgetFeature = 4
	QDockWidget__DockWidgetVerticalTitleBar QDockWidget__DockWidgetFeature = 8
	QDockWidget__DockWidgetFeatureMask      QDockWidget__DockWidgetFeature = 15
	QDockWidget__AllDockWidgetFeatures      QDockWidget__DockWidgetFeature = 7
	QDockWidget__NoDockWidgetFeatures       QDockWidget__DockWidgetFeature = 0
	QDockWidget__Reserved                   QDockWidget__DockWidgetFeature = 255
)

type QDockWidget struct {
	h *C.QDockWidget
	*QWidget
}

func (this *QDockWidget) cPointer() *C.QDockWidget {
	if this == nil {
		return nil
	}
	return this.h
}

func newQDockWidget(h *C.QDockWidget) *QDockWidget {
	if h == nil {
		return nil
	}
	return &QDockWidget{h: h, QWidget: newQWidget_U(unsafe.Pointer(h))}
}

func newQDockWidget_U(h unsafe.Pointer) *QDockWidget {
	return newQDockWidget((*C.QDockWidget)(h))
}

// NewQDockWidget constructs a new QDockWidget object.
func NewQDockWidget(title string) *QDockWidget {
	title_ms := miqt_strdupg(title)
	defer C.free(title_ms)
	ret := C.QDockWidget_new((*C.struct_miqt_string)(title_ms))
	return newQDockWidget(ret)
}

// NewQDockWidget2 constructs a new QDockWidget object.
func NewQDockWidget2() *QDockWidget {
	ret := C.QDockWidget_new2()
	return newQDockWidget(ret)
}

// NewQDockWidget3 constructs a new QDockWidget object.
func NewQDockWidget3(title string, parent *QWidget) *QDockWidget {
	title_ms := miqt_strdupg(title)
	defer C.free(title_ms)
	ret := C.QDockWidget_new3((*C.struct_miqt_string)(title_ms), parent.cPointer())
	return newQDockWidget(ret)
}

// NewQDockWidget4 constructs a new QDockWidget object.
func NewQDockWidget4(title string, parent *QWidget, flags WindowType) *QDockWidget {
	title_ms := miqt_strdupg(title)
	defer C.free(title_ms)
	ret := C.QDockWidget_new4((*C.struct_miqt_string)(title_ms), parent.cPointer(), (C.int)(flags))
	return newQDockWidget(ret)
}

// NewQDockWidget5 constructs a new QDockWidget object.
func NewQDockWidget5(parent *QWidget) *QDockWidget {
	ret := C.QDockWidget_new5(parent.cPointer())
	return newQDockWidget(ret)
}

// NewQDockWidget6 constructs a new QDockWidget object.
func NewQDockWidget6(parent *QWidget, flags WindowType) *QDockWidget {
	ret := C.QDockWidget_new6(parent.cPointer(), (C.int)(flags))
	return newQDockWidget(ret)
}

func (this *QDockWidget) MetaObject() *QMetaObject {
	return newQMetaObject_U(unsafe.Pointer(C.QDockWidget_MetaObject(this.h)))
}

func (this *QDockWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := C.CString(param1)
	defer C.free(unsafe.Pointer(param1_Cstring))
	return C.QDockWidget_Metacast(this.h, param1_Cstring)
}

func QDockWidget_Tr(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms *C.struct_miqt_string = C.QDockWidget_Tr(s_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QDockWidget_TrUtf8(s string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	var _ms *C.struct_miqt_string = C.QDockWidget_TrUtf8(s_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func (this *QDockWidget) Widget() *QWidget {
	return newQWidget_U(unsafe.Pointer(C.QDockWidget_Widget(this.h)))
}

func (this *QDockWidget) SetWidget(widget *QWidget) {
	C.QDockWidget_SetWidget(this.h, widget.cPointer())
}

func (this *QDockWidget) SetFeatures(features QDockWidget__DockWidgetFeature) {
	C.QDockWidget_SetFeatures(this.h, (C.int)(features))
}

func (this *QDockWidget) Features() QDockWidget__DockWidgetFeature {
	return (QDockWidget__DockWidgetFeature)(C.QDockWidget_Features(this.h))
}

func (this *QDockWidget) SetFloating(floating bool) {
	C.QDockWidget_SetFloating(this.h, (C.bool)(floating))
}

func (this *QDockWidget) IsFloating() bool {
	return (bool)(C.QDockWidget_IsFloating(this.h))
}

func (this *QDockWidget) SetAllowedAreas(areas DockWidgetArea) {
	C.QDockWidget_SetAllowedAreas(this.h, (C.int)(areas))
}

func (this *QDockWidget) AllowedAreas() DockWidgetArea {
	return (DockWidgetArea)(C.QDockWidget_AllowedAreas(this.h))
}

func (this *QDockWidget) SetTitleBarWidget(widget *QWidget) {
	C.QDockWidget_SetTitleBarWidget(this.h, widget.cPointer())
}

func (this *QDockWidget) TitleBarWidget() *QWidget {
	return newQWidget_U(unsafe.Pointer(C.QDockWidget_TitleBarWidget(this.h)))
}

func (this *QDockWidget) IsAreaAllowed(area DockWidgetArea) bool {
	return (bool)(C.QDockWidget_IsAreaAllowed(this.h, (C.int)(area)))
}

func (this *QDockWidget) ToggleViewAction() *QAction {
	return newQAction_U(unsafe.Pointer(C.QDockWidget_ToggleViewAction(this.h)))
}

func (this *QDockWidget) FeaturesChanged(features QDockWidget__DockWidgetFeature) {
	C.QDockWidget_FeaturesChanged(this.h, (C.int)(features))
}
func (this *QDockWidget) OnFeaturesChanged(slot func(features QDockWidget__DockWidgetFeature)) {
	C.QDockWidget_connect_FeaturesChanged(this.h, unsafe.Pointer(uintptr(cgo.NewHandle(slot))))
}

//export miqt_exec_callback_QDockWidget_FeaturesChanged
func miqt_exec_callback_QDockWidget_FeaturesChanged(cb *C.void, features C.int) {
	gofunc, ok := (cgo.Handle(uintptr(unsafe.Pointer(cb))).Value()).(func(features QDockWidget__DockWidgetFeature))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QDockWidget__DockWidgetFeature)(features)

	gofunc(slotval1)
}

func (this *QDockWidget) TopLevelChanged(topLevel bool) {
	C.QDockWidget_TopLevelChanged(this.h, (C.bool)(topLevel))
}
func (this *QDockWidget) OnTopLevelChanged(slot func(topLevel bool)) {
	C.QDockWidget_connect_TopLevelChanged(this.h, unsafe.Pointer(uintptr(cgo.NewHandle(slot))))
}

//export miqt_exec_callback_QDockWidget_TopLevelChanged
func miqt_exec_callback_QDockWidget_TopLevelChanged(cb *C.void, topLevel C.bool) {
	gofunc, ok := (cgo.Handle(uintptr(unsafe.Pointer(cb))).Value()).(func(topLevel bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(topLevel)

	gofunc(slotval1)
}

func (this *QDockWidget) AllowedAreasChanged(allowedAreas DockWidgetArea) {
	C.QDockWidget_AllowedAreasChanged(this.h, (C.int)(allowedAreas))
}
func (this *QDockWidget) OnAllowedAreasChanged(slot func(allowedAreas DockWidgetArea)) {
	C.QDockWidget_connect_AllowedAreasChanged(this.h, unsafe.Pointer(uintptr(cgo.NewHandle(slot))))
}

//export miqt_exec_callback_QDockWidget_AllowedAreasChanged
func miqt_exec_callback_QDockWidget_AllowedAreasChanged(cb *C.void, allowedAreas C.int) {
	gofunc, ok := (cgo.Handle(uintptr(unsafe.Pointer(cb))).Value()).(func(allowedAreas DockWidgetArea))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DockWidgetArea)(allowedAreas)

	gofunc(slotval1)
}

func (this *QDockWidget) VisibilityChanged(visible bool) {
	C.QDockWidget_VisibilityChanged(this.h, (C.bool)(visible))
}
func (this *QDockWidget) OnVisibilityChanged(slot func(visible bool)) {
	C.QDockWidget_connect_VisibilityChanged(this.h, unsafe.Pointer(uintptr(cgo.NewHandle(slot))))
}

//export miqt_exec_callback_QDockWidget_VisibilityChanged
func miqt_exec_callback_QDockWidget_VisibilityChanged(cb *C.void, visible C.bool) {
	gofunc, ok := (cgo.Handle(uintptr(unsafe.Pointer(cb))).Value()).(func(visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc(slotval1)
}

func (this *QDockWidget) DockLocationChanged(area DockWidgetArea) {
	C.QDockWidget_DockLocationChanged(this.h, (C.int)(area))
}
func (this *QDockWidget) OnDockLocationChanged(slot func(area DockWidgetArea)) {
	C.QDockWidget_connect_DockLocationChanged(this.h, unsafe.Pointer(uintptr(cgo.NewHandle(slot))))
}

//export miqt_exec_callback_QDockWidget_DockLocationChanged
func miqt_exec_callback_QDockWidget_DockLocationChanged(cb *C.void, area C.int) {
	gofunc, ok := (cgo.Handle(uintptr(unsafe.Pointer(cb))).Value()).(func(area DockWidgetArea))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DockWidgetArea)(area)

	gofunc(slotval1)
}

func QDockWidget_Tr2(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QDockWidget_Tr2(s_Cstring, c_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QDockWidget_Tr3(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QDockWidget_Tr3(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QDockWidget_TrUtf82(s string, c string) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QDockWidget_TrUtf82(s_Cstring, c_Cstring)
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

func QDockWidget_TrUtf83(s string, c string, n int) string {
	s_Cstring := C.CString(s)
	defer C.free(unsafe.Pointer(s_Cstring))
	c_Cstring := C.CString(c)
	defer C.free(unsafe.Pointer(c_Cstring))
	var _ms *C.struct_miqt_string = C.QDockWidget_TrUtf83(s_Cstring, c_Cstring, (C.int)(n))
	_ret := C.GoStringN(&_ms.data, C.int(int64(_ms.len)))
	C.free(unsafe.Pointer(_ms))
	return _ret
}

// Delete this object from C++ memory.
func (this *QDockWidget) Delete() {
	C.QDockWidget_Delete(this.h)
}

// GoGC adds a Go Finalizer to this pointer, so that it will be deleted
// from C++ memory once it is unreachable from Go memory.
func (this *QDockWidget) GoGC() {
	runtime.SetFinalizer(this, func(this *QDockWidget) {
		this.Delete()
		runtime.KeepAlive(this.h)
	})
}
