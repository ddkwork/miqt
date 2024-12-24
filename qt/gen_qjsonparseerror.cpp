// +build ignore

#include <QJsonParseError>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qjsonparseerror.h>
#include "gen_qjsonparseerror.h"

#ifndef _Bool
#define _Bool bool
#endif

void _GUID_Delete(_GUID* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<_GUID*>( self );
	} else {
		delete self;
	}
}

void type_info_Delete(type_info* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<type_info*>( self );
	} else {
		delete self;
	}
}

struct miqt_string QJsonParseError_ErrorString(const QJsonParseError* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QJsonParseError_Delete(QJsonParseError* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QJsonParseError*>( self );
	} else {
		delete self;
	}
}

