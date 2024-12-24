// +build ignore

#include <QLibrary>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qlibrary.h>
#include "gen_qlibrary.h"

class MiqtVirtualQLibrary : public virtual QLibrary {
public:

	MiqtVirtualQLibrary(): QLibrary() {};
	MiqtVirtualQLibrary(const QString& fileName): QLibrary(fileName) {};
	MiqtVirtualQLibrary(const QString& fileName, int verNum): QLibrary(fileName, verNum) {};
	MiqtVirtualQLibrary(const QString& fileName, const QString& version): QLibrary(fileName, version) {};
	MiqtVirtualQLibrary(QObject* parent): QLibrary(parent) {};
	MiqtVirtualQLibrary(const QString& fileName, QObject* parent): QLibrary(fileName, parent) {};
	MiqtVirtualQLibrary(const QString& fileName, int verNum, QObject* parent): QLibrary(fileName, verNum, parent) {};
	MiqtVirtualQLibrary(const QString& fileName, const QString& version, QObject* parent): QLibrary(fileName, version, parent) {};

	virtual ~MiqtVirtualQLibrary() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QLibrary::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QLibrary_MetaObject(const_cast<MiqtVirtualQLibrary*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QLibrary::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QLibrary::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QLibrary_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QLibrary::qt_metacast(param1);

	}

};

QLibrary* QLibrary_new() {
	return new MiqtVirtualQLibrary();
}

QLibrary* QLibrary_new2(struct miqt_string fileName) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return new MiqtVirtualQLibrary(fileName_QString);
}

QLibrary* QLibrary_new3(struct miqt_string fileName, int verNum) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return new MiqtVirtualQLibrary(fileName_QString, static_cast<int>(verNum));
}

QLibrary* QLibrary_new4(struct miqt_string fileName, struct miqt_string version) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	QString version_QString = QString::fromUtf8(version.data, version.len);
	return new MiqtVirtualQLibrary(fileName_QString, version_QString);
}

QLibrary* QLibrary_new5(QObject* parent) {
	return new MiqtVirtualQLibrary(parent);
}

QLibrary* QLibrary_new6(struct miqt_string fileName, QObject* parent) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return new MiqtVirtualQLibrary(fileName_QString, parent);
}

QLibrary* QLibrary_new7(struct miqt_string fileName, int verNum, QObject* parent) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return new MiqtVirtualQLibrary(fileName_QString, static_cast<int>(verNum), parent);
}

QLibrary* QLibrary_new8(struct miqt_string fileName, struct miqt_string version, QObject* parent) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	QString version_QString = QString::fromUtf8(version.data, version.len);
	return new MiqtVirtualQLibrary(fileName_QString, version_QString, parent);
}

void QLibrary_virtbase(QLibrary* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QLibrary_MetaObject(const QLibrary* self) {
	return (QMetaObject*) self->metaObject();
}

void* QLibrary_Metacast(QLibrary* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QLibrary_Tr(const char* s) {
	QString _ret = QLibrary::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QLibrary_Load(QLibrary* self) {
	return self->load();
}

bool QLibrary_Unload(QLibrary* self) {
	return self->unload();
}

bool QLibrary_IsLoaded(const QLibrary* self) {
	return self->isLoaded();
}

bool QLibrary_IsLibrary(struct miqt_string fileName) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return QLibrary::isLibrary(fileName_QString);
}

void QLibrary_SetFileName(QLibrary* self, struct miqt_string fileName) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	self->setFileName(fileName_QString);
}

struct miqt_string QLibrary_FileName(const QLibrary* self) {
	QString _ret = self->fileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLibrary_SetFileNameAndVersion(QLibrary* self, struct miqt_string fileName, int verNum) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	self->setFileNameAndVersion(fileName_QString, static_cast<int>(verNum));
}

void QLibrary_SetFileNameAndVersion2(QLibrary* self, struct miqt_string fileName, struct miqt_string version) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	QString version_QString = QString::fromUtf8(version.data, version.len);
	self->setFileNameAndVersion(fileName_QString, version_QString);
}

struct miqt_string QLibrary_ErrorString(const QLibrary* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLibrary_SetLoadHints(QLibrary* self, LoadHints hints) {
	self->setLoadHints(hints);
}

LoadHints QLibrary_LoadHints(const QLibrary* self) {
	return self->loadHints();
}

struct miqt_string QLibrary_Tr2(const char* s, const char* c) {
	QString _ret = QLibrary::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QLibrary_Tr3(const char* s, const char* c, int n) {
	QString _ret = QLibrary::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QLibrary_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLibrary*>( (QLibrary*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QLibrary_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQLibrary*)(self) )->virtualbase_MetaObject();
}

void QLibrary_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQLibrary*>( (QLibrary*)(self) )->handle__Metacast = slot;
}

void* QLibrary_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQLibrary*)(self) )->virtualbase_Metacast(param1);
}

void QLibrary_Delete(QLibrary* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQLibrary*>( self );
	} else {
		delete self;
	}
}

