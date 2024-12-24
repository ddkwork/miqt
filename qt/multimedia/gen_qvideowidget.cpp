// +build ignore

#include <QEvent>
#include <QHideEvent>
#include <QMetaObject>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVideoSink>
#include <QVideoWidget>
#include <QWidget>
#include <qvideowidget.h>
#include "gen_qvideowidget.h"
class MiqtVirtualQVideoWidget : public virtual QVideoWidget {
public:
MiqtVirtualQVideoWidget(QWidget* parent): QVideoWidget(parent) {};
MiqtVirtualQVideoWidget(): QVideoWidget() {};

virtual ~MiqtVirtualQVideoWidget() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QVideoWidget::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QVideoWidget_MetaObject(const_cast<MiqtVirtualQVideoWidget*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QVideoWidget::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QVideoWidget::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QVideoWidget_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QVideoWidget::qt_metacast(param1);

	}
};
QVideoWidget* QVideoWidget_new(QWidget* parent) {
return new MiqtVirtualQVideoWidget(parent);
}
QVideoWidget* QVideoWidget_new2() {
return new MiqtVirtualQVideoWidget();
}
void QVideoWidget_virtbase(QVideoWidget* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QVideoWidget_MetaObject(const QVideoWidget* self) {
	return (QMetaObject*) self->metaObject();
}
void* QVideoWidget_Metacast(QVideoWidget* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QVideoWidget_Tr(const char* s) {
	QString _ret = QVideoWidget::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QVideoSink* QVideoWidget_VideoSink(const QVideoWidget* self) {
	return self->videoSink();
}
int QVideoWidget_AspectRatioMode(const QVideoWidget* self) {
	Qt::AspectRatioMode _ret = self->aspectRatioMode();
	return static_cast<int>(_ret);
}
QSize* QVideoWidget_SizeHint(const QVideoWidget* self) {
	return new QSize(self->sizeHint());
}
void QVideoWidget_SetFullScreen(QVideoWidget* self, bool fullScreen) {
	self->setFullScreen(fullScreen);
}
void QVideoWidget_SetAspectRatioMode(QVideoWidget* self, int mode) {
	self->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}
void QVideoWidget_FullScreenChanged(QVideoWidget* self, bool fullScreen) {
	self->fullScreenChanged(fullScreen);
}
void QVideoWidget_connect_FullScreenChanged(QVideoWidget* self, intptr_t slot) {
	MiqtVirtualQVideoWidget::connect(self, static_cast<void (QVideoWidget::*)(bool)>(&QVideoWidget::fullScreenChanged), self, [=](bool fullScreen) {
		bool sigval1 = fullScreen;
		miqt_exec_callback_QVideoWidget_FullScreenChanged(slot, sigval1);
	});
}
void QVideoWidget_AspectRatioModeChanged(QVideoWidget* self, int mode) {
	self->aspectRatioModeChanged(static_cast<Qt::AspectRatioMode>(mode));
}
void QVideoWidget_connect_AspectRatioModeChanged(QVideoWidget* self, intptr_t slot) {
	MiqtVirtualQVideoWidget::connect(self, static_cast<void (QVideoWidget::*)(Qt::AspectRatioMode)>(&QVideoWidget::aspectRatioModeChanged), self, [=](Qt::AspectRatioMode mode) {
		Qt::AspectRatioMode mode_ret = mode;
		int sigval1 = static_cast<int>(mode_ret);
		miqt_exec_callback_QVideoWidget_AspectRatioModeChanged(slot, sigval1);
	});
}
struct miqt_string QVideoWidget_Tr2(const char* s, const char* c) {
	QString _ret = QVideoWidget::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QVideoWidget_Tr3(const char* s, const char* c, int n) {
	QString _ret = QVideoWidget::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QVideoWidget_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQVideoWidget*>( (QVideoWidget*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QVideoWidget_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQVideoWidget*)(self) )->virtualbase_MetaObject();
}
void QVideoWidget_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQVideoWidget*>( (QVideoWidget*)(self) )->handle__Metacast = slot;
}
void* QVideoWidget_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQVideoWidget*)(self) )->virtualbase_Metacast(param1);
}
void QVideoWidget_Delete(QVideoWidget* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQVideoWidget*>( self );
	} else {
		delete self;
	}
}
