// +build ignore

#include <QAction>
#include <QActionGroup>
#include <QIcon>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qactiongroup.h>
#include "gen_qactiongroup.h"
class MiqtVirtualQActionGroup : public virtual QActionGroup {
public:
MiqtVirtualQActionGroup(QObject* parent): QActionGroup(parent) {};

virtual ~MiqtVirtualQActionGroup() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QActionGroup::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QActionGroup_MetaObject(const_cast<MiqtVirtualQActionGroup*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QActionGroup::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QActionGroup::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QActionGroup_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QActionGroup::qt_metacast(param1);

	}
};
QActionGroup* QActionGroup_new(QObject* parent) {
return new MiqtVirtualQActionGroup(parent);
}
void QActionGroup_virtbase(QActionGroup* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QActionGroup_MetaObject(const QActionGroup* self) {
	return (QMetaObject*) self->metaObject();
}
void* QActionGroup_Metacast(QActionGroup* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QActionGroup_Tr(const char* s) {
	QString _ret = QActionGroup::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QAction* QActionGroup_AddAction(QActionGroup* self, QAction* a) {
	return self->addAction(a);
}
QAction* QActionGroup_AddActionWithText(QActionGroup* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->addAction(text_QString);
}
QAction* QActionGroup_AddAction2(QActionGroup* self, QIcon* icon, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->addAction(*icon, text_QString);
}
void QActionGroup_RemoveAction(QActionGroup* self, QAction* a) {
	self->removeAction(a);
}
struct miqt_array /* of QAction* */  QActionGroup_Actions(const QActionGroup* self) {
	QList<QAction *> _ret = self->actions();
	// Convert QList<> from C++ memory to manually-managed C memory
	QAction** _arr = static_cast<QAction**>(malloc(sizeof(QAction*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
QAction* QActionGroup_CheckedAction(const QActionGroup* self) {
	return self->checkedAction();
}
bool QActionGroup_IsExclusive(const QActionGroup* self) {
	return self->isExclusive();
}
bool QActionGroup_IsEnabled(const QActionGroup* self) {
	return self->isEnabled();
}
bool QActionGroup_IsVisible(const QActionGroup* self) {
	return self->isVisible();
}
ExclusionPolicy QActionGroup_ExclusionPolicy(const QActionGroup* self) {
	return self->exclusionPolicy();
}
void QActionGroup_SetEnabled(QActionGroup* self, bool enabled) {
	self->setEnabled(enabled);
}
void QActionGroup_SetDisabled(QActionGroup* self, bool b) {
	self->setDisabled(b);
}
void QActionGroup_SetVisible(QActionGroup* self, bool visible) {
	self->setVisible(visible);
}
void QActionGroup_SetExclusive(QActionGroup* self, bool exclusive) {
	self->setExclusive(exclusive);
}
void QActionGroup_SetExclusionPolicy(QActionGroup* self, ExclusionPolicy policy) {
	self->setExclusionPolicy(policy);
}
void QActionGroup_Triggered(QActionGroup* self, QAction* param1) {
	self->triggered(param1);
}
void QActionGroup_connect_Triggered(QActionGroup* self, intptr_t slot) {
	MiqtVirtualQActionGroup::connect(self, static_cast<void (QActionGroup::*)(QAction*)>(&QActionGroup::triggered), self, [=](QAction* param1) {
		QAction* sigval1 = param1;
		miqt_exec_callback_QActionGroup_Triggered(slot, sigval1);
	});
}
void QActionGroup_Hovered(QActionGroup* self, QAction* param1) {
	self->hovered(param1);
}
void QActionGroup_connect_Hovered(QActionGroup* self, intptr_t slot) {
	MiqtVirtualQActionGroup::connect(self, static_cast<void (QActionGroup::*)(QAction*)>(&QActionGroup::hovered), self, [=](QAction* param1) {
		QAction* sigval1 = param1;
		miqt_exec_callback_QActionGroup_Hovered(slot, sigval1);
	});
}
struct miqt_string QActionGroup_Tr2(const char* s, const char* c) {
	QString _ret = QActionGroup::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QActionGroup_Tr3(const char* s, const char* c, int n) {
	QString _ret = QActionGroup::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QActionGroup_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQActionGroup*>( (QActionGroup*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QActionGroup_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQActionGroup*)(self) )->virtualbase_MetaObject();
}
void QActionGroup_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQActionGroup*>( (QActionGroup*)(self) )->handle__Metacast = slot;
}
void* QActionGroup_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQActionGroup*)(self) )->virtualbase_Metacast(param1);
}
void QActionGroup_Delete(QActionGroup* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQActionGroup*>( self );
	} else {
		delete self;
	}
}
