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

#ifndef _Bool
#define _Bool bool
#endif

class MiqtVirtualQWindowsMimeConverter : public virtual QWindowsMimeConverter {
public:

	MiqtVirtualQWindowsMimeConverter(): QWindowsMimeConverter() {};

	virtual ~MiqtVirtualQWindowsMimeConverter() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CanConvertFromMime = 0;

	// Subclass to allow providing a Go implementation
	virtual bool canConvertFromMime(const FORMATETC& formatetc, const QMimeData* mimeData) const override {
		if (handle__CanConvertFromMime == 0) {
			return false; // Pure virtual, there is no base we can call
		}
		
		const FORMATETC& formatetc_ret = formatetc;
		const tagFORMATETC* sigval1 = static_cast<const tagFORMATETC*>(formatetc_ret);
		QMimeData* sigval2 = (QMimeData*) mimeData;

		bool callback_return_value = miqt_exec_callback_QWindowsMimeConverter_CanConvertFromMime(const_cast<MiqtVirtualQWindowsMimeConverter*>(this), handle__CanConvertFromMime, sigval1, sigval2);

		return callback_return_value;
	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ConvertFromMime = 0;

	// Subclass to allow providing a Go implementation
	virtual bool convertFromMime(const FORMATETC& formatetc, const QMimeData* mimeData, STGMEDIUM* pmedium) const override {
		if (handle__ConvertFromMime == 0) {
			return false; // Pure virtual, there is no base we can call
		}
		
		const FORMATETC& formatetc_ret = formatetc;
		const tagFORMATETC* sigval1 = static_cast<const tagFORMATETC*>(formatetc_ret);
		QMimeData* sigval2 = (QMimeData*) mimeData;
		STGMEDIUM* pmedium_ret = pmedium;
		tagSTGMEDIUM* sigval3 = static_cast<tagSTGMEDIUM*>(pmedium_ret);

		bool callback_return_value = miqt_exec_callback_QWindowsMimeConverter_ConvertFromMime(const_cast<MiqtVirtualQWindowsMimeConverter*>(this), handle__ConvertFromMime, sigval1, sigval2, sigval3);

		return callback_return_value;
	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__FormatsForMime = 0;

	// Subclass to allow providing a Go implementation
	virtual QList<FORMATETC> formatsForMime(const QString& mimeType, const QMimeData* mimeData) const override {
		if (handle__FormatsForMime == 0) {
			return QList<FORMATETC>(); // Pure virtual, there is no base we can call
		}
		
		const QString mimeType_ret = mimeType;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray mimeType_b = mimeType_ret.toUtf8();
		struct miqt_string mimeType_ms;
		mimeType_ms.len = mimeType_b.length();
		mimeType_ms.data = static_cast<char*>(malloc(mimeType_ms.len));
		memcpy(mimeType_ms.data, mimeType_b.data(), mimeType_ms.len);
		struct miqt_string sigval1 = mimeType_ms;
		QMimeData* sigval2 = (QMimeData*) mimeData;

		struct miqt_array /* of tagFORMATETC */  callback_return_value = miqt_exec_callback_QWindowsMimeConverter_FormatsForMime(const_cast<MiqtVirtualQWindowsMimeConverter*>(this), handle__FormatsForMime, sigval1, sigval2);
		QList<FORMATETC> callback_return_value_QList;
		callback_return_value_QList.reserve(callback_return_value.len);
		tagFORMATETC* callback_return_value_arr = static_cast<tagFORMATETC*>(callback_return_value.data);
		for(size_t i = 0; i < callback_return_value.len; ++i) {
			callback_return_value_QList.push_back(callback_return_value_arr[i]);
		}

		return callback_return_value_QList;
	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__CanConvertToMime = 0;

	// Subclass to allow providing a Go implementation
	virtual bool canConvertToMime(const QString& mimeType, IDataObject* pDataObj) const override {
		if (handle__CanConvertToMime == 0) {
			return false; // Pure virtual, there is no base we can call
		}
		
		const QString mimeType_ret = mimeType;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray mimeType_b = mimeType_ret.toUtf8();
		struct miqt_string mimeType_ms;
		mimeType_ms.len = mimeType_b.length();
		mimeType_ms.data = static_cast<char*>(malloc(mimeType_ms.len));
		memcpy(mimeType_ms.data, mimeType_b.data(), mimeType_ms.len);
		struct miqt_string sigval1 = mimeType_ms;
		IDataObject* sigval2 = pDataObj;

		bool callback_return_value = miqt_exec_callback_QWindowsMimeConverter_CanConvertToMime(const_cast<MiqtVirtualQWindowsMimeConverter*>(this), handle__CanConvertToMime, sigval1, sigval2);

		return callback_return_value;
	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ConvertToMime = 0;

	// Subclass to allow providing a Go implementation
	virtual QVariant convertToMime(const QString& mimeType, IDataObject* pDataObj, QMetaType preferredType) const override {
		if (handle__ConvertToMime == 0) {
			return QVariant(); // Pure virtual, there is no base we can call
		}
		
		const QString mimeType_ret = mimeType;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray mimeType_b = mimeType_ret.toUtf8();
		struct miqt_string mimeType_ms;
		mimeType_ms.len = mimeType_b.length();
		mimeType_ms.data = static_cast<char*>(malloc(mimeType_ms.len));
		memcpy(mimeType_ms.data, mimeType_b.data(), mimeType_ms.len);
		struct miqt_string sigval1 = mimeType_ms;
		IDataObject* sigval2 = pDataObj;
		QMetaType* sigval3 = new QMetaType(preferredType);

		QVariant* callback_return_value = miqt_exec_callback_QWindowsMimeConverter_ConvertToMime(const_cast<MiqtVirtualQWindowsMimeConverter*>(this), handle__ConvertToMime, sigval1, sigval2, sigval3);

		return *callback_return_value;
	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MimeForFormat = 0;

	// Subclass to allow providing a Go implementation
	virtual QString mimeForFormat(const FORMATETC& formatetc) const override {
		if (handle__MimeForFormat == 0) {
			return QString(); // Pure virtual, there is no base we can call
		}
		
		const FORMATETC& formatetc_ret = formatetc;
		const tagFORMATETC* sigval1 = static_cast<const tagFORMATETC*>(formatetc_ret);

		struct miqt_string callback_return_value = miqt_exec_callback_QWindowsMimeConverter_MimeForFormat(const_cast<MiqtVirtualQWindowsMimeConverter*>(this), handle__MimeForFormat, sigval1);
		QString callback_return_value_QString = QString::fromUtf8(callback_return_value.data, callback_return_value.len);

		return callback_return_value_QString;
	}

};

QWindowsMimeConverter* QWindowsMimeConverter_new() {
	return new MiqtVirtualQWindowsMimeConverter();
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

void QWindowsMimeConverter_override_virtual_CanConvertFromMime(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWindowsMimeConverter*>( (QWindowsMimeConverter*)(self) )->handle__CanConvertFromMime = slot;
}

void QWindowsMimeConverter_override_virtual_ConvertFromMime(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWindowsMimeConverter*>( (QWindowsMimeConverter*)(self) )->handle__ConvertFromMime = slot;
}

void QWindowsMimeConverter_override_virtual_FormatsForMime(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWindowsMimeConverter*>( (QWindowsMimeConverter*)(self) )->handle__FormatsForMime = slot;
}

void QWindowsMimeConverter_override_virtual_CanConvertToMime(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWindowsMimeConverter*>( (QWindowsMimeConverter*)(self) )->handle__CanConvertToMime = slot;
}

void QWindowsMimeConverter_override_virtual_ConvertToMime(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWindowsMimeConverter*>( (QWindowsMimeConverter*)(self) )->handle__ConvertToMime = slot;
}

void QWindowsMimeConverter_override_virtual_MimeForFormat(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWindowsMimeConverter*>( (QWindowsMimeConverter*)(self) )->handle__MimeForFormat = slot;
}

void QWindowsMimeConverter_Delete(QWindowsMimeConverter* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQWindowsMimeConverter*>( self );
	} else {
		delete self;
	}
}

