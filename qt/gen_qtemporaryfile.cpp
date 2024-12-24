// +build ignore

#include <QFile>
#include <QFileDevice>
#include <QIODevice>
#include <QIODeviceBase>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTemporaryFile>
#include <qtemporaryfile.h>
#include "gen_qtemporaryfile.h"
class MiqtVirtualQTemporaryFile : public virtual QTemporaryFile {
public:
MiqtVirtualQTemporaryFile(): QTemporaryFile() {};
MiqtVirtualQTemporaryFile(const QString& templateName): QTemporaryFile(templateName) {};
MiqtVirtualQTemporaryFile(QObject* parent): QTemporaryFile(parent) {};
MiqtVirtualQTemporaryFile(const QString& templateName, QObject* parent): QTemporaryFile(templateName, parent) {};

virtual ~MiqtVirtualQTemporaryFile() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QTemporaryFile::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QTemporaryFile_MetaObject(const_cast<MiqtVirtualQTemporaryFile*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QTemporaryFile::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QTemporaryFile::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QTemporaryFile_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QTemporaryFile::qt_metacast(param1);

	}
};
QTemporaryFile* QTemporaryFile_new() {
return new MiqtVirtualQTemporaryFile();
}
QTemporaryFile* QTemporaryFile_new2(struct miqt_string templateName) {
QString templateName_QString = QString::fromUtf8(templateName.data, templateName.len);
	return new MiqtVirtualQTemporaryFile(templateName_QString);
}
QTemporaryFile* QTemporaryFile_new3(QObject* parent) {
return new MiqtVirtualQTemporaryFile(parent);
}
QTemporaryFile* QTemporaryFile_new4(struct miqt_string templateName, QObject* parent) {
QString templateName_QString = QString::fromUtf8(templateName.data, templateName.len);
	return new MiqtVirtualQTemporaryFile(templateName_QString, parent);
}
void QTemporaryFile_virtbase(QTemporaryFile* src
, QFile** outptr_QFile
) {
*outptr_QFile = static_cast<QFile*>(src);
}
QMetaObject* QTemporaryFile_MetaObject(const QTemporaryFile* self) {
	return (QMetaObject*) self->metaObject();
}
void* QTemporaryFile_Metacast(QTemporaryFile* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QTemporaryFile_Tr(const char* s) {
	QString _ret = QTemporaryFile::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QTemporaryFile_AutoRemove(const QTemporaryFile* self) {
	return self->autoRemove();
}
void QTemporaryFile_SetAutoRemove(QTemporaryFile* self, bool b) {
	self->setAutoRemove(b);
}
bool QTemporaryFile_Open(QTemporaryFile* self) {
	return self->open();
}
struct miqt_string QTemporaryFile_FileName(const QTemporaryFile* self) {
	QString _ret = self->fileName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QTemporaryFile_FileTemplate(const QTemporaryFile* self) {
	QString _ret = self->fileTemplate();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTemporaryFile_SetFileTemplate(QTemporaryFile* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->setFileTemplate(name_QString);
}
bool QTemporaryFile_Rename(QTemporaryFile* self, struct miqt_string newName) {
	QString newName_QString = QString::fromUtf8(newName.data, newName.len);
	return self->rename(newName_QString);
}
QTemporaryFile* QTemporaryFile_CreateNativeFile(struct miqt_string fileName) {
	QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return QTemporaryFile::createNativeFile(fileName_QString);
}
QTemporaryFile* QTemporaryFile_CreateNativeFileWithFile(QFile* file) {
	return QTemporaryFile::createNativeFile(*file);
}
struct miqt_string QTemporaryFile_Tr2(const char* s, const char* c) {
	QString _ret = QTemporaryFile::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QTemporaryFile_Tr3(const char* s, const char* c, int n) {
	QString _ret = QTemporaryFile::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTemporaryFile_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTemporaryFile*>( (QTemporaryFile*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QTemporaryFile_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQTemporaryFile*)(self) )->virtualbase_MetaObject();
}
void QTemporaryFile_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTemporaryFile*>( (QTemporaryFile*)(self) )->handle__Metacast = slot;
}
void* QTemporaryFile_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQTemporaryFile*)(self) )->virtualbase_Metacast(param1);
}
void QTemporaryFile_Delete(QTemporaryFile* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQTemporaryFile*>( self );
	} else {
		delete self;
	}
}
