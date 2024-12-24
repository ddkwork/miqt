// +build ignore

#include <QEvent>
#include <QGesture>
#include <QGestureEvent>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QPanGesture>
#include <QPinchGesture>
#include <QPointF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QSwipeGesture>
#include <QTapAndHoldGesture>
#include <QTapGesture>
#include <QWidget>
#include <qgesture.h>
#include "gen_qgesture.h"
class MiqtVirtualQGesture : public virtual QGesture {
public:
MiqtVirtualQGesture(): QGesture() {};
MiqtVirtualQGesture(QObject* parent): QGesture(parent) {};

virtual ~MiqtVirtualQGesture() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QGesture::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QGesture_MetaObject(const_cast<MiqtVirtualQGesture*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QGesture::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QGesture::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QGesture_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QGesture::qt_metacast(param1);

	}
};
QGesture* QGesture_new() {
return new MiqtVirtualQGesture();
}
QGesture* QGesture_new2(QObject* parent) {
return new MiqtVirtualQGesture(parent);
}
void QGesture_virtbase(QGesture* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QGesture_MetaObject(const QGesture* self) {
	return (QMetaObject*) self->metaObject();
}
void* QGesture_Metacast(QGesture* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QGesture_Tr(const char* s) {
	QString _ret = QGesture::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QGesture_GestureType(const QGesture* self) {
	Qt::GestureType _ret = self->gestureType();
	return static_cast<int>(_ret);
}
int QGesture_State(const QGesture* self) {
	Qt::GestureState _ret = self->state();
	return static_cast<int>(_ret);
}
QPointF* QGesture_HotSpot(const QGesture* self) {
	return new QPointF(self->hotSpot());
}
void QGesture_SetHotSpot(QGesture* self, QPointF* value) {
	self->setHotSpot(*value);
}
bool QGesture_HasHotSpot(const QGesture* self) {
	return self->hasHotSpot();
}
void QGesture_UnsetHotSpot(QGesture* self) {
	self->unsetHotSpot();
}
void QGesture_SetGestureCancelPolicy(QGesture* self, GestureCancelPolicy policy) {
	self->setGestureCancelPolicy(policy);
}
GestureCancelPolicy QGesture_GestureCancelPolicy(const QGesture* self) {
	return self->gestureCancelPolicy();
}
struct miqt_string QGesture_Tr2(const char* s, const char* c) {
	QString _ret = QGesture::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QGesture_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGesture::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QGesture_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGesture*>( (QGesture*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QGesture_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQGesture*)(self) )->virtualbase_MetaObject();
}
void QGesture_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGesture*>( (QGesture*)(self) )->handle__Metacast = slot;
}
void* QGesture_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQGesture*)(self) )->virtualbase_Metacast(param1);
}
void QGesture_Delete(QGesture* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGesture*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQPanGesture : public virtual QPanGesture {
public:
MiqtVirtualQPanGesture(): QPanGesture() {};
MiqtVirtualQPanGesture(QObject* parent): QPanGesture(parent) {};

virtual ~MiqtVirtualQPanGesture() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QPanGesture::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QPanGesture_MetaObject(const_cast<MiqtVirtualQPanGesture*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QPanGesture::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QPanGesture::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QPanGesture_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QPanGesture::qt_metacast(param1);

	}
};
QPanGesture* QPanGesture_new() {
return new MiqtVirtualQPanGesture();
}
QPanGesture* QPanGesture_new2(QObject* parent) {
return new MiqtVirtualQPanGesture(parent);
}
void QPanGesture_virtbase(QPanGesture* src
, QGesture** outptr_QGesture
) {
*outptr_QGesture = static_cast<QGesture*>(src);
}
QMetaObject* QPanGesture_MetaObject(const QPanGesture* self) {
	return (QMetaObject*) self->metaObject();
}
void* QPanGesture_Metacast(QPanGesture* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QPanGesture_Tr(const char* s) {
	QString _ret = QPanGesture::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QPointF* QPanGesture_LastOffset(const QPanGesture* self) {
	return new QPointF(self->lastOffset());
}
QPointF* QPanGesture_Offset(const QPanGesture* self) {
	return new QPointF(self->offset());
}
QPointF* QPanGesture_Delta(const QPanGesture* self) {
	return new QPointF(self->delta());
}
double QPanGesture_Acceleration(const QPanGesture* self) {
	qreal _ret = self->acceleration();
	return static_cast<double>(_ret);
}
void QPanGesture_SetLastOffset(QPanGesture* self, QPointF* value) {
	self->setLastOffset(*value);
}
void QPanGesture_SetOffset(QPanGesture* self, QPointF* value) {
	self->setOffset(*value);
}
void QPanGesture_SetAcceleration(QPanGesture* self, double value) {
	self->setAcceleration(static_cast<qreal>(value));
}
struct miqt_string QPanGesture_Tr2(const char* s, const char* c) {
	QString _ret = QPanGesture::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QPanGesture_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPanGesture::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPanGesture_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPanGesture*>( (QPanGesture*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QPanGesture_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQPanGesture*)(self) )->virtualbase_MetaObject();
}
void QPanGesture_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPanGesture*>( (QPanGesture*)(self) )->handle__Metacast = slot;
}
void* QPanGesture_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQPanGesture*)(self) )->virtualbase_Metacast(param1);
}
void QPanGesture_Delete(QPanGesture* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQPanGesture*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQPinchGesture : public virtual QPinchGesture {
public:
MiqtVirtualQPinchGesture(): QPinchGesture() {};
MiqtVirtualQPinchGesture(QObject* parent): QPinchGesture(parent) {};

virtual ~MiqtVirtualQPinchGesture() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QPinchGesture::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QPinchGesture_MetaObject(const_cast<MiqtVirtualQPinchGesture*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QPinchGesture::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QPinchGesture::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QPinchGesture_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QPinchGesture::qt_metacast(param1);

	}
};
QPinchGesture* QPinchGesture_new() {
return new MiqtVirtualQPinchGesture();
}
QPinchGesture* QPinchGesture_new2(QObject* parent) {
return new MiqtVirtualQPinchGesture(parent);
}
void QPinchGesture_virtbase(QPinchGesture* src
, QGesture** outptr_QGesture
) {
*outptr_QGesture = static_cast<QGesture*>(src);
}
QMetaObject* QPinchGesture_MetaObject(const QPinchGesture* self) {
	return (QMetaObject*) self->metaObject();
}
void* QPinchGesture_Metacast(QPinchGesture* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QPinchGesture_Tr(const char* s) {
	QString _ret = QPinchGesture::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
ChangeFlags QPinchGesture_TotalChangeFlags(const QPinchGesture* self) {
	return self->totalChangeFlags();
}
void QPinchGesture_SetTotalChangeFlags(QPinchGesture* self, ChangeFlags value) {
	self->setTotalChangeFlags(value);
}
ChangeFlags QPinchGesture_ChangeFlags(const QPinchGesture* self) {
	return self->changeFlags();
}
void QPinchGesture_SetChangeFlags(QPinchGesture* self, ChangeFlags value) {
	self->setChangeFlags(value);
}
QPointF* QPinchGesture_StartCenterPoint(const QPinchGesture* self) {
	return new QPointF(self->startCenterPoint());
}
QPointF* QPinchGesture_LastCenterPoint(const QPinchGesture* self) {
	return new QPointF(self->lastCenterPoint());
}
QPointF* QPinchGesture_CenterPoint(const QPinchGesture* self) {
	return new QPointF(self->centerPoint());
}
void QPinchGesture_SetStartCenterPoint(QPinchGesture* self, QPointF* value) {
	self->setStartCenterPoint(*value);
}
void QPinchGesture_SetLastCenterPoint(QPinchGesture* self, QPointF* value) {
	self->setLastCenterPoint(*value);
}
void QPinchGesture_SetCenterPoint(QPinchGesture* self, QPointF* value) {
	self->setCenterPoint(*value);
}
double QPinchGesture_TotalScaleFactor(const QPinchGesture* self) {
	qreal _ret = self->totalScaleFactor();
	return static_cast<double>(_ret);
}
double QPinchGesture_LastScaleFactor(const QPinchGesture* self) {
	qreal _ret = self->lastScaleFactor();
	return static_cast<double>(_ret);
}
double QPinchGesture_ScaleFactor(const QPinchGesture* self) {
	qreal _ret = self->scaleFactor();
	return static_cast<double>(_ret);
}
void QPinchGesture_SetTotalScaleFactor(QPinchGesture* self, double value) {
	self->setTotalScaleFactor(static_cast<qreal>(value));
}
void QPinchGesture_SetLastScaleFactor(QPinchGesture* self, double value) {
	self->setLastScaleFactor(static_cast<qreal>(value));
}
void QPinchGesture_SetScaleFactor(QPinchGesture* self, double value) {
	self->setScaleFactor(static_cast<qreal>(value));
}
double QPinchGesture_TotalRotationAngle(const QPinchGesture* self) {
	qreal _ret = self->totalRotationAngle();
	return static_cast<double>(_ret);
}
double QPinchGesture_LastRotationAngle(const QPinchGesture* self) {
	qreal _ret = self->lastRotationAngle();
	return static_cast<double>(_ret);
}
double QPinchGesture_RotationAngle(const QPinchGesture* self) {
	qreal _ret = self->rotationAngle();
	return static_cast<double>(_ret);
}
void QPinchGesture_SetTotalRotationAngle(QPinchGesture* self, double value) {
	self->setTotalRotationAngle(static_cast<qreal>(value));
}
void QPinchGesture_SetLastRotationAngle(QPinchGesture* self, double value) {
	self->setLastRotationAngle(static_cast<qreal>(value));
}
void QPinchGesture_SetRotationAngle(QPinchGesture* self, double value) {
	self->setRotationAngle(static_cast<qreal>(value));
}
struct miqt_string QPinchGesture_Tr2(const char* s, const char* c) {
	QString _ret = QPinchGesture::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QPinchGesture_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPinchGesture::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPinchGesture_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPinchGesture*>( (QPinchGesture*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QPinchGesture_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQPinchGesture*)(self) )->virtualbase_MetaObject();
}
void QPinchGesture_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPinchGesture*>( (QPinchGesture*)(self) )->handle__Metacast = slot;
}
void* QPinchGesture_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQPinchGesture*)(self) )->virtualbase_Metacast(param1);
}
void QPinchGesture_Delete(QPinchGesture* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQPinchGesture*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQSwipeGesture : public virtual QSwipeGesture {
public:
MiqtVirtualQSwipeGesture(): QSwipeGesture() {};
MiqtVirtualQSwipeGesture(QObject* parent): QSwipeGesture(parent) {};

virtual ~MiqtVirtualQSwipeGesture() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QSwipeGesture::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QSwipeGesture_MetaObject(const_cast<MiqtVirtualQSwipeGesture*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSwipeGesture::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QSwipeGesture::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSwipeGesture_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSwipeGesture::qt_metacast(param1);

	}
};
QSwipeGesture* QSwipeGesture_new() {
return new MiqtVirtualQSwipeGesture();
}
QSwipeGesture* QSwipeGesture_new2(QObject* parent) {
return new MiqtVirtualQSwipeGesture(parent);
}
void QSwipeGesture_virtbase(QSwipeGesture* src
, QGesture** outptr_QGesture
) {
*outptr_QGesture = static_cast<QGesture*>(src);
}
QMetaObject* QSwipeGesture_MetaObject(const QSwipeGesture* self) {
	return (QMetaObject*) self->metaObject();
}
void* QSwipeGesture_Metacast(QSwipeGesture* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QSwipeGesture_Tr(const char* s) {
	QString _ret = QSwipeGesture::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
SwipeDirection QSwipeGesture_HorizontalDirection(const QSwipeGesture* self) {
	return self->horizontalDirection();
}
SwipeDirection QSwipeGesture_VerticalDirection(const QSwipeGesture* self) {
	return self->verticalDirection();
}
double QSwipeGesture_SwipeAngle(const QSwipeGesture* self) {
	qreal _ret = self->swipeAngle();
	return static_cast<double>(_ret);
}
void QSwipeGesture_SetSwipeAngle(QSwipeGesture* self, double value) {
	self->setSwipeAngle(static_cast<qreal>(value));
}
struct miqt_string QSwipeGesture_Tr2(const char* s, const char* c) {
	QString _ret = QSwipeGesture::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QSwipeGesture_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSwipeGesture::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSwipeGesture_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSwipeGesture*>( (QSwipeGesture*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QSwipeGesture_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSwipeGesture*)(self) )->virtualbase_MetaObject();
}
void QSwipeGesture_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSwipeGesture*>( (QSwipeGesture*)(self) )->handle__Metacast = slot;
}
void* QSwipeGesture_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSwipeGesture*)(self) )->virtualbase_Metacast(param1);
}
void QSwipeGesture_Delete(QSwipeGesture* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSwipeGesture*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQTapGesture : public virtual QTapGesture {
public:
MiqtVirtualQTapGesture(): QTapGesture() {};
MiqtVirtualQTapGesture(QObject* parent): QTapGesture(parent) {};

virtual ~MiqtVirtualQTapGesture() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QTapGesture::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QTapGesture_MetaObject(const_cast<MiqtVirtualQTapGesture*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QTapGesture::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QTapGesture::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QTapGesture_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QTapGesture::qt_metacast(param1);

	}
};
QTapGesture* QTapGesture_new() {
return new MiqtVirtualQTapGesture();
}
QTapGesture* QTapGesture_new2(QObject* parent) {
return new MiqtVirtualQTapGesture(parent);
}
void QTapGesture_virtbase(QTapGesture* src
, QGesture** outptr_QGesture
) {
*outptr_QGesture = static_cast<QGesture*>(src);
}
QMetaObject* QTapGesture_MetaObject(const QTapGesture* self) {
	return (QMetaObject*) self->metaObject();
}
void* QTapGesture_Metacast(QTapGesture* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QTapGesture_Tr(const char* s) {
	QString _ret = QTapGesture::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QPointF* QTapGesture_Position(const QTapGesture* self) {
	return new QPointF(self->position());
}
void QTapGesture_SetPosition(QTapGesture* self, QPointF* pos) {
	self->setPosition(*pos);
}
struct miqt_string QTapGesture_Tr2(const char* s, const char* c) {
	QString _ret = QTapGesture::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QTapGesture_Tr3(const char* s, const char* c, int n) {
	QString _ret = QTapGesture::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTapGesture_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTapGesture*>( (QTapGesture*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QTapGesture_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQTapGesture*)(self) )->virtualbase_MetaObject();
}
void QTapGesture_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTapGesture*>( (QTapGesture*)(self) )->handle__Metacast = slot;
}
void* QTapGesture_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQTapGesture*)(self) )->virtualbase_Metacast(param1);
}
void QTapGesture_Delete(QTapGesture* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQTapGesture*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQTapAndHoldGesture : public virtual QTapAndHoldGesture {
public:
MiqtVirtualQTapAndHoldGesture(): QTapAndHoldGesture() {};
MiqtVirtualQTapAndHoldGesture(QObject* parent): QTapAndHoldGesture(parent) {};

virtual ~MiqtVirtualQTapAndHoldGesture() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QTapAndHoldGesture::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QTapAndHoldGesture_MetaObject(const_cast<MiqtVirtualQTapAndHoldGesture*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QTapAndHoldGesture::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QTapAndHoldGesture::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QTapAndHoldGesture_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QTapAndHoldGesture::qt_metacast(param1);

	}
};
QTapAndHoldGesture* QTapAndHoldGesture_new() {
return new MiqtVirtualQTapAndHoldGesture();
}
QTapAndHoldGesture* QTapAndHoldGesture_new2(QObject* parent) {
return new MiqtVirtualQTapAndHoldGesture(parent);
}
void QTapAndHoldGesture_virtbase(QTapAndHoldGesture* src
, QGesture** outptr_QGesture
) {
*outptr_QGesture = static_cast<QGesture*>(src);
}
QMetaObject* QTapAndHoldGesture_MetaObject(const QTapAndHoldGesture* self) {
	return (QMetaObject*) self->metaObject();
}
void* QTapAndHoldGesture_Metacast(QTapAndHoldGesture* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QTapAndHoldGesture_Tr(const char* s) {
	QString _ret = QTapAndHoldGesture::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QPointF* QTapAndHoldGesture_Position(const QTapAndHoldGesture* self) {
	return new QPointF(self->position());
}
void QTapAndHoldGesture_SetPosition(QTapAndHoldGesture* self, QPointF* pos) {
	self->setPosition(*pos);
}
void QTapAndHoldGesture_SetTimeout(int msecs) {
	QTapAndHoldGesture::setTimeout(static_cast<int>(msecs));
}
int QTapAndHoldGesture_Timeout() {
	return QTapAndHoldGesture::timeout();
}
struct miqt_string QTapAndHoldGesture_Tr2(const char* s, const char* c) {
	QString _ret = QTapAndHoldGesture::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QTapAndHoldGesture_Tr3(const char* s, const char* c, int n) {
	QString _ret = QTapAndHoldGesture::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTapAndHoldGesture_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTapAndHoldGesture*>( (QTapAndHoldGesture*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QTapAndHoldGesture_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQTapAndHoldGesture*)(self) )->virtualbase_MetaObject();
}
void QTapAndHoldGesture_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTapAndHoldGesture*>( (QTapAndHoldGesture*)(self) )->handle__Metacast = slot;
}
void* QTapAndHoldGesture_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQTapAndHoldGesture*)(self) )->virtualbase_Metacast(param1);
}
void QTapAndHoldGesture_Delete(QTapAndHoldGesture* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQTapAndHoldGesture*>( self );
	} else {
		delete self;
	}
}
QGestureEvent* QGestureEvent_new(struct miqt_array /* of QGesture* */  gestures) {
QList<QGesture *> gestures_QList;
	gestures_QList.reserve(gestures.len);
	QGesture** gestures_arr = static_cast<QGesture**>(gestures.data);
	for(size_t i = 0; i < gestures.len; ++i) {
		gestures_QList.push_back(gestures_arr[i]);
	}
	return new QGestureEvent(gestures_QList);
}
QGestureEvent* QGestureEvent_new2(QGestureEvent* param1) {
return new QGestureEvent(*param1);
}
void QGestureEvent_virtbase(QGestureEvent* src
, QEvent** outptr_QEvent
) {
*outptr_QEvent = static_cast<QEvent*>(src);
}
struct miqt_array /* of QGesture* */  QGestureEvent_Gestures(const QGestureEvent* self) {
	QList<QGesture *> _ret = self->gestures();
	// Convert QList<> from C++ memory to manually-managed C memory
	QGesture** _arr = static_cast<QGesture**>(malloc(sizeof(QGesture*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
QGesture* QGestureEvent_Gesture(const QGestureEvent* self, int typeVal) {
	return self->gesture(static_cast<Qt::GestureType>(typeVal));
}
struct miqt_array /* of QGesture* */  QGestureEvent_ActiveGestures(const QGestureEvent* self) {
	QList<QGesture *> _ret = self->activeGestures();
	// Convert QList<> from C++ memory to manually-managed C memory
	QGesture** _arr = static_cast<QGesture**>(malloc(sizeof(QGesture*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
struct miqt_array /* of QGesture* */  QGestureEvent_CanceledGestures(const QGestureEvent* self) {
	QList<QGesture *> _ret = self->canceledGestures();
	// Convert QList<> from C++ memory to manually-managed C memory
	QGesture** _arr = static_cast<QGesture**>(malloc(sizeof(QGesture*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QGestureEvent_SetAccepted(QGestureEvent* self, QGesture* param1, bool param2) {
	self->setAccepted(param1, param2);
}
void QGestureEvent_Accept(QGestureEvent* self, QGesture* param1) {
	self->accept(param1);
}
void QGestureEvent_Ignore(QGestureEvent* self, QGesture* param1) {
	self->ignore(param1);
}
bool QGestureEvent_IsAccepted(const QGestureEvent* self, QGesture* param1) {
	return self->isAccepted(param1);
}
void QGestureEvent_SetAccepted2(QGestureEvent* self, int param1, bool param2) {
	self->setAccepted(static_cast<Qt::GestureType>(param1), param2);
}
void QGestureEvent_AcceptWithQtGestureType(QGestureEvent* self, int param1) {
	self->accept(static_cast<Qt::GestureType>(param1));
}
void QGestureEvent_IgnoreWithQtGestureType(QGestureEvent* self, int param1) {
	self->ignore(static_cast<Qt::GestureType>(param1));
}
bool QGestureEvent_IsAcceptedWithQtGestureType(const QGestureEvent* self, int param1) {
	return self->isAccepted(static_cast<Qt::GestureType>(param1));
}
void QGestureEvent_SetWidget(QGestureEvent* self, QWidget* widget) {
	self->setWidget(widget);
}
QWidget* QGestureEvent_Widget(const QGestureEvent* self) {
	return self->widget();
}
QPointF* QGestureEvent_MapToGraphicsScene(const QGestureEvent* self, QPointF* gesturePoint) {
	return new QPointF(self->mapToGraphicsScene(*gesturePoint));
}
void QGestureEvent_Delete(QGestureEvent* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QGestureEvent*>( self );
	} else {
		delete self;
	}
}
