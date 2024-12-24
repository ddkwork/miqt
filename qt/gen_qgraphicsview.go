package qt

import (
	"unsafe"
)

type QGraphicsView__ViewportAnchor int

const (
	QGraphicsView__NoAnchor         QGraphicsView__ViewportAnchor = 0
	QGraphicsView__AnchorViewCenter QGraphicsView__ViewportAnchor = 1
	QGraphicsView__AnchorUnderMouse QGraphicsView__ViewportAnchor = 2
)

type QGraphicsView__CacheModeFlag int

const (
	QGraphicsView__CacheNone       QGraphicsView__CacheModeFlag = 0
	QGraphicsView__CacheBackground QGraphicsView__CacheModeFlag = 1
)

type QGraphicsView__DragMode int

const (
	QGraphicsView__NoDrag         QGraphicsView__DragMode = 0
	QGraphicsView__ScrollHandDrag QGraphicsView__DragMode = 1
	QGraphicsView__RubberBandDrag QGraphicsView__DragMode = 2
)

type QGraphicsView__ViewportUpdateMode int

const (
	QGraphicsView__FullViewportUpdate         QGraphicsView__ViewportUpdateMode = 0
	QGraphicsView__MinimalViewportUpdate      QGraphicsView__ViewportUpdateMode = 1
	QGraphicsView__SmartViewportUpdate        QGraphicsView__ViewportUpdateMode = 2
	QGraphicsView__NoViewportUpdate           QGraphicsView__ViewportUpdateMode = 3
	QGraphicsView__BoundingRectViewportUpdate QGraphicsView__ViewportUpdateMode = 4
)

type QGraphicsView__OptimizationFlag int

const (
	QGraphicsView__DontSavePainterState      QGraphicsView__OptimizationFlag = 1
	QGraphicsView__DontAdjustForAntialiasing QGraphicsView__OptimizationFlag = 2
	QGraphicsView__IndirectPainting          QGraphicsView__OptimizationFlag = 4
)

type QGraphicsView struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsView constructs a new QGraphicsView object.
func NewQGraphicsView(parent *QWidget) *QGraphicsView {

	ret := newQGraphicsView(QGraphicsView_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQGraphicsView2 constructs a new QGraphicsView object.
func NewQGraphicsView2() *QGraphicsView {

	ret := newQGraphicsView(QGraphicsView_new2())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsView3 constructs a new QGraphicsView object.
func NewQGraphicsView3(scene *QGraphicsScene) *QGraphicsView {

	ret := newQGraphicsView(QGraphicsView_new3(scene.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQGraphicsView4 constructs a new QGraphicsView object.
func NewQGraphicsView4(scene *QGraphicsScene, parent *QWidget) *QGraphicsView {

	ret := newQGraphicsView(QGraphicsView_new4(scene.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsView) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsView_MetaObject(this.h))
}

func (this *QGraphicsView) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsView_Metacast(this.h, param1_Cstring))
}

func QGraphicsView_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsView_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsView) SizeHint() *QSize {
	_goptr := newQSize(QGraphicsView_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) RenderHints() RenderHint {
	return (RenderHint)(QGraphicsView_RenderHints(this.h))
}

func (this *QGraphicsView) SetRenderHint(hint QPainter__RenderHint) {
	QGraphicsView_SetRenderHint(this.h, (int)(hint))
}

func (this *QGraphicsView) SetRenderHints(hints RenderHint) {
	QGraphicsView_SetRenderHints(this.h, (int)(hints))
}

func (this *QGraphicsView) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QGraphicsView_Alignment(this.h))
}

func (this *QGraphicsView) SetAlignment(alignment AlignmentFlag) {
	QGraphicsView_SetAlignment(this.h, (int)(alignment))
}

func (this *QGraphicsView) TransformationAnchor() ViewportAnchor {
	xxxxxxxxx
}

func (this *QGraphicsView) SetTransformationAnchor(anchor ViewportAnchor) {
	QGraphicsView_SetTransformationAnchor(this.h, anchor)
}

func (this *QGraphicsView) ResizeAnchor() ViewportAnchor {
	xxxxxxxxx
}

func (this *QGraphicsView) SetResizeAnchor(anchor ViewportAnchor) {
	QGraphicsView_SetResizeAnchor(this.h, anchor)
}

func (this *QGraphicsView) ViewportUpdateMode() ViewportUpdateMode {
	xxxxxxxxx
}

func (this *QGraphicsView) SetViewportUpdateMode(mode ViewportUpdateMode) {
	QGraphicsView_SetViewportUpdateMode(this.h, mode)
}

func (this *QGraphicsView) OptimizationFlags() OptimizationFlags {
	xxxxxxxxx
}

func (this *QGraphicsView) SetOptimizationFlag(flag OptimizationFlag) {
	QGraphicsView_SetOptimizationFlag(this.h, flag)
}

func (this *QGraphicsView) SetOptimizationFlags(flags OptimizationFlags) {
	QGraphicsView_SetOptimizationFlags(this.h, flags)
}

func (this *QGraphicsView) DragMode() DragMode {
	xxxxxxxxx
}

func (this *QGraphicsView) SetDragMode(mode DragMode) {
	QGraphicsView_SetDragMode(this.h, mode)
}

func (this *QGraphicsView) RubberBandSelectionMode() ItemSelectionMode {
	return (ItemSelectionMode)(QGraphicsView_RubberBandSelectionMode(this.h))
}

func (this *QGraphicsView) SetRubberBandSelectionMode(mode ItemSelectionMode) {
	QGraphicsView_SetRubberBandSelectionMode(this.h, (int)(mode))
}

func (this *QGraphicsView) RubberBandRect() *QRect {
	_goptr := newQRect(QGraphicsView_RubberBandRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) CacheMode() CacheMode {
	xxxxxxxxx
}

func (this *QGraphicsView) SetCacheMode(mode CacheMode) {
	QGraphicsView_SetCacheMode(this.h, mode)
}

func (this *QGraphicsView) ResetCachedContent() {
	QGraphicsView_ResetCachedContent(this.h)
}

func (this *QGraphicsView) IsInteractive() bool {
	return (bool)(QGraphicsView_IsInteractive(this.h))
}

func (this *QGraphicsView) SetInteractive(allowed bool) {
	QGraphicsView_SetInteractive(this.h, (bool)(allowed))
}

func (this *QGraphicsView) Scene() *QGraphicsScene {
	return newQGraphicsScene(QGraphicsView_Scene(this.h))
}

func (this *QGraphicsView) SetScene(scene *QGraphicsScene) {
	QGraphicsView_SetScene(this.h, scene.cPointer())
}

func (this *QGraphicsView) SceneRect() *QRectF {
	_goptr := newQRectF(QGraphicsView_SceneRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) SetSceneRect(rect *QRectF) {
	QGraphicsView_SetSceneRect(this.h, rect.cPointer())
}

func (this *QGraphicsView) SetSceneRect2(x float64, y float64, w float64, h float64) {
	QGraphicsView_SetSceneRect2(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QGraphicsView) Transform() *QTransform {
	_goptr := newQTransform(QGraphicsView_Transform(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) ViewportTransform() *QTransform {
	_goptr := newQTransform(QGraphicsView_ViewportTransform(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) IsTransformed() bool {
	return (bool)(QGraphicsView_IsTransformed(this.h))
}

func (this *QGraphicsView) SetTransform(matrix *QTransform) {
	QGraphicsView_SetTransform(this.h, matrix.cPointer())
}

func (this *QGraphicsView) ResetTransform() {
	QGraphicsView_ResetTransform(this.h)
}

func (this *QGraphicsView) Rotate(angle float64) {
	QGraphicsView_Rotate(this.h, (double)(angle))
}

func (this *QGraphicsView) Scale(sx float64, sy float64) {
	QGraphicsView_Scale(this.h, (double)(sx), (double)(sy))
}

func (this *QGraphicsView) Shear(sh float64, sv float64) {
	QGraphicsView_Shear(this.h, (double)(sh), (double)(sv))
}

func (this *QGraphicsView) Translate(dx float64, dy float64) {
	QGraphicsView_Translate(this.h, (double)(dx), (double)(dy))
}

func (this *QGraphicsView) CenterOn(pos *QPointF) {
	QGraphicsView_CenterOn(this.h, pos.cPointer())
}

func (this *QGraphicsView) CenterOn2(x float64, y float64) {
	QGraphicsView_CenterOn2(this.h, (double)(x), (double)(y))
}

func (this *QGraphicsView) CenterOnWithItem(item *QGraphicsItem) {
	QGraphicsView_CenterOnWithItem(this.h, item.cPointer())
}

func (this *QGraphicsView) EnsureVisible(rect *QRectF) {
	QGraphicsView_EnsureVisible(this.h, rect.cPointer())
}

func (this *QGraphicsView) EnsureVisible2(x float64, y float64, w float64, h float64) {
	QGraphicsView_EnsureVisible2(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QGraphicsView) EnsureVisibleWithItem(item *QGraphicsItem) {
	QGraphicsView_EnsureVisibleWithItem(this.h, item.cPointer())
}

func (this *QGraphicsView) FitInView(rect *QRectF) {
	QGraphicsView_FitInView(this.h, rect.cPointer())
}

func (this *QGraphicsView) FitInView2(x float64, y float64, w float64, h float64) {
	QGraphicsView_FitInView2(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QGraphicsView) FitInViewWithItem(item *QGraphicsItem) {
	QGraphicsView_FitInViewWithItem(this.h, item.cPointer())
}

func (this *QGraphicsView) Render(painter *QPainter) {
	QGraphicsView_Render(this.h, painter.cPointer())
}

func (this *QGraphicsView) Items() []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsView_Items(this.h)
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsView) ItemsWithPos(pos *QPoint) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsView_ItemsWithPos(this.h, pos.cPointer())
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsView) Items2(x int, y int) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsView_Items2(this.h, (int)(x), (int)(y))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsView) ItemsWithRect(rect *QRect) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsView_ItemsWithRect(this.h, rect.cPointer())
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsView) Items3(x int, y int, w int, h int) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsView_Items3(this.h, (int)(x), (int)(y), (int)(w), (int)(h))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsView) ItemsWithPath(path *QPainterPath) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsView_ItemsWithPath(this.h, path.cPointer())
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsView) ItemAt(pos *QPoint) *QGraphicsItem {
	return newQGraphicsItem(QGraphicsView_ItemAt(this.h, pos.cPointer()))
}

func (this *QGraphicsView) ItemAt2(x int, y int) *QGraphicsItem {
	return newQGraphicsItem(QGraphicsView_ItemAt2(this.h, (int)(x), (int)(y)))
}

func (this *QGraphicsView) MapToScene(point *QPoint) *QPointF {
	_goptr := newQPointF(QGraphicsView_MapToScene(this.h, point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) MapToSceneWithPath(path *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QGraphicsView_MapToSceneWithPath(this.h, path.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) MapFromScene(point *QPointF) *QPoint {
	_goptr := newQPoint(QGraphicsView_MapFromScene(this.h, point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) MapFromSceneWithPath(path *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QGraphicsView_MapFromSceneWithPath(this.h, path.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) MapToScene2(x int, y int) *QPointF {
	_goptr := newQPointF(QGraphicsView_MapToScene2(this.h, (int)(x), (int)(y)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) MapFromScene2(x float64, y float64) *QPoint {
	_goptr := newQPoint(QGraphicsView_MapFromScene2(this.h, (double)(x), (double)(y)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) InputMethodQuery(query InputMethodQuery) *QVariant {
	_goptr := newQVariant(QGraphicsView_InputMethodQuery(this.h, (int)(query)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) BackgroundBrush() *QBrush {
	_goptr := newQBrush(QGraphicsView_BackgroundBrush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) SetBackgroundBrush(brush *QBrush) {
	QGraphicsView_SetBackgroundBrush(this.h, brush.cPointer())
}

func (this *QGraphicsView) ForegroundBrush() *QBrush {
	_goptr := newQBrush(QGraphicsView_ForegroundBrush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsView) SetForegroundBrush(brush *QBrush) {
	QGraphicsView_SetForegroundBrush(this.h, brush.cPointer())
}

func (this *QGraphicsView) UpdateScene(rects []QRectF) {
	rects_CArray := (*[0xffff]*QRectF)(malloc(size_t(8 * len(rects))))
	defer free(unsafe.Pointer(rects_CArray))
	for i := range rects {
		rects_CArray[i] = rects[i].cPointer()
	}
	rects_ma := struct_miqt_array{len: size_t(len(rects)), data: unsafe.Pointer(rects_CArray)}
	QGraphicsView_UpdateScene(this.h, rects_ma)
}

func (this *QGraphicsView) InvalidateScene() {
	QGraphicsView_InvalidateScene(this.h)
}

func (this *QGraphicsView) UpdateSceneRect(rect *QRectF) {
	QGraphicsView_UpdateSceneRect(this.h, rect.cPointer())
}

func (this *QGraphicsView) RubberBandChanged(viewportRect QRect, fromScenePoint QPointF, toScenePoint QPointF) {
	QGraphicsView_RubberBandChanged(this.h, viewportRect.cPointer(), fromScenePoint.cPointer(), toScenePoint.cPointer())
}
func (this *QGraphicsView) OnRubberBandChanged(slot func(viewportRect QRect, fromScenePoint QPointF, toScenePoint QPointF)) {
	QGraphicsView_connect_RubberBandChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_RubberBandChanged
func miqt_exec_callback_QGraphicsView_RubberBandChanged(cb intptr_t, viewportRect *QRect, fromScenePoint *QPointF, toScenePoint *QPointF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(viewportRect QRect, fromScenePoint QPointF, toScenePoint QPointF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	viewportRect_goptr := newQRect(viewportRect)
	viewportRect_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *viewportRect_goptr

	fromScenePoint_goptr := newQPointF(fromScenePoint)
	fromScenePoint_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval2 := *fromScenePoint_goptr

	toScenePoint_goptr := newQPointF(toScenePoint)
	toScenePoint_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval3 := *toScenePoint_goptr

	gofunc(slotval1, slotval2, slotval3)
}

func QGraphicsView_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsView_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsView_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsView_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsView) SetRenderHint2(hint QPainter__RenderHint, enabled bool) {
	QGraphicsView_SetRenderHint2(this.h, (int)(hint), (bool)(enabled))
}

func (this *QGraphicsView) SetOptimizationFlag2(flag OptimizationFlag, enabled bool) {
	QGraphicsView_SetOptimizationFlag2(this.h, flag, (bool)(enabled))
}

func (this *QGraphicsView) SetTransform2(matrix *QTransform, combine bool) {
	QGraphicsView_SetTransform2(this.h, matrix.cPointer(), (bool)(combine))
}

func (this *QGraphicsView) EnsureVisible22(rect *QRectF, xmargin int) {
	QGraphicsView_EnsureVisible22(this.h, rect.cPointer(), (int)(xmargin))
}

func (this *QGraphicsView) EnsureVisible3(rect *QRectF, xmargin int, ymargin int) {
	QGraphicsView_EnsureVisible3(this.h, rect.cPointer(), (int)(xmargin), (int)(ymargin))
}

func (this *QGraphicsView) EnsureVisible5(x float64, y float64, w float64, h float64, xmargin int) {
	QGraphicsView_EnsureVisible5(this.h, (double)(x), (double)(y), (double)(w), (double)(h), (int)(xmargin))
}

func (this *QGraphicsView) EnsureVisible6(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	QGraphicsView_EnsureVisible6(this.h, (double)(x), (double)(y), (double)(w), (double)(h), (int)(xmargin), (int)(ymargin))
}

func (this *QGraphicsView) EnsureVisible23(item *QGraphicsItem, xmargin int) {
	QGraphicsView_EnsureVisible23(this.h, item.cPointer(), (int)(xmargin))
}

func (this *QGraphicsView) EnsureVisible32(item *QGraphicsItem, xmargin int, ymargin int) {
	QGraphicsView_EnsureVisible32(this.h, item.cPointer(), (int)(xmargin), (int)(ymargin))
}

func (this *QGraphicsView) FitInView22(rect *QRectF, aspectRadioMode AspectRatioMode) {
	QGraphicsView_FitInView22(this.h, rect.cPointer(), (int)(aspectRadioMode))
}

func (this *QGraphicsView) FitInView5(x float64, y float64, w float64, h float64, aspectRadioMode AspectRatioMode) {
	QGraphicsView_FitInView5(this.h, (double)(x), (double)(y), (double)(w), (double)(h), (int)(aspectRadioMode))
}

func (this *QGraphicsView) FitInView23(item *QGraphicsItem, aspectRadioMode AspectRatioMode) {
	QGraphicsView_FitInView23(this.h, item.cPointer(), (int)(aspectRadioMode))
}

func (this *QGraphicsView) Render2(painter *QPainter, target *QRectF) {
	QGraphicsView_Render2(this.h, painter.cPointer(), target.cPointer())
}

func (this *QGraphicsView) Render3(painter *QPainter, target *QRectF, source *QRect) {
	QGraphicsView_Render3(this.h, painter.cPointer(), target.cPointer(), source.cPointer())
}

func (this *QGraphicsView) Render4(painter *QPainter, target *QRectF, source *QRect, aspectRatioMode AspectRatioMode) {
	QGraphicsView_Render4(this.h, painter.cPointer(), target.cPointer(), source.cPointer(), (int)(aspectRatioMode))
}

func (this *QGraphicsView) Items22(rect *QRect, mode ItemSelectionMode) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsView_Items22(this.h, rect.cPointer(), (int)(mode))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsView) Items5(x int, y int, w int, h int, mode ItemSelectionMode) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsView_Items5(this.h, (int)(x), (int)(y), (int)(w), (int)(h), (int)(mode))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsView) Items24(path *QPainterPath, mode ItemSelectionMode) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsView_Items24(this.h, path.cPointer(), (int)(mode))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsView) InvalidateScene1(rect *QRectF) {
	QGraphicsView_InvalidateScene1(this.h, rect.cPointer())
}

func (this *QGraphicsView) InvalidateScene2(rect *QRectF, layers SceneLayer) {
	QGraphicsView_InvalidateScene2(this.h, rect.cPointer(), (int)(layers))
}

func (this *QGraphicsView) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QGraphicsView_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsView) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_SizeHint
func miqt_exec_callback_QGraphicsView_SizeHint(self QGraphicsView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsView{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QGraphicsView) callVirtualBase_InputMethodQuery(query InputMethodQuery) *QVariant {

	_goptr := newQVariant(QGraphicsView_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(query)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsView) OnInputMethodQuery(slot func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_InputMethodQuery
func miqt_exec_callback_QGraphicsView_InputMethodQuery(self QGraphicsView, cb intptr_t, query int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(query)

	virtualReturn := gofunc((&QGraphicsView{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGraphicsView) callVirtualBase_SetupViewport(widget *QWidget) {

	QGraphicsView_virtualbase_SetupViewport(unsafe.Pointer(this.h), widget.cPointer())

}
func (this *QGraphicsView) OnSetupViewport(slot func(super func(widget *QWidget), widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_SetupViewport(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_SetupViewport
func miqt_exec_callback_QGraphicsView_SetupViewport(self QGraphicsView, cb intptr_t, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(widget *QWidget), widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(widget)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_SetupViewport, slotval1)

}

func (this *QGraphicsView) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QGraphicsView_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QGraphicsView) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_Event
func miqt_exec_callback_QGraphicsView_Event(self QGraphicsView, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QGraphicsView{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsView) callVirtualBase_ViewportEvent(event *QEvent) bool {

	return (bool)(QGraphicsView_virtualbase_ViewportEvent(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QGraphicsView) OnViewportEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_ViewportEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_ViewportEvent
func miqt_exec_callback_QGraphicsView_ViewportEvent(self QGraphicsView, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QGraphicsView{h: self}).callVirtualBase_ViewportEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsView) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QGraphicsView_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_ContextMenuEvent
func miqt_exec_callback_QGraphicsView_ContextMenuEvent(self QGraphicsView, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QGraphicsView_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_DragEnterEvent
func miqt_exec_callback_QGraphicsView_DragEnterEvent(self QGraphicsView, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QGraphicsView_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_DragLeaveEvent
func miqt_exec_callback_QGraphicsView_DragLeaveEvent(self QGraphicsView, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QGraphicsView_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_DragMoveEvent
func miqt_exec_callback_QGraphicsView_DragMoveEvent(self QGraphicsView, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_DropEvent(event *QDropEvent) {

	QGraphicsView_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_DropEvent
func miqt_exec_callback_QGraphicsView_DropEvent(self QGraphicsView, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QGraphicsView_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_FocusInEvent
func miqt_exec_callback_QGraphicsView_FocusInEvent(self QGraphicsView, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QGraphicsView_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QGraphicsView) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_FocusNextPrevChild
func miqt_exec_callback_QGraphicsView_FocusNextPrevChild(self QGraphicsView, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QGraphicsView{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsView) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QGraphicsView_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_FocusOutEvent
func miqt_exec_callback_QGraphicsView_FocusOutEvent(self QGraphicsView, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QGraphicsView_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_KeyPressEvent
func miqt_exec_callback_QGraphicsView_KeyPressEvent(self QGraphicsView, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QGraphicsView_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_KeyReleaseEvent
func miqt_exec_callback_QGraphicsView_KeyReleaseEvent(self QGraphicsView, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QGraphicsView_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_MouseDoubleClickEvent
func miqt_exec_callback_QGraphicsView_MouseDoubleClickEvent(self QGraphicsView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QGraphicsView_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_MousePressEvent
func miqt_exec_callback_QGraphicsView_MousePressEvent(self QGraphicsView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QGraphicsView_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_MouseMoveEvent
func miqt_exec_callback_QGraphicsView_MouseMoveEvent(self QGraphicsView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QGraphicsView_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_MouseReleaseEvent
func miqt_exec_callback_QGraphicsView_MouseReleaseEvent(self QGraphicsView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QGraphicsView_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_WheelEvent
func miqt_exec_callback_QGraphicsView_WheelEvent(self QGraphicsView, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QGraphicsView_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_PaintEvent
func miqt_exec_callback_QGraphicsView_PaintEvent(self QGraphicsView, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QGraphicsView_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_ResizeEvent
func miqt_exec_callback_QGraphicsView_ResizeEvent(self QGraphicsView, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QGraphicsView_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QGraphicsView) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_ScrollContentsBy
func miqt_exec_callback_QGraphicsView_ScrollContentsBy(self QGraphicsView, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QGraphicsView) callVirtualBase_ShowEvent(event *QShowEvent) {

	QGraphicsView_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_ShowEvent
func miqt_exec_callback_QGraphicsView_ShowEvent(self QGraphicsView, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_InputMethodEvent(event *QInputMethodEvent) {

	QGraphicsView_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsView) OnInputMethodEvent(slot func(super func(event *QInputMethodEvent), event *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_InputMethodEvent
func miqt_exec_callback_QGraphicsView_InputMethodEvent(self QGraphicsView, cb intptr_t, event *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QInputMethodEvent), event *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(event)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QGraphicsView) callVirtualBase_DrawBackground(painter *QPainter, rect *QRectF) {

	QGraphicsView_virtualbase_DrawBackground(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer())

}
func (this *QGraphicsView) OnDrawBackground(slot func(super func(painter *QPainter, rect *QRectF), painter *QPainter, rect *QRectF)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_DrawBackground(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_DrawBackground
func miqt_exec_callback_QGraphicsView_DrawBackground(self QGraphicsView, cb intptr_t, painter *QPainter, rect *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, rect *QRectF), painter *QPainter, rect *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQRectF(rect)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_DrawBackground, slotval1, slotval2)

}

func (this *QGraphicsView) callVirtualBase_DrawForeground(painter *QPainter, rect *QRectF) {

	QGraphicsView_virtualbase_DrawForeground(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer())

}
func (this *QGraphicsView) OnDrawForeground(slot func(super func(painter *QPainter, rect *QRectF), painter *QPainter, rect *QRectF)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_DrawForeground(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_DrawForeground
func miqt_exec_callback_QGraphicsView_DrawForeground(self QGraphicsView, cb intptr_t, painter *QPainter, rect *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, rect *QRectF), painter *QPainter, rect *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQRectF(rect)

	gofunc((&QGraphicsView{h: self}).callVirtualBase_DrawForeground, slotval1, slotval2)

}

func (this *QGraphicsView) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QGraphicsView_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsView) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_MinimumSizeHint
func miqt_exec_callback_QGraphicsView_MinimumSizeHint(self QGraphicsView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsView{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QGraphicsView) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(QGraphicsView_virtualbase_EventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QGraphicsView) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_EventFilter
func miqt_exec_callback_QGraphicsView_EventFilter(self QGraphicsView, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QGraphicsView{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QGraphicsView) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QGraphicsView_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsView) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsView_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsView_ViewportSizeHint
func miqt_exec_callback_QGraphicsView_ViewportSizeHint(self QGraphicsView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsView{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

}
