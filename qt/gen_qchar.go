package qt

/*

#include "gen_qchar.h"
#include <stdlib.h>

*/
import "C"

import (
	"runtime"
	"unsafe"
)

type QChar__SpecialCharacter int

const (
	QChar__SpecialCharacter__Null                       QChar__SpecialCharacter = 0
	QChar__SpecialCharacter__Tabulation                 QChar__SpecialCharacter = 9
	QChar__SpecialCharacter__LineFeed                   QChar__SpecialCharacter = 10
	QChar__SpecialCharacter__FormFeed                   QChar__SpecialCharacter = 12
	QChar__SpecialCharacter__CarriageReturn             QChar__SpecialCharacter = 13
	QChar__SpecialCharacter__Space                      QChar__SpecialCharacter = 32
	QChar__SpecialCharacter__Nbsp                       QChar__SpecialCharacter = 160
	QChar__SpecialCharacter__SoftHyphen                 QChar__SpecialCharacter = 173
	QChar__SpecialCharacter__ReplacementCharacter       QChar__SpecialCharacter = 65533
	QChar__SpecialCharacter__ObjectReplacementCharacter QChar__SpecialCharacter = 65532
	QChar__SpecialCharacter__ByteOrderMark              QChar__SpecialCharacter = 65279
	QChar__SpecialCharacter__ByteOrderSwapped           QChar__SpecialCharacter = 65534
	QChar__SpecialCharacter__ParagraphSeparator         QChar__SpecialCharacter = 8233
	QChar__SpecialCharacter__LineSeparator              QChar__SpecialCharacter = 8232
	QChar__SpecialCharacter__LastValidCodePoint         QChar__SpecialCharacter = 1114111
)

type QChar__Category int

const (
	QChar__Category__Mark_NonSpacing          QChar__Category = 0
	QChar__Category__Mark_SpacingCombining    QChar__Category = 1
	QChar__Category__Mark_Enclosing           QChar__Category = 2
	QChar__Category__Number_DecimalDigit      QChar__Category = 3
	QChar__Category__Number_Letter            QChar__Category = 4
	QChar__Category__Number_Other             QChar__Category = 5
	QChar__Category__Separator_Space          QChar__Category = 6
	QChar__Category__Separator_Line           QChar__Category = 7
	QChar__Category__Separator_Paragraph      QChar__Category = 8
	QChar__Category__Other_Control            QChar__Category = 9
	QChar__Category__Other_Format             QChar__Category = 10
	QChar__Category__Other_Surrogate          QChar__Category = 11
	QChar__Category__Other_PrivateUse         QChar__Category = 12
	QChar__Category__Other_NotAssigned        QChar__Category = 13
	QChar__Category__Letter_Uppercase         QChar__Category = 14
	QChar__Category__Letter_Lowercase         QChar__Category = 15
	QChar__Category__Letter_Titlecase         QChar__Category = 16
	QChar__Category__Letter_Modifier          QChar__Category = 17
	QChar__Category__Letter_Other             QChar__Category = 18
	QChar__Category__Punctuation_Connector    QChar__Category = 19
	QChar__Category__Punctuation_Dash         QChar__Category = 20
	QChar__Category__Punctuation_Open         QChar__Category = 21
	QChar__Category__Punctuation_Close        QChar__Category = 22
	QChar__Category__Punctuation_InitialQuote QChar__Category = 23
	QChar__Category__Punctuation_FinalQuote   QChar__Category = 24
	QChar__Category__Punctuation_Other        QChar__Category = 25
	QChar__Category__Symbol_Math              QChar__Category = 26
	QChar__Category__Symbol_Currency          QChar__Category = 27
	QChar__Category__Symbol_Modifier          QChar__Category = 28
	QChar__Category__Symbol_Other             QChar__Category = 29
)

type QChar__Script int

const (
	QChar__Script__Script_Unknown               QChar__Script = 0
	QChar__Script__Script_Inherited             QChar__Script = 1
	QChar__Script__Script_Common                QChar__Script = 2
	QChar__Script__Script_Latin                 QChar__Script = 3
	QChar__Script__Script_Greek                 QChar__Script = 4
	QChar__Script__Script_Cyrillic              QChar__Script = 5
	QChar__Script__Script_Armenian              QChar__Script = 6
	QChar__Script__Script_Hebrew                QChar__Script = 7
	QChar__Script__Script_Arabic                QChar__Script = 8
	QChar__Script__Script_Syriac                QChar__Script = 9
	QChar__Script__Script_Thaana                QChar__Script = 10
	QChar__Script__Script_Devanagari            QChar__Script = 11
	QChar__Script__Script_Bengali               QChar__Script = 12
	QChar__Script__Script_Gurmukhi              QChar__Script = 13
	QChar__Script__Script_Gujarati              QChar__Script = 14
	QChar__Script__Script_Oriya                 QChar__Script = 15
	QChar__Script__Script_Tamil                 QChar__Script = 16
	QChar__Script__Script_Telugu                QChar__Script = 17
	QChar__Script__Script_Kannada               QChar__Script = 18
	QChar__Script__Script_Malayalam             QChar__Script = 19
	QChar__Script__Script_Sinhala               QChar__Script = 20
	QChar__Script__Script_Thai                  QChar__Script = 21
	QChar__Script__Script_Lao                   QChar__Script = 22
	QChar__Script__Script_Tibetan               QChar__Script = 23
	QChar__Script__Script_Myanmar               QChar__Script = 24
	QChar__Script__Script_Georgian              QChar__Script = 25
	QChar__Script__Script_Hangul                QChar__Script = 26
	QChar__Script__Script_Ethiopic              QChar__Script = 27
	QChar__Script__Script_Cherokee              QChar__Script = 28
	QChar__Script__Script_CanadianAboriginal    QChar__Script = 29
	QChar__Script__Script_Ogham                 QChar__Script = 30
	QChar__Script__Script_Runic                 QChar__Script = 31
	QChar__Script__Script_Khmer                 QChar__Script = 32
	QChar__Script__Script_Mongolian             QChar__Script = 33
	QChar__Script__Script_Hiragana              QChar__Script = 34
	QChar__Script__Script_Katakana              QChar__Script = 35
	QChar__Script__Script_Bopomofo              QChar__Script = 36
	QChar__Script__Script_Han                   QChar__Script = 37
	QChar__Script__Script_Yi                    QChar__Script = 38
	QChar__Script__Script_OldItalic             QChar__Script = 39
	QChar__Script__Script_Gothic                QChar__Script = 40
	QChar__Script__Script_Deseret               QChar__Script = 41
	QChar__Script__Script_Tagalog               QChar__Script = 42
	QChar__Script__Script_Hanunoo               QChar__Script = 43
	QChar__Script__Script_Buhid                 QChar__Script = 44
	QChar__Script__Script_Tagbanwa              QChar__Script = 45
	QChar__Script__Script_Coptic                QChar__Script = 46
	QChar__Script__Script_Limbu                 QChar__Script = 47
	QChar__Script__Script_TaiLe                 QChar__Script = 48
	QChar__Script__Script_LinearB               QChar__Script = 49
	QChar__Script__Script_Ugaritic              QChar__Script = 50
	QChar__Script__Script_Shavian               QChar__Script = 51
	QChar__Script__Script_Osmanya               QChar__Script = 52
	QChar__Script__Script_Cypriot               QChar__Script = 53
	QChar__Script__Script_Braille               QChar__Script = 54
	QChar__Script__Script_Buginese              QChar__Script = 55
	QChar__Script__Script_NewTaiLue             QChar__Script = 56
	QChar__Script__Script_Glagolitic            QChar__Script = 57
	QChar__Script__Script_Tifinagh              QChar__Script = 58
	QChar__Script__Script_SylotiNagri           QChar__Script = 59
	QChar__Script__Script_OldPersian            QChar__Script = 60
	QChar__Script__Script_Kharoshthi            QChar__Script = 61
	QChar__Script__Script_Balinese              QChar__Script = 62
	QChar__Script__Script_Cuneiform             QChar__Script = 63
	QChar__Script__Script_Phoenician            QChar__Script = 64
	QChar__Script__Script_PhagsPa               QChar__Script = 65
	QChar__Script__Script_Nko                   QChar__Script = 66
	QChar__Script__Script_Sundanese             QChar__Script = 67
	QChar__Script__Script_Lepcha                QChar__Script = 68
	QChar__Script__Script_OlChiki               QChar__Script = 69
	QChar__Script__Script_Vai                   QChar__Script = 70
	QChar__Script__Script_Saurashtra            QChar__Script = 71
	QChar__Script__Script_KayahLi               QChar__Script = 72
	QChar__Script__Script_Rejang                QChar__Script = 73
	QChar__Script__Script_Lycian                QChar__Script = 74
	QChar__Script__Script_Carian                QChar__Script = 75
	QChar__Script__Script_Lydian                QChar__Script = 76
	QChar__Script__Script_Cham                  QChar__Script = 77
	QChar__Script__Script_TaiTham               QChar__Script = 78
	QChar__Script__Script_TaiViet               QChar__Script = 79
	QChar__Script__Script_Avestan               QChar__Script = 80
	QChar__Script__Script_EgyptianHieroglyphs   QChar__Script = 81
	QChar__Script__Script_Samaritan             QChar__Script = 82
	QChar__Script__Script_Lisu                  QChar__Script = 83
	QChar__Script__Script_Bamum                 QChar__Script = 84
	QChar__Script__Script_Javanese              QChar__Script = 85
	QChar__Script__Script_MeeteiMayek           QChar__Script = 86
	QChar__Script__Script_ImperialAramaic       QChar__Script = 87
	QChar__Script__Script_OldSouthArabian       QChar__Script = 88
	QChar__Script__Script_InscriptionalParthian QChar__Script = 89
	QChar__Script__Script_InscriptionalPahlavi  QChar__Script = 90
	QChar__Script__Script_OldTurkic             QChar__Script = 91
	QChar__Script__Script_Kaithi                QChar__Script = 92
	QChar__Script__Script_Batak                 QChar__Script = 93
	QChar__Script__Script_Brahmi                QChar__Script = 94
	QChar__Script__Script_Mandaic               QChar__Script = 95
	QChar__Script__Script_Chakma                QChar__Script = 96
	QChar__Script__Script_MeroiticCursive       QChar__Script = 97
	QChar__Script__Script_MeroiticHieroglyphs   QChar__Script = 98
	QChar__Script__Script_Miao                  QChar__Script = 99
	QChar__Script__Script_Sharada               QChar__Script = 100
	QChar__Script__Script_SoraSompeng           QChar__Script = 101
	QChar__Script__Script_Takri                 QChar__Script = 102
	QChar__Script__Script_CaucasianAlbanian     QChar__Script = 103
	QChar__Script__Script_BassaVah              QChar__Script = 104
	QChar__Script__Script_Duployan              QChar__Script = 105
	QChar__Script__Script_Elbasan               QChar__Script = 106
	QChar__Script__Script_Grantha               QChar__Script = 107
	QChar__Script__Script_PahawhHmong           QChar__Script = 108
	QChar__Script__Script_Khojki                QChar__Script = 109
	QChar__Script__Script_LinearA               QChar__Script = 110
	QChar__Script__Script_Mahajani              QChar__Script = 111
	QChar__Script__Script_Manichaean            QChar__Script = 112
	QChar__Script__Script_MendeKikakui          QChar__Script = 113
	QChar__Script__Script_Modi                  QChar__Script = 114
	QChar__Script__Script_Mro                   QChar__Script = 115
	QChar__Script__Script_OldNorthArabian       QChar__Script = 116
	QChar__Script__Script_Nabataean             QChar__Script = 117
	QChar__Script__Script_Palmyrene             QChar__Script = 118
	QChar__Script__Script_PauCinHau             QChar__Script = 119
	QChar__Script__Script_OldPermic             QChar__Script = 120
	QChar__Script__Script_PsalterPahlavi        QChar__Script = 121
	QChar__Script__Script_Siddham               QChar__Script = 122
	QChar__Script__Script_Khudawadi             QChar__Script = 123
	QChar__Script__Script_Tirhuta               QChar__Script = 124
	QChar__Script__Script_WarangCiti            QChar__Script = 125
	QChar__Script__Script_Ahom                  QChar__Script = 126
	QChar__Script__Script_AnatolianHieroglyphs  QChar__Script = 127
	QChar__Script__Script_Hatran                QChar__Script = 128
	QChar__Script__Script_Multani               QChar__Script = 129
	QChar__Script__Script_OldHungarian          QChar__Script = 130
	QChar__Script__Script_SignWriting           QChar__Script = 131
	QChar__Script__Script_Adlam                 QChar__Script = 132
	QChar__Script__Script_Bhaiksuki             QChar__Script = 133
	QChar__Script__Script_Marchen               QChar__Script = 134
	QChar__Script__Script_Newa                  QChar__Script = 135
	QChar__Script__Script_Osage                 QChar__Script = 136
	QChar__Script__Script_Tangut                QChar__Script = 137
	QChar__Script__Script_MasaramGondi          QChar__Script = 138
	QChar__Script__Script_Nushu                 QChar__Script = 139
	QChar__Script__Script_Soyombo               QChar__Script = 140
	QChar__Script__Script_ZanabazarSquare       QChar__Script = 141
	QChar__Script__Script_Dogra                 QChar__Script = 142
	QChar__Script__Script_GunjalaGondi          QChar__Script = 143
	QChar__Script__Script_HanifiRohingya        QChar__Script = 144
	QChar__Script__Script_Makasar               QChar__Script = 145
	QChar__Script__Script_Medefaidrin           QChar__Script = 146
	QChar__Script__Script_OldSogdian            QChar__Script = 147
	QChar__Script__Script_Sogdian               QChar__Script = 148
	QChar__Script__Script_Elymaic               QChar__Script = 149
	QChar__Script__Script_Nandinagari           QChar__Script = 150
	QChar__Script__Script_NyiakengPuachueHmong  QChar__Script = 151
	QChar__Script__Script_Wancho                QChar__Script = 152
	QChar__Script__Script_Chorasmian            QChar__Script = 153
	QChar__Script__Script_DivesAkuru            QChar__Script = 154
	QChar__Script__Script_KhitanSmallScript     QChar__Script = 155
	QChar__Script__Script_Yezidi                QChar__Script = 156
	QChar__Script__ScriptCount                  QChar__Script = 157
)

type QChar__Direction int

const (
	QChar__Direction__DirL   QChar__Direction = 0
	QChar__Direction__DirR   QChar__Direction = 1
	QChar__Direction__DirEN  QChar__Direction = 2
	QChar__Direction__DirES  QChar__Direction = 3
	QChar__Direction__DirET  QChar__Direction = 4
	QChar__Direction__DirAN  QChar__Direction = 5
	QChar__Direction__DirCS  QChar__Direction = 6
	QChar__Direction__DirB   QChar__Direction = 7
	QChar__Direction__DirS   QChar__Direction = 8
	QChar__Direction__DirWS  QChar__Direction = 9
	QChar__Direction__DirON  QChar__Direction = 10
	QChar__Direction__DirLRE QChar__Direction = 11
	QChar__Direction__DirLRO QChar__Direction = 12
	QChar__Direction__DirAL  QChar__Direction = 13
	QChar__Direction__DirRLE QChar__Direction = 14
	QChar__Direction__DirRLO QChar__Direction = 15
	QChar__Direction__DirPDF QChar__Direction = 16
	QChar__Direction__DirNSM QChar__Direction = 17
	QChar__Direction__DirBN  QChar__Direction = 18
	QChar__Direction__DirLRI QChar__Direction = 19
	QChar__Direction__DirRLI QChar__Direction = 20
	QChar__Direction__DirFSI QChar__Direction = 21
	QChar__Direction__DirPDI QChar__Direction = 22
)

type QChar__Decomposition int

const (
	QChar__Decomposition__NoDecomposition QChar__Decomposition = 0
	QChar__Decomposition__Canonical       QChar__Decomposition = 1
	QChar__Decomposition__Font            QChar__Decomposition = 2
	QChar__Decomposition__NoBreak         QChar__Decomposition = 3
	QChar__Decomposition__Initial         QChar__Decomposition = 4
	QChar__Decomposition__Medial          QChar__Decomposition = 5
	QChar__Decomposition__Final           QChar__Decomposition = 6
	QChar__Decomposition__Isolated        QChar__Decomposition = 7
	QChar__Decomposition__Circle          QChar__Decomposition = 8
	QChar__Decomposition__Super           QChar__Decomposition = 9
	QChar__Decomposition__Sub             QChar__Decomposition = 10
	QChar__Decomposition__Vertical        QChar__Decomposition = 11
	QChar__Decomposition__Wide            QChar__Decomposition = 12
	QChar__Decomposition__Narrow          QChar__Decomposition = 13
	QChar__Decomposition__Small           QChar__Decomposition = 14
	QChar__Decomposition__Square          QChar__Decomposition = 15
	QChar__Decomposition__Compat          QChar__Decomposition = 16
	QChar__Decomposition__Fraction        QChar__Decomposition = 17
)

type QChar__JoiningType int

const (
	QChar__JoiningType__Joining_None        QChar__JoiningType = 0
	QChar__JoiningType__Joining_Causing     QChar__JoiningType = 1
	QChar__JoiningType__Joining_Dual        QChar__JoiningType = 2
	QChar__JoiningType__Joining_Right       QChar__JoiningType = 3
	QChar__JoiningType__Joining_Left        QChar__JoiningType = 4
	QChar__JoiningType__Joining_Transparent QChar__JoiningType = 5
)

type QChar__Joining int

const (
	QChar__Joining__OtherJoining QChar__Joining = 0
	QChar__Joining__Dual         QChar__Joining = 1
	QChar__Joining__Right        QChar__Joining = 2
	QChar__Joining__Center       QChar__Joining = 3
)

type QChar__CombiningClass int

const (
	QChar__CombiningClass__Combining_BelowLeftAttached  QChar__CombiningClass = 200
	QChar__CombiningClass__Combining_BelowAttached      QChar__CombiningClass = 202
	QChar__CombiningClass__Combining_BelowRightAttached QChar__CombiningClass = 204
	QChar__CombiningClass__Combining_LeftAttached       QChar__CombiningClass = 208
	QChar__CombiningClass__Combining_RightAttached      QChar__CombiningClass = 210
	QChar__CombiningClass__Combining_AboveLeftAttached  QChar__CombiningClass = 212
	QChar__CombiningClass__Combining_AboveAttached      QChar__CombiningClass = 214
	QChar__CombiningClass__Combining_AboveRightAttached QChar__CombiningClass = 216
	QChar__CombiningClass__Combining_BelowLeft          QChar__CombiningClass = 218
	QChar__CombiningClass__Combining_Below              QChar__CombiningClass = 220
	QChar__CombiningClass__Combining_BelowRight         QChar__CombiningClass = 222
	QChar__CombiningClass__Combining_Left               QChar__CombiningClass = 224
	QChar__CombiningClass__Combining_Right              QChar__CombiningClass = 226
	QChar__CombiningClass__Combining_AboveLeft          QChar__CombiningClass = 228
	QChar__CombiningClass__Combining_Above              QChar__CombiningClass = 230
	QChar__CombiningClass__Combining_AboveRight         QChar__CombiningClass = 232
	QChar__CombiningClass__Combining_DoubleBelow        QChar__CombiningClass = 233
	QChar__CombiningClass__Combining_DoubleAbove        QChar__CombiningClass = 234
	QChar__CombiningClass__Combining_IotaSubscript      QChar__CombiningClass = 240
)

type QChar__UnicodeVersion int

const (
	QChar__UnicodeVersion__Unicode_Unassigned QChar__UnicodeVersion = 0
	QChar__UnicodeVersion__Unicode_1_1        QChar__UnicodeVersion = 1
	QChar__UnicodeVersion__Unicode_2_0        QChar__UnicodeVersion = 2
	QChar__UnicodeVersion__Unicode_2_1_2      QChar__UnicodeVersion = 3
	QChar__UnicodeVersion__Unicode_3_0        QChar__UnicodeVersion = 4
	QChar__UnicodeVersion__Unicode_3_1        QChar__UnicodeVersion = 5
	QChar__UnicodeVersion__Unicode_3_2        QChar__UnicodeVersion = 6
	QChar__UnicodeVersion__Unicode_4_0        QChar__UnicodeVersion = 7
	QChar__UnicodeVersion__Unicode_4_1        QChar__UnicodeVersion = 8
	QChar__UnicodeVersion__Unicode_5_0        QChar__UnicodeVersion = 9
	QChar__UnicodeVersion__Unicode_5_1        QChar__UnicodeVersion = 10
	QChar__UnicodeVersion__Unicode_5_2        QChar__UnicodeVersion = 11
	QChar__UnicodeVersion__Unicode_6_0        QChar__UnicodeVersion = 12
	QChar__UnicodeVersion__Unicode_6_1        QChar__UnicodeVersion = 13
	QChar__UnicodeVersion__Unicode_6_2        QChar__UnicodeVersion = 14
	QChar__UnicodeVersion__Unicode_6_3        QChar__UnicodeVersion = 15
	QChar__UnicodeVersion__Unicode_7_0        QChar__UnicodeVersion = 16
	QChar__UnicodeVersion__Unicode_8_0        QChar__UnicodeVersion = 17
	QChar__UnicodeVersion__Unicode_9_0        QChar__UnicodeVersion = 18
	QChar__UnicodeVersion__Unicode_10_0       QChar__UnicodeVersion = 19
	QChar__UnicodeVersion__Unicode_11_0       QChar__UnicodeVersion = 20
	QChar__UnicodeVersion__Unicode_12_0       QChar__UnicodeVersion = 21
	QChar__UnicodeVersion__Unicode_12_1       QChar__UnicodeVersion = 22
	QChar__UnicodeVersion__Unicode_13_0       QChar__UnicodeVersion = 23
)

type QLatin1Char struct {
	h *C.QLatin1Char
}

func (this *QLatin1Char) cPointer() *C.QLatin1Char {
	if this == nil {
		return nil
	}
	return this.h
}

func newQLatin1Char(h *C.QLatin1Char) *QLatin1Char {
	if h == nil {
		return nil
	}
	return &QLatin1Char{h: h}
}

func newQLatin1Char_U(h unsafe.Pointer) *QLatin1Char {
	return newQLatin1Char((*C.QLatin1Char)(h))
}

// NewQLatin1Char constructs a new QLatin1Char object.
func NewQLatin1Char(c byte) *QLatin1Char {
	ret := C.QLatin1Char_new((C.char)(c))
	return newQLatin1Char(ret)
}

// NewQLatin1Char2 constructs a new QLatin1Char object.
func NewQLatin1Char2(param1 *QLatin1Char) *QLatin1Char {
	ret := C.QLatin1Char_new2(param1.cPointer())
	return newQLatin1Char(ret)
}

func (this *QLatin1Char) ToLatin1() byte {
	ret := C.QLatin1Char_ToLatin1(this.h)
	return (byte)(ret)
}

func (this *QLatin1Char) Unicode() uint16 {
	ret := C.QLatin1Char_Unicode(this.h)
	return (uint16)(ret)
}

func (this *QLatin1Char) Delete() {
	C.QLatin1Char_Delete(this.h)
}

type QChar struct {
	h *C.QChar
}

func (this *QChar) cPointer() *C.QChar {
	if this == nil {
		return nil
	}
	return this.h
}

func newQChar(h *C.QChar) *QChar {
	if h == nil {
		return nil
	}
	return &QChar{h: h}
}

func newQChar_U(h unsafe.Pointer) *QChar {
	return newQChar((*C.QChar)(h))
}

// NewQChar constructs a new QChar object.
func NewQChar() *QChar {
	ret := C.QChar_new()
	return newQChar(ret)
}

// NewQChar2 constructs a new QChar object.
func NewQChar2(rc uint16) *QChar {
	ret := C.QChar_new2((C.uint16_t)(rc))
	return newQChar(ret)
}

// NewQChar3 constructs a new QChar object.
func NewQChar3(c byte, r byte) *QChar {
	ret := C.QChar_new3((C.uchar)(c), (C.uchar)(r))
	return newQChar(ret)
}

// NewQChar4 constructs a new QChar object.
func NewQChar4(rc int16) *QChar {
	ret := C.QChar_new4((C.int16_t)(rc))
	return newQChar(ret)
}

// NewQChar5 constructs a new QChar object.
func NewQChar5(rc uint) *QChar {
	ret := C.QChar_new5((C.uint)(rc))
	return newQChar(ret)
}

// NewQChar6 constructs a new QChar object.
func NewQChar6(rc int) *QChar {
	ret := C.QChar_new6((C.int)(rc))
	return newQChar(ret)
}

// NewQChar7 constructs a new QChar object.
func NewQChar7(s QChar__SpecialCharacter) *QChar {
	ret := C.QChar_new7((C.uintptr_t)(s))
	return newQChar(ret)
}

// NewQChar8 constructs a new QChar object.
func NewQChar8(ch QLatin1Char) *QChar {
	ret := C.QChar_new8(ch.cPointer())
	return newQChar(ret)
}

// NewQChar9 constructs a new QChar object.
func NewQChar9(c byte) *QChar {
	ret := C.QChar_new9((C.char)(c))
	return newQChar(ret)
}

// NewQChar10 constructs a new QChar object.
func NewQChar10(c byte) *QChar {
	ret := C.QChar_new10((C.uchar)(c))
	return newQChar(ret)
}

// NewQChar11 constructs a new QChar object.
func NewQChar11(param1 *QChar) *QChar {
	ret := C.QChar_new11(param1.cPointer())
	return newQChar(ret)
}

func (this *QChar) Category() QChar__Category {
	ret := C.QChar_Category(this.h)
	return (QChar__Category)(ret)
}

func (this *QChar) Direction() QChar__Direction {
	ret := C.QChar_Direction(this.h)
	return (QChar__Direction)(ret)
}

func (this *QChar) JoiningType() QChar__JoiningType {
	ret := C.QChar_JoiningType(this.h)
	return (QChar__JoiningType)(ret)
}

func (this *QChar) Joining() QChar__Joining {
	ret := C.QChar_Joining(this.h)
	return (QChar__Joining)(ret)
}

func (this *QChar) CombiningClass() byte {
	ret := C.QChar_CombiningClass(this.h)
	return (byte)(ret)
}

func (this *QChar) MirroredChar() *QChar {
	ret := C.QChar_MirroredChar(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQChar(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QChar) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QChar) HasMirrored() bool {
	ret := C.QChar_HasMirrored(this.h)
	return (bool)(ret)
}

func (this *QChar) Decomposition() string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QChar_Decomposition(this.h, &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func (this *QChar) DecompositionTag() QChar__Decomposition {
	ret := C.QChar_DecompositionTag(this.h)
	return (QChar__Decomposition)(ret)
}

func (this *QChar) DigitValue() int {
	ret := C.QChar_DigitValue(this.h)
	return (int)(ret)
}

func (this *QChar) ToLower() *QChar {
	ret := C.QChar_ToLower(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQChar(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QChar) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QChar) ToUpper() *QChar {
	ret := C.QChar_ToUpper(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQChar(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QChar) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QChar) ToTitleCase() *QChar {
	ret := C.QChar_ToTitleCase(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQChar(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QChar) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QChar) ToCaseFolded() *QChar {
	ret := C.QChar_ToCaseFolded(this.h)
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQChar(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QChar) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QChar) Script() QChar__Script {
	ret := C.QChar_Script(this.h)
	return (QChar__Script)(ret)
}

func (this *QChar) UnicodeVersion() QChar__UnicodeVersion {
	ret := C.QChar_UnicodeVersion(this.h)
	return (QChar__UnicodeVersion)(ret)
}

func (this *QChar) ToLatin1() byte {
	ret := C.QChar_ToLatin1(this.h)
	return (byte)(ret)
}

func (this *QChar) Unicode() uint16 {
	ret := C.QChar_Unicode(this.h)
	return (uint16)(ret)
}

func QChar_FromLatin1(c byte) *QChar {
	ret := C.QChar_FromLatin1((C.char)(c))
	// Qt uses pass-by-value semantics for this type. Mimic with finalizer
	ret1 := newQChar(ret)
	runtime.SetFinalizer(ret1, func(ret2 *QChar) {
		ret2.Delete()
		runtime.KeepAlive(ret2.h)
	})
	return ret1
}

func (this *QChar) IsNull() bool {
	ret := C.QChar_IsNull(this.h)
	return (bool)(ret)
}

func (this *QChar) IsPrint() bool {
	ret := C.QChar_IsPrint(this.h)
	return (bool)(ret)
}

func (this *QChar) IsSpace() bool {
	ret := C.QChar_IsSpace(this.h)
	return (bool)(ret)
}

func (this *QChar) IsMark() bool {
	ret := C.QChar_IsMark(this.h)
	return (bool)(ret)
}

func (this *QChar) IsPunct() bool {
	ret := C.QChar_IsPunct(this.h)
	return (bool)(ret)
}

func (this *QChar) IsSymbol() bool {
	ret := C.QChar_IsSymbol(this.h)
	return (bool)(ret)
}

func (this *QChar) IsLetter() bool {
	ret := C.QChar_IsLetter(this.h)
	return (bool)(ret)
}

func (this *QChar) IsNumber() bool {
	ret := C.QChar_IsNumber(this.h)
	return (bool)(ret)
}

func (this *QChar) IsLetterOrNumber() bool {
	ret := C.QChar_IsLetterOrNumber(this.h)
	return (bool)(ret)
}

func (this *QChar) IsDigit() bool {
	ret := C.QChar_IsDigit(this.h)
	return (bool)(ret)
}

func (this *QChar) IsLower() bool {
	ret := C.QChar_IsLower(this.h)
	return (bool)(ret)
}

func (this *QChar) IsUpper() bool {
	ret := C.QChar_IsUpper(this.h)
	return (bool)(ret)
}

func (this *QChar) IsTitleCase() bool {
	ret := C.QChar_IsTitleCase(this.h)
	return (bool)(ret)
}

func (this *QChar) IsNonCharacter() bool {
	ret := C.QChar_IsNonCharacter(this.h)
	return (bool)(ret)
}

func (this *QChar) IsHighSurrogate() bool {
	ret := C.QChar_IsHighSurrogate(this.h)
	return (bool)(ret)
}

func (this *QChar) IsLowSurrogate() bool {
	ret := C.QChar_IsLowSurrogate(this.h)
	return (bool)(ret)
}

func (this *QChar) IsSurrogate() bool {
	ret := C.QChar_IsSurrogate(this.h)
	return (bool)(ret)
}

func (this *QChar) Cell() byte {
	ret := C.QChar_Cell(this.h)
	return (byte)(ret)
}

func (this *QChar) Row() byte {
	ret := C.QChar_Row(this.h)
	return (byte)(ret)
}

func (this *QChar) SetCell(acell byte) {
	C.QChar_SetCell(this.h, (C.uchar)(acell))
}

func (this *QChar) SetRow(arow byte) {
	C.QChar_SetRow(this.h, (C.uchar)(arow))
}

func QChar_IsNonCharacterWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsNonCharacterWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsHighSurrogateWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsHighSurrogateWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsLowSurrogateWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsLowSurrogateWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsSurrogateWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsSurrogateWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_RequiresSurrogates(ucs4 uint) bool {
	ret := C.QChar_RequiresSurrogates((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_SurrogateToUcs4(high uint16, low uint16) uint {
	ret := C.QChar_SurrogateToUcs4((C.uint16_t)(high), (C.uint16_t)(low))
	return (uint)(ret)
}

func QChar_SurrogateToUcs42(high QChar, low QChar) uint {
	ret := C.QChar_SurrogateToUcs42(high.cPointer(), low.cPointer())
	return (uint)(ret)
}

func QChar_HighSurrogate(ucs4 uint) uint16 {
	ret := C.QChar_HighSurrogate((C.uint)(ucs4))
	return (uint16)(ret)
}

func QChar_LowSurrogate(ucs4 uint) uint16 {
	ret := C.QChar_LowSurrogate((C.uint)(ucs4))
	return (uint16)(ret)
}

func QChar_CategoryWithUcs4(ucs4 uint) QChar__Category {
	ret := C.QChar_CategoryWithUcs4((C.uint)(ucs4))
	return (QChar__Category)(ret)
}

func QChar_DirectionWithUcs4(ucs4 uint) QChar__Direction {
	ret := C.QChar_DirectionWithUcs4((C.uint)(ucs4))
	return (QChar__Direction)(ret)
}

func QChar_JoiningTypeWithUcs4(ucs4 uint) QChar__JoiningType {
	ret := C.QChar_JoiningTypeWithUcs4((C.uint)(ucs4))
	return (QChar__JoiningType)(ret)
}

func QChar_JoiningWithUcs4(ucs4 uint) QChar__Joining {
	ret := C.QChar_JoiningWithUcs4((C.uint)(ucs4))
	return (QChar__Joining)(ret)
}

func QChar_CombiningClassWithUcs4(ucs4 uint) byte {
	ret := C.QChar_CombiningClassWithUcs4((C.uint)(ucs4))
	return (byte)(ret)
}

func QChar_MirroredCharWithUcs4(ucs4 uint) uint {
	ret := C.QChar_MirroredCharWithUcs4((C.uint)(ucs4))
	return (uint)(ret)
}

func QChar_HasMirroredWithUcs4(ucs4 uint) bool {
	ret := C.QChar_HasMirroredWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_DecompositionWithUcs4(ucs4 uint) string {
	var _out *C.char = nil
	var _out_Strlen C.int = 0
	C.QChar_DecompositionWithUcs4((C.uint)(ucs4), &_out, &_out_Strlen)
	ret := C.GoStringN(_out, _out_Strlen)
	C.free(unsafe.Pointer(_out))
	return ret
}

func QChar_DecompositionTagWithUcs4(ucs4 uint) QChar__Decomposition {
	ret := C.QChar_DecompositionTagWithUcs4((C.uint)(ucs4))
	return (QChar__Decomposition)(ret)
}

func QChar_DigitValueWithUcs4(ucs4 uint) int {
	ret := C.QChar_DigitValueWithUcs4((C.uint)(ucs4))
	return (int)(ret)
}

func QChar_ToLowerWithUcs4(ucs4 uint) uint {
	ret := C.QChar_ToLowerWithUcs4((C.uint)(ucs4))
	return (uint)(ret)
}

func QChar_ToUpperWithUcs4(ucs4 uint) uint {
	ret := C.QChar_ToUpperWithUcs4((C.uint)(ucs4))
	return (uint)(ret)
}

func QChar_ToTitleCaseWithUcs4(ucs4 uint) uint {
	ret := C.QChar_ToTitleCaseWithUcs4((C.uint)(ucs4))
	return (uint)(ret)
}

func QChar_ToCaseFoldedWithUcs4(ucs4 uint) uint {
	ret := C.QChar_ToCaseFoldedWithUcs4((C.uint)(ucs4))
	return (uint)(ret)
}

func QChar_ScriptWithUcs4(ucs4 uint) QChar__Script {
	ret := C.QChar_ScriptWithUcs4((C.uint)(ucs4))
	return (QChar__Script)(ret)
}

func QChar_UnicodeVersionWithUcs4(ucs4 uint) QChar__UnicodeVersion {
	ret := C.QChar_UnicodeVersionWithUcs4((C.uint)(ucs4))
	return (QChar__UnicodeVersion)(ret)
}

func QChar_CurrentUnicodeVersion() QChar__UnicodeVersion {
	ret := C.QChar_CurrentUnicodeVersion()
	return (QChar__UnicodeVersion)(ret)
}

func QChar_IsPrintWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsPrintWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsSpaceWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsSpaceWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsMarkWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsMarkWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsPunctWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsPunctWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsSymbolWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsSymbolWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsLetterWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsLetterWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsNumberWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsNumberWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsLetterOrNumberWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsLetterOrNumberWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsDigitWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsDigitWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsLowerWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsLowerWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsUpperWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsUpperWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func QChar_IsTitleCaseWithUcs4(ucs4 uint) bool {
	ret := C.QChar_IsTitleCaseWithUcs4((C.uint)(ucs4))
	return (bool)(ret)
}

func (this *QChar) Delete() {
	C.QChar_Delete(this.h)
}
