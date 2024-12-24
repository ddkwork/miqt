package qt

import (
	"unsafe"
)

type QPalette__ColorGroup int

const (
	QPalette__Active       QPalette__ColorGroup = 0
	QPalette__Disabled     QPalette__ColorGroup = 1
	QPalette__Inactive     QPalette__ColorGroup = 2
	QPalette__NColorGroups QPalette__ColorGroup = 3
	QPalette__Current      QPalette__ColorGroup = 4
	QPalette__All          QPalette__ColorGroup = 5
	QPalette__Normal       QPalette__ColorGroup = 0
)

type QPalette__ColorRole int

const (
	QPalette__WindowText      QPalette__ColorRole = 0
	QPalette__Button          QPalette__ColorRole = 1
	QPalette__Light           QPalette__ColorRole = 2
	QPalette__Midlight        QPalette__ColorRole = 3
	QPalette__Dark            QPalette__ColorRole = 4
	QPalette__Mid             QPalette__ColorRole = 5
	QPalette__Text            QPalette__ColorRole = 6
	QPalette__BrightText      QPalette__ColorRole = 7
	QPalette__ButtonText      QPalette__ColorRole = 8
	QPalette__Base            QPalette__ColorRole = 9
	QPalette__Window          QPalette__ColorRole = 10
	QPalette__Shadow          QPalette__ColorRole = 11
	QPalette__Highlight       QPalette__ColorRole = 12
	QPalette__HighlightedText QPalette__ColorRole = 13
	QPalette__Link            QPalette__ColorRole = 14
	QPalette__LinkVisited     QPalette__ColorRole = 15
	QPalette__AlternateBase   QPalette__ColorRole = 16
	QPalette__NoRole          QPalette__ColorRole = 17
	QPalette__ToolTipBase     QPalette__ColorRole = 18
	QPalette__ToolTipText     QPalette__ColorRole = 19
	QPalette__PlaceholderText QPalette__ColorRole = 20
	QPalette__Accent          QPalette__ColorRole = 21
	QPalette__NColorRoles     QPalette__ColorRole = 22
)

type QPalette struct {
	h          uintptr
	isSubclass bool
}

// NewQPalette constructs a new QPalette object.
func NewQPalette() *QPalette {
	ret := newQPalette(QPalette_new())
	ret.isSubclass = true
	return ret
}

// NewQPalette2 constructs a new QPalette object.
func NewQPalette2(button *QColor) *QPalette {
	ret := newQPalette(QPalette_new2(button.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPalette3 constructs a new QPalette object.
func NewQPalette3(button GlobalColor) *QPalette {
	ret := newQPalette(QPalette_new3((int)(button)))
	ret.isSubclass = true
	return ret
}

// NewQPalette4 constructs a new QPalette object.
func NewQPalette4(button *QColor, window *QColor) *QPalette {
	ret := newQPalette(QPalette_new4(button.cPointer(), window.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPalette5 constructs a new QPalette object.
func NewQPalette5(windowText *QBrush, button *QBrush, light *QBrush, dark *QBrush, mid *QBrush, text *QBrush, bright_text *QBrush, base *QBrush, window *QBrush) *QPalette {
	ret := newQPalette(QPalette_new5(windowText.cPointer(), button.cPointer(), light.cPointer(), dark.cPointer(), mid.cPointer(), text.cPointer(), bright_text.cPointer(), base.cPointer(), window.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPalette6 constructs a new QPalette object.
func NewQPalette6(windowText *QColor, window *QColor, light *QColor, dark *QColor, mid *QColor, text *QColor, base *QColor) *QPalette {
	ret := newQPalette(QPalette_new6(windowText.cPointer(), window.cPointer(), light.cPointer(), dark.cPointer(), mid.cPointer(), text.cPointer(), base.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPalette7 constructs a new QPalette object.
func NewQPalette7(palette *QPalette) *QPalette {
	ret := newQPalette(QPalette_new7(palette.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPalette) OperatorAssign(palette *QPalette) {
	QPalette_OperatorAssign(this.h, palette.cPointer())
}

func (this *QPalette) Swap(other *QPalette) {
	QPalette_Swap(this.h, other.cPointer())
}

func (this *QPalette) CurrentColorGroup() ColorGroup {
	xxxxxxxxx
}

func (this *QPalette) SetCurrentColorGroup(cg ColorGroup) {
	QPalette_SetCurrentColorGroup(this.h, cg)
}

func (this *QPalette) Color(cg ColorGroup, cr ColorRole) *QColor {
	return newQColor(QPalette_Color(this.h, cg, cr))
}

func (this *QPalette) Brush(cg ColorGroup, cr ColorRole) *QBrush {
	return newQBrush(QPalette_Brush(this.h, cg, cr))
}

func (this *QPalette) SetColor(cg ColorGroup, cr ColorRole, color *QColor) {
	QPalette_SetColor(this.h, cg, cr, color.cPointer())
}

func (this *QPalette) SetColor2(cr ColorRole, color *QColor) {
	QPalette_SetColor2(this.h, cr, color.cPointer())
}

func (this *QPalette) SetBrush(cr ColorRole, brush *QBrush) {
	QPalette_SetBrush(this.h, cr, brush.cPointer())
}

func (this *QPalette) IsBrushSet(cg ColorGroup, cr ColorRole) bool {
	return (bool)(QPalette_IsBrushSet(this.h, cg, cr))
}

func (this *QPalette) SetBrush2(cg ColorGroup, cr ColorRole, brush *QBrush) {
	QPalette_SetBrush2(this.h, cg, cr, brush.cPointer())
}

func (this *QPalette) SetColorGroup(cr ColorGroup, windowText *QBrush, button *QBrush, light *QBrush, dark *QBrush, mid *QBrush, text *QBrush, bright_text *QBrush, base *QBrush, window *QBrush) {
	QPalette_SetColorGroup(this.h, cr, windowText.cPointer(), button.cPointer(), light.cPointer(), dark.cPointer(), mid.cPointer(), text.cPointer(), bright_text.cPointer(), base.cPointer(), window.cPointer())
}

func (this *QPalette) IsEqual(cr1 ColorGroup, cr2 ColorGroup) bool {
	return (bool)(QPalette_IsEqual(this.h, cr1, cr2))
}

func (this *QPalette) ColorWithCr(cr ColorRole) *QColor {
	return newQColor(QPalette_ColorWithCr(this.h, cr))
}

func (this *QPalette) BrushWithCr(cr ColorRole) *QBrush {
	return newQBrush(QPalette_BrushWithCr(this.h, cr))
}

func (this *QPalette) WindowText() *QBrush {
	return newQBrush(QPalette_WindowText(this.h))
}

func (this *QPalette) Button() *QBrush {
	return newQBrush(QPalette_Button(this.h))
}

func (this *QPalette) Light() *QBrush {
	return newQBrush(QPalette_Light(this.h))
}

func (this *QPalette) Dark() *QBrush {
	return newQBrush(QPalette_Dark(this.h))
}

func (this *QPalette) Mid() *QBrush {
	return newQBrush(QPalette_Mid(this.h))
}

func (this *QPalette) Text() *QBrush {
	return newQBrush(QPalette_Text(this.h))
}

func (this *QPalette) Base() *QBrush {
	return newQBrush(QPalette_Base(this.h))
}

func (this *QPalette) AlternateBase() *QBrush {
	return newQBrush(QPalette_AlternateBase(this.h))
}

func (this *QPalette) ToolTipBase() *QBrush {
	return newQBrush(QPalette_ToolTipBase(this.h))
}

func (this *QPalette) ToolTipText() *QBrush {
	return newQBrush(QPalette_ToolTipText(this.h))
}

func (this *QPalette) Window() *QBrush {
	return newQBrush(QPalette_Window(this.h))
}

func (this *QPalette) Midlight() *QBrush {
	return newQBrush(QPalette_Midlight(this.h))
}

func (this *QPalette) BrightText() *QBrush {
	return newQBrush(QPalette_BrightText(this.h))
}

func (this *QPalette) ButtonText() *QBrush {
	return newQBrush(QPalette_ButtonText(this.h))
}

func (this *QPalette) Shadow() *QBrush {
	return newQBrush(QPalette_Shadow(this.h))
}

func (this *QPalette) Highlight() *QBrush {
	return newQBrush(QPalette_Highlight(this.h))
}

func (this *QPalette) HighlightedText() *QBrush {
	return newQBrush(QPalette_HighlightedText(this.h))
}

func (this *QPalette) Link() *QBrush {
	return newQBrush(QPalette_Link(this.h))
}

func (this *QPalette) LinkVisited() *QBrush {
	return newQBrush(QPalette_LinkVisited(this.h))
}

func (this *QPalette) PlaceholderText() *QBrush {
	return newQBrush(QPalette_PlaceholderText(this.h))
}

func (this *QPalette) Accent() *QBrush {
	return newQBrush(QPalette_Accent(this.h))
}

func (this *QPalette) OperatorEqual(p *QPalette) bool {
	return (bool)(QPalette_OperatorEqual(this.h, p.cPointer()))
}

func (this *QPalette) OperatorNotEqual(p *QPalette) bool {
	return (bool)(QPalette_OperatorNotEqual(this.h, p.cPointer()))
}

func (this *QPalette) IsCopyOf(p *QPalette) bool {
	return (bool)(QPalette_IsCopyOf(this.h, p.cPointer()))
}

func (this *QPalette) CacheKey() int64 {
	return (int64)(QPalette_CacheKey(this.h))
}

func (this *QPalette) Resolve(other *QPalette) *QPalette {
	_goptr := newQPalette(QPalette_Resolve(this.h, other.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPalette) ResolveMask() ResolveMask {
	xxxxxxxxx
}

func (this *QPalette) SetResolveMask(mask ResolveMask) {
	QPalette_SetResolveMask(this.h, mask)
}
