// +build ignore

#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStringConverter>
#include <QStringConverterBase>
#define WORKAROUND_INNER_CLASS_DEFINITION_QStringConverterBase__State
#include <qstringconverter_base.h>
#include "gen_qstringconverter_base.h"

#ifndef _Bool
#define _Bool bool
#endif

QStringConverterBase* QStringConverterBase_new(QStringConverterBase* param1) {
	return new QStringConverterBase(*param1);
}

QStringConverterBase* QStringConverterBase_new2() {
	return new QStringConverterBase();
}

void QStringConverter_virtbase(QStringConverter* src, QStringConverterBase** outptr_QStringConverterBase) {
	*outptr_QStringConverterBase = static_cast<QStringConverterBase*>(src);
}

bool QStringConverter_IsValid(const QStringConverter* self) {
	return self->isValid();
}

void QStringConverter_ResetState(QStringConverter* self) {
	self->resetState();
}

bool QStringConverter_HasError(const QStringConverter* self) {
	return self->hasError();
}

const char* QStringConverter_Name(const QStringConverter* self) {
	return (const char*) self->name();
}

const char* QStringConverter_NameForEncoding(Encoding e) {
	return (const char*) QStringConverter::nameForEncoding(e);
}

struct miqt_array /* of struct miqt_string */  QStringConverter_AvailableCodecs() {
	QStringList _ret = QStringConverter::availableCodecs();
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

QStringConverterBase__State* QStringConverterBase__State_new() {
	return new QStringConverterBase::State();
}

QStringConverterBase__State* QStringConverterBase__State_new2(Flags f) {
	return new QStringConverterBase::State(f);
}

void QStringConverterBase__State_Clear(QStringConverterBase__State* self) {
	self->clear();
}

void QStringConverterBase__State_Reset(QStringConverterBase__State* self) {
	self->reset();
}

void QStringConverterBase__State_Delete(QStringConverterBase__State* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QStringConverterBase::State*>( self );
	} else {
		delete self;
	}
}

