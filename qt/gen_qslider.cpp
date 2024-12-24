// +build ignore

#include <QAbstractSlider>
#include <QEvent>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QSize>
#include <QSlider>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionSlider>
#include <QWidget>
#include <qslider.h>
#include "gen_qslider.h"
class MiqtVirtualQSlider : public virtual QSlider {
public:
MiqtVirtualQSlider(QWidget* parent): QSlider(parent) {};
MiqtVirtualQSlider(): QSlider() {};
MiqtVirtualQSlider(Qt::Orientation orientation): QSlider(orientation) {};
MiqtVirtualQSlider(Qt::Orientation orientation, QWidget* parent): QSlider(orientation, parent) {};

virtual ~MiqtVirtualQSlider() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QSlider::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QSlider_MetaObject(const_cast<MiqtVirtualQSlider*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSlider::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QSlider::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSlider_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSlider::qt_metacast(param1);

	}
};
QSlider* QSlider_new(QWidget* parent) {
return new MiqtVirtualQSlider(parent);
}
QSlider* QSlider_new2() {
return new MiqtVirtualQSlider();
}
QSlider* QSlider_new3(int orientation) {
return new MiqtVirtualQSlider(static_cast<Qt::Orientation>(orientation));
}
QSlider* QSlider_new4(int orientation, QWidget* parent) {
return new MiqtVirtualQSlider(static_cast<Qt::Orientation>(orientation), parent);
}
void QSlider_virtbase(QSlider* src
, QAbstractSlider** outptr_QAbstractSlider
) {
*outptr_QAbstractSlider = static_cast<QAbstractSlider*>(src);
}
QMetaObject* QSlider_MetaObject(const QSlider* self) {
	return (QMetaObject*) self->metaObject();
}
void* QSlider_Metacast(QSlider* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QSlider_Tr(const char* s) {
	QString _ret = QSlider::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QSize* QSlider_SizeHint(const QSlider* self) {
	return new QSize(self->sizeHint());
}
QSize* QSlider_MinimumSizeHint(const QSlider* self) {
	return new QSize(self->minimumSizeHint());
}
void QSlider_SetTickPosition(QSlider* self, TickPosition position) {
	self->setTickPosition(position);
}
TickPosition QSlider_TickPosition(const QSlider* self) {
	return self->tickPosition();
}
void QSlider_SetTickInterval(QSlider* self, int ti) {
	self->setTickInterval(static_cast<int>(ti));
}
int QSlider_TickInterval(const QSlider* self) {
	return self->tickInterval();
}
bool QSlider_Event(QSlider* self, QEvent* event) {
	return self->event(event);
}
struct miqt_string QSlider_Tr2(const char* s, const char* c) {
	QString _ret = QSlider::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QSlider_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSlider::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSlider_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSlider*>( (QSlider*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QSlider_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSlider*)(self) )->virtualbase_MetaObject();
}
void QSlider_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSlider*>( (QSlider*)(self) )->handle__Metacast = slot;
}
void* QSlider_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSlider*)(self) )->virtualbase_Metacast(param1);
}
void QSlider_Delete(QSlider* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSlider*>( self );
	} else {
		delete self;
	}
}
