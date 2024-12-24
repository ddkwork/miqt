// +build ignore

#include <QGraphicsItem>
#include <QGraphicsObject>
#include <QGraphicsSvgItem>
#include <QMetaObject>
#include <QObject>
#include <QPainter>
#include <QRectF>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionGraphicsItem>
#include <QSvgRenderer>
#include <QWidget>
#include <qgraphicssvgitem.h>
#include "gen_qgraphicssvgitem.h"
class MiqtVirtualQGraphicsSvgItem : public virtual QGraphicsSvgItem {
public:
MiqtVirtualQGraphicsSvgItem(): QGraphicsSvgItem() {};
MiqtVirtualQGraphicsSvgItem(const QString& fileName): QGraphicsSvgItem(fileName) {};
MiqtVirtualQGraphicsSvgItem(QGraphicsItem* parentItem): QGraphicsSvgItem(parentItem) {};
MiqtVirtualQGraphicsSvgItem(const QString& fileName, QGraphicsItem* parentItem): QGraphicsSvgItem(fileName, parentItem) {};

virtual ~MiqtVirtualQGraphicsSvgItem() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QGraphicsSvgItem::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QGraphicsSvgItem_MetaObject(const_cast<MiqtVirtualQGraphicsSvgItem*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QGraphicsSvgItem::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QGraphicsSvgItem::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QGraphicsSvgItem_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QGraphicsSvgItem::qt_metacast(param1);

	}
};
QGraphicsSvgItem* QGraphicsSvgItem_new() {
return new MiqtVirtualQGraphicsSvgItem();
}
QGraphicsSvgItem* QGraphicsSvgItem_new2(struct miqt_string fileName) {
QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return new MiqtVirtualQGraphicsSvgItem(fileName_QString);
}
QGraphicsSvgItem* QGraphicsSvgItem_new3(QGraphicsItem* parentItem) {
return new MiqtVirtualQGraphicsSvgItem(parentItem);
}
QGraphicsSvgItem* QGraphicsSvgItem_new4(struct miqt_string fileName, QGraphicsItem* parentItem) {
QString fileName_QString = QString::fromUtf8(fileName.data, fileName.len);
	return new MiqtVirtualQGraphicsSvgItem(fileName_QString, parentItem);
}
void QGraphicsSvgItem_virtbase(QGraphicsSvgItem* src
, QGraphicsObject** outptr_QGraphicsObject
) {
*outptr_QGraphicsObject = static_cast<QGraphicsObject*>(src);
}
QMetaObject* QGraphicsSvgItem_MetaObject(const QGraphicsSvgItem* self) {
	return (QMetaObject*) self->metaObject();
}
void* QGraphicsSvgItem_Metacast(QGraphicsSvgItem* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QGraphicsSvgItem_Tr(const char* s) {
	QString _ret = QGraphicsSvgItem::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QGraphicsSvgItem_SetSharedRenderer(QGraphicsSvgItem* self, QSvgRenderer* renderer) {
	self->setSharedRenderer(renderer);
}
QSvgRenderer* QGraphicsSvgItem_Renderer(const QGraphicsSvgItem* self) {
	return self->renderer();
}
void QGraphicsSvgItem_SetElementId(QGraphicsSvgItem* self, struct miqt_string id) {
	QString id_QString = QString::fromUtf8(id.data, id.len);
	self->setElementId(id_QString);
}
struct miqt_string QGraphicsSvgItem_ElementId(const QGraphicsSvgItem* self) {
	QString _ret = self->elementId();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QGraphicsSvgItem_SetCachingEnabled(QGraphicsSvgItem* self, bool cachingEnabled) {
	self->setCachingEnabled(cachingEnabled);
}
bool QGraphicsSvgItem_IsCachingEnabled(const QGraphicsSvgItem* self) {
	return self->isCachingEnabled();
}
void QGraphicsSvgItem_SetMaximumCacheSize(QGraphicsSvgItem* self, QSize* size) {
	self->setMaximumCacheSize(*size);
}
QSize* QGraphicsSvgItem_MaximumCacheSize(const QGraphicsSvgItem* self) {
	return new QSize(self->maximumCacheSize());
}
QRectF* QGraphicsSvgItem_BoundingRect(const QGraphicsSvgItem* self) {
	return new QRectF(self->boundingRect());
}
void QGraphicsSvgItem_Paint(QGraphicsSvgItem* self, QPainter* painter, QStyleOptionGraphicsItem* option, QWidget* widget) {
	self->paint(painter, option, widget);
}
int QGraphicsSvgItem_Type(const QGraphicsSvgItem* self) {
	return self->type();
}
struct miqt_string QGraphicsSvgItem_Tr2(const char* s, const char* c) {
	QString _ret = QGraphicsSvgItem::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QGraphicsSvgItem_Tr3(const char* s, const char* c, int n) {
	QString _ret = QGraphicsSvgItem::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QGraphicsSvgItem_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsSvgItem*>( (QGraphicsSvgItem*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QGraphicsSvgItem_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQGraphicsSvgItem*)(self) )->virtualbase_MetaObject();
}
void QGraphicsSvgItem_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQGraphicsSvgItem*>( (QGraphicsSvgItem*)(self) )->handle__Metacast = slot;
}
void* QGraphicsSvgItem_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQGraphicsSvgItem*)(self) )->virtualbase_Metacast(param1);
}
void QGraphicsSvgItem_Delete(QGraphicsSvgItem* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQGraphicsSvgItem*>( self );
	} else {
		delete self;
	}
}
