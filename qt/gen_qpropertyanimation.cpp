// +build ignore

#include <QAbstractAnimation>
#include <QByteArray>
#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QPropertyAnimation>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include <QVariantAnimation>
#include <qpropertyanimation.h>
#include "gen_qpropertyanimation.h"
class MiqtVirtualQPropertyAnimation : public virtual QPropertyAnimation {
public:
MiqtVirtualQPropertyAnimation(): QPropertyAnimation() {};
MiqtVirtualQPropertyAnimation(QObject* target, const QByteArray& propertyName): QPropertyAnimation(target, propertyName) {};
MiqtVirtualQPropertyAnimation(QObject* parent): QPropertyAnimation(parent) {};
MiqtVirtualQPropertyAnimation(QObject* target, const QByteArray& propertyName, QObject* parent): QPropertyAnimation(target, propertyName, parent) {};

virtual ~MiqtVirtualQPropertyAnimation() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QPropertyAnimation::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QPropertyAnimation_MetaObject(const_cast<MiqtVirtualQPropertyAnimation*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QPropertyAnimation::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QPropertyAnimation::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QPropertyAnimation_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QPropertyAnimation::qt_metacast(param1);

	}
};
QPropertyAnimation* QPropertyAnimation_new() {
return new MiqtVirtualQPropertyAnimation();
}
QPropertyAnimation* QPropertyAnimation_new2(QObject* target, struct miqt_string propertyName) {
QByteArray propertyName_QByteArray(propertyName.data, propertyName.len);
	return new MiqtVirtualQPropertyAnimation(target, propertyName_QByteArray);
}
QPropertyAnimation* QPropertyAnimation_new3(QObject* parent) {
return new MiqtVirtualQPropertyAnimation(parent);
}
QPropertyAnimation* QPropertyAnimation_new4(QObject* target, struct miqt_string propertyName, QObject* parent) {
QByteArray propertyName_QByteArray(propertyName.data, propertyName.len);
	return new MiqtVirtualQPropertyAnimation(target, propertyName_QByteArray, parent);
}
void QPropertyAnimation_virtbase(QPropertyAnimation* src
, QVariantAnimation** outptr_QVariantAnimation
) {
*outptr_QVariantAnimation = static_cast<QVariantAnimation*>(src);
}
QMetaObject* QPropertyAnimation_MetaObject(const QPropertyAnimation* self) {
	return (QMetaObject*) self->metaObject();
}
void* QPropertyAnimation_Metacast(QPropertyAnimation* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QPropertyAnimation_Tr(const char* s) {
	QString _ret = QPropertyAnimation::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QObject* QPropertyAnimation_TargetObject(const QPropertyAnimation* self) {
	return self->targetObject();
}
void QPropertyAnimation_SetTargetObject(QPropertyAnimation* self, QObject* target) {
	self->setTargetObject(target);
}
struct miqt_string QPropertyAnimation_PropertyName(const QPropertyAnimation* self) {
	QByteArray _qb = self->propertyName();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}
void QPropertyAnimation_SetPropertyName(QPropertyAnimation* self, struct miqt_string propertyName) {
	QByteArray propertyName_QByteArray(propertyName.data, propertyName.len);
	self->setPropertyName(propertyName_QByteArray);
}
struct miqt_string QPropertyAnimation_Tr2(const char* s, const char* c) {
	QString _ret = QPropertyAnimation::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QPropertyAnimation_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPropertyAnimation::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPropertyAnimation_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPropertyAnimation*>( (QPropertyAnimation*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QPropertyAnimation_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQPropertyAnimation*)(self) )->virtualbase_MetaObject();
}
void QPropertyAnimation_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPropertyAnimation*>( (QPropertyAnimation*)(self) )->handle__Metacast = slot;
}
void* QPropertyAnimation_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQPropertyAnimation*)(self) )->virtualbase_Metacast(param1);
}
void QPropertyAnimation_Delete(QPropertyAnimation* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQPropertyAnimation*>( self );
	} else {
		delete self;
	}
}
