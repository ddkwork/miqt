// +build ignore

#include <QAction>
#include <QActionEvent>
#include <QBackingStore>
#include <QBitmap>
#include <QByteArray>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QCursor>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEnterEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QFont>
#include <QFontInfo>
#include <QFontMetrics>
#include <QGraphicsEffect>
#include <QGraphicsProxyWidget>
#include <QHideEvent>
#include <QIcon>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QKeySequence>
#include <QLayout>
#include <QList>
#include <QLocale>
#include <QMargins>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEvent>
#include <QPainter>
#include <QPalette>
#include <QPixmap>
#include <QPoint>
#include <QPointF>
#include <QRect>
#include <QRegion>
#include <QResizeEvent>
#include <QScreen>
#include <QShowEvent>
#include <QSize>
#include <QSizePolicy>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyle>
#include <QTabletEvent>
#include <QTimerEvent>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>
#include <QWidgetData>
#include <QWindow>
#include <qwidget.h>
#include "gen_qwidget.h"

#ifndef _Bool
#define _Bool bool
#endif

//void _GUID_Delete(_GUID* self, bool isSubclass) {
//    delete self;
//}
//
//void type_info_Delete(type_info* self, bool isSubclass) {
//	if (isSubclass) {
//		delete dynamic_cast<type_info*>( self );
//	} else {
//		delete self;
//	}
//}

QWidgetData* QWidgetData_new(QWidgetData* param1) {
	return new QWidgetData(*param1);
}

void QWidgetData_OperatorAssign(QWidgetData* self, QWidgetData* param1) {
	self->operator=(*param1);
}

void QWidgetData_Delete(QWidgetData* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QWidgetData*>( self );
	} else {
		delete self;
	}
}

 
QWidget* QWidget_new(QWidget* parent) {
	return new QWidget(parent);
}

QWidget* QWidget_new2() {
	return new QWidget();
}

QWidget* QWidget_new3(QWidget* parent, int f) {
	return new QWidget(parent, static_cast<Qt::WindowFlags>(f));
}

void QWidget_virtbase(QWidget* src, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice) {
	*outptr_QObject = static_cast<QObject*>(src);
	*outptr_QPaintDevice = static_cast<QPaintDevice*>(src);
}

QMetaObject* QWidget_MetaObject(const QWidget* self) {
	return (QMetaObject*) self->metaObject();
}

void* QWidget_Metacast(QWidget* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QWidget_Tr(const char* s) {
	QString _ret = QWidget::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QWidget_DevType(const QWidget* self) {
	return self->devType();
}

uintptr_t QWidget_WinId(const QWidget* self) {
	WId _ret = self->winId();
	return static_cast<uintptr_t>(_ret);
}

void QWidget_CreateWinId(QWidget* self) {
	self->createWinId();
}

uintptr_t QWidget_InternalWinId(const QWidget* self) {
	WId _ret = self->internalWinId();
	return static_cast<uintptr_t>(_ret);
}

uintptr_t QWidget_EffectiveWinId(const QWidget* self) {
	WId _ret = self->effectiveWinId();
	return static_cast<uintptr_t>(_ret);
}

QStyle* QWidget_Style(const QWidget* self) {
	return self->style();
}

void QWidget_SetStyle(QWidget* self, QStyle* style) {
	self->setStyle(style);
}

bool QWidget_IsTopLevel(const QWidget* self) {
	return self->isTopLevel();
}

bool QWidget_IsWindow(const QWidget* self) {
	return self->isWindow();
}

bool QWidget_IsModal(const QWidget* self) {
	return self->isModal();
}

int QWidget_WindowModality(const QWidget* self) {
	Qt::WindowModality _ret = self->windowModality();
	return static_cast<int>(_ret);
}

void QWidget_SetWindowModality(QWidget* self, int windowModality) {
	self->setWindowModality(static_cast<Qt::WindowModality>(windowModality));
}

bool QWidget_IsEnabled(const QWidget* self) {
	return self->isEnabled();
}

bool QWidget_IsEnabledTo(const QWidget* self, QWidget* param1) {
	return self->isEnabledTo(param1);
}

void QWidget_SetEnabled(QWidget* self, bool enabled) {
	self->setEnabled(enabled);
}

void QWidget_SetDisabled(QWidget* self, bool disabled) {
	self->setDisabled(disabled);
}

void QWidget_SetWindowModified(QWidget* self, bool windowModified) {
	self->setWindowModified(windowModified);
}

QRect* QWidget_FrameGeometry(const QWidget* self) {
	return new QRect(self->frameGeometry());
}

QRect* QWidget_Geometry(const QWidget* self) {
	const QRect& _ret = self->geometry();
	// Cast returned reference into pointer
	return const_cast<QRect*>(&_ret);
}

QRect* QWidget_NormalGeometry(const QWidget* self) {
	return new QRect(self->normalGeometry());
}

int QWidget_X(const QWidget* self) {
	return self->x();
}

int QWidget_Y(const QWidget* self) {
	return self->y();
}

QPoint* QWidget_Pos(const QWidget* self) {
	return new QPoint(self->pos());
}

QSize* QWidget_FrameSize(const QWidget* self) {
	return new QSize(self->frameSize());
}

QSize* QWidget_Size(const QWidget* self) {
	return new QSize(self->size());
}

int QWidget_Width(const QWidget* self) {
	return self->width();
}

int QWidget_Height(const QWidget* self) {
	return self->height();
}

QRect* QWidget_Rect(const QWidget* self) {
	return new QRect(self->rect());
}

QRect* QWidget_ChildrenRect(const QWidget* self) {
	return new QRect(self->childrenRect());
}

QRegion* QWidget_ChildrenRegion(const QWidget* self) {
	return new QRegion(self->childrenRegion());
}

QSize* QWidget_MinimumSize(const QWidget* self) {
	return new QSize(self->minimumSize());
}

QSize* QWidget_MaximumSize(const QWidget* self) {
	return new QSize(self->maximumSize());
}

int QWidget_MinimumWidth(const QWidget* self) {
	return self->minimumWidth();
}

int QWidget_MinimumHeight(const QWidget* self) {
	return self->minimumHeight();
}

int QWidget_MaximumWidth(const QWidget* self) {
	return self->maximumWidth();
}

int QWidget_MaximumHeight(const QWidget* self) {
	return self->maximumHeight();
}

void QWidget_SetMinimumSize(QWidget* self, QSize* minimumSize) {
	self->setMinimumSize(*minimumSize);
}

void QWidget_SetMinimumSize2(QWidget* self, int minw, int minh) {
	self->setMinimumSize(static_cast<int>(minw), static_cast<int>(minh));
}

void QWidget_SetMaximumSize(QWidget* self, QSize* maximumSize) {
	self->setMaximumSize(*maximumSize);
}

void QWidget_SetMaximumSize2(QWidget* self, int maxw, int maxh) {
	self->setMaximumSize(static_cast<int>(maxw), static_cast<int>(maxh));
}

void QWidget_SetMinimumWidth(QWidget* self, int minw) {
	self->setMinimumWidth(static_cast<int>(minw));
}

void QWidget_SetMinimumHeight(QWidget* self, int minh) {
	self->setMinimumHeight(static_cast<int>(minh));
}

void QWidget_SetMaximumWidth(QWidget* self, int maxw) {
	self->setMaximumWidth(static_cast<int>(maxw));
}

void QWidget_SetMaximumHeight(QWidget* self, int maxh) {
	self->setMaximumHeight(static_cast<int>(maxh));
}

QSize* QWidget_SizeIncrement(const QWidget* self) {
	return new QSize(self->sizeIncrement());
}

void QWidget_SetSizeIncrement(QWidget* self, QSize* sizeIncrement) {
	self->setSizeIncrement(*sizeIncrement);
}

void QWidget_SetSizeIncrement2(QWidget* self, int w, int h) {
	self->setSizeIncrement(static_cast<int>(w), static_cast<int>(h));
}

QSize* QWidget_BaseSize(const QWidget* self) {
	return new QSize(self->baseSize());
}

void QWidget_SetBaseSize(QWidget* self, QSize* baseSize) {
	self->setBaseSize(*baseSize);
}

void QWidget_SetBaseSize2(QWidget* self, int basew, int baseh) {
	self->setBaseSize(static_cast<int>(basew), static_cast<int>(baseh));
}

void QWidget_SetFixedSize(QWidget* self, QSize* fixedSize) {
	self->setFixedSize(*fixedSize);
}

void QWidget_SetFixedSize2(QWidget* self, int w, int h) {
	self->setFixedSize(static_cast<int>(w), static_cast<int>(h));
}

void QWidget_SetFixedWidth(QWidget* self, int w) {
	self->setFixedWidth(static_cast<int>(w));
}

void QWidget_SetFixedHeight(QWidget* self, int h) {
	self->setFixedHeight(static_cast<int>(h));
}

QPointF* QWidget_MapToGlobal(const QWidget* self, QPointF* param1) {
	return new QPointF(self->mapToGlobal(*param1));
}

QPoint* QWidget_MapToGlobalWithQPoint(const QWidget* self, QPoint* param1) {
	return new QPoint(self->mapToGlobal(*param1));
}

QPointF* QWidget_MapFromGlobal(const QWidget* self, QPointF* param1) {
	return new QPointF(self->mapFromGlobal(*param1));
}

QPoint* QWidget_MapFromGlobalWithQPoint(const QWidget* self, QPoint* param1) {
	return new QPoint(self->mapFromGlobal(*param1));
}

QPointF* QWidget_MapToParent(const QWidget* self, QPointF* param1) {
	return new QPointF(self->mapToParent(*param1));
}

QPoint* QWidget_MapToParentWithQPoint(const QWidget* self, QPoint* param1) {
	return new QPoint(self->mapToParent(*param1));
}

QPointF* QWidget_MapFromParent(const QWidget* self, QPointF* param1) {
	return new QPointF(self->mapFromParent(*param1));
}

QPoint* QWidget_MapFromParentWithQPoint(const QWidget* self, QPoint* param1) {
	return new QPoint(self->mapFromParent(*param1));
}

QPointF* QWidget_MapTo(const QWidget* self, QWidget* param1, QPointF* param2) {
	return new QPointF(self->mapTo(param1, *param2));
}

QPoint* QWidget_MapTo2(const QWidget* self, QWidget* param1, QPoint* param2) {
	return new QPoint(self->mapTo(param1, *param2));
}

QPointF* QWidget_MapFrom(const QWidget* self, QWidget* param1, QPointF* param2) {
	return new QPointF(self->mapFrom(param1, *param2));
}

QPoint* QWidget_MapFrom2(const QWidget* self, QWidget* param1, QPoint* param2) {
	return new QPoint(self->mapFrom(param1, *param2));
}

QWidget* QWidget_Window(const QWidget* self) {
	return self->window();
}

QWidget* QWidget_NativeParentWidget(const QWidget* self) {
	return self->nativeParentWidget();
}

QWidget* QWidget_TopLevelWidget(const QWidget* self) {
	return self->topLevelWidget();
}

QPalette* QWidget_Palette(const QWidget* self) {
	const QPalette& _ret = self->palette();
	// Cast returned reference into pointer
	return const_cast<QPalette*>(&_ret);
}

void QWidget_SetPalette(QWidget* self, QPalette* palette) {
	self->setPalette(*palette);
}

void QWidget_SetBackgroundRole(QWidget* self, int backgroundRole) {
	self->setBackgroundRole(static_cast<QPalette::ColorRole>(backgroundRole));
}

int QWidget_BackgroundRole(const QWidget* self) {
	QPalette::ColorRole _ret = self->backgroundRole();
	return static_cast<int>(_ret);
}

void QWidget_SetForegroundRole(QWidget* self, int foregroundRole) {
	self->setForegroundRole(static_cast<QPalette::ColorRole>(foregroundRole));
}

int QWidget_ForegroundRole(const QWidget* self) {
	QPalette::ColorRole _ret = self->foregroundRole();
	return static_cast<int>(_ret);
}

QFont* QWidget_Font(const QWidget* self) {
	const QFont& _ret = self->font();
	// Cast returned reference into pointer
	return const_cast<QFont*>(&_ret);
}

void QWidget_SetFont(QWidget* self, QFont* font) {
	self->setFont(*font);
}

QFontMetrics* QWidget_FontMetrics(const QWidget* self) {
	return new QFontMetrics(self->fontMetrics());
}

QFontInfo* QWidget_FontInfo(const QWidget* self) {
	return new QFontInfo(self->fontInfo());
}

QCursor* QWidget_Cursor(const QWidget* self) {
	return new QCursor(self->cursor());
}

void QWidget_SetCursor(QWidget* self, QCursor* cursor) {
	self->setCursor(*cursor);
}

void QWidget_UnsetCursor(QWidget* self) {
	self->unsetCursor();
}

void QWidget_SetMouseTracking(QWidget* self, bool enable) {
	self->setMouseTracking(enable);
}

bool QWidget_HasMouseTracking(const QWidget* self) {
	return self->hasMouseTracking();
}

bool QWidget_UnderMouse(const QWidget* self) {
	return self->underMouse();
}

void QWidget_SetTabletTracking(QWidget* self, bool enable) {
	self->setTabletTracking(enable);
}

bool QWidget_HasTabletTracking(const QWidget* self) {
	return self->hasTabletTracking();
}

void QWidget_SetMask(QWidget* self, QBitmap* mask) {
	self->setMask(*mask);
}

void QWidget_SetMaskWithMask(QWidget* self, QRegion* mask) {
	self->setMask(*mask);
}

QRegion* QWidget_Mask(const QWidget* self) {
	return new QRegion(self->mask());
}

void QWidget_ClearMask(QWidget* self) {
	self->clearMask();
}

void QWidget_Render(QWidget* self, QPaintDevice* target) {
	self->render(target);
}

void QWidget_RenderWithPainter(QWidget* self, QPainter* painter) {
	self->render(painter);
}

QPixmap* QWidget_Grab(QWidget* self) {
	return new QPixmap(self->grab());
}

QGraphicsEffect* QWidget_GraphicsEffect(const QWidget* self) {
	return self->graphicsEffect();
}

void QWidget_SetGraphicsEffect(QWidget* self, QGraphicsEffect* effect) {
	self->setGraphicsEffect(effect);
}

void QWidget_GrabGesture(QWidget* self, int typeVal) {
	self->grabGesture(static_cast<Qt::GestureType>(typeVal));
}

void QWidget_UngrabGesture(QWidget* self, int typeVal) {
	self->ungrabGesture(static_cast<Qt::GestureType>(typeVal));
}

void QWidget_SetWindowTitle(QWidget* self, struct miqt_string windowTitle) {
	QString windowTitle_QString = QString::fromUtf8(windowTitle.data, windowTitle.len);
	self->setWindowTitle(windowTitle_QString);
}

void QWidget_SetStyleSheet(QWidget* self, struct miqt_string styleSheet) {
	QString styleSheet_QString = QString::fromUtf8(styleSheet.data, styleSheet.len);
	self->setStyleSheet(styleSheet_QString);
}

struct miqt_string QWidget_StyleSheet(const QWidget* self) {
	QString _ret = self->styleSheet();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QWidget_WindowTitle(const QWidget* self) {
	QString _ret = self->windowTitle();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWidget_SetWindowIcon(QWidget* self, QIcon* icon) {
	self->setWindowIcon(*icon);
}

QIcon* QWidget_WindowIcon(const QWidget* self) {
	return new QIcon(self->windowIcon());
}

void QWidget_SetWindowIconText(QWidget* self, struct miqt_string windowIconText) {
	QString windowIconText_QString = QString::fromUtf8(windowIconText.data, windowIconText.len);
	self->setWindowIconText(windowIconText_QString);
}

struct miqt_string QWidget_WindowIconText(const QWidget* self) {
	QString _ret = self->windowIconText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWidget_SetWindowRole(QWidget* self, struct miqt_string windowRole) {
	QString windowRole_QString = QString::fromUtf8(windowRole.data, windowRole.len);
	self->setWindowRole(windowRole_QString);
}

struct miqt_string QWidget_WindowRole(const QWidget* self) {
	QString _ret = self->windowRole();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWidget_SetWindowFilePath(QWidget* self, struct miqt_string filePath) {
	QString filePath_QString = QString::fromUtf8(filePath.data, filePath.len);
	self->setWindowFilePath(filePath_QString);
}

struct miqt_string QWidget_WindowFilePath(const QWidget* self) {
	QString _ret = self->windowFilePath();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWidget_SetWindowOpacity(QWidget* self, double level) {
	self->setWindowOpacity(static_cast<qreal>(level));
}

double QWidget_WindowOpacity(const QWidget* self) {
	qreal _ret = self->windowOpacity();
	return static_cast<double>(_ret);
}

bool QWidget_IsWindowModified(const QWidget* self) {
	return self->isWindowModified();
}

void QWidget_SetToolTip(QWidget* self, struct miqt_string toolTip) {
	QString toolTip_QString = QString::fromUtf8(toolTip.data, toolTip.len);
	self->setToolTip(toolTip_QString);
}

struct miqt_string QWidget_ToolTip(const QWidget* self) {
	QString _ret = self->toolTip();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWidget_SetToolTipDuration(QWidget* self, int msec) {
	self->setToolTipDuration(static_cast<int>(msec));
}

int QWidget_ToolTipDuration(const QWidget* self) {
	return self->toolTipDuration();
}

void QWidget_SetStatusTip(QWidget* self, struct miqt_string statusTip) {
	QString statusTip_QString = QString::fromUtf8(statusTip.data, statusTip.len);
	self->setStatusTip(statusTip_QString);
}

struct miqt_string QWidget_StatusTip(const QWidget* self) {
	QString _ret = self->statusTip();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWidget_SetWhatsThis(QWidget* self, struct miqt_string whatsThis) {
	QString whatsThis_QString = QString::fromUtf8(whatsThis.data, whatsThis.len);
	self->setWhatsThis(whatsThis_QString);
}

struct miqt_string QWidget_WhatsThis(const QWidget* self) {
	QString _ret = self->whatsThis();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWidget_SetLayoutDirection(QWidget* self, int direction) {
	self->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

int QWidget_LayoutDirection(const QWidget* self) {
	Qt::LayoutDirection _ret = self->layoutDirection();
	return static_cast<int>(_ret);
}

void QWidget_UnsetLayoutDirection(QWidget* self) {
	self->unsetLayoutDirection();
}

void QWidget_SetLocale(QWidget* self, QLocale* locale) {
	self->setLocale(*locale);
}

QLocale* QWidget_Locale(const QWidget* self) {
	return new QLocale(self->locale());
}

void QWidget_UnsetLocale(QWidget* self) {
	self->unsetLocale();
}

bool QWidget_IsRightToLeft(const QWidget* self) {
	return self->isRightToLeft();
}

bool QWidget_IsLeftToRight(const QWidget* self) {
	return self->isLeftToRight();
}

void QWidget_SetFocus(QWidget* self) {
	self->setFocus();
}

bool QWidget_IsActiveWindow(const QWidget* self) {
	return self->isActiveWindow();
}

void QWidget_ActivateWindow(QWidget* self) {
	self->activateWindow();
}

void QWidget_ClearFocus(QWidget* self) {
	self->clearFocus();
}

void QWidget_SetFocusWithReason(QWidget* self, int reason) {
	self->setFocus(static_cast<Qt::FocusReason>(reason));
}

int QWidget_FocusPolicy(const QWidget* self) {
	Qt::FocusPolicy _ret = self->focusPolicy();
	return static_cast<int>(_ret);
}

void QWidget_SetFocusPolicy(QWidget* self, int policy) {
	self->setFocusPolicy(static_cast<Qt::FocusPolicy>(policy));
}

bool QWidget_HasFocus(const QWidget* self) {
	return self->hasFocus();
}

void QWidget_SetTabOrder(QWidget* param1, QWidget* param2) {
	QWidget::setTabOrder(param1, param2);
}

void QWidget_SetFocusProxy(QWidget* self, QWidget* focusProxy) {
	self->setFocusProxy(focusProxy);
}

QWidget* QWidget_FocusProxy(const QWidget* self) {
	return self->focusProxy();
}

int QWidget_ContextMenuPolicy(const QWidget* self) {
	Qt::ContextMenuPolicy _ret = self->contextMenuPolicy();
	return static_cast<int>(_ret);
}

void QWidget_SetContextMenuPolicy(QWidget* self, int policy) {
	self->setContextMenuPolicy(static_cast<Qt::ContextMenuPolicy>(policy));
}

void QWidget_GrabMouse(QWidget* self) {
	self->grabMouse();
}

void QWidget_GrabMouseWithQCursor(QWidget* self, QCursor* param1) {
	self->grabMouse(*param1);
}

void QWidget_ReleaseMouse(QWidget* self) {
	self->releaseMouse();
}

void QWidget_GrabKeyboard(QWidget* self) {
	self->grabKeyboard();
}

void QWidget_ReleaseKeyboard(QWidget* self) {
	self->releaseKeyboard();
}

int QWidget_GrabShortcut(QWidget* self, QKeySequence* key) {
	return self->grabShortcut(*key);
}

void QWidget_ReleaseShortcut(QWidget* self, int id) {
	self->releaseShortcut(static_cast<int>(id));
}

void QWidget_SetShortcutEnabled(QWidget* self, int id) {
	self->setShortcutEnabled(static_cast<int>(id));
}

void QWidget_SetShortcutAutoRepeat(QWidget* self, int id) {
	self->setShortcutAutoRepeat(static_cast<int>(id));
}

QWidget* QWidget_MouseGrabber() {
	return QWidget::mouseGrabber();
}

QWidget* QWidget_KeyboardGrabber() {
	return QWidget::keyboardGrabber();
}

bool QWidget_UpdatesEnabled(const QWidget* self) {
	return self->updatesEnabled();
}

void QWidget_SetUpdatesEnabled(QWidget* self, bool enable) {
	self->setUpdatesEnabled(enable);
}

QGraphicsProxyWidget* QWidget_GraphicsProxyWidget(const QWidget* self) {
	return self->graphicsProxyWidget();
}

void QWidget_Update(QWidget* self) {
	self->update();
}

void QWidget_Repaint(QWidget* self) {
	self->repaint();
}

void QWidget_Update2(QWidget* self, int x, int y, int w, int h) {
	self->update(static_cast<int>(x), static_cast<int>(y), static_cast<int>(w), static_cast<int>(h));
}

void QWidget_UpdateWithQRect(QWidget* self, QRect* param1) {
	self->update(*param1);
}

void QWidget_UpdateWithQRegion(QWidget* self, QRegion* param1) {
	self->update(*param1);
}

void QWidget_Repaint2(QWidget* self, int x, int y, int w, int h) {
	self->repaint(static_cast<int>(x), static_cast<int>(y), static_cast<int>(w), static_cast<int>(h));
}

void QWidget_RepaintWithQRect(QWidget* self, QRect* param1) {
	self->repaint(*param1);
}

void QWidget_RepaintWithQRegion(QWidget* self, QRegion* param1) {
	self->repaint(*param1);
}

void QWidget_SetVisible(QWidget* self, bool visible) {
	self->setVisible(visible);
}

void QWidget_SetHidden(QWidget* self, bool hidden) {
	self->setHidden(hidden);
}

void QWidget_Show(QWidget* self) {
	self->show();
}

void QWidget_Hide(QWidget* self) {
	self->hide();
}

void QWidget_ShowMinimized(QWidget* self) {
	self->showMinimized();
}

void QWidget_ShowMaximized(QWidget* self) {
	self->showMaximized();
}

void QWidget_ShowFullScreen(QWidget* self) {
	self->showFullScreen();
}

void QWidget_ShowNormal(QWidget* self) {
	self->showNormal();
}

bool QWidget_Close(QWidget* self) {
	return self->close();
}

void QWidget_Raise(QWidget* self) {
	self->raise();
}

void QWidget_Lower(QWidget* self) {
	self->lower();
}

void QWidget_StackUnder(QWidget* self, QWidget* param1) {
	self->stackUnder(param1);
}

void QWidget_Move(QWidget* self, int x, int y) {
	self->move(static_cast<int>(x), static_cast<int>(y));
}

void QWidget_MoveWithQPoint(QWidget* self, QPoint* param1) {
	self->move(*param1);
}

void QWidget_Resize(QWidget* self, int w, int h) {
	self->resize(static_cast<int>(w), static_cast<int>(h));
}

void QWidget_ResizeWithQSize(QWidget* self, QSize* param1) {
	self->resize(*param1);
}

void QWidget_SetGeometry(QWidget* self, int x, int y, int w, int h) {
	self->setGeometry(static_cast<int>(x), static_cast<int>(y), static_cast<int>(w), static_cast<int>(h));
}

void QWidget_SetGeometryWithGeometry(QWidget* self, QRect* geometry) {
	self->setGeometry(*geometry);
}

struct miqt_string QWidget_SaveGeometry(const QWidget* self) {
	QByteArray _qb = self->saveGeometry();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

bool QWidget_RestoreGeometry(QWidget* self, struct miqt_string geometry) {
	QByteArray geometry_QByteArray(geometry.data, geometry.len);
	return self->restoreGeometry(geometry_QByteArray);
}

void QWidget_AdjustSize(QWidget* self) {
	self->adjustSize();
}

bool QWidget_IsVisible(const QWidget* self) {
	return self->isVisible();
}

bool QWidget_IsVisibleTo(const QWidget* self, QWidget* param1) {
	return self->isVisibleTo(param1);
}

bool QWidget_IsHidden(const QWidget* self) {
	return self->isHidden();
}

bool QWidget_IsMinimized(const QWidget* self) {
	return self->isMinimized();
}

bool QWidget_IsMaximized(const QWidget* self) {
	return self->isMaximized();
}

bool QWidget_IsFullScreen(const QWidget* self) {
	return self->isFullScreen();
}

int QWidget_WindowState(const QWidget* self) {
	Qt::WindowStates _ret = self->windowState();
	return static_cast<int>(_ret);
}

void QWidget_SetWindowState(QWidget* self, int state) {
	self->setWindowState(static_cast<Qt::WindowStates>(state));
}

void QWidget_OverrideWindowState(QWidget* self, int state) {
	self->overrideWindowState(static_cast<Qt::WindowStates>(state));
}

QSize* QWidget_SizeHint(const QWidget* self) {
	return new QSize(self->sizeHint());
}

QSize* QWidget_MinimumSizeHint(const QWidget* self) {
	return new QSize(self->minimumSizeHint());
}

QSizePolicy* QWidget_SizePolicy(const QWidget* self) {
	return new QSizePolicy(self->sizePolicy());
}

void QWidget_SetSizePolicy(QWidget* self, QSizePolicy* sizePolicy) {
	self->setSizePolicy(*sizePolicy);
}

void QWidget_SetSizePolicy2(QWidget* self, int horizontal, int vertical) {
	self->setSizePolicy(static_cast<QSizePolicy::Policy>(horizontal), static_cast<QSizePolicy::Policy>(vertical));
}

int QWidget_HeightForWidth(const QWidget* self, int param1) {
	return self->heightForWidth(static_cast<int>(param1));
}

bool QWidget_HasHeightForWidth(const QWidget* self) {
	return self->hasHeightForWidth();
}

QRegion* QWidget_VisibleRegion(const QWidget* self) {
	return new QRegion(self->visibleRegion());
}

void QWidget_SetContentsMargins(QWidget* self, int left, int top, int right, int bottom) {
	self->setContentsMargins(static_cast<int>(left), static_cast<int>(top), static_cast<int>(right), static_cast<int>(bottom));
}

void QWidget_SetContentsMarginsWithMargins(QWidget* self, QMargins* margins) {
	self->setContentsMargins(*margins);
}

QMargins* QWidget_ContentsMargins(const QWidget* self) {
	return new QMargins(self->contentsMargins());
}

QRect* QWidget_ContentsRect(const QWidget* self) {
	return new QRect(self->contentsRect());
}

QLayout* QWidget_Layout(const QWidget* self) {
	return self->layout();
}

void QWidget_SetLayout(QWidget* self, QLayout* layout) {
	self->setLayout(layout);
}

void QWidget_UpdateGeometry(QWidget* self) {
	self->updateGeometry();
}

void QWidget_SetParent(QWidget* self, QWidget* parent) {
	self->setParent(parent);
}

void QWidget_SetParent2(QWidget* self, QWidget* parent, int f) {
	self->setParent(parent, static_cast<Qt::WindowFlags>(f));
}

void QWidget_Scroll(QWidget* self, int dx, int dy) {
	self->scroll(static_cast<int>(dx), static_cast<int>(dy));
}

void QWidget_Scroll2(QWidget* self, int dx, int dy, QRect* param3) {
	self->scroll(static_cast<int>(dx), static_cast<int>(dy), *param3);
}

QWidget* QWidget_FocusWidget(const QWidget* self) {
	return self->focusWidget();
}

QWidget* QWidget_NextInFocusChain(const QWidget* self) {
	return self->nextInFocusChain();
}

QWidget* QWidget_PreviousInFocusChain(const QWidget* self) {
	return self->previousInFocusChain();
}

bool QWidget_AcceptDrops(const QWidget* self) {
	return self->acceptDrops();
}

void QWidget_SetAcceptDrops(QWidget* self, bool on) {
	self->setAcceptDrops(on);
}

void QWidget_AddAction(QWidget* self, QAction* action) {
	self->addAction(action);
}

void QWidget_AddActions(QWidget* self, struct miqt_array /* of QAction* */  actions) {
	QList<QAction *> actions_QList;
	actions_QList.reserve(actions.len);
	QAction** actions_arr = static_cast<QAction**>(actions.data);
	for(size_t i = 0; i < actions.len; ++i) {
		actions_QList.push_back(actions_arr[i]);
	}
	self->addActions(actions_QList);
}

void QWidget_InsertActions(QWidget* self, QAction* before, struct miqt_array /* of QAction* */  actions) {
	QList<QAction *> actions_QList;
	actions_QList.reserve(actions.len);
	QAction** actions_arr = static_cast<QAction**>(actions.data);
	for(size_t i = 0; i < actions.len; ++i) {
		actions_QList.push_back(actions_arr[i]);
	}
	self->insertActions(before, actions_QList);
}

void QWidget_InsertAction(QWidget* self, QAction* before, QAction* action) {
	self->insertAction(before, action);
}

void QWidget_RemoveAction(QWidget* self, QAction* action) {
	self->removeAction(action);
}

struct miqt_array /* of QAction* */  QWidget_Actions(const QWidget* self) {
	QList<QAction *> _ret = self->actions();
	// Convert QList<> from C++ memory to manually-managed C memory
	QAction** _arr = static_cast<QAction**>(malloc(sizeof(QAction*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

QAction* QWidget_AddActionWithText(QWidget* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->addAction(text_QString);
}

QAction* QWidget_AddAction2(QWidget* self, QIcon* icon, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->addAction(*icon, text_QString);
}

QAction* QWidget_AddAction3(QWidget* self, struct miqt_string text, QKeySequence* shortcut) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->addAction(text_QString, *shortcut);
}

QAction* QWidget_AddAction4(QWidget* self, QIcon* icon, struct miqt_string text, QKeySequence* shortcut) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->addAction(*icon, text_QString, *shortcut);
}

QWidget* QWidget_ParentWidget(const QWidget* self) {
	return self->parentWidget();
}

void QWidget_SetWindowFlags(QWidget* self, int typeVal) {
	self->setWindowFlags(static_cast<Qt::WindowFlags>(typeVal));
}

int QWidget_WindowFlags(const QWidget* self) {
	Qt::WindowFlags _ret = self->windowFlags();
	return static_cast<int>(_ret);
}

void QWidget_SetWindowFlag(QWidget* self, int param1) {
	self->setWindowFlag(static_cast<Qt::WindowType>(param1));
}

void QWidget_OverrideWindowFlags(QWidget* self, int typeVal) {
	self->overrideWindowFlags(static_cast<Qt::WindowFlags>(typeVal));
}

int QWidget_WindowType(const QWidget* self) {
	Qt::WindowType _ret = self->windowType();
	return static_cast<int>(_ret);
}

QWidget* QWidget_Find(uintptr_t param1) {
	return QWidget::find(static_cast<WId>(param1));
}

QWidget* QWidget_ChildAt(const QWidget* self, int x, int y) {
	return self->childAt(static_cast<int>(x), static_cast<int>(y));
}

QWidget* QWidget_ChildAtWithQPoint(const QWidget* self, QPoint* p) {
	return self->childAt(*p);
}

QWidget* QWidget_ChildAtWithQPointF(const QWidget* self, QPointF* p) {
	return self->childAt(*p);
}

void QWidget_SetAttribute(QWidget* self, int param1) {
	self->setAttribute(static_cast<Qt::WidgetAttribute>(param1));
}

bool QWidget_TestAttribute(const QWidget* self, int param1) {
	return self->testAttribute(static_cast<Qt::WidgetAttribute>(param1));
}

QPaintEngine* QWidget_PaintEngine(const QWidget* self) {
	return self->paintEngine();
}

void QWidget_EnsurePolished(const QWidget* self) {
	self->ensurePolished();
}

bool QWidget_IsAncestorOf(const QWidget* self, QWidget* child) {
	return self->isAncestorOf(child);
}

bool QWidget_AutoFillBackground(const QWidget* self) {
	return self->autoFillBackground();
}

void QWidget_SetAutoFillBackground(QWidget* self, bool enabled) {
	self->setAutoFillBackground(enabled);
}

QBackingStore* QWidget_BackingStore(const QWidget* self) {
	return self->backingStore();
}

QWindow* QWidget_WindowHandle(const QWidget* self) {
	return self->windowHandle();
}

QScreen* QWidget_Screen(const QWidget* self) {
	return self->screen();
}

void QWidget_SetScreen(QWidget* self, QScreen* screen) {
	self->setScreen(screen);
}

QWidget* QWidget_CreateWindowContainer(QWindow* window) {
	return QWidget::createWindowContainer(window);
}

void QWidget_WindowTitleChanged(QWidget* self, struct miqt_string title) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	self->windowTitleChanged(title_QString);
}

void QWidget_connect_WindowTitleChanged(QWidget* self, intptr_t slot) {
	QWidget::connect(self, static_cast<void (QWidget::*)(const QString&)>(&QWidget::windowTitleChanged), self, [=](const QString& title) {
		const QString title_ret = title;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray title_b = title_ret.toUtf8();
		struct miqt_string title_ms;
		title_ms.len = title_b.length();
		title_ms.data = static_cast<char*>(malloc(title_ms.len));
		memcpy(title_ms.data, title_b.data(), title_ms.len);
		struct miqt_string sigval1 = title_ms;
//		miqt_exec_callback_QWidget_WindowTitleChanged(slot, sigval1); //todo
	});
}

void QWidget_WindowIconChanged(QWidget* self, QIcon* icon) {
	self->windowIconChanged(*icon);
}

void QWidget_connect_WindowIconChanged(QWidget* self, intptr_t slot) {
	QWidget::connect(self, static_cast<void (QWidget::*)(const QIcon&)>(&QWidget::windowIconChanged), self, [=](const QIcon& icon) {
		const QIcon& icon_ret = icon;
		// Cast returned reference into pointer
		QIcon* sigval1 = const_cast<QIcon*>(&icon_ret);
//		miqt_exec_callback_QWidget_WindowIconChanged(slot, sigval1);
	});
}

void QWidget_WindowIconTextChanged(QWidget* self, struct miqt_string iconText) {
	QString iconText_QString = QString::fromUtf8(iconText.data, iconText.len);
	self->windowIconTextChanged(iconText_QString);
}

void QWidget_connect_WindowIconTextChanged(QWidget* self, intptr_t slot) {
	QWidget::connect(self, static_cast<void (QWidget::*)(const QString&)>(&QWidget::windowIconTextChanged), self, [=](const QString& iconText) {
		const QString iconText_ret = iconText;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray iconText_b = iconText_ret.toUtf8();
		struct miqt_string iconText_ms;
		iconText_ms.len = iconText_b.length();
		iconText_ms.data = static_cast<char*>(malloc(iconText_ms.len));
		memcpy(iconText_ms.data, iconText_b.data(), iconText_ms.len);
		struct miqt_string sigval1 = iconText_ms;
//		miqt_exec_callback_QWidget_WindowIconTextChanged(slot, sigval1);
	});
}

void QWidget_CustomContextMenuRequested(QWidget* self, QPoint* pos) {
	self->customContextMenuRequested(*pos);
}

void QWidget_connect_CustomContextMenuRequested(QWidget* self, intptr_t slot) {
	QWidget::connect(self, static_cast<void (QWidget::*)(const QPoint&)>(&QWidget::customContextMenuRequested), self, [=](const QPoint& pos) {
		const QPoint& pos_ret = pos;
		// Cast returned reference into pointer
		QPoint* sigval1 = const_cast<QPoint*>(&pos_ret);
//		miqt_exec_callback_QWidget_CustomContextMenuRequested(slot, sigval1);
	});
}

QVariant* QWidget_InputMethodQuery(const QWidget* self, int param1) {
	return new QVariant(self->inputMethodQuery(static_cast<Qt::InputMethodQuery>(param1)));
}

int QWidget_InputMethodHints(const QWidget* self) {
	Qt::InputMethodHints _ret = self->inputMethodHints();
	return static_cast<int>(_ret);
}

void QWidget_SetInputMethodHints(QWidget* self, int hints) {
	self->setInputMethodHints(static_cast<Qt::InputMethodHints>(hints));
}

struct miqt_string QWidget_Tr2(const char* s, const char* c) {
	QString _ret = QWidget::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QWidget_Tr3(const char* s, const char* c, int n) {
	QString _ret = QWidget::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QWidget_Render2(QWidget* self, QPaintDevice* target, QPoint* targetOffset) {
	self->render(target, *targetOffset);
}

void QWidget_Render3(QWidget* self, QPaintDevice* target, QPoint* targetOffset, QRegion* sourceRegion) {
	self->render(target, *targetOffset, *sourceRegion);
}

//void QWidget_Render4(QWidget* self, QPaintDevice* target, QPoint* targetOffset, QRegion* sourceRegion, RenderFlags renderFlags) {
//	self->render(target, *targetOffset, *sourceRegion, renderFlags);
//}

void QWidget_Render22(QWidget* self, QPainter* painter, QPoint* targetOffset) {
	self->render(painter, *targetOffset);
}

void QWidget_Render32(QWidget* self, QPainter* painter, QPoint* targetOffset, QRegion* sourceRegion) {
	self->render(painter, *targetOffset, *sourceRegion);
}

//void QWidget_Render42(QWidget* self, QPainter* painter, QPoint* targetOffset, QRegion* sourceRegion, RenderFlags renderFlags) {
//	self->render(painter, *targetOffset, *sourceRegion, renderFlags);
//}

QPixmap* QWidget_Grab1(QWidget* self, QRect* rectangle) {
	return new QPixmap(self->grab(*rectangle));
}

void QWidget_GrabGesture2(QWidget* self, int typeVal, int flags) {
	self->grabGesture(static_cast<Qt::GestureType>(typeVal), static_cast<Qt::GestureFlags>(flags));
}

int QWidget_GrabShortcut2(QWidget* self, QKeySequence* key, int context) {
	return self->grabShortcut(*key, static_cast<Qt::ShortcutContext>(context));
}

void QWidget_SetShortcutEnabled2(QWidget* self, int id, bool enable) {
	self->setShortcutEnabled(static_cast<int>(id), enable);
}

void QWidget_SetShortcutAutoRepeat2(QWidget* self, int id, bool enable) {
	self->setShortcutAutoRepeat(static_cast<int>(id), enable);
}

void QWidget_SetWindowFlag2(QWidget* self, int param1, bool on) {
	self->setWindowFlag(static_cast<Qt::WindowType>(param1), on);
}

void QWidget_SetAttribute2(QWidget* self, int param1, bool on) {
	self->setAttribute(static_cast<Qt::WidgetAttribute>(param1), on);
}

QWidget* QWidget_CreateWindowContainer2(QWindow* window, QWidget* parent) {
	return QWidget::createWindowContainer(window, parent);
}

QWidget* QWidget_CreateWindowContainer3(QWindow* window, QWidget* parent, int flags) {
	return QWidget::createWindowContainer(window, parent, static_cast<Qt::WindowFlags>(flags));
}

void QWidget_override_virtual_DevType(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__DevType = slot;
}

int QWidget_virtualbase_DevType(const void* self) {
//	return ( (const QWidget*)(self) )->virtualbase_DevType();
return 0;
}

void QWidget_override_virtual_SetVisible(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__SetVisible = slot;
}

void QWidget_virtualbase_SetVisible(void* self, bool visible) {
//	( (QWidget*)(self) )->virtualbase_SetVisible(visible);
}

void QWidget_override_virtual_SizeHint(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__SizeHint = slot;
}

QSize* QWidget_virtualbase_SizeHint(const void* self) {
    auto size= ( (const QWidget*)(self) )->sizeHint();
    return new QSize(size.width(), size.height());
}

void QWidget_override_virtual_MinimumSizeHint(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__MinimumSizeHint = slot;
}

QSize* QWidget_virtualbase_MinimumSizeHint(const void* self) {
	auto size= ( (const QWidget*)(self) )->minimumSizeHint();
    return new QSize(size.width(), size.height());
}

void QWidget_override_virtual_HeightForWidth(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__HeightForWidth = slot;
}

int QWidget_heightForWidth(const void* self, int param1) {
	return ( (const QWidget*)(self) )->heightForWidth(param1);
}

void QWidget_override_virtual_HasHeightForWidth(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__HasHeightForWidth = slot;
}

bool QWidget_hasHeightForWidth(const void* self) {
	return ( (const QWidget*)(self) )->hasHeightForWidth();
}

void QWidget_override_virtual_PaintEngine(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__PaintEngine = slot;
}

QPaintEngine* QWidget_paintEngine(const void* self) {
	return ( (const QWidget*)(self) )->paintEngine();
}

void QWidget_override_virtual_Event(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__Event = slot;
}

bool QWidget_virtualbase_Event(void* self, QEvent* event) {
//	return ( (QWidget*)(self) )->virtualbase_Event(event);
return false;
}

void QWidget_override_virtual_MousePressEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__MousePressEvent = slot;
}

void QWidget_virtualbase_MousePressEvent(void* self, QMouseEvent* event) {
//	( (QWidget*)(self) )->virtualbase_MousePressEvent(event);
}

void QWidget_override_virtual_MouseReleaseEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__MouseReleaseEvent = slot;
}

void QWidget_virtualbase_MouseReleaseEvent(void* self, QMouseEvent* event) {
//	( (QWidget*)(self) )->virtualbase_MouseReleaseEvent(event);
}

void QWidget_override_virtual_MouseDoubleClickEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__MouseDoubleClickEvent = slot;
}

void QWidget_virtualbase_MouseDoubleClickEvent(void* self, QMouseEvent* event) {
//	( (QWidget*)(self) )->virtualbase_MouseDoubleClickEvent(event);
}

void QWidget_override_virtual_MouseMoveEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__MouseMoveEvent = slot;
}

void QWidget_virtualbase_MouseMoveEvent(void* self, QMouseEvent* event) {
//	( (QWidget*)(self) )->virtualbase_MouseMoveEvent(event);
}

void QWidget_override_virtual_WheelEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__WheelEvent = slot;
}

void QWidget_virtualbase_WheelEvent(void* self, QWheelEvent* event) {
//	( (QWidget*)(self) )->virtualbase_WheelEvent(event);
}

void QWidget_override_virtual_KeyPressEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__KeyPressEvent = slot;
}

void QWidget_virtualbase_KeyPressEvent(void* self, QKeyEvent* event) {
//	( (QWidget*)(self) )->virtualbase_KeyPressEvent(event);
}

void QWidget_override_virtual_KeyReleaseEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__KeyReleaseEvent = slot;
}

void QWidget_virtualbase_KeyReleaseEvent(void* self, QKeyEvent* event) {
//	( (QWidget*)(self) )->virtualbase_KeyReleaseEvent(event);
}

void QWidget_override_virtual_FocusInEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__FocusInEvent = slot;
}

void QWidget_virtualbase_FocusInEvent(void* self, QFocusEvent* event) {
//	( (QWidget*)(self) )->virtualbase_FocusInEvent(event);
}

void QWidget_override_virtual_FocusOutEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__FocusOutEvent = slot;
}

void QWidget_virtualbase_FocusOutEvent(void* self, QFocusEvent* event) {
//	( (QWidget*)(self) )->virtualbase_FocusOutEvent(event);
}

void QWidget_override_virtual_EnterEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__EnterEvent = slot;
}

void QWidget_virtualbase_EnterEvent(void* self, QEnterEvent* event) {
//	( (QWidget*)(self) )->virtualbase_EnterEvent(event);
}

void QWidget_override_virtual_LeaveEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__LeaveEvent = slot;
}

void QWidget_virtualbase_LeaveEvent(void* self, QEvent* event) {
//	( (QWidget*)(self) )->virtualbase_LeaveEvent(event);
}

void QWidget_override_virtual_PaintEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__PaintEvent = slot;
}

void QWidget_paintEvent(void* self, QPaintEvent* event) {
//	( (QWidget*)(self) )->paintEvent(event);
}

void QWidget_override_virtual_MoveEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__MoveEvent = slot;
}

void QWidget_virtualbase_MoveEvent(void* self, QMoveEvent* event) {
//	( (QWidget*)(self) )->virtualbase_MoveEvent(event);
}

void QWidget_override_virtual_ResizeEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__ResizeEvent = slot;
}

void QWidget_resizeEvent(void* self, QResizeEvent* event) {
//	( (QWidget*)(self) )->resizeEvent(event);
}

void QWidget_override_virtual_CloseEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__CloseEvent = slot;
}

void QWidget_virtualbase_CloseEvent(void* self, QCloseEvent* event) {
//	( (QWidget*)(self) )->virtualbase_CloseEvent(event);
}

void QWidget_override_virtual_ContextMenuEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__ContextMenuEvent = slot;
}

void QWidget_virtualbase_ContextMenuEvent(void* self, QContextMenuEvent* event) {
//	( (QWidget*)(self) )->virtualbase_ContextMenuEvent(event);
}

void QWidget_override_virtual_TabletEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__TabletEvent = slot;
}

void QWidget_virtualbase_TabletEvent(void* self, QTabletEvent* event) {
//	( (QWidget*)(self) )->virtualbase_TabletEvent(event);
}

void QWidget_override_virtual_ActionEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__ActionEvent = slot;
}

void QWidget_virtualbase_ActionEvent(void* self, QActionEvent* event) {
//	( (QWidget*)(self) )->virtualbase_ActionEvent(event);
}

void QWidget_override_virtual_DragEnterEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__DragEnterEvent = slot;
}

void QWidget_virtualbase_DragEnterEvent(void* self, QDragEnterEvent* event) {
//	( (QWidget*)(self) )->virtualbase_DragEnterEvent(event);
}

void QWidget_override_virtual_DragMoveEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__DragMoveEvent = slot;
}

void QWidget_virtualbase_DragMoveEvent(void* self, QDragMoveEvent* event) {
//	( (QWidget*)(self) )->virtualbase_DragMoveEvent(event);
}

void QWidget_override_virtual_DragLeaveEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__DragLeaveEvent = slot;
}

void QWidget_virtualbase_DragLeaveEvent(void* self, QDragLeaveEvent* event) {
//	( (QWidget*)(self) )->virtualbase_DragLeaveEvent(event);
}

void QWidget_override_virtual_DropEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__DropEvent = slot;
}

void QWidget_virtualbase_DropEvent(void* self, QDropEvent* event) {
//	( (QWidget*)(self) )->virtualbase_DropEvent(event);
}

void QWidget_override_virtual_ShowEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__ShowEvent = slot;
}

void QWidget_virtualbase_ShowEvent(void* self, QShowEvent* event) {
//	( (QWidget*)(self) )->virtualbase_ShowEvent(event);
}

void QWidget_override_virtual_HideEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__HideEvent = slot;
}

void QWidget_hideEvent(void* self, QHideEvent* event) {
//	( (QWidget*)(self) )->hideEvent(event);
}

void QWidget_override_virtual_NativeEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__NativeEvent = slot;
}

bool QWidget_nativeEvent(void* self, struct miqt_string eventType, void* message, intptr_t* result) {
//	return ( (QWidget*)(self) )->nativeEvent(eventType, message, result);
return false;//todo
}

void QWidget_override_virtual_ChangeEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__ChangeEvent = slot;
}

void QWidget_virtualbase_ChangeEvent(void* self, QEvent* param1) {
//	( (QWidget*)(self) )->virtualbase_ChangeEvent(param1);
}

void QWidget_override_virtual_Metric(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__Metric = slot;
}

//int QWidget_virtualbase_Metric(const void* self, PaintDeviceMetric param1) {
//	return ( (const QWidget*)(self) )->virtualbase_Metric(param1);
//}

void QWidget_override_virtual_InitPainter(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__InitPainter = slot;
}

void QWidget_virtualbase_InitPainter(const void* self, QPainter* painter) {
//	( (const QWidget*)(self) )->virtualbase_InitPainter(painter);
}

void QWidget_override_virtual_Redirected(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__Redirected = slot;
}

QPaintDevice* QWidget_redirected(const void* self, QPoint* offset) {
//	return ( (const QWidget*)(self) )->redirected(offset);
return nullptr;
}

void QWidget_override_virtual_SharedPainter(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__SharedPainter = slot;
}

QPainter* QWidget_virtualbase_SharedPainter(const void* self) {
//	return ( (const QWidget*)(self) )->virtualbase_SharedPainter();
return nullptr;
}

void QWidget_override_virtual_InputMethodEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__InputMethodEvent = slot;
}

void QWidget_virtualbase_InputMethodEvent(void* self, QInputMethodEvent* param1) {
//	( (QWidget*)(self) )->virtualbase_InputMethodEvent(param1);
}

void QWidget_override_virtual_InputMethodQuery(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__InputMethodQuery = slot;
}

QVariant* QWidget_virtualbase_InputMethodQuery(const void* self, int param1) {
//	return ( (const QWidget*)(self) )->virtualbase_InputMethodQuery(param1);
return nullptr;
}

void QWidget_override_virtual_FocusNextPrevChild(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__FocusNextPrevChild = slot;
}

bool QWidget_virtualbase_FocusNextPrevChild(void* self, bool next) {
//	return ( (QWidget*)(self) )->virtualbase_FocusNextPrevChild(next);
    return false;
}

void QWidget_override_virtual_EventFilter(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__EventFilter = slot;
}

bool QWidget_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event) {
//	return ( (QWidget*)(self) )->virtualbase_EventFilter(watched, event);
return false;
}

void QWidget_override_virtual_TimerEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__TimerEvent = slot;
}

void QWidget_virtualbase_TimerEvent(void* self, QTimerEvent* event) {
//	( (QWidget*)(self) )->virtualbase_TimerEvent(event);
}

void QWidget_override_virtual_ChildEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__ChildEvent = slot;
}

void QWidget_virtualbase_ChildEvent(void* self, QChildEvent* event) {
//	( (QWidget*)(self) )->virtualbase_ChildEvent(event);
}

void QWidget_override_virtual_CustomEvent(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__CustomEvent = slot;
}

void QWidget_virtualbase_CustomEvent(void* self, QEvent* event) {
//	( (QWidget*)(self) )->virtualbase_CustomEvent(event);
}

void QWidget_override_virtual_ConnectNotify(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__ConnectNotify = slot;
}

void QWidget_virtualbase_ConnectNotify(void* self, QMetaMethod* signal) {
//	( (QWidget*)(self) )->virtualbase_ConnectNotify(signal);
}

void QWidget_override_virtual_DisconnectNotify(void* self, intptr_t slot) {
//	dynamic_cast<QWidget*>( (QWidget*)(self) )->handle__DisconnectNotify = slot;
}

void QWidget_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal) {
//	( (QWidget*)(self) )->virtualbase_DisconnectNotify(signal);
}

void QWidget_Delete(QWidget* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QWidget*>( self );
	} else {
		delete self;
	}
}
