// +build ignore

#include <QAbstractAnimation>
#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QPauseAnimation>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qpauseanimation.h>
#include "gen_qpauseanimation.h"

class MiqtVirtualQPauseAnimation : public virtual QPauseAnimation {
public:

	MiqtVirtualQPauseAnimation(): QPauseAnimation() {};
	MiqtVirtualQPauseAnimation(int msecs): QPauseAnimation(msecs) {};
	MiqtVirtualQPauseAnimation(QObject* parent): QPauseAnimation(parent) {};
	MiqtVirtualQPauseAnimation(int msecs, QObject* parent): QPauseAnimation(msecs, parent) {};

	virtual ~MiqtVirtualQPauseAnimation() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QPauseAnimation::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QPauseAnimation_MetaObject(const_cast<MiqtVirtualQPauseAnimation*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QPauseAnimation::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QPauseAnimation::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QPauseAnimation_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QPauseAnimation::qt_metacast(param1);

	}

};

QPauseAnimation* QPauseAnimation_new() {
	return new MiqtVirtualQPauseAnimation();
}

QPauseAnimation* QPauseAnimation_new2(int msecs) {
	return new MiqtVirtualQPauseAnimation(static_cast<int>(msecs));
}

QPauseAnimation* QPauseAnimation_new3(QObject* parent) {
	return new MiqtVirtualQPauseAnimation(parent);
}

QPauseAnimation* QPauseAnimation_new4(int msecs, QObject* parent) {
	return new MiqtVirtualQPauseAnimation(static_cast<int>(msecs), parent);
}

void QPauseAnimation_virtbase(QPauseAnimation* src, QAbstractAnimation** outptr_QAbstractAnimation) {
	*outptr_QAbstractAnimation = static_cast<QAbstractAnimation*>(src);
}

QMetaObject* QPauseAnimation_MetaObject(const QPauseAnimation* self) {
	return (QMetaObject*) self->metaObject();
}

void* QPauseAnimation_Metacast(QPauseAnimation* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QPauseAnimation_Tr(const char* s) {
	QString _ret = QPauseAnimation::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QPauseAnimation_Duration(const QPauseAnimation* self) {
	return self->duration();
}

void QPauseAnimation_SetDuration(QPauseAnimation* self, int msecs) {
	self->setDuration(static_cast<int>(msecs));
}

struct miqt_string QPauseAnimation_Tr2(const char* s, const char* c) {
	QString _ret = QPauseAnimation::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QPauseAnimation_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPauseAnimation::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPauseAnimation_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPauseAnimation*>( (QPauseAnimation*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QPauseAnimation_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQPauseAnimation*)(self) )->virtualbase_MetaObject();
}

void QPauseAnimation_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPauseAnimation*>( (QPauseAnimation*)(self) )->handle__Metacast = slot;
}

void* QPauseAnimation_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQPauseAnimation*)(self) )->virtualbase_Metacast(param1);
}

void QPauseAnimation_Delete(QPauseAnimation* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQPauseAnimation*>( self );
	} else {
		delete self;
	}
}

