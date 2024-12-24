package qt

import (
	"unsafe"
)

type QDrawBorderPixmap__DrawingHint int

const (
	QDrawBorderPixmap__OpaqueTopLeft     QDrawBorderPixmap__DrawingHint = 1
	QDrawBorderPixmap__OpaqueTop         QDrawBorderPixmap__DrawingHint = 2
	QDrawBorderPixmap__OpaqueTopRight    QDrawBorderPixmap__DrawingHint = 4
	QDrawBorderPixmap__OpaqueLeft        QDrawBorderPixmap__DrawingHint = 8
	QDrawBorderPixmap__OpaqueCenter      QDrawBorderPixmap__DrawingHint = 16
	QDrawBorderPixmap__OpaqueRight       QDrawBorderPixmap__DrawingHint = 32
	QDrawBorderPixmap__OpaqueBottomLeft  QDrawBorderPixmap__DrawingHint = 64
	QDrawBorderPixmap__OpaqueBottom      QDrawBorderPixmap__DrawingHint = 128
	QDrawBorderPixmap__OpaqueBottomRight QDrawBorderPixmap__DrawingHint = 256
	QDrawBorderPixmap__OpaqueCorners     QDrawBorderPixmap__DrawingHint = 325
	QDrawBorderPixmap__OpaqueEdges       QDrawBorderPixmap__DrawingHint = 170
	QDrawBorderPixmap__OpaqueFrame       QDrawBorderPixmap__DrawingHint = 495
	QDrawBorderPixmap__OpaqueAll         QDrawBorderPixmap__DrawingHint = 511
)

type QTileRules struct {
	h          uintptr
	isSubclass bool
}

// NewQTileRules constructs a new QTileRules object.
func NewQTileRules(horizontalRule TileRule, verticalRule TileRule) *QTileRules {
	g := newQTileRules(QTileRules_new((int)(horizontalRule), (int)(verticalRule)))
	g.isSubclass = true
	return g
}

// NewQTileRules2 constructs a new QTileRules object.
func NewQTileRules2() *QTileRules {
	g := newQTileRules(QTileRules_new2())
	g.isSubclass = true
	return g
}

// NewQTileRules3 constructs a new QTileRules object.
func NewQTileRules3(param1 *QTileRules) *QTileRules {
	g := newQTileRules(QTileRules_new3(param1.cPointer()))
	g.isSubclass = true
	return g
}

// NewQTileRules4 constructs a new QTileRules object.
func NewQTileRules4(rule TileRule) *QTileRules {
	g := newQTileRules(QTileRules_new4((int)(rule)))
	g.isSubclass = true
	return g
}
