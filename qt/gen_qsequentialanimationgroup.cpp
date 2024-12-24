// +build ignore

#include <QAbstractAnimation>
#include <QAnimationGroup>
#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QPauseAnimation>
#include <QSequentialAnimationGroup>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qsequentialanimationgroup.h>
#include "gen_qsequentialanimationgroup.h"
class MiqtVirtualQSequentialAnimationGroup : public virtual QSequentialAnimationGroup {
public:
MiqtVirtualQSequentialAnimationGroup(): QSequentialAnimationGroup() {};
MiqtVirtualQSequentialAnimationGroup(QObject* parent): QSequentialAnimationGroup(parent) {};

virtual ~MiqtVirtualQSequentialAnimationGroup() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QSequentialAnimationGroup::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QSequentialAnimationGroup_MetaObject(const_cast<MiqtVirtualQSequentialAnimationGroup*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSequentialAnimationGroup::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QSequentialAnimationGroup::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSequentialAnimationGroup_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSequentialAnimationGroup::qt_metacast(param1);

	}
};
QSequentialAnimationGroup* QSequentialAnimationGroup_new() {
return new MiqtVirtualQSequentialAnimationGroup();
}
QSequentialAnimationGroup* QSequentialAnimationGroup_new2(QObject* parent) {
return new MiqtVirtualQSequentialAnimationGroup(parent);
}
void QSequentialAnimationGroup_virtbase(QSequentialAnimationGroup* src
, QAnimationGroup** outptr_QAnimationGroup
) {
*outptr_QAnimationGroup = static_cast<QAnimationGroup*>(src);
}
QMetaObject* QSequentialAnimationGroup_MetaObject(const QSequentialAnimationGroup* self) {
	return (QMetaObject*) self->metaObject();
}
void* QSequentialAnimationGroup_Metacast(QSequentialAnimationGroup* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QSequentialAnimationGroup_Tr(const char* s) {
	QString _ret = QSequentialAnimationGroup::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QPauseAnimation* QSequentialAnimationGroup_AddPause(QSequentialAnimationGroup* self, int msecs) {
	return self->addPause(static_cast<int>(msecs));
}
QPauseAnimation* QSequentialAnimationGroup_InsertPause(QSequentialAnimationGroup* self, int index, int msecs) {
	return self->insertPause(static_cast<int>(index), static_cast<int>(msecs));
}
QAbstractAnimation* QSequentialAnimationGroup_CurrentAnimation(const QSequentialAnimationGroup* self) {
	return self->currentAnimation();
}
int QSequentialAnimationGroup_Duration(const QSequentialAnimationGroup* self) {
	return self->duration();
}
void QSequentialAnimationGroup_CurrentAnimationChanged(QSequentialAnimationGroup* self, QAbstractAnimation* current) {
	self->currentAnimationChanged(current);
}
void QSequentialAnimationGroup_connect_CurrentAnimationChanged(QSequentialAnimationGroup* self, intptr_t slot) {
	MiqtVirtualQSequentialAnimationGroup::connect(self, static_cast<void (QSequentialAnimationGroup::*)(QAbstractAnimation*)>(&QSequentialAnimationGroup::currentAnimationChanged), self, [=](QAbstractAnimation* current) {
		QAbstractAnimation* sigval1 = current;
		miqt_exec_callback_QSequentialAnimationGroup_CurrentAnimationChanged(slot, sigval1);
	});
}
struct miqt_string QSequentialAnimationGroup_Tr2(const char* s, const char* c) {
	QString _ret = QSequentialAnimationGroup::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QSequentialAnimationGroup_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSequentialAnimationGroup::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSequentialAnimationGroup_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSequentialAnimationGroup*>( (QSequentialAnimationGroup*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QSequentialAnimationGroup_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSequentialAnimationGroup*)(self) )->virtualbase_MetaObject();
}
void QSequentialAnimationGroup_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSequentialAnimationGroup*>( (QSequentialAnimationGroup*)(self) )->handle__Metacast = slot;
}
void* QSequentialAnimationGroup_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSequentialAnimationGroup*)(self) )->virtualbase_Metacast(param1);
}
void QSequentialAnimationGroup_Delete(QSequentialAnimationGroup* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSequentialAnimationGroup*>( self );
	} else {
		delete self;
	}
}
