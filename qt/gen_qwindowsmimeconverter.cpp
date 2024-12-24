// +build ignore

#include <QList>
#include <QMetaType>
#include <QMimeData>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include <QWindowsMimeConverter>
#include <qwindowsmimeconverter.h>
#include "gen_qwindowsmimeconverter.h"

QWindowsMimeConverter* QWindowsMimeConverter_new() {
	return new QWindowsMimeConverter();
}

int QWindowsMimeConverter_RegisterMimeType(struct miqt_string mimeType) {
	QString mimeType_QString = QString::fromUtf8(mimeType.data, mimeType.len);
	return QWindowsMimeConverter::registerMimeType(mimeType_QString);
}

bool QWindowsMimeConverter_CanConvertFromMime(const QWindowsMimeConverter* self, const tagFORMATETC* formatetc, QMimeData* mimeData) {
	return self->canConvertFromMime(*formatetc, mimeData);
}

bool QWindowsMimeConverter_ConvertFromMime(const QWindowsMimeConverter* self, const tagFORMATETC* formatetc, QMimeData* mimeData, tagSTGMEDIUM* pmedium) {
	return self->convertFromMime(*formatetc, mimeData, pmedium);
}

struct miqt_array /* of tagFORMATETC */  QWindowsMimeConverter_FormatsForMime(const QWindowsMimeConverter* self, struct miqt_string mimeType, QMimeData* mimeData) {
	QString mimeType_QString = QString::fromUtf8(mimeType.data, mimeType.len);
	QList<FORMATETC> _ret = self->formatsForMime(mimeType_QString, mimeData);
	// Convert QList<> from C++ memory to manually-managed C memory
	tagFORMATETC* _arr = static_cast<tagFORMATETC*>(malloc(sizeof(tagFORMATETC) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

bool QWindowsMimeConverter_CanConvertToMime(const QWindowsMimeConverter* self, struct miqt_string mimeType, IDataObject* pDataObj) {
	QString mimeType_QString = QString::fromUtf8(mimeType.data, mimeType.len);
	return self->canConvertToMime(mimeType_QString, pDataObj);
}

QVariant* QWindowsMimeConverter_ConvertToMime(const QWindowsMimeConverter* self, struct miqt_string mimeType, IDataObject* pDataObj, QMetaType* preferredType) {
	QString mimeType_QString = QString::fromUtf8(mimeType.data, mimeType.len);
	return new QVariant(self->convertToMime(mimeType_QString, pDataObj, *preferredType));
}

struct miqt_string QWindowsMimeConverter_MimeForFormat(const QWindowsMimeConverter* self, const tagFORMATETC* formatetc) {
	QString _ret = self->mimeForFormat(*formatetc);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWindowsMimeConverter_Delete(QWindowsMimeConverter* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QWindowsMimeConverter*>( self );
	} else {
		delete self;
	}
}

