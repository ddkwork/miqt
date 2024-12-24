package qt

import (
	"unsafe"
)

type QCursor struct {
	h          uintptr
	isSubclass bool
}

// NewQCursor constructs a new QCursor object.
func NewQCursor() *QCursor {
	g := newQCursor(QCursor_new())
	g.isSubclass = true
	return g
}

// NewQCursor2 constructs a new QCursor object.
func NewQCursor2(shape CursorShape) *QCursor {
	g := newQCursor(QCursor_new2((int)(shape)))
	g.isSubclass = true
	return g
}

// NewQCursor3 constructs a new QCursor object.
func NewQCursor3(bitmap *QBitmap, mask *QBitmap) *QCursor {
	g := newQCursor(QCursor_new3(bitmap.cPointer(), mask.cPointer()))
	g.isSubclass = true
	return g
}

// NewQCursor4 constructs a new QCursor object.
func NewQCursor4(pixmap *QPixmap) *QCursor {
	g := newQCursor(QCursor_new4(pixmap.cPointer()))
	g.isSubclass = true
	return g
}

// NewQCursor5 constructs a new QCursor object.
func NewQCursor5(cursor *QCursor) *QCursor {
	g := newQCursor(QCursor_new5(cursor.cPointer()))
	g.isSubclass = true
	return g
}

// NewQCursor6 constructs a new QCursor object.
func NewQCursor6(bitmap *QBitmap, mask *QBitmap, hotX int) *QCursor {
	g := newQCursor(QCursor_new6(bitmap.cPointer(), mask.cPointer(), (int)(hotX)))
	g.isSubclass = true
	return g
}

// NewQCursor7 constructs a new QCursor object.
func NewQCursor7(bitmap *QBitmap, mask *QBitmap, hotX int, hotY int) *QCursor {
	g := newQCursor(QCursor_new7(bitmap.cPointer(), mask.cPointer(), (int)(hotX), (int)(hotY)))
	g.isSubclass = true
	return g
}

// NewQCursor8 constructs a new QCursor object.
func NewQCursor8(pixmap *QPixmap, hotX int) *QCursor {
	g := newQCursor(QCursor_new8(pixmap.cPointer(), (int)(hotX)))
	g.isSubclass = true
	return g
}

// NewQCursor9 constructs a new QCursor object.
func NewQCursor9(pixmap *QPixmap, hotX int, hotY int) *QCursor {
	g := newQCursor(QCursor_new9(pixmap.cPointer(), (int)(hotX), (int)(hotY)))
	g.isSubclass = true
	return g
}

func (this *QCursor) OperatorAssign(cursor *QCursor) {
	QCursor_OperatorAssign(this.h, cursor.cPointer())
}

func (this *QCursor) Swap(other *QCursor) {
	QCursor_Swap(this.h, other.cPointer())
}

func (this *QCursor) Shape() CursorShape {
	return (CursorShape)(QCursor_Shape(this.h))
}

func (this *QCursor) SetShape(newShape CursorShape) {
	QCursor_SetShape(this.h, (int)(newShape))
}

func (this *QCursor) Bitmap(param1 ReturnByValueConstant) *QBitmap {
	_goptr := newQBitmap(QCursor_Bitmap(this.h, (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCursor) Mask(param1 ReturnByValueConstant) *QBitmap {
	_goptr := newQBitmap(QCursor_Mask(this.h, (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCursor) Bitmap2() *QBitmap {
	_goptr := newQBitmap(QCursor_Bitmap2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCursor) Mask2() *QBitmap {
	_goptr := newQBitmap(QCursor_Mask2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCursor) Pixmap() *QPixmap {
	_goptr := newQPixmap(QCursor_Pixmap(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCursor) HotSpot() *QPoint {
	_goptr := newQPoint(QCursor_HotSpot(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCursor_Pos() *QPoint {
	_goptr := newQPoint(QCursor_Pos())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCursor_PosWithScreen(screen *QScreen) *QPoint {
	_goptr := newQPoint(QCursor_PosWithScreen(screen.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QCursor_SetPos(x int, y int) {
	QCursor_SetPos((int)(x), (int)(y))
}

func QCursor_SetPos2(screen *QScreen, x int, y int) {
	QCursor_SetPos2(screen.cPointer(), (int)(x), (int)(y))
}

func QCursor_SetPosWithQPoint(p *QPoint) {
	QCursor_SetPosWithQPoint(p.cPointer())
}

func QCursor_SetPos3(screen *QScreen, p *QPoint) {
	QCursor_SetPos3(screen.cPointer(), p.cPointer())
}
