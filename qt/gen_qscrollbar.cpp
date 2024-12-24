// +build ignore

#include <QAbstractSlider>
#include <QContextMenuEvent>
#include <QEvent>
#include <QHideEvent>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QScrollBar>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionSlider>
#include <QWheelEvent>
#include <QWidget>
#include <qscrollbar.h>
#include "gen_qscrollbar.h"

class MiqtVirtualQScrollBar : public virtual QScrollBar {
public:

	MiqtVirtualQScrollBar(QWidget* parent): QScrollBar(parent) {};
	MiqtVirtualQScrollBar(): QScrollBar() {};
	MiqtVirtualQScrollBar(Qt::Orientation param1): QScrollBar(param1) {};
	MiqtVirtualQScrollBar(Qt::Orientation param1, QWidget* parent): QScrollBar(param1, parent) {};

	virtual ~MiqtVirtualQScrollBar() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QScrollBar::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QScrollBar_MetaObject(const_cast<MiqtVirtualQScrollBar*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QScrollBar::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QScrollBar::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QScrollBar_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QScrollBar::qt_metacast(param1);

	}

};

QScrollBar* QScrollBar_new(QWidget* parent) {
	return new MiqtVirtualQScrollBar(parent);
}

QScrollBar* QScrollBar_new2() {
	return new MiqtVirtualQScrollBar();
}

QScrollBar* QScrollBar_new3(int param1) {
	return new MiqtVirtualQScrollBar(static_cast<Qt::Orientation>(param1));
}

QScrollBar* QScrollBar_new4(int param1, QWidget* parent) {
	return new MiqtVirtualQScrollBar(static_cast<Qt::Orientation>(param1), parent);
}

void QScrollBar_virtbase(QScrollBar* src, QAbstractSlider** outptr_QAbstractSlider) {
	*outptr_QAbstractSlider = static_cast<QAbstractSlider*>(src);
}

QMetaObject* QScrollBar_MetaObject(const QScrollBar* self) {
	return (QMetaObject*) self->metaObject();
}

void* QScrollBar_Metacast(QScrollBar* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QScrollBar_Tr(const char* s) {
	QString _ret = QScrollBar::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QSize* QScrollBar_SizeHint(const QScrollBar* self) {
	return new QSize(self->sizeHint());
}

bool QScrollBar_Event(QScrollBar* self, QEvent* event) {
	return self->event(event);
}

struct miqt_string QScrollBar_Tr2(const char* s, const char* c) {
	QString _ret = QScrollBar::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QScrollBar_Tr3(const char* s, const char* c, int n) {
	QString _ret = QScrollBar::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QScrollBar_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQScrollBar*>( (QScrollBar*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QScrollBar_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQScrollBar*)(self) )->virtualbase_MetaObject();
}

void QScrollBar_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQScrollBar*>( (QScrollBar*)(self) )->handle__Metacast = slot;
}

void* QScrollBar_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQScrollBar*)(self) )->virtualbase_Metacast(param1);
}

void QScrollBar_Delete(QScrollBar* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQScrollBar*>( self );
	} else {
		delete self;
	}
}

