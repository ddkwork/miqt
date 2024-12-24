package qt

import (
	"unsafe"
)

type QGraphicsSceneContextMenuEvent__Reason int

const (
	QGraphicsSceneContextMenuEvent__Mouse    QGraphicsSceneContextMenuEvent__Reason = 0
	QGraphicsSceneContextMenuEvent__Keyboard QGraphicsSceneContextMenuEvent__Reason = 1
	QGraphicsSceneContextMenuEvent__Other    QGraphicsSceneContextMenuEvent__Reason = 2
)

type QGraphicsSceneEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsSceneEvent constructs a new QGraphicsSceneEvent object.
func NewQGraphicsSceneEvent(typeVal Type) *QGraphicsSceneEvent {
	ret := newQGraphicsSceneEvent(QGraphicsSceneEvent_new(typeVal))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsSceneEvent) Widget() *QWidget {
	return newQWidget(QGraphicsSceneEvent_Widget(this.h))
}

func (this *QGraphicsSceneEvent) SetWidget(widget *QWidget) {
	QGraphicsSceneEvent_SetWidget(this.h, widget.cPointer())
}

func (this *QGraphicsSceneEvent) Timestamp() uint64 {
	return (uint64)(QGraphicsSceneEvent_Timestamp(this.h))
}

func (this *QGraphicsSceneEvent) SetTimestamp(ts uint64) {
	QGraphicsSceneEvent_SetTimestamp(this.h, (ulonglong)(ts))
}

type QGraphicsSceneMouseEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsSceneMouseEvent constructs a new QGraphicsSceneMouseEvent object.
func NewQGraphicsSceneMouseEvent() *QGraphicsSceneMouseEvent {
	ret := newQGraphicsSceneMouseEvent(QGraphicsSceneMouseEvent_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsSceneMouseEvent2 constructs a new QGraphicsSceneMouseEvent object.
func NewQGraphicsSceneMouseEvent2(typeVal Type) *QGraphicsSceneMouseEvent {
	ret := newQGraphicsSceneMouseEvent(QGraphicsSceneMouseEvent_new2(typeVal))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsSceneMouseEvent) Pos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneMouseEvent_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneMouseEvent) SetPos(pos *QPointF) {
	QGraphicsSceneMouseEvent_SetPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneMouseEvent) ScenePos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneMouseEvent_ScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneMouseEvent) SetScenePos(pos *QPointF) {
	QGraphicsSceneMouseEvent_SetScenePos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneMouseEvent) ScreenPos() *QPoint {
	_goptr := newQPoint(QGraphicsSceneMouseEvent_ScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneMouseEvent) SetScreenPos(pos *QPoint) {
	QGraphicsSceneMouseEvent_SetScreenPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneMouseEvent) ButtonDownPos(button MouseButton) *QPointF {
	_goptr := newQPointF(QGraphicsSceneMouseEvent_ButtonDownPos(this.h, (int)(button)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneMouseEvent) SetButtonDownPos(button MouseButton, pos *QPointF) {
	QGraphicsSceneMouseEvent_SetButtonDownPos(this.h, (int)(button), pos.cPointer())
}

func (this *QGraphicsSceneMouseEvent) ButtonDownScenePos(button MouseButton) *QPointF {
	_goptr := newQPointF(QGraphicsSceneMouseEvent_ButtonDownScenePos(this.h, (int)(button)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneMouseEvent) SetButtonDownScenePos(button MouseButton, pos *QPointF) {
	QGraphicsSceneMouseEvent_SetButtonDownScenePos(this.h, (int)(button), pos.cPointer())
}

func (this *QGraphicsSceneMouseEvent) ButtonDownScreenPos(button MouseButton) *QPoint {
	_goptr := newQPoint(QGraphicsSceneMouseEvent_ButtonDownScreenPos(this.h, (int)(button)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneMouseEvent) SetButtonDownScreenPos(button MouseButton, pos *QPoint) {
	QGraphicsSceneMouseEvent_SetButtonDownScreenPos(this.h, (int)(button), pos.cPointer())
}

func (this *QGraphicsSceneMouseEvent) LastPos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneMouseEvent_LastPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneMouseEvent) SetLastPos(pos *QPointF) {
	QGraphicsSceneMouseEvent_SetLastPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneMouseEvent) LastScenePos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneMouseEvent_LastScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneMouseEvent) SetLastScenePos(pos *QPointF) {
	QGraphicsSceneMouseEvent_SetLastScenePos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneMouseEvent) LastScreenPos() *QPoint {
	_goptr := newQPoint(QGraphicsSceneMouseEvent_LastScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneMouseEvent) SetLastScreenPos(pos *QPoint) {
	QGraphicsSceneMouseEvent_SetLastScreenPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneMouseEvent) Buttons() MouseButton {
	return (MouseButton)(QGraphicsSceneMouseEvent_Buttons(this.h))
}

func (this *QGraphicsSceneMouseEvent) SetButtons(buttons MouseButton) {
	QGraphicsSceneMouseEvent_SetButtons(this.h, (int)(buttons))
}

func (this *QGraphicsSceneMouseEvent) Button() MouseButton {
	return (MouseButton)(QGraphicsSceneMouseEvent_Button(this.h))
}

func (this *QGraphicsSceneMouseEvent) SetButton(button MouseButton) {
	QGraphicsSceneMouseEvent_SetButton(this.h, (int)(button))
}

func (this *QGraphicsSceneMouseEvent) Modifiers() KeyboardModifier {
	return (KeyboardModifier)(QGraphicsSceneMouseEvent_Modifiers(this.h))
}

func (this *QGraphicsSceneMouseEvent) SetModifiers(modifiers KeyboardModifier) {
	QGraphicsSceneMouseEvent_SetModifiers(this.h, (int)(modifiers))
}

func (this *QGraphicsSceneMouseEvent) Source() MouseEventSource {
	return (MouseEventSource)(QGraphicsSceneMouseEvent_Source(this.h))
}

func (this *QGraphicsSceneMouseEvent) SetSource(source MouseEventSource) {
	QGraphicsSceneMouseEvent_SetSource(this.h, (int)(source))
}

func (this *QGraphicsSceneMouseEvent) Flags() MouseEventFlag {
	return (MouseEventFlag)(QGraphicsSceneMouseEvent_Flags(this.h))
}

func (this *QGraphicsSceneMouseEvent) SetFlags(flags MouseEventFlag) {
	QGraphicsSceneMouseEvent_SetFlags(this.h, (int)(flags))
}

type QGraphicsSceneWheelEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsSceneWheelEvent constructs a new QGraphicsSceneWheelEvent object.
func NewQGraphicsSceneWheelEvent() *QGraphicsSceneWheelEvent {
	ret := newQGraphicsSceneWheelEvent(QGraphicsSceneWheelEvent_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsSceneWheelEvent2 constructs a new QGraphicsSceneWheelEvent object.
func NewQGraphicsSceneWheelEvent2(typeVal Type) *QGraphicsSceneWheelEvent {
	ret := newQGraphicsSceneWheelEvent(QGraphicsSceneWheelEvent_new2(typeVal))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsSceneWheelEvent) Pos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneWheelEvent_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneWheelEvent) SetPos(pos *QPointF) {
	QGraphicsSceneWheelEvent_SetPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneWheelEvent) ScenePos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneWheelEvent_ScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneWheelEvent) SetScenePos(pos *QPointF) {
	QGraphicsSceneWheelEvent_SetScenePos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneWheelEvent) ScreenPos() *QPoint {
	_goptr := newQPoint(QGraphicsSceneWheelEvent_ScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneWheelEvent) SetScreenPos(pos *QPoint) {
	QGraphicsSceneWheelEvent_SetScreenPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneWheelEvent) Buttons() MouseButton {
	return (MouseButton)(QGraphicsSceneWheelEvent_Buttons(this.h))
}

func (this *QGraphicsSceneWheelEvent) SetButtons(buttons MouseButton) {
	QGraphicsSceneWheelEvent_SetButtons(this.h, (int)(buttons))
}

func (this *QGraphicsSceneWheelEvent) Modifiers() KeyboardModifier {
	return (KeyboardModifier)(QGraphicsSceneWheelEvent_Modifiers(this.h))
}

func (this *QGraphicsSceneWheelEvent) SetModifiers(modifiers KeyboardModifier) {
	QGraphicsSceneWheelEvent_SetModifiers(this.h, (int)(modifiers))
}

func (this *QGraphicsSceneWheelEvent) Delta() int {
	return (int)(QGraphicsSceneWheelEvent_Delta(this.h))
}

func (this *QGraphicsSceneWheelEvent) SetDelta(delta int) {
	QGraphicsSceneWheelEvent_SetDelta(this.h, (int)(delta))
}

func (this *QGraphicsSceneWheelEvent) Orientation() Orientation {
	return (Orientation)(QGraphicsSceneWheelEvent_Orientation(this.h))
}

func (this *QGraphicsSceneWheelEvent) SetOrientation(orientation Orientation) {
	QGraphicsSceneWheelEvent_SetOrientation(this.h, (int)(orientation))
}

func (this *QGraphicsSceneWheelEvent) Phase() ScrollPhase {
	return (ScrollPhase)(QGraphicsSceneWheelEvent_Phase(this.h))
}

func (this *QGraphicsSceneWheelEvent) SetPhase(scrollPhase ScrollPhase) {
	QGraphicsSceneWheelEvent_SetPhase(this.h, (int)(scrollPhase))
}

func (this *QGraphicsSceneWheelEvent) PixelDelta() *QPoint {
	_goptr := newQPoint(QGraphicsSceneWheelEvent_PixelDelta(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneWheelEvent) SetPixelDelta(delta QPoint) {
	QGraphicsSceneWheelEvent_SetPixelDelta(this.h, delta.cPointer())
}

func (this *QGraphicsSceneWheelEvent) IsInverted() bool {
	return (bool)(QGraphicsSceneWheelEvent_IsInverted(this.h))
}

func (this *QGraphicsSceneWheelEvent) SetInverted(inverted bool) {
	QGraphicsSceneWheelEvent_SetInverted(this.h, (bool)(inverted))
}

type QGraphicsSceneContextMenuEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsSceneContextMenuEvent constructs a new QGraphicsSceneContextMenuEvent object.
func NewQGraphicsSceneContextMenuEvent() *QGraphicsSceneContextMenuEvent {
	ret := newQGraphicsSceneContextMenuEvent(QGraphicsSceneContextMenuEvent_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsSceneContextMenuEvent2 constructs a new QGraphicsSceneContextMenuEvent object.
func NewQGraphicsSceneContextMenuEvent2(typeVal Type) *QGraphicsSceneContextMenuEvent {
	ret := newQGraphicsSceneContextMenuEvent(QGraphicsSceneContextMenuEvent_new2(typeVal))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsSceneContextMenuEvent) Pos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneContextMenuEvent_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneContextMenuEvent) SetPos(pos *QPointF) {
	QGraphicsSceneContextMenuEvent_SetPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneContextMenuEvent) ScenePos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneContextMenuEvent_ScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneContextMenuEvent) SetScenePos(pos *QPointF) {
	QGraphicsSceneContextMenuEvent_SetScenePos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneContextMenuEvent) ScreenPos() *QPoint {
	_goptr := newQPoint(QGraphicsSceneContextMenuEvent_ScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneContextMenuEvent) SetScreenPos(pos *QPoint) {
	QGraphicsSceneContextMenuEvent_SetScreenPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneContextMenuEvent) Modifiers() KeyboardModifier {
	return (KeyboardModifier)(QGraphicsSceneContextMenuEvent_Modifiers(this.h))
}

func (this *QGraphicsSceneContextMenuEvent) SetModifiers(modifiers KeyboardModifier) {
	QGraphicsSceneContextMenuEvent_SetModifiers(this.h, (int)(modifiers))
}

func (this *QGraphicsSceneContextMenuEvent) Reason() Reason {
	xxxxxxxxx
}

func (this *QGraphicsSceneContextMenuEvent) SetReason(reason Reason) {
	QGraphicsSceneContextMenuEvent_SetReason(this.h, reason)
}

type QGraphicsSceneHoverEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsSceneHoverEvent constructs a new QGraphicsSceneHoverEvent object.
func NewQGraphicsSceneHoverEvent() *QGraphicsSceneHoverEvent {
	ret := newQGraphicsSceneHoverEvent(QGraphicsSceneHoverEvent_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsSceneHoverEvent2 constructs a new QGraphicsSceneHoverEvent object.
func NewQGraphicsSceneHoverEvent2(typeVal Type) *QGraphicsSceneHoverEvent {
	ret := newQGraphicsSceneHoverEvent(QGraphicsSceneHoverEvent_new2(typeVal))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsSceneHoverEvent) Pos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneHoverEvent_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneHoverEvent) SetPos(pos *QPointF) {
	QGraphicsSceneHoverEvent_SetPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneHoverEvent) ScenePos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneHoverEvent_ScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneHoverEvent) SetScenePos(pos *QPointF) {
	QGraphicsSceneHoverEvent_SetScenePos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneHoverEvent) ScreenPos() *QPoint {
	_goptr := newQPoint(QGraphicsSceneHoverEvent_ScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneHoverEvent) SetScreenPos(pos *QPoint) {
	QGraphicsSceneHoverEvent_SetScreenPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneHoverEvent) LastPos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneHoverEvent_LastPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneHoverEvent) SetLastPos(pos *QPointF) {
	QGraphicsSceneHoverEvent_SetLastPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneHoverEvent) LastScenePos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneHoverEvent_LastScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneHoverEvent) SetLastScenePos(pos *QPointF) {
	QGraphicsSceneHoverEvent_SetLastScenePos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneHoverEvent) LastScreenPos() *QPoint {
	_goptr := newQPoint(QGraphicsSceneHoverEvent_LastScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneHoverEvent) SetLastScreenPos(pos *QPoint) {
	QGraphicsSceneHoverEvent_SetLastScreenPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneHoverEvent) Modifiers() KeyboardModifier {
	return (KeyboardModifier)(QGraphicsSceneHoverEvent_Modifiers(this.h))
}

func (this *QGraphicsSceneHoverEvent) SetModifiers(modifiers KeyboardModifier) {
	QGraphicsSceneHoverEvent_SetModifiers(this.h, (int)(modifiers))
}

type QGraphicsSceneHelpEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsSceneHelpEvent constructs a new QGraphicsSceneHelpEvent object.
func NewQGraphicsSceneHelpEvent() *QGraphicsSceneHelpEvent {
	ret := newQGraphicsSceneHelpEvent(QGraphicsSceneHelpEvent_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsSceneHelpEvent2 constructs a new QGraphicsSceneHelpEvent object.
func NewQGraphicsSceneHelpEvent2(typeVal Type) *QGraphicsSceneHelpEvent {
	ret := newQGraphicsSceneHelpEvent(QGraphicsSceneHelpEvent_new2(typeVal))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsSceneHelpEvent) ScenePos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneHelpEvent_ScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneHelpEvent) SetScenePos(pos *QPointF) {
	QGraphicsSceneHelpEvent_SetScenePos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneHelpEvent) ScreenPos() *QPoint {
	_goptr := newQPoint(QGraphicsSceneHelpEvent_ScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneHelpEvent) SetScreenPos(pos *QPoint) {
	QGraphicsSceneHelpEvent_SetScreenPos(this.h, pos.cPointer())
}

type QGraphicsSceneDragDropEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsSceneDragDropEvent constructs a new QGraphicsSceneDragDropEvent object.
func NewQGraphicsSceneDragDropEvent() *QGraphicsSceneDragDropEvent {
	ret := newQGraphicsSceneDragDropEvent(QGraphicsSceneDragDropEvent_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsSceneDragDropEvent2 constructs a new QGraphicsSceneDragDropEvent object.
func NewQGraphicsSceneDragDropEvent2(typeVal Type) *QGraphicsSceneDragDropEvent {
	ret := newQGraphicsSceneDragDropEvent(QGraphicsSceneDragDropEvent_new2(typeVal))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsSceneDragDropEvent) Pos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneDragDropEvent_Pos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneDragDropEvent) SetPos(pos *QPointF) {
	QGraphicsSceneDragDropEvent_SetPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneDragDropEvent) ScenePos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneDragDropEvent_ScenePos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneDragDropEvent) SetScenePos(pos *QPointF) {
	QGraphicsSceneDragDropEvent_SetScenePos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneDragDropEvent) ScreenPos() *QPoint {
	_goptr := newQPoint(QGraphicsSceneDragDropEvent_ScreenPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneDragDropEvent) SetScreenPos(pos *QPoint) {
	QGraphicsSceneDragDropEvent_SetScreenPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneDragDropEvent) Buttons() MouseButton {
	return (MouseButton)(QGraphicsSceneDragDropEvent_Buttons(this.h))
}

func (this *QGraphicsSceneDragDropEvent) SetButtons(buttons MouseButton) {
	QGraphicsSceneDragDropEvent_SetButtons(this.h, (int)(buttons))
}

func (this *QGraphicsSceneDragDropEvent) Modifiers() KeyboardModifier {
	return (KeyboardModifier)(QGraphicsSceneDragDropEvent_Modifiers(this.h))
}

func (this *QGraphicsSceneDragDropEvent) SetModifiers(modifiers KeyboardModifier) {
	QGraphicsSceneDragDropEvent_SetModifiers(this.h, (int)(modifiers))
}

func (this *QGraphicsSceneDragDropEvent) PossibleActions() DropAction {
	return (DropAction)(QGraphicsSceneDragDropEvent_PossibleActions(this.h))
}

func (this *QGraphicsSceneDragDropEvent) SetPossibleActions(actions DropAction) {
	QGraphicsSceneDragDropEvent_SetPossibleActions(this.h, (int)(actions))
}

func (this *QGraphicsSceneDragDropEvent) ProposedAction() DropAction {
	return (DropAction)(QGraphicsSceneDragDropEvent_ProposedAction(this.h))
}

func (this *QGraphicsSceneDragDropEvent) SetProposedAction(action DropAction) {
	QGraphicsSceneDragDropEvent_SetProposedAction(this.h, (int)(action))
}

func (this *QGraphicsSceneDragDropEvent) AcceptProposedAction() {
	QGraphicsSceneDragDropEvent_AcceptProposedAction(this.h)
}

func (this *QGraphicsSceneDragDropEvent) DropAction() DropAction {
	return (DropAction)(QGraphicsSceneDragDropEvent_DropAction(this.h))
}

func (this *QGraphicsSceneDragDropEvent) SetDropAction(action DropAction) {
	QGraphicsSceneDragDropEvent_SetDropAction(this.h, (int)(action))
}

func (this *QGraphicsSceneDragDropEvent) Source() *QWidget {
	return newQWidget(QGraphicsSceneDragDropEvent_Source(this.h))
}

func (this *QGraphicsSceneDragDropEvent) SetSource(source *QWidget) {
	QGraphicsSceneDragDropEvent_SetSource(this.h, source.cPointer())
}

func (this *QGraphicsSceneDragDropEvent) MimeData() *QMimeData {
	return newQMimeData(QGraphicsSceneDragDropEvent_MimeData(this.h))
}

func (this *QGraphicsSceneDragDropEvent) SetMimeData(data *QMimeData) {
	QGraphicsSceneDragDropEvent_SetMimeData(this.h, data.cPointer())
}

type QGraphicsSceneResizeEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsSceneResizeEvent constructs a new QGraphicsSceneResizeEvent object.
func NewQGraphicsSceneResizeEvent() *QGraphicsSceneResizeEvent {
	ret := newQGraphicsSceneResizeEvent(QGraphicsSceneResizeEvent_new())
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsSceneResizeEvent) OldSize() *QSizeF {
	_goptr := newQSizeF(QGraphicsSceneResizeEvent_OldSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneResizeEvent) SetOldSize(size *QSizeF) {
	QGraphicsSceneResizeEvent_SetOldSize(this.h, size.cPointer())
}

func (this *QGraphicsSceneResizeEvent) NewSize() *QSizeF {
	_goptr := newQSizeF(QGraphicsSceneResizeEvent_NewSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneResizeEvent) SetNewSize(size *QSizeF) {
	QGraphicsSceneResizeEvent_SetNewSize(this.h, size.cPointer())
}

type QGraphicsSceneMoveEvent struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsSceneMoveEvent constructs a new QGraphicsSceneMoveEvent object.
func NewQGraphicsSceneMoveEvent() *QGraphicsSceneMoveEvent {
	ret := newQGraphicsSceneMoveEvent(QGraphicsSceneMoveEvent_new())
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsSceneMoveEvent) OldPos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneMoveEvent_OldPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneMoveEvent) SetOldPos(pos *QPointF) {
	QGraphicsSceneMoveEvent_SetOldPos(this.h, pos.cPointer())
}

func (this *QGraphicsSceneMoveEvent) NewPos() *QPointF {
	_goptr := newQPointF(QGraphicsSceneMoveEvent_NewPos(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSceneMoveEvent) SetNewPos(pos *QPointF) {
	QGraphicsSceneMoveEvent_SetNewPos(this.h, pos.cPointer())
}
