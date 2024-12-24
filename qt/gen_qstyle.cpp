// +build ignore

#include <QApplication>
#include <QFontMetrics>
#include <QIcon>
#include <QMetaObject>
#include <QObject>
#include <QPainter>
#include <QPalette>
#include <QPixmap>
#include <QPoint>
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
#include <qstyle.h>
#include "gen_qstyle.h"
class MiqtVirtualQStyle : public virtual QStyle {
public:
MiqtVirtualQStyle(): QStyle() {};

virtual ~MiqtVirtualQStyle() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QStyle::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QStyle_MetaObject(const_cast<MiqtVirtualQStyle*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QStyle::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QStyle::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QStyle_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QStyle::qt_metacast(param1);

	}
};
QStyle* QStyle_new() {
return new MiqtVirtualQStyle();
}
void QStyle_virtbase(QStyle* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QStyle_MetaObject(const QStyle* self) {
	return (QMetaObject*) self->metaObject();
}
void* QStyle_Metacast(QStyle* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QStyle_Tr(const char* s) {
	QString _ret = QStyle::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QStyle_Name(const QStyle* self) {
	QString _ret = self->name();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QStyle_Polish(QStyle* self, QWidget* widget) {
	self->polish(widget);
}
void QStyle_Unpolish(QStyle* self, QWidget* widget) {
	self->unpolish(widget);
}
void QStyle_PolishWithApplication(QStyle* self, QApplication* application) {
	self->polish(application);
}
void QStyle_UnpolishWithApplication(QStyle* self, QApplication* application) {
	self->unpolish(application);
}
void QStyle_PolishWithPalette(QStyle* self, QPalette* palette) {
	self->polish(*palette);
}
QRect* QStyle_ItemTextRect(const QStyle* self, QFontMetrics* fm, QRect* r, int flags, bool enabled, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return new QRect(self->itemTextRect(*fm, *r, static_cast<int>(flags), enabled, text_QString));
}
QRect* QStyle_ItemPixmapRect(const QStyle* self, QRect* r, int flags, QPixmap* pixmap) {
	return new QRect(self->itemPixmapRect(*r, static_cast<int>(flags), *pixmap));
}
void QStyle_DrawItemText(const QStyle* self, QPainter* painter, QRect* rect, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->drawItemText(painter, *rect, static_cast<int>(flags), *pal, enabled, text_QString, static_cast<QPalette::ColorRole>(textRole));
}
void QStyle_DrawItemPixmap(const QStyle* self, QPainter* painter, QRect* rect, int alignment, QPixmap* pixmap) {
	self->drawItemPixmap(painter, *rect, static_cast<int>(alignment), *pixmap);
}
QPalette* QStyle_StandardPalette(const QStyle* self) {
	return new QPalette(self->standardPalette());
}
void QStyle_DrawPrimitive(const QStyle* self, PrimitiveElement pe, QStyleOption* opt, QPainter* p, QWidget* w) {
	self->drawPrimitive(pe, opt, p, w);
}
void QStyle_DrawControl(const QStyle* self, ControlElement element, QStyleOption* opt, QPainter* p, QWidget* w) {
	self->drawControl(element, opt, p, w);
}
QRect* QStyle_SubElementRect(const QStyle* self, SubElement subElement, QStyleOption* option, QWidget* widget) {
	return new QRect(self->subElementRect(subElement, option, widget));
}
void QStyle_DrawComplexControl(const QStyle* self, ComplexControl cc, QStyleOptionComplex* opt, QPainter* p, QWidget* widget) {
	self->drawComplexControl(cc, opt, p, widget);
}
SubControl QStyle_HitTestComplexControl(const QStyle* self, ComplexControl cc, QStyleOptionComplex* opt, QPoint* pt, QWidget* widget) {
	return self->hitTestComplexControl(cc, opt, *pt, widget);
}
QRect* QStyle_SubControlRect(const QStyle* self, ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* widget) {
	return new QRect(self->subControlRect(cc, opt, sc, widget));
}
int QStyle_PixelMetric(const QStyle* self, PixelMetric metric, QStyleOption* option, QWidget* widget) {
	return self->pixelMetric(metric, option, widget);
}
QSize* QStyle_SizeFromContents(const QStyle* self, ContentsType ct, QStyleOption* opt, QSize* contentsSize, QWidget* w) {
	return new QSize(self->sizeFromContents(ct, opt, *contentsSize, w));
}
int QStyle_StyleHint(const QStyle* self, StyleHint stylehint, QStyleOption* opt, QWidget* widget, QStyleHintReturn* returnData) {
	return self->styleHint(stylehint, opt, widget, returnData);
}
QPixmap* QStyle_StandardPixmap(const QStyle* self, StandardPixmap standardPixmap, QStyleOption* opt, QWidget* widget) {
	return new QPixmap(self->standardPixmap(standardPixmap, opt, widget));
}
QIcon* QStyle_StandardIcon(const QStyle* self, StandardPixmap standardIcon, QStyleOption* option, QWidget* widget) {
	return new QIcon(self->standardIcon(standardIcon, option, widget));
}
QPixmap* QStyle_GeneratedIconPixmap(const QStyle* self, int iconMode, QPixmap* pixmap, QStyleOption* opt) {
	return new QPixmap(self->generatedIconPixmap(static_cast<QIcon::Mode>(iconMode), *pixmap, opt));
}
QRect* QStyle_VisualRect(int direction, QRect* boundingRect, QRect* logicalRect) {
	return new QRect(QStyle::visualRect(static_cast<Qt::LayoutDirection>(direction), *boundingRect, *logicalRect));
}
QPoint* QStyle_VisualPos(int direction, QRect* boundingRect, QPoint* logicalPos) {
	return new QPoint(QStyle::visualPos(static_cast<Qt::LayoutDirection>(direction), *boundingRect, *logicalPos));
}
int QStyle_SliderPositionFromValue(int min, int max, int val, int space) {
	return QStyle::sliderPositionFromValue(static_cast<int>(min), static_cast<int>(max), static_cast<int>(val), static_cast<int>(space));
}
int QStyle_SliderValueFromPosition(int min, int max, int pos, int space) {
	return QStyle::sliderValueFromPosition(static_cast<int>(min), static_cast<int>(max), static_cast<int>(pos), static_cast<int>(space));
}
int QStyle_VisualAlignment(int direction, int alignment) {
	Qt::Alignment _ret = QStyle::visualAlignment(static_cast<Qt::LayoutDirection>(direction), static_cast<Qt::Alignment>(alignment));
	return static_cast<int>(_ret);
}
QRect* QStyle_AlignedRect(int direction, int alignment, QSize* size, QRect* rectangle) {
	return new QRect(QStyle::alignedRect(static_cast<Qt::LayoutDirection>(direction), static_cast<Qt::Alignment>(alignment), *size, *rectangle));
}
int QStyle_LayoutSpacing(const QStyle* self, int control1, int control2, int orientation, QStyleOption* option, QWidget* widget) {
	return self->layoutSpacing(static_cast<QSizePolicy::ControlType>(control1), static_cast<QSizePolicy::ControlType>(control2), static_cast<Qt::Orientation>(orientation), option, widget);
}
int QStyle_CombinedLayoutSpacing(const QStyle* self, int controls1, int controls2, int orientation) {
	return self->combinedLayoutSpacing(static_cast<QSizePolicy::ControlTypes>(controls1), static_cast<QSizePolicy::ControlTypes>(controls2), static_cast<Qt::Orientation>(orientation));
}
QStyle* QStyle_Proxy(const QStyle* self) {
	return (QStyle*) self->proxy();
}
struct miqt_string QStyle_Tr2(const char* s, const char* c) {
	QString _ret = QStyle::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QStyle_Tr3(const char* s, const char* c, int n) {
	QString _ret = QStyle::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QStyle_SliderPositionFromValue5(int min, int max, int val, int space, bool upsideDown) {
	return QStyle::sliderPositionFromValue(static_cast<int>(min), static_cast<int>(max), static_cast<int>(val), static_cast<int>(space), upsideDown);
}
int QStyle_SliderValueFromPosition5(int min, int max, int pos, int space, bool upsideDown) {
	return QStyle::sliderValueFromPosition(static_cast<int>(min), static_cast<int>(max), static_cast<int>(pos), static_cast<int>(space), upsideDown);
}
int QStyle_CombinedLayoutSpacing4(const QStyle* self, int controls1, int controls2, int orientation, QStyleOption* option) {
	return self->combinedLayoutSpacing(static_cast<QSizePolicy::ControlTypes>(controls1), static_cast<QSizePolicy::ControlTypes>(controls2), static_cast<Qt::Orientation>(orientation), option);
}
int QStyle_CombinedLayoutSpacing5(const QStyle* self, int controls1, int controls2, int orientation, QStyleOption* option, QWidget* widget) {
	return self->combinedLayoutSpacing(static_cast<QSizePolicy::ControlTypes>(controls1), static_cast<QSizePolicy::ControlTypes>(controls2), static_cast<Qt::Orientation>(orientation), option, widget);
}
void QStyle_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStyle*>( (QStyle*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QStyle_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQStyle*)(self) )->virtualbase_MetaObject();
}
void QStyle_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStyle*>( (QStyle*)(self) )->handle__Metacast = slot;
}
void* QStyle_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQStyle*)(self) )->virtualbase_Metacast(param1);
}
void QStyle_Delete(QStyle* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQStyle*>( self );
	} else {
		delete self;
	}
}
