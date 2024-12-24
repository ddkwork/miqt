// +build ignore

#include <QGenericPlugin>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qgenericplugin.h>
#include "gen_qgenericplugin.h"

class MiqtVirtualQGenericPlugin : public virtual QGenericPlugin {
public:

	MiqtVirtualQGenericPlugin(): QGenericPlugin() {};
	MiqtVirtualQGenericPlugin(QObject* parent): QGenericPlugin(parent) {};

	virtual ~MiqtVirtualQGenericPlugin() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QGenericPlugin::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QGenericPlugin_MetaObject(const_cast<MiqtVirtualQGenericPlugin*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QGenericPlugin::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QGenericPlugin::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QGenericPlugin_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QGenericPlugin::qt_metacast(param1);

	}

};

QGenericPlugin* QGenericPlugin_new() {
	return new MiqtVirtualQGenericPlugin();
}

QGenericPlugin* QGenericPlugin_new2(QObject* parent) {
	return new MiqtVirtualQGenericPlugin(parent);
}

void QGenericPlugin_virtbase(QGenericPlugin* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QGenericPlugin_MetaObject(const QGenericPlugin* self) {
	return (QMetaObject*) self->metaObject();
}

void* QGenericPlugin_Metacast(QGenericPlugin* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QGenericPlugin_Tr(const char* s) {
	QString _ret = QGenericPlugin::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QObject* QGenericPlugin_Create(QGenericPlugin* self, struct miqt_string name, struct miqt_string spec) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	QString spec_QString = QString::fromUtf8(spec.data, spec.len);
	return self->create(name_QString, spec_QString);
}

struct miqt_string QGenericPlugin_Tr2(const char* s, const char* c) {
	QString _ret = QGenericPlugin::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QGenericPlugin_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGenericPlugin::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QGenericPlugin_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGenericPlugin*>( (QGenericPlugin*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QGenericPlugin_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQGenericPlugin*)(self) )->virtualbase_MetaObject();
}

void QGenericPlugin_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGenericPlugin*>( (QGenericPlugin*)(self) )->handle__Metacast = slot;
}

void* QGenericPlugin_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQGenericPlugin*)(self) )->virtualbase_Metacast(param1);
}

void QGenericPlugin_Delete(QGenericPlugin* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGenericPlugin*>( self );
	} else {
		delete self;
	}
}

