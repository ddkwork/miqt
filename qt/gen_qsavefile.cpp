// +build ignore

#include <QFileDevice>
#include <QIODevice>
#include <QIODeviceBase>
#include <QMetaObject>
#include <QObject>
#include <QSaveFile>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qsavefile.h>
#include "gen_qsavefile.h"

class MiqtVirtualQSaveFile : public virtual QSaveFile {
public:

	MiqtVirtualQSaveFile(const QString& name): QSaveFile(name) {};
	MiqtVirtualQSaveFile(): QSaveFile() {};
	MiqtVirtualQSaveFile(const QString& name, QObject* parent): QSaveFile(name, parent) {};
	MiqtVirtualQSaveFile(QObject* parent): QSaveFile(parent) {};

	virtual ~MiqtVirtualQSaveFile() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QSaveFile::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QSaveFile_MetaObject(const_cast<MiqtVirtualQSaveFile*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSaveFile::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QSaveFile::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSaveFile_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSaveFile::qt_metacast(param1);

	}

};

QSaveFile* QSaveFile_new(struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQSaveFile(name_QString);
}

QSaveFile* QSaveFile_new2() {
	return new MiqtVirtualQSaveFile();
}

QSaveFile* QSaveFile_new3(struct miqt_string name, QObject* parent) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new MiqtVirtualQSaveFile(name_QString, parent);
}

QSaveFile* QSaveFile_new4(QObject* parent) {
	return new MiqtVirtualQSaveFile(parent);
}

void QSaveFile_virtbase(QSaveFile* src, QFileDevice** outptr_QFileDevice) {
	*outptr_QFileDevice = static_cast<QFileDevice*>(src);
}

QMetaObject* QSaveFile_MetaObject(const QSaveFile* self) {
	return (QMetaObject*) self->metaObject();
}

void* QSaveFile_Metacast(QSaveFile* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QSaveFile_Tr(const char* s) {
	QString _ret = QSaveFile::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QSaveFile_FileName(const QSaveFile* self) {
	QString _ret = self->fileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QSaveFile_SetFileName(QSaveFile* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->setFileName(name_QString);
}

bool QSaveFile_Open(QSaveFile* self, OpenMode flags) {
	return self->open(flags);
}

bool QSaveFile_Commit(QSaveFile* self) {
	return self->commit();
}

void QSaveFile_CancelWriting(QSaveFile* self) {
	self->cancelWriting();
}

void QSaveFile_SetDirectWriteFallback(QSaveFile* self, bool enabled) {
	self->setDirectWriteFallback(enabled);
}

bool QSaveFile_DirectWriteFallback(const QSaveFile* self) {
	return self->directWriteFallback();
}

struct miqt_string QSaveFile_Tr2(const char* s, const char* c) {
	QString _ret = QSaveFile::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QSaveFile_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSaveFile::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QSaveFile_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSaveFile*>( (QSaveFile*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QSaveFile_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSaveFile*)(self) )->virtualbase_MetaObject();
}

void QSaveFile_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSaveFile*>( (QSaveFile*)(self) )->handle__Metacast = slot;
}

void* QSaveFile_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSaveFile*)(self) )->virtualbase_Metacast(param1);
}

void QSaveFile_Delete(QSaveFile* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSaveFile*>( self );
	} else {
		delete self;
	}
}

