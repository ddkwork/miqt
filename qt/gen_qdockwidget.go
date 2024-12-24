package qt

import (
	"unsafe"
)

type QDockWidget__DockWidgetFeature int

const (
	QDockWidget__DockWidgetClosable         QDockWidget__DockWidgetFeature = 1
	QDockWidget__DockWidgetMovable          QDockWidget__DockWidgetFeature = 2
	QDockWidget__DockWidgetFloatable        QDockWidget__DockWidgetFeature = 4
	QDockWidget__DockWidgetVerticalTitleBar QDockWidget__DockWidgetFeature = 8
	QDockWidget__DockWidgetFeatureMask      QDockWidget__DockWidgetFeature = 15
	QDockWidget__NoDockWidgetFeatures       QDockWidget__DockWidgetFeature = 0
	QDockWidget__Reserved                   QDockWidget__DockWidgetFeature = 255
)

type QDockWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQDockWidget constructs a new QDockWidget object.
func NewQDockWidget(parent *QWidget) *QDockWidget {
	ret := newQDockWidget(QDockWidget_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDockWidget2 constructs a new QDockWidget object.
func NewQDockWidget2(title string) *QDockWidget {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))

	ret := newQDockWidget(QDockWidget_new2(title_ms))
	ret.isSubclass = true
	return ret
}

// NewQDockWidget3 constructs a new QDockWidget object.
func NewQDockWidget3() *QDockWidget {
	ret := newQDockWidget(QDockWidget_new3())
	ret.isSubclass = true
	return ret
}

// NewQDockWidget4 constructs a new QDockWidget object.
func NewQDockWidget4(title string, parent *QWidget) *QDockWidget {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))

	ret := newQDockWidget(QDockWidget_new4(title_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDockWidget5 constructs a new QDockWidget object.
func NewQDockWidget5(title string, parent *QWidget, flags WindowType) *QDockWidget {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))

	ret := newQDockWidget(QDockWidget_new5(title_ms, parent.cPointer(), (int)(flags)))
	ret.isSubclass = true
	return ret
}

// NewQDockWidget6 constructs a new QDockWidget object.
func NewQDockWidget6(parent *QWidget, flags WindowType) *QDockWidget {
	ret := newQDockWidget(QDockWidget_new6(parent.cPointer(), (int)(flags)))
	ret.isSubclass = true
	return ret
}

func (this *QDockWidget) MetaObject() *QMetaObject {
	return newQMetaObject(QDockWidget_MetaObject(this.h))
}

func (this *QDockWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDockWidget_Metacast(this.h, param1_Cstring))
}

func QDockWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDockWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDockWidget) Widget() *QWidget {
	return newQWidget(QDockWidget_Widget(this.h))
}

func (this *QDockWidget) SetWidget(widget *QWidget) {
	QDockWidget_SetWidget(this.h, widget.cPointer())
}

func (this *QDockWidget) SetFeatures(features DockWidgetFeatures) {
	QDockWidget_SetFeatures(this.h, features)
}

func (this *QDockWidget) Features() DockWidgetFeatures {
	xxxxxxxxx
}

func (this *QDockWidget) SetFloating(floating bool) {
	QDockWidget_SetFloating(this.h, (bool)(floating))
}

func (this *QDockWidget) IsFloating() bool {
	return (bool)(QDockWidget_IsFloating(this.h))
}

func (this *QDockWidget) SetAllowedAreas(areas DockWidgetArea) {
	QDockWidget_SetAllowedAreas(this.h, (int)(areas))
}

func (this *QDockWidget) AllowedAreas() DockWidgetArea {
	return (DockWidgetArea)(QDockWidget_AllowedAreas(this.h))
}

func (this *QDockWidget) SetTitleBarWidget(widget *QWidget) {
	QDockWidget_SetTitleBarWidget(this.h, widget.cPointer())
}

func (this *QDockWidget) TitleBarWidget() *QWidget {
	return newQWidget(QDockWidget_TitleBarWidget(this.h))
}

func (this *QDockWidget) SetDockLocation(area DockWidgetArea) {
	QDockWidget_SetDockLocation(this.h, (int)(area))
}

func (this *QDockWidget) DockLocation() DockWidgetArea {
	return (DockWidgetArea)(QDockWidget_DockLocation(this.h))
}

func (this *QDockWidget) IsAreaAllowed(area DockWidgetArea) bool {
	return (bool)(QDockWidget_IsAreaAllowed(this.h, (int)(area)))
}

func (this *QDockWidget) ToggleViewAction() *QAction {
	return newQAction(QDockWidget_ToggleViewAction(this.h))
}

func (this *QDockWidget) FeaturesChanged(features DockWidgetFeature) {
	QDockWidget_FeaturesChanged(this.h, (int)(features))
}

func (this *QDockWidget) OnFeaturesChanged(slot func(features DockWidgetFeature)) {
	QDockWidget_connect_FeaturesChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_FeaturesChanged
func miqt_exec_callback_QDockWidget_FeaturesChanged(cb intptr_t, features int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(features DockWidgetFeature))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DockWidgetFeature)(features)

	gofunc(slotval1)
}

func (this *QDockWidget) TopLevelChanged(topLevel bool) {
	QDockWidget_TopLevelChanged(this.h, (bool)(topLevel))
}

func (this *QDockWidget) OnTopLevelChanged(slot func(topLevel bool)) {
	QDockWidget_connect_TopLevelChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_TopLevelChanged
func miqt_exec_callback_QDockWidget_TopLevelChanged(cb intptr_t, topLevel bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(topLevel bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(topLevel)

	gofunc(slotval1)
}

func (this *QDockWidget) AllowedAreasChanged(allowedAreas DockWidgetArea) {
	QDockWidget_AllowedAreasChanged(this.h, (int)(allowedAreas))
}

func (this *QDockWidget) OnAllowedAreasChanged(slot func(allowedAreas DockWidgetArea)) {
	QDockWidget_connect_AllowedAreasChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_AllowedAreasChanged
func miqt_exec_callback_QDockWidget_AllowedAreasChanged(cb intptr_t, allowedAreas int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(allowedAreas DockWidgetArea))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DockWidgetArea)(allowedAreas)

	gofunc(slotval1)
}

func (this *QDockWidget) VisibilityChanged(visible bool) {
	QDockWidget_VisibilityChanged(this.h, (bool)(visible))
}

func (this *QDockWidget) OnVisibilityChanged(slot func(visible bool)) {
	QDockWidget_connect_VisibilityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_VisibilityChanged
func miqt_exec_callback_QDockWidget_VisibilityChanged(cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc(slotval1)
}

func (this *QDockWidget) DockLocationChanged(area DockWidgetArea) {
	QDockWidget_DockLocationChanged(this.h, (int)(area))
}

func (this *QDockWidget) OnDockLocationChanged(slot func(area DockWidgetArea)) {
	QDockWidget_connect_DockLocationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_DockLocationChanged
func miqt_exec_callback_QDockWidget_DockLocationChanged(cb intptr_t, area int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(area DockWidgetArea))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DockWidgetArea)(area)

	gofunc(slotval1)
}

func QDockWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDockWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDockWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDockWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDockWidget) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QDockWidget_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QDockWidget) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_MetaObject
func miqt_exec_callback_QDockWidget_MetaObject(self QDockWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QDockWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QDockWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QDockWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_Metacast
func miqt_exec_callback_QDockWidget_Metacast(self QDockWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
