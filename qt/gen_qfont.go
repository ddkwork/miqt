package qt

import (
	"unsafe"
)

type QFont__StyleHint int

const (
	QFont__Helvetica  QFont__StyleHint = 0
	QFont__SansSerif  QFont__StyleHint = 0
	QFont__Times      QFont__StyleHint = 1
	QFont__Serif      QFont__StyleHint = 1
	QFont__Courier    QFont__StyleHint = 2
	QFont__TypeWriter QFont__StyleHint = 2
	QFont__OldEnglish QFont__StyleHint = 3
	QFont__Decorative QFont__StyleHint = 3
	QFont__System     QFont__StyleHint = 4
	QFont__AnyStyle   QFont__StyleHint = 5
	QFont__Cursive    QFont__StyleHint = 6
	QFont__Monospace  QFont__StyleHint = 7
	QFont__Fantasy    QFont__StyleHint = 8
)

type QFont__StyleStrategy int

const (
	QFont__PreferDefault         QFont__StyleStrategy = 1
	QFont__PreferBitmap          QFont__StyleStrategy = 2
	QFont__PreferDevice          QFont__StyleStrategy = 4
	QFont__PreferOutline         QFont__StyleStrategy = 8
	QFont__ForceOutline          QFont__StyleStrategy = 16
	QFont__PreferMatch           QFont__StyleStrategy = 32
	QFont__PreferQuality         QFont__StyleStrategy = 64
	QFont__PreferAntialias       QFont__StyleStrategy = 128
	QFont__NoAntialias           QFont__StyleStrategy = 256
	QFont__NoSubpixelAntialias   QFont__StyleStrategy = 2048
	QFont__PreferNoShaping       QFont__StyleStrategy = 4096
	QFont__ContextFontMerging    QFont__StyleStrategy = 8192
	QFont__PreferTypoLineMetrics QFont__StyleStrategy = 16384
	QFont__NoFontMerging         QFont__StyleStrategy = 32768
)

type QFont__HintingPreference int

const (
	QFont__PreferDefaultHinting  QFont__HintingPreference = 0
	QFont__PreferNoHinting       QFont__HintingPreference = 1
	QFont__PreferVerticalHinting QFont__HintingPreference = 2
	QFont__PreferFullHinting     QFont__HintingPreference = 3
)

type QFont__Weight int

const (
	QFont__Thin       QFont__Weight = 100
	QFont__ExtraLight QFont__Weight = 200
	QFont__Light      QFont__Weight = 300
	QFont__Normal     QFont__Weight = 400
	QFont__Medium     QFont__Weight = 500
	QFont__DemiBold   QFont__Weight = 600
	QFont__Bold       QFont__Weight = 700
	QFont__ExtraBold  QFont__Weight = 800
	QFont__Black      QFont__Weight = 900
)

type QFont__Style int

const (
	QFont__StyleNormal  QFont__Style = 0
	QFont__StyleItalic  QFont__Style = 1
	QFont__StyleOblique QFont__Style = 2
)

type QFont__Stretch int

const (
	QFont__AnyStretch     QFont__Stretch = 0
	QFont__UltraCondensed QFont__Stretch = 50
	QFont__ExtraCondensed QFont__Stretch = 62
	QFont__Condensed      QFont__Stretch = 75
	QFont__SemiCondensed  QFont__Stretch = 87
	QFont__Unstretched    QFont__Stretch = 100
	QFont__SemiExpanded   QFont__Stretch = 112
	QFont__Expanded       QFont__Stretch = 125
	QFont__ExtraExpanded  QFont__Stretch = 150
	QFont__UltraExpanded  QFont__Stretch = 200
)

type QFont__Capitalization int

const (
	QFont__MixedCase    QFont__Capitalization = 0
	QFont__AllUppercase QFont__Capitalization = 1
	QFont__AllLowercase QFont__Capitalization = 2
	QFont__SmallCaps    QFont__Capitalization = 3
	QFont__Capitalize   QFont__Capitalization = 4
)

type QFont__SpacingType int

const (
	QFont__PercentageSpacing QFont__SpacingType = 0
	QFont__AbsoluteSpacing   QFont__SpacingType = 1
)

type QFont__ResolveProperties int

const (
	QFont__NoPropertiesResolved      QFont__ResolveProperties = 0
	QFont__FamilyResolved            QFont__ResolveProperties = 1
	QFont__SizeResolved              QFont__ResolveProperties = 2
	QFont__StyleHintResolved         QFont__ResolveProperties = 4
	QFont__StyleStrategyResolved     QFont__ResolveProperties = 8
	QFont__WeightResolved            QFont__ResolveProperties = 16
	QFont__StyleResolved             QFont__ResolveProperties = 32
	QFont__UnderlineResolved         QFont__ResolveProperties = 64
	QFont__OverlineResolved          QFont__ResolveProperties = 128
	QFont__StrikeOutResolved         QFont__ResolveProperties = 256
	QFont__FixedPitchResolved        QFont__ResolveProperties = 512
	QFont__StretchResolved           QFont__ResolveProperties = 1024
	QFont__KerningResolved           QFont__ResolveProperties = 2048
	QFont__CapitalizationResolved    QFont__ResolveProperties = 4096
	QFont__LetterSpacingResolved     QFont__ResolveProperties = 8192
	QFont__WordSpacingResolved       QFont__ResolveProperties = 16384
	QFont__HintingPreferenceResolved QFont__ResolveProperties = 32768
	QFont__StyleNameResolved         QFont__ResolveProperties = 65536
	QFont__FamiliesResolved          QFont__ResolveProperties = 131072
	QFont__FeaturesResolved          QFont__ResolveProperties = 262144
	QFont__VariableAxesResolved      QFont__ResolveProperties = 524288
	QFont__AllPropertiesResolved     QFont__ResolveProperties = 1048575
)

type QFont struct {
	h          uintptr
	isSubclass bool
}

// NewQFont constructs a new QFont object.
func NewQFont() *QFont {
	ret := newQFont(QFont_new())
	ret.isSubclass = true
	return ret
}

// NewQFont2 constructs a new QFont object.
func NewQFont2(family string) *QFont {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))

	ret := newQFont(QFont_new2(family_ms))
	ret.isSubclass = true
	return ret
}

// NewQFont3 constructs a new QFont object.
func NewQFont3(families []string) *QFont {
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

	ret := newQFont(QFont_new3(families_ma))
	ret.isSubclass = true
	return ret
}

// NewQFont4 constructs a new QFont object.
func NewQFont4(font *QFont, pd *QPaintDevice) *QFont {
	ret := newQFont(QFont_new4(font.cPointer(), pd.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFont5 constructs a new QFont object.
func NewQFont5(font *QFont) *QFont {
	ret := newQFont(QFont_new5(font.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFont6 constructs a new QFont object.
func NewQFont6(family string, pointSize int) *QFont {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))

	ret := newQFont(QFont_new6(family_ms, (int)(pointSize)))
	ret.isSubclass = true
	return ret
}

// NewQFont7 constructs a new QFont object.
func NewQFont7(family string, pointSize int, weight int) *QFont {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))

	ret := newQFont(QFont_new7(family_ms, (int)(pointSize), (int)(weight)))
	ret.isSubclass = true
	return ret
}

// NewQFont8 constructs a new QFont object.
func NewQFont8(family string, pointSize int, weight int, italic bool) *QFont {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))

	ret := newQFont(QFont_new8(family_ms, (int)(pointSize), (int)(weight), (bool)(italic)))
	ret.isSubclass = true
	return ret
}

// NewQFont9 constructs a new QFont object.
func NewQFont9(families []string, pointSize int) *QFont {
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

	ret := newQFont(QFont_new9(families_ma, (int)(pointSize)))
	ret.isSubclass = true
	return ret
}

// NewQFont10 constructs a new QFont object.
func NewQFont10(families []string, pointSize int, weight int) *QFont {
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

	ret := newQFont(QFont_new10(families_ma, (int)(pointSize), (int)(weight)))
	ret.isSubclass = true
	return ret
}

// NewQFont11 constructs a new QFont object.
func NewQFont11(families []string, pointSize int, weight int, italic bool) *QFont {
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

	ret := newQFont(QFont_new11(families_ma, (int)(pointSize), (int)(weight), (bool)(italic)))
	ret.isSubclass = true
	return ret
}

func (this *QFont) Swap(other *QFont) {
	QFont_Swap(this.h, other.cPointer())
}

func (this *QFont) Family() string {
	var _ms struct_miqt_string = QFont_Family(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFont) SetFamily(family string) {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	QFont_SetFamily(this.h, family_ms)
}

func (this *QFont) Families() []string {
	var _ma struct_miqt_array = QFont_Families(this.h)
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

func (this *QFont) SetFamilies(families []string) {
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
	QFont_SetFamilies(this.h, families_ma)
}

func (this *QFont) StyleName() string {
	var _ms struct_miqt_string = QFont_StyleName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFont) SetStyleName(styleName string) {
	styleName_ms := struct_miqt_string{}
	styleName_ms.data = CString(styleName)
	styleName_ms.len = size_t(len(styleName))
	defer free(unsafe.Pointer(styleName_ms.data))
	QFont_SetStyleName(this.h, styleName_ms)
}

func (this *QFont) PointSize() int {
	return (int)(QFont_PointSize(this.h))
}

func (this *QFont) SetPointSize(pointSize int) {
	QFont_SetPointSize(this.h, (int)(pointSize))
}

func (this *QFont) PointSizeF() float64 {
	return (float64)(QFont_PointSizeF(this.h))
}

func (this *QFont) SetPointSizeF(pointSizeF float64) {
	QFont_SetPointSizeF(this.h, (double)(pointSizeF))
}

func (this *QFont) PixelSize() int {
	return (int)(QFont_PixelSize(this.h))
}

func (this *QFont) SetPixelSize(pixelSize int) {
	QFont_SetPixelSize(this.h, (int)(pixelSize))
}

func (this *QFont) Weight() Weight {
	xxxxxxxxx
}

func (this *QFont) SetWeight(weight Weight) {
	QFont_SetWeight(this.h, weight)
}

func (this *QFont) Bold() bool {
	return (bool)(QFont_Bold(this.h))
}

func (this *QFont) SetBold(bold bool) {
	QFont_SetBold(this.h, (bool)(bold))
}

func (this *QFont) SetStyle(style Style) {
	QFont_SetStyle(this.h, style)
}

func (this *QFont) Style() Style {
	xxxxxxxxx
}

func (this *QFont) Italic() bool {
	return (bool)(QFont_Italic(this.h))
}

func (this *QFont) SetItalic(b bool) {
	QFont_SetItalic(this.h, (bool)(b))
}

func (this *QFont) Underline() bool {
	return (bool)(QFont_Underline(this.h))
}

func (this *QFont) SetUnderline(underline bool) {
	QFont_SetUnderline(this.h, (bool)(underline))
}

func (this *QFont) Overline() bool {
	return (bool)(QFont_Overline(this.h))
}

func (this *QFont) SetOverline(overline bool) {
	QFont_SetOverline(this.h, (bool)(overline))
}

func (this *QFont) StrikeOut() bool {
	return (bool)(QFont_StrikeOut(this.h))
}

func (this *QFont) SetStrikeOut(strikeOut bool) {
	QFont_SetStrikeOut(this.h, (bool)(strikeOut))
}

func (this *QFont) FixedPitch() bool {
	return (bool)(QFont_FixedPitch(this.h))
}

func (this *QFont) SetFixedPitch(fixedPitch bool) {
	QFont_SetFixedPitch(this.h, (bool)(fixedPitch))
}

func (this *QFont) Kerning() bool {
	return (bool)(QFont_Kerning(this.h))
}

func (this *QFont) SetKerning(kerning bool) {
	QFont_SetKerning(this.h, (bool)(kerning))
}

func (this *QFont) StyleHint() StyleHint {
	xxxxxxxxx
}

func (this *QFont) StyleStrategy() StyleStrategy {
	xxxxxxxxx
}

func (this *QFont) SetStyleHint(param1 StyleHint) {
	QFont_SetStyleHint(this.h, param1)
}

func (this *QFont) SetStyleStrategy(s StyleStrategy) {
	QFont_SetStyleStrategy(this.h, s)
}

func (this *QFont) Stretch() int {
	return (int)(QFont_Stretch(this.h))
}

func (this *QFont) SetStretch(stretch int) {
	QFont_SetStretch(this.h, (int)(stretch))
}

func (this *QFont) LetterSpacing() float64 {
	return (float64)(QFont_LetterSpacing(this.h))
}

func (this *QFont) LetterSpacingType() SpacingType {
	xxxxxxxxx
}

func (this *QFont) SetLetterSpacing(typeVal SpacingType, spacing float64) {
	QFont_SetLetterSpacing(this.h, typeVal, (double)(spacing))
}

func (this *QFont) WordSpacing() float64 {
	return (float64)(QFont_WordSpacing(this.h))
}

func (this *QFont) SetWordSpacing(spacing float64) {
	QFont_SetWordSpacing(this.h, (double)(spacing))
}

func (this *QFont) SetCapitalization(capitalization Capitalization) {
	QFont_SetCapitalization(this.h, capitalization)
}

func (this *QFont) Capitalization() Capitalization {
	xxxxxxxxx
}

func (this *QFont) SetHintingPreference(hintingPreference HintingPreference) {
	QFont_SetHintingPreference(this.h, hintingPreference)
}

func (this *QFont) HintingPreference() HintingPreference {
	xxxxxxxxx
}

func (this *QFont) SetFeature(tag Tag, value uint) {
	QFont_SetFeature(this.h, tag, (uint)(value))
}

func (this *QFont) UnsetFeature(tag Tag) {
	QFont_UnsetFeature(this.h, tag)
}

func (this *QFont) FeatureValue(tag Tag) uint {
	return (uint)(QFont_FeatureValue(this.h, tag))
}

func (this *QFont) IsFeatureSet(tag Tag) bool {
	return (bool)(QFont_IsFeatureSet(this.h, tag))
}

func (this *QFont) FeatureTags() []Tag {
	var _ma struct_miqt_array = QFont_FeatureTags(this.h)
	_ret := make([]Tag, int(_ma.len))
	_outCast := (*[0xffff]Tag)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QFont) ClearFeatures() {
	QFont_ClearFeatures(this.h)
}

func (this *QFont) SetVariableAxis(tag Tag, value float32) {
	QFont_SetVariableAxis(this.h, tag, (float)(value))
}

func (this *QFont) UnsetVariableAxis(tag Tag) {
	QFont_UnsetVariableAxis(this.h, tag)
}

func (this *QFont) IsVariableAxisSet(tag Tag) bool {
	return (bool)(QFont_IsVariableAxisSet(this.h, tag))
}

func (this *QFont) VariableAxisValue(tag Tag) float32 {
	return (float32)(QFont_VariableAxisValue(this.h, tag))
}

func (this *QFont) ClearVariableAxes() {
	QFont_ClearVariableAxes(this.h)
}

func (this *QFont) VariableAxisTags() []Tag {
	var _ma struct_miqt_array = QFont_VariableAxisTags(this.h)
	_ret := make([]Tag, int(_ma.len))
	_outCast := (*[0xffff]Tag)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QFont) ExactMatch() bool {
	return (bool)(QFont_ExactMatch(this.h))
}

func (this *QFont) OperatorAssign(param1 *QFont) {
	QFont_OperatorAssign(this.h, param1.cPointer())
}

func (this *QFont) OperatorEqual(param1 *QFont) bool {
	return (bool)(QFont_OperatorEqual(this.h, param1.cPointer()))
}

func (this *QFont) OperatorNotEqual(param1 *QFont) bool {
	return (bool)(QFont_OperatorNotEqual(this.h, param1.cPointer()))
}

func (this *QFont) OperatorLesser(param1 *QFont) bool {
	return (bool)(QFont_OperatorLesser(this.h, param1.cPointer()))
}

func (this *QFont) IsCopyOf(param1 *QFont) bool {
	return (bool)(QFont_IsCopyOf(this.h, param1.cPointer()))
}

func (this *QFont) Key() string {
	var _ms struct_miqt_string = QFont_Key(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFont) ToString() string {
	var _ms struct_miqt_string = QFont_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFont) FromString(param1 string) bool {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	return (bool)(QFont_FromString(this.h, param1_ms))
}

func QFont_Substitute(param1 string) string {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	var _ms struct_miqt_string = QFont_Substitute(param1_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFont_Substitutes(param1 string) []string {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	var _ma struct_miqt_array = QFont_Substitutes(param1_ms)
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

func QFont_Substitutions() []string {
	var _ma struct_miqt_array = QFont_Substitutions()
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

func QFont_InsertSubstitution(param1 string, param2 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	param2_ms := struct_miqt_string{}
	param2_ms.data = CString(param2)
	param2_ms.len = size_t(len(param2))
	defer free(unsafe.Pointer(param2_ms.data))
	QFont_InsertSubstitution(param1_ms, param2_ms)
}

func QFont_InsertSubstitutions(param1 string, param2 []string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	param2_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(param2))))
	defer free(unsafe.Pointer(param2_CArray))
	for i := range param2 {
		param2_i_ms := struct_miqt_string{}
		param2_i_ms.data = CString(param2[i])
		param2_i_ms.len = size_t(len(param2[i]))
		defer free(unsafe.Pointer(param2_i_ms.data))
		param2_CArray[i] = param2_i_ms
	}
	param2_ma := struct_miqt_array{len: size_t(len(param2)), data: unsafe.Pointer(param2_CArray)}
	QFont_InsertSubstitutions(param1_ms, param2_ma)
}

func QFont_RemoveSubstitutions(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QFont_RemoveSubstitutions(param1_ms)
}

func QFont_Initialize() {
	QFont_Initialize()
}

func QFont_Cleanup() {
	QFont_Cleanup()
}

func QFont_CacheStatistics() {
	QFont_CacheStatistics()
}

func (this *QFont) DefaultFamily() string {
	var _ms struct_miqt_string = QFont_DefaultFamily(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFont) Resolve(param1 *QFont) *QFont {
	_goptr := newQFont(QFont_Resolve(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFont) ResolveMask() uint {
	return (uint)(QFont_ResolveMask(this.h))
}

func (this *QFont) SetResolveMask(mask uint) {
	QFont_SetResolveMask(this.h, (uint)(mask))
}

func (this *QFont) SetLegacyWeight(legacyWeight int) {
	QFont_SetLegacyWeight(this.h, (int)(legacyWeight))
}

func (this *QFont) LegacyWeight() int {
	return (int)(QFont_LegacyWeight(this.h))
}

func (this *QFont) SetStyleHint2(param1 StyleHint, param2 StyleStrategy) {
	QFont_SetStyleHint2(this.h, param1, param2)
}

type QFont__Tag struct {
	h          uintptr
	isSubclass bool
}

// NewQFont__Tag constructs a new QFont::Tag object.
func NewQFont__Tag() *QFont__Tag {
	ret := newQFont__Tag(QFont__Tag_new())
	ret.isSubclass = true
	return ret
}

// NewQFont__Tag2 constructs a new QFont::Tag object.
func NewQFont__Tag2(param1 *Tag) *QFont__Tag {
	ret := newQFont__Tag(QFont__Tag_new2(param1))
	ret.isSubclass = true
	return ret
}

func (this *QFont__Tag) IsValid() bool {
	return (bool)(QFont__Tag_IsValid(this.h))
}

func (this *QFont__Tag) Value() uint {
	return (uint)(QFont__Tag_Value(this.h))
}

func (this *QFont__Tag) ToString() []byte {
	var _bytearray struct_miqt_string = QFont__Tag_ToString(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}
