// +build ignore

#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QScreen>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QSurface>
#include <QSurfaceFormat>
#include <qoffscreensurface.h>
#include "gen_qoffscreensurface.h"
class MiqtVirtualQOffscreenSurface : public virtual QOffscreenSurface {
public:
MiqtVirtualQOffscreenSurface(): QOffscreenSurface() {};
MiqtVirtualQOffscreenSurface(QScreen* screen): QOffscreenSurface(screen) {};
MiqtVirtualQOffscreenSurface(QScreen* screen, QObject* parent): QOffscreenSurface(screen, parent) {};

virtual ~MiqtVirtualQOffscreenSurface() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QOffscreenSurface::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QOffscreenSurface_MetaObject(const_cast<MiqtVirtualQOffscreenSurface*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QOffscreenSurface::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QOffscreenSurface::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QOffscreenSurface_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QOffscreenSurface::qt_metacast(param1);

	}
};
QOffscreenSurface* QOffscreenSurface_new() {
return new MiqtVirtualQOffscreenSurface();
}
QOffscreenSurface* QOffscreenSurface_new2(QScreen* screen) {
return new MiqtVirtualQOffscreenSurface(screen);
}
QOffscreenSurface* QOffscreenSurface_new3(QScreen* screen, QObject* parent) {
return new MiqtVirtualQOffscreenSurface(screen, parent);
}
void QOffscreenSurface_virtbase(QOffscreenSurface* src
, QObject** outptr_QObject
, QSurface** outptr_QSurface
) {
*outptr_QObject = static_cast<QObject*>(src);
*outptr_QSurface = static_cast<QSurface*>(src);
}
QMetaObject* QOffscreenSurface_MetaObject(const QOffscreenSurface* self) {
	return (QMetaObject*) self->metaObject();
}
void* QOffscreenSurface_Metacast(QOffscreenSurface* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QOffscreenSurface_Tr(const char* s) {
	QString _ret = QOffscreenSurface::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
SurfaceType QOffscreenSurface_SurfaceType(const QOffscreenSurface* self) {
	return self->surfaceType();
}
void QOffscreenSurface_Create(QOffscreenSurface* self) {
	self->create();
}
void QOffscreenSurface_Destroy(QOffscreenSurface* self) {
	self->destroy();
}
bool QOffscreenSurface_IsValid(const QOffscreenSurface* self) {
	return self->isValid();
}
void QOffscreenSurface_SetFormat(QOffscreenSurface* self, QSurfaceFormat* format) {
	self->setFormat(*format);
}
QSurfaceFormat* QOffscreenSurface_Format(const QOffscreenSurface* self) {
	return new QSurfaceFormat(self->format());
}
QSurfaceFormat* QOffscreenSurface_RequestedFormat(const QOffscreenSurface* self) {
	return new QSurfaceFormat(self->requestedFormat());
}
QSize* QOffscreenSurface_Size(const QOffscreenSurface* self) {
	return new QSize(self->size());
}
QScreen* QOffscreenSurface_Screen(const QOffscreenSurface* self) {
	return self->screen();
}
void QOffscreenSurface_SetScreen(QOffscreenSurface* self, QScreen* screen) {
	self->setScreen(screen);
}
void QOffscreenSurface_ScreenChanged(QOffscreenSurface* self, QScreen* screen) {
	self->screenChanged(screen);
}
void QOffscreenSurface_connect_ScreenChanged(QOffscreenSurface* self, intptr_t slot) {
	MiqtVirtualQOffscreenSurface::connect(self, static_cast<void (QOffscreenSurface::*)(QScreen*)>(&QOffscreenSurface::screenChanged), self, [=](QScreen* screen) {
		QScreen* sigval1 = screen;
		miqt_exec_callback_QOffscreenSurface_ScreenChanged(slot, sigval1);
	});
}
struct miqt_string QOffscreenSurface_Tr2(const char* s, const char* c) {
	QString _ret = QOffscreenSurface::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QOffscreenSurface_Tr3(const char* s, const char* c, int n) {
	QString _ret = QOffscreenSurface::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QOffscreenSurface_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQOffscreenSurface*>( (QOffscreenSurface*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QOffscreenSurface_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQOffscreenSurface*)(self) )->virtualbase_MetaObject();
}
void QOffscreenSurface_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQOffscreenSurface*>( (QOffscreenSurface*)(self) )->handle__Metacast = slot;
}
void* QOffscreenSurface_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQOffscreenSurface*)(self) )->virtualbase_Metacast(param1);
}
void QOffscreenSurface_Delete(QOffscreenSurface* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQOffscreenSurface*>( self );
	} else {
		delete self;
	}
}
