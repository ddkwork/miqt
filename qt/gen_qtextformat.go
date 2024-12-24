package qt

import (
	"unsafe"
)

type QTextLength__Type int

const (
	QTextLength__VariableLength   QTextLength__Type = 0
	QTextLength__FixedLength      QTextLength__Type = 1
	QTextLength__PercentageLength QTextLength__Type = 2
)

type QTextFormat__FormatType int

const (
	QTextFormat__InvalidFormat QTextFormat__FormatType = -1
	QTextFormat__BlockFormat   QTextFormat__FormatType = 1
	QTextFormat__CharFormat    QTextFormat__FormatType = 2
	QTextFormat__ListFormat    QTextFormat__FormatType = 3
	QTextFormat__FrameFormat   QTextFormat__FormatType = 5
	QTextFormat__UserFormat    QTextFormat__FormatType = 100
)

type QTextFormat__Property int

const (
	QTextFormat__ObjectIndex                       QTextFormat__Property = 0
	QTextFormat__CssFloat                          QTextFormat__Property = 2048
	QTextFormat__LayoutDirection                   QTextFormat__Property = 2049
	QTextFormat__OutlinePen                        QTextFormat__Property = 2064
	QTextFormat__BackgroundBrush                   QTextFormat__Property = 2080
	QTextFormat__ForegroundBrush                   QTextFormat__Property = 2081
	QTextFormat__BackgroundImageUrl                QTextFormat__Property = 2083
	QTextFormat__BlockAlignment                    QTextFormat__Property = 4112
	QTextFormat__BlockTopMargin                    QTextFormat__Property = 4144
	QTextFormat__BlockBottomMargin                 QTextFormat__Property = 4145
	QTextFormat__BlockLeftMargin                   QTextFormat__Property = 4146
	QTextFormat__BlockRightMargin                  QTextFormat__Property = 4147
	QTextFormat__TextIndent                        QTextFormat__Property = 4148
	QTextFormat__TabPositions                      QTextFormat__Property = 4149
	QTextFormat__BlockIndent                       QTextFormat__Property = 4160
	QTextFormat__LineHeight                        QTextFormat__Property = 4168
	QTextFormat__LineHeightType                    QTextFormat__Property = 4169
	QTextFormat__BlockNonBreakableLines            QTextFormat__Property = 4176
	QTextFormat__BlockTrailingHorizontalRulerWidth QTextFormat__Property = 4192
	QTextFormat__HeadingLevel                      QTextFormat__Property = 4208
	QTextFormat__BlockQuoteLevel                   QTextFormat__Property = 4224
	QTextFormat__BlockCodeLanguage                 QTextFormat__Property = 4240
	QTextFormat__BlockCodeFence                    QTextFormat__Property = 4241
	QTextFormat__BlockMarker                       QTextFormat__Property = 4256
	QTextFormat__FirstFontProperty                 QTextFormat__Property = 8160
	QTextFormat__FontCapitalization                QTextFormat__Property = 8160
	QTextFormat__FontLetterSpacing                 QTextFormat__Property = 8161
	QTextFormat__FontWordSpacing                   QTextFormat__Property = 8162
	QTextFormat__FontStyleHint                     QTextFormat__Property = 8163
	QTextFormat__FontStyleStrategy                 QTextFormat__Property = 8164
	QTextFormat__FontKerning                       QTextFormat__Property = 8165
	QTextFormat__FontHintingPreference             QTextFormat__Property = 8166
	QTextFormat__FontFamilies                      QTextFormat__Property = 8167
	QTextFormat__FontStyleName                     QTextFormat__Property = 8168
	QTextFormat__FontLetterSpacingType             QTextFormat__Property = 8169
	QTextFormat__FontStretch                       QTextFormat__Property = 8170
	QTextFormat__FontFamily                        QTextFormat__Property = 8192
	QTextFormat__FontPointSize                     QTextFormat__Property = 8193
	QTextFormat__FontSizeAdjustment                QTextFormat__Property = 8194
	QTextFormat__FontSizeIncrement                 QTextFormat__Property = 8194
	QTextFormat__FontWeight                        QTextFormat__Property = 8195
	QTextFormat__FontItalic                        QTextFormat__Property = 8196
	QTextFormat__FontUnderline                     QTextFormat__Property = 8197
	QTextFormat__FontOverline                      QTextFormat__Property = 8198
	QTextFormat__FontStrikeOut                     QTextFormat__Property = 8199
	QTextFormat__FontFixedPitch                    QTextFormat__Property = 8200
	QTextFormat__FontPixelSize                     QTextFormat__Property = 8201
	QTextFormat__LastFontProperty                  QTextFormat__Property = 8201
	QTextFormat__TextUnderlineColor                QTextFormat__Property = 8224
	QTextFormat__TextVerticalAlignment             QTextFormat__Property = 8225
	QTextFormat__TextOutline                       QTextFormat__Property = 8226
	QTextFormat__TextUnderlineStyle                QTextFormat__Property = 8227
	QTextFormat__TextToolTip                       QTextFormat__Property = 8228
	QTextFormat__TextSuperScriptBaseline           QTextFormat__Property = 8229
	QTextFormat__TextSubScriptBaseline             QTextFormat__Property = 8230
	QTextFormat__TextBaselineOffset                QTextFormat__Property = 8231
	QTextFormat__IsAnchor                          QTextFormat__Property = 8240
	QTextFormat__AnchorHref                        QTextFormat__Property = 8241
	QTextFormat__AnchorName                        QTextFormat__Property = 8242
	QTextFormat__OldFontLetterSpacingType          QTextFormat__Property = 8243
	QTextFormat__OldFontStretch                    QTextFormat__Property = 8244
	QTextFormat__OldTextUnderlineColor             QTextFormat__Property = 8208
	QTextFormat__OldFontFamily                     QTextFormat__Property = 8192
	QTextFormat__ObjectType                        QTextFormat__Property = 12032
	QTextFormat__ListStyle                         QTextFormat__Property = 12288
	QTextFormat__ListIndent                        QTextFormat__Property = 12289
	QTextFormat__ListNumberPrefix                  QTextFormat__Property = 12290
	QTextFormat__ListNumberSuffix                  QTextFormat__Property = 12291
	QTextFormat__ListStart                         QTextFormat__Property = 12292
	QTextFormat__FrameBorder                       QTextFormat__Property = 16384
	QTextFormat__FrameMargin                       QTextFormat__Property = 16385
	QTextFormat__FramePadding                      QTextFormat__Property = 16386
	QTextFormat__FrameWidth                        QTextFormat__Property = 16387
	QTextFormat__FrameHeight                       QTextFormat__Property = 16388
	QTextFormat__FrameTopMargin                    QTextFormat__Property = 16389
	QTextFormat__FrameBottomMargin                 QTextFormat__Property = 16390
	QTextFormat__FrameLeftMargin                   QTextFormat__Property = 16391
	QTextFormat__FrameRightMargin                  QTextFormat__Property = 16392
	QTextFormat__FrameBorderBrush                  QTextFormat__Property = 16393
	QTextFormat__FrameBorderStyle                  QTextFormat__Property = 16400
	QTextFormat__TableColumns                      QTextFormat__Property = 16640
	QTextFormat__TableColumnWidthConstraints       QTextFormat__Property = 16641
	QTextFormat__TableCellSpacing                  QTextFormat__Property = 16642
	QTextFormat__TableCellPadding                  QTextFormat__Property = 16643
	QTextFormat__TableHeaderRowCount               QTextFormat__Property = 16644
	QTextFormat__TableBorderCollapse               QTextFormat__Property = 16645
	QTextFormat__TableCellRowSpan                  QTextFormat__Property = 18448
	QTextFormat__TableCellColumnSpan               QTextFormat__Property = 18449
	QTextFormat__TableCellTopPadding               QTextFormat__Property = 18450
	QTextFormat__TableCellBottomPadding            QTextFormat__Property = 18451
	QTextFormat__TableCellLeftPadding              QTextFormat__Property = 18452
	QTextFormat__TableCellRightPadding             QTextFormat__Property = 18453
	QTextFormat__TableCellTopBorder                QTextFormat__Property = 18454
	QTextFormat__TableCellBottomBorder             QTextFormat__Property = 18455
	QTextFormat__TableCellLeftBorder               QTextFormat__Property = 18456
	QTextFormat__TableCellRightBorder              QTextFormat__Property = 18457
	QTextFormat__TableCellTopBorderStyle           QTextFormat__Property = 18458
	QTextFormat__TableCellBottomBorderStyle        QTextFormat__Property = 18459
	QTextFormat__TableCellLeftBorderStyle          QTextFormat__Property = 18460
	QTextFormat__TableCellRightBorderStyle         QTextFormat__Property = 18461
	QTextFormat__TableCellTopBorderBrush           QTextFormat__Property = 18462
	QTextFormat__TableCellBottomBorderBrush        QTextFormat__Property = 18463
	QTextFormat__TableCellLeftBorderBrush          QTextFormat__Property = 18464
	QTextFormat__TableCellRightBorderBrush         QTextFormat__Property = 18465
	QTextFormat__ImageName                         QTextFormat__Property = 20480
	QTextFormat__ImageTitle                        QTextFormat__Property = 20481
	QTextFormat__ImageAltText                      QTextFormat__Property = 20482
	QTextFormat__ImageWidth                        QTextFormat__Property = 20496
	QTextFormat__ImageHeight                       QTextFormat__Property = 20497
	QTextFormat__ImageQuality                      QTextFormat__Property = 20500
	QTextFormat__ImageMaxWidth                     QTextFormat__Property = 20501
	QTextFormat__FullWidthSelection                QTextFormat__Property = 24576
	QTextFormat__PageBreakPolicy                   QTextFormat__Property = 28672
	QTextFormat__UserProperty                      QTextFormat__Property = 1048576
)

type QTextFormat__ObjectTypes int

const (
	QTextFormat__NoObject        QTextFormat__ObjectTypes = 0
	QTextFormat__ImageObject     QTextFormat__ObjectTypes = 1
	QTextFormat__TableObject     QTextFormat__ObjectTypes = 2
	QTextFormat__TableCellObject QTextFormat__ObjectTypes = 3
	QTextFormat__UserObject      QTextFormat__ObjectTypes = 4096
)

type QTextFormat__PageBreakFlag int

const (
	QTextFormat__PageBreak_Auto         QTextFormat__PageBreakFlag = 0
	QTextFormat__PageBreak_AlwaysBefore QTextFormat__PageBreakFlag = 1
	QTextFormat__PageBreak_AlwaysAfter  QTextFormat__PageBreakFlag = 16
)

type QTextCharFormat__VerticalAlignment int

const (
	QTextCharFormat__AlignNormal      QTextCharFormat__VerticalAlignment = 0
	QTextCharFormat__AlignSuperScript QTextCharFormat__VerticalAlignment = 1
	QTextCharFormat__AlignSubScript   QTextCharFormat__VerticalAlignment = 2
	QTextCharFormat__AlignMiddle      QTextCharFormat__VerticalAlignment = 3
	QTextCharFormat__AlignTop         QTextCharFormat__VerticalAlignment = 4
	QTextCharFormat__AlignBottom      QTextCharFormat__VerticalAlignment = 5
	QTextCharFormat__AlignBaseline    QTextCharFormat__VerticalAlignment = 6
)

type QTextCharFormat__UnderlineStyle int

const (
	QTextCharFormat__NoUnderline         QTextCharFormat__UnderlineStyle = 0
	QTextCharFormat__SingleUnderline     QTextCharFormat__UnderlineStyle = 1
	QTextCharFormat__DashUnderline       QTextCharFormat__UnderlineStyle = 2
	QTextCharFormat__DotLine             QTextCharFormat__UnderlineStyle = 3
	QTextCharFormat__DashDotLine         QTextCharFormat__UnderlineStyle = 4
	QTextCharFormat__DashDotDotLine      QTextCharFormat__UnderlineStyle = 5
	QTextCharFormat__WaveUnderline       QTextCharFormat__UnderlineStyle = 6
	QTextCharFormat__SpellCheckUnderline QTextCharFormat__UnderlineStyle = 7
)

type QTextCharFormat__FontPropertiesInheritanceBehavior int

const (
	QTextCharFormat__FontPropertiesSpecifiedOnly QTextCharFormat__FontPropertiesInheritanceBehavior = 0
	QTextCharFormat__FontPropertiesAll           QTextCharFormat__FontPropertiesInheritanceBehavior = 1
)

type QTextBlockFormat__LineHeightTypes int

const (
	QTextBlockFormat__SingleHeight       QTextBlockFormat__LineHeightTypes = 0
	QTextBlockFormat__ProportionalHeight QTextBlockFormat__LineHeightTypes = 1
	QTextBlockFormat__FixedHeight        QTextBlockFormat__LineHeightTypes = 2
	QTextBlockFormat__MinimumHeight      QTextBlockFormat__LineHeightTypes = 3
	QTextBlockFormat__LineDistanceHeight QTextBlockFormat__LineHeightTypes = 4
)

type QTextBlockFormat__MarkerType int

const (
	QTextBlockFormat__NoMarker  QTextBlockFormat__MarkerType = 0
	QTextBlockFormat__Unchecked QTextBlockFormat__MarkerType = 1
	QTextBlockFormat__Checked   QTextBlockFormat__MarkerType = 2
)

type QTextListFormat__Style int

const (
	QTextListFormat__ListDisc           QTextListFormat__Style = -1
	QTextListFormat__ListCircle         QTextListFormat__Style = -2
	QTextListFormat__ListSquare         QTextListFormat__Style = -3
	QTextListFormat__ListDecimal        QTextListFormat__Style = -4
	QTextListFormat__ListLowerAlpha     QTextListFormat__Style = -5
	QTextListFormat__ListUpperAlpha     QTextListFormat__Style = -6
	QTextListFormat__ListLowerRoman     QTextListFormat__Style = -7
	QTextListFormat__ListUpperRoman     QTextListFormat__Style = -8
	QTextListFormat__ListStyleUndefined QTextListFormat__Style = 0
)

type QTextFrameFormat__Position int

const (
	QTextFrameFormat__InFlow     QTextFrameFormat__Position = 0
	QTextFrameFormat__FloatLeft  QTextFrameFormat__Position = 1
	QTextFrameFormat__FloatRight QTextFrameFormat__Position = 2
)

type QTextFrameFormat__BorderStyle int

const (
	QTextFrameFormat__BorderStyle_None       QTextFrameFormat__BorderStyle = 0
	QTextFrameFormat__BorderStyle_Dotted     QTextFrameFormat__BorderStyle = 1
	QTextFrameFormat__BorderStyle_Dashed     QTextFrameFormat__BorderStyle = 2
	QTextFrameFormat__BorderStyle_Solid      QTextFrameFormat__BorderStyle = 3
	QTextFrameFormat__BorderStyle_Double     QTextFrameFormat__BorderStyle = 4
	QTextFrameFormat__BorderStyle_DotDash    QTextFrameFormat__BorderStyle = 5
	QTextFrameFormat__BorderStyle_DotDotDash QTextFrameFormat__BorderStyle = 6
	QTextFrameFormat__BorderStyle_Groove     QTextFrameFormat__BorderStyle = 7
	QTextFrameFormat__BorderStyle_Ridge      QTextFrameFormat__BorderStyle = 8
	QTextFrameFormat__BorderStyle_Inset      QTextFrameFormat__BorderStyle = 9
	QTextFrameFormat__BorderStyle_Outset     QTextFrameFormat__BorderStyle = 10
)

type QTextLength struct {
	h          uintptr
	isSubclass bool
}

// NewQTextLength constructs a new QTextLength object.
func NewQTextLength() *QTextLength {

	ret := newQTextLength(QTextLength_new())
	ret.isSubclass = true
	return ret
}

// NewQTextLength2 constructs a new QTextLength object.
func NewQTextLength2(typeVal Type, value float64) *QTextLength {

	ret := newQTextLength(QTextLength_new2(typeVal, (double)(value)))
	ret.isSubclass = true
	return ret
}

// NewQTextLength3 constructs a new QTextLength object.
func NewQTextLength3(param1 *QTextLength) *QTextLength {

	ret := newQTextLength(QTextLength_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextLength) Type() Type {
	xxxxxxxxx
}

func (this *QTextLength) Value(maximumLength float64) float64 {
	return (float64)(QTextLength_Value(this.h, (double)(maximumLength)))
}

func (this *QTextLength) RawValue() float64 {
	return (float64)(QTextLength_RawValue(this.h))
}

func (this *QTextLength) OperatorEqual(other *QTextLength) bool {
	return (bool)(QTextLength_OperatorEqual(this.h, other.cPointer()))
}

func (this *QTextLength) OperatorNotEqual(other *QTextLength) bool {
	return (bool)(QTextLength_OperatorNotEqual(this.h, other.cPointer()))
}

type QTextFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQTextFormat constructs a new QTextFormat object.
func NewQTextFormat() *QTextFormat {

	ret := newQTextFormat(QTextFormat_new())
	ret.isSubclass = true
	return ret
}

// NewQTextFormat2 constructs a new QTextFormat object.
func NewQTextFormat2(typeVal int) *QTextFormat {

	ret := newQTextFormat(QTextFormat_new2((int)(typeVal)))
	ret.isSubclass = true
	return ret
}

// NewQTextFormat3 constructs a new QTextFormat object.
func NewQTextFormat3(rhs *QTextFormat) *QTextFormat {

	ret := newQTextFormat(QTextFormat_new3(rhs.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextFormat) OperatorAssign(rhs *QTextFormat) {
	QTextFormat_OperatorAssign(this.h, rhs.cPointer())
}

func (this *QTextFormat) Swap(other *QTextFormat) {
	QTextFormat_Swap(this.h, other.cPointer())
}

func (this *QTextFormat) Merge(other *QTextFormat) {
	QTextFormat_Merge(this.h, other.cPointer())
}

func (this *QTextFormat) IsValid() bool {
	return (bool)(QTextFormat_IsValid(this.h))
}

func (this *QTextFormat) IsEmpty() bool {
	return (bool)(QTextFormat_IsEmpty(this.h))
}

func (this *QTextFormat) Type() int {
	return (int)(QTextFormat_Type(this.h))
}

func (this *QTextFormat) ObjectIndex() int {
	return (int)(QTextFormat_ObjectIndex(this.h))
}

func (this *QTextFormat) SetObjectIndex(object int) {
	QTextFormat_SetObjectIndex(this.h, (int)(object))
}

func (this *QTextFormat) Property(propertyId int) *QVariant {
	_goptr := newQVariant(QTextFormat_Property(this.h, (int)(propertyId)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) SetProperty(propertyId int, value *QVariant) {
	QTextFormat_SetProperty(this.h, (int)(propertyId), value.cPointer())
}

func (this *QTextFormat) ClearProperty(propertyId int) {
	QTextFormat_ClearProperty(this.h, (int)(propertyId))
}

func (this *QTextFormat) HasProperty(propertyId int) bool {
	return (bool)(QTextFormat_HasProperty(this.h, (int)(propertyId)))
}

func (this *QTextFormat) BoolProperty(propertyId int) bool {
	return (bool)(QTextFormat_BoolProperty(this.h, (int)(propertyId)))
}

func (this *QTextFormat) IntProperty(propertyId int) int {
	return (int)(QTextFormat_IntProperty(this.h, (int)(propertyId)))
}

func (this *QTextFormat) DoubleProperty(propertyId int) float64 {
	return (float64)(QTextFormat_DoubleProperty(this.h, (int)(propertyId)))
}

func (this *QTextFormat) StringProperty(propertyId int) string {
	var _ms struct_miqt_string = QTextFormat_StringProperty(this.h, (int)(propertyId))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextFormat) ColorProperty(propertyId int) *QColor {
	_goptr := newQColor(QTextFormat_ColorProperty(this.h, (int)(propertyId)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) PenProperty(propertyId int) *QPen {
	_goptr := newQPen(QTextFormat_PenProperty(this.h, (int)(propertyId)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) BrushProperty(propertyId int) *QBrush {
	_goptr := newQBrush(QTextFormat_BrushProperty(this.h, (int)(propertyId)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) LengthProperty(propertyId int) *QTextLength {
	_goptr := newQTextLength(QTextFormat_LengthProperty(this.h, (int)(propertyId)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) LengthVectorProperty(propertyId int) []QTextLength {
	var _ma struct_miqt_array = QTextFormat_LengthVectorProperty(this.h, (int)(propertyId))
	_ret := make([]QTextLength, int(_ma.len))
	_outCast := (*[0xffff]*QTextLength)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQTextLength(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextFormat) SetProperty2(propertyId int, lengths []QTextLength) {
	lengths_CArray := (*[0xffff]*QTextLength)(malloc(size_t(8 * len(lengths))))
	defer free(unsafe.Pointer(lengths_CArray))
	for i := range lengths {
		lengths_CArray[i] = lengths[i].cPointer()
	}
	lengths_ma := struct_miqt_array{len: size_t(len(lengths)), data: unsafe.Pointer(lengths_CArray)}
	QTextFormat_SetProperty2(this.h, (int)(propertyId), lengths_ma)
}

func (this *QTextFormat) Properties() map[int]QVariant {
	var _mm struct_miqt_map = QTextFormat_Properties(this.h)
	_ret := make(map[int]QVariant, int(_mm.len))
	_Keys := (*[0xffff]int)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]*QVariant)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		_entry_Key := (int)(_Keys[i])

		_mapval_goptr := newQVariant(_Values[i])
		_mapval_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_entry_Value := *_mapval_goptr

		_ret[_entry_Key] = _entry_Value
	}
	return _ret
}

func (this *QTextFormat) PropertyCount() int {
	return (int)(QTextFormat_PropertyCount(this.h))
}

func (this *QTextFormat) SetObjectType(typeVal int) {
	QTextFormat_SetObjectType(this.h, (int)(typeVal))
}

func (this *QTextFormat) ObjectType() int {
	return (int)(QTextFormat_ObjectType(this.h))
}

func (this *QTextFormat) IsCharFormat() bool {
	return (bool)(QTextFormat_IsCharFormat(this.h))
}

func (this *QTextFormat) IsBlockFormat() bool {
	return (bool)(QTextFormat_IsBlockFormat(this.h))
}

func (this *QTextFormat) IsListFormat() bool {
	return (bool)(QTextFormat_IsListFormat(this.h))
}

func (this *QTextFormat) IsFrameFormat() bool {
	return (bool)(QTextFormat_IsFrameFormat(this.h))
}

func (this *QTextFormat) IsImageFormat() bool {
	return (bool)(QTextFormat_IsImageFormat(this.h))
}

func (this *QTextFormat) IsTableFormat() bool {
	return (bool)(QTextFormat_IsTableFormat(this.h))
}

func (this *QTextFormat) IsTableCellFormat() bool {
	return (bool)(QTextFormat_IsTableCellFormat(this.h))
}

func (this *QTextFormat) ToBlockFormat() *QTextBlockFormat {
	_goptr := newQTextBlockFormat(QTextFormat_ToBlockFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) ToCharFormat() *QTextCharFormat {
	_goptr := newQTextCharFormat(QTextFormat_ToCharFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) ToListFormat() *QTextListFormat {
	_goptr := newQTextListFormat(QTextFormat_ToListFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) ToTableFormat() *QTextTableFormat {
	_goptr := newQTextTableFormat(QTextFormat_ToTableFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) ToFrameFormat() *QTextFrameFormat {
	_goptr := newQTextFrameFormat(QTextFormat_ToFrameFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) ToImageFormat() *QTextImageFormat {
	_goptr := newQTextImageFormat(QTextFormat_ToImageFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) ToTableCellFormat() *QTextTableCellFormat {
	_goptr := newQTextTableCellFormat(QTextFormat_ToTableCellFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) OperatorEqual(rhs *QTextFormat) bool {
	return (bool)(QTextFormat_OperatorEqual(this.h, rhs.cPointer()))
}

func (this *QTextFormat) OperatorNotEqual(rhs *QTextFormat) bool {
	return (bool)(QTextFormat_OperatorNotEqual(this.h, rhs.cPointer()))
}

func (this *QTextFormat) SetLayoutDirection(direction LayoutDirection) {
	QTextFormat_SetLayoutDirection(this.h, (int)(direction))
}

func (this *QTextFormat) LayoutDirection() LayoutDirection {
	return (LayoutDirection)(QTextFormat_LayoutDirection(this.h))
}

func (this *QTextFormat) SetBackground(brush *QBrush) {
	QTextFormat_SetBackground(this.h, brush.cPointer())
}

func (this *QTextFormat) Background() *QBrush {
	_goptr := newQBrush(QTextFormat_Background(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) ClearBackground() {
	QTextFormat_ClearBackground(this.h)
}

func (this *QTextFormat) SetForeground(brush *QBrush) {
	QTextFormat_SetForeground(this.h, brush.cPointer())
}

func (this *QTextFormat) Foreground() *QBrush {
	_goptr := newQBrush(QTextFormat_Foreground(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFormat) ClearForeground() {
	QTextFormat_ClearForeground(this.h)
}

type QTextCharFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQTextCharFormat constructs a new QTextCharFormat object.
func NewQTextCharFormat() *QTextCharFormat {

	ret := newQTextCharFormat(QTextCharFormat_new())
	ret.isSubclass = true
	return ret
}

func (this *QTextCharFormat) IsValid() bool {
	return (bool)(QTextCharFormat_IsValid(this.h))
}

func (this *QTextCharFormat) SetFont(font *QFont) {
	QTextCharFormat_SetFont(this.h, font.cPointer())
}

func (this *QTextCharFormat) Font() *QFont {
	_goptr := newQFont(QTextCharFormat_Font(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextCharFormat) SetFontFamily(family string) {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	QTextCharFormat_SetFontFamily(this.h, family_ms)
}

func (this *QTextCharFormat) FontFamily() string {
	var _ms struct_miqt_string = QTextCharFormat_FontFamily(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextCharFormat) SetFontFamilies(families []string) {
	families_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(families))))
	defer free(unsafe.Pointer(families_CArray))
	for i := range families {
		families_i_ms := struct_miqt_string{}
		families_i_ms.data = CString(families[i])
		families_i_ms.len = size_t(len(families[i]))
		defer free(unsafe.Pointer(families_i_ms.data))
		families_CArray[i] = families_i_ms
	}
	families_ma := struct_miqt_array{len: size_t(len(families)), data: unsafe.Pointer(families_CArray)}
	QTextCharFormat_SetFontFamilies(this.h, families_ma)
}

func (this *QTextCharFormat) FontFamilies() *QVariant {
	_goptr := newQVariant(QTextCharFormat_FontFamilies(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextCharFormat) SetFontStyleName(styleName string) {
	styleName_ms := struct_miqt_string{}
	styleName_ms.data = CString(styleName)
	styleName_ms.len = size_t(len(styleName))
	defer free(unsafe.Pointer(styleName_ms.data))
	QTextCharFormat_SetFontStyleName(this.h, styleName_ms)
}

func (this *QTextCharFormat) FontStyleName() *QVariant {
	_goptr := newQVariant(QTextCharFormat_FontStyleName(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextCharFormat) SetFontPointSize(size float64) {
	QTextCharFormat_SetFontPointSize(this.h, (double)(size))
}

func (this *QTextCharFormat) FontPointSize() float64 {
	return (float64)(QTextCharFormat_FontPointSize(this.h))
}

func (this *QTextCharFormat) SetFontWeight(weight int) {
	QTextCharFormat_SetFontWeight(this.h, (int)(weight))
}

func (this *QTextCharFormat) FontWeight() int {
	return (int)(QTextCharFormat_FontWeight(this.h))
}

func (this *QTextCharFormat) SetFontItalic(italic bool) {
	QTextCharFormat_SetFontItalic(this.h, (bool)(italic))
}

func (this *QTextCharFormat) FontItalic() bool {
	return (bool)(QTextCharFormat_FontItalic(this.h))
}

func (this *QTextCharFormat) SetFontCapitalization(capitalization QFont__Capitalization) {
	QTextCharFormat_SetFontCapitalization(this.h, (int)(capitalization))
}

func (this *QTextCharFormat) FontCapitalization() QFont__Capitalization {
	return (QFont__Capitalization)(QTextCharFormat_FontCapitalization(this.h))
}

func (this *QTextCharFormat) SetFontLetterSpacingType(letterSpacingType QFont__SpacingType) {
	QTextCharFormat_SetFontLetterSpacingType(this.h, (int)(letterSpacingType))
}

func (this *QTextCharFormat) FontLetterSpacingType() QFont__SpacingType {
	return (QFont__SpacingType)(QTextCharFormat_FontLetterSpacingType(this.h))
}

func (this *QTextCharFormat) SetFontLetterSpacing(spacing float64) {
	QTextCharFormat_SetFontLetterSpacing(this.h, (double)(spacing))
}

func (this *QTextCharFormat) FontLetterSpacing() float64 {
	return (float64)(QTextCharFormat_FontLetterSpacing(this.h))
}

func (this *QTextCharFormat) SetFontWordSpacing(spacing float64) {
	QTextCharFormat_SetFontWordSpacing(this.h, (double)(spacing))
}

func (this *QTextCharFormat) FontWordSpacing() float64 {
	return (float64)(QTextCharFormat_FontWordSpacing(this.h))
}

func (this *QTextCharFormat) SetFontUnderline(underline bool) {
	QTextCharFormat_SetFontUnderline(this.h, (bool)(underline))
}

func (this *QTextCharFormat) FontUnderline() bool {
	return (bool)(QTextCharFormat_FontUnderline(this.h))
}

func (this *QTextCharFormat) SetFontOverline(overline bool) {
	QTextCharFormat_SetFontOverline(this.h, (bool)(overline))
}

func (this *QTextCharFormat) FontOverline() bool {
	return (bool)(QTextCharFormat_FontOverline(this.h))
}

func (this *QTextCharFormat) SetFontStrikeOut(strikeOut bool) {
	QTextCharFormat_SetFontStrikeOut(this.h, (bool)(strikeOut))
}

func (this *QTextCharFormat) FontStrikeOut() bool {
	return (bool)(QTextCharFormat_FontStrikeOut(this.h))
}

func (this *QTextCharFormat) SetUnderlineColor(color *QColor) {
	QTextCharFormat_SetUnderlineColor(this.h, color.cPointer())
}

func (this *QTextCharFormat) UnderlineColor() *QColor {
	_goptr := newQColor(QTextCharFormat_UnderlineColor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextCharFormat) SetFontFixedPitch(fixedPitch bool) {
	QTextCharFormat_SetFontFixedPitch(this.h, (bool)(fixedPitch))
}

func (this *QTextCharFormat) FontFixedPitch() bool {
	return (bool)(QTextCharFormat_FontFixedPitch(this.h))
}

func (this *QTextCharFormat) SetFontStretch(factor int) {
	QTextCharFormat_SetFontStretch(this.h, (int)(factor))
}

func (this *QTextCharFormat) FontStretch() int {
	return (int)(QTextCharFormat_FontStretch(this.h))
}

func (this *QTextCharFormat) SetFontStyleHint(hint QFont__StyleHint) {
	QTextCharFormat_SetFontStyleHint(this.h, (int)(hint))
}

func (this *QTextCharFormat) SetFontStyleStrategy(strategy QFont__StyleStrategy) {
	QTextCharFormat_SetFontStyleStrategy(this.h, (int)(strategy))
}

func (this *QTextCharFormat) FontStyleHint() QFont__StyleHint {
	return (QFont__StyleHint)(QTextCharFormat_FontStyleHint(this.h))
}

func (this *QTextCharFormat) FontStyleStrategy() QFont__StyleStrategy {
	return (QFont__StyleStrategy)(QTextCharFormat_FontStyleStrategy(this.h))
}

func (this *QTextCharFormat) SetFontHintingPreference(hintingPreference QFont__HintingPreference) {
	QTextCharFormat_SetFontHintingPreference(this.h, (int)(hintingPreference))
}

func (this *QTextCharFormat) FontHintingPreference() QFont__HintingPreference {
	return (QFont__HintingPreference)(QTextCharFormat_FontHintingPreference(this.h))
}

func (this *QTextCharFormat) SetFontKerning(enable bool) {
	QTextCharFormat_SetFontKerning(this.h, (bool)(enable))
}

func (this *QTextCharFormat) FontKerning() bool {
	return (bool)(QTextCharFormat_FontKerning(this.h))
}

func (this *QTextCharFormat) SetUnderlineStyle(style UnderlineStyle) {
	QTextCharFormat_SetUnderlineStyle(this.h, style)
}

func (this *QTextCharFormat) UnderlineStyle() UnderlineStyle {
	xxxxxxxxx
}

func (this *QTextCharFormat) SetVerticalAlignment(alignment VerticalAlignment) {
	QTextCharFormat_SetVerticalAlignment(this.h, alignment)
}

func (this *QTextCharFormat) VerticalAlignment() VerticalAlignment {
	xxxxxxxxx
}

func (this *QTextCharFormat) SetTextOutline(pen *QPen) {
	QTextCharFormat_SetTextOutline(this.h, pen.cPointer())
}

func (this *QTextCharFormat) TextOutline() *QPen {
	_goptr := newQPen(QTextCharFormat_TextOutline(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextCharFormat) SetToolTip(tip string) {
	tip_ms := struct_miqt_string{}
	tip_ms.data = CString(tip)
	tip_ms.len = size_t(len(tip))
	defer free(unsafe.Pointer(tip_ms.data))
	QTextCharFormat_SetToolTip(this.h, tip_ms)
}

func (this *QTextCharFormat) ToolTip() string {
	var _ms struct_miqt_string = QTextCharFormat_ToolTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextCharFormat) SetSuperScriptBaseline(baseline float64) {
	QTextCharFormat_SetSuperScriptBaseline(this.h, (double)(baseline))
}

func (this *QTextCharFormat) SuperScriptBaseline() float64 {
	return (float64)(QTextCharFormat_SuperScriptBaseline(this.h))
}

func (this *QTextCharFormat) SetSubScriptBaseline(baseline float64) {
	QTextCharFormat_SetSubScriptBaseline(this.h, (double)(baseline))
}

func (this *QTextCharFormat) SubScriptBaseline() float64 {
	return (float64)(QTextCharFormat_SubScriptBaseline(this.h))
}

func (this *QTextCharFormat) SetBaselineOffset(baseline float64) {
	QTextCharFormat_SetBaselineOffset(this.h, (double)(baseline))
}

func (this *QTextCharFormat) BaselineOffset() float64 {
	return (float64)(QTextCharFormat_BaselineOffset(this.h))
}

func (this *QTextCharFormat) SetAnchor(anchor bool) {
	QTextCharFormat_SetAnchor(this.h, (bool)(anchor))
}

func (this *QTextCharFormat) IsAnchor() bool {
	return (bool)(QTextCharFormat_IsAnchor(this.h))
}

func (this *QTextCharFormat) SetAnchorHref(value string) {
	value_ms := struct_miqt_string{}
	value_ms.data = CString(value)
	value_ms.len = size_t(len(value))
	defer free(unsafe.Pointer(value_ms.data))
	QTextCharFormat_SetAnchorHref(this.h, value_ms)
}

func (this *QTextCharFormat) AnchorHref() string {
	var _ms struct_miqt_string = QTextCharFormat_AnchorHref(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextCharFormat) SetAnchorNames(names []string) {
	names_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(names))))
	defer free(unsafe.Pointer(names_CArray))
	for i := range names {
		names_i_ms := struct_miqt_string{}
		names_i_ms.data = CString(names[i])
		names_i_ms.len = size_t(len(names[i]))
		defer free(unsafe.Pointer(names_i_ms.data))
		names_CArray[i] = names_i_ms
	}
	names_ma := struct_miqt_array{len: size_t(len(names)), data: unsafe.Pointer(names_CArray)}
	QTextCharFormat_SetAnchorNames(this.h, names_ma)
}

func (this *QTextCharFormat) AnchorNames() []string {
	var _ma struct_miqt_array = QTextCharFormat_AnchorNames(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QTextCharFormat) SetTableCellRowSpan(tableCellRowSpan int) {
	QTextCharFormat_SetTableCellRowSpan(this.h, (int)(tableCellRowSpan))
}

func (this *QTextCharFormat) TableCellRowSpan() int {
	return (int)(QTextCharFormat_TableCellRowSpan(this.h))
}

func (this *QTextCharFormat) SetTableCellColumnSpan(tableCellColumnSpan int) {
	QTextCharFormat_SetTableCellColumnSpan(this.h, (int)(tableCellColumnSpan))
}

func (this *QTextCharFormat) TableCellColumnSpan() int {
	return (int)(QTextCharFormat_TableCellColumnSpan(this.h))
}

func (this *QTextCharFormat) SetFont2(font *QFont, behavior FontPropertiesInheritanceBehavior) {
	QTextCharFormat_SetFont2(this.h, font.cPointer(), behavior)
}

func (this *QTextCharFormat) SetFontStyleHint2(hint QFont__StyleHint, strategy QFont__StyleStrategy) {
	QTextCharFormat_SetFontStyleHint2(this.h, (int)(hint), (int)(strategy))
}

type QTextBlockFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQTextBlockFormat constructs a new QTextBlockFormat object.
func NewQTextBlockFormat() *QTextBlockFormat {

	ret := newQTextBlockFormat(QTextBlockFormat_new())
	ret.isSubclass = true
	return ret
}

func (this *QTextBlockFormat) IsValid() bool {
	return (bool)(QTextBlockFormat_IsValid(this.h))
}

func (this *QTextBlockFormat) SetAlignment(alignment AlignmentFlag) {
	QTextBlockFormat_SetAlignment(this.h, (int)(alignment))
}

func (this *QTextBlockFormat) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QTextBlockFormat_Alignment(this.h))
}

func (this *QTextBlockFormat) SetTopMargin(margin float64) {
	QTextBlockFormat_SetTopMargin(this.h, (double)(margin))
}

func (this *QTextBlockFormat) TopMargin() float64 {
	return (float64)(QTextBlockFormat_TopMargin(this.h))
}

func (this *QTextBlockFormat) SetBottomMargin(margin float64) {
	QTextBlockFormat_SetBottomMargin(this.h, (double)(margin))
}

func (this *QTextBlockFormat) BottomMargin() float64 {
	return (float64)(QTextBlockFormat_BottomMargin(this.h))
}

func (this *QTextBlockFormat) SetLeftMargin(margin float64) {
	QTextBlockFormat_SetLeftMargin(this.h, (double)(margin))
}

func (this *QTextBlockFormat) LeftMargin() float64 {
	return (float64)(QTextBlockFormat_LeftMargin(this.h))
}

func (this *QTextBlockFormat) SetRightMargin(margin float64) {
	QTextBlockFormat_SetRightMargin(this.h, (double)(margin))
}

func (this *QTextBlockFormat) RightMargin() float64 {
	return (float64)(QTextBlockFormat_RightMargin(this.h))
}

func (this *QTextBlockFormat) SetTextIndent(aindent float64) {
	QTextBlockFormat_SetTextIndent(this.h, (double)(aindent))
}

func (this *QTextBlockFormat) TextIndent() float64 {
	return (float64)(QTextBlockFormat_TextIndent(this.h))
}

func (this *QTextBlockFormat) SetIndent(indent int) {
	QTextBlockFormat_SetIndent(this.h, (int)(indent))
}

func (this *QTextBlockFormat) Indent() int {
	return (int)(QTextBlockFormat_Indent(this.h))
}

func (this *QTextBlockFormat) SetHeadingLevel(alevel int) {
	QTextBlockFormat_SetHeadingLevel(this.h, (int)(alevel))
}

func (this *QTextBlockFormat) HeadingLevel() int {
	return (int)(QTextBlockFormat_HeadingLevel(this.h))
}

func (this *QTextBlockFormat) SetLineHeight(height float64, heightType int) {
	QTextBlockFormat_SetLineHeight(this.h, (double)(height), (int)(heightType))
}

func (this *QTextBlockFormat) LineHeight(scriptLineHeight float64, scaling float64) float64 {
	return (float64)(QTextBlockFormat_LineHeight(this.h, (double)(scriptLineHeight), (double)(scaling)))
}

func (this *QTextBlockFormat) LineHeight2() float64 {
	return (float64)(QTextBlockFormat_LineHeight2(this.h))
}

func (this *QTextBlockFormat) LineHeightType() int {
	return (int)(QTextBlockFormat_LineHeightType(this.h))
}

func (this *QTextBlockFormat) SetNonBreakableLines(b bool) {
	QTextBlockFormat_SetNonBreakableLines(this.h, (bool)(b))
}

func (this *QTextBlockFormat) NonBreakableLines() bool {
	return (bool)(QTextBlockFormat_NonBreakableLines(this.h))
}

func (this *QTextBlockFormat) SetPageBreakPolicy(flags PageBreakFlags) {
	QTextBlockFormat_SetPageBreakPolicy(this.h, flags)
}

func (this *QTextBlockFormat) PageBreakPolicy() PageBreakFlags {
	xxxxxxxxx
}

func (this *QTextBlockFormat) SetTabPositions(tabs []QTextOption__Tab) {
	tabs_CArray := (*[0xffff]*QTextOption__Tab)(malloc(size_t(8 * len(tabs))))
	defer free(unsafe.Pointer(tabs_CArray))
	for i := range tabs {
		tabs_CArray[i] = tabs[i].cPointer()
	}
	tabs_ma := struct_miqt_array{len: size_t(len(tabs)), data: unsafe.Pointer(tabs_CArray)}
	QTextBlockFormat_SetTabPositions(this.h, tabs_ma)
}

func (this *QTextBlockFormat) TabPositions() []QTextOption__Tab {
	var _ma struct_miqt_array = QTextBlockFormat_TabPositions(this.h)
	_ret := make([]QTextOption__Tab, int(_ma.len))
	_outCast := (*[0xffff]*QTextOption__Tab)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQTextOption__Tab(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextBlockFormat) SetMarker(marker MarkerType) {
	QTextBlockFormat_SetMarker(this.h, marker)
}

func (this *QTextBlockFormat) Marker() MarkerType {
	xxxxxxxxx
}

type QTextListFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQTextListFormat constructs a new QTextListFormat object.
func NewQTextListFormat() *QTextListFormat {

	ret := newQTextListFormat(QTextListFormat_new())
	ret.isSubclass = true
	return ret
}

func (this *QTextListFormat) IsValid() bool {
	return (bool)(QTextListFormat_IsValid(this.h))
}

func (this *QTextListFormat) SetStyle(style Style) {
	QTextListFormat_SetStyle(this.h, style)
}

func (this *QTextListFormat) Style() Style {
	xxxxxxxxx
}

func (this *QTextListFormat) SetIndent(indent int) {
	QTextListFormat_SetIndent(this.h, (int)(indent))
}

func (this *QTextListFormat) Indent() int {
	return (int)(QTextListFormat_Indent(this.h))
}

func (this *QTextListFormat) SetNumberPrefix(numberPrefix string) {
	numberPrefix_ms := struct_miqt_string{}
	numberPrefix_ms.data = CString(numberPrefix)
	numberPrefix_ms.len = size_t(len(numberPrefix))
	defer free(unsafe.Pointer(numberPrefix_ms.data))
	QTextListFormat_SetNumberPrefix(this.h, numberPrefix_ms)
}

func (this *QTextListFormat) NumberPrefix() string {
	var _ms struct_miqt_string = QTextListFormat_NumberPrefix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextListFormat) SetNumberSuffix(numberSuffix string) {
	numberSuffix_ms := struct_miqt_string{}
	numberSuffix_ms.data = CString(numberSuffix)
	numberSuffix_ms.len = size_t(len(numberSuffix))
	defer free(unsafe.Pointer(numberSuffix_ms.data))
	QTextListFormat_SetNumberSuffix(this.h, numberSuffix_ms)
}

func (this *QTextListFormat) NumberSuffix() string {
	var _ms struct_miqt_string = QTextListFormat_NumberSuffix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextListFormat) SetStart(indent int) {
	QTextListFormat_SetStart(this.h, (int)(indent))
}

func (this *QTextListFormat) Start() int {
	return (int)(QTextListFormat_Start(this.h))
}

type QTextImageFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQTextImageFormat constructs a new QTextImageFormat object.
func NewQTextImageFormat() *QTextImageFormat {

	ret := newQTextImageFormat(QTextImageFormat_new())
	ret.isSubclass = true
	return ret
}

func (this *QTextImageFormat) IsValid() bool {
	return (bool)(QTextImageFormat_IsValid(this.h))
}

func (this *QTextImageFormat) SetName(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QTextImageFormat_SetName(this.h, name_ms)
}

func (this *QTextImageFormat) Name() string {
	var _ms struct_miqt_string = QTextImageFormat_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextImageFormat) SetWidth(width float64) {
	QTextImageFormat_SetWidth(this.h, (double)(width))
}

func (this *QTextImageFormat) Width() float64 {
	return (float64)(QTextImageFormat_Width(this.h))
}

func (this *QTextImageFormat) SetMaximumWidth(maxWidth QTextLength) {
	QTextImageFormat_SetMaximumWidth(this.h, maxWidth.cPointer())
}

func (this *QTextImageFormat) MaximumWidth() *QTextLength {
	_goptr := newQTextLength(QTextImageFormat_MaximumWidth(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextImageFormat) SetHeight(height float64) {
	QTextImageFormat_SetHeight(this.h, (double)(height))
}

func (this *QTextImageFormat) Height() float64 {
	return (float64)(QTextImageFormat_Height(this.h))
}

func (this *QTextImageFormat) SetQuality(quality int) {
	QTextImageFormat_SetQuality(this.h, (int)(quality))
}

func (this *QTextImageFormat) SetQuality2() {
	QTextImageFormat_SetQuality2(this.h)
}

func (this *QTextImageFormat) Quality() int {
	return (int)(QTextImageFormat_Quality(this.h))
}

type QTextFrameFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQTextFrameFormat constructs a new QTextFrameFormat object.
func NewQTextFrameFormat() *QTextFrameFormat {

	ret := newQTextFrameFormat(QTextFrameFormat_new())
	ret.isSubclass = true
	return ret
}

func (this *QTextFrameFormat) IsValid() bool {
	return (bool)(QTextFrameFormat_IsValid(this.h))
}

func (this *QTextFrameFormat) SetPosition(f Position) {
	QTextFrameFormat_SetPosition(this.h, f)
}

func (this *QTextFrameFormat) Position() Position {
	xxxxxxxxx
}

func (this *QTextFrameFormat) SetBorder(border float64) {
	QTextFrameFormat_SetBorder(this.h, (double)(border))
}

func (this *QTextFrameFormat) Border() float64 {
	return (float64)(QTextFrameFormat_Border(this.h))
}

func (this *QTextFrameFormat) SetBorderBrush(brush *QBrush) {
	QTextFrameFormat_SetBorderBrush(this.h, brush.cPointer())
}

func (this *QTextFrameFormat) BorderBrush() *QBrush {
	_goptr := newQBrush(QTextFrameFormat_BorderBrush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFrameFormat) SetBorderStyle(style BorderStyle) {
	QTextFrameFormat_SetBorderStyle(this.h, style)
}

func (this *QTextFrameFormat) BorderStyle() BorderStyle {
	xxxxxxxxx
}

func (this *QTextFrameFormat) SetMargin(margin float64) {
	QTextFrameFormat_SetMargin(this.h, (double)(margin))
}

func (this *QTextFrameFormat) Margin() float64 {
	return (float64)(QTextFrameFormat_Margin(this.h))
}

func (this *QTextFrameFormat) SetTopMargin(margin float64) {
	QTextFrameFormat_SetTopMargin(this.h, (double)(margin))
}

func (this *QTextFrameFormat) TopMargin() float64 {
	return (float64)(QTextFrameFormat_TopMargin(this.h))
}

func (this *QTextFrameFormat) SetBottomMargin(margin float64) {
	QTextFrameFormat_SetBottomMargin(this.h, (double)(margin))
}

func (this *QTextFrameFormat) BottomMargin() float64 {
	return (float64)(QTextFrameFormat_BottomMargin(this.h))
}

func (this *QTextFrameFormat) SetLeftMargin(margin float64) {
	QTextFrameFormat_SetLeftMargin(this.h, (double)(margin))
}

func (this *QTextFrameFormat) LeftMargin() float64 {
	return (float64)(QTextFrameFormat_LeftMargin(this.h))
}

func (this *QTextFrameFormat) SetRightMargin(margin float64) {
	QTextFrameFormat_SetRightMargin(this.h, (double)(margin))
}

func (this *QTextFrameFormat) RightMargin() float64 {
	return (float64)(QTextFrameFormat_RightMargin(this.h))
}

func (this *QTextFrameFormat) SetPadding(padding float64) {
	QTextFrameFormat_SetPadding(this.h, (double)(padding))
}

func (this *QTextFrameFormat) Padding() float64 {
	return (float64)(QTextFrameFormat_Padding(this.h))
}

func (this *QTextFrameFormat) SetWidth(width float64) {
	QTextFrameFormat_SetWidth(this.h, (double)(width))
}

func (this *QTextFrameFormat) SetWidthWithLength(length *QTextLength) {
	QTextFrameFormat_SetWidthWithLength(this.h, length.cPointer())
}

func (this *QTextFrameFormat) Width() *QTextLength {
	_goptr := newQTextLength(QTextFrameFormat_Width(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFrameFormat) SetHeight(height float64) {
	QTextFrameFormat_SetHeight(this.h, (double)(height))
}

func (this *QTextFrameFormat) SetHeightWithHeight(height *QTextLength) {
	QTextFrameFormat_SetHeightWithHeight(this.h, height.cPointer())
}

func (this *QTextFrameFormat) Height() *QTextLength {
	_goptr := newQTextLength(QTextFrameFormat_Height(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFrameFormat) SetPageBreakPolicy(flags PageBreakFlags) {
	QTextFrameFormat_SetPageBreakPolicy(this.h, flags)
}

func (this *QTextFrameFormat) PageBreakPolicy() PageBreakFlags {
	xxxxxxxxx
}

type QTextTableFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQTextTableFormat constructs a new QTextTableFormat object.
func NewQTextTableFormat() *QTextTableFormat {

	ret := newQTextTableFormat(QTextTableFormat_new())
	ret.isSubclass = true
	return ret
}

func (this *QTextTableFormat) IsValid() bool {
	return (bool)(QTextTableFormat_IsValid(this.h))
}

func (this *QTextTableFormat) Columns() int {
	return (int)(QTextTableFormat_Columns(this.h))
}

func (this *QTextTableFormat) SetColumns(columns int) {
	QTextTableFormat_SetColumns(this.h, (int)(columns))
}

func (this *QTextTableFormat) SetColumnWidthConstraints(constraints []QTextLength) {
	constraints_CArray := (*[0xffff]*QTextLength)(malloc(size_t(8 * len(constraints))))
	defer free(unsafe.Pointer(constraints_CArray))
	for i := range constraints {
		constraints_CArray[i] = constraints[i].cPointer()
	}
	constraints_ma := struct_miqt_array{len: size_t(len(constraints)), data: unsafe.Pointer(constraints_CArray)}
	QTextTableFormat_SetColumnWidthConstraints(this.h, constraints_ma)
}

func (this *QTextTableFormat) ColumnWidthConstraints() []QTextLength {
	var _ma struct_miqt_array = QTextTableFormat_ColumnWidthConstraints(this.h)
	_ret := make([]QTextLength, int(_ma.len))
	_outCast := (*[0xffff]*QTextLength)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQTextLength(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextTableFormat) ClearColumnWidthConstraints() {
	QTextTableFormat_ClearColumnWidthConstraints(this.h)
}

func (this *QTextTableFormat) CellSpacing() float64 {
	return (float64)(QTextTableFormat_CellSpacing(this.h))
}

func (this *QTextTableFormat) SetCellSpacing(spacing float64) {
	QTextTableFormat_SetCellSpacing(this.h, (double)(spacing))
}

func (this *QTextTableFormat) CellPadding() float64 {
	return (float64)(QTextTableFormat_CellPadding(this.h))
}

func (this *QTextTableFormat) SetCellPadding(padding float64) {
	QTextTableFormat_SetCellPadding(this.h, (double)(padding))
}

func (this *QTextTableFormat) SetAlignment(alignment AlignmentFlag) {
	QTextTableFormat_SetAlignment(this.h, (int)(alignment))
}

func (this *QTextTableFormat) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QTextTableFormat_Alignment(this.h))
}

func (this *QTextTableFormat) SetHeaderRowCount(count int) {
	QTextTableFormat_SetHeaderRowCount(this.h, (int)(count))
}

func (this *QTextTableFormat) HeaderRowCount() int {
	return (int)(QTextTableFormat_HeaderRowCount(this.h))
}

func (this *QTextTableFormat) SetBorderCollapse(borderCollapse bool) {
	QTextTableFormat_SetBorderCollapse(this.h, (bool)(borderCollapse))
}

func (this *QTextTableFormat) BorderCollapse() bool {
	return (bool)(QTextTableFormat_BorderCollapse(this.h))
}

type QTextTableCellFormat struct {
	h          uintptr
	isSubclass bool
}

// NewQTextTableCellFormat constructs a new QTextTableCellFormat object.
func NewQTextTableCellFormat() *QTextTableCellFormat {

	ret := newQTextTableCellFormat(QTextTableCellFormat_new())
	ret.isSubclass = true
	return ret
}

func (this *QTextTableCellFormat) IsValid() bool {
	return (bool)(QTextTableCellFormat_IsValid(this.h))
}

func (this *QTextTableCellFormat) SetTopPadding(padding float64) {
	QTextTableCellFormat_SetTopPadding(this.h, (double)(padding))
}

func (this *QTextTableCellFormat) TopPadding() float64 {
	return (float64)(QTextTableCellFormat_TopPadding(this.h))
}

func (this *QTextTableCellFormat) SetBottomPadding(padding float64) {
	QTextTableCellFormat_SetBottomPadding(this.h, (double)(padding))
}

func (this *QTextTableCellFormat) BottomPadding() float64 {
	return (float64)(QTextTableCellFormat_BottomPadding(this.h))
}

func (this *QTextTableCellFormat) SetLeftPadding(padding float64) {
	QTextTableCellFormat_SetLeftPadding(this.h, (double)(padding))
}

func (this *QTextTableCellFormat) LeftPadding() float64 {
	return (float64)(QTextTableCellFormat_LeftPadding(this.h))
}

func (this *QTextTableCellFormat) SetRightPadding(padding float64) {
	QTextTableCellFormat_SetRightPadding(this.h, (double)(padding))
}

func (this *QTextTableCellFormat) RightPadding() float64 {
	return (float64)(QTextTableCellFormat_RightPadding(this.h))
}

func (this *QTextTableCellFormat) SetPadding(padding float64) {
	QTextTableCellFormat_SetPadding(this.h, (double)(padding))
}

func (this *QTextTableCellFormat) SetTopBorder(width float64) {
	QTextTableCellFormat_SetTopBorder(this.h, (double)(width))
}

func (this *QTextTableCellFormat) TopBorder() float64 {
	return (float64)(QTextTableCellFormat_TopBorder(this.h))
}

func (this *QTextTableCellFormat) SetBottomBorder(width float64) {
	QTextTableCellFormat_SetBottomBorder(this.h, (double)(width))
}

func (this *QTextTableCellFormat) BottomBorder() float64 {
	return (float64)(QTextTableCellFormat_BottomBorder(this.h))
}

func (this *QTextTableCellFormat) SetLeftBorder(width float64) {
	QTextTableCellFormat_SetLeftBorder(this.h, (double)(width))
}

func (this *QTextTableCellFormat) LeftBorder() float64 {
	return (float64)(QTextTableCellFormat_LeftBorder(this.h))
}

func (this *QTextTableCellFormat) SetRightBorder(width float64) {
	QTextTableCellFormat_SetRightBorder(this.h, (double)(width))
}

func (this *QTextTableCellFormat) RightBorder() float64 {
	return (float64)(QTextTableCellFormat_RightBorder(this.h))
}

func (this *QTextTableCellFormat) SetBorder(width float64) {
	QTextTableCellFormat_SetBorder(this.h, (double)(width))
}

func (this *QTextTableCellFormat) SetTopBorderStyle(style QTextFrameFormat__BorderStyle) {
	QTextTableCellFormat_SetTopBorderStyle(this.h, (int)(style))
}

func (this *QTextTableCellFormat) TopBorderStyle() QTextFrameFormat__BorderStyle {
	return (QTextFrameFormat__BorderStyle)(QTextTableCellFormat_TopBorderStyle(this.h))
}

func (this *QTextTableCellFormat) SetBottomBorderStyle(style QTextFrameFormat__BorderStyle) {
	QTextTableCellFormat_SetBottomBorderStyle(this.h, (int)(style))
}

func (this *QTextTableCellFormat) BottomBorderStyle() QTextFrameFormat__BorderStyle {
	return (QTextFrameFormat__BorderStyle)(QTextTableCellFormat_BottomBorderStyle(this.h))
}

func (this *QTextTableCellFormat) SetLeftBorderStyle(style QTextFrameFormat__BorderStyle) {
	QTextTableCellFormat_SetLeftBorderStyle(this.h, (int)(style))
}

func (this *QTextTableCellFormat) LeftBorderStyle() QTextFrameFormat__BorderStyle {
	return (QTextFrameFormat__BorderStyle)(QTextTableCellFormat_LeftBorderStyle(this.h))
}

func (this *QTextTableCellFormat) SetRightBorderStyle(style QTextFrameFormat__BorderStyle) {
	QTextTableCellFormat_SetRightBorderStyle(this.h, (int)(style))
}

func (this *QTextTableCellFormat) RightBorderStyle() QTextFrameFormat__BorderStyle {
	return (QTextFrameFormat__BorderStyle)(QTextTableCellFormat_RightBorderStyle(this.h))
}

func (this *QTextTableCellFormat) SetBorderStyle(style QTextFrameFormat__BorderStyle) {
	QTextTableCellFormat_SetBorderStyle(this.h, (int)(style))
}

func (this *QTextTableCellFormat) SetTopBorderBrush(brush *QBrush) {
	QTextTableCellFormat_SetTopBorderBrush(this.h, brush.cPointer())
}

func (this *QTextTableCellFormat) TopBorderBrush() *QBrush {
	_goptr := newQBrush(QTextTableCellFormat_TopBorderBrush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTableCellFormat) SetBottomBorderBrush(brush *QBrush) {
	QTextTableCellFormat_SetBottomBorderBrush(this.h, brush.cPointer())
}

func (this *QTextTableCellFormat) BottomBorderBrush() *QBrush {
	_goptr := newQBrush(QTextTableCellFormat_BottomBorderBrush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTableCellFormat) SetLeftBorderBrush(brush *QBrush) {
	QTextTableCellFormat_SetLeftBorderBrush(this.h, brush.cPointer())
}

func (this *QTextTableCellFormat) LeftBorderBrush() *QBrush {
	_goptr := newQBrush(QTextTableCellFormat_LeftBorderBrush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTableCellFormat) SetRightBorderBrush(brush *QBrush) {
	QTextTableCellFormat_SetRightBorderBrush(this.h, brush.cPointer())
}

func (this *QTextTableCellFormat) RightBorderBrush() *QBrush {
	_goptr := newQBrush(QTextTableCellFormat_RightBorderBrush(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextTableCellFormat) SetBorderBrush(brush *QBrush) {
	QTextTableCellFormat_SetBorderBrush(this.h, brush.cPointer())
}
