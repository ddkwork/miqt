// +build ignore

#include <QMetaObject>
#include <QObject>
#include <QSignalMapper>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qsignalmapper.h>
#include "gen_qsignalmapper.h"
class MiqtVirtualQSignalMapper : public virtual QSignalMapper {
public:
MiqtVirtualQSignalMapper(): QSignalMapper() {};
MiqtVirtualQSignalMapper(QObject* parent): QSignalMapper(parent) {};

virtual ~MiqtVirtualQSignalMapper() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QSignalMapper::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QSignalMapper_MetaObject(const_cast<MiqtVirtualQSignalMapper*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSignalMapper::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QSignalMapper::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSignalMapper_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSignalMapper::qt_metacast(param1);

	}
};
QSignalMapper* QSignalMapper_new() {
return new MiqtVirtualQSignalMapper();
}
QSignalMapper* QSignalMapper_new2(QObject* parent) {
return new MiqtVirtualQSignalMapper(parent);
}
void QSignalMapper_virtbase(QSignalMapper* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QSignalMapper_MetaObject(const QSignalMapper* self) {
	return (QMetaObject*) self->metaObject();
}
void* QSignalMapper_Metacast(QSignalMapper* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QSignalMapper_Tr(const char* s) {
	QString _ret = QSignalMapper::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSignalMapper_SetMapping(QSignalMapper* self, QObject* sender, int id) {
	self->setMapping(sender, static_cast<int>(id));
}
void QSignalMapper_SetMapping2(QSignalMapper* self, QObject* sender, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setMapping(sender, text_QString);
}
void QSignalMapper_SetMapping3(QSignalMapper* self, QObject* sender, QObject* object) {
	self->setMapping(sender, object);
}
void QSignalMapper_RemoveMappings(QSignalMapper* self, QObject* sender) {
	self->removeMappings(sender);
}
QObject* QSignalMapper_Mapping(const QSignalMapper* self, int id) {
	return self->mapping(static_cast<int>(id));
}
QObject* QSignalMapper_MappingWithText(const QSignalMapper* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->mapping(text_QString);
}
QObject* QSignalMapper_MappingWithObject(const QSignalMapper* self, QObject* object) {
	return self->mapping(object);
}
void QSignalMapper_MappedInt(QSignalMapper* self, int param1) {
	self->mappedInt(static_cast<int>(param1));
}
void QSignalMapper_connect_MappedInt(QSignalMapper* self, intptr_t slot) {
	MiqtVirtualQSignalMapper::connect(self, static_cast<void (QSignalMapper::*)(int)>(&QSignalMapper::mappedInt), self, [=](int param1) {
		int sigval1 = param1;
		miqt_exec_callback_QSignalMapper_MappedInt(slot, sigval1);
	});
}
void QSignalMapper_MappedString(QSignalMapper* self, struct miqt_string param1) {
	QString param1_QString = QString::fromUtf8(param1.data, param1.len);
	self->mappedString(param1_QString);
}
void QSignalMapper_connect_MappedString(QSignalMapper* self, intptr_t slot) {
	MiqtVirtualQSignalMapper::connect(self, static_cast<void (QSignalMapper::*)(const QString&)>(&QSignalMapper::mappedString), self, [=](const QString& param1) {
		const QString param1_ret = param1;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray param1_b = param1_ret.toUtf8();
		struct miqt_string param1_ms;
		param1_ms.len = param1_b.length();
		param1_ms.data = static_cast<char*>(malloc(param1_ms.len));
		memcpy(param1_ms.data, param1_b.data(), param1_ms.len);
		struct miqt_string sigval1 = param1_ms;
		miqt_exec_callback_QSignalMapper_MappedString(slot, sigval1);
	});
}
void QSignalMapper_MappedObject(QSignalMapper* self, QObject* param1) {
	self->mappedObject(param1);
}
void QSignalMapper_connect_MappedObject(QSignalMapper* self, intptr_t slot) {
	MiqtVirtualQSignalMapper::connect(self, static_cast<void (QSignalMapper::*)(QObject*)>(&QSignalMapper::mappedObject), self, [=](QObject* param1) {
		QObject* sigval1 = param1;
		miqt_exec_callback_QSignalMapper_MappedObject(slot, sigval1);
	});
}
void QSignalMapper_Map(QSignalMapper* self) {
	self->map();
}
void QSignalMapper_MapWithSender(QSignalMapper* self, QObject* sender) {
	self->map(sender);
}
struct miqt_string QSignalMapper_Tr2(const char* s, const char* c) {
	QString _ret = QSignalMapper::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QSignalMapper_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSignalMapper::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSignalMapper_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSignalMapper*>( (QSignalMapper*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QSignalMapper_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSignalMapper*)(self) )->virtualbase_MetaObject();
}
void QSignalMapper_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSignalMapper*>( (QSignalMapper*)(self) )->handle__Metacast = slot;
}
void* QSignalMapper_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSignalMapper*)(self) )->virtualbase_Metacast(param1);
}
void QSignalMapper_Delete(QSignalMapper* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSignalMapper*>( self );
	} else {
		delete self;
	}
}
