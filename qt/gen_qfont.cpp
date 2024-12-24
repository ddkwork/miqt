// +build ignore

#include <QByteArray>
#include <QFont>
#define WORKAROUND_INNER_CLASS_DEFINITION_QFont__Tag
#include <QList>
#include <QPaintDevice>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qfont.h>
#include "gen_qfont.h"

#ifndef _Bool
#define _Bool bool
#endif

QFont* QFont_new() {
	return new QFont();
}

QFont* QFont_new2(struct miqt_string family) {
	QString family_QString = QString::fromUtf8(family.data, family.len);
	return new QFont(family_QString);
}

QFont* QFont_new3(struct miqt_array /* of struct miqt_string */  families) {
	QStringList families_QList;
	families_QList.reserve(families.len);
	struct miqt_string* families_arr = static_cast<struct miqt_string*>(families.data);
	for(size_t i = 0; i < families.len; ++i) {
		QString families_arr_i_QString = QString::fromUtf8(families_arr[i].data, families_arr[i].len);
		families_QList.push_back(families_arr_i_QString);
	}
	return new QFont(families_QList);
}

QFont* QFont_new4(QFont* font, QPaintDevice* pd) {
	return new QFont(*font, pd);
}

QFont* QFont_new5(QFont* font) {
	return new QFont(*font);
}

QFont* QFont_new6(struct miqt_string family, int pointSize) {
	QString family_QString = QString::fromUtf8(family.data, family.len);
	return new QFont(family_QString, static_cast<int>(pointSize));
}

QFont* QFont_new7(struct miqt_string family, int pointSize, int weight) {
	QString family_QString = QString::fromUtf8(family.data, family.len);
	return new QFont(family_QString, static_cast<int>(pointSize), static_cast<int>(weight));
}

QFont* QFont_new8(struct miqt_string family, int pointSize, int weight, bool italic) {
	QString family_QString = QString::fromUtf8(family.data, family.len);
	return new QFont(family_QString, static_cast<int>(pointSize), static_cast<int>(weight), italic);
}

QFont* QFont_new9(struct miqt_array /* of struct miqt_string */  families, int pointSize) {
	QStringList families_QList;
	families_QList.reserve(families.len);
	struct miqt_string* families_arr = static_cast<struct miqt_string*>(families.data);
	for(size_t i = 0; i < families.len; ++i) {
		QString families_arr_i_QString = QString::fromUtf8(families_arr[i].data, families_arr[i].len);
		families_QList.push_back(families_arr_i_QString);
	}
	return new QFont(families_QList, static_cast<int>(pointSize));
}

QFont* QFont_new10(struct miqt_array /* of struct miqt_string */  families, int pointSize, int weight) {
	QStringList families_QList;
	families_QList.reserve(families.len);
	struct miqt_string* families_arr = static_cast<struct miqt_string*>(families.data);
	for(size_t i = 0; i < families.len; ++i) {
		QString families_arr_i_QString = QString::fromUtf8(families_arr[i].data, families_arr[i].len);
		families_QList.push_back(families_arr_i_QString);
	}
	return new QFont(families_QList, static_cast<int>(pointSize), static_cast<int>(weight));
}

QFont* QFont_new11(struct miqt_array /* of struct miqt_string */  families, int pointSize, int weight, bool italic) {
	QStringList families_QList;
	families_QList.reserve(families.len);
	struct miqt_string* families_arr = static_cast<struct miqt_string*>(families.data);
	for(size_t i = 0; i < families.len; ++i) {
		QString families_arr_i_QString = QString::fromUtf8(families_arr[i].data, families_arr[i].len);
		families_QList.push_back(families_arr_i_QString);
	}
	return new QFont(families_QList, static_cast<int>(pointSize), static_cast<int>(weight), italic);
}

void QFont_Swap(QFont* self, QFont* other) {
	self->swap(*other);
}

struct miqt_string QFont_Family(const QFont* self) {
	QString _ret = self->family();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QFont_SetFamily(QFont* self, struct miqt_string family) {
	QString family_QString = QString::fromUtf8(family.data, family.len);
	self->setFamily(family_QString);
}

struct miqt_array /* of struct miqt_string */  QFont_Families(const QFont* self) {
	QStringList _ret = self->families();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QFont_SetFamilies(QFont* self, struct miqt_array /* of struct miqt_string */  families) {
	QStringList families_QList;
	families_QList.reserve(families.len);
	struct miqt_string* families_arr = static_cast<struct miqt_string*>(families.data);
	for(size_t i = 0; i < families.len; ++i) {
		QString families_arr_i_QString = QString::fromUtf8(families_arr[i].data, families_arr[i].len);
		families_QList.push_back(families_arr_i_QString);
	}
	self->setFamilies(families_QList);
}

struct miqt_string QFont_StyleName(const QFont* self) {
	QString _ret = self->styleName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QFont_SetStyleName(QFont* self, struct miqt_string styleName) {
	QString styleName_QString = QString::fromUtf8(styleName.data, styleName.len);
	self->setStyleName(styleName_QString);
}

int QFont_PointSize(const QFont* self) {
	return self->pointSize();
}

void QFont_SetPointSize(QFont* self, int pointSize) {
	self->setPointSize(static_cast<int>(pointSize));
}

double QFont_PointSizeF(const QFont* self) {
	qreal _ret = self->pointSizeF();
	return static_cast<double>(_ret);
}

void QFont_SetPointSizeF(QFont* self, double pointSizeF) {
	self->setPointSizeF(static_cast<qreal>(pointSizeF));
}

int QFont_PixelSize(const QFont* self) {
	return self->pixelSize();
}

void QFont_SetPixelSize(QFont* self, int pixelSize) {
	self->setPixelSize(static_cast<int>(pixelSize));
}

Weight QFont_Weight(const QFont* self) {
	return self->weight();
}

void QFont_SetWeight(QFont* self, Weight weight) {
	self->setWeight(weight);
}

bool QFont_Bold(const QFont* self) {
	return self->bold();
}

void QFont_SetBold(QFont* self, bool bold) {
	self->setBold(bold);
}

void QFont_SetStyle(QFont* self, Style style) {
	self->setStyle(style);
}

Style QFont_Style(const QFont* self) {
	return self->style();
}

bool QFont_Italic(const QFont* self) {
	return self->italic();
}

void QFont_SetItalic(QFont* self, bool b) {
	self->setItalic(b);
}

bool QFont_Underline(const QFont* self) {
	return self->underline();
}

void QFont_SetUnderline(QFont* self, bool underline) {
	self->setUnderline(underline);
}

bool QFont_Overline(const QFont* self) {
	return self->overline();
}

void QFont_SetOverline(QFont* self, bool overline) {
	self->setOverline(overline);
}

bool QFont_StrikeOut(const QFont* self) {
	return self->strikeOut();
}

void QFont_SetStrikeOut(QFont* self, bool strikeOut) {
	self->setStrikeOut(strikeOut);
}

bool QFont_FixedPitch(const QFont* self) {
	return self->fixedPitch();
}

void QFont_SetFixedPitch(QFont* self, bool fixedPitch) {
	self->setFixedPitch(fixedPitch);
}

bool QFont_Kerning(const QFont* self) {
	return self->kerning();
}

void QFont_SetKerning(QFont* self, bool kerning) {
	self->setKerning(kerning);
}

StyleHint QFont_StyleHint(const QFont* self) {
	return self->styleHint();
}

StyleStrategy QFont_StyleStrategy(const QFont* self) {
	return self->styleStrategy();
}

void QFont_SetStyleHint(QFont* self, StyleHint param1) {
	self->setStyleHint(param1);
}

void QFont_SetStyleStrategy(QFont* self, StyleStrategy s) {
	self->setStyleStrategy(s);
}

int QFont_Stretch(const QFont* self) {
	return self->stretch();
}

void QFont_SetStretch(QFont* self, int stretch) {
	self->setStretch(static_cast<int>(stretch));
}

double QFont_LetterSpacing(const QFont* self) {
	qreal _ret = self->letterSpacing();
	return static_cast<double>(_ret);
}

SpacingType QFont_LetterSpacingType(const QFont* self) {
	return self->letterSpacingType();
}

void QFont_SetLetterSpacing(QFont* self, SpacingType typeVal, double spacing) {
	self->setLetterSpacing(typeVal, static_cast<qreal>(spacing));
}

double QFont_WordSpacing(const QFont* self) {
	qreal _ret = self->wordSpacing();
	return static_cast<double>(_ret);
}

void QFont_SetWordSpacing(QFont* self, double spacing) {
	self->setWordSpacing(static_cast<qreal>(spacing));
}

void QFont_SetCapitalization(QFont* self, Capitalization capitalization) {
	self->setCapitalization(capitalization);
}

Capitalization QFont_Capitalization(const QFont* self) {
	return self->capitalization();
}

void QFont_SetHintingPreference(QFont* self, HintingPreference hintingPreference) {
	self->setHintingPreference(hintingPreference);
}

HintingPreference QFont_HintingPreference(const QFont* self) {
	return self->hintingPreference();
}

void QFont_SetFeature(QFont* self, Tag tag, unsigned int value) {
	self->setFeature(tag, static_cast<quint32>(value));
}

void QFont_UnsetFeature(QFont* self, Tag tag) {
	self->unsetFeature(tag);
}

unsigned int QFont_FeatureValue(const QFont* self, Tag tag) {
	quint32 _ret = self->featureValue(tag);
	return static_cast<unsigned int>(_ret);
}

bool QFont_IsFeatureSet(const QFont* self, Tag tag) {
	return self->isFeatureSet(tag);
}

struct miqt_array /* of Tag */  QFont_FeatureTags(const QFont* self) {
	QList<Tag> _ret = self->featureTags();
	// Convert QList<> from C++ memory to manually-managed C memory
	Tag* _arr = static_cast<Tag*>(malloc(sizeof(Tag) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QFont_ClearFeatures(QFont* self) {
	self->clearFeatures();
}

void QFont_SetVariableAxis(QFont* self, Tag tag, float value) {
	self->setVariableAxis(tag, static_cast<float>(value));
}

void QFont_UnsetVariableAxis(QFont* self, Tag tag) {
	self->unsetVariableAxis(tag);
}

bool QFont_IsVariableAxisSet(const QFont* self, Tag tag) {
	return self->isVariableAxisSet(tag);
}

float QFont_VariableAxisValue(const QFont* self, Tag tag) {
	return self->variableAxisValue(tag);
}

void QFont_ClearVariableAxes(QFont* self) {
	self->clearVariableAxes();
}

struct miqt_array /* of Tag */  QFont_VariableAxisTags(const QFont* self) {
	QList<Tag> _ret = self->variableAxisTags();
	// Convert QList<> from C++ memory to manually-managed C memory
	Tag* _arr = static_cast<Tag*>(malloc(sizeof(Tag) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

bool QFont_ExactMatch(const QFont* self) {
	return self->exactMatch();
}

void QFont_OperatorAssign(QFont* self, QFont* param1) {
	self->operator=(*param1);
}

bool QFont_OperatorEqual(const QFont* self, QFont* param1) {
	return (*self == *param1);
}

bool QFont_OperatorNotEqual(const QFont* self, QFont* param1) {
	return (*self != *param1);
}

bool QFont_OperatorLesser(const QFont* self, QFont* param1) {
	return (*self < *param1);
}

bool QFont_IsCopyOf(const QFont* self, QFont* param1) {
	return self->isCopyOf(*param1);
}

struct miqt_string QFont_Key(const QFont* self) {
	QString _ret = self->key();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QFont_ToString(const QFont* self) {
	QString _ret = self->toString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QFont_FromString(QFont* self, struct miqt_string param1) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	return self->fromString(param1_QString);
}

struct miqt_string QFont_Substitute(struct miqt_string param1) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QString _ret = QFont::substitute(param1_QString);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_array /* of struct miqt_string */  QFont_Substitutes(struct miqt_string param1) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QStringList _ret = QFont::substitutes(param1_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of struct miqt_string */  QFont_Substitutions() {
	QStringList _ret = QFont::substitutions();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QFont_InsertSubstitution(struct miqt_string param1, struct miqt_string param2) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QString param2_QString = QString::fromUtf8(param2.data, param2.len);
	QFont::insertSubstitution(param1_QString, param2_QString);
}

void QFont_InsertSubstitutions(struct miqt_string param1, struct miqt_array /* of struct miqt_string */  param2) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QStringList param2_QList;
	param2_QList.reserve(param2.len);
	struct miqt_string* param2_arr = static_cast<struct miqt_string*>(param2.data);
	for(size_t i = 0; i < param2.len; ++i) {
		QString param2_arr_i_QString = QString::fromUtf8(param2_arr[i].data, param2_arr[i].len);
		param2_QList.push_back(param2_arr_i_QString);
	}
	QFont::insertSubstitutions(param1_QString, param2_QList);
}

void QFont_RemoveSubstitutions(struct miqt_string param1) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	QFont::removeSubstitutions(param1_QString);
}

void QFont_Initialize() {
	QFont::initialize();
}

void QFont_Cleanup() {
	QFont::cleanup();
}

void QFont_CacheStatistics() {
	QFont::cacheStatistics();
}

struct miqt_string QFont_DefaultFamily(const QFont* self) {
	QString _ret = self->defaultFamily();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QFont* QFont_Resolve(const QFont* self, QFont* param1) {
	return new QFont(self->resolve(*param1));
}

unsigned int QFont_ResolveMask(const QFont* self) {
	uint _ret = self->resolveMask();
	return static_cast<unsigned int>(_ret);
}

void QFont_SetResolveMask(QFont* self, unsigned int mask) {
	self->setResolveMask(static_cast<uint>(mask));
}

void QFont_SetLegacyWeight(QFont* self, int legacyWeight) {
	self->setLegacyWeight(static_cast<int>(legacyWeight));
}

int QFont_LegacyWeight(const QFont* self) {
	return self->legacyWeight();
}

void QFont_SetStyleHint2(QFont* self, StyleHint param1, StyleStrategy param2) {
	self->setStyleHint(param1, param2);
}

void QFont_Delete(QFont* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QFont*>( self );
	} else {
		delete self;
	}
}

QFont__Tag* QFont__Tag_new() {
	return new QFont::Tag();
}

QFont__Tag* QFont__Tag_new2(const Tag* param1) {
	return new QFont::Tag(*param1);
}

bool QFont__Tag_IsValid(const QFont__Tag* self) {
	return self->isValid();
}

unsigned int QFont__Tag_Value(const QFont__Tag* self) {
	quint32 _ret = self->value();
	return static_cast<unsigned int>(_ret);
}

struct miqt_string QFont__Tag_ToString(const QFont__Tag* self) {
	QByteArray _qb = self->toString();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

void QFont__Tag_Delete(QFont__Tag* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QFont::Tag*>( self );
	} else {
		delete self;
	}
}

