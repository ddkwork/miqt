// +build ignore

#include <QApplication>
#include <QCommonStyle>
#include <QEvent>
#include <QFontMetrics>
#include <QIcon>
#include <QMetaObject>
#include <QObject>
#include <QPainter>
#include <QPalette>
#include <QPixmap>
#include <QPoint>
#include <QProxyStyle>
#include <QRect>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyle>
#include <QStyleHintReturn>
#include <QStyleOption>
#include <QStyleOptionComplex>
#include <QWidget>
#include <qproxystyle.h>
#include "gen_qproxystyle.h"
class MiqtVirtualQProxyStyle : public virtual QProxyStyle {
public:
MiqtVirtualQProxyStyle(): QProxyStyle() {};
MiqtVirtualQProxyStyle(const QString& key): QProxyStyle(key) {};
MiqtVirtualQProxyStyle(QStyle* style): QProxyStyle(style) {};

virtual ~MiqtVirtualQProxyStyle() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QProxyStyle::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QProxyStyle_MetaObject(const_cast<MiqtVirtualQProxyStyle*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QProxyStyle::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QProxyStyle::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QProxyStyle_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QProxyStyle::qt_metacast(param1);

	}
};
QProxyStyle* QProxyStyle_new() {
return new MiqtVirtualQProxyStyle();
}
QProxyStyle* QProxyStyle_new2(struct miqt_string key) {
QString key_QString = QString::fromUtf8(key.data, key.len);
	return new MiqtVirtualQProxyStyle(key_QString);
}
QProxyStyle* QProxyStyle_new3(QStyle* style) {
return new MiqtVirtualQProxyStyle(style);
}
void QProxyStyle_virtbase(QProxyStyle* src
, QCommonStyle** outptr_QCommonStyle
) {
*outptr_QCommonStyle = static_cast<QCommonStyle*>(src);
}
QMetaObject* QProxyStyle_MetaObject(const QProxyStyle* self) {
	return (QMetaObject*) self->metaObject();
}
void* QProxyStyle_Metacast(QProxyStyle* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QProxyStyle_Tr(const char* s) {
	QString _ret = QProxyStyle::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QStyle* QProxyStyle_BaseStyle(const QProxyStyle* self) {
	return self->baseStyle();
}
void QProxyStyle_SetBaseStyle(QProxyStyle* self, QStyle* style) {
	self->setBaseStyle(style);
}
void QProxyStyle_DrawPrimitive(const QProxyStyle* self, PrimitiveElement element, QStyleOption* option, QPainter* painter, QWidget* widget) {
	self->drawPrimitive(element, option, painter, widget);
}
void QProxyStyle_DrawControl(const QProxyStyle* self, ControlElement element, QStyleOption* option, QPainter* painter, QWidget* widget) {
	self->drawControl(element, option, painter, widget);
}
void QProxyStyle_DrawComplexControl(const QProxyStyle* self, ComplexControl control, QStyleOptionComplex* option, QPainter* painter, QWidget* widget) {
	self->drawComplexControl(control, option, painter, widget);
}
void QProxyStyle_DrawItemText(const QProxyStyle* self, QPainter* painter, QRect* rect, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->drawItemText(painter, *rect, static_cast<int>(flags), *pal, enabled, text_QString, static_cast<QPalette::ColorRole>(textRole));
}
void QProxyStyle_DrawItemPixmap(const QProxyStyle* self, QPainter* painter, QRect* rect, int alignment, QPixmap* pixmap) {
	self->drawItemPixmap(painter, *rect, static_cast<int>(alignment), *pixmap);
}
QSize* QProxyStyle_SizeFromContents(const QProxyStyle* self, ContentsType typeVal, QStyleOption* option, QSize* size, QWidget* widget) {
	return new QSize(self->sizeFromContents(typeVal, option, *size, widget));
}
QRect* QProxyStyle_SubElementRect(const QProxyStyle* self, SubElement element, QStyleOption* option, QWidget* widget) {
	return new QRect(self->subElementRect(element, option, widget));
}
QRect* QProxyStyle_SubControlRect(const QProxyStyle* self, ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* widget) {
	return new QRect(self->subControlRect(cc, opt, sc, widget));
}
QRect* QProxyStyle_ItemTextRect(const QProxyStyle* self, QFontMetrics* fm, QRect* r, int flags, bool enabled, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return new QRect(self->itemTextRect(*fm, *r, static_cast<int>(flags), enabled, text_QString));
}
QRect* QProxyStyle_ItemPixmapRect(const QProxyStyle* self, QRect* r, int flags, QPixmap* pixmap) {
	return new QRect(self->itemPixmapRect(*r, static_cast<int>(flags), *pixmap));
}
SubControl QProxyStyle_HitTestComplexControl(const QProxyStyle* self, ComplexControl control, QStyleOptionComplex* option, QPoint* pos, QWidget* widget) {
	return self->hitTestComplexControl(control, option, *pos, widget);
}
int QProxyStyle_StyleHint(const QProxyStyle* self, StyleHint hint, QStyleOption* option, QWidget* widget, QStyleHintReturn* returnData) {
	return self->styleHint(hint, option, widget, returnData);
}
int QProxyStyle_PixelMetric(const QProxyStyle* self, PixelMetric metric, QStyleOption* option, QWidget* widget) {
	return self->pixelMetric(metric, option, widget);
}
int QProxyStyle_LayoutSpacing(const QProxyStyle* self, int control1, int control2, int orientation, QStyleOption* option, QWidget* widget) {
	return self->layoutSpacing(static_cast<QSizePolicy::ControlType>(control1), static_cast<QSizePolicy::ControlType>(control2), static_cast<Qt::Orientation>(orientation), option, widget);
}
QIcon* QProxyStyle_StandardIcon(const QProxyStyle* self, StandardPixmap standardIcon, QStyleOption* option, QWidget* widget) {
	return new QIcon(self->standardIcon(standardIcon, option, widget));
}
QPixmap* QProxyStyle_StandardPixmap(const QProxyStyle* self, StandardPixmap standardPixmap, QStyleOption* opt, QWidget* widget) {
	return new QPixmap(self->standardPixmap(standardPixmap, opt, widget));
}
QPixmap* QProxyStyle_GeneratedIconPixmap(const QProxyStyle* self, int iconMode, QPixmap* pixmap, QStyleOption* opt) {
	return new QPixmap(self->generatedIconPixmap(static_cast<QIcon::Mode>(iconMode), *pixmap, opt));
}
QPalette* QProxyStyle_StandardPalette(const QProxyStyle* self) {
	return new QPalette(self->standardPalette());
}
void QProxyStyle_Polish(QProxyStyle* self, QWidget* widget) {
	self->polish(widget);
}
void QProxyStyle_PolishWithPal(QProxyStyle* self, QPalette* pal) {
	self->polish(*pal);
}
void QProxyStyle_PolishWithApp(QProxyStyle* self, QApplication* app) {
	self->polish(app);
}
void QProxyStyle_Unpolish(QProxyStyle* self, QWidget* widget) {
	self->unpolish(widget);
}
void QProxyStyle_UnpolishWithApp(QProxyStyle* self, QApplication* app) {
	self->unpolish(app);
}
struct miqt_string QProxyStyle_Tr2(const char* s, const char* c) {
	QString _ret = QProxyStyle::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QProxyStyle_Tr3(const char* s, const char* c, int n) {
	QString _ret = QProxyStyle::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QProxyStyle_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQProxyStyle*>( (QProxyStyle*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QProxyStyle_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQProxyStyle*)(self) )->virtualbase_MetaObject();
}
void QProxyStyle_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQProxyStyle*>( (QProxyStyle*)(self) )->handle__Metacast = slot;
}
void* QProxyStyle_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQProxyStyle*)(self) )->virtualbase_Metacast(param1);
}
void QProxyStyle_Delete(QProxyStyle* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQProxyStyle*>( self );
	} else {
		delete self;
	}
}
