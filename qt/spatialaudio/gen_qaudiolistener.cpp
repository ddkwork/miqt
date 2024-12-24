// +build ignore

#include <QAudioEngine>
#include <QAudioListener>
#include <QMetaObject>
#include <QObject>
#include <QQuaternion>
#include <QVector3D>
#include <qaudiolistener.h>
#include "gen_qaudiolistener.h"

class MiqtVirtualQAudioListener : public virtual QAudioListener {
public:

	MiqtVirtualQAudioListener(QAudioEngine* engine): QAudioListener(engine) {};

	virtual ~MiqtVirtualQAudioListener() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QAudioListener::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QAudioListener_MetaObject(const_cast<MiqtVirtualQAudioListener*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAudioListener::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QAudioListener::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAudioListener_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAudioListener::qt_metacast(param1);

	}

};

QAudioListener* QAudioListener_new(QAudioEngine* engine) {
	return new MiqtVirtualQAudioListener(engine);
}

void QAudioListener_virtbase(QAudioListener* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

void QAudioListener_SetPosition(QAudioListener* self, QVector3D* pos) {
	self->setPosition(*pos);
}

QVector3D* QAudioListener_Position(const QAudioListener* self) {
	return new QVector3D(self->position());
}

void QAudioListener_SetRotation(QAudioListener* self, QQuaternion* q) {
	self->setRotation(*q);
}

QQuaternion* QAudioListener_Rotation(const QAudioListener* self) {
	return new QQuaternion(self->rotation());
}

QAudioEngine* QAudioListener_Engine(const QAudioListener* self) {
	return self->engine();
}

void QAudioListener_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioListener*>( (QAudioListener*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QAudioListener_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAudioListener*)(self) )->virtualbase_MetaObject();
}

void QAudioListener_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioListener*>( (QAudioListener*)(self) )->handle__Metacast = slot;
}

void* QAudioListener_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAudioListener*)(self) )->virtualbase_Metacast(param1);
}

void QAudioListener_Delete(QAudioListener* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAudioListener*>( self );
	} else {
		delete self;
	}
}

