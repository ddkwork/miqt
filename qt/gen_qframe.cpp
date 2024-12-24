// +build ignore

#include <QEvent>
#include <QFrame>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QRect>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionFrame>
#include <QWidget>
#include <qframe.h>
#include "gen_qframe.h"
class MiqtVirtualQFrame : public virtual QFrame {
public:
MiqtVirtualQFrame(QWidget* parent): QFrame(parent) {};
MiqtVirtualQFrame(): QFrame() {};
MiqtVirtualQFrame(QWidget* parent, Qt::WindowFlags f): QFrame(parent, f) {};

virtual ~MiqtVirtualQFrame() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QFrame::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QFrame_MetaObject(const_cast<MiqtVirtualQFrame*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QFrame::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QFrame::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QFrame_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QFrame::qt_metacast(param1);

	}
};
QFrame* QFrame_new(QWidget* parent) {
return new MiqtVirtualQFrame(parent);
}
QFrame* QFrame_new2() {
return new MiqtVirtualQFrame();
}
QFrame* QFrame_new3(QWidget* parent, int f) {
return new MiqtVirtualQFrame(parent, static_cast<Qt::WindowFlags>(f));
}
void QFrame_virtbase(QFrame* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QFrame_MetaObject(const QFrame* self) {
	return (QMetaObject*) self->metaObject();
}
void* QFrame_Metacast(QFrame* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QFrame_Tr(const char* s) {
	QString _ret = QFrame::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QFrame_FrameStyle(const QFrame* self) {
	return self->frameStyle();
}
void QFrame_SetFrameStyle(QFrame* self, int frameStyle) {
	self->setFrameStyle(static_cast<int>(frameStyle));
}
int QFrame_FrameWidth(const QFrame* self) {
	return self->frameWidth();
}
QSize* QFrame_SizeHint(const QFrame* self) {
	return new QSize(self->sizeHint());
}
Shape QFrame_FrameShape(const QFrame* self) {
	return self->frameShape();
}
void QFrame_SetFrameShape(QFrame* self, Shape frameShape) {
	self->setFrameShape(frameShape);
}
Shadow QFrame_FrameShadow(const QFrame* self) {
	return self->frameShadow();
}
void QFrame_SetFrameShadow(QFrame* self, Shadow frameShadow) {
	self->setFrameShadow(frameShadow);
}
int QFrame_LineWidth(const QFrame* self) {
	return self->lineWidth();
}
void QFrame_SetLineWidth(QFrame* self, int lineWidth) {
	self->setLineWidth(static_cast<int>(lineWidth));
}
int QFrame_MidLineWidth(const QFrame* self) {
	return self->midLineWidth();
}
void QFrame_SetMidLineWidth(QFrame* self, int midLineWidth) {
	self->setMidLineWidth(static_cast<int>(midLineWidth));
}
QRect* QFrame_FrameRect(const QFrame* self) {
	return new QRect(self->frameRect());
}
void QFrame_SetFrameRect(QFrame* self, QRect* frameRect) {
	self->setFrameRect(*frameRect);
}
struct miqt_string QFrame_Tr2(const char* s, const char* c) {
	QString _ret = QFrame::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QFrame_Tr3(const char* s, const char* c, int n) {
	QString _ret = QFrame::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QFrame_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQFrame*>( (QFrame*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QFrame_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQFrame*)(self) )->virtualbase_MetaObject();
}
void QFrame_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQFrame*>( (QFrame*)(self) )->handle__Metacast = slot;
}
void* QFrame_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQFrame*)(self) )->virtualbase_Metacast(param1);
}
void QFrame_Delete(QFrame* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQFrame*>( self );
	} else {
		delete self;
	}
}
