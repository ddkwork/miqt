// +build ignore

#include <QApplication>
#include <QCommonStyle>
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
#include <qcommonstyle.h>
#include "gen_qcommonstyle.h"

class MiqtVirtualQCommonStyle : public virtual QCommonStyle {
public:

	MiqtVirtualQCommonStyle(): QCommonStyle() {};

	virtual ~MiqtVirtualQCommonStyle() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DrawPrimitive = 0;

	// Subclass to allow providing a Go implementation
	virtual void drawPrimitive(PrimitiveElement pe, const QStyleOption* opt, QPainter* p, const QWidget* w) const override {
		if (handle__DrawPrimitive == 0) {
			QCommonStyle::drawPrimitive(pe, opt, p, w);
			return;
		}
		
		PrimitiveElement sigval1 = pe;
		QStyleOption* sigval2 = (QStyleOption*) opt;
		QPainter* sigval3 = p;
		QWidget* sigval4 = (QWidget*) w;

		miqt_exec_callback_QCommonStyle_DrawPrimitive(const_cast<MiqtVirtualQCommonStyle*>(this), handle__DrawPrimitive, sigval1, sigval2, sigval3, sigval4);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DrawPrimitive(PrimitiveElement pe, QStyleOption* opt, QPainter* p, QWidget* w) const {

		QCommonStyle::drawPrimitive(pe, opt, p, w);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DrawControl = 0;

	// Subclass to allow providing a Go implementation
	virtual void drawControl(ControlElement element, const QStyleOption* opt, QPainter* p, const QWidget* w) const override {
		if (handle__DrawControl == 0) {
			QCommonStyle::drawControl(element, opt, p, w);
			return;
		}
		
		ControlElement sigval1 = element;
		QStyleOption* sigval2 = (QStyleOption*) opt;
		QPainter* sigval3 = p;
		QWidget* sigval4 = (QWidget*) w;

		miqt_exec_callback_QCommonStyle_DrawControl(const_cast<MiqtVirtualQCommonStyle*>(this), handle__DrawControl, sigval1, sigval2, sigval3, sigval4);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DrawControl(ControlElement element, QStyleOption* opt, QPainter* p, QWidget* w) const {

		QCommonStyle::drawControl(element, opt, p, w);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SubElementRect = 0;

	// Subclass to allow providing a Go implementation
	virtual QRect subElementRect(SubElement r, const QStyleOption* opt, const QWidget* widget) const override {
		if (handle__SubElementRect == 0) {
			return QCommonStyle::subElementRect(r, opt, widget);
		}
		
		SubElement sigval1 = r;
		QStyleOption* sigval2 = (QStyleOption*) opt;
		QWidget* sigval3 = (QWidget*) widget;

		QRect* callback_return_value = miqt_exec_callback_QCommonStyle_SubElementRect(const_cast<MiqtVirtualQCommonStyle*>(this), handle__SubElementRect, sigval1, sigval2, sigval3);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QRect* virtualbase_SubElementRect(SubElement r, QStyleOption* opt, QWidget* widget) const {

		return new QRect(QCommonStyle::subElementRect(r, opt, widget));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DrawComplexControl = 0;

	// Subclass to allow providing a Go implementation
	virtual void drawComplexControl(ComplexControl cc, const QStyleOptionComplex* opt, QPainter* p, const QWidget* w) const override {
		if (handle__DrawComplexControl == 0) {
			QCommonStyle::drawComplexControl(cc, opt, p, w);
			return;
		}
		
		ComplexControl sigval1 = cc;
		QStyleOptionComplex* sigval2 = (QStyleOptionComplex*) opt;
		QPainter* sigval3 = p;
		QWidget* sigval4 = (QWidget*) w;

		miqt_exec_callback_QCommonStyle_DrawComplexControl(const_cast<MiqtVirtualQCommonStyle*>(this), handle__DrawComplexControl, sigval1, sigval2, sigval3, sigval4);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DrawComplexControl(ComplexControl cc, QStyleOptionComplex* opt, QPainter* p, QWidget* w) const {

		QCommonStyle::drawComplexControl(cc, opt, p, w);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__HitTestComplexControl = 0;

	// Subclass to allow providing a Go implementation
	virtual SubControl hitTestComplexControl(ComplexControl cc, const QStyleOptionComplex* opt, const QPoint& pt, const QWidget* w) const override {
		if (handle__HitTestComplexControl == 0) {
			return QCommonStyle::hitTestComplexControl(cc, opt, pt, w);
		}
		
		ComplexControl sigval1 = cc;
		QStyleOptionComplex* sigval2 = (QStyleOptionComplex*) opt;
		const QPoint& pt_ret = pt;
		// Cast returned reference into pointer
		QPoint* sigval3 = const_cast<QPoint*>(&pt_ret);
		QWidget* sigval4 = (QWidget*) w;

		SubControl callback_return_value = miqt_exec_callback_QCommonStyle_HitTestComplexControl(const_cast<MiqtVirtualQCommonStyle*>(this), handle__HitTestComplexControl, sigval1, sigval2, sigval3, sigval4);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	SubControl virtualbase_HitTestComplexControl(ComplexControl cc, QStyleOptionComplex* opt, QPoint* pt, QWidget* w) const {

		return QCommonStyle::hitTestComplexControl(cc, opt, *pt, w);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SubControlRect = 0;

	// Subclass to allow providing a Go implementation
	virtual QRect subControlRect(ComplexControl cc, const QStyleOptionComplex* opt, SubControl sc, const QWidget* w) const override {
		if (handle__SubControlRect == 0) {
			return QCommonStyle::subControlRect(cc, opt, sc, w);
		}
		
		ComplexControl sigval1 = cc;
		QStyleOptionComplex* sigval2 = (QStyleOptionComplex*) opt;
		SubControl sigval3 = sc;
		QWidget* sigval4 = (QWidget*) w;

		QRect* callback_return_value = miqt_exec_callback_QCommonStyle_SubControlRect(const_cast<MiqtVirtualQCommonStyle*>(this), handle__SubControlRect, sigval1, sigval2, sigval3, sigval4);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QRect* virtualbase_SubControlRect(ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* w) const {

		return new QRect(QCommonStyle::subControlRect(cc, opt, sc, w));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__SizeFromContents = 0;

	// Subclass to allow providing a Go implementation
	virtual QSize sizeFromContents(ContentsType ct, const QStyleOption* opt, const QSize& contentsSize, const QWidget* widget) const override {
		if (handle__SizeFromContents == 0) {
			return QCommonStyle::sizeFromContents(ct, opt, contentsSize, widget);
		}
		
		ContentsType sigval1 = ct;
		QStyleOption* sigval2 = (QStyleOption*) opt;
		const QSize& contentsSize_ret = contentsSize;
		// Cast returned reference into pointer
		QSize* sigval3 = const_cast<QSize*>(&contentsSize_ret);
		QWidget* sigval4 = (QWidget*) widget;

		QSize* callback_return_value = miqt_exec_callback_QCommonStyle_SizeFromContents(const_cast<MiqtVirtualQCommonStyle*>(this), handle__SizeFromContents, sigval1, sigval2, sigval3, sigval4);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QSize* virtualbase_SizeFromContents(ContentsType ct, QStyleOption* opt, QSize* contentsSize, QWidget* widget) const {

		return new QSize(QCommonStyle::sizeFromContents(ct, opt, *contentsSize, widget));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PixelMetric = 0;

	// Subclass to allow providing a Go implementation
	virtual int pixelMetric(PixelMetric m, const QStyleOption* opt, const QWidget* widget) const override {
		if (handle__PixelMetric == 0) {
			return QCommonStyle::pixelMetric(m, opt, widget);
		}
		
		PixelMetric sigval1 = m;
		QStyleOption* sigval2 = (QStyleOption*) opt;
		QWidget* sigval3 = (QWidget*) widget;

		int callback_return_value = miqt_exec_callback_QCommonStyle_PixelMetric(const_cast<MiqtVirtualQCommonStyle*>(this), handle__PixelMetric, sigval1, sigval2, sigval3);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_PixelMetric(PixelMetric m, QStyleOption* opt, QWidget* widget) const {

		return QCommonStyle::pixelMetric(m, opt, widget);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__StyleHint = 0;

	// Subclass to allow providing a Go implementation
	virtual int styleHint(StyleHint sh, const QStyleOption* opt, const QWidget* w, QStyleHintReturn* shret) const override {
		if (handle__StyleHint == 0) {
			return QCommonStyle::styleHint(sh, opt, w, shret);
		}
		
		StyleHint sigval1 = sh;
		QStyleOption* sigval2 = (QStyleOption*) opt;
		QWidget* sigval3 = (QWidget*) w;
		QStyleHintReturn* sigval4 = shret;

		int callback_return_value = miqt_exec_callback_QCommonStyle_StyleHint(const_cast<MiqtVirtualQCommonStyle*>(this), handle__StyleHint, sigval1, sigval2, sigval3, sigval4);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_StyleHint(StyleHint sh, QStyleOption* opt, QWidget* w, QStyleHintReturn* shret) const {

		return QCommonStyle::styleHint(sh, opt, w, shret);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__StandardIcon = 0;

	// Subclass to allow providing a Go implementation
	virtual QIcon standardIcon(StandardPixmap standardIcon, const QStyleOption* opt, const QWidget* widget) const override {
		if (handle__StandardIcon == 0) {
			return QCommonStyle::standardIcon(standardIcon, opt, widget);
		}
		
		StandardPixmap sigval1 = standardIcon;
		QStyleOption* sigval2 = (QStyleOption*) opt;
		QWidget* sigval3 = (QWidget*) widget;

		QIcon* callback_return_value = miqt_exec_callback_QCommonStyle_StandardIcon(const_cast<MiqtVirtualQCommonStyle*>(this), handle__StandardIcon, sigval1, sigval2, sigval3);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QIcon* virtualbase_StandardIcon(StandardPixmap standardIcon, QStyleOption* opt, QWidget* widget) const {

		return new QIcon(QCommonStyle::standardIcon(standardIcon, opt, widget));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__StandardPixmap = 0;

	// Subclass to allow providing a Go implementation
	virtual QPixmap standardPixmap(StandardPixmap sp, const QStyleOption* opt, const QWidget* widget) const override {
		if (handle__StandardPixmap == 0) {
			return QCommonStyle::standardPixmap(sp, opt, widget);
		}
		
		StandardPixmap sigval1 = sp;
		QStyleOption* sigval2 = (QStyleOption*) opt;
		QWidget* sigval3 = (QWidget*) widget;

		QPixmap* callback_return_value = miqt_exec_callback_QCommonStyle_StandardPixmap(const_cast<MiqtVirtualQCommonStyle*>(this), handle__StandardPixmap, sigval1, sigval2, sigval3);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPixmap* virtualbase_StandardPixmap(StandardPixmap sp, QStyleOption* opt, QWidget* widget) const {

		return new QPixmap(QCommonStyle::standardPixmap(sp, opt, widget));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__GeneratedIconPixmap = 0;

	// Subclass to allow providing a Go implementation
	virtual QPixmap generatedIconPixmap(QIcon::Mode iconMode, const QPixmap& pixmap, const QStyleOption* opt) const override {
		if (handle__GeneratedIconPixmap == 0) {
			return QCommonStyle::generatedIconPixmap(iconMode, pixmap, opt);
		}
		
		QIcon::Mode iconMode_ret = iconMode;
		int sigval1 = static_cast<int>(iconMode_ret);
		const QPixmap& pixmap_ret = pixmap;
		// Cast returned reference into pointer
		QPixmap* sigval2 = const_cast<QPixmap*>(&pixmap_ret);
		QStyleOption* sigval3 = (QStyleOption*) opt;

		QPixmap* callback_return_value = miqt_exec_callback_QCommonStyle_GeneratedIconPixmap(const_cast<MiqtVirtualQCommonStyle*>(this), handle__GeneratedIconPixmap, sigval1, sigval2, sigval3);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPixmap* virtualbase_GeneratedIconPixmap(int iconMode, QPixmap* pixmap, QStyleOption* opt) const {

		return new QPixmap(QCommonStyle::generatedIconPixmap(static_cast<QIcon::Mode>(iconMode), *pixmap, opt));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__LayoutSpacing = 0;

	// Subclass to allow providing a Go implementation
	virtual int layoutSpacing(QSizePolicy::ControlType control1, QSizePolicy::ControlType control2, Qt::Orientation orientation, const QStyleOption* option, const QWidget* widget) const override {
		if (handle__LayoutSpacing == 0) {
			return QCommonStyle::layoutSpacing(control1, control2, orientation, option, widget);
		}
		
		QSizePolicy::ControlType control1_ret = control1;
		int sigval1 = static_cast<int>(control1_ret);
		QSizePolicy::ControlType control2_ret = control2;
		int sigval2 = static_cast<int>(control2_ret);
		Qt::Orientation orientation_ret = orientation;
		int sigval3 = static_cast<int>(orientation_ret);
		QStyleOption* sigval4 = (QStyleOption*) option;
		QWidget* sigval5 = (QWidget*) widget;

		int callback_return_value = miqt_exec_callback_QCommonStyle_LayoutSpacing(const_cast<MiqtVirtualQCommonStyle*>(this), handle__LayoutSpacing, sigval1, sigval2, sigval3, sigval4, sigval5);

		return static_cast<int>(callback_return_value);
	}

	// Wrapper to allow calling protected method
	int virtualbase_LayoutSpacing(int control1, int control2, int orientation, QStyleOption* option, QWidget* widget) const {

		return QCommonStyle::layoutSpacing(static_cast<QSizePolicy::ControlType>(control1), static_cast<QSizePolicy::ControlType>(control2), static_cast<Qt::Orientation>(orientation), option, widget);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Polish = 0;

	// Subclass to allow providing a Go implementation
	virtual void polish(QPalette& param1) override {
		if (handle__Polish == 0) {
			QCommonStyle::polish(param1);
			return;
		}
		
		QPalette& param1_ret = param1;
		// Cast returned reference into pointer
		QPalette* sigval1 = &param1_ret;

		miqt_exec_callback_QCommonStyle_Polish(this, handle__Polish, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Polish(QPalette* param1) {

		QCommonStyle::polish(*param1);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PolishWithApp = 0;

	// Subclass to allow providing a Go implementation
	virtual void polish(QApplication* app) override {
		if (handle__PolishWithApp == 0) {
			QCommonStyle::polish(app);
			return;
		}
		
		QApplication* sigval1 = app;

		miqt_exec_callback_QCommonStyle_PolishWithApp(this, handle__PolishWithApp, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_PolishWithApp(QApplication* app) {

		QCommonStyle::polish(app);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__PolishWithWidget = 0;

	// Subclass to allow providing a Go implementation
	virtual void polish(QWidget* widget) override {
		if (handle__PolishWithWidget == 0) {
			QCommonStyle::polish(widget);
			return;
		}
		
		QWidget* sigval1 = widget;

		miqt_exec_callback_QCommonStyle_PolishWithWidget(this, handle__PolishWithWidget, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_PolishWithWidget(QWidget* widget) {

		QCommonStyle::polish(widget);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Unpolish = 0;

	// Subclass to allow providing a Go implementation
	virtual void unpolish(QWidget* widget) override {
		if (handle__Unpolish == 0) {
			QCommonStyle::unpolish(widget);
			return;
		}
		
		QWidget* sigval1 = widget;

		miqt_exec_callback_QCommonStyle_Unpolish(this, handle__Unpolish, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_Unpolish(QWidget* widget) {

		QCommonStyle::unpolish(widget);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__UnpolishWithApplication = 0;

	// Subclass to allow providing a Go implementation
	virtual void unpolish(QApplication* application) override {
		if (handle__UnpolishWithApplication == 0) {
			QCommonStyle::unpolish(application);
			return;
		}
		
		QApplication* sigval1 = application;

		miqt_exec_callback_QCommonStyle_UnpolishWithApplication(this, handle__UnpolishWithApplication, sigval1);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_UnpolishWithApplication(QApplication* application) {

		QCommonStyle::unpolish(application);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ItemTextRect = 0;

	// Subclass to allow providing a Go implementation
	virtual QRect itemTextRect(const QFontMetrics& fm, const QRect& r, int flags, bool enabled, const QString& text) const override {
		if (handle__ItemTextRect == 0) {
			return QCommonStyle::itemTextRect(fm, r, flags, enabled, text);
		}
		
		const QFontMetrics& fm_ret = fm;
		// Cast returned reference into pointer
		QFontMetrics* sigval1 = const_cast<QFontMetrics*>(&fm_ret);
		const QRect& r_ret = r;
		// Cast returned reference into pointer
		QRect* sigval2 = const_cast<QRect*>(&r_ret);
		int sigval3 = flags;
		bool sigval4 = enabled;
		const QString text_ret = text;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray text_b = text_ret.toUtf8();
		struct miqt_string text_ms;
		text_ms.len = text_b.length();
		text_ms.data = static_cast<char*>(malloc(text_ms.len));
		memcpy(text_ms.data, text_b.data(), text_ms.len);
		struct miqt_string sigval5 = text_ms;

		QRect* callback_return_value = miqt_exec_callback_QCommonStyle_ItemTextRect(const_cast<MiqtVirtualQCommonStyle*>(this), handle__ItemTextRect, sigval1, sigval2, sigval3, sigval4, sigval5);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QRect* virtualbase_ItemTextRect(QFontMetrics* fm, QRect* r, int flags, bool enabled, struct miqt_string text) const {
		QString text_QString = QString::fromUtf8(text.data, text.len);

		return new QRect(QCommonStyle::itemTextRect(*fm, *r, static_cast<int>(flags), enabled, text_QString));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__ItemPixmapRect = 0;

	// Subclass to allow providing a Go implementation
	virtual QRect itemPixmapRect(const QRect& r, int flags, const QPixmap& pixmap) const override {
		if (handle__ItemPixmapRect == 0) {
			return QCommonStyle::itemPixmapRect(r, flags, pixmap);
		}
		
		const QRect& r_ret = r;
		// Cast returned reference into pointer
		QRect* sigval1 = const_cast<QRect*>(&r_ret);
		int sigval2 = flags;
		const QPixmap& pixmap_ret = pixmap;
		// Cast returned reference into pointer
		QPixmap* sigval3 = const_cast<QPixmap*>(&pixmap_ret);

		QRect* callback_return_value = miqt_exec_callback_QCommonStyle_ItemPixmapRect(const_cast<MiqtVirtualQCommonStyle*>(this), handle__ItemPixmapRect, sigval1, sigval2, sigval3);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QRect* virtualbase_ItemPixmapRect(QRect* r, int flags, QPixmap* pixmap) const {

		return new QRect(QCommonStyle::itemPixmapRect(*r, static_cast<int>(flags), *pixmap));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DrawItemText = 0;

	// Subclass to allow providing a Go implementation
	virtual void drawItemText(QPainter* painter, const QRect& rect, int flags, const QPalette& pal, bool enabled, const QString& text, QPalette::ColorRole textRole) const override {
		if (handle__DrawItemText == 0) {
			QCommonStyle::drawItemText(painter, rect, flags, pal, enabled, text, textRole);
			return;
		}
		
		QPainter* sigval1 = painter;
		const QRect& rect_ret = rect;
		// Cast returned reference into pointer
		QRect* sigval2 = const_cast<QRect*>(&rect_ret);
		int sigval3 = flags;
		const QPalette& pal_ret = pal;
		// Cast returned reference into pointer
		QPalette* sigval4 = const_cast<QPalette*>(&pal_ret);
		bool sigval5 = enabled;
		const QString text_ret = text;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray text_b = text_ret.toUtf8();
		struct miqt_string text_ms;
		text_ms.len = text_b.length();
		text_ms.data = static_cast<char*>(malloc(text_ms.len));
		memcpy(text_ms.data, text_b.data(), text_ms.len);
		struct miqt_string sigval6 = text_ms;
		QPalette::ColorRole textRole_ret = textRole;
		int sigval7 = static_cast<int>(textRole_ret);

		miqt_exec_callback_QCommonStyle_DrawItemText(const_cast<MiqtVirtualQCommonStyle*>(this), handle__DrawItemText, sigval1, sigval2, sigval3, sigval4, sigval5, sigval6, sigval7);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DrawItemText(QPainter* painter, QRect* rect, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole) const {
		QString text_QString = QString::fromUtf8(text.data, text.len);

		QCommonStyle::drawItemText(painter, *rect, static_cast<int>(flags), *pal, enabled, text_QString, static_cast<QPalette::ColorRole>(textRole));

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__DrawItemPixmap = 0;

	// Subclass to allow providing a Go implementation
	virtual void drawItemPixmap(QPainter* painter, const QRect& rect, int alignment, const QPixmap& pixmap) const override {
		if (handle__DrawItemPixmap == 0) {
			QCommonStyle::drawItemPixmap(painter, rect, alignment, pixmap);
			return;
		}
		
		QPainter* sigval1 = painter;
		const QRect& rect_ret = rect;
		// Cast returned reference into pointer
		QRect* sigval2 = const_cast<QRect*>(&rect_ret);
		int sigval3 = alignment;
		const QPixmap& pixmap_ret = pixmap;
		// Cast returned reference into pointer
		QPixmap* sigval4 = const_cast<QPixmap*>(&pixmap_ret);

		miqt_exec_callback_QCommonStyle_DrawItemPixmap(const_cast<MiqtVirtualQCommonStyle*>(this), handle__DrawItemPixmap, sigval1, sigval2, sigval3, sigval4);

		
	}

	// Wrapper to allow calling protected method
	void virtualbase_DrawItemPixmap(QPainter* painter, QRect* rect, int alignment, QPixmap* pixmap) const {

		QCommonStyle::drawItemPixmap(painter, *rect, static_cast<int>(alignment), *pixmap);

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__StandardPalette = 0;

	// Subclass to allow providing a Go implementation
	virtual QPalette standardPalette() const override {
		if (handle__StandardPalette == 0) {
			return QCommonStyle::standardPalette();
		}
		

		QPalette* callback_return_value = miqt_exec_callback_QCommonStyle_StandardPalette(const_cast<MiqtVirtualQCommonStyle*>(this), handle__StandardPalette);

		return *callback_return_value;
	}

	// Wrapper to allow calling protected method
	QPalette* virtualbase_StandardPalette() const {

		return new QPalette(QCommonStyle::standardPalette());

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

void QCommonStyle_override_virtual_DrawPrimitive(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__DrawPrimitive = slot;
}

void QCommonStyle_virtualbase_DrawPrimitive(const void* self, PrimitiveElement pe, QStyleOption* opt, QPainter* p, QWidget* w) {
	( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_DrawPrimitive(pe, opt, p, w);
}

void QCommonStyle_override_virtual_DrawControl(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__DrawControl = slot;
}

void QCommonStyle_virtualbase_DrawControl(const void* self, ControlElement element, QStyleOption* opt, QPainter* p, QWidget* w) {
	( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_DrawControl(element, opt, p, w);
}

void QCommonStyle_override_virtual_SubElementRect(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__SubElementRect = slot;
}

QRect* QCommonStyle_virtualbase_SubElementRect(const void* self, SubElement r, QStyleOption* opt, QWidget* widget) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_SubElementRect(r, opt, widget);
}

void QCommonStyle_override_virtual_DrawComplexControl(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__DrawComplexControl = slot;
}

void QCommonStyle_virtualbase_DrawComplexControl(const void* self, ComplexControl cc, QStyleOptionComplex* opt, QPainter* p, QWidget* w) {
	( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_DrawComplexControl(cc, opt, p, w);
}

void QCommonStyle_override_virtual_HitTestComplexControl(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__HitTestComplexControl = slot;
}

SubControl QCommonStyle_virtualbase_HitTestComplexControl(const void* self, ComplexControl cc, QStyleOptionComplex* opt, QPoint* pt, QWidget* w) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_HitTestComplexControl(cc, opt, pt, w);
}

void QCommonStyle_override_virtual_SubControlRect(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__SubControlRect = slot;
}

QRect* QCommonStyle_virtualbase_SubControlRect(const void* self, ComplexControl cc, QStyleOptionComplex* opt, SubControl sc, QWidget* w) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_SubControlRect(cc, opt, sc, w);
}

void QCommonStyle_override_virtual_SizeFromContents(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__SizeFromContents = slot;
}

QSize* QCommonStyle_virtualbase_SizeFromContents(const void* self, ContentsType ct, QStyleOption* opt, QSize* contentsSize, QWidget* widget) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_SizeFromContents(ct, opt, contentsSize, widget);
}

void QCommonStyle_override_virtual_PixelMetric(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__PixelMetric = slot;
}

int QCommonStyle_virtualbase_PixelMetric(const void* self, PixelMetric m, QStyleOption* opt, QWidget* widget) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_PixelMetric(m, opt, widget);
}

void QCommonStyle_override_virtual_StyleHint(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__StyleHint = slot;
}

int QCommonStyle_virtualbase_StyleHint(const void* self, StyleHint sh, QStyleOption* opt, QWidget* w, QStyleHintReturn* shret) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_StyleHint(sh, opt, w, shret);
}

void QCommonStyle_override_virtual_StandardIcon(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__StandardIcon = slot;
}

QIcon* QCommonStyle_virtualbase_StandardIcon(const void* self, StandardPixmap standardIcon, QStyleOption* opt, QWidget* widget) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_StandardIcon(standardIcon, opt, widget);
}

void QCommonStyle_override_virtual_StandardPixmap(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__StandardPixmap = slot;
}

QPixmap* QCommonStyle_virtualbase_StandardPixmap(const void* self, StandardPixmap sp, QStyleOption* opt, QWidget* widget) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_StandardPixmap(sp, opt, widget);
}

void QCommonStyle_override_virtual_GeneratedIconPixmap(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__GeneratedIconPixmap = slot;
}

QPixmap* QCommonStyle_virtualbase_GeneratedIconPixmap(const void* self, int iconMode, QPixmap* pixmap, QStyleOption* opt) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_GeneratedIconPixmap(iconMode, pixmap, opt);
}

void QCommonStyle_override_virtual_LayoutSpacing(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__LayoutSpacing = slot;
}

int QCommonStyle_virtualbase_LayoutSpacing(const void* self, int control1, int control2, int orientation, QStyleOption* option, QWidget* widget) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_LayoutSpacing(control1, control2, orientation, option, widget);
}

void QCommonStyle_override_virtual_Polish(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__Polish = slot;
}

void QCommonStyle_virtualbase_Polish(void* self, QPalette* param1) {
	( (MiqtVirtualQCommonStyle*)(self) )->virtualbase_Polish(param1);
}

void QCommonStyle_override_virtual_PolishWithApp(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__PolishWithApp = slot;
}

void QCommonStyle_virtualbase_PolishWithApp(void* self, QApplication* app) {
	( (MiqtVirtualQCommonStyle*)(self) )->virtualbase_PolishWithApp(app);
}

void QCommonStyle_override_virtual_PolishWithWidget(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__PolishWithWidget = slot;
}

void QCommonStyle_virtualbase_PolishWithWidget(void* self, QWidget* widget) {
	( (MiqtVirtualQCommonStyle*)(self) )->virtualbase_PolishWithWidget(widget);
}

void QCommonStyle_override_virtual_Unpolish(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__Unpolish = slot;
}

void QCommonStyle_virtualbase_Unpolish(void* self, QWidget* widget) {
	( (MiqtVirtualQCommonStyle*)(self) )->virtualbase_Unpolish(widget);
}

void QCommonStyle_override_virtual_UnpolishWithApplication(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__UnpolishWithApplication = slot;
}

void QCommonStyle_virtualbase_UnpolishWithApplication(void* self, QApplication* application) {
	( (MiqtVirtualQCommonStyle*)(self) )->virtualbase_UnpolishWithApplication(application);
}

void QCommonStyle_override_virtual_ItemTextRect(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__ItemTextRect = slot;
}

QRect* QCommonStyle_virtualbase_ItemTextRect(const void* self, QFontMetrics* fm, QRect* r, int flags, bool enabled, struct miqt_string text) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_ItemTextRect(fm, r, flags, enabled, text);
}

void QCommonStyle_override_virtual_ItemPixmapRect(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__ItemPixmapRect = slot;
}

QRect* QCommonStyle_virtualbase_ItemPixmapRect(const void* self, QRect* r, int flags, QPixmap* pixmap) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_ItemPixmapRect(r, flags, pixmap);
}

void QCommonStyle_override_virtual_DrawItemText(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__DrawItemText = slot;
}

void QCommonStyle_virtualbase_DrawItemText(const void* self, QPainter* painter, QRect* rect, int flags, QPalette* pal, bool enabled, struct miqt_string text, int textRole) {
	( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_DrawItemText(painter, rect, flags, pal, enabled, text, textRole);
}

void QCommonStyle_override_virtual_DrawItemPixmap(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__DrawItemPixmap = slot;
}

void QCommonStyle_virtualbase_DrawItemPixmap(const void* self, QPainter* painter, QRect* rect, int alignment, QPixmap* pixmap) {
	( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_DrawItemPixmap(painter, rect, alignment, pixmap);
}

void QCommonStyle_override_virtual_StandardPalette(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCommonStyle*>( (QCommonStyle*)(self) )->handle__StandardPalette = slot;
}

QPalette* QCommonStyle_virtualbase_StandardPalette(const void* self) {
	return ( (const MiqtVirtualQCommonStyle*)(self) )->virtualbase_StandardPalette();
}

void QCommonStyle_Delete(QCommonStyle* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQCommonStyle*>( self );
	} else {
		delete self;
	}
}

