// +build ignore

#include <QAbstractSlider>
#include <QDial>
#include <QEvent>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QResizeEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionSlider>
#include <QWidget>
#include <qdial.h>
#include "gen_qdial.h"

class MiqtVirtualQDial : public virtual QDial {
public:

	MiqtVirtualQDial(QWidget* parent): QDial(parent) {};
	MiqtVirtualQDial(): QDial() {};

	virtual ~MiqtVirtualQDial() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QDial::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QDial_MetaObject(const_cast<MiqtVirtualQDial*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QDial::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QDial::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QDial_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QDial::qt_metacast(param1);

	}

};

QDial* QDial_new(QWidget* parent) {
	return new MiqtVirtualQDial(parent);
}

QDial* QDial_new2() {
	return new MiqtVirtualQDial();
}

void QDial_virtbase(QDial* src, QAbstractSlider** outptr_QAbstractSlider) {
	*outptr_QAbstractSlider = static_cast<QAbstractSlider*>(src);
}

QMetaObject* QDial_MetaObject(const QDial* self) {
	return (QMetaObject*) self->metaObject();
}

void* QDial_Metacast(QDial* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QDial_Tr(const char* s) {
	QString _ret = QDial::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QDial_Wrapping(const QDial* self) {
	return self->wrapping();
}

int QDial_NotchSize(const QDial* self) {
	return self->notchSize();
}

void QDial_SetNotchTarget(QDial* self, double target) {
	self->setNotchTarget(static_cast<double>(target));
}

double QDial_NotchTarget(const QDial* self) {
	qreal _ret = self->notchTarget();
	return static_cast<double>(_ret);
}

bool QDial_NotchesVisible(const QDial* self) {
	return self->notchesVisible();
}

QSize* QDial_SizeHint(const QDial* self) {
	return new QSize(self->sizeHint());
}

QSize* QDial_MinimumSizeHint(const QDial* self) {
	return new QSize(self->minimumSizeHint());
}

void QDial_SetNotchesVisible(QDial* self, bool visible) {
	self->setNotchesVisible(visible);
}

void QDial_SetWrapping(QDial* self, bool on) {
	self->setWrapping(on);
}

struct miqt_string QDial_Tr2(const char* s, const char* c) {
	QString _ret = QDial::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDial_Tr3(const char* s, const char* c, int n) {
	QString _ret = QDial::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QDial_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDial*>( (QDial*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QDial_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQDial*)(self) )->virtualbase_MetaObject();
}

void QDial_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDial*>( (QDial*)(self) )->handle__Metacast = slot;
}

void* QDial_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQDial*)(self) )->virtualbase_Metacast(param1);
}

void QDial_Delete(QDial* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQDial*>( self );
	} else {
		delete self;
	}
}

