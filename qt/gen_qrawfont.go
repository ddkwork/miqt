package qt

import (
	"unsafe"
)

type QRawFont__AntialiasingType int

const (
	QRawFont__PixelAntialiasing    QRawFont__AntialiasingType = 0
	QRawFont__SubPixelAntialiasing QRawFont__AntialiasingType = 1
)

type QRawFont__LayoutFlag int

const (
	QRawFont__SeparateAdvances QRawFont__LayoutFlag = 0
	QRawFont__KernedAdvances   QRawFont__LayoutFlag = 1
	QRawFont__UseDesignMetrics QRawFont__LayoutFlag = 2
)

type QRawFont struct {
	h          uintptr
	isSubclass bool
}

// NewQRawFont constructs a new QRawFont object.
func NewQRawFont() *QRawFont {

	ret := newQRawFont(QRawFont_new())
	ret.isSubclass = true
	return ret
}

// NewQRawFont2 constructs a new QRawFont object.
func NewQRawFont2(fileName string, pixelSize float64) *QRawFont {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQRawFont(QRawFont_new2(fileName_ms, (double)(pixelSize)))
	ret.isSubclass = true
	return ret
}

// NewQRawFont3 constructs a new QRawFont object.
func NewQRawFont3(fontData []byte, pixelSize float64) *QRawFont {
	fontData_alias := struct_miqt_string{}
	fontData_alias.data = (char)(unsafe.Pointer(&fontData[0]))
	fontData_alias.len = size_t(len(fontData))

	ret := newQRawFont(QRawFont_new3(fontData_alias, (double)(pixelSize)))
	ret.isSubclass = true
	return ret
}

// NewQRawFont4 constructs a new QRawFont object.
func NewQRawFont4(other *QRawFont) *QRawFont {

	ret := newQRawFont(QRawFont_new4(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQRawFont5 constructs a new QRawFont object.
func NewQRawFont5(fileName string, pixelSize float64, hintingPreference QFont__HintingPreference) *QRawFont {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQRawFont(QRawFont_new5(fileName_ms, (double)(pixelSize), (int)(hintingPreference)))
	ret.isSubclass = true
	return ret
}

// NewQRawFont6 constructs a new QRawFont object.
func NewQRawFont6(fontData []byte, pixelSize float64, hintingPreference QFont__HintingPreference) *QRawFont {
	fontData_alias := struct_miqt_string{}
	fontData_alias.data = (char)(unsafe.Pointer(&fontData[0]))
	fontData_alias.len = size_t(len(fontData))

	ret := newQRawFont(QRawFont_new6(fontData_alias, (double)(pixelSize), (int)(hintingPreference)))
	ret.isSubclass = true
	return ret
}

func (this *QRawFont) OperatorAssign(other *QRawFont) {
	QRawFont_OperatorAssign(this.h, other.cPointer())
}

func (this *QRawFont) Swap(other *QRawFont) {
	QRawFont_Swap(this.h, other.cPointer())
}

func (this *QRawFont) IsValid() bool {
	return (bool)(QRawFont_IsValid(this.h))
}

func (this *QRawFont) OperatorEqual(other *QRawFont) bool {
	return (bool)(QRawFont_OperatorEqual(this.h, other.cPointer()))
}

func (this *QRawFont) OperatorNotEqual(other *QRawFont) bool {
	return (bool)(QRawFont_OperatorNotEqual(this.h, other.cPointer()))
}

func (this *QRawFont) FamilyName() string {
	var _ms struct_miqt_string = QRawFont_FamilyName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRawFont) StyleName() string {
	var _ms struct_miqt_string = QRawFont_StyleName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRawFont) Style() QFont__Style {
	return (QFont__Style)(QRawFont_Style(this.h))
}

func (this *QRawFont) Weight() int {
	return (int)(QRawFont_Weight(this.h))
}

func (this *QRawFont) GlyphIndexesForString(text string) []uint {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ma struct_miqt_array = QRawFont_GlyphIndexesForString(this.h, text_ms)
	_ret := make([]uint, int(_ma.len))
	_outCast := (*[0xffff]uint)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (uint)(_outCast[i])
	}
	return _ret
}

func (this *QRawFont) AdvancesForGlyphIndexes(glyphIndexes []uint) []QPointF {
	glyphIndexes_CArray := (*[0xffff]uint)(malloc(size_t(8 * len(glyphIndexes))))
	defer free(unsafe.Pointer(glyphIndexes_CArray))
	for i := range glyphIndexes {
		glyphIndexes_CArray[i] = (uint)(glyphIndexes[i])
	}
	glyphIndexes_ma := struct_miqt_array{len: size_t(len(glyphIndexes)), data: unsafe.Pointer(glyphIndexes_CArray)}
	var _ma struct_miqt_array = QRawFont_AdvancesForGlyphIndexes(this.h, glyphIndexes_ma)
	_ret := make([]QPointF, int(_ma.len))
	_outCast := (*[0xffff]*QPointF)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQPointF(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QRawFont) AdvancesForGlyphIndexes2(glyphIndexes []uint, layoutFlags LayoutFlags) []QPointF {
	glyphIndexes_CArray := (*[0xffff]uint)(malloc(size_t(8 * len(glyphIndexes))))
	defer free(unsafe.Pointer(glyphIndexes_CArray))
	for i := range glyphIndexes {
		glyphIndexes_CArray[i] = (uint)(glyphIndexes[i])
	}
	glyphIndexes_ma := struct_miqt_array{len: size_t(len(glyphIndexes)), data: unsafe.Pointer(glyphIndexes_CArray)}
	var _ma struct_miqt_array = QRawFont_AdvancesForGlyphIndexes2(this.h, glyphIndexes_ma, layoutFlags)
	_ret := make([]QPointF, int(_ma.len))
	_outCast := (*[0xffff]*QPointF)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQPointF(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QRawFont) GlyphIndexesForChars(chars *QChar, numChars int, glyphIndexes *uint, numGlyphs *int) bool {
	return (bool)(QRawFont_GlyphIndexesForChars(this.h, chars.cPointer(), (int)(numChars), (*uint)(unsafe.Pointer(glyphIndexes)), (*int)(unsafe.Pointer(numGlyphs))))
}

func (this *QRawFont) AdvancesForGlyphIndexes3(glyphIndexes *uint, advances *QPointF, numGlyphs int) bool {
	return (bool)(QRawFont_AdvancesForGlyphIndexes3(this.h, (*uint)(unsafe.Pointer(glyphIndexes)), advances.cPointer(), (int)(numGlyphs)))
}

func (this *QRawFont) AdvancesForGlyphIndexes4(glyphIndexes *uint, advances *QPointF, numGlyphs int, layoutFlags LayoutFlags) bool {
	return (bool)(QRawFont_AdvancesForGlyphIndexes4(this.h, (*uint)(unsafe.Pointer(glyphIndexes)), advances.cPointer(), (int)(numGlyphs), layoutFlags))
}

func (this *QRawFont) AlphaMapForGlyph(glyphIndex uint) *QImage {
	_goptr := newQImage(QRawFont_AlphaMapForGlyph(this.h, (uint)(glyphIndex)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRawFont) PathForGlyph(glyphIndex uint) *QPainterPath {
	_goptr := newQPainterPath(QRawFont_PathForGlyph(this.h, (uint)(glyphIndex)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRawFont) BoundingRect(glyphIndex uint) *QRectF {
	_goptr := newQRectF(QRawFont_BoundingRect(this.h, (uint)(glyphIndex)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRawFont) SetPixelSize(pixelSize float64) {
	QRawFont_SetPixelSize(this.h, (double)(pixelSize))
}

func (this *QRawFont) PixelSize() float64 {
	return (float64)(QRawFont_PixelSize(this.h))
}

func (this *QRawFont) HintingPreference() QFont__HintingPreference {
	return (QFont__HintingPreference)(QRawFont_HintingPreference(this.h))
}

func (this *QRawFont) Ascent() float64 {
	return (float64)(QRawFont_Ascent(this.h))
}

func (this *QRawFont) CapHeight() float64 {
	return (float64)(QRawFont_CapHeight(this.h))
}

func (this *QRawFont) Descent() float64 {
	return (float64)(QRawFont_Descent(this.h))
}

func (this *QRawFont) Leading() float64 {
	return (float64)(QRawFont_Leading(this.h))
}

func (this *QRawFont) XHeight() float64 {
	return (float64)(QRawFont_XHeight(this.h))
}

func (this *QRawFont) AverageCharWidth() float64 {
	return (float64)(QRawFont_AverageCharWidth(this.h))
}

func (this *QRawFont) MaxCharWidth() float64 {
	return (float64)(QRawFont_MaxCharWidth(this.h))
}

func (this *QRawFont) LineThickness() float64 {
	return (float64)(QRawFont_LineThickness(this.h))
}

func (this *QRawFont) UnderlinePosition() float64 {
	return (float64)(QRawFont_UnderlinePosition(this.h))
}

func (this *QRawFont) UnitsPerEm() float64 {
	return (float64)(QRawFont_UnitsPerEm(this.h))
}

func (this *QRawFont) LoadFromFile(fileName string, pixelSize float64, hintingPreference QFont__HintingPreference) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QRawFont_LoadFromFile(this.h, fileName_ms, (double)(pixelSize), (int)(hintingPreference))
}

func (this *QRawFont) LoadFromData(fontData []byte, pixelSize float64, hintingPreference QFont__HintingPreference) {
	fontData_alias := struct_miqt_string{}
	fontData_alias.data = (char)(unsafe.Pointer(&fontData[0]))
	fontData_alias.len = size_t(len(fontData))
	QRawFont_LoadFromData(this.h, fontData_alias, (double)(pixelSize), (int)(hintingPreference))
}

func (this *QRawFont) SupportsCharacter(ucs4 uint) bool {
	return (bool)(QRawFont_SupportsCharacter(this.h, (uint)(ucs4)))
}

func (this *QRawFont) SupportsCharacterWithCharacter(character QChar) bool {
	return (bool)(QRawFont_SupportsCharacterWithCharacter(this.h, character.cPointer()))
}

func (this *QRawFont) SupportedWritingSystems() []QFontDatabase__WritingSystem {
	var _ma struct_miqt_array = QRawFont_SupportedWritingSystems(this.h)
	_ret := make([]QFontDatabase__WritingSystem, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (QFontDatabase__WritingSystem)(_outCast[i])
	}
	return _ret
}

func (this *QRawFont) FontTable(tagName string) []byte {
	tagName_Cstring := CString(tagName)
	defer free(unsafe.Pointer(tagName_Cstring))
	var _bytearray struct_miqt_string = QRawFont_FontTable(this.h, tagName_Cstring)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QRawFont) FontTableWithTag(tag QFont__Tag) []byte {
	var _bytearray struct_miqt_string = QRawFont_FontTableWithTag(this.h, tag.cPointer())
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QRawFont_FromFont(font *QFont) *QRawFont {
	_goptr := newQRawFont(QRawFont_FromFont(font.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRawFont) AlphaMapForGlyph2(glyphIndex uint, antialiasingType AntialiasingType) *QImage {
	_goptr := newQImage(QRawFont_AlphaMapForGlyph2(this.h, (uint)(glyphIndex), antialiasingType))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRawFont) AlphaMapForGlyph3(glyphIndex uint, antialiasingType AntialiasingType, transform *QTransform) *QImage {
	_goptr := newQImage(QRawFont_AlphaMapForGlyph3(this.h, (uint)(glyphIndex), antialiasingType, transform.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QRawFont_FromFont2(font *QFont, writingSystem QFontDatabase__WritingSystem) *QRawFont {
	_goptr := newQRawFont(QRawFont_FromFont2(font.cPointer(), (int)(writingSystem)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
