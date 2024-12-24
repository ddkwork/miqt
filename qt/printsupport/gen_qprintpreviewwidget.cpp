// +build ignore

#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPrintPreviewWidget>
#include <QPrinter>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qprintpreviewwidget.h>
#include "gen_qprintpreviewwidget.h"
class MiqtVirtualQPrintPreviewWidget : public virtual QPrintPreviewWidget {
public:
MiqtVirtualQPrintPreviewWidget(QWidget* parent): QPrintPreviewWidget(parent) {};
MiqtVirtualQPrintPreviewWidget(QPrinter* printer): QPrintPreviewWidget(printer) {};
MiqtVirtualQPrintPreviewWidget(): QPrintPreviewWidget() {};
MiqtVirtualQPrintPreviewWidget(QPrinter* printer, QWidget* parent): QPrintPreviewWidget(printer, parent) {};
MiqtVirtualQPrintPreviewWidget(QPrinter* printer, QWidget* parent, Qt::WindowFlags flags): QPrintPreviewWidget(printer, parent, flags) {};
MiqtVirtualQPrintPreviewWidget(QWidget* parent, Qt::WindowFlags flags): QPrintPreviewWidget(parent, flags) {};

virtual ~MiqtVirtualQPrintPreviewWidget() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QPrintPreviewWidget::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QPrintPreviewWidget_MetaObject(const_cast<MiqtVirtualQPrintPreviewWidget*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QPrintPreviewWidget::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QPrintPreviewWidget::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QPrintPreviewWidget_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QPrintPreviewWidget::qt_metacast(param1);

	}
};
QPrintPreviewWidget* QPrintPreviewWidget_new(QWidget* parent) {
return new MiqtVirtualQPrintPreviewWidget(parent);
}
QPrintPreviewWidget* QPrintPreviewWidget_new2(QPrinter* printer) {
return new MiqtVirtualQPrintPreviewWidget(printer);
}
QPrintPreviewWidget* QPrintPreviewWidget_new3() {
return new MiqtVirtualQPrintPreviewWidget();
}
QPrintPreviewWidget* QPrintPreviewWidget_new4(QPrinter* printer, QWidget* parent) {
return new MiqtVirtualQPrintPreviewWidget(printer, parent);
}
QPrintPreviewWidget* QPrintPreviewWidget_new5(QPrinter* printer, QWidget* parent, int flags) {
return new MiqtVirtualQPrintPreviewWidget(printer, parent, static_cast<Qt::WindowFlags>(flags));
}
QPrintPreviewWidget* QPrintPreviewWidget_new6(QWidget* parent, int flags) {
return new MiqtVirtualQPrintPreviewWidget(parent, static_cast<Qt::WindowFlags>(flags));
}
void QPrintPreviewWidget_virtbase(QPrintPreviewWidget* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QPrintPreviewWidget_MetaObject(const QPrintPreviewWidget* self) {
	return (QMetaObject*) self->metaObject();
}
void* QPrintPreviewWidget_Metacast(QPrintPreviewWidget* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QPrintPreviewWidget_Tr(const char* s) {
	QString _ret = QPrintPreviewWidget::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
double QPrintPreviewWidget_ZoomFactor(const QPrintPreviewWidget* self) {
	qreal _ret = self->zoomFactor();
	return static_cast<double>(_ret);
}
int QPrintPreviewWidget_Orientation(const QPrintPreviewWidget* self) {
	QPageLayout::Orientation _ret = self->orientation();
	return static_cast<int>(_ret);
}
ViewMode QPrintPreviewWidget_ViewMode(const QPrintPreviewWidget* self) {
	return self->viewMode();
}
ZoomMode QPrintPreviewWidget_ZoomMode(const QPrintPreviewWidget* self) {
	return self->zoomMode();
}
int QPrintPreviewWidget_CurrentPage(const QPrintPreviewWidget* self) {
	return self->currentPage();
}
int QPrintPreviewWidget_PageCount(const QPrintPreviewWidget* self) {
	return self->pageCount();
}
void QPrintPreviewWidget_SetVisible(QPrintPreviewWidget* self, bool visible) {
	self->setVisible(visible);
}
void QPrintPreviewWidget_Print(QPrintPreviewWidget* self) {
	self->print();
}
void QPrintPreviewWidget_ZoomIn(QPrintPreviewWidget* self) {
	self->zoomIn();
}
void QPrintPreviewWidget_ZoomOut(QPrintPreviewWidget* self) {
	self->zoomOut();
}
void QPrintPreviewWidget_SetZoomFactor(QPrintPreviewWidget* self, double zoomFactor) {
	self->setZoomFactor(static_cast<qreal>(zoomFactor));
}
void QPrintPreviewWidget_SetOrientation(QPrintPreviewWidget* self, int orientation) {
	self->setOrientation(static_cast<QPageLayout::Orientation>(orientation));
}
void QPrintPreviewWidget_SetViewMode(QPrintPreviewWidget* self, ViewMode viewMode) {
	self->setViewMode(viewMode);
}
void QPrintPreviewWidget_SetZoomMode(QPrintPreviewWidget* self, ZoomMode zoomMode) {
	self->setZoomMode(zoomMode);
}
void QPrintPreviewWidget_SetCurrentPage(QPrintPreviewWidget* self, int pageNumber) {
	self->setCurrentPage(static_cast<int>(pageNumber));
}
void QPrintPreviewWidget_FitToWidth(QPrintPreviewWidget* self) {
	self->fitToWidth();
}
void QPrintPreviewWidget_FitInView(QPrintPreviewWidget* self) {
	self->fitInView();
}
void QPrintPreviewWidget_SetLandscapeOrientation(QPrintPreviewWidget* self) {
	self->setLandscapeOrientation();
}
void QPrintPreviewWidget_SetPortraitOrientation(QPrintPreviewWidget* self) {
	self->setPortraitOrientation();
}
void QPrintPreviewWidget_SetSinglePageViewMode(QPrintPreviewWidget* self) {
	self->setSinglePageViewMode();
}
void QPrintPreviewWidget_SetFacingPagesViewMode(QPrintPreviewWidget* self) {
	self->setFacingPagesViewMode();
}
void QPrintPreviewWidget_SetAllPagesViewMode(QPrintPreviewWidget* self) {
	self->setAllPagesViewMode();
}
void QPrintPreviewWidget_UpdatePreview(QPrintPreviewWidget* self) {
	self->updatePreview();
}
void QPrintPreviewWidget_PaintRequested(QPrintPreviewWidget* self, QPrinter* printer) {
	self->paintRequested(printer);
}
void QPrintPreviewWidget_connect_PaintRequested(QPrintPreviewWidget* self, intptr_t slot) {
	MiqtVirtualQPrintPreviewWidget::connect(self, static_cast<void (QPrintPreviewWidget::*)(QPrinter*)>(&QPrintPreviewWidget::paintRequested), self, [=](QPrinter* printer) {
		QPrinter* sigval1 = printer;
		miqt_exec_callback_QPrintPreviewWidget_PaintRequested(slot, sigval1);
	});
}
void QPrintPreviewWidget_PreviewChanged(QPrintPreviewWidget* self) {
	self->previewChanged();
}
void QPrintPreviewWidget_connect_PreviewChanged(QPrintPreviewWidget* self, intptr_t slot) {
	MiqtVirtualQPrintPreviewWidget::connect(self, static_cast<void (QPrintPreviewWidget::*)()>(&QPrintPreviewWidget::previewChanged), self, [=]() {
		miqt_exec_callback_QPrintPreviewWidget_PreviewChanged(slot);
	});
}
struct miqt_string QPrintPreviewWidget_Tr2(const char* s, const char* c) {
	QString _ret = QPrintPreviewWidget::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QPrintPreviewWidget_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPrintPreviewWidget::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPrintPreviewWidget_ZoomIn1(QPrintPreviewWidget* self, double zoom) {
	self->zoomIn(static_cast<qreal>(zoom));
}
void QPrintPreviewWidget_ZoomOut1(QPrintPreviewWidget* self, double zoom) {
	self->zoomOut(static_cast<qreal>(zoom));
}
void QPrintPreviewWidget_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrintPreviewWidget*>( (QPrintPreviewWidget*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QPrintPreviewWidget_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQPrintPreviewWidget*)(self) )->virtualbase_MetaObject();
}
void QPrintPreviewWidget_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrintPreviewWidget*>( (QPrintPreviewWidget*)(self) )->handle__Metacast = slot;
}
void* QPrintPreviewWidget_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQPrintPreviewWidget*)(self) )->virtualbase_Metacast(param1);
}
void QPrintPreviewWidget_Delete(QPrintPreviewWidget* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQPrintPreviewWidget*>( self );
	} else {
		delete self;
	}
}
