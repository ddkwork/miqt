// +build ignore

#include <QApplication>
#include <QCommonStyle>
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
#include <qcommonstyle.h>
#include "gen_qcommonstyle.h"

class MiqtVirtualQCommonStyle : public virtual QCommonStyle {
public:

	MiqtVirtualQCommonStyle(): QCommonStyle() {};

	virtual ~MiqtVirtualQCommonStyle() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QCommonStyle::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QCommonStyle_MetaObject(const_cast<MiqtVirtualQCommonStyle*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QCommonStyle::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QCommonStyle::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QCommonStyle_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QCommonStyle::qt_metacast(param1);

	}

};

QCommonStyle* QCommonStyle_new() {
	return new MiqtVirtualQCommonStyle();
}

void QCommonStyle_virtbase(QCommonStyle* src, QStyle** outptr_QStyle) {
	*outptr_QStyle = static_cast<QStyle*>(src);
}

QMetaObject* QCommonStyle_MetaObject(const QCommonStyle* self) {
	return (QMetaObject*) self->metaObject();
}

void* QCommonStyle_Metacast(QCommonStyle* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QCommonStyle_Tr(const char* s) {
	QString _ret = QCommonStyle::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QCommonStyle_DrawPrimitive(const QCommonStyle* self, PrimitiveElement pe, QStyleOption* opt, QPainter* p, QWidget* w) {
	self->drawPrimitive(pe, opt, p, w);
}

void QCommonStyle_DrawControl(const QCommonStyle* self, ControlElement element, QStyleOption* opt, QPainter* p, QWidget* w) {
	self->drawControl(element, opt, p, w);
}

QRect* QCommonStyle_SubElementRect(const QCommonStyle* self, SubElement r, QStyleOption* opt, QWidget* widget) {
	return new QRect(self->subElementRect(r, opt, widget));
}

void QCommonStyle_DrawComplexControl(const QCommonStyle* self, ComplexControl cc, QStyleOptionComplex* opt, QPainter* p, QWidget* w) {
	self->drawComplexControl(cc, opt, p, w);
}

SubControl QCommonStyle_HitTestComplexControl(const QCommonStyle* self, ComplexControl cc, QStyleOptionComplex* opt, QPoint* pt, QWidget* w) {
	return self->hitTestComplexControl(cc, opt, *pt, w);
}

QRect* QCommonStyle_SubControlRect(const QCommonStyle* self, ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* w) {
	return new QRect(self->subControlRect(cc, opt, sc, w));
}

QSize* QCommonStyle_SizeFromContents(const QCommonStyle* self, ContentsType ct, QStyleOption* opt, QSize* contentsSize, QWidget* widget) {
	return new QSize(self->sizeFromContents(ct, opt, *contentsSize, widget));
}

int QCommonStyle_PixelMetric(const QCommonStyle* self, PixelMetric m, QStyleOption* opt, QWidget* widget) {
	return self->pixelMetric(m, opt, widget);
}

int QCommonStyle_StyleHint(const QCommonStyle* self, StyleHint sh, QStyleOption* opt, QWidget* w, QStyleHintReturn* shret) {
	return self->styleHint(sh, opt, w, shret);
}

QIcon* QCommonStyle_StandardIcon(const QCommonStyle* self, StandardPixmap standardIcon, QStyleOption* opt, QWidget* widget) {
	return new QIcon(self->standardIcon(standardIcon, opt, widget));
}

QPixmap* QCommonStyle_StandardPixmap(const QCommonStyle* self, StandardPixmap sp, QStyleOption* opt, QWidget* widget) {
	return new QPixmap(self->standardPixmap(sp, opt, widget));
}

QPixmap* QCommonStyle_GeneratedIconPixmap(const QCommonStyle* self, int iconMode, QPixmap* pixmap, QStyleOption* opt) {
	return new QPixmap(self->generatedIconPixmap(static_cast<QIcon::Mode>(iconMode), *pixmap, opt));
}

int QCommonStyle_LayoutSpacing(const QCommonStyle* self, int control1, int control2, int orientation, QStyleOption* option, QWidget* widget) {
	return self->layoutSpacing(static_cast<QSizePolicy::ControlType>(control1), static_cast<QSizePolicy::ControlType>(control2), static_cast<Qt::Orientation>(orientation), option, widget);
}

void QCommonStyle_Polish(QCommonStyle* self, QPalette* param1) {
	self->polish(*param1);
}

void QCommonStyle_PolishWithApp(QCommonStyle* self, QApplication* app) {
	self->polish(app);
}

void QCommonStyle_PolishWithWidget(QCommonStyle* self, QWidget* widget) {
	self->polish(widget);
}

void QCommonStyle_Unpolish(QCommonStyle* self, QWidget* widget) {
	self->unpolish(widget);
}

void QCommonStyle_UnpolishWithApplication(QCommonStyle* self, QApplication* application) {
	self->unpolish(application);
}

struct miqt_string QCommonStyle_Tr2(const char* s, const char* c) {
	QString _ret = QCommonStyle::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QCommonStyle_Tr3(const char* s, const char* c, int n) {
	QString _ret = QCommonStyle::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QCommonStyle_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QCommonStyle_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_MetaObject();
}

void QCommonStyle_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__Metacast = slot;
}

void* QCommonStyle_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQCommonStyle*)(self) )->virtualbase_Metacast(param1);
}

void QCommonStyle_Delete(QCommonStyle* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQCommonStyle*>( self );
	} else {
		delete self;
	}
}

