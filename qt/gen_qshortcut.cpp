// +build ignore

#include <QEvent>
#include <QKeySequence>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QShortcut>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qshortcut.h>
#include "gen_qshortcut.h"
class MiqtVirtualQShortcut : public virtual QShortcut {
public:
MiqtVirtualQShortcut(QObject* parent): QShortcut(parent) {};
MiqtVirtualQShortcut(const QKeySequence& key, QObject* parent): QShortcut(key, parent) {};
MiqtVirtualQShortcut(QKeySequence::StandardKey key, QObject* parent): QShortcut(key, parent) {};
MiqtVirtualQShortcut(const QKeySequence& key, QObject* parent, const char* member): QShortcut(key, parent, member) {};
MiqtVirtualQShortcut(const QKeySequence& key, QObject* parent, const char* member, const char* ambiguousMember): QShortcut(key, parent, member, ambiguousMember) {};
MiqtVirtualQShortcut(const QKeySequence& key, QObject* parent, const char* member, const char* ambiguousMember, Qt::ShortcutContext context): QShortcut(key, parent, member, ambiguousMember, context) {};
MiqtVirtualQShortcut(QKeySequence::StandardKey key, QObject* parent, const char* member): QShortcut(key, parent, member) {};
MiqtVirtualQShortcut(QKeySequence::StandardKey key, QObject* parent, const char* member, const char* ambiguousMember): QShortcut(key, parent, member, ambiguousMember) {};
MiqtVirtualQShortcut(QKeySequence::StandardKey key, QObject* parent, const char* member, const char* ambiguousMember, Qt::ShortcutContext context): QShortcut(key, parent, member, ambiguousMember, context) {};

virtual ~MiqtVirtualQShortcut() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QShortcut::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QShortcut_MetaObject(const_cast<MiqtVirtualQShortcut*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QShortcut::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QShortcut::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QShortcut_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QShortcut::qt_metacast(param1);

	}
};
QShortcut* QShortcut_new(QObject* parent) {
return new MiqtVirtualQShortcut(parent);
}
QShortcut* QShortcut_new2(QKeySequence* key, QObject* parent) {
return new MiqtVirtualQShortcut(*key, parent);
}
QShortcut* QShortcut_new3(int key, QObject* parent) {
return new MiqtVirtualQShortcut(static_cast<QKeySequence::StandardKey>(key), parent);
}
QShortcut* QShortcut_new4(QKeySequence* key, QObject* parent, const char* member) {
return new MiqtVirtualQShortcut(*key, parent, member);
}
QShortcut* QShortcut_new5(QKeySequence* key, QObject* parent, const char* member, const char* ambiguousMember) {
return new MiqtVirtualQShortcut(*key, parent, member, ambiguousMember);
}
QShortcut* QShortcut_new6(QKeySequence* key, QObject* parent, const char* member, const char* ambiguousMember, int context) {
return new MiqtVirtualQShortcut(*key, parent, member, ambiguousMember, static_cast<Qt::ShortcutContext>(context));
}
QShortcut* QShortcut_new7(int key, QObject* parent, const char* member) {
return new MiqtVirtualQShortcut(static_cast<QKeySequence::StandardKey>(key), parent, member);
}
QShortcut* QShortcut_new8(int key, QObject* parent, const char* member, const char* ambiguousMember) {
return new MiqtVirtualQShortcut(static_cast<QKeySequence::StandardKey>(key), parent, member, ambiguousMember);
}
QShortcut* QShortcut_new9(int key, QObject* parent, const char* member, const char* ambiguousMember, int context) {
return new MiqtVirtualQShortcut(static_cast<QKeySequence::StandardKey>(key), parent, member, ambiguousMember, static_cast<Qt::ShortcutContext>(context));
}
void QShortcut_virtbase(QShortcut* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QShortcut_MetaObject(const QShortcut* self) {
	return (QMetaObject*) self->metaObject();
}
void* QShortcut_Metacast(QShortcut* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QShortcut_Tr(const char* s) {
	QString _ret = QShortcut::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QShortcut_SetKey(QShortcut* self, QKeySequence* key) {
	self->setKey(*key);
}
QKeySequence* QShortcut_Key(const QShortcut* self) {
	return new QKeySequence(self->key());
}
void QShortcut_SetKeys(QShortcut* self, int key) {
	self->setKeys(static_cast<QKeySequence::StandardKey>(key));
}
void QShortcut_SetKeysWithKeys(QShortcut* self, struct miqt_array /* of QKeySequence* */  keys) {
	QList<QKeySequence> keys_QList;
	keys_QList.reserve(keys.len);
	QKeySequence** keys_arr = static_cast<QKeySequence**>(keys.data);
	for(size_t i = 0; i < keys.len; ++i) {
		keys_QList.push_back(*(keys_arr[i]));
	}
	self->setKeys(keys_QList);
}
struct miqt_array /* of QKeySequence* */  QShortcut_Keys(const QShortcut* self) {
	QList<QKeySequence> _ret = self->keys();
	// Convert QList<> from C++ memory to manually-managed C memory
	QKeySequence** _arr = static_cast<QKeySequence**>(malloc(sizeof(QKeySequence*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QKeySequence(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QShortcut_SetEnabled(QShortcut* self, bool enable) {
	self->setEnabled(enable);
}
bool QShortcut_IsEnabled(const QShortcut* self) {
	return self->isEnabled();
}
void QShortcut_SetContext(QShortcut* self, int context) {
	self->setContext(static_cast<Qt::ShortcutContext>(context));
}
int QShortcut_Context(const QShortcut* self) {
	Qt::ShortcutContext _ret = self->context();
	return static_cast<int>(_ret);
}
void QShortcut_SetAutoRepeat(QShortcut* self, bool on) {
	self->setAutoRepeat(on);
}
bool QShortcut_AutoRepeat(const QShortcut* self) {
	return self->autoRepeat();
}
int QShortcut_Id(const QShortcut* self) {
	return self->id();
}
void QShortcut_SetWhatsThis(QShortcut* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setWhatsThis(text_QString);
}
struct miqt_string QShortcut_WhatsThis(const QShortcut* self) {
	QString _ret = self->whatsThis();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QShortcut_Activated(QShortcut* self) {
	self->activated();
}
void QShortcut_connect_Activated(QShortcut* self, intptr_t slot) {
	MiqtVirtualQShortcut::connect(self, static_cast<void (QShortcut::*)()>(&QShortcut::activated), self, [=]() {
		miqt_exec_callback_QShortcut_Activated(slot);
	});
}
void QShortcut_ActivatedAmbiguously(QShortcut* self) {
	self->activatedAmbiguously();
}
void QShortcut_connect_ActivatedAmbiguously(QShortcut* self, intptr_t slot) {
	MiqtVirtualQShortcut::connect(self, static_cast<void (QShortcut::*)()>(&QShortcut::activatedAmbiguously), self, [=]() {
		miqt_exec_callback_QShortcut_ActivatedAmbiguously(slot);
	});
}
struct miqt_string QShortcut_Tr2(const char* s, const char* c) {
	QString _ret = QShortcut::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QShortcut_Tr3(const char* s, const char* c, int n) {
	QString _ret = QShortcut::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QShortcut_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQShortcut*>( (QShortcut*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QShortcut_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQShortcut*)(self) )->virtualbase_MetaObject();
}
void QShortcut_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQShortcut*>( (QShortcut*)(self) )->handle__Metacast = slot;
}
void* QShortcut_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQShortcut*)(self) )->virtualbase_Metacast(param1);
}
void QShortcut_Delete(QShortcut* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQShortcut*>( self );
	} else {
		delete self;
	}
}
