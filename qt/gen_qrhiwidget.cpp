// +build ignore

#include <QEvent>
#include <QImage>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QResizeEvent>
#include <QRhiWidget>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qrhiwidget.h>
#include "gen_qrhiwidget.h"

class MiqtVirtualQRhiWidget : public virtual QRhiWidget {
public:

	MiqtVirtualQRhiWidget(QWidget* parent): QRhiWidget(parent) {};
	MiqtVirtualQRhiWidget(): QRhiWidget() {};
	MiqtVirtualQRhiWidget(QWidget* parent, Qt::WindowFlags f): QRhiWidget(parent, f) {};

	virtual ~MiqtVirtualQRhiWidget() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QRhiWidget::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QRhiWidget_MetaObject(const_cast<MiqtVirtualQRhiWidget*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QRhiWidget::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QRhiWidget::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QRhiWidget_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QRhiWidget::qt_metacast(param1);

	}

};

QRhiWidget* QRhiWidget_new(QWidget* parent) {
	return new MiqtVirtualQRhiWidget(parent);
}

QRhiWidget* QRhiWidget_new2() {
	return new MiqtVirtualQRhiWidget();
}

QRhiWidget* QRhiWidget_new3(QWidget* parent, int f) {
	return new MiqtVirtualQRhiWidget(parent, static_cast<Qt::WindowFlags>(f));
}

void QRhiWidget_virtbase(QRhiWidget* src, QWidget** outptr_QWidget) {
	*outptr_QWidget = static_cast<QWidget*>(src);
}

QMetaObject* QRhiWidget_MetaObject(const QRhiWidget* self) {
	return (QMetaObject*) self->metaObject();
}

void* QRhiWidget_Metacast(QRhiWidget* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QRhiWidget_Tr(const char* s) {
	QString _ret = QRhiWidget::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

Api QRhiWidget_Api(const QRhiWidget* self) {
	return self->api();
}

void QRhiWidget_SetApi(QRhiWidget* self, Api api) {
	self->setApi(api);
}

bool QRhiWidget_IsDebugLayerEnabled(const QRhiWidget* self) {
	return self->isDebugLayerEnabled();
}

void QRhiWidget_SetDebugLayerEnabled(QRhiWidget* self, bool enable) {
	self->setDebugLayerEnabled(enable);
}

int QRhiWidget_SampleCount(const QRhiWidget* self) {
	return self->sampleCount();
}

void QRhiWidget_SetSampleCount(QRhiWidget* self, int samples) {
	self->setSampleCount(static_cast<int>(samples));
}

TextureFormat QRhiWidget_ColorBufferFormat(const QRhiWidget* self) {
	return self->colorBufferFormat();
}

void QRhiWidget_SetColorBufferFormat(QRhiWidget* self, TextureFormat format) {
	self->setColorBufferFormat(format);
}

QSize* QRhiWidget_FixedColorBufferSize(const QRhiWidget* self) {
	return new QSize(self->fixedColorBufferSize());
}

void QRhiWidget_SetFixedColorBufferSize(QRhiWidget* self, QSize* pixelSize) {
	self->setFixedColorBufferSize(*pixelSize);
}

void QRhiWidget_SetFixedColorBufferSize2(QRhiWidget* self, int w, int h) {
	self->setFixedColorBufferSize(static_cast<int>(w), static_cast<int>(h));
}

bool QRhiWidget_IsMirrorVerticallyEnabled(const QRhiWidget* self) {
	return self->isMirrorVerticallyEnabled();
}

void QRhiWidget_SetMirrorVertically(QRhiWidget* self, bool enabled) {
	self->setMirrorVertically(enabled);
}

QImage* QRhiWidget_GrabFramebuffer(const QRhiWidget* self) {
	return new QImage(self->grabFramebuffer());
}

void QRhiWidget_FrameSubmitted(QRhiWidget* self) {
	self->frameSubmitted();
}

void QRhiWidget_connect_FrameSubmitted(QRhiWidget* self, intptr_t slot) {
	MiqtVirtualQRhiWidget::connect(self, static_cast<void (QRhiWidget::*)()>(&QRhiWidget::frameSubmitted), self, [=]() {
		miqt_exec_callback_QRhiWidget_FrameSubmitted(slot);
	});
}

void QRhiWidget_RenderFailed(QRhiWidget* self) {
	self->renderFailed();
}

void QRhiWidget_connect_RenderFailed(QRhiWidget* self, intptr_t slot) {
	MiqtVirtualQRhiWidget::connect(self, static_cast<void (QRhiWidget::*)()>(&QRhiWidget::renderFailed), self, [=]() {
		miqt_exec_callback_QRhiWidget_RenderFailed(slot);
	});
}

void QRhiWidget_SampleCountChanged(QRhiWidget* self, int samples) {
	self->sampleCountChanged(static_cast<int>(samples));
}

void QRhiWidget_connect_SampleCountChanged(QRhiWidget* self, intptr_t slot) {
	MiqtVirtualQRhiWidget::connect(self, static_cast<void (QRhiWidget::*)(int)>(&QRhiWidget::sampleCountChanged), self, [=](int samples) {
		int sigval1 = samples;
		miqt_exec_callback_QRhiWidget_SampleCountChanged(slot, sigval1);
	});
}

void QRhiWidget_ColorBufferFormatChanged(QRhiWidget* self, TextureFormat format) {
	self->colorBufferFormatChanged(format);
}

void QRhiWidget_connect_ColorBufferFormatChanged(QRhiWidget* self, intptr_t slot) {
	MiqtVirtualQRhiWidget::connect(self, static_cast<void (QRhiWidget::*)(TextureFormat)>(&QRhiWidget::colorBufferFormatChanged), self, [=](TextureFormat format) {
		TextureFormat sigval1 = format;
		miqt_exec_callback_QRhiWidget_ColorBufferFormatChanged(slot, sigval1);
	});
}

void QRhiWidget_FixedColorBufferSizeChanged(QRhiWidget* self, QSize* pixelSize) {
	self->fixedColorBufferSizeChanged(*pixelSize);
}

void QRhiWidget_connect_FixedColorBufferSizeChanged(QRhiWidget* self, intptr_t slot) {
	MiqtVirtualQRhiWidget::connect(self, static_cast<void (QRhiWidget::*)(const QSize&)>(&QRhiWidget::fixedColorBufferSizeChanged), self, [=](const QSize& pixelSize) {
		const QSize& pixelSize_ret = pixelSize;
		// Cast returned reference into pointer
		QSize* sigval1 = const_cast<QSize*>(&pixelSize_ret);
		miqt_exec_callback_QRhiWidget_FixedColorBufferSizeChanged(slot, sigval1);
	});
}

void QRhiWidget_MirrorVerticallyChanged(QRhiWidget* self, bool enabled) {
	self->mirrorVerticallyChanged(enabled);
}

void QRhiWidget_connect_MirrorVerticallyChanged(QRhiWidget* self, intptr_t slot) {
	MiqtVirtualQRhiWidget::connect(self, static_cast<void (QRhiWidget::*)(bool)>(&QRhiWidget::mirrorVerticallyChanged), self, [=](bool enabled) {
		bool sigval1 = enabled;
		miqt_exec_callback_QRhiWidget_MirrorVerticallyChanged(slot, sigval1);
	});
}

struct miqt_string QRhiWidget_Tr2(const char* s, const char* c) {
	QString _ret = QRhiWidget::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QRhiWidget_Tr3(const char* s, const char* c, int n) {
	QString _ret = QRhiWidget::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QRhiWidget_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QRhiWidget_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQRhiWidget*)(self) )->virtualbase_MetaObject();
}

void QRhiWidget_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRhiWidget*>( (QRhiWidget*)(self) )->handle__Metacast = slot;
}

void* QRhiWidget_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQRhiWidget*)(self) )->virtualbase_Metacast(param1);
}

void QRhiWidget_Delete(QRhiWidget* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQRhiWidget*>( self );
	} else {
		delete self;
	}
}

