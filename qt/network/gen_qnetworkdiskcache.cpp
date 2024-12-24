// +build ignore

#include <QAbstractNetworkCache>
#include <QIODevice>
#include <QMetaObject>
#include <QNetworkCacheMetaData>
#include <QNetworkDiskCache>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUrl>
#include <qnetworkdiskcache.h>
#include "gen_qnetworkdiskcache.h"
class MiqtVirtualQNetworkDiskCache : public virtual QNetworkDiskCache {
public:
MiqtVirtualQNetworkDiskCache(): QNetworkDiskCache() {};
MiqtVirtualQNetworkDiskCache(QObject* parent): QNetworkDiskCache(parent) {};

virtual ~MiqtVirtualQNetworkDiskCache() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QNetworkDiskCache::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QNetworkDiskCache_MetaObject(const_cast<MiqtVirtualQNetworkDiskCache*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QNetworkDiskCache::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QNetworkDiskCache::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QNetworkDiskCache_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QNetworkDiskCache::qt_metacast(param1);

	}
};
QNetworkDiskCache* QNetworkDiskCache_new() {
return new MiqtVirtualQNetworkDiskCache();
}
QNetworkDiskCache* QNetworkDiskCache_new2(QObject* parent) {
return new MiqtVirtualQNetworkDiskCache(parent);
}
void QNetworkDiskCache_virtbase(QNetworkDiskCache* src
, QAbstractNetworkCache** outptr_QAbstractNetworkCache
) {
*outptr_QAbstractNetworkCache = static_cast<QAbstractNetworkCache*>(src);
}
QMetaObject* QNetworkDiskCache_MetaObject(const QNetworkDiskCache* self) {
	return (QMetaObject*) self->metaObject();
}
void* QNetworkDiskCache_Metacast(QNetworkDiskCache* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QNetworkDiskCache_Tr(const char* s) {
	QString _ret = QNetworkDiskCache::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QNetworkDiskCache_CacheDirectory(const QNetworkDiskCache* self) {
	QString _ret = self->cacheDirectory();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QNetworkDiskCache_SetCacheDirectory(QNetworkDiskCache* self, struct miqt_string cacheDir) {
	QString cacheDir_QString = QString::fromUtf8(cacheDir.data, cacheDir.len);
	self->setCacheDirectory(cacheDir_QString);
}
long long QNetworkDiskCache_MaximumCacheSize(const QNetworkDiskCache* self) {
	qint64 _ret = self->maximumCacheSize();
	return static_cast<long long>(_ret);
}
void QNetworkDiskCache_SetMaximumCacheSize(QNetworkDiskCache* self, long long size) {
	self->setMaximumCacheSize(static_cast<qint64>(size));
}
long long QNetworkDiskCache_CacheSize(const QNetworkDiskCache* self) {
	qint64 _ret = self->cacheSize();
	return static_cast<long long>(_ret);
}
QNetworkCacheMetaData* QNetworkDiskCache_MetaData(QNetworkDiskCache* self, QUrl* url) {
	return new QNetworkCacheMetaData(self->metaData(*url));
}
void QNetworkDiskCache_UpdateMetaData(QNetworkDiskCache* self, QNetworkCacheMetaData* metaData) {
	self->updateMetaData(*metaData);
}
QIODevice* QNetworkDiskCache_Data(QNetworkDiskCache* self, QUrl* url) {
	return self->data(*url);
}
bool QNetworkDiskCache_Remove(QNetworkDiskCache* self, QUrl* url) {
	return self->remove(*url);
}
QIODevice* QNetworkDiskCache_Prepare(QNetworkDiskCache* self, QNetworkCacheMetaData* metaData) {
	return self->prepare(*metaData);
}
void QNetworkDiskCache_Insert(QNetworkDiskCache* self, QIODevice* device) {
	self->insert(device);
}
QNetworkCacheMetaData* QNetworkDiskCache_FileMetaData(const QNetworkDiskCache* self, struct miqt_string fileName) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return new QNetworkCacheMetaData(self->fileMetaData(fileName_QString));
}
void QNetworkDiskCache_Clear(QNetworkDiskCache* self) {
	self->clear();
}
struct miqt_string QNetworkDiskCache_Tr2(const char* s, const char* c) {
	QString _ret = QNetworkDiskCache::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QNetworkDiskCache_Tr3(const char* s, const char* c, int n) {
	QString _ret = QNetworkDiskCache::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QNetworkDiskCache_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQNetworkDiskCache*>( (QNetworkDiskCache*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QNetworkDiskCache_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQNetworkDiskCache*)(self) )->virtualbase_MetaObject();
}
void QNetworkDiskCache_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQNetworkDiskCache*>( (QNetworkDiskCache*)(self) )->handle__Metacast = slot;
}
void* QNetworkDiskCache_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQNetworkDiskCache*)(self) )->virtualbase_Metacast(param1);
}
void QNetworkDiskCache_Delete(QNetworkDiskCache* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQNetworkDiskCache*>( self );
	} else {
		delete self;
	}
}
