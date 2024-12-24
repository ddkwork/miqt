package qt

import (
	"unsafe"
)

type QGraphicsItem__GraphicsItemFlag int

const (
	QGraphicsItem__ItemIsMovable                        QGraphicsItem__GraphicsItemFlag = 1
	QGraphicsItem__ItemIsSelectable                     QGraphicsItem__GraphicsItemFlag = 2
	QGraphicsItem__ItemIsFocusable                      QGraphicsItem__GraphicsItemFlag = 4
	QGraphicsItem__ItemClipsToShape                     QGraphicsItem__GraphicsItemFlag = 8
	QGraphicsItem__ItemClipsChildrenToShape             QGraphicsItem__GraphicsItemFlag = 16
	QGraphicsItem__ItemIgnoresTransformations           QGraphicsItem__GraphicsItemFlag = 32
	QGraphicsItem__ItemIgnoresParentOpacity             QGraphicsItem__GraphicsItemFlag = 64
	QGraphicsItem__ItemDoesntPropagateOpacityToChildren QGraphicsItem__GraphicsItemFlag = 128
	QGraphicsItem__ItemStacksBehindParent               QGraphicsItem__GraphicsItemFlag = 256
	QGraphicsItem__ItemUsesExtendedStyleOption          QGraphicsItem__GraphicsItemFlag = 512
	QGraphicsItem__ItemHasNoContents                    QGraphicsItem__GraphicsItemFlag = 1024
	QGraphicsItem__ItemSendsGeometryChanges             QGraphicsItem__GraphicsItemFlag = 2048
	QGraphicsItem__ItemAcceptsInputMethod               QGraphicsItem__GraphicsItemFlag = 4096
	QGraphicsItem__ItemNegativeZStacksBehindParent      QGraphicsItem__GraphicsItemFlag = 8192
	QGraphicsItem__ItemIsPanel                          QGraphicsItem__GraphicsItemFlag = 16384
	QGraphicsItem__ItemIsFocusScope                     QGraphicsItem__GraphicsItemFlag = 32768
	QGraphicsItem__ItemSendsScenePositionChanges        QGraphicsItem__GraphicsItemFlag = 65536
	QGraphicsItem__ItemStopsClickFocusPropagation       QGraphicsItem__GraphicsItemFlag = 131072
	QGraphicsItem__ItemStopsFocusHandling               QGraphicsItem__GraphicsItemFlag = 262144
	QGraphicsItem__ItemContainsChildrenInShape          QGraphicsItem__GraphicsItemFlag = 524288
)

type QGraphicsItem__GraphicsItemChange int

const (
	QGraphicsItem__ItemPositionChange                 QGraphicsItem__GraphicsItemChange = 0
	QGraphicsItem__ItemVisibleChange                  QGraphicsItem__GraphicsItemChange = 2
	QGraphicsItem__ItemEnabledChange                  QGraphicsItem__GraphicsItemChange = 3
	QGraphicsItem__ItemSelectedChange                 QGraphicsItem__GraphicsItemChange = 4
	QGraphicsItem__ItemParentChange                   QGraphicsItem__GraphicsItemChange = 5
	QGraphicsItem__ItemChildAddedChange               QGraphicsItem__GraphicsItemChange = 6
	QGraphicsItem__ItemChildRemovedChange             QGraphicsItem__GraphicsItemChange = 7
	QGraphicsItem__ItemTransformChange                QGraphicsItem__GraphicsItemChange = 8
	QGraphicsItem__ItemPositionHasChanged             QGraphicsItem__GraphicsItemChange = 9
	QGraphicsItem__ItemTransformHasChanged            QGraphicsItem__GraphicsItemChange = 10
	QGraphicsItem__ItemSceneChange                    QGraphicsItem__GraphicsItemChange = 11
	QGraphicsItem__ItemVisibleHasChanged              QGraphicsItem__GraphicsItemChange = 12
	QGraphicsItem__ItemEnabledHasChanged              QGraphicsItem__GraphicsItemChange = 13
	QGraphicsItem__ItemSelectedHasChanged             QGraphicsItem__GraphicsItemChange = 14
	QGraphicsItem__ItemParentHasChanged               QGraphicsItem__GraphicsItemChange = 15
	QGraphicsItem__ItemSceneHasChanged                QGraphicsItem__GraphicsItemChange = 16
	QGraphicsItem__ItemCursorChange                   QGraphicsItem__GraphicsItemChange = 17
	QGraphicsItem__ItemCursorHasChanged               QGraphicsItem__GraphicsItemChange = 18
	QGraphicsItem__ItemToolTipChange                  QGraphicsItem__GraphicsItemChange = 19
	QGraphicsItem__ItemToolTipHasChanged              QGraphicsItem__GraphicsItemChange = 20
	QGraphicsItem__ItemFlagsChange                    QGraphicsItem__GraphicsItemChange = 21
	QGraphicsItem__ItemFlagsHaveChanged               QGraphicsItem__GraphicsItemChange = 22
	QGraphicsItem__ItemZValueChange                   QGraphicsItem__GraphicsItemChange = 23
	QGraphicsItem__ItemZValueHasChanged               QGraphicsItem__GraphicsItemChange = 24
	QGraphicsItem__ItemOpacityChange                  QGraphicsItem__GraphicsItemChange = 25
	QGraphicsItem__ItemOpacityHasChanged              QGraphicsItem__GraphicsItemChange = 26
	QGraphicsItem__ItemScenePositionHasChanged        QGraphicsItem__GraphicsItemChange = 27
	QGraphicsItem__ItemRotationChange                 QGraphicsItem__GraphicsItemChange = 28
	QGraphicsItem__ItemRotationHasChanged             QGraphicsItem__GraphicsItemChange = 29
	QGraphicsItem__ItemScaleChange                    QGraphicsItem__GraphicsItemChange = 30
	QGraphicsItem__ItemScaleHasChanged                QGraphicsItem__GraphicsItemChange = 31
	QGraphicsItem__ItemTransformOriginPointChange     QGraphicsItem__GraphicsItemChange = 32
	QGraphicsItem__ItemTransformOriginPointHasChanged QGraphicsItem__GraphicsItemChange = 33
)

type QGraphicsItem__CacheMode int

const (
	QGraphicsItem__NoCache               QGraphicsItem__CacheMode = 0
	QGraphicsItem__ItemCoordinateCache   QGraphicsItem__CacheMode = 1
	QGraphicsItem__DeviceCoordinateCache QGraphicsItem__CacheMode = 2
)

type QGraphicsItem__PanelModality int

const (
	QGraphicsItem__NonModal   QGraphicsItem__PanelModality = 0
	QGraphicsItem__PanelModal QGraphicsItem__PanelModality = 1
	QGraphicsItem__SceneModal QGraphicsItem__PanelModality = 2
)

type QGraphicsItem__ int

const (
	QGraphicsItem__Type     QGraphicsItem__ = 1
	QGraphicsItem__UserType QGraphicsItem__ = 65536
)

type QGraphicsItem__Extension int

const (
	QGraphicsItem__UserExtension QGraphicsItem__Extension = 2147483648
)

type QGraphicsPathItem__ int

const (
	QGraphicsPathItem__Type QGraphicsPathItem__ = 2
)

type QGraphicsRectItem__ int

const (
	QGraphicsRectItem__Type QGraphicsRectItem__ = 3
)

type QGraphicsEllipseItem__ int

const (
	QGraphicsEllipseItem__Type QGraphicsEllipseItem__ = 4
)

type QGraphicsPolygonItem__ int

const (
	QGraphicsPolygonItem__Type QGraphicsPolygonItem__ = 5
)

type QGraphicsLineItem__ int

const (
	QGraphicsLineItem__Type QGraphicsLineItem__ = 6
)

type QGraphicsPixmapItem__ShapeMode int

const (
	QGraphicsPixmapItem__MaskShape          QGraphicsPixmapItem__ShapeMode = 0
	QGraphicsPixmapItem__BoundingRectShape  QGraphicsPixmapItem__ShapeMode = 1
	QGraphicsPixmapItem__HeuristicMaskShape QGraphicsPixmapItem__ShapeMode = 2
)

type QGraphicsPixmapItem__ int

const (
	QGraphicsPixmapItem__Type QGraphicsPixmapItem__ = 7
)

type QGraphicsTextItem__ int

const (
	QGraphicsTextItem__Type QGraphicsTextItem__ = 8
)

type QGraphicsSimpleTextItem__ int

const (
	QGraphicsSimpleTextItem__Type QGraphicsSimpleTextItem__ = 9
)

type QGraphicsItemGroup__ int

const (
	QGraphicsItemGroup__Type QGraphicsItemGroup__ = 10
)

type QGraphicsItem struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsItem constructs a new QGraphicsItem object.
func NewQGraphicsItem() *QGraphicsItem {
	g := newQGraphicsItem(QGraphicsItem_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsItem2 constructs a new QGraphicsItem object.
func NewQGraphicsItem2(parent *QGraphicsItem) *QGraphicsItem {
	g := newQGraphicsItem(QGraphicsItem_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsItem) Scene() *QGraphicsScene {
	return newQGraphicsScene(QGraphicsItem_Scene(this.h))
}

func (this *QGraphicsItem) ParentItem() *QGraphicsItem {
	return newQGraphicsItem(QGraphicsItem_ParentItem(this.h))
}

func (this *QGraphicsItem) TopLevelItem() *QGraphicsItem {
	return newQGraphicsItem(QGraphicsItem_TopLevelItem(this.h))
}

func (this *QGraphicsItem) ParentObject() *QGraphicsObject {
	return newQGraphicsObject(QGraphicsItem_ParentObject(this.h))
}

func (this *QGraphicsItem) ParentWidget() *QGraphicsWidget {
	return newQGraphicsWidget(QGraphicsItem_ParentWidget(this.h))
}

func (this *QGraphicsItem) TopLevelWidget() *QGraphicsWidget {
	return newQGraphicsWidget(QGraphicsItem_TopLevelWidget(this.h))
}

func (this *QGraphicsItem) Window() *QGraphicsWidget {
	return newQGraphicsWidget(QGraphicsItem_Window(this.h))
}

func (this *QGraphicsItem) Panel() *QGraphicsItem {
	return newQGraphicsItem(QGraphicsItem_Panel(this.h))
}

func (this *QGraphicsItem) SetParentItem(parent *QGraphicsItem) {
	QGraphicsItem_SetParentItem(this.h, parent.cPointer())
}

func (this *QGraphicsItem) ChildItems() []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsItem_ChildItems(this.h)
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsItem) IsWidget() bool {
	return (bool)(QGraphicsItem_IsWidget(this.h))
}

func (this *QGraphicsItem) IsWindow() bool {
	return (bool)(QGraphicsItem_IsWindow(this.h))
}

func (this *QGraphicsItem) IsPanel() bool {
	return (bool)(QGraphicsItem_IsPanel(this.h))
}

func (this *QGraphicsItem) ToGraphicsObject() *QGraphicsObject {
	return newQGraphicsObject(QGraphicsItem_ToGraphicsObject(this.h))
}

func (this *QGraphicsItem) ToGraphicsObject2() *QGraphicsObject {
	return newQGraphicsObject(QGraphicsItem_ToGraphicsObject2(this.h))
}

func (this *QGraphicsItem) Group() *QGraphicsItemGroup {
	return newQGraphicsItemGroup(QGraphicsItem_Group(this.h))
}

func (this *QGraphicsItem) SetGroup(group *QGraphicsItemGroup) {
	QGraphicsItem_SetGroup(this.h, group.cPointer())
}

func (this *QGraphicsItem) Flags() GraphicsItemFlags {
	xxxxxxxxx
}

func (this *QGraphicsItem) SetFlag(flag GraphicsItemFlag) {
	QGraphicsItem_SetFlag(this.h, flag)
}

func (this *QGraphicsItem) SetFlags(flags GraphicsItemFlags) {
	QGraphicsItem_SetFlags(this.h, flags)
}

func (this *QGraphicsItem) CacheMode() CacheMode {
	xxxxxxxxx
}

func (this *QGraphicsItem) SetCacheMode(mode CacheMode) {
	QGraphicsItem_SetCacheMode(this.h, mode)
}

func (this *QGraphicsItem) PanelModality() PanelModality {
	xxxxxxxxx
}

func (this *QGraphicsItem) SetPanelModality(panelModality PanelModality) {
	QGraphicsItem_SetPanelModality(this.h, panelModality)
}

func (this *QGraphicsItem) IsBlockedByModalPanel() bool {
	return (bool)(QGraphicsItem_IsBlockedByModalPanel(this.h))
}

func (this *QGraphicsItem) ToolTip() string {
	var _ms struct_miqt_string = QGraphicsItem_ToolTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsItem) SetToolTip(toolTip string) {
	toolTip_ms := struct_miqt_string{}
	toolTip_ms.data = CString(toolTip)
	toolTip_ms.len = size_t(len(toolTip))
	defer free(unsafe.Pointer(toolTip_ms.data))
	QGraphicsItem_SetToolTip(this.h, toolTip_ms)
}

func (this *QGraphicsItem) Cursor() *QCursor {
	_goptr := newQCursor(QGraphicsItem_Cursor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) SetCursor(cursor *QCursor) {
	QGraphicsItem_SetCursor(this.h, cursor.cPointer())
}

func (this *QGraphicsItem) HasCursor() bool {
	return (bool)(QGraphicsItem_HasCursor(this.h))
}

func (this *QGraphicsItem) UnsetCursor() {
	QGraphicsItem_UnsetCursor(this.h)
}

func (this *QGraphicsItem) IsVisible() bool {
	return (bool)(QGraphicsItem_IsVisible(this.h))
}

func (this *QGraphicsItem) IsVisibleTo(parent *QGraphicsItem) bool {
	return (bool)(QGraphicsItem_IsVisibleTo(this.h, parent.cPointer()))
}

func (this *QGraphicsItem) SetVisible(visible bool) {
	QGraphicsItem_SetVisible(this.h, (bool)(visible))
}

func (this *QGraphicsItem) Hide() {
	QGraphicsItem_Hide(this.h)
}

func (this *QGraphicsItem) Show() {
	QGraphicsItem_Show(this.h)
}

func (this *QGraphicsItem) IsEnabled() bool {
	return (bool)(QGraphicsItem_IsEnabled(this.h))
}

func (this *QGraphicsItem) SetEnabled(enabled bool) {
	QGraphicsItem_SetEnabled(this.h, (bool)(enabled))
}

func (this *QGraphicsItem) IsSelected() bool {
	return (bool)(QGraphicsItem_IsSelected(this.h))
}

func (this *QGraphicsItem) SetSelected(selected bool) {
	QGraphicsItem_SetSelected(this.h, (bool)(selected))
}

func (this *QGraphicsItem) AcceptDrops() bool {
	return (bool)(QGraphicsItem_AcceptDrops(this.h))
}

func (this *QGraphicsItem) SetAcceptDrops(on bool) {
	QGraphicsItem_SetAcceptDrops(this.h, (bool)(on))
}

func (this *QGraphicsItem) Opacity() float64 {
	return (float64)(QGraphicsItem_Opacity(this.h))
}

func (this *QGraphicsItem) EffectiveOpacity() float64 {
	return (float64)(QGraphicsItem_EffectiveOpacity(this.h))
}

func (this *QGraphicsItem) SetOpacity(opacity float64) {
	QGraphicsItem_SetOpacity(this.h, (double)(opacity))
}

func (this *QGraphicsItem) GraphicsEffect() *QGraphicsEffect {
	return newQGraphicsEffect(QGraphicsItem_GraphicsEffect(this.h))
}

func (this *QGraphicsItem) SetGraphicsEffect(effect *QGraphicsEffect) {
	QGraphicsItem_SetGraphicsEffect(this.h, effect.cPointer())
}

func (this *QGraphicsItem) AcceptedMouseButtons() MouseButton {
	return (MouseButton)(QGraphicsItem_AcceptedMouseButtons(this.h))
}

func (this *QGraphicsItem) SetAcceptedMouseButtons(buttons MouseButton) {
	QGraphicsItem_SetAcceptedMouseButtons(this.h, (int)(buttons))
}

func (this *QGraphicsItem) AcceptHoverEvents() bool {
	return (bool)(QGraphicsItem_AcceptHoverEvents(this.h))
}

func (this *QGraphicsItem) SetAcceptHoverEvents(enabled bool) {
	QGraphicsItem_SetAcceptHoverEvents(this.h, (bool)(enabled))
}

func (this *QGraphicsItem) AcceptTouchEvents() bool {
	return (bool)(QGraphicsItem_AcceptTouchEvents(this.h))
}

func (this *QGraphicsItem) SetAcceptTouchEvents(enabled bool) {
	QGraphicsItem_SetAcceptTouchEvents(this.h, (bool)(enabled))
}

func (this *QGraphicsItem) FiltersChildEvents() bool {
	return (bool)(QGraphicsItem_FiltersChildEvents(this.h))
}

func (this *QGraphicsItem) SetFiltersChildEvents(enabled bool) {
	QGraphicsItem_SetFiltersChildEvents(this.h, (bool)(enabled))
}

func (this *QGraphicsItem) HandlesChildEvents() bool {
	return (bool)(QGraphicsItem_HandlesChildEvents(this.h))
}

func (this *QGraphicsItem) SetHandlesChildEvents(enabled bool) {
	QGraphicsItem_SetHandlesChildEvents(this.h, (bool)(enabled))
}

func (this *QGraphicsItem) IsActive() bool {
	return (bool)(QGraphicsItem_IsActive(this.h))
}

func (this *QGraphicsItem) SetActive(active bool) {
	QGraphicsItem_SetActive(this.h, (bool)(active))
}

func (this *QGraphicsItem) HasFocus() bool {
	return (bool)(QGraphicsItem_HasFocus(this.h))
}

func (this *QGraphicsItem) SetFocus() {
	QGraphicsItem_SetFocus(this.h)
}

func (this *QGraphicsItem) ClearFocus() {
	QGraphicsItem_ClearFocus(this.h)
}

func (this *QGraphicsItem) FocusProxy() *QGraphicsItem {
	return newQGraphicsItem(QGraphicsItem_FocusProxy(this.h))
}

func (this *QGraphicsItem) SetFocusProxy(item *QGraphicsItem) {
	QGraphicsItem_SetFocusProxy(this.h, item.cPointer())
}

func (this *QGraphicsItem) FocusItem() *QGraphicsItem {
	return newQGraphicsItem(QGraphicsItem_FocusItem(this.h))
}

func (this *QGraphicsItem) FocusScopeItem() *QGraphicsItem {
	return newQGraphicsItem(QGraphicsItem_FocusScopeItem(this.h))
}

func (this *QGraphicsItem) GrabMouse() {
	QGraphicsItem_GrabMouse(this.h)
}

func (this *QGraphicsItem) UngrabMouse() {
	QGraphicsItem_UngrabMouse(this.h)
}

func (this *QGraphicsItem) GrabKeyboard() {
	QGraphicsItem_GrabKeyboard(this.h)
}

func (this *QGraphicsItem) UngrabKeyboard() {
	QGraphicsItem_UngrabKeyboard(this.h)
}

func (this *QGraphicsItem) Pos() *QPointF {
	_goptr := newQPointF(QGraphicsItem_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) X() float64 {
	return (float64)(QGraphicsItem_X(this.h))
}

func (this *QGraphicsItem) SetX(x float64) {
	QGraphicsItem_SetX(this.h, (double)(x))
}

func (this *QGraphicsItem) Y() float64 {
	return (float64)(QGraphicsItem_Y(this.h))
}

func (this *QGraphicsItem) SetY(y float64) {
	QGraphicsItem_SetY(this.h, (double)(y))
}

func (this *QGraphicsItem) ScenePos() *QPointF {
	_goptr := newQPointF(QGraphicsItem_ScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) SetPos(pos *QPointF) {
	QGraphicsItem_SetPos(this.h, pos.cPointer())
}

func (this *QGraphicsItem) SetPos2(x float64, y float64) {
	QGraphicsItem_SetPos2(this.h, (double)(x), (double)(y))
}

func (this *QGraphicsItem) MoveBy(dx float64, dy float64) {
	QGraphicsItem_MoveBy(this.h, (double)(dx), (double)(dy))
}

func (this *QGraphicsItem) EnsureVisible() {
	QGraphicsItem_EnsureVisible(this.h)
}

func (this *QGraphicsItem) EnsureVisible2(x float64, y float64, w float64, h float64) {
	QGraphicsItem_EnsureVisible2(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QGraphicsItem) Transform() *QTransform {
	_goptr := newQTransform(QGraphicsItem_Transform(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) SceneTransform() *QTransform {
	_goptr := newQTransform(QGraphicsItem_SceneTransform(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) DeviceTransform(viewportTransform *QTransform) *QTransform {
	_goptr := newQTransform(QGraphicsItem_DeviceTransform(this.h, viewportTransform.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) ItemTransform(other *QGraphicsItem) *QTransform {
	_goptr := newQTransform(QGraphicsItem_ItemTransform(this.h, other.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) SetTransform(matrix *QTransform) {
	QGraphicsItem_SetTransform(this.h, matrix.cPointer())
}

func (this *QGraphicsItem) ResetTransform() {
	QGraphicsItem_ResetTransform(this.h)
}

func (this *QGraphicsItem) SetRotation(angle float64) {
	QGraphicsItem_SetRotation(this.h, (double)(angle))
}

func (this *QGraphicsItem) Rotation() float64 {
	return (float64)(QGraphicsItem_Rotation(this.h))
}

func (this *QGraphicsItem) SetScale(scale float64) {
	QGraphicsItem_SetScale(this.h, (double)(scale))
}

func (this *QGraphicsItem) Scale() float64 {
	return (float64)(QGraphicsItem_Scale(this.h))
}

func (this *QGraphicsItem) Transformations() []*QGraphicsTransform {
	var _ma struct_miqt_array = QGraphicsItem_Transformations(this.h)
	_ret := make([]*QGraphicsTransform, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsTransform)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsTransform(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsItem) SetTransformations(transformations []*QGraphicsTransform) {
	transformations_CArray := (*[0xffff]*QGraphicsTransform)(malloc(size_t(8 * len(transformations))))
	defer free(unsafe.Pointer(transformations_CArray))
	for i := range transformations {
		transformations_CArray[i] = transformations[i].cPointer()
	}
	transformations_ma := struct_miqt_array{len: size_t(len(transformations)), data: unsafe.Pointer(transformations_CArray)}
	QGraphicsItem_SetTransformations(this.h, transformations_ma)
}

func (this *QGraphicsItem) TransformOriginPoint() *QPointF {
	_goptr := newQPointF(QGraphicsItem_TransformOriginPoint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) SetTransformOriginPoint(origin *QPointF) {
	QGraphicsItem_SetTransformOriginPoint(this.h, origin.cPointer())
}

func (this *QGraphicsItem) SetTransformOriginPoint2(ax float64, ay float64) {
	QGraphicsItem_SetTransformOriginPoint2(this.h, (double)(ax), (double)(ay))
}

func (this *QGraphicsItem) Advance(phase int) {
	QGraphicsItem_Advance(this.h, (int)(phase))
}

func (this *QGraphicsItem) ZValue() float64 {
	return (float64)(QGraphicsItem_ZValue(this.h))
}

func (this *QGraphicsItem) SetZValue(z float64) {
	QGraphicsItem_SetZValue(this.h, (double)(z))
}

func (this *QGraphicsItem) StackBefore(sibling *QGraphicsItem) {
	QGraphicsItem_StackBefore(this.h, sibling.cPointer())
}

func (this *QGraphicsItem) BoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsItem_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) ChildrenBoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsItem_ChildrenBoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) SceneBoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsItem_SceneBoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) Shape() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsItem_Shape(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) IsClipped() bool {
	return (bool)(QGraphicsItem_IsClipped(this.h))
}

func (this *QGraphicsItem) ClipPath() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsItem_ClipPath(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) Contains(point *QPointF) bool {
	return (bool)(QGraphicsItem_Contains(this.h, point.cPointer()))
}

func (this *QGraphicsItem) CollidesWithItem(other *QGraphicsItem, mode ItemSelectionMode) bool {
	return (bool)(QGraphicsItem_CollidesWithItem(this.h, other.cPointer(), (int)(mode)))
}

func (this *QGraphicsItem) CollidesWithPath(path *QPainterPath, mode ItemSelectionMode) bool {
	return (bool)(QGraphicsItem_CollidesWithPath(this.h, path.cPointer(), (int)(mode)))
}

func (this *QGraphicsItem) CollidingItems() []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsItem_CollidingItems(this.h)
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsItem) IsObscured() bool {
	return (bool)(QGraphicsItem_IsObscured(this.h))
}

func (this *QGraphicsItem) IsObscured2(x float64, y float64, w float64, h float64) bool {
	return (bool)(QGraphicsItem_IsObscured2(this.h, (double)(x), (double)(y), (double)(w), (double)(h)))
}

func (this *QGraphicsItem) IsObscuredBy(item *QGraphicsItem) bool {
	return (bool)(QGraphicsItem_IsObscuredBy(this.h, item.cPointer()))
}

func (this *QGraphicsItem) OpaqueArea() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsItem_OpaqueArea(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) BoundingRegion(itemToDeviceTransform *QTransform) *QRegion {
	_goptr := newQRegion(QGraphicsItem_BoundingRegion(this.h, itemToDeviceTransform.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) BoundingRegionGranularity() float64 {
	return (float64)(QGraphicsItem_BoundingRegionGranularity(this.h))
}

func (this *QGraphicsItem) SetBoundingRegionGranularity(granularity float64) {
	QGraphicsItem_SetBoundingRegionGranularity(this.h, (double)(granularity))
}

func (this *QGraphicsItem) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsItem_Paint(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsItem) Update() {
	QGraphicsItem_Update(this.h)
}

func (this *QGraphicsItem) Update2(x float64, y float64, width float64, height float64) {
	QGraphicsItem_Update2(this.h, (double)(x), (double)(y), (double)(width), (double)(height))
}

func (this *QGraphicsItem) Scroll(dx float64, dy float64) {
	QGraphicsItem_Scroll(this.h, (double)(dx), (double)(dy))
}

func (this *QGraphicsItem) MapToItem(item *QGraphicsItem, point *QPointF) *QPointF {
	_goptr := newQPointF(QGraphicsItem_MapToItem(this.h, item.cPointer(), point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapToParent(point *QPointF) *QPointF {
	_goptr := newQPointF(QGraphicsItem_MapToParent(this.h, point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapToScene(point *QPointF) *QPointF {
	_goptr := newQPointF(QGraphicsItem_MapToScene(this.h, point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapRectToItem(item *QGraphicsItem, rect *QRectF) *QRectF {
	_goptr := newQRectF(QGraphicsItem_MapRectToItem(this.h, item.cPointer(), rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapRectToParent(rect *QRectF) *QRectF {
	_goptr := newQRectF(QGraphicsItem_MapRectToParent(this.h, rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapRectToScene(rect *QRectF) *QRectF {
	_goptr := newQRectF(QGraphicsItem_MapRectToScene(this.h, rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapToItem4(item *QGraphicsItem, path *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QGraphicsItem_MapToItem4(this.h, item.cPointer(), path.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapToParentWithPath(path *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QGraphicsItem_MapToParentWithPath(this.h, path.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapToSceneWithPath(path *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QGraphicsItem_MapToSceneWithPath(this.h, path.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapFromItem(item *QGraphicsItem, point *QPointF) *QPointF {
	_goptr := newQPointF(QGraphicsItem_MapFromItem(this.h, item.cPointer(), point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapFromParent(point *QPointF) *QPointF {
	_goptr := newQPointF(QGraphicsItem_MapFromParent(this.h, point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapFromScene(point *QPointF) *QPointF {
	_goptr := newQPointF(QGraphicsItem_MapFromScene(this.h, point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapRectFromItem(item *QGraphicsItem, rect *QRectF) *QRectF {
	_goptr := newQRectF(QGraphicsItem_MapRectFromItem(this.h, item.cPointer(), rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapRectFromParent(rect *QRectF) *QRectF {
	_goptr := newQRectF(QGraphicsItem_MapRectFromParent(this.h, rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapRectFromScene(rect *QRectF) *QRectF {
	_goptr := newQRectF(QGraphicsItem_MapRectFromScene(this.h, rect.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapFromItem4(item *QGraphicsItem, path *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QGraphicsItem_MapFromItem4(this.h, item.cPointer(), path.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapFromParentWithPath(path *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QGraphicsItem_MapFromParentWithPath(this.h, path.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapFromSceneWithPath(path *QPainterPath) *QPainterPath {
	_goptr := newQPainterPath(QGraphicsItem_MapFromSceneWithPath(this.h, path.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapToItem5(item *QGraphicsItem, x float64, y float64) *QPointF {
	_goptr := newQPointF(QGraphicsItem_MapToItem5(this.h, item.cPointer(), (double)(x), (double)(y)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapToParent2(x float64, y float64) *QPointF {
	_goptr := newQPointF(QGraphicsItem_MapToParent2(this.h, (double)(x), (double)(y)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapToScene2(x float64, y float64) *QPointF {
	_goptr := newQPointF(QGraphicsItem_MapToScene2(this.h, (double)(x), (double)(y)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapRectToItem2(item *QGraphicsItem, x float64, y float64, w float64, h float64) *QRectF {
	_goptr := newQRectF(QGraphicsItem_MapRectToItem2(this.h, item.cPointer(), (double)(x), (double)(y), (double)(w), (double)(h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapRectToParent2(x float64, y float64, w float64, h float64) *QRectF {
	_goptr := newQRectF(QGraphicsItem_MapRectToParent2(this.h, (double)(x), (double)(y), (double)(w), (double)(h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapRectToScene2(x float64, y float64, w float64, h float64) *QRectF {
	_goptr := newQRectF(QGraphicsItem_MapRectToScene2(this.h, (double)(x), (double)(y), (double)(w), (double)(h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapFromItem5(item *QGraphicsItem, x float64, y float64) *QPointF {
	_goptr := newQPointF(QGraphicsItem_MapFromItem5(this.h, item.cPointer(), (double)(x), (double)(y)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapFromParent2(x float64, y float64) *QPointF {
	_goptr := newQPointF(QGraphicsItem_MapFromParent2(this.h, (double)(x), (double)(y)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapFromScene2(x float64, y float64) *QPointF {
	_goptr := newQPointF(QGraphicsItem_MapFromScene2(this.h, (double)(x), (double)(y)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapRectFromItem2(item *QGraphicsItem, x float64, y float64, w float64, h float64) *QRectF {
	_goptr := newQRectF(QGraphicsItem_MapRectFromItem2(this.h, item.cPointer(), (double)(x), (double)(y), (double)(w), (double)(h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapRectFromParent2(x float64, y float64, w float64, h float64) *QRectF {
	_goptr := newQRectF(QGraphicsItem_MapRectFromParent2(this.h, (double)(x), (double)(y), (double)(w), (double)(h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) MapRectFromScene2(x float64, y float64, w float64, h float64) *QRectF {
	_goptr := newQRectF(QGraphicsItem_MapRectFromScene2(this.h, (double)(x), (double)(y), (double)(w), (double)(h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) IsAncestorOf(child *QGraphicsItem) bool {
	return (bool)(QGraphicsItem_IsAncestorOf(this.h, child.cPointer()))
}

func (this *QGraphicsItem) CommonAncestorItem(other *QGraphicsItem) *QGraphicsItem {
	return newQGraphicsItem(QGraphicsItem_CommonAncestorItem(this.h, other.cPointer()))
}

func (this *QGraphicsItem) IsUnderMouse() bool {
	return (bool)(QGraphicsItem_IsUnderMouse(this.h))
}

func (this *QGraphicsItem) Data(key int) *QVariant {
	_goptr := newQVariant(QGraphicsItem_Data(this.h, (int)(key)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) SetData(key int, value *QVariant) {
	QGraphicsItem_SetData(this.h, (int)(key), value.cPointer())
}

func (this *QGraphicsItem) InputMethodHints() InputMethodHint {
	return (InputMethodHint)(QGraphicsItem_InputMethodHints(this.h))
}

func (this *QGraphicsItem) SetInputMethodHints(hints InputMethodHint) {
	QGraphicsItem_SetInputMethodHints(this.h, (int)(hints))
}

func (this *QGraphicsItem) Type() int {
	return (int)(QGraphicsItem_Type(this.h))
}

func (this *QGraphicsItem) InstallSceneEventFilter(filterItem *QGraphicsItem) {
	QGraphicsItem_InstallSceneEventFilter(this.h, filterItem.cPointer())
}

func (this *QGraphicsItem) RemoveSceneEventFilter(filterItem *QGraphicsItem) {
	QGraphicsItem_RemoveSceneEventFilter(this.h, filterItem.cPointer())
}

func (this *QGraphicsItem) SetFlag2(flag GraphicsItemFlag, enabled bool) {
	QGraphicsItem_SetFlag2(this.h, flag, (bool)(enabled))
}

func (this *QGraphicsItem) SetCacheMode2(mode CacheMode, cacheSize *QSize) {
	QGraphicsItem_SetCacheMode2(this.h, mode, cacheSize.cPointer())
}

func (this *QGraphicsItem) SetFocus1(focusReason FocusReason) {
	QGraphicsItem_SetFocus1(this.h, (int)(focusReason))
}

func (this *QGraphicsItem) EnsureVisible1(rect *QRectF) {
	QGraphicsItem_EnsureVisible1(this.h, rect.cPointer())
}

func (this *QGraphicsItem) EnsureVisible22(rect *QRectF, xmargin int) {
	QGraphicsItem_EnsureVisible22(this.h, rect.cPointer(), (int)(xmargin))
}

func (this *QGraphicsItem) EnsureVisible3(rect *QRectF, xmargin int, ymargin int) {
	QGraphicsItem_EnsureVisible3(this.h, rect.cPointer(), (int)(xmargin), (int)(ymargin))
}

func (this *QGraphicsItem) EnsureVisible5(x float64, y float64, w float64, h float64, xmargin int) {
	QGraphicsItem_EnsureVisible5(this.h, (double)(x), (double)(y), (double)(w), (double)(h), (int)(xmargin))
}

func (this *QGraphicsItem) EnsureVisible6(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	QGraphicsItem_EnsureVisible6(this.h, (double)(x), (double)(y), (double)(w), (double)(h), (int)(xmargin), (int)(ymargin))
}

func (this *QGraphicsItem) ItemTransform2(other *QGraphicsItem, ok *bool) *QTransform {
	_goptr := newQTransform(QGraphicsItem_ItemTransform2(this.h, other.cPointer(), (*bool)(unsafe.Pointer(ok))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItem) SetTransform2(matrix *QTransform, combine bool) {
	QGraphicsItem_SetTransform2(this.h, matrix.cPointer(), (bool)(combine))
}

func (this *QGraphicsItem) CollidingItems1(mode ItemSelectionMode) []*QGraphicsItem {
	var _ma struct_miqt_array = QGraphicsItem_CollidingItems1(this.h, (int)(mode))
	_ret := make([]*QGraphicsItem, int(_ma.len))
	_outCast := (*[0xffff]*QGraphicsItem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQGraphicsItem(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsItem) IsObscured1(rect *QRectF) bool {
	return (bool)(QGraphicsItem_IsObscured1(this.h, rect.cPointer()))
}

func (this *QGraphicsItem) Update1(rect *QRectF) {
	QGraphicsItem_Update1(this.h, rect.cPointer())
}

func (this *QGraphicsItem) Scroll3(dx float64, dy float64, rect *QRectF) {
	QGraphicsItem_Scroll3(this.h, (double)(dx), (double)(dy), rect.cPointer())
}

type QGraphicsObject struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsObject constructs a new QGraphicsObject object.
func NewQGraphicsObject() *QGraphicsObject {
	g := newQGraphicsObject(QGraphicsObject_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsObject2 constructs a new QGraphicsObject object.
func NewQGraphicsObject2(parent *QGraphicsItem) *QGraphicsObject {
	g := newQGraphicsObject(QGraphicsObject_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsObject) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsObject_MetaObject(this.h))
}

func (this *QGraphicsObject) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsObject_Metacast(this.h, param1_Cstring))
}

func QGraphicsObject_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsObject_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsObject) GrabGesture(typeVal GestureType) {
	QGraphicsObject_GrabGesture(this.h, (int)(typeVal))
}

func (this *QGraphicsObject) UngrabGesture(typeVal GestureType) {
	QGraphicsObject_UngrabGesture(this.h, (int)(typeVal))
}

func (this *QGraphicsObject) ParentChanged() {
	QGraphicsObject_ParentChanged(this.h)
}

func (this *QGraphicsObject) OnParentChanged(slot func()) {
	QGraphicsObject_connect_ParentChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_ParentChanged
func miqt_exec_callback_QGraphicsObject_ParentChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsObject) OpacityChanged() {
	QGraphicsObject_OpacityChanged(this.h)
}

func (this *QGraphicsObject) OnOpacityChanged(slot func()) {
	QGraphicsObject_connect_OpacityChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_OpacityChanged
func miqt_exec_callback_QGraphicsObject_OpacityChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsObject) VisibleChanged() {
	QGraphicsObject_VisibleChanged(this.h)
}

func (this *QGraphicsObject) OnVisibleChanged(slot func()) {
	QGraphicsObject_connect_VisibleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_VisibleChanged
func miqt_exec_callback_QGraphicsObject_VisibleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsObject) EnabledChanged() {
	QGraphicsObject_EnabledChanged(this.h)
}

func (this *QGraphicsObject) OnEnabledChanged(slot func()) {
	QGraphicsObject_connect_EnabledChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_EnabledChanged
func miqt_exec_callback_QGraphicsObject_EnabledChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsObject) XChanged() {
	QGraphicsObject_XChanged(this.h)
}

func (this *QGraphicsObject) OnXChanged(slot func()) {
	QGraphicsObject_connect_XChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_XChanged
func miqt_exec_callback_QGraphicsObject_XChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsObject) YChanged() {
	QGraphicsObject_YChanged(this.h)
}

func (this *QGraphicsObject) OnYChanged(slot func()) {
	QGraphicsObject_connect_YChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_YChanged
func miqt_exec_callback_QGraphicsObject_YChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsObject) ZChanged() {
	QGraphicsObject_ZChanged(this.h)
}

func (this *QGraphicsObject) OnZChanged(slot func()) {
	QGraphicsObject_connect_ZChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_ZChanged
func miqt_exec_callback_QGraphicsObject_ZChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsObject) RotationChanged() {
	QGraphicsObject_RotationChanged(this.h)
}

func (this *QGraphicsObject) OnRotationChanged(slot func()) {
	QGraphicsObject_connect_RotationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_RotationChanged
func miqt_exec_callback_QGraphicsObject_RotationChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsObject) ScaleChanged() {
	QGraphicsObject_ScaleChanged(this.h)
}

func (this *QGraphicsObject) OnScaleChanged(slot func()) {
	QGraphicsObject_connect_ScaleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_ScaleChanged
func miqt_exec_callback_QGraphicsObject_ScaleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsObject) ChildrenChanged() {
	QGraphicsObject_ChildrenChanged(this.h)
}

func (this *QGraphicsObject) OnChildrenChanged(slot func()) {
	QGraphicsObject_connect_ChildrenChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_ChildrenChanged
func miqt_exec_callback_QGraphicsObject_ChildrenChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsObject) WidthChanged() {
	QGraphicsObject_WidthChanged(this.h)
}

func (this *QGraphicsObject) OnWidthChanged(slot func()) {
	QGraphicsObject_connect_WidthChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_WidthChanged
func miqt_exec_callback_QGraphicsObject_WidthChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsObject) HeightChanged() {
	QGraphicsObject_HeightChanged(this.h)
}

func (this *QGraphicsObject) OnHeightChanged(slot func()) {
	QGraphicsObject_connect_HeightChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_HeightChanged
func miqt_exec_callback_QGraphicsObject_HeightChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QGraphicsObject_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsObject_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsObject_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsObject_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsObject) GrabGesture2(typeVal GestureType, flags GestureFlag) {
	QGraphicsObject_GrabGesture2(this.h, (int)(typeVal), (int)(flags))
}

func (this *QGraphicsObject) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsObject_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QGraphicsObject) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsObject_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_MetaObject
func miqt_exec_callback_QGraphicsObject_MetaObject(self QGraphicsObject, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsObject{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QGraphicsObject) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGraphicsObject_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGraphicsObject) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsObject_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsObject_Metacast
func miqt_exec_callback_QGraphicsObject_Metacast(self QGraphicsObject, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QGraphicsObject{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QAbstractGraphicsShapeItem struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractGraphicsShapeItem constructs a new QAbstractGraphicsShapeItem object.
func NewQAbstractGraphicsShapeItem() *QAbstractGraphicsShapeItem {
	g := newQAbstractGraphicsShapeItem(QAbstractGraphicsShapeItem_new())
	g.isSubclass = true
	return g
}

// NewQAbstractGraphicsShapeItem2 constructs a new QAbstractGraphicsShapeItem object.
func NewQAbstractGraphicsShapeItem2(parent *QGraphicsItem) *QAbstractGraphicsShapeItem {
	g := newQAbstractGraphicsShapeItem(QAbstractGraphicsShapeItem_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QAbstractGraphicsShapeItem) Pen() *QPen {
	_goptr := newQPen(QAbstractGraphicsShapeItem_Pen(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractGraphicsShapeItem) SetPen(pen *QPen) {
	QAbstractGraphicsShapeItem_SetPen(this.h, pen.cPointer())
}

func (this *QAbstractGraphicsShapeItem) Brush() *QBrush {
	_goptr := newQBrush(QAbstractGraphicsShapeItem_Brush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractGraphicsShapeItem) SetBrush(brush *QBrush) {
	QAbstractGraphicsShapeItem_SetBrush(this.h, brush.cPointer())
}

func (this *QAbstractGraphicsShapeItem) IsObscuredBy(item *QGraphicsItem) bool {
	return (bool)(QAbstractGraphicsShapeItem_IsObscuredBy(this.h, item.cPointer()))
}

func (this *QAbstractGraphicsShapeItem) OpaqueArea() *QPainterPath {
	_goptr := newQPainterPath(QAbstractGraphicsShapeItem_OpaqueArea(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QGraphicsPathItem struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsPathItem constructs a new QGraphicsPathItem object.
func NewQGraphicsPathItem() *QGraphicsPathItem {
	g := newQGraphicsPathItem(QGraphicsPathItem_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsPathItem2 constructs a new QGraphicsPathItem object.
func NewQGraphicsPathItem2(path *QPainterPath) *QGraphicsPathItem {
	g := newQGraphicsPathItem(QGraphicsPathItem_new2(path.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsPathItem3 constructs a new QGraphicsPathItem object.
func NewQGraphicsPathItem3(parent *QGraphicsItem) *QGraphicsPathItem {
	g := newQGraphicsPathItem(QGraphicsPathItem_new3(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsPathItem4 constructs a new QGraphicsPathItem object.
func NewQGraphicsPathItem4(path *QPainterPath, parent *QGraphicsItem) *QGraphicsPathItem {
	g := newQGraphicsPathItem(QGraphicsPathItem_new4(path.cPointer(), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsPathItem) Path() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsPathItem_Path(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsPathItem) SetPath(path *QPainterPath) {
	QGraphicsPathItem_SetPath(this.h, path.cPointer())
}

func (this *QGraphicsPathItem) BoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsPathItem_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsPathItem) Shape() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsPathItem_Shape(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsPathItem) Contains(point *QPointF) bool {
	return (bool)(QGraphicsPathItem_Contains(this.h, point.cPointer()))
}

func (this *QGraphicsPathItem) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsPathItem_Paint(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsPathItem) IsObscuredBy(item *QGraphicsItem) bool {
	return (bool)(QGraphicsPathItem_IsObscuredBy(this.h, item.cPointer()))
}

func (this *QGraphicsPathItem) OpaqueArea() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsPathItem_OpaqueArea(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsPathItem) Type() int {
	return (int)(QGraphicsPathItem_Type(this.h))
}

type QGraphicsRectItem struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsRectItem constructs a new QGraphicsRectItem object.
func NewQGraphicsRectItem() *QGraphicsRectItem {
	g := newQGraphicsRectItem(QGraphicsRectItem_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsRectItem2 constructs a new QGraphicsRectItem object.
func NewQGraphicsRectItem2(rect *QRectF) *QGraphicsRectItem {
	g := newQGraphicsRectItem(QGraphicsRectItem_new2(rect.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsRectItem3 constructs a new QGraphicsRectItem object.
func NewQGraphicsRectItem3(x float64, y float64, w float64, h float64) *QGraphicsRectItem {
	g := newQGraphicsRectItem(QGraphicsRectItem_new3((double)(x), (double)(y), (double)(w), (double)(h)))
	g.isSubclass = true
	return g
}

// NewQGraphicsRectItem4 constructs a new QGraphicsRectItem object.
func NewQGraphicsRectItem4(parent *QGraphicsItem) *QGraphicsRectItem {
	g := newQGraphicsRectItem(QGraphicsRectItem_new4(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsRectItem5 constructs a new QGraphicsRectItem object.
func NewQGraphicsRectItem5(rect *QRectF, parent *QGraphicsItem) *QGraphicsRectItem {
	g := newQGraphicsRectItem(QGraphicsRectItem_new5(rect.cPointer(), parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsRectItem6 constructs a new QGraphicsRectItem object.
func NewQGraphicsRectItem6(x float64, y float64, w float64, h float64, parent *QGraphicsItem) *QGraphicsRectItem {
	g := newQGraphicsRectItem(QGraphicsRectItem_new6((double)(x), (double)(y), (double)(w), (double)(h), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsRectItem) Rect() *QRectF {
	_goptr := newQRectF(QGraphicsRectItem_Rect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsRectItem) SetRect(rect *QRectF) {
	QGraphicsRectItem_SetRect(this.h, rect.cPointer())
}

func (this *QGraphicsRectItem) SetRect2(x float64, y float64, w float64, h float64) {
	QGraphicsRectItem_SetRect2(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QGraphicsRectItem) BoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsRectItem_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsRectItem) Shape() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsRectItem_Shape(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsRectItem) Contains(point *QPointF) bool {
	return (bool)(QGraphicsRectItem_Contains(this.h, point.cPointer()))
}

func (this *QGraphicsRectItem) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsRectItem_Paint(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsRectItem) IsObscuredBy(item *QGraphicsItem) bool {
	return (bool)(QGraphicsRectItem_IsObscuredBy(this.h, item.cPointer()))
}

func (this *QGraphicsRectItem) OpaqueArea() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsRectItem_OpaqueArea(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsRectItem) Type() int {
	return (int)(QGraphicsRectItem_Type(this.h))
}

type QGraphicsEllipseItem struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsEllipseItem constructs a new QGraphicsEllipseItem object.
func NewQGraphicsEllipseItem() *QGraphicsEllipseItem {
	g := newQGraphicsEllipseItem(QGraphicsEllipseItem_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsEllipseItem2 constructs a new QGraphicsEllipseItem object.
func NewQGraphicsEllipseItem2(rect *QRectF) *QGraphicsEllipseItem {
	g := newQGraphicsEllipseItem(QGraphicsEllipseItem_new2(rect.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsEllipseItem3 constructs a new QGraphicsEllipseItem object.
func NewQGraphicsEllipseItem3(x float64, y float64, w float64, h float64) *QGraphicsEllipseItem {
	g := newQGraphicsEllipseItem(QGraphicsEllipseItem_new3((double)(x), (double)(y), (double)(w), (double)(h)))
	g.isSubclass = true
	return g
}

// NewQGraphicsEllipseItem4 constructs a new QGraphicsEllipseItem object.
func NewQGraphicsEllipseItem4(parent *QGraphicsItem) *QGraphicsEllipseItem {
	g := newQGraphicsEllipseItem(QGraphicsEllipseItem_new4(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsEllipseItem5 constructs a new QGraphicsEllipseItem object.
func NewQGraphicsEllipseItem5(rect *QRectF, parent *QGraphicsItem) *QGraphicsEllipseItem {
	g := newQGraphicsEllipseItem(QGraphicsEllipseItem_new5(rect.cPointer(), parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsEllipseItem6 constructs a new QGraphicsEllipseItem object.
func NewQGraphicsEllipseItem6(x float64, y float64, w float64, h float64, parent *QGraphicsItem) *QGraphicsEllipseItem {
	g := newQGraphicsEllipseItem(QGraphicsEllipseItem_new6((double)(x), (double)(y), (double)(w), (double)(h), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsEllipseItem) Rect() *QRectF {
	_goptr := newQRectF(QGraphicsEllipseItem_Rect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsEllipseItem) SetRect(rect *QRectF) {
	QGraphicsEllipseItem_SetRect(this.h, rect.cPointer())
}

func (this *QGraphicsEllipseItem) SetRect2(x float64, y float64, w float64, h float64) {
	QGraphicsEllipseItem_SetRect2(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QGraphicsEllipseItem) StartAngle() int {
	return (int)(QGraphicsEllipseItem_StartAngle(this.h))
}

func (this *QGraphicsEllipseItem) SetStartAngle(angle int) {
	QGraphicsEllipseItem_SetStartAngle(this.h, (int)(angle))
}

func (this *QGraphicsEllipseItem) SpanAngle() int {
	return (int)(QGraphicsEllipseItem_SpanAngle(this.h))
}

func (this *QGraphicsEllipseItem) SetSpanAngle(angle int) {
	QGraphicsEllipseItem_SetSpanAngle(this.h, (int)(angle))
}

func (this *QGraphicsEllipseItem) BoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsEllipseItem_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsEllipseItem) Shape() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsEllipseItem_Shape(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsEllipseItem) Contains(point *QPointF) bool {
	return (bool)(QGraphicsEllipseItem_Contains(this.h, point.cPointer()))
}

func (this *QGraphicsEllipseItem) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsEllipseItem_Paint(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsEllipseItem) IsObscuredBy(item *QGraphicsItem) bool {
	return (bool)(QGraphicsEllipseItem_IsObscuredBy(this.h, item.cPointer()))
}

func (this *QGraphicsEllipseItem) OpaqueArea() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsEllipseItem_OpaqueArea(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsEllipseItem) Type() int {
	return (int)(QGraphicsEllipseItem_Type(this.h))
}

type QGraphicsPolygonItem struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsPolygonItem constructs a new QGraphicsPolygonItem object.
func NewQGraphicsPolygonItem() *QGraphicsPolygonItem {
	g := newQGraphicsPolygonItem(QGraphicsPolygonItem_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsPolygonItem2 constructs a new QGraphicsPolygonItem object.
func NewQGraphicsPolygonItem2(parent *QGraphicsItem) *QGraphicsPolygonItem {
	g := newQGraphicsPolygonItem(QGraphicsPolygonItem_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsPolygonItem) FillRule() FillRule {
	return (FillRule)(QGraphicsPolygonItem_FillRule(this.h))
}

func (this *QGraphicsPolygonItem) SetFillRule(rule FillRule) {
	QGraphicsPolygonItem_SetFillRule(this.h, (int)(rule))
}

func (this *QGraphicsPolygonItem) BoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsPolygonItem_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsPolygonItem) Shape() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsPolygonItem_Shape(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsPolygonItem) Contains(point *QPointF) bool {
	return (bool)(QGraphicsPolygonItem_Contains(this.h, point.cPointer()))
}

func (this *QGraphicsPolygonItem) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsPolygonItem_Paint(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsPolygonItem) IsObscuredBy(item *QGraphicsItem) bool {
	return (bool)(QGraphicsPolygonItem_IsObscuredBy(this.h, item.cPointer()))
}

func (this *QGraphicsPolygonItem) OpaqueArea() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsPolygonItem_OpaqueArea(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsPolygonItem) Type() int {
	return (int)(QGraphicsPolygonItem_Type(this.h))
}

type QGraphicsLineItem struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsLineItem constructs a new QGraphicsLineItem object.
func NewQGraphicsLineItem() *QGraphicsLineItem {
	g := newQGraphicsLineItem(QGraphicsLineItem_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsLineItem2 constructs a new QGraphicsLineItem object.
func NewQGraphicsLineItem2(line *QLineF) *QGraphicsLineItem {
	g := newQGraphicsLineItem(QGraphicsLineItem_new2(line.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsLineItem3 constructs a new QGraphicsLineItem object.
func NewQGraphicsLineItem3(x1 float64, y1 float64, x2 float64, y2 float64) *QGraphicsLineItem {
	g := newQGraphicsLineItem(QGraphicsLineItem_new3((double)(x1), (double)(y1), (double)(x2), (double)(y2)))
	g.isSubclass = true
	return g
}

// NewQGraphicsLineItem4 constructs a new QGraphicsLineItem object.
func NewQGraphicsLineItem4(parent *QGraphicsItem) *QGraphicsLineItem {
	g := newQGraphicsLineItem(QGraphicsLineItem_new4(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsLineItem5 constructs a new QGraphicsLineItem object.
func NewQGraphicsLineItem5(line *QLineF, parent *QGraphicsItem) *QGraphicsLineItem {
	g := newQGraphicsLineItem(QGraphicsLineItem_new5(line.cPointer(), parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsLineItem6 constructs a new QGraphicsLineItem object.
func NewQGraphicsLineItem6(x1 float64, y1 float64, x2 float64, y2 float64, parent *QGraphicsItem) *QGraphicsLineItem {
	g := newQGraphicsLineItem(QGraphicsLineItem_new6((double)(x1), (double)(y1), (double)(x2), (double)(y2), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsLineItem) Pen() *QPen {
	_goptr := newQPen(QGraphicsLineItem_Pen(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLineItem) SetPen(pen *QPen) {
	QGraphicsLineItem_SetPen(this.h, pen.cPointer())
}

func (this *QGraphicsLineItem) Line() *QLineF {
	_goptr := newQLineF(QGraphicsLineItem_Line(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLineItem) SetLine(line *QLineF) {
	QGraphicsLineItem_SetLine(this.h, line.cPointer())
}

func (this *QGraphicsLineItem) SetLine2(x1 float64, y1 float64, x2 float64, y2 float64) {
	QGraphicsLineItem_SetLine2(this.h, (double)(x1), (double)(y1), (double)(x2), (double)(y2))
}

func (this *QGraphicsLineItem) BoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsLineItem_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLineItem) Shape() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsLineItem_Shape(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLineItem) Contains(point *QPointF) bool {
	return (bool)(QGraphicsLineItem_Contains(this.h, point.cPointer()))
}

func (this *QGraphicsLineItem) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsLineItem_Paint(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsLineItem) IsObscuredBy(item *QGraphicsItem) bool {
	return (bool)(QGraphicsLineItem_IsObscuredBy(this.h, item.cPointer()))
}

func (this *QGraphicsLineItem) OpaqueArea() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsLineItem_OpaqueArea(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsLineItem) Type() int {
	return (int)(QGraphicsLineItem_Type(this.h))
}

type QGraphicsPixmapItem struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsPixmapItem constructs a new QGraphicsPixmapItem object.
func NewQGraphicsPixmapItem() *QGraphicsPixmapItem {
	g := newQGraphicsPixmapItem(QGraphicsPixmapItem_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsPixmapItem2 constructs a new QGraphicsPixmapItem object.
func NewQGraphicsPixmapItem2(pixmap *QPixmap) *QGraphicsPixmapItem {
	g := newQGraphicsPixmapItem(QGraphicsPixmapItem_new2(pixmap.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsPixmapItem3 constructs a new QGraphicsPixmapItem object.
func NewQGraphicsPixmapItem3(parent *QGraphicsItem) *QGraphicsPixmapItem {
	g := newQGraphicsPixmapItem(QGraphicsPixmapItem_new3(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsPixmapItem4 constructs a new QGraphicsPixmapItem object.
func NewQGraphicsPixmapItem4(pixmap *QPixmap, parent *QGraphicsItem) *QGraphicsPixmapItem {
	g := newQGraphicsPixmapItem(QGraphicsPixmapItem_new4(pixmap.cPointer(), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsPixmapItem) Pixmap() *QPixmap {
	_goptr := newQPixmap(QGraphicsPixmapItem_Pixmap(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsPixmapItem) SetPixmap(pixmap *QPixmap) {
	QGraphicsPixmapItem_SetPixmap(this.h, pixmap.cPointer())
}

func (this *QGraphicsPixmapItem) TransformationMode() TransformationMode {
	return (TransformationMode)(QGraphicsPixmapItem_TransformationMode(this.h))
}

func (this *QGraphicsPixmapItem) SetTransformationMode(mode TransformationMode) {
	QGraphicsPixmapItem_SetTransformationMode(this.h, (int)(mode))
}

func (this *QGraphicsPixmapItem) Offset() *QPointF {
	_goptr := newQPointF(QGraphicsPixmapItem_Offset(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsPixmapItem) SetOffset(offset *QPointF) {
	QGraphicsPixmapItem_SetOffset(this.h, offset.cPointer())
}

func (this *QGraphicsPixmapItem) SetOffset2(x float64, y float64) {
	QGraphicsPixmapItem_SetOffset2(this.h, (double)(x), (double)(y))
}

func (this *QGraphicsPixmapItem) BoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsPixmapItem_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsPixmapItem) Shape() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsPixmapItem_Shape(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsPixmapItem) Contains(point *QPointF) bool {
	return (bool)(QGraphicsPixmapItem_Contains(this.h, point.cPointer()))
}

func (this *QGraphicsPixmapItem) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsPixmapItem_Paint(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsPixmapItem) IsObscuredBy(item *QGraphicsItem) bool {
	return (bool)(QGraphicsPixmapItem_IsObscuredBy(this.h, item.cPointer()))
}

func (this *QGraphicsPixmapItem) OpaqueArea() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsPixmapItem_OpaqueArea(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsPixmapItem) Type() int {
	return (int)(QGraphicsPixmapItem_Type(this.h))
}

func (this *QGraphicsPixmapItem) ShapeMode() ShapeMode {
	xxxxxxxxx
}

func (this *QGraphicsPixmapItem) SetShapeMode(mode ShapeMode) {
	QGraphicsPixmapItem_SetShapeMode(this.h, mode)
}

type QGraphicsTextItem struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsTextItem constructs a new QGraphicsTextItem object.
func NewQGraphicsTextItem() *QGraphicsTextItem {
	g := newQGraphicsTextItem(QGraphicsTextItem_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsTextItem2 constructs a new QGraphicsTextItem object.
func NewQGraphicsTextItem2(text string) *QGraphicsTextItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQGraphicsTextItem(QGraphicsTextItem_new2(text_ms))
	g.isSubclass = true
	return g
}

// NewQGraphicsTextItem3 constructs a new QGraphicsTextItem object.
func NewQGraphicsTextItem3(parent *QGraphicsItem) *QGraphicsTextItem {
	g := newQGraphicsTextItem(QGraphicsTextItem_new3(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsTextItem4 constructs a new QGraphicsTextItem object.
func NewQGraphicsTextItem4(text string, parent *QGraphicsItem) *QGraphicsTextItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQGraphicsTextItem(QGraphicsTextItem_new4(text_ms, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsTextItem) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsTextItem_MetaObject(this.h))
}

func (this *QGraphicsTextItem) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsTextItem_Metacast(this.h, param1_Cstring))
}

func QGraphicsTextItem_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsTextItem_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsTextItem) ToHtml() string {
	var _ms struct_miqt_string = QGraphicsTextItem_ToHtml(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsTextItem) SetHtml(html string) {
	html_ms := struct_miqt_string{}
	html_ms.data = CString(html)
	html_ms.len = size_t(len(html))
	defer free(unsafe.Pointer(html_ms.data))
	QGraphicsTextItem_SetHtml(this.h, html_ms)
}

func (this *QGraphicsTextItem) ToPlainText() string {
	var _ms struct_miqt_string = QGraphicsTextItem_ToPlainText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsTextItem) SetPlainText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QGraphicsTextItem_SetPlainText(this.h, text_ms)
}

func (this *QGraphicsTextItem) Font() *QFont {
	_goptr := newQFont(QGraphicsTextItem_Font(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsTextItem) SetFont(font *QFont) {
	QGraphicsTextItem_SetFont(this.h, font.cPointer())
}

func (this *QGraphicsTextItem) SetDefaultTextColor(c *QColor) {
	QGraphicsTextItem_SetDefaultTextColor(this.h, c.cPointer())
}

func (this *QGraphicsTextItem) DefaultTextColor() *QColor {
	_goptr := newQColor(QGraphicsTextItem_DefaultTextColor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsTextItem) BoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsTextItem_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsTextItem) Shape() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsTextItem_Shape(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsTextItem) Contains(point *QPointF) bool {
	return (bool)(QGraphicsTextItem_Contains(this.h, point.cPointer()))
}

func (this *QGraphicsTextItem) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsTextItem_Paint(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsTextItem) IsObscuredBy(item *QGraphicsItem) bool {
	return (bool)(QGraphicsTextItem_IsObscuredBy(this.h, item.cPointer()))
}

func (this *QGraphicsTextItem) OpaqueArea() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsTextItem_OpaqueArea(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsTextItem) Type() int {
	return (int)(QGraphicsTextItem_Type(this.h))
}

func (this *QGraphicsTextItem) SetTextWidth(width float64) {
	QGraphicsTextItem_SetTextWidth(this.h, (double)(width))
}

func (this *QGraphicsTextItem) TextWidth() float64 {
	return (float64)(QGraphicsTextItem_TextWidth(this.h))
}

func (this *QGraphicsTextItem) AdjustSize() {
	QGraphicsTextItem_AdjustSize(this.h)
}

func (this *QGraphicsTextItem) SetDocument(document *QTextDocument) {
	QGraphicsTextItem_SetDocument(this.h, document.cPointer())
}

func (this *QGraphicsTextItem) Document() *QTextDocument {
	return newQTextDocument(QGraphicsTextItem_Document(this.h))
}

func (this *QGraphicsTextItem) SetTextInteractionFlags(flags TextInteractionFlag) {
	QGraphicsTextItem_SetTextInteractionFlags(this.h, (int)(flags))
}

func (this *QGraphicsTextItem) TextInteractionFlags() TextInteractionFlag {
	return (TextInteractionFlag)(QGraphicsTextItem_TextInteractionFlags(this.h))
}

func (this *QGraphicsTextItem) SetTabChangesFocus(b bool) {
	QGraphicsTextItem_SetTabChangesFocus(this.h, (bool)(b))
}

func (this *QGraphicsTextItem) TabChangesFocus() bool {
	return (bool)(QGraphicsTextItem_TabChangesFocus(this.h))
}

func (this *QGraphicsTextItem) SetOpenExternalLinks(open bool) {
	QGraphicsTextItem_SetOpenExternalLinks(this.h, (bool)(open))
}

func (this *QGraphicsTextItem) OpenExternalLinks() bool {
	return (bool)(QGraphicsTextItem_OpenExternalLinks(this.h))
}

func (this *QGraphicsTextItem) SetTextCursor(cursor *QTextCursor) {
	QGraphicsTextItem_SetTextCursor(this.h, cursor.cPointer())
}

func (this *QGraphicsTextItem) TextCursor() *QTextCursor {
	_goptr := newQTextCursor(QGraphicsTextItem_TextCursor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsTextItem) LinkActivated(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QGraphicsTextItem_LinkActivated(this.h, param1_ms)
}

func (this *QGraphicsTextItem) OnLinkActivated(slot func(param1 string)) {
	QGraphicsTextItem_connect_LinkActivated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTextItem_LinkActivated
func miqt_exec_callback_QGraphicsTextItem_LinkActivated(cb intptr_t, param1 struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var param1_ms struct_miqt_string = param1
	param1_ret := GoStringN(param1_ms.data, int(int64(param1_ms.len)))
	free(unsafe.Pointer(param1_ms.data))
	slotval1 := param1_ret

	gofunc(slotval1)
}

func (this *QGraphicsTextItem) LinkHovered(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QGraphicsTextItem_LinkHovered(this.h, param1_ms)
}

func (this *QGraphicsTextItem) OnLinkHovered(slot func(param1 string)) {
	QGraphicsTextItem_connect_LinkHovered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTextItem_LinkHovered
func miqt_exec_callback_QGraphicsTextItem_LinkHovered(cb intptr_t, param1 struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var param1_ms struct_miqt_string = param1
	param1_ret := GoStringN(param1_ms.data, int(int64(param1_ms.len)))
	free(unsafe.Pointer(param1_ms.data))
	slotval1 := param1_ret

	gofunc(slotval1)
}

func QGraphicsTextItem_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsTextItem_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsTextItem_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsTextItem_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsTextItem) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsTextItem_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QGraphicsTextItem) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsTextItem_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTextItem_MetaObject
func miqt_exec_callback_QGraphicsTextItem_MetaObject(self QGraphicsTextItem, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsTextItem{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QGraphicsTextItem) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGraphicsTextItem_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGraphicsTextItem) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsTextItem_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsTextItem_Metacast
func miqt_exec_callback_QGraphicsTextItem_Metacast(self QGraphicsTextItem, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QGraphicsTextItem{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QGraphicsSimpleTextItem struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsSimpleTextItem constructs a new QGraphicsSimpleTextItem object.
func NewQGraphicsSimpleTextItem() *QGraphicsSimpleTextItem {
	g := newQGraphicsSimpleTextItem(QGraphicsSimpleTextItem_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsSimpleTextItem2 constructs a new QGraphicsSimpleTextItem object.
func NewQGraphicsSimpleTextItem2(text string) *QGraphicsSimpleTextItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQGraphicsSimpleTextItem(QGraphicsSimpleTextItem_new2(text_ms))
	g.isSubclass = true
	return g
}

// NewQGraphicsSimpleTextItem3 constructs a new QGraphicsSimpleTextItem object.
func NewQGraphicsSimpleTextItem3(parent *QGraphicsItem) *QGraphicsSimpleTextItem {
	g := newQGraphicsSimpleTextItem(QGraphicsSimpleTextItem_new3(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsSimpleTextItem4 constructs a new QGraphicsSimpleTextItem object.
func NewQGraphicsSimpleTextItem4(text string, parent *QGraphicsItem) *QGraphicsSimpleTextItem {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQGraphicsSimpleTextItem(QGraphicsSimpleTextItem_new4(text_ms, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsSimpleTextItem) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QGraphicsSimpleTextItem_SetText(this.h, text_ms)
}

func (this *QGraphicsSimpleTextItem) Text() string {
	var _ms struct_miqt_string = QGraphicsSimpleTextItem_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsSimpleTextItem) SetFont(font *QFont) {
	QGraphicsSimpleTextItem_SetFont(this.h, font.cPointer())
}

func (this *QGraphicsSimpleTextItem) Font() *QFont {
	_goptr := newQFont(QGraphicsSimpleTextItem_Font(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSimpleTextItem) BoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsSimpleTextItem_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSimpleTextItem) Shape() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsSimpleTextItem_Shape(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSimpleTextItem) Contains(point *QPointF) bool {
	return (bool)(QGraphicsSimpleTextItem_Contains(this.h, point.cPointer()))
}

func (this *QGraphicsSimpleTextItem) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsSimpleTextItem_Paint(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsSimpleTextItem) IsObscuredBy(item *QGraphicsItem) bool {
	return (bool)(QGraphicsSimpleTextItem_IsObscuredBy(this.h, item.cPointer()))
}

func (this *QGraphicsSimpleTextItem) OpaqueArea() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsSimpleTextItem_OpaqueArea(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSimpleTextItem) Type() int {
	return (int)(QGraphicsSimpleTextItem_Type(this.h))
}

type QGraphicsItemGroup struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsItemGroup constructs a new QGraphicsItemGroup object.
func NewQGraphicsItemGroup() *QGraphicsItemGroup {
	g := newQGraphicsItemGroup(QGraphicsItemGroup_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsItemGroup2 constructs a new QGraphicsItemGroup object.
func NewQGraphicsItemGroup2(parent *QGraphicsItem) *QGraphicsItemGroup {
	g := newQGraphicsItemGroup(QGraphicsItemGroup_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QGraphicsItemGroup) AddToGroup(item *QGraphicsItem) {
	QGraphicsItemGroup_AddToGroup(this.h, item.cPointer())
}

func (this *QGraphicsItemGroup) RemoveFromGroup(item *QGraphicsItem) {
	QGraphicsItemGroup_RemoveFromGroup(this.h, item.cPointer())
}

func (this *QGraphicsItemGroup) BoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsItemGroup_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItemGroup) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsItemGroup_Paint(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsItemGroup) IsObscuredBy(item *QGraphicsItem) bool {
	return (bool)(QGraphicsItemGroup_IsObscuredBy(this.h, item.cPointer()))
}

func (this *QGraphicsItemGroup) OpaqueArea() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsItemGroup_OpaqueArea(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsItemGroup) Type() int {
	return (int)(QGraphicsItemGroup_Type(this.h))
}
