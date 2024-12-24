// +build ignore

#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyle>
#include <QStylePlugin>
#include <qstyleplugin.h>
#include "gen_qstyleplugin.h"
class MiqtVirtualQStylePlugin : public virtual QStylePlugin {
public:
MiqtVirtualQStylePlugin(): QStylePlugin() {};
MiqtVirtualQStylePlugin(QObject* parent): QStylePlugin(parent) {};

virtual ~MiqtVirtualQStylePlugin() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QStylePlugin::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QStylePlugin_MetaObject(const_cast<MiqtVirtualQStylePlugin*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QStylePlugin::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QStylePlugin::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QStylePlugin_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QStylePlugin::qt_metacast(param1);

	}
};
QStylePlugin* QStylePlugin_new() {
return new MiqtVirtualQStylePlugin();
}
QStylePlugin* QStylePlugin_new2(QObject* parent) {
return new MiqtVirtualQStylePlugin(parent);
}
void QStylePlugin_virtbase(QStylePlugin* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QStylePlugin_MetaObject(const QStylePlugin* self) {
	return (QMetaObject*) self->metaObject();
}
void* QStylePlugin_Metacast(QStylePlugin* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QStylePlugin_Tr(const char* s) {
	QString _ret = QStylePlugin::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QStyle* QStylePlugin_Create(QStylePlugin* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return self->create(key_QString);
}
struct miqt_string QStylePlugin_Tr2(const char* s, const char* c) {
	QString _ret = QStylePlugin::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QStylePlugin_Tr3(const char* s, const char* c, int n) {
	QString _ret = QStylePlugin::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QStylePlugin_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStylePlugin*>( (QStylePlugin*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QStylePlugin_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQStylePlugin*)(self) )->virtualbase_MetaObject();
}
void QStylePlugin_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStylePlugin*>( (QStylePlugin*)(self) )->handle__Metacast = slot;
}
void* QStylePlugin_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQStylePlugin*)(self) )->virtualbase_Metacast(param1);
}
void QStylePlugin_Delete(QStylePlugin* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQStylePlugin*>( self );
	} else {
		delete self;
	}
}
