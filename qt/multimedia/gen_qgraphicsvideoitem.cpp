// +build ignore

#include <QGraphicsItem>
#include <QGraphicsObject>
#include <QGraphicsVideoItem>
#include <QMetaObject>
#include <QObject>
#include <QPainter>
#include <QPointF>
#include <QRectF>
#include <QSizeF>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionGraphicsItem>
#include <QTimerEvent>
#include <QVariant>
#include <QVideoSink>
#include <QWidget>
#include <qgraphicsvideoitem.h>
#include "gen_qgraphicsvideoitem.h"
class MiqtVirtualQGraphicsVideoItem : public virtual QGraphicsVideoItem {
public:
MiqtVirtualQGraphicsVideoItem(): QGraphicsVideoItem() {};
MiqtVirtualQGraphicsVideoItem(QGraphicsItem* parent): QGraphicsVideoItem(parent) {};

virtual ~MiqtVirtualQGraphicsVideoItem() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QGraphicsVideoItem::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QGraphicsVideoItem_MetaObject(const_cast<MiqtVirtualQGraphicsVideoItem*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QGraphicsVideoItem::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QGraphicsVideoItem::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QGraphicsVideoItem_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QGraphicsVideoItem::qt_metacast(param1);

	}
};
QGraphicsVideoItem* QGraphicsVideoItem_new() {
return new MiqtVirtualQGraphicsVideoItem();
}
QGraphicsVideoItem* QGraphicsVideoItem_new2(QGraphicsItem* parent) {
return new MiqtVirtualQGraphicsVideoItem(parent);
}
void QGraphicsVideoItem_virtbase(QGraphicsVideoItem* src
, QGraphicsObject** outptr_QGraphicsObject
) {
*outptr_QGraphicsObject = static_cast<QGraphicsObject*>(src);
}
QMetaObject* QGraphicsVideoItem_MetaObject(const QGraphicsVideoItem* self) {
	return (QMetaObject*) self->metaObject();
}
void* QGraphicsVideoItem_Metacast(QGraphicsVideoItem* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QGraphicsVideoItem_Tr(const char* s) {
	QString _ret = QGraphicsVideoItem::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QVideoSink* QGraphicsVideoItem_VideoSink(const QGraphicsVideoItem* self) {
	return self->videoSink();
}
int QGraphicsVideoItem_AspectRatioMode(const QGraphicsVideoItem* self) {
	Qt::AspectRatioMode _ret = self->aspectRatioMode();
	return static_cast<int>(_ret);
}
void QGraphicsVideoItem_SetAspectRatioMode(QGraphicsVideoItem* self, int mode) {
	self->setAspectRatioMode(static_cast<Qt::AspectRatioMode>(mode));
}
QPointF* QGraphicsVideoItem_Offset(const QGraphicsVideoItem* self) {
	return new QPointF(self->offset());
}
void QGraphicsVideoItem_SetOffset(QGraphicsVideoItem* self, QPointF* offset) {
	self->setOffset(*offset);
}
QSizeF* QGraphicsVideoItem_Size(const QGraphicsVideoItem* self) {
	return new QSizeF(self->size());
}
void QGraphicsVideoItem_SetSize(QGraphicsVideoItem* self, QSizeF* size) {
	self->setSize(*size);
}
QSizeF* QGraphicsVideoItem_NativeSize(const QGraphicsVideoItem* self) {
	return new QSizeF(self->nativeSize());
}
QRectF* QGraphicsVideoItem_BoundingRect(const QGraphicsVideoItem* self) {
	return new QRectF(self->boundingRect());
}
void QGraphicsVideoItem_Paint(QGraphicsVideoItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget) {
	self->paint(painter, option, widget);
}
int QGraphicsVideoItem_Type(const QGraphicsVideoItem* self) {
	return self->type();
}
void QGraphicsVideoItem_NativeSizeChanged(QGraphicsVideoItem* self, QSizeF* size) {
	self->nativeSizeChanged(*size);
}
void QGraphicsVideoItem_connect_NativeSizeChanged(QGraphicsVideoItem* self, intptr_t slot) {
	MiqtVirtualQGraphicsVideoItem::connect(self, static_cast<void (QGraphicsVideoItem::*)(const QSizeF&)>(&QGraphicsVideoItem::nativeSizeChanged), self, [=](const QSizeF& size) {
		const QSizeF& size_ret = size;
		// Cast returned reference into pointer
		QSizeF* sigval1 = const_cast<QSizeF*>(&size_ret);
		miqt_exec_callback_QGraphicsVideoItem_NativeSizeChanged(slot, sigval1);
	});
}
struct miqt_string QGraphicsVideoItem_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsVideoItem::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QGraphicsVideoItem_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsVideoItem::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QGraphicsVideoItem_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsVideoItem*>( (QGraphicsVideoItem*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QGraphicsVideoItem_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQGraphicsVideoItem*)(self) )->virtualbase_MetaObject();
}
void QGraphicsVideoItem_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsVideoItem*>( (QGraphicsVideoItem*)(self) )->handle__Metacast = slot;
}
void* QGraphicsVideoItem_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQGraphicsVideoItem*)(self) )->virtualbase_Metacast(param1);
}
void QGraphicsVideoItem_Delete(QGraphicsVideoItem* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsVideoItem*>( self );
	} else {
		delete self;
	}
}
