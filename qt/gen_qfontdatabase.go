package qt

import (
	"unsafe"
)

type QFontDatabase__WritingSystem int

const (
	QFontDatabase__Any                 QFontDatabase__WritingSystem = 0
	QFontDatabase__Latin               QFontDatabase__WritingSystem = 1
	QFontDatabase__Greek               QFontDatabase__WritingSystem = 2
	QFontDatabase__Cyrillic            QFontDatabase__WritingSystem = 3
	QFontDatabase__Armenian            QFontDatabase__WritingSystem = 4
	QFontDatabase__Hebrew              QFontDatabase__WritingSystem = 5
	QFontDatabase__Arabic              QFontDatabase__WritingSystem = 6
	QFontDatabase__Syriac              QFontDatabase__WritingSystem = 7
	QFontDatabase__Thaana              QFontDatabase__WritingSystem = 8
	QFontDatabase__Devanagari          QFontDatabase__WritingSystem = 9
	QFontDatabase__Bengali             QFontDatabase__WritingSystem = 10
	QFontDatabase__Gurmukhi            QFontDatabase__WritingSystem = 11
	QFontDatabase__Gujarati            QFontDatabase__WritingSystem = 12
	QFontDatabase__Oriya               QFontDatabase__WritingSystem = 13
	QFontDatabase__Tamil               QFontDatabase__WritingSystem = 14
	QFontDatabase__Telugu              QFontDatabase__WritingSystem = 15
	QFontDatabase__Kannada             QFontDatabase__WritingSystem = 16
	QFontDatabase__Malayalam           QFontDatabase__WritingSystem = 17
	QFontDatabase__Sinhala             QFontDatabase__WritingSystem = 18
	QFontDatabase__Thai                QFontDatabase__WritingSystem = 19
	QFontDatabase__Lao                 QFontDatabase__WritingSystem = 20
	QFontDatabase__Tibetan             QFontDatabase__WritingSystem = 21
	QFontDatabase__Myanmar             QFontDatabase__WritingSystem = 22
	QFontDatabase__Georgian            QFontDatabase__WritingSystem = 23
	QFontDatabase__Khmer               QFontDatabase__WritingSystem = 24
	QFontDatabase__SimplifiedChinese   QFontDatabase__WritingSystem = 25
	QFontDatabase__TraditionalChinese  QFontDatabase__WritingSystem = 26
	QFontDatabase__Japanese            QFontDatabase__WritingSystem = 27
	QFontDatabase__Korean              QFontDatabase__WritingSystem = 28
	QFontDatabase__Vietnamese          QFontDatabase__WritingSystem = 29
	QFontDatabase__Symbol              QFontDatabase__WritingSystem = 30
	QFontDatabase__Other               QFontDatabase__WritingSystem = 30
	QFontDatabase__Ogham               QFontDatabase__WritingSystem = 31
	QFontDatabase__Runic               QFontDatabase__WritingSystem = 32
	QFontDatabase__Nko                 QFontDatabase__WritingSystem = 33
	QFontDatabase__WritingSystemsCount QFontDatabase__WritingSystem = 34
)

type QFontDatabase__SystemFont int

const (
	QFontDatabase__GeneralFont          QFontDatabase__SystemFont = 0
	QFontDatabase__FixedFont            QFontDatabase__SystemFont = 1
	QFontDatabase__TitleFont            QFontDatabase__SystemFont = 2
	QFontDatabase__SmallestReadableFont QFontDatabase__SystemFont = 3
)

type QFontDatabase struct {
	h          uintptr
	isSubclass bool
}

// NewQFontDatabase constructs a new QFontDatabase object.
func NewQFontDatabase() *QFontDatabase {

	ret := newQFontDatabase(QFontDatabase_new())
	ret.isSubclass = true
	return ret
}

func QFontDatabase_StandardSizes() []int {
	var _ma struct_miqt_array = QFontDatabase_StandardSizes()
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func QFontDatabase_WritingSystems() []WritingSystem {
	var _ma struct_miqt_array = QFontDatabase_WritingSystems()
	_ret := make([]WritingSystem, int(_ma.len))
	_outCast := (*[0xffff]WritingSystem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func QFontDatabase_WritingSystemsWithFamily(family string) []WritingSystem {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	var _ma struct_miqt_array = QFontDatabase_WritingSystemsWithFamily(family_ms)
	_ret := make([]WritingSystem, int(_ma.len))
	_outCast := (*[0xffff]WritingSystem)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func QFontDatabase_Families() []string {
	var _ma struct_miqt_array = QFontDatabase_Families()
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

func QFontDatabase_Styles(family string) []string {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	var _ma struct_miqt_array = QFontDatabase_Styles(family_ms)
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

func QFontDatabase_PointSizes(family string) []int {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	var _ma struct_miqt_array = QFontDatabase_PointSizes(family_ms)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func QFontDatabase_SmoothSizes(family string, style string) []int {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	style_ms := struct_miqt_string{}
	style_ms.data = CString(style)
	style_ms.len = size_t(len(style))
	defer free(unsafe.Pointer(style_ms.data))
	var _ma struct_miqt_array = QFontDatabase_SmoothSizes(family_ms, style_ms)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func QFontDatabase_StyleString(font *QFont) string {
	var _ms struct_miqt_string = QFontDatabase_StyleString(font.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFontDatabase_StyleStringWithFontInfo(fontInfo *QFontInfo) string {
	var _ms struct_miqt_string = QFontDatabase_StyleStringWithFontInfo(fontInfo.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFontDatabase_Font(family string, style string, pointSize int) *QFont {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	style_ms := struct_miqt_string{}
	style_ms.data = CString(style)
	style_ms.len = size_t(len(style))
	defer free(unsafe.Pointer(style_ms.data))
	_goptr := newQFont(QFontDatabase_Font(family_ms, style_ms, (int)(pointSize)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFontDatabase_IsBitmapScalable(family string) bool {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	return (bool)(QFontDatabase_IsBitmapScalable(family_ms))
}

func QFontDatabase_IsSmoothlyScalable(family string) bool {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	return (bool)(QFontDatabase_IsSmoothlyScalable(family_ms))
}

func QFontDatabase_IsScalable(family string) bool {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	return (bool)(QFontDatabase_IsScalable(family_ms))
}

func QFontDatabase_IsFixedPitch(family string) bool {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	return (bool)(QFontDatabase_IsFixedPitch(family_ms))
}

func QFontDatabase_Italic(family string, style string) bool {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	style_ms := struct_miqt_string{}
	style_ms.data = CString(style)
	style_ms.len = size_t(len(style))
	defer free(unsafe.Pointer(style_ms.data))
	return (bool)(QFontDatabase_Italic(family_ms, style_ms))
}

func QFontDatabase_Bold(family string, style string) bool {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	style_ms := struct_miqt_string{}
	style_ms.data = CString(style)
	style_ms.len = size_t(len(style))
	defer free(unsafe.Pointer(style_ms.data))
	return (bool)(QFontDatabase_Bold(family_ms, style_ms))
}

func QFontDatabase_Weight(family string, style string) int {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	style_ms := struct_miqt_string{}
	style_ms.data = CString(style)
	style_ms.len = size_t(len(style))
	defer free(unsafe.Pointer(style_ms.data))
	return (int)(QFontDatabase_Weight(family_ms, style_ms))
}

func QFontDatabase_HasFamily(family string) bool {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	return (bool)(QFontDatabase_HasFamily(family_ms))
}

func QFontDatabase_IsPrivateFamily(family string) bool {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	return (bool)(QFontDatabase_IsPrivateFamily(family_ms))
}

func QFontDatabase_WritingSystemName(writingSystem WritingSystem) string {
	var _ms struct_miqt_string = QFontDatabase_WritingSystemName(writingSystem)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFontDatabase_WritingSystemSample(writingSystem WritingSystem) string {
	var _ms struct_miqt_string = QFontDatabase_WritingSystemSample(writingSystem)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFontDatabase_AddApplicationFont(fileName string) int {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (int)(QFontDatabase_AddApplicationFont(fileName_ms))
}

func QFontDatabase_AddApplicationFontFromData(fontData []byte) int {
	fontData_alias := struct_miqt_string{}
	fontData_alias.data = (char)(unsafe.Pointer(&fontData[0]))
	fontData_alias.len = size_t(len(fontData))
	return (int)(QFontDatabase_AddApplicationFontFromData(fontData_alias))
}

func QFontDatabase_ApplicationFontFamilies(id int) []string {
	var _ma struct_miqt_array = QFontDatabase_ApplicationFontFamilies((int)(id))
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

func QFontDatabase_RemoveApplicationFont(id int) bool {
	return (bool)(QFontDatabase_RemoveApplicationFont((int)(id)))
}

func QFontDatabase_RemoveAllApplicationFonts() bool {
	return (bool)(QFontDatabase_RemoveAllApplicationFonts())
}

func QFontDatabase_AddApplicationFallbackFontFamily(script QChar__Script, familyName string) {
	familyName_ms := struct_miqt_string{}
	familyName_ms.data = CString(familyName)
	familyName_ms.len = size_t(len(familyName))
	defer free(unsafe.Pointer(familyName_ms.data))
	QFontDatabase_AddApplicationFallbackFontFamily((int)(script), familyName_ms)
}

func QFontDatabase_RemoveApplicationFallbackFontFamily(script QChar__Script, familyName string) bool {
	familyName_ms := struct_miqt_string{}
	familyName_ms.data = CString(familyName)
	familyName_ms.len = size_t(len(familyName))
	defer free(unsafe.Pointer(familyName_ms.data))
	return (bool)(QFontDatabase_RemoveApplicationFallbackFontFamily((int)(script), familyName_ms))
}

func QFontDatabase_SetApplicationFallbackFontFamilies(param1 QChar__Script, familyNames []string) {
	familyNames_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(familyNames))))
	defer free(unsafe.Pointer(familyNames_CArray))
	for i := range familyNames {
		familyNames_i_ms := struct_miqt_string{}
		familyNames_i_ms.data = CString(familyNames[i])
		familyNames_i_ms.len = size_t(len(familyNames[i]))
		defer free(unsafe.Pointer(familyNames_i_ms.data))
		familyNames_CArray[i] = familyNames_i_ms
	}
	familyNames_ma := struct_miqt_array{len: size_t(len(familyNames)), data: unsafe.Pointer(familyNames_CArray)}
	QFontDatabase_SetApplicationFallbackFontFamilies((int)(param1), familyNames_ma)
}

func QFontDatabase_ApplicationFallbackFontFamilies(script QChar__Script) []string {
	var _ma struct_miqt_array = QFontDatabase_ApplicationFallbackFontFamilies((int)(script))
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

func QFontDatabase_AddApplicationEmojiFontFamily(familyName string) {
	familyName_ms := struct_miqt_string{}
	familyName_ms.data = CString(familyName)
	familyName_ms.len = size_t(len(familyName))
	defer free(unsafe.Pointer(familyName_ms.data))
	QFontDatabase_AddApplicationEmojiFontFamily(familyName_ms)
}

func QFontDatabase_RemoveApplicationEmojiFontFamily(familyName string) bool {
	familyName_ms := struct_miqt_string{}
	familyName_ms.data = CString(familyName)
	familyName_ms.len = size_t(len(familyName))
	defer free(unsafe.Pointer(familyName_ms.data))
	return (bool)(QFontDatabase_RemoveApplicationEmojiFontFamily(familyName_ms))
}

func QFontDatabase_SetApplicationEmojiFontFamilies(familyNames []string) {
	familyNames_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(familyNames))))
	defer free(unsafe.Pointer(familyNames_CArray))
	for i := range familyNames {
		familyNames_i_ms := struct_miqt_string{}
		familyNames_i_ms.data = CString(familyNames[i])
		familyNames_i_ms.len = size_t(len(familyNames[i]))
		defer free(unsafe.Pointer(familyNames_i_ms.data))
		familyNames_CArray[i] = familyNames_i_ms
	}
	familyNames_ma := struct_miqt_array{len: size_t(len(familyNames)), data: unsafe.Pointer(familyNames_CArray)}
	QFontDatabase_SetApplicationEmojiFontFamilies(familyNames_ma)
}

func QFontDatabase_ApplicationEmojiFontFamilies() []string {
	var _ma struct_miqt_array = QFontDatabase_ApplicationEmojiFontFamilies()
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

func QFontDatabase_SystemFont(typeVal SystemFont) *QFont {
	_goptr := newQFont(QFontDatabase_SystemFont(typeVal))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFontDatabase_Families1(writingSystem WritingSystem) []string {
	var _ma struct_miqt_array = QFontDatabase_Families1(writingSystem)
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

func QFontDatabase_PointSizes2(family string, style string) []int {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	style_ms := struct_miqt_string{}
	style_ms.data = CString(style)
	style_ms.len = size_t(len(style))
	defer free(unsafe.Pointer(style_ms.data))
	var _ma struct_miqt_array = QFontDatabase_PointSizes2(family_ms, style_ms)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func QFontDatabase_IsBitmapScalable2(family string, style string) bool {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	style_ms := struct_miqt_string{}
	style_ms.data = CString(style)
	style_ms.len = size_t(len(style))
	defer free(unsafe.Pointer(style_ms.data))
	return (bool)(QFontDatabase_IsBitmapScalable2(family_ms, style_ms))
}

func QFontDatabase_IsSmoothlyScalable2(family string, style string) bool {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	style_ms := struct_miqt_string{}
	style_ms.data = CString(style)
	style_ms.len = size_t(len(style))
	defer free(unsafe.Pointer(style_ms.data))
	return (bool)(QFontDatabase_IsSmoothlyScalable2(family_ms, style_ms))
}

func QFontDatabase_IsScalable2(family string, style string) bool {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	style_ms := struct_miqt_string{}
	style_ms.data = CString(style)
	style_ms.len = size_t(len(style))
	defer free(unsafe.Pointer(style_ms.data))
	return (bool)(QFontDatabase_IsScalable2(family_ms, style_ms))
}

func QFontDatabase_IsFixedPitch2(family string, style string) bool {
	family_ms := struct_miqt_string{}
	family_ms.data = CString(family)
	family_ms.len = size_t(len(family))
	defer free(unsafe.Pointer(family_ms.data))
	style_ms := struct_miqt_string{}
	style_ms.data = CString(style)
	style_ms.len = size_t(len(style))
	defer free(unsafe.Pointer(style_ms.data))
	return (bool)(QFontDatabase_IsFixedPitch2(family_ms, style_ms))
}
