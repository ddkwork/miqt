// +build ignore

#include <QEvent>
#include <QMetaObject>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QPoint>
#include <QRect>
#include <QResizeEvent>
#include <QRubberBand>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionRubberBand>
#include <QWidget>
#include <qrubberband.h>
#include "gen_qrubberband.h"
class MiqtVirtualQRubberBand : public virtual QRubberBand {
public:
MiqtVirtualQRubberBand(Shape param1): QRubberBand(param1) {};
MiqtVirtualQRubberBand(Shape param1, QWidget* param2): QRubberBand(param1, param2) {};

virtual ~MiqtVirtualQRubberBand() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QRubberBand::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QRubberBand_MetaObject(const_cast<MiqtVirtualQRubberBand*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QRubberBand::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QRubberBand::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QRubberBand_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QRubberBand::qt_metacast(param1);

	}
};
QRubberBand* QRubberBand_new(Shape param1) {
return new MiqtVirtualQRubberBand(param1);
}
QRubberBand* QRubberBand_new2(Shape param1, QWidget* param2) {
return new MiqtVirtualQRubberBand(param1, param2);
}
void QRubberBand_virtbase(QRubberBand* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QRubberBand_MetaObject(const QRubberBand* self) {
	return (QMetaObject*) self->metaObject();
}
void* QRubberBand_Metacast(QRubberBand* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QRubberBand_Tr(const char* s) {
	QString _ret = QRubberBand::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
Shape QRubberBand_Shape(const QRubberBand* self) {
	return self->shape();
}
void QRubberBand_SetGeometry(QRubberBand* self, QRect* r) {
	self->setGeometry(*r);
}
void QRubberBand_SetGeometry2(QRubberBand* self, int x, int y, int w, int h) {
	self->setGeometry(static_cast<int>(x), static_cast<int>(y), static_cast<int>(w), static_cast<int>(h));
}
void QRubberBand_Move(QRubberBand* self, int x, int y) {
	self->move(static_cast<int>(x), static_cast<int>(y));
}
void QRubberBand_MoveWithQPoint(QRubberBand* self, QPoint* p) {
	self->move(*p);
}
void QRubberBand_Resize(QRubberBand* self, int w, int h) {
	self->resize(static_cast<int>(w), static_cast<int>(h));
}
void QRubberBand_ResizeWithQSize(QRubberBand* self, QSize* s) {
	self->resize(*s);
}
struct miqt_string QRubberBand_Tr2(const char* s, const char* c) {
	QString _ret = QRubberBand::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QRubberBand_Tr3(const char* s, const char* c, int n) {
	QString _ret = QRubberBand::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QRubberBand_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRubberBand*>( (QRubberBand*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QRubberBand_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQRubberBand*)(self) )->virtualbase_MetaObject();
}
void QRubberBand_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRubberBand*>( (QRubberBand*)(self) )->handle__Metacast = slot;
}
void* QRubberBand_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQRubberBand*)(self) )->virtualbase_Metacast(param1);
}
void QRubberBand_Delete(QRubberBand* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQRubberBand*>( self );
	} else {
		delete self;
	}
}
