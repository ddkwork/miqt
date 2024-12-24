package qt

import (
	"unsafe"
)

type QGraphicsScene__ItemIndexMethod int

const (
	QGraphicsScene__BspTreeIndex QGraphicsScene__ItemIndexMethod = 0
	QGraphicsScene__NoIndex      QGraphicsScene__ItemIndexMethod = -1
)

type QGraphicsScene__SceneLayer int

const (
	QGraphicsScene__ItemLayer       QGraphicsScene__SceneLayer = 1
	QGraphicsScene__BackgroundLayer QGraphicsScene__SceneLayer = 2
	QGraphicsScene__ForegroundLayer QGraphicsScene__SceneLayer = 4
	QGraphicsScene__AllLayers       QGraphicsScene__SceneLayer = 65535
)

type QGraphicsScene struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsScene constructs a new QGraphicsScene object.
func NewQGraphicsScene() *QGraphicsScene {
	g := newQGraphicsScene(QGraphicsScene_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsScene2 constructs a new QGraphicsScene object.
func NewQGraphicsScene2(sceneRect *QRectF) *QGraphicsScene {
	g := newQGraphicsScene(QGraphicsScene_new2(sceneRect.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsScene3 constructs a new QGraphicsScene object.
func NewQGraphicsScene3(x float64, y float64, width float64, height float64) *QGraphicsScene {
	g := newQGraphicsScene(QGraphicsScene_new3((double)(x), (double)(y), (double)(width), (double)(height)))
	g.isSubclass = true
	return g
}

// NewQGraphicsScene4 constructs a new QGraphicsScene object.
func NewQGraphicsScene4(parent *QObject) *QGraphicsScene {
	g := newQGraphicsScene(QGraphicsScene_new4(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsScene5 constructs a new QGraphicsScene object.
func NewQGraphicsScene5(sceneRect *QRectF, parent *QObject) *QGraphicsScene {
	g := newQGraphicsScene(QGraphicsScene_new5(sceneRect.cPointer(), parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsScene6 constructs a new QGraphicsScene object.
func NewQGraphicsScene6(x float64, y float64, width float64, height float64, parent *QObject) *QGraphicsScene {
	g := newQGraphicsScene(QGraphicsScene_new6((double)(x), (double)(y), (double)(width), (double)(height), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsScene) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsScene_MetaObject(this.h))
}

func (this *QGraphicsScene) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsScene_Metacast(this.h, param1_Cstring))
}

func QGraphicsScene_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsScene_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsScene) SceneRect() *QRectF {
	_goptr := newQRectF(QGraphicsScene_SceneRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsScene) Width() float64 {
	return (float64)(QGraphicsScene_Width(this.h))
}

func (this *QGraphicsScene) Height() float64 {
	return (float64)(QGraphicsScene_Height(this.h))
}

func (this *QGraphicsScene) SetSceneRect(rect *QRectF) {
	QGraphicsScene_SetSceneRect(this.h, rect.cPointer())
}

func (this *QGraphicsScene) SetSceneRect2(x float64, y float64, w float64, h float64) {
	QGraphicsScene_SetSceneRect2(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QGraphicsScene) Render(painter *QPainter) {
	QGraphicsScene_Render(this.h, painter.cPointer())
}

func (this *QGraphicsScene) ItemIndexMethod() ItemIndexMethod {
	xxxxxxxxx
}

func (this *QGraphicsScene) SetItemIndexMethod(method ItemIndexMethod) {
	QGraphicsScene_SetItemIndexMethod(this.h, method)
}

func (this *QGraphicsScene) BspTreeDepth() int {
	return (int)(QGraphicsScene_BspTreeDepth(this.h))
}

func (this *QGraphicsScene) SetBspTreeDepth(depth int) {
	QGraphicsScene_SetBspTreeDepth(this.h, (int)(depth))
}

func (this *QGraphicsScene) ItemsBoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsScene_ItemsBoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsScene) Items() []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items(this.h)
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) ItemsWithPos(pos *QPointF) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_ItemsWithPos(this.h, pos.cPointer())
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) ItemsWithRect(rect *QRectF) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_ItemsWithRect(this.h, rect.cPointer())
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) ItemsWithPath(path *QPainterPath) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_ItemsWithPath(this.h, path.cPointer())
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) Items2(x float64, y float64, w float64, h float64, mode ItemSelectionMode, order SortOrder) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items2(this.h, (double)(x), (double)(y), (double)(w), (double)(h), (int)(mode), (int)(order))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) CollidingItems(item *QGraphicsItem) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_CollidingItems(this.h, item.cPointer())
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) ItemAt(pos *QPointF, deviceTransform *QTransform) *QGraphicsItem {
	return newQGraphicsItem(QGraphicsScene_ItemAt(this.h, pos.cPointer(), deviceTransform.cPointer()))
}

func (this *QGraphicsScene) ItemAt2(x float64, y float64, deviceTransform *QTransform) *QGraphicsItem {
	return newQGraphicsItem(QGraphicsScene_ItemAt2(this.h, (double)(x), (double)(y), deviceTransform.cPointer()))
}

func (this *QGraphicsScene) SelectedItems() []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_SelectedItems(this.h)
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) SelectionArea() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsScene_SelectionArea(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsScene) SetSelectionArea(path *QPainterPath, deviceTransform *QTransform) {
	QGraphicsScene_SetSelectionArea(this.h, path.cPointer(), deviceTransform.cPointer())
}

func (this *QGraphicsScene) SetSelectionAreaWithPath(path *QPainterPath) {
	QGraphicsScene_SetSelectionAreaWithPath(this.h, path.cPointer())
}

func (this *QGraphicsScene) CreateItemGroup(items []*QGraphicsItem) *QGraphicsItemGroup {
	items_CArray := (*[0xffff]*QGraphicsItem)(malloc(size_t(8 * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_CArray[i] = items[i].cPointer()
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	return newQGraphicsItemGroup(QGraphicsScene_CreateItemGroup(this.h, items_ma))
}

func (this *QGraphicsScene) DestroyItemGroup(group *QGraphicsItemGroup) {
	QGraphicsScene_DestroyItemGroup(this.h, group.cPointer())
}

func (this *QGraphicsScene) AddItem(item *QGraphicsItem) {
	QGraphicsScene_AddItem(this.h, item.cPointer())
}

func (this *QGraphicsScene) AddEllipse(rect *QRectF) *QGraphicsEllipseItem {
	return newQGraphicsEllipseItem(QGraphicsScene_AddEllipse(this.h, rect.cPointer()))
}

func (this *QGraphicsScene) AddLine(line *QLineF) *QGraphicsLineItem {
	return newQGraphicsLineItem(QGraphicsScene_AddLine(this.h, line.cPointer()))
}

func (this *QGraphicsScene) AddPath(path *QPainterPath) *QGraphicsPathItem {
	return newQGraphicsPathItem(QGraphicsScene_AddPath(this.h, path.cPointer()))
}

func (this *QGraphicsScene) AddPixmap(pixmap *QPixmap) *QGraphicsPixmapItem {
	return newQGraphicsPixmapItem(QGraphicsScene_AddPixmap(this.h, pixmap.cPointer()))
}

func (this *QGraphicsScene) AddRect(rect *QRectF) *QGraphicsRectItem {
	return newQGraphicsRectItem(QGraphicsScene_AddRect(this.h, rect.cPointer()))
}

func (this *QGraphicsScene) AddText(text string) *QGraphicsTextItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQGraphicsTextItem(QGraphicsScene_AddText(this.h, text_ms))
}

func (this *QGraphicsScene) AddSimpleText(text string) *QGraphicsSimpleTextItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQGraphicsSimpleTextItem(QGraphicsScene_AddSimpleText(this.h, text_ms))
}

func (this *QGraphicsScene) AddWidget(widget *QWidget) *QGraphicsProxyWidget {
	return newQGraphicsProxyWidget(QGraphicsScene_AddWidget(this.h, widget.cPointer()))
}

func (this *QGraphicsScene) AddEllipse2(x float64, y float64, w float64, h float64) *QGraphicsEllipseItem {
	return newQGraphicsEllipseItem(QGraphicsScene_AddEllipse2(this.h, (double)(x), (double)(y), (double)(w), (double)(h)))
}

func (this *QGraphicsScene) AddLine2(x1 float64, y1 float64, x2 float64, y2 float64) *QGraphicsLineItem {
	return newQGraphicsLineItem(QGraphicsScene_AddLine2(this.h, (double)(x1), (double)(y1), (double)(x2), (double)(y2)))
}

func (this *QGraphicsScene) AddRect2(x float64, y float64, w float64, h float64) *QGraphicsRectItem {
	return newQGraphicsRectItem(QGraphicsScene_AddRect2(this.h, (double)(x), (double)(y), (double)(w), (double)(h)))
}

func (this *QGraphicsScene) RemoveItem(item *QGraphicsItem) {
	QGraphicsScene_RemoveItem(this.h, item.cPointer())
}

func (this *QGraphicsScene) FocusItem() *QGraphicsItem {
	return newQGraphicsItem(QGraphicsScene_FocusItem(this.h))
}

func (this *QGraphicsScene) SetFocusItem(item *QGraphicsItem) {
	QGraphicsScene_SetFocusItem(this.h, item.cPointer())
}

func (this *QGraphicsScene) HasFocus() bool {
	return (bool)(QGraphicsScene_HasFocus(this.h))
}

func (this *QGraphicsScene) SetFocus() {
	QGraphicsScene_SetFocus(this.h)
}

func (this *QGraphicsScene) ClearFocus() {
	QGraphicsScene_ClearFocus(this.h)
}

func (this *QGraphicsScene) SetStickyFocus(enabled bool) {
	QGraphicsScene_SetStickyFocus(this.h, (bool)(enabled))
}

func (this *QGraphicsScene) StickyFocus() bool {
	return (bool)(QGraphicsScene_StickyFocus(this.h))
}

func (this *QGraphicsScene) MouseGrabberItem() *QGraphicsItem {
	return newQGraphicsItem(QGraphicsScene_MouseGrabberItem(this.h))
}

func (this *QGraphicsScene) BackgroundBrush() *QBrush {
	_goptr := newQBrush(QGraphicsScene_BackgroundBrush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsScene) SetBackgroundBrush(brush *QBrush) {
	QGraphicsScene_SetBackgroundBrush(this.h, brush.cPointer())
}

func (this *QGraphicsScene) ForegroundBrush() *QBrush {
	_goptr := newQBrush(QGraphicsScene_ForegroundBrush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsScene) SetForegroundBrush(brush *QBrush) {
	QGraphicsScene_SetForegroundBrush(this.h, brush.cPointer())
}

func (this *QGraphicsScene) InputMethodQuery(query InputMethodQuery) *QVariant {
	_goptr := newQVariant(QGraphicsScene_InputMethodQuery(this.h, (int)(query)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsScene) Views() []*QGraphicsView {
	var _ma struct_miqt_array = QGraphicsScene_Views(this.h)
	_ret := make([]*QGraphicsView, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsView)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsView(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) Update(x float64, y float64, w float64, h float64) {
	QGraphicsScene_Update(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QGraphicsScene) Invalidate(x float64, y float64, w float64, h float64) {
	QGraphicsScene_Invalidate(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QGraphicsScene) Style() *QStyle {
	return newQStyle(QGraphicsScene_Style(this.h))
}

func (this *QGraphicsScene) SetStyle(style *QStyle) {
	QGraphicsScene_SetStyle(this.h, style.cPointer())
}

func (this *QGraphicsScene) Font() *QFont {
	_goptr := newQFont(QGraphicsScene_Font(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsScene) SetFont(font *QFont) {
	QGraphicsScene_SetFont(this.h, font.cPointer())
}

func (this *QGraphicsScene) Palette() *QPalette {
	_goptr := newQPalette(QGraphicsScene_Palette(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsScene) SetPalette(palette *QPalette) {
	QGraphicsScene_SetPalette(this.h, palette.cPointer())
}

func (this *QGraphicsScene) IsActive() bool {
	return (bool)(QGraphicsScene_IsActive(this.h))
}

func (this *QGraphicsScene) ActivePanel() *QGraphicsItem {
	return newQGraphicsItem(QGraphicsScene_ActivePanel(this.h))
}

func (this *QGraphicsScene) SetActivePanel(item *QGraphicsItem) {
	QGraphicsScene_SetActivePanel(this.h, item.cPointer())
}

func (this *QGraphicsScene) ActiveWindow() *QGraphicsWidget {
	return newQGraphicsWidget(QGraphicsScene_ActiveWindow(this.h))
}

func (this *QGraphicsScene) SetActiveWindow(widget *QGraphicsWidget) {
	QGraphicsScene_SetActiveWindow(this.h, widget.cPointer())
}

func (this *QGraphicsScene) SendEvent(item *QGraphicsItem, event *QEvent) bool {
	return (bool)(QGraphicsScene_SendEvent(this.h, item.cPointer(), event.cPointer()))
}

func (this *QGraphicsScene) MinimumRenderSize() float64 {
	return (float64)(QGraphicsScene_MinimumRenderSize(this.h))
}

func (this *QGraphicsScene) SetMinimumRenderSize(minSize float64) {
	QGraphicsScene_SetMinimumRenderSize(this.h, (double)(minSize))
}

func (this *QGraphicsScene) FocusOnTouch() bool {
	return (bool)(QGraphicsScene_FocusOnTouch(this.h))
}

func (this *QGraphicsScene) SetFocusOnTouch(enabled bool) {
	QGraphicsScene_SetFocusOnTouch(this.h, (bool)(enabled))
}

func (this *QGraphicsScene) Update2() {
	QGraphicsScene_Update2(this.h)
}

func (this *QGraphicsScene) Invalidate2() {
	QGraphicsScene_Invalidate2(this.h)
}

func (this *QGraphicsScene) Advance() {
	QGraphicsScene_Advance(this.h)
}

func (this *QGraphicsScene) ClearSelection() {
	QGraphicsScene_ClearSelection(this.h)
}

func (this *QGraphicsScene) Clear() {
	QGraphicsScene_Clear(this.h)
}

func (this *QGraphicsScene) Changed(region []QRectF) {
	region_CArray := (*[0xffff]*QRectF)(malloc(size_t(8 * len(region))))
	defer free(unsafe.Pointer(region_CArray))
	for i := range region {
		region_CArray[i] = region[i].cPointer()
	}
	region_ma := struct_miqt_array{len: size_t(len(region)), data: unsafe.Pointer(region_CArray)}
	QGraphicsScene_Changed(this.h, region_ma)
}

func (this *QGraphicsScene) OnChanged(slot func(region []QRectF)) {
	QGraphicsScene_connect_Changed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScene_Changed
func miqt_exec_callback_QGraphicsScene_Changed(cb intptr_t, region struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(region []QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var region_ma struct_miqt_array = region
	region_ret := make([]QRectF, int(region_ma.len))
	region_outCast := (*[0xffff]*QRectF)(unsafe.Pointer(region_ma.data)) // hey ya
	for i := 0; i < int(region_ma.len); i++ {
		region_lv_goptr := newQRectF(region_outCast[i])
		region_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		region_ret[i] = *region_lv_goptr
	}
	slotval1 := region_ret

	gofunc(slotval1)
}

func (this *QGraphicsScene) SceneRectChanged(rect *QRectF) {
	QGraphicsScene_SceneRectChanged(this.h, rect.cPointer())
}

func (this *QGraphicsScene) OnSceneRectChanged(slot func(rect *QRectF)) {
	QGraphicsScene_connect_SceneRectChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScene_SceneRectChanged
func miqt_exec_callback_QGraphicsScene_SceneRectChanged(cb intptr_t, rect *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(rect *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(rect)

	gofunc(slotval1)
}

func (this *QGraphicsScene) SelectionChanged() {
	QGraphicsScene_SelectionChanged(this.h)
}

func (this *QGraphicsScene) OnSelectionChanged(slot func()) {
	QGraphicsScene_connect_SelectionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScene_SelectionChanged
func miqt_exec_callback_QGraphicsScene_SelectionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsScene) FocusItemChanged(newFocus *QGraphicsItem, oldFocus *QGraphicsItem, reason FocusReason) {
	QGraphicsScene_FocusItemChanged(this.h, newFocus.cPointer(), oldFocus.cPointer(), (int)(reason))
}

func (this *QGraphicsScene) OnFocusItemChanged(slot func(newFocus *QGraphicsItem, oldFocus *QGraphicsItem, reason FocusReason)) {
	QGraphicsScene_connect_FocusItemChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScene_FocusItemChanged
func miqt_exec_callback_QGraphicsScene_FocusItemChanged(cb intptr_t, newFocus *QGraphicsItem, oldFocus *QGraphicsItem, reason int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newFocus *QGraphicsItem, oldFocus *QGraphicsItem, reason FocusReason))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsItem(newFocus)

	slotval2 := newQGraphicsItem(oldFocus)

	slotval3 := (FocusReason)(reason)

	gofunc(slotval1, slotval2, slotval3)
}

func QGraphicsScene_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsScene_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsScene_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsScene_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsScene) Render2(painter *QPainter, target *QRectF) {
	QGraphicsScene_Render2(this.h, painter.cPointer(), target.cPointer())
}

func (this *QGraphicsScene) Render3(painter *QPainter, target *QRectF, source *QRectF) {
	QGraphicsScene_Render3(this.h, painter.cPointer(), target.cPointer(), source.cPointer())
}

func (this *QGraphicsScene) Render4(painter *QPainter, target *QRectF, source *QRectF, aspectRatioMode AspectRatioMode) {
	QGraphicsScene_Render4(this.h, painter.cPointer(), target.cPointer(), source.cPointer(), (int)(aspectRatioMode))
}

func (this *QGraphicsScene) Items1(order SortOrder) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items1(this.h, (int)(order))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) Items22(pos *QPointF, mode ItemSelectionMode) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items22(this.h, pos.cPointer(), (int)(mode))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) Items3(pos *QPointF, mode ItemSelectionMode, order SortOrder) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items3(this.h, pos.cPointer(), (int)(mode), (int)(order))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) Items4(pos *QPointF, mode ItemSelectionMode, order SortOrder, deviceTransform *QTransform) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items4(this.h, pos.cPointer(), (int)(mode), (int)(order), deviceTransform.cPointer())
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) Items23(rect *QRectF, mode ItemSelectionMode) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items23(this.h, rect.cPointer(), (int)(mode))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) Items32(rect *QRectF, mode ItemSelectionMode, order SortOrder) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items32(this.h, rect.cPointer(), (int)(mode), (int)(order))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) Items42(rect *QRectF, mode ItemSelectionMode, order SortOrder, deviceTransform *QTransform) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items42(this.h, rect.cPointer(), (int)(mode), (int)(order), deviceTransform.cPointer())
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) Items25(path *QPainterPath, mode ItemSelectionMode) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items25(this.h, path.cPointer(), (int)(mode))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) Items34(path *QPainterPath, mode ItemSelectionMode, order SortOrder) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items34(this.h, path.cPointer(), (int)(mode), (int)(order))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) Items44(path *QPainterPath, mode ItemSelectionMode, order SortOrder, deviceTransform *QTransform) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items44(this.h, path.cPointer(), (int)(mode), (int)(order), deviceTransform.cPointer())
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) Items7(x float64, y float64, w float64, h float64, mode ItemSelectionMode, order SortOrder, deviceTransform *QTransform) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_Items7(this.h, (double)(x), (double)(y), (double)(w), (double)(h), (int)(mode), (int)(order), deviceTransform.cPointer())
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) CollidingItems2(item *QGraphicsItem, mode ItemSelectionMode) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsScene_CollidingItems2(this.h, item.cPointer(), (int)(mode))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsScene) SetSelectionArea2(path *QPainterPath, selectionOperation ItemSelectionOperation) {
	QGraphicsScene_SetSelectionArea2(this.h, path.cPointer(), (int)(selectionOperation))
}

func (this *QGraphicsScene) SetSelectionArea3(path *QPainterPath, selectionOperation ItemSelectionOperation, mode ItemSelectionMode) {
	QGraphicsScene_SetSelectionArea3(this.h, path.cPointer(), (int)(selectionOperation), (int)(mode))
}

func (this *QGraphicsScene) SetSelectionArea4(path *QPainterPath, selectionOperation ItemSelectionOperation, mode ItemSelectionMode, deviceTransform *QTransform) {
	QGraphicsScene_SetSelectionArea4(this.h, path.cPointer(), (int)(selectionOperation), (int)(mode), deviceTransform.cPointer())
}

func (this *QGraphicsScene) AddEllipse22(rect *QRectF, pen *QPen) *QGraphicsEllipseItem {
	return newQGraphicsEllipseItem(QGraphicsScene_AddEllipse22(this.h, rect.cPointer(), pen.cPointer()))
}

func (this *QGraphicsScene) AddEllipse3(rect *QRectF, pen *QPen, brush *QBrush) *QGraphicsEllipseItem {
	return newQGraphicsEllipseItem(QGraphicsScene_AddEllipse3(this.h, rect.cPointer(), pen.cPointer(), brush.cPointer()))
}

func (this *QGraphicsScene) AddLine22(line *QLineF, pen *QPen) *QGraphicsLineItem {
	return newQGraphicsLineItem(QGraphicsScene_AddLine22(this.h, line.cPointer(), pen.cPointer()))
}

func (this *QGraphicsScene) AddPath2(path *QPainterPath, pen *QPen) *QGraphicsPathItem {
	return newQGraphicsPathItem(QGraphicsScene_AddPath2(this.h, path.cPointer(), pen.cPointer()))
}

func (this *QGraphicsScene) AddPath3(path *QPainterPath, pen *QPen, brush *QBrush) *QGraphicsPathItem {
	return newQGraphicsPathItem(QGraphicsScene_AddPath3(this.h, path.cPointer(), pen.cPointer(), brush.cPointer()))
}

func (this *QGraphicsScene) AddRect22(rect *QRectF, pen *QPen) *QGraphicsRectItem {
	return newQGraphicsRectItem(QGraphicsScene_AddRect22(this.h, rect.cPointer(), pen.cPointer()))
}

func (this *QGraphicsScene) AddRect3(rect *QRectF, pen *QPen, brush *QBrush) *QGraphicsRectItem {
	return newQGraphicsRectItem(QGraphicsScene_AddRect3(this.h, rect.cPointer(), pen.cPointer(), brush.cPointer()))
}

func (this *QGraphicsScene) AddText2(text string, font *QFont) *QGraphicsTextItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQGraphicsTextItem(QGraphicsScene_AddText2(this.h, text_ms, font.cPointer()))
}

func (this *QGraphicsScene) AddSimpleText2(text string, font *QFont) *QGraphicsSimpleTextItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQGraphicsSimpleTextItem(QGraphicsScene_AddSimpleText2(this.h, text_ms, font.cPointer()))
}

func (this *QGraphicsScene) AddWidget2(widget *QWidget, wFlags WindowType) *QGraphicsProxyWidget {
	return newQGraphicsProxyWidget(QGraphicsScene_AddWidget2(this.h, widget.cPointer(), (int)(wFlags)))
}

func (this *QGraphicsScene) AddEllipse5(x float64, y float64, w float64, h float64, pen *QPen) *QGraphicsEllipseItem {
	return newQGraphicsEllipseItem(QGraphicsScene_AddEllipse5(this.h, (double)(x), (double)(y), (double)(w), (double)(h), pen.cPointer()))
}

func (this *QGraphicsScene) AddEllipse6(x float64, y float64, w float64, h float64, pen *QPen, brush *QBrush) *QGraphicsEllipseItem {
	return newQGraphicsEllipseItem(QGraphicsScene_AddEllipse6(this.h, (double)(x), (double)(y), (double)(w), (double)(h), pen.cPointer(), brush.cPointer()))
}

func (this *QGraphicsScene) AddLine5(x1 float64, y1 float64, x2 float64, y2 float64, pen *QPen) *QGraphicsLineItem {
	return newQGraphicsLineItem(QGraphicsScene_AddLine5(this.h, (double)(x1), (double)(y1), (double)(x2), (double)(y2), pen.cPointer()))
}

func (this *QGraphicsScene) AddRect5(x float64, y float64, w float64, h float64, pen *QPen) *QGraphicsRectItem {
	return newQGraphicsRectItem(QGraphicsScene_AddRect5(this.h, (double)(x), (double)(y), (double)(w), (double)(h), pen.cPointer()))
}

func (this *QGraphicsScene) AddRect6(x float64, y float64, w float64, h float64, pen *QPen, brush *QBrush) *QGraphicsRectItem {
	return newQGraphicsRectItem(QGraphicsScene_AddRect6(this.h, (double)(x), (double)(y), (double)(w), (double)(h), pen.cPointer(), brush.cPointer()))
}

func (this *QGraphicsScene) SetFocusItem2(item *QGraphicsItem, focusReason FocusReason) {
	QGraphicsScene_SetFocusItem2(this.h, item.cPointer(), (int)(focusReason))
}

func (this *QGraphicsScene) SetFocus1(focusReason FocusReason) {
	QGraphicsScene_SetFocus1(this.h, (int)(focusReason))
}

func (this *QGraphicsScene) Invalidate5(x float64, y float64, w float64, h float64, layers SceneLayers) {
	QGraphicsScene_Invalidate5(this.h, (double)(x), (double)(y), (double)(w), (double)(h), layers)
}

func (this *QGraphicsScene) Update1(rect *QRectF) {
	QGraphicsScene_Update1(this.h, rect.cPointer())
}

func (this *QGraphicsScene) Invalidate1(rect *QRectF) {
	QGraphicsScene_Invalidate1(this.h, rect.cPointer())
}

func (this *QGraphicsScene) Invalidate22(rect *QRectF, layers SceneLayers) {
	QGraphicsScene_Invalidate22(this.h, rect.cPointer(), layers)
}

func (this *QGraphicsScene) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsScene_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QGraphicsScene) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsScene_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScene_MetaObject
func miqt_exec_callback_QGraphicsScene_MetaObject(self QGraphicsScene, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsScene{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QGraphicsScene) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGraphicsScene_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGraphicsScene) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsScene_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsScene_Metacast
func miqt_exec_callback_QGraphicsScene_Metacast(self QGraphicsScene, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QGraphicsScene{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
