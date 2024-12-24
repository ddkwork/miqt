// +build ignore

#include <QAbstractAnimation>
#include <QAnimationGroup>
#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qanimationgroup.h>
#include "gen_qanimationgroup.h"
class MiqtVirtualQAnimationGroup : public virtual QAnimationGroup {
public:
MiqtVirtualQAnimationGroup(): QAnimationGroup() {};
MiqtVirtualQAnimationGroup(QObject* parent): QAnimationGroup(parent) {};

virtual ~MiqtVirtualQAnimationGroup() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QAnimationGroup::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QAnimationGroup_MetaObject(const_cast<MiqtVirtualQAnimationGroup*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAnimationGroup::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QAnimationGroup::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAnimationGroup_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAnimationGroup::qt_metacast(param1);

	}
};
QAnimationGroup* QAnimationGroup_new() {
return new MiqtVirtualQAnimationGroup();
}
QAnimationGroup* QAnimationGroup_new2(QObject* parent) {
return new MiqtVirtualQAnimationGroup(parent);
}
void QAnimationGroup_virtbase(QAnimationGroup* src
, QAbstractAnimation** outptr_QAbstractAnimation
) {
*outptr_QAbstractAnimation = static_cast<QAbstractAnimation*>(src);
}
QMetaObject* QAnimationGroup_MetaObject(const QAnimationGroup* self) {
	return (QMetaObject*) self->metaObject();
}
void* QAnimationGroup_Metacast(QAnimationGroup* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QAnimationGroup_Tr(const char* s) {
	QString _ret = QAnimationGroup::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QAbstractAnimation* QAnimationGroup_AnimationAt(const QAnimationGroup* self, int index) {
	return self->animationAt(static_cast<int>(index));
}
int QAnimationGroup_AnimationCount(const QAnimationGroup* self) {
	return self->animationCount();
}
int QAnimationGroup_IndexOfAnimation(const QAnimationGroup* self, QAbstractAnimation* animation) {
	return self->indexOfAnimation(animation);
}
void QAnimationGroup_AddAnimation(QAnimationGroup* self, QAbstractAnimation* animation) {
	self->addAnimation(animation);
}
void QAnimationGroup_InsertAnimation(QAnimationGroup* self, int index, QAbstractAnimation* animation) {
	self->insertAnimation(static_cast<int>(index), animation);
}
void QAnimationGroup_RemoveAnimation(QAnimationGroup* self, QAbstractAnimation* animation) {
	self->removeAnimation(animation);
}
QAbstractAnimation* QAnimationGroup_TakeAnimation(QAnimationGroup* self, int index) {
	return self->takeAnimation(static_cast<int>(index));
}
void QAnimationGroup_Clear(QAnimationGroup* self) {
	self->clear();
}
struct miqt_string QAnimationGroup_Tr2(const char* s, const char* c) {
	QString _ret = QAnimationGroup::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QAnimationGroup_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAnimationGroup::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QAnimationGroup_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAnimationGroup*>( (QAnimationGroup*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QAnimationGroup_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAnimationGroup*)(self) )->virtualbase_MetaObject();
}
void QAnimationGroup_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAnimationGroup*>( (QAnimationGroup*)(self) )->handle__Metacast = slot;
}
void* QAnimationGroup_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAnimationGroup*)(self) )->virtualbase_Metacast(param1);
}
void QAnimationGroup_Delete(QAnimationGroup* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAnimationGroup*>( self );
	} else {
		delete self;
	}
}
