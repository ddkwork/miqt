// +build ignore

#include <QIconEngine>
#include <QIconEnginePlugin>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qiconengineplugin.h>
#include "gen_qiconengineplugin.h"
class MiqtVirtualQIconEnginePlugin : public virtual QIconEnginePlugin {
public:
MiqtVirtualQIconEnginePlugin(): QIconEnginePlugin() {};
MiqtVirtualQIconEnginePlugin(QObject* parent): QIconEnginePlugin(parent) {};

virtual ~MiqtVirtualQIconEnginePlugin() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QIconEnginePlugin::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QIconEnginePlugin_MetaObject(const_cast<MiqtVirtualQIconEnginePlugin*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QIconEnginePlugin::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QIconEnginePlugin::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QIconEnginePlugin_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QIconEnginePlugin::qt_metacast(param1);

	}
};
QIconEnginePlugin* QIconEnginePlugin_new() {
return new MiqtVirtualQIconEnginePlugin();
}
QIconEnginePlugin* QIconEnginePlugin_new2(QObject* parent) {
return new MiqtVirtualQIconEnginePlugin(parent);
}
void QIconEnginePlugin_virtbase(QIconEnginePlugin* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QIconEnginePlugin_MetaObject(const QIconEnginePlugin* self) {
	return (QMetaObject*) self->metaObject();
}
void* QIconEnginePlugin_Metacast(QIconEnginePlugin* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QIconEnginePlugin_Tr(const char* s) {
	QString _ret = QIconEnginePlugin::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QIconEngine* QIconEnginePlugin_Create(QIconEnginePlugin* self, struct miqt_string filename) {
	QString filename_QString = QString::fromUtf8(filename.data, filename.len);
	return self->create(filename_QString);
}
struct miqt_string QIconEnginePlugin_Tr2(const char* s, const char* c) {
	QString _ret = QIconEnginePlugin::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QIconEnginePlugin_Tr3(const char* s, const char* c, int n) {
	QString _ret = QIconEnginePlugin::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QIconEnginePlugin_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQIconEnginePlugin*>( (QIconEnginePlugin*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QIconEnginePlugin_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQIconEnginePlugin*)(self) )->virtualbase_MetaObject();
}
void QIconEnginePlugin_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQIconEnginePlugin*>( (QIconEnginePlugin*)(self) )->handle__Metacast = slot;
}
void* QIconEnginePlugin_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQIconEnginePlugin*)(self) )->virtualbase_Metacast(param1);
}
void QIconEnginePlugin_Delete(QIconEnginePlugin* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQIconEnginePlugin*>( self );
	} else {
		delete self;
	}
}
