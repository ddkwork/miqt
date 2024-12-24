// +build ignore

#include <QEvent>
#include <QFocusEvent>
#include <QKeyCombination>
#include <QKeyEvent>
#include <QKeySequence>
#include <QKeySequenceEdit>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <QWidget>
#include <qkeysequenceedit.h>
#include "gen_qkeysequenceedit.h"

class MiqtVirtualQKeySequenceEdit : public virtual QKeySequenceEdit {
public:

	MiqtVirtualQKeySequenceEdit(QWidget* parent): QKeySequenceEdit(parent) {};
	MiqtVirtualQKeySequenceEdit(): QKeySequenceEdit() {};
	MiqtVirtualQKeySequenceEdit(const QKeySequence& keySequence): QKeySequenceEdit(keySequence) {};
	MiqtVirtualQKeySequenceEdit(const QKeySequence& keySequence, QWidget* parent): QKeySequenceEdit(keySequence, parent) {};

	virtual ~MiqtVirtualQKeySequenceEdit() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QKeySequenceEdit::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QKeySequenceEdit_MetaObject(const_cast<MiqtVirtualQKeySequenceEdit*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QKeySequenceEdit::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QKeySequenceEdit::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QKeySequenceEdit_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QKeySequenceEdit::qt_metacast(param1);

	}

};

QKeySequenceEdit* QKeySequenceEdit_new(QWidget* parent) {
	return new MiqtVirtualQKeySequenceEdit(parent);
}

QKeySequenceEdit* QKeySequenceEdit_new2() {
	return new MiqtVirtualQKeySequenceEdit();
}

QKeySequenceEdit* QKeySequenceEdit_new3(QKeySequence* keySequence) {
	return new MiqtVirtualQKeySequenceEdit(*keySequence);
}

QKeySequenceEdit* QKeySequenceEdit_new4(QKeySequence* keySequence, QWidget* parent) {
	return new MiqtVirtualQKeySequenceEdit(*keySequence, parent);
}

void QKeySequenceEdit_virtbase(QKeySequenceEdit* src, QWidget** outptr_QWidget) {
	*outptr_QWidget = static_cast<QWidget*>(src);
}

QMetaObject* QKeySequenceEdit_MetaObject(const QKeySequenceEdit* self) {
	return (QMetaObject*) self->metaObject();
}

void* QKeySequenceEdit_Metacast(QKeySequenceEdit* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QKeySequenceEdit_Tr(const char* s) {
	QString _ret = QKeySequenceEdit::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QKeySequence* QKeySequenceEdit_KeySequence(const QKeySequenceEdit* self) {
	return new QKeySequence(self->keySequence());
}

ptrdiff_t QKeySequenceEdit_MaximumSequenceLength(const QKeySequenceEdit* self) {
	qsizetype _ret = self->maximumSequenceLength();
	return static_cast<ptrdiff_t>(_ret);
}

void QKeySequenceEdit_SetClearButtonEnabled(QKeySequenceEdit* self, bool enable) {
	self->setClearButtonEnabled(enable);
}

bool QKeySequenceEdit_IsClearButtonEnabled(const QKeySequenceEdit* self) {
	return self->isClearButtonEnabled();
}

void QKeySequenceEdit_SetFinishingKeyCombinations(QKeySequenceEdit* self, struct miqt_array /* of QKeyCombination* */  finishingKeyCombinations) {
	QList<QKeyCombination> finishingKeyCombinations_QList;
	finishingKeyCombinations_QList.reserve(finishingKeyCombinations.len);
	QKeyCombination** finishingKeyCombinations_arr = static_cast<QKeyCombination**>(finishingKeyCombinations.data);
	for(size_t i = 0; i < finishingKeyCombinations.len; ++i) {
		finishingKeyCombinations_QList.push_back(*(finishingKeyCombinations_arr[i]));
	}
	self->setFinishingKeyCombinations(finishingKeyCombinations_QList);
}

struct miqt_array /* of QKeyCombination* */  QKeySequenceEdit_FinishingKeyCombinations(const QKeySequenceEdit* self) {
	QList<QKeyCombination> _ret = self->finishingKeyCombinations();
	// Convert QList<> from C++ memory to manually-managed C memory
	QKeyCombination** _arr = static_cast<QKeyCombination**>(malloc(sizeof(QKeyCombination*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QKeyCombination(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QKeySequenceEdit_SetKeySequence(QKeySequenceEdit* self, QKeySequence* keySequence) {
	self->setKeySequence(*keySequence);
}

void QKeySequenceEdit_Clear(QKeySequenceEdit* self) {
	self->clear();
}

void QKeySequenceEdit_SetMaximumSequenceLength(QKeySequenceEdit* self, ptrdiff_t count) {
	self->setMaximumSequenceLength((qsizetype)(count));
}

void QKeySequenceEdit_EditingFinished(QKeySequenceEdit* self) {
	self->editingFinished();
}

void QKeySequenceEdit_connect_EditingFinished(QKeySequenceEdit* self, intptr_t slot) {
	MiqtVirtualQKeySequenceEdit::connect(self, static_cast<void (QKeySequenceEdit::*)()>(&QKeySequenceEdit::editingFinished), self, [=]() {
		miqt_exec_callback_QKeySequenceEdit_EditingFinished(slot);
	});
}

void QKeySequenceEdit_KeySequenceChanged(QKeySequenceEdit* self, QKeySequence* keySequence) {
	self->keySequenceChanged(*keySequence);
}

void QKeySequenceEdit_connect_KeySequenceChanged(QKeySequenceEdit* self, intptr_t slot) {
	MiqtVirtualQKeySequenceEdit::connect(self, static_cast<void (QKeySequenceEdit::*)(const QKeySequence&)>(&QKeySequenceEdit::keySequenceChanged), self, [=](const QKeySequence& keySequence) {
		const QKeySequence& keySequence_ret = keySequence;
		// Cast returned reference into pointer
		QKeySequence* sigval1 = const_cast<QKeySequence*>(&keySequence_ret);
		miqt_exec_callback_QKeySequenceEdit_KeySequenceChanged(slot, sigval1);
	});
}

struct miqt_string QKeySequenceEdit_Tr2(const char* s, const char* c) {
	QString _ret = QKeySequenceEdit::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QKeySequenceEdit_Tr3(const char* s, const char* c, int n) {
	QString _ret = QKeySequenceEdit::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QKeySequenceEdit_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQKeySequenceEdit*>( (QKeySequenceEdit*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QKeySequenceEdit_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQKeySequenceEdit*)(self) )->virtualbase_MetaObject();
}

void QKeySequenceEdit_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQKeySequenceEdit*>( (QKeySequenceEdit*)(self) )->handle__Metacast = slot;
}

void* QKeySequenceEdit_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQKeySequenceEdit*)(self) )->virtualbase_Metacast(param1);
}

void QKeySequenceEdit_Delete(QKeySequenceEdit* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQKeySequenceEdit*>( self );
	} else {
		delete self;
	}
}

