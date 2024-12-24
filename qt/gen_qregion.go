package qt
import (
	"unsafe"
)

		type QRegion__RegionType int
		const (
QRegion__Rectangle QRegion__RegionType = 0
QRegion__Ellipse QRegion__RegionType = 1

)


		type QRegion struct {
			h uintptr
			isSubclass bool
}
		
			// NewQRegion constructs a new QRegion object.
			func NewQRegion() *QRegion {
								
				ret := newQRegion(QRegion_new())
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQRegion2 constructs a new QRegion object.
			func NewQRegion2(x int, y int, w int, h int) *QRegion {
								
				ret := newQRegion(QRegion_new2((int)(x), (int)(y), (int)(w), (int)(h)))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQRegion3 constructs a new QRegion object.
			func NewQRegion3(r *QRect) *QRegion {
								
				ret := newQRegion(QRegion_new3(r.cPointer()))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQRegion4 constructs a new QRegion object.
			func NewQRegion4(region *QRegion) *QRegion {
								
				ret := newQRegion(QRegion_new4(region.cPointer()))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQRegion5 constructs a new QRegion object.
			func NewQRegion5(bitmap *QBitmap) *QRegion {
								
				ret := newQRegion(QRegion_new5(bitmap.cPointer()))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQRegion6 constructs a new QRegion object.
			func NewQRegion6(x int, y int, w int, h int, t RegionType) *QRegion {
								
				ret := newQRegion(QRegion_new6((int)(x), (int)(y), (int)(w), (int)(h), t))
				ret.isSubclass = true
				return ret
			}
			
			
			// NewQRegion7 constructs a new QRegion object.
			func NewQRegion7(r *QRect, t RegionType) *QRegion {
								
				ret := newQRegion(QRegion_new7(r.cPointer(), t))
				ret.isSubclass = true
				return ret
			}
			
			
			func (this *QRegion) OperatorAssign(param1 *QRegion)  {
				 QRegion_OperatorAssign(this.h, param1.cPointer())
}
			
			func (this *QRegion) Swap(other *QRegion)  {
				 QRegion_Swap(this.h, other.cPointer())
}
			
			func (this *QRegion) IsEmpty() bool {
				return (bool)(QRegion_IsEmpty(this.h))
}
			
			func (this *QRegion) IsNull() bool {
				return (bool)(QRegion_IsNull(this.h))
}
			
			func (this *QRegion) Begin() const_iterator {
				xxxxxxxxx}
			
			func (this *QRegion) Cbegin() const_iterator {
				xxxxxxxxx}
			
			func (this *QRegion) End() const_iterator {
				xxxxxxxxx}
			
			func (this *QRegion) Cend() const_iterator {
				xxxxxxxxx}
			
			func (this *QRegion) Rbegin() const_reverse_iterator {
				xxxxxxxxx}
			
			func (this *QRegion) Crbegin() const_reverse_iterator {
				xxxxxxxxx}
			
			func (this *QRegion) Rend() const_reverse_iterator {
				xxxxxxxxx}
			
			func (this *QRegion) Crend() const_reverse_iterator {
				xxxxxxxxx}
			
			func (this *QRegion) Contains(p *QPoint) bool {
				return (bool)(QRegion_Contains(this.h, p.cPointer()))
}
			
			func (this *QRegion) ContainsWithQRect(r *QRect) bool {
				return (bool)(QRegion_ContainsWithQRect(this.h, r.cPointer()))
}
			
			func (this *QRegion) Translate(dx int, dy int)  {
				 QRegion_Translate(this.h, (int)(dx), (int)(dy))
}
			
			func (this *QRegion) TranslateWithQPoint(p *QPoint)  {
				 QRegion_TranslateWithQPoint(this.h, p.cPointer())
}
			
			func (this *QRegion) Translated(dx int, dy int) *QRegion {
				_goptr := newQRegion(QRegion_Translated(this.h, (int)(dx), (int)(dy)))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) TranslatedWithQPoint(p *QPoint) *QRegion {
				_goptr := newQRegion(QRegion_TranslatedWithQPoint(this.h, p.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) United(r *QRegion) *QRegion {
				_goptr := newQRegion(QRegion_United(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) UnitedWithQRect(r *QRect) *QRegion {
				_goptr := newQRegion(QRegion_UnitedWithQRect(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) Intersected(r *QRegion) *QRegion {
				_goptr := newQRegion(QRegion_Intersected(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) IntersectedWithQRect(r *QRect) *QRegion {
				_goptr := newQRegion(QRegion_IntersectedWithQRect(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) Subtracted(r *QRegion) *QRegion {
				_goptr := newQRegion(QRegion_Subtracted(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) Xored(r *QRegion) *QRegion {
				_goptr := newQRegion(QRegion_Xored(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) Intersects(r *QRegion) bool {
				return (bool)(QRegion_Intersects(this.h, r.cPointer()))
}
			
			func (this *QRegion) IntersectsWithQRect(r *QRect) bool {
				return (bool)(QRegion_IntersectsWithQRect(this.h, r.cPointer()))
}
			
			func (this *QRegion) BoundingRect() *QRect {
				_goptr := newQRect(QRegion_BoundingRect(this.h))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) SetRects(rect *QRect, num int)  {
				 QRegion_SetRects(this.h, rect.cPointer(), (int)(num))
}
			
			func (this *QRegion) SetRectsWithQSpanLesserconstQRectGreater(r QSpan<const QRect>)  {
				 QRegion_SetRectsWithQSpanLesserconstQRectGreater(this.h, r)
}
			
			func (this *QRegion) Rects() QSpan<const QRect> {
				xxxxxxxxx}
			
			func (this *QRegion) RectCount() int {
				return (int)(QRegion_RectCount(this.h))
}
			
			func (this *QRegion) OperatorBitwiseOr(r *QRegion) *QRegion {
				_goptr := newQRegion(QRegion_OperatorBitwiseOr(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) OperatorPlus(r *QRegion) *QRegion {
				_goptr := newQRegion(QRegion_OperatorPlus(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) OperatorPlusWithQRect(r *QRect) *QRegion {
				_goptr := newQRegion(QRegion_OperatorPlusWithQRect(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) OperatorBitwiseAnd(r *QRegion) *QRegion {
				_goptr := newQRegion(QRegion_OperatorBitwiseAnd(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) OperatorBitwiseAndWithQRect(r *QRect) *QRegion {
				_goptr := newQRegion(QRegion_OperatorBitwiseAndWithQRect(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) OperatorMinus(r *QRegion) *QRegion {
				_goptr := newQRegion(QRegion_OperatorMinus(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) OperatorBitwiseNot(r *QRegion) *QRegion {
				_goptr := newQRegion(QRegion_OperatorBitwiseNot(this.h, r.cPointer()))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			
			func (this *QRegion) OperatorBitwiseOrAssign(r *QRegion)  {
				 QRegion_OperatorBitwiseOrAssign(this.h, r.cPointer())
}
			
			func (this *QRegion) OperatorPlusAssign(r *QRegion) *QRegion {
				return newQRegion(QRegion_OperatorPlusAssign(this.h, r.cPointer()))
}
			
			func (this *QRegion) OperatorPlusAssignWithQRect(r *QRect) *QRegion {
				return newQRegion(QRegion_OperatorPlusAssignWithQRect(this.h, r.cPointer()))
}
			
			func (this *QRegion) OperatorBitwiseAndAssign(r *QRegion)  {
				 QRegion_OperatorBitwiseAndAssign(this.h, r.cPointer())
}
			
			func (this *QRegion) OperatorBitwiseAndAssignWithQRect(r *QRect)  {
				 QRegion_OperatorBitwiseAndAssignWithQRect(this.h, r.cPointer())
}
			
			func (this *QRegion) OperatorMinusAssign(r *QRegion) *QRegion {
				return newQRegion(QRegion_OperatorMinusAssign(this.h, r.cPointer()))
}
			
			func (this *QRegion) OperatorBitwiseNotAssign(r *QRegion)  {
				 QRegion_OperatorBitwiseNotAssign(this.h, r.cPointer())
}
			
			func (this *QRegion) OperatorEqual(r *QRegion) bool {
				return (bool)(QRegion_OperatorEqual(this.h, r.cPointer()))
}
			
			func (this *QRegion) OperatorNotEqual(r *QRegion) bool {
				return (bool)(QRegion_OperatorNotEqual(this.h, r.cPointer()))
}
			
			func (this *QRegion) ToHRGN() *struct HRGN__ {
				return (*struct HRGN__)(unsafe.Pointer(QRegion_ToHRGN(this.h)))
}
			
			func QRegion_FromHRGN(hrgn *struct HRGN__) *QRegion {
				_goptr := newQRegion(QRegion_FromHRGN(hrgn))
_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
return _goptr
}
			