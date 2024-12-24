// +build ignore

#include <QAbstractButton>
#include <QButtonGroup>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qbuttongroup.h>
#include "gen_qbuttongroup.h"
class MiqtVirtualQButtonGroup : public virtual QButtonGroup {
public:
MiqtVirtualQButtonGroup(): QButtonGroup() {};
MiqtVirtualQButtonGroup(QObject* parent): QButtonGroup(parent) {};

virtual ~MiqtVirtualQButtonGroup() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QButtonGroup::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QButtonGroup_MetaObject(const_cast<MiqtVirtualQButtonGroup*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QButtonGroup::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QButtonGroup::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QButtonGroup_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QButtonGroup::qt_metacast(param1);

	}
};
QButtonGroup* QButtonGroup_new() {
return new MiqtVirtualQButtonGroup();
}
QButtonGroup* QButtonGroup_new2(QObject* parent) {
return new MiqtVirtualQButtonGroup(parent);
}
void QButtonGroup_virtbase(QButtonGroup* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QButtonGroup_MetaObject(const QButtonGroup* self) {
	return (QMetaObject*) self->metaObject();
}
void* QButtonGroup_Metacast(QButtonGroup* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QButtonGroup_Tr(const char* s) {
	QString _ret = QButtonGroup::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QButtonGroup_SetExclusive(QButtonGroup* self, bool exclusive) {
	self->setExclusive(exclusive);
}
bool QButtonGroup_Exclusive(const QButtonGroup* self) {
	return self->exclusive();
}
void QButtonGroup_AddButton(QButtonGroup* self, QAbstractButton* param1) {
	self->addButton(param1);
}
void QButtonGroup_RemoveButton(QButtonGroup* self, QAbstractButton* param1) {
	self->removeButton(param1);
}
struct miqt_array /* of QAbstractButton* */  QButtonGroup_Buttons(const QButtonGroup* self) {
	QList<QAbstractButton *> _ret = self->buttons();
	// Convert QList<> from C++ memory to manually-managed C memory
	QAbstractButton** _arr = static_cast<QAbstractButton**>(malloc(sizeof(QAbstractButton*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
QAbstractButton* QButtonGroup_CheckedButton(const QButtonGroup* self) {
	return self->checkedButton();
}
QAbstractButton* QButtonGroup_Button(const QButtonGroup* self, int id) {
	return self->button(static_cast<int>(id));
}
void QButtonGroup_SetId(QButtonGroup* self, QAbstractButton* button, int id) {
	self->setId(button, static_cast<int>(id));
}
int QButtonGroup_Id(const QButtonGroup* self, QAbstractButton* button) {
	return self->id(button);
}
int QButtonGroup_CheckedId(const QButtonGroup* self) {
	return self->checkedId();
}
void QButtonGroup_ButtonClicked(QButtonGroup* self, QAbstractButton* param1) {
	self->buttonClicked(param1);
}
void QButtonGroup_connect_ButtonClicked(QButtonGroup* self, intptr_t slot) {
	MiqtVirtualQButtonGroup::connect(self, static_cast<void (QButtonGroup::*)(QAbstractButton*)>(&QButtonGroup::buttonClicked), self, [=](QAbstractButton* param1) {
		QAbstractButton* sigval1 = param1;
		miqt_exec_callback_QButtonGroup_ButtonClicked(slot, sigval1);
	});
}
void QButtonGroup_ButtonPressed(QButtonGroup* self, QAbstractButton* param1) {
	self->buttonPressed(param1);
}
void QButtonGroup_connect_ButtonPressed(QButtonGroup* self, intptr_t slot) {
	MiqtVirtualQButtonGroup::connect(self, static_cast<void (QButtonGroup::*)(QAbstractButton*)>(&QButtonGroup::buttonPressed), self, [=](QAbstractButton* param1) {
		QAbstractButton* sigval1 = param1;
		miqt_exec_callback_QButtonGroup_ButtonPressed(slot, sigval1);
	});
}
void QButtonGroup_ButtonReleased(QButtonGroup* self, QAbstractButton* param1) {
	self->buttonReleased(param1);
}
void QButtonGroup_connect_ButtonReleased(QButtonGroup* self, intptr_t slot) {
	MiqtVirtualQButtonGroup::connect(self, static_cast<void (QButtonGroup::*)(QAbstractButton*)>(&QButtonGroup::buttonReleased), self, [=](QAbstractButton* param1) {
		QAbstractButton* sigval1 = param1;
		miqt_exec_callback_QButtonGroup_ButtonReleased(slot, sigval1);
	});
}
void QButtonGroup_ButtonToggled(QButtonGroup* self, QAbstractButton* param1, bool param2) {
	self->buttonToggled(param1, param2);
}
void QButtonGroup_connect_ButtonToggled(QButtonGroup* self, intptr_t slot) {
	MiqtVirtualQButtonGroup::connect(self, static_cast<void (QButtonGroup::*)(QAbstractButton*, bool)>(&QButtonGroup::buttonToggled), self, [=](QAbstractButton* param1, bool param2) {
		QAbstractButton* sigval1 = param1;
		bool sigval2 = param2;
		miqt_exec_callback_QButtonGroup_ButtonToggled(slot, sigval1, sigval2);
	});
}
void QButtonGroup_IdClicked(QButtonGroup* self, int param1) {
	self->idClicked(static_cast<int>(param1));
}
void QButtonGroup_connect_IdClicked(QButtonGroup* self, intptr_t slot) {
	MiqtVirtualQButtonGroup::connect(self, static_cast<void (QButtonGroup::*)(int)>(&QButtonGroup::idClicked), self, [=](int param1) {
		int sigval1 = param1;
		miqt_exec_callback_QButtonGroup_IdClicked(slot, sigval1);
	});
}
void QButtonGroup_IdPressed(QButtonGroup* self, int param1) {
	self->idPressed(static_cast<int>(param1));
}
void QButtonGroup_connect_IdPressed(QButtonGroup* self, intptr_t slot) {
	MiqtVirtualQButtonGroup::connect(self, static_cast<void (QButtonGroup::*)(int)>(&QButtonGroup::idPressed), self, [=](int param1) {
		int sigval1 = param1;
		miqt_exec_callback_QButtonGroup_IdPressed(slot, sigval1);
	});
}
void QButtonGroup_IdReleased(QButtonGroup* self, int param1) {
	self->idReleased(static_cast<int>(param1));
}
void QButtonGroup_connect_IdReleased(QButtonGroup* self, intptr_t slot) {
	MiqtVirtualQButtonGroup::connect(self, static_cast<void (QButtonGroup::*)(int)>(&QButtonGroup::idReleased), self, [=](int param1) {
		int sigval1 = param1;
		miqt_exec_callback_QButtonGroup_IdReleased(slot, sigval1);
	});
}
void QButtonGroup_IdToggled(QButtonGroup* self, int param1, bool param2) {
	self->idToggled(static_cast<int>(param1), param2);
}
void QButtonGroup_connect_IdToggled(QButtonGroup* self, intptr_t slot) {
	MiqtVirtualQButtonGroup::connect(self, static_cast<void (QButtonGroup::*)(int, bool)>(&QButtonGroup::idToggled), self, [=](int param1, bool param2) {
		int sigval1 = param1;
		bool sigval2 = param2;
		miqt_exec_callback_QButtonGroup_IdToggled(slot, sigval1, sigval2);
	});
}
struct miqt_string QButtonGroup_Tr2(const char* s, const char* c) {
	QString _ret = QButtonGroup::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QButtonGroup_Tr3(const char* s, const char* c, int n) {
	QString _ret = QButtonGroup::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QButtonGroup_AddButton2(QButtonGroup* self, QAbstractButton* param1, int id) {
	self->addButton(param1, static_cast<int>(id));
}
void QButtonGroup_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQButtonGroup*>( (QButtonGroup*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QButtonGroup_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQButtonGroup*)(self) )->virtualbase_MetaObject();
}
void QButtonGroup_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQButtonGroup*>( (QButtonGroup*)(self) )->handle__Metacast = slot;
}
void* QButtonGroup_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQButtonGroup*)(self) )->virtualbase_Metacast(param1);
}
void QButtonGroup_Delete(QButtonGroup* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQButtonGroup*>( self );
	} else {
		delete self;
	}
}
