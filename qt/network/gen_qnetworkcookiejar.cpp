// +build ignore

#include <QList>
#include <QMetaObject>
#include <QNetworkCookie>
#include <QNetworkCookieJar>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUrl>
#include <qnetworkcookiejar.h>
#include "gen_qnetworkcookiejar.h"

class MiqtVirtualQNetworkCookieJar : public virtual QNetworkCookieJar {
public:

	MiqtVirtualQNetworkCookieJar(): QNetworkCookieJar() {};
	MiqtVirtualQNetworkCookieJar(QObject* parent): QNetworkCookieJar(parent) {};

	virtual ~MiqtVirtualQNetworkCookieJar() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QNetworkCookieJar::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QNetworkCookieJar_MetaObject(const_cast<MiqtVirtualQNetworkCookieJar*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QNetworkCookieJar::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QNetworkCookieJar::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QNetworkCookieJar_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QNetworkCookieJar::qt_metacast(param1);

	}

};

QNetworkCookieJar* QNetworkCookieJar_new() {
	return new MiqtVirtualQNetworkCookieJar();
}

QNetworkCookieJar* QNetworkCookieJar_new2(QObject* parent) {
	return new MiqtVirtualQNetworkCookieJar(parent);
}

void QNetworkCookieJar_virtbase(QNetworkCookieJar* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QNetworkCookieJar_MetaObject(const QNetworkCookieJar* self) {
	return (QMetaObject*) self->metaObject();
}

void* QNetworkCookieJar_Metacast(QNetworkCookieJar* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QNetworkCookieJar_Tr(const char* s) {
	QString _ret = QNetworkCookieJar::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_array /* of QNetworkCookie* */  QNetworkCookieJar_CookiesForUrl(const QNetworkCookieJar* self, QUrl* url) {
	QList<QNetworkCookie> _ret = self->cookiesForUrl(*url);
	// Convert QList<> from C++ memory to manually-managed C memory
	QNetworkCookie** _arr = static_cast<QNetworkCookie**>(malloc(sizeof(QNetworkCookie*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QNetworkCookie(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

bool QNetworkCookieJar_SetCookiesFromUrl(QNetworkCookieJar* self, struct miqt_array /* of QNetworkCookie* */  cookieList, QUrl* url) {
	QList<QNetworkCookie> cookieList_QList;
	cookieList_QList.reserve(cookieList.len);
	QNetworkCookie** cookieList_arr = static_cast<QNetworkCookie**>(cookieList.data);
	for(size_t i = 0; i < cookieList.len; ++i) {
		cookieList_QList.push_back(*(cookieList_arr[i]));
	}
	return self->setCookiesFromUrl(cookieList_QList, *url);
}

bool QNetworkCookieJar_InsertCookie(QNetworkCookieJar* self, QNetworkCookie* cookie) {
	return self->insertCookie(*cookie);
}

bool QNetworkCookieJar_UpdateCookie(QNetworkCookieJar* self, QNetworkCookie* cookie) {
	return self->updateCookie(*cookie);
}

bool QNetworkCookieJar_DeleteCookie(QNetworkCookieJar* self, QNetworkCookie* cookie) {
	return self->deleteCookie(*cookie);
}

struct miqt_string QNetworkCookieJar_Tr2(const char* s, const char* c) {
	QString _ret = QNetworkCookieJar::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QNetworkCookieJar_Tr3(const char* s, const char* c, int n) {
	QString _ret = QNetworkCookieJar::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QNetworkCookieJar_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQNetworkCookieJar*>( (QNetworkCookieJar*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QNetworkCookieJar_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQNetworkCookieJar*)(self) )->virtualbase_MetaObject();
}

void QNetworkCookieJar_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQNetworkCookieJar*>( (QNetworkCookieJar*)(self) )->handle__Metacast = slot;
}

void* QNetworkCookieJar_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQNetworkCookieJar*)(self) )->virtualbase_Metacast(param1);
}

void QNetworkCookieJar_Delete(QNetworkCookieJar* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQNetworkCookieJar*>( self );
	} else {
		delete self;
	}
}

