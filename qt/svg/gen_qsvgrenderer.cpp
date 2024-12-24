// +build ignore

#include <QByteArray>
#include <QMetaObject>
#include <QObject>
#include <QPainter>
#include <QRect>
#include <QRectF>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QSvgRenderer>
#include <QTransform>
#include <QXmlStreamReader>
#include <qsvgrenderer.h>
#include "gen_qsvgrenderer.h"

class MiqtVirtualQSvgRenderer : public virtual QSvgRenderer {
public:

	MiqtVirtualQSvgRenderer(): QSvgRenderer() {};
	MiqtVirtualQSvgRenderer(const QString& filename): QSvgRenderer(filename) {};
	MiqtVirtualQSvgRenderer(const QByteArray& contents): QSvgRenderer(contents) {};
	MiqtVirtualQSvgRenderer(QXmlStreamReader* contents): QSvgRenderer(contents) {};
	MiqtVirtualQSvgRenderer(QObject* parent): QSvgRenderer(parent) {};
	MiqtVirtualQSvgRenderer(const QString& filename, QObject* parent): QSvgRenderer(filename, parent) {};
	MiqtVirtualQSvgRenderer(const QByteArray& contents, QObject* parent): QSvgRenderer(contents, parent) {};
	MiqtVirtualQSvgRenderer(QXmlStreamReader* contents, QObject* parent): QSvgRenderer(contents, parent) {};

	virtual ~MiqtVirtualQSvgRenderer() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QSvgRenderer::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QSvgRenderer_MetaObject(const_cast<MiqtVirtualQSvgRenderer*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSvgRenderer::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QSvgRenderer::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSvgRenderer_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSvgRenderer::qt_metacast(param1);

	}

};

QSvgRenderer* QSvgRenderer_new() {
	return new MiqtVirtualQSvgRenderer();
}

QSvgRenderer* QSvgRenderer_new2(struct miqt_string filename) {
	QString filename_QString = QString::fromUtf8(filename.data, filename.len);
	return new MiqtVirtualQSvgRenderer(filename_QString);
}

QSvgRenderer* QSvgRenderer_new3(struct miqt_string contents) {
	QByteArray contents_QByteArray(contents.data, contents.len);
	return new MiqtVirtualQSvgRenderer(contents_QByteArray);
}

QSvgRenderer* QSvgRenderer_new4(QXmlStreamReader* contents) {
	return new MiqtVirtualQSvgRenderer(contents);
}

QSvgRenderer* QSvgRenderer_new5(QObject* parent) {
	return new MiqtVirtualQSvgRenderer(parent);
}

QSvgRenderer* QSvgRenderer_new6(struct miqt_string filename, QObject* parent) {
	QString filename_QString = QString::fromUtf8(filename.data, filename.len);
	return new MiqtVirtualQSvgRenderer(filename_QString, parent);
}

QSvgRenderer* QSvgRenderer_new7(struct miqt_string contents, QObject* parent) {
	QByteArray contents_QByteArray(contents.data, contents.len);
	return new MiqtVirtualQSvgRenderer(contents_QByteArray, parent);
}

QSvgRenderer* QSvgRenderer_new8(QXmlStreamReader* contents, QObject* parent) {
	return new MiqtVirtualQSvgRenderer(contents, parent);
}

void QSvgRenderer_virtbase(QSvgRenderer* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QSvgRenderer_MetaObject(const QSvgRenderer* self) {
	return (QMetaObject*) self->metaObject();
}

void* QSvgRenderer_Metacast(QSvgRenderer* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QSvgRenderer_Tr(const char* s) {
	QString _ret = QSvgRenderer::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QSvgRenderer_IsValid(const QSvgRenderer* self) {
	return self->isValid();
}

QSize* QSvgRenderer_DefaultSize(const QSvgRenderer* self) {
	return new QSize(self->defaultSize());
}

QRect* QSvgRenderer_ViewBox(const QSvgRenderer* self) {
	return new QRect(self->viewBox());
}

QRectF* QSvgRenderer_ViewBoxF(const QSvgRenderer* self) {
	return new QRectF(self->viewBoxF());
}

void QSvgRenderer_SetViewBox(QSvgRenderer* self, QRect* viewbox) {
	self->setViewBox(*viewbox);
}

void QSvgRenderer_SetViewBoxWithViewbox(QSvgRenderer* self, QRectF* viewbox) {
	self->setViewBox(*viewbox);
}

int QSvgRenderer_AspectRatioMode(const QSvgRenderer* self) {
	Qt::AspectRatioMode _ret = self->aspectRatioMode();
	return static_cast<int>(_ret);
}

void QSvgRenderer_SetAspectRatioMode(QSvgRenderer* self, int mode) {
	self->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}

int QSvgRenderer_Options(const QSvgRenderer* self) {
	QtSvg::Options _ret = self->options();
	return static_cast<int>(_ret);
}

void QSvgRenderer_SetOptions(QSvgRenderer* self, int flags) {
	self->setOptions(static_cast<QtSvg::Options>(flags));
}

bool QSvgRenderer_Animated(const QSvgRenderer* self) {
	return self->animated();
}

int QSvgRenderer_FramesPerSecond(const QSvgRenderer* self) {
	return self->framesPerSecond();
}

void QSvgRenderer_SetFramesPerSecond(QSvgRenderer* self, int num) {
	self->setFramesPerSecond(static_cast<int>(num));
}

int QSvgRenderer_CurrentFrame(const QSvgRenderer* self) {
	return self->currentFrame();
}

void QSvgRenderer_SetCurrentFrame(QSvgRenderer* self, int currentFrame) {
	self->setCurrentFrame(static_cast<int>(currentFrame));
}

int QSvgRenderer_AnimationDuration(const QSvgRenderer* self) {
	return self->animationDuration();
}

bool QSvgRenderer_IsAnimationEnabled(const QSvgRenderer* self) {
	return self->isAnimationEnabled();
}

void QSvgRenderer_SetAnimationEnabled(QSvgRenderer* self, bool enable) {
	self->setAnimationEnabled(enable);
}

QRectF* QSvgRenderer_BoundsOnElement(const QSvgRenderer* self, struct miqt_string id) {
	QString id_QString = QString::fromUtf8(id.data, id.len);
	return new QRectF(self->boundsOnElement(id_QString));
}

bool QSvgRenderer_ElementExists(const QSvgRenderer* self, struct miqt_string id) {
	QString id_QString = QString::fromUtf8(id.data, id.len);
	return self->elementExists(id_QString);
}

QTransform* QSvgRenderer_TransformForElement(const QSvgRenderer* self, struct miqt_string id) {
	QString id_QString = QString::fromUtf8(id.data, id.len);
	return new QTransform(self->transformForElement(id_QString));
}

void QSvgRenderer_SetDefaultOptions(int flags) {
	QSvgRenderer::setDefaultOptions(static_cast<QtSvg::Options>(flags));
}

bool QSvgRenderer_Load(QSvgRenderer* self, struct miqt_string filename) {
	QString filename_QString = QString::fromUtf8(filename.data, filename.len);
	return self->load(filename_QString);
}

bool QSvgRenderer_LoadWithContents(QSvgRenderer* self, struct miqt_string contents) {
	QByteArray contents_QByteArray(contents.data, contents.len);
	return self->load(contents_QByteArray);
}

bool QSvgRenderer_Load2(QSvgRenderer* self, QXmlStreamReader* contents) {
	return self->load(contents);
}

void QSvgRenderer_Render(QSvgRenderer* self, QPainter* p) {
	self->render(p);
}

void QSvgRenderer_Render2(QSvgRenderer* self, QPainter* p, QRectF* bounds) {
	self->render(p, *bounds);
}

void QSvgRenderer_Render3(QSvgRenderer* self, QPainter* p, struct miqt_string elementId) {
	QString elementId_QString = QString::fromUtf8(elementId.data, elementId.len);
	self->render(p, elementId_QString);
}

void QSvgRenderer_RepaintNeeded(QSvgRenderer* self) {
	self->repaintNeeded();
}

void QSvgRenderer_connect_RepaintNeeded(QSvgRenderer* self, intptr_t slot) {
	MiqtVirtualQSvgRenderer::connect(self, static_cast<void (QSvgRenderer::*)()>(&QSvgRenderer::repaintNeeded), self, [=]() {
		miqt_exec_callback_QSvgRenderer_RepaintNeeded(slot);
	});
}

struct miqt_string QSvgRenderer_Tr2(const char* s, const char* c) {
	QString _ret = QSvgRenderer::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QSvgRenderer_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSvgRenderer::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QSvgRenderer_Render32(QSvgRenderer* self, QPainter* p, struct miqt_string elementId, QRectF* bounds) {
	QString elementId_QString = QString::fromUtf8(elementId.data, elementId.len);
	self->render(p, elementId_QString, *bounds);
}

void QSvgRenderer_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSvgRenderer*>( (QSvgRenderer*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QSvgRenderer_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSvgRenderer*)(self) )->virtualbase_MetaObject();
}

void QSvgRenderer_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSvgRenderer*>( (QSvgRenderer*)(self) )->handle__Metacast = slot;
}

void* QSvgRenderer_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSvgRenderer*)(self) )->virtualbase_Metacast(param1);
}

void QSvgRenderer_Delete(QSvgRenderer* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSvgRenderer*>( self );
	} else {
		delete self;
	}
}

