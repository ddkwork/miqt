// +build ignore

#include <QAudioEngine>
#include <QAudioRoom>
#include <QMetaObject>
#include <QObject>
#include <QQuaternion>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVector3D>
#include <qaudioroom.h>
#include "gen_qaudioroom.h"

class MiqtVirtualQAudioRoom : public virtual QAudioRoom {
public:

	MiqtVirtualQAudioRoom(QAudioEngine* engine): QAudioRoom(engine) {};

	virtual ~MiqtVirtualQAudioRoom() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QAudioRoom::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QAudioRoom_MetaObject(const_cast<MiqtVirtualQAudioRoom*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAudioRoom::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QAudioRoom::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAudioRoom_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAudioRoom::qt_metacast(param1);

	}

};

QAudioRoom* QAudioRoom_new(QAudioEngine* engine) {
	return new MiqtVirtualQAudioRoom(engine);
}

void QAudioRoom_virtbase(QAudioRoom* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QAudioRoom_MetaObject(const QAudioRoom* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAudioRoom_Metacast(QAudioRoom* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAudioRoom_Tr(const char* s) {
	QString _ret = QAudioRoom::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAudioRoom_SetPosition(QAudioRoom* self, QVector3D* pos) {
	self->setPosition(*pos);
}

QVector3D* QAudioRoom_Position(const QAudioRoom* self) {
	return new QVector3D(self->position());
}

void QAudioRoom_SetDimensions(QAudioRoom* self, QVector3D* dim) {
	self->setDimensions(*dim);
}

QVector3D* QAudioRoom_Dimensions(const QAudioRoom* self) {
	return new QVector3D(self->dimensions());
}

void QAudioRoom_SetRotation(QAudioRoom* self, QQuaternion* q) {
	self->setRotation(*q);
}

QQuaternion* QAudioRoom_Rotation(const QAudioRoom* self) {
	return new QQuaternion(self->rotation());
}

void QAudioRoom_SetWallMaterial(QAudioRoom* self, Wall wall, Material material) {
	self->setWallMaterial(wall, material);
}

Material QAudioRoom_WallMaterial(const QAudioRoom* self, Wall wall) {
	return self->wallMaterial(wall);
}

void QAudioRoom_SetReflectionGain(QAudioRoom* self, float factor) {
	self->setReflectionGain(static_cast<float>(factor));
}

float QAudioRoom_ReflectionGain(const QAudioRoom* self) {
	return self->reflectionGain();
}

void QAudioRoom_SetReverbGain(QAudioRoom* self, float factor) {
	self->setReverbGain(static_cast<float>(factor));
}

float QAudioRoom_ReverbGain(const QAudioRoom* self) {
	return self->reverbGain();
}

void QAudioRoom_SetReverbTime(QAudioRoom* self, float factor) {
	self->setReverbTime(static_cast<float>(factor));
}

float QAudioRoom_ReverbTime(const QAudioRoom* self) {
	return self->reverbTime();
}

void QAudioRoom_SetReverbBrightness(QAudioRoom* self, float factor) {
	self->setReverbBrightness(static_cast<float>(factor));
}

float QAudioRoom_ReverbBrightness(const QAudioRoom* self) {
	return self->reverbBrightness();
}

void QAudioRoom_PositionChanged(QAudioRoom* self) {
	self->positionChanged();
}

void QAudioRoom_connect_PositionChanged(QAudioRoom* self, intptr_t slot) {
	MiqtVirtualQAudioRoom::connect(self, static_cast<void (QAudioRoom::*)()>(&QAudioRoom::positionChanged), self, [=]() {
		miqt_exec_callback_QAudioRoom_PositionChanged(slot);
	});
}

void QAudioRoom_DimensionsChanged(QAudioRoom* self) {
	self->dimensionsChanged();
}

void QAudioRoom_connect_DimensionsChanged(QAudioRoom* self, intptr_t slot) {
	MiqtVirtualQAudioRoom::connect(self, static_cast<void (QAudioRoom::*)()>(&QAudioRoom::dimensionsChanged), self, [=]() {
		miqt_exec_callback_QAudioRoom_DimensionsChanged(slot);
	});
}

void QAudioRoom_RotationChanged(QAudioRoom* self) {
	self->rotationChanged();
}

void QAudioRoom_connect_RotationChanged(QAudioRoom* self, intptr_t slot) {
	MiqtVirtualQAudioRoom::connect(self, static_cast<void (QAudioRoom::*)()>(&QAudioRoom::rotationChanged), self, [=]() {
		miqt_exec_callback_QAudioRoom_RotationChanged(slot);
	});
}

void QAudioRoom_WallsChanged(QAudioRoom* self) {
	self->wallsChanged();
}

void QAudioRoom_connect_WallsChanged(QAudioRoom* self, intptr_t slot) {
	MiqtVirtualQAudioRoom::connect(self, static_cast<void (QAudioRoom::*)()>(&QAudioRoom::wallsChanged), self, [=]() {
		miqt_exec_callback_QAudioRoom_WallsChanged(slot);
	});
}

void QAudioRoom_ReflectionGainChanged(QAudioRoom* self) {
	self->reflectionGainChanged();
}

void QAudioRoom_connect_ReflectionGainChanged(QAudioRoom* self, intptr_t slot) {
	MiqtVirtualQAudioRoom::connect(self, static_cast<void (QAudioRoom::*)()>(&QAudioRoom::reflectionGainChanged), self, [=]() {
		miqt_exec_callback_QAudioRoom_ReflectionGainChanged(slot);
	});
}

void QAudioRoom_ReverbGainChanged(QAudioRoom* self) {
	self->reverbGainChanged();
}

void QAudioRoom_connect_ReverbGainChanged(QAudioRoom* self, intptr_t slot) {
	MiqtVirtualQAudioRoom::connect(self, static_cast<void (QAudioRoom::*)()>(&QAudioRoom::reverbGainChanged), self, [=]() {
		miqt_exec_callback_QAudioRoom_ReverbGainChanged(slot);
	});
}

void QAudioRoom_ReverbTimeChanged(QAudioRoom* self) {
	self->reverbTimeChanged();
}

void QAudioRoom_connect_ReverbTimeChanged(QAudioRoom* self, intptr_t slot) {
	MiqtVirtualQAudioRoom::connect(self, static_cast<void (QAudioRoom::*)()>(&QAudioRoom::reverbTimeChanged), self, [=]() {
		miqt_exec_callback_QAudioRoom_ReverbTimeChanged(slot);
	});
}

void QAudioRoom_ReverbBrightnessChanged(QAudioRoom* self) {
	self->reverbBrightnessChanged();
}

void QAudioRoom_connect_ReverbBrightnessChanged(QAudioRoom* self, intptr_t slot) {
	MiqtVirtualQAudioRoom::connect(self, static_cast<void (QAudioRoom::*)()>(&QAudioRoom::reverbBrightnessChanged), self, [=]() {
		miqt_exec_callback_QAudioRoom_ReverbBrightnessChanged(slot);
	});
}

struct miqt_string QAudioRoom_Tr2(const char* s, const char* c) {
	QString _ret = QAudioRoom::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAudioRoom_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAudioRoom::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAudioRoom_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioRoom*>( (QAudioRoom*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QAudioRoom_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAudioRoom*)(self) )->virtualbase_MetaObject();
}

void QAudioRoom_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAudioRoom*>( (QAudioRoom*)(self) )->handle__Metacast = slot;
}

void* QAudioRoom_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAudioRoom*)(self) )->virtualbase_Metacast(param1);
}

void QAudioRoom_Delete(QAudioRoom* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAudioRoom*>( self );
	} else {
		delete self;
	}
}

