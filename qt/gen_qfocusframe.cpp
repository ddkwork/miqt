// +build ignore

#include <QEvent>
#include <QFocusFrame>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOption>
#include <QWidget>
#include <qfocusframe.h>
#include "gen_qfocusframe.h"
class MiqtVirtualQFocusFrame : public virtual QFocusFrame {
public:
MiqtVirtualQFocusFrame(QWidget* parent): QFocusFrame(parent) {};
MiqtVirtualQFocusFrame(): QFocusFrame() {};

virtual ~MiqtVirtualQFocusFrame() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QFocusFrame::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QFocusFrame_MetaObject(const_cast<MiqtVirtualQFocusFrame*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QFocusFrame::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QFocusFrame::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QFocusFrame_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QFocusFrame::qt_metacast(param1);

	}
};
QFocusFrame* QFocusFrame_new(QWidget* parent) {
return new MiqtVirtualQFocusFrame(parent);
}
QFocusFrame* QFocusFrame_new2() {
return new MiqtVirtualQFocusFrame();
}
void QFocusFrame_virtbase(QFocusFrame* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QFocusFrame_MetaObject(const QFocusFrame* self) {
	return (QMetaObject*) self->metaObject();
}
void* QFocusFrame_Metacast(QFocusFrame* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QFocusFrame_Tr(const char* s) {
	QString _ret = QFocusFrame::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QFocusFrame_SetWidget(QFocusFrame* self, QWidget* widget) {
	self->setWidget(widget);
}
QWidget* QFocusFrame_Widget(const QFocusFrame* self) {
	return self->widget();
}
struct miqt_string QFocusFrame_Tr2(const char* s, const char* c) {
	QString _ret = QFocusFrame::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QFocusFrame_Tr3(const char* s, const char* c, int n) {
	QString _ret = QFocusFrame::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QFocusFrame_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQFocusFrame*>( (QFocusFrame*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QFocusFrame_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQFocusFrame*)(self) )->virtualbase_MetaObject();
}
void QFocusFrame_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQFocusFrame*>( (QFocusFrame*)(self) )->handle__Metacast = slot;
}
void* QFocusFrame_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQFocusFrame*)(self) )->virtualbase_Metacast(param1);
}
void QFocusFrame_Delete(QFocusFrame* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQFocusFrame*>( self );
	} else {
		delete self;
	}
}
