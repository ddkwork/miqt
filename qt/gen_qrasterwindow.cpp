// +build ignore

#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPoint>
#include <QRasterWindow>
#include <QResizeEvent>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QSurface>
#include <QWindow>
#include <qrasterwindow.h>
#include "gen_qrasterwindow.h"
class MiqtVirtualQRasterWindow : public virtual QRasterWindow {
public:
MiqtVirtualQRasterWindow(): QRasterWindow() {};
MiqtVirtualQRasterWindow(QWindow* parent): QRasterWindow(parent) {};

virtual ~MiqtVirtualQRasterWindow() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QRasterWindow::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QRasterWindow_MetaObject(const_cast<MiqtVirtualQRasterWindow*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QRasterWindow::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QRasterWindow::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QRasterWindow_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QRasterWindow::qt_metacast(param1);

	}
};
QRasterWindow* QRasterWindow_new() {
return new MiqtVirtualQRasterWindow();
}
QRasterWindow* QRasterWindow_new2(QWindow* parent) {
return new MiqtVirtualQRasterWindow(parent);
}
void QRasterWindow_virtbase(QRasterWindow* src
, QPaintDeviceWindow** outptr_QPaintDeviceWindow
) {
*outptr_QPaintDeviceWindow = static_cast<QPaintDeviceWindow*>(src);
}
QMetaObject* QRasterWindow_MetaObject(const QRasterWindow* self) {
	return (QMetaObject*) self->metaObject();
}
void* QRasterWindow_Metacast(QRasterWindow* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QRasterWindow_Tr(const char* s) {
	QString _ret = QRasterWindow::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QRasterWindow_Tr2(const char* s, const char* c) {
	QString _ret = QRasterWindow::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QRasterWindow_Tr3(const char* s, const char* c, int n) {
	QString _ret = QRasterWindow::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QRasterWindow_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRasterWindow*>( (QRasterWindow*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QRasterWindow_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQRasterWindow*)(self) )->virtualbase_MetaObject();
}
void QRasterWindow_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRasterWindow*>( (QRasterWindow*)(self) )->handle__Metacast = slot;
}
void* QRasterWindow_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQRasterWindow*)(self) )->virtualbase_Metacast(param1);
}
void QRasterWindow_Delete(QRasterWindow* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQRasterWindow*>( self );
	} else {
		delete self;
	}
}
