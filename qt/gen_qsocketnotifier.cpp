// +build ignore

#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QSocketDescriptor>
#include <QSocketNotifier>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qsocketnotifier.h>
#include "gen_qsocketnotifier.h"

class MiqtVirtualQSocketNotifier : public virtual QSocketNotifier {
public:

	MiqtVirtualQSocketNotifier(Type param1): QSocketNotifier(param1) {};
	MiqtVirtualQSocketNotifier(qintptr socket, Type param2): QSocketNotifier(socket, param2) {};
	MiqtVirtualQSocketNotifier(Type param1, QObject* parent): QSocketNotifier(param1, parent) {};
	MiqtVirtualQSocketNotifier(qintptr socket, Type param2, QObject* parent): QSocketNotifier(socket, param2, parent) {};

	virtual ~MiqtVirtualQSocketNotifier() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QSocketNotifier::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QSocketNotifier_MetaObject(const_cast<MiqtVirtualQSocketNotifier*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSocketNotifier::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QSocketNotifier::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSocketNotifier_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSocketNotifier::qt_metacast(param1);

	}

};

QSocketNotifier* QSocketNotifier_new(Type param1) {
	return new MiqtVirtualQSocketNotifier(param1);
}

QSocketNotifier* QSocketNotifier_new2(intptr_t socket, Type param2) {
	return new MiqtVirtualQSocketNotifier((qintptr)(socket), param2);
}

QSocketNotifier* QSocketNotifier_new3(Type param1, QObject* parent) {
	return new MiqtVirtualQSocketNotifier(param1, parent);
}

QSocketNotifier* QSocketNotifier_new4(intptr_t socket, Type param2, QObject* parent) {
	return new MiqtVirtualQSocketNotifier((qintptr)(socket), param2, parent);
}

void QSocketNotifier_virtbase(QSocketNotifier* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QSocketNotifier_MetaObject(const QSocketNotifier* self) {
	return (QMetaObject*) self->metaObject();
}

void* QSocketNotifier_Metacast(QSocketNotifier* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QSocketNotifier_Tr(const char* s) {
	QString _ret = QSocketNotifier::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QSocketNotifier_SetSocket(QSocketNotifier* self, intptr_t socket) {
	self->setSocket((qintptr)(socket));
}

intptr_t QSocketNotifier_Socket(const QSocketNotifier* self) {
	qintptr _ret = self->socket();
	return (intptr_t)(_ret);
}

Type QSocketNotifier_Type(const QSocketNotifier* self) {
	return self->type();
}

bool QSocketNotifier_IsValid(const QSocketNotifier* self) {
	return self->isValid();
}

bool QSocketNotifier_IsEnabled(const QSocketNotifier* self) {
	return self->isEnabled();
}

void QSocketNotifier_SetEnabled(QSocketNotifier* self, bool enabled) {
	self->setEnabled(enabled);
}

struct miqt_string QSocketNotifier_Tr2(const char* s, const char* c) {
	QString _ret = QSocketNotifier::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QSocketNotifier_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSocketNotifier::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QSocketNotifier_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSocketNotifier*>( (QSocketNotifier*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QSocketNotifier_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSocketNotifier*)(self) )->virtualbase_MetaObject();
}

void QSocketNotifier_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSocketNotifier*>( (QSocketNotifier*)(self) )->handle__Metacast = slot;
}

void* QSocketNotifier_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSocketNotifier*)(self) )->virtualbase_Metacast(param1);
}

void QSocketNotifier_Delete(QSocketNotifier* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSocketNotifier*>( self );
	} else {
		delete self;
	}
}

QSocketDescriptor* QSocketDescriptor_new() {
	return new QSocketDescriptor();
}

QSocketDescriptor* QSocketDescriptor_new2(intptr_t desc) {
	return new QSocketDescriptor((qintptr)(desc));
}

QSocketDescriptor* QSocketDescriptor_new3(QSocketDescriptor* param1) {
	return new QSocketDescriptor(*param1);
}

QSocketDescriptor* QSocketDescriptor_new4(DescriptorType descriptor) {
	return new QSocketDescriptor(descriptor);
}

void* QSocketDescriptor_WinHandle(const QSocketDescriptor* self) {
	Qt::HANDLE _ret = self->winHandle();
	return static_cast<void*>(_ret);
}

bool QSocketDescriptor_IsValid(const QSocketDescriptor* self) {
	return self->isValid();
}

void QSocketDescriptor_Delete(QSocketDescriptor* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QSocketDescriptor*>( self );
	} else {
		delete self;
	}
}

