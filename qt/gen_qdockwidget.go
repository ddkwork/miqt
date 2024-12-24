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

func (this *QDockWidget) callVirtualBase_ChangeEvent(event *QEvent) {

	QDockWidget_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_ChangeEvent
func miqt_exec_callback_QDockWidget_ChangeEvent(self QDockWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QDockWidget_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_CloseEvent
func miqt_exec_callback_QDockWidget_CloseEvent(self QDockWidget, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QDockWidget_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_PaintEvent
func miqt_exec_callback_QDockWidget_PaintEvent(self QDockWidget, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QDockWidget_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QDockWidget) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_Event
func miqt_exec_callback_QDockWidget_Event(self QDockWidget, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QDockWidget) callVirtualBase_InitStyleOption(option *QStyleOptionDockWidget) {

	QDockWidget_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QDockWidget) OnInitStyleOption(slot func(super func(option *QStyleOptionDockWidget), option *QStyleOptionDockWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_InitStyleOption
func miqt_exec_callback_QDockWidget_InitStyleOption(self QDockWidget, cb intptr_t, option *QStyleOptionDockWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionDockWidget), option *QStyleOptionDockWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionDockWidget(option)

	gofunc((&QDockWidget{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QDockWidget) callVirtualBase_DevType() int {

	return (int)(QDockWidget_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QDockWidget) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_DevType
func miqt_exec_callback_QDockWidget_DevType(self QDockWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QDockWidget) callVirtualBase_SetVisible(visible bool) {

	QDockWidget_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QDockWidget) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_SetVisible
func miqt_exec_callback_QDockWidget_SetVisible(self QDockWidget, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QDockWidget{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QDockWidget) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QDockWidget_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDockWidget) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_SizeHint
func miqt_exec_callback_QDockWidget_SizeHint(self QDockWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QDockWidget) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QDockWidget_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDockWidget) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_MinimumSizeHint
func miqt_exec_callback_QDockWidget_MinimumSizeHint(self QDockWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QDockWidget) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QDockWidget_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QDockWidget) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_HeightForWidth
func miqt_exec_callback_QDockWidget_HeightForWidth(self QDockWidget, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QDockWidget) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QDockWidget_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QDockWidget) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_HasHeightForWidth
func miqt_exec_callback_QDockWidget_HasHeightForWidth(self QDockWidget, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QDockWidget) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QDockWidget_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QDockWidget) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_PaintEngine
func miqt_exec_callback_QDockWidget_PaintEngine(self QDockWidget, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QDockWidget) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QDockWidget_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_MousePressEvent
func miqt_exec_callback_QDockWidget_MousePressEvent(self QDockWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QDockWidget_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_MouseReleaseEvent
func miqt_exec_callback_QDockWidget_MouseReleaseEvent(self QDockWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QDockWidget_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_MouseDoubleClickEvent
func miqt_exec_callback_QDockWidget_MouseDoubleClickEvent(self QDockWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QDockWidget_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_MouseMoveEvent
func miqt_exec_callback_QDockWidget_MouseMoveEvent(self QDockWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QDockWidget_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_WheelEvent
func miqt_exec_callback_QDockWidget_WheelEvent(self QDockWidget, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QDockWidget_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_KeyPressEvent
func miqt_exec_callback_QDockWidget_KeyPressEvent(self QDockWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QDockWidget_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_KeyReleaseEvent
func miqt_exec_callback_QDockWidget_KeyReleaseEvent(self QDockWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QDockWidget_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_FocusInEvent
func miqt_exec_callback_QDockWidget_FocusInEvent(self QDockWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QDockWidget_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_FocusOutEvent
func miqt_exec_callback_QDockWidget_FocusOutEvent(self QDockWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QDockWidget_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_EnterEvent
func miqt_exec_callback_QDockWidget_EnterEvent(self QDockWidget, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_LeaveEvent(event *QEvent) {

	QDockWidget_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_LeaveEvent
func miqt_exec_callback_QDockWidget_LeaveEvent(self QDockWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QDockWidget_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_MoveEvent
func miqt_exec_callback_QDockWidget_MoveEvent(self QDockWidget, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QDockWidget_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_ResizeEvent
func miqt_exec_callback_QDockWidget_ResizeEvent(self QDockWidget, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QDockWidget_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_ContextMenuEvent
func miqt_exec_callback_QDockWidget_ContextMenuEvent(self QDockWidget, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QDockWidget_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_TabletEvent
func miqt_exec_callback_QDockWidget_TabletEvent(self QDockWidget, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_ActionEvent(event *QActionEvent) {

	QDockWidget_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_ActionEvent
func miqt_exec_callback_QDockWidget_ActionEvent(self QDockWidget, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QDockWidget_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_DragEnterEvent
func miqt_exec_callback_QDockWidget_DragEnterEvent(self QDockWidget, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QDockWidget_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_DragMoveEvent
func miqt_exec_callback_QDockWidget_DragMoveEvent(self QDockWidget, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QDockWidget_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_DragLeaveEvent
func miqt_exec_callback_QDockWidget_DragLeaveEvent(self QDockWidget, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_DropEvent(event *QDropEvent) {

	QDockWidget_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_DropEvent
func miqt_exec_callback_QDockWidget_DropEvent(self QDockWidget, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_ShowEvent(event *QShowEvent) {

	QDockWidget_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_ShowEvent
func miqt_exec_callback_QDockWidget_ShowEvent(self QDockWidget, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_HideEvent(event *QHideEvent) {

	QDockWidget_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDockWidget) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_HideEvent
func miqt_exec_callback_QDockWidget_HideEvent(self QDockWidget, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QDockWidget{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QDockWidget_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QDockWidget) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_NativeEvent
func miqt_exec_callback_QDockWidget_NativeEvent(self QDockWidget, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QDockWidget) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QDockWidget_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QDockWidget) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_Metric
func miqt_exec_callback_QDockWidget_Metric(self QDockWidget, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QDockWidget) callVirtualBase_InitPainter(painter *QPainter) {

	QDockWidget_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QDockWidget) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_InitPainter
func miqt_exec_callback_QDockWidget_InitPainter(self QDockWidget, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QDockWidget{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QDockWidget) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QDockWidget_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QDockWidget) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_Redirected
func miqt_exec_callback_QDockWidget_Redirected(self QDockWidget, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QDockWidget) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QDockWidget_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QDockWidget) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_SharedPainter
func miqt_exec_callback_QDockWidget_SharedPainter(self QDockWidget, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QDockWidget) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QDockWidget_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QDockWidget) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_InputMethodEvent
func miqt_exec_callback_QDockWidget_InputMethodEvent(self QDockWidget, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QDockWidget{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QDockWidget) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QDockWidget_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDockWidget) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_InputMethodQuery
func miqt_exec_callback_QDockWidget_InputMethodQuery(self QDockWidget, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QDockWidget) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QDockWidget_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QDockWidget) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDockWidget_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDockWidget_FocusNextPrevChild
func miqt_exec_callback_QDockWidget_FocusNextPrevChild(self QDockWidget, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QDockWidget{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
