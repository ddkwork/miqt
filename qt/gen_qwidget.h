#pragma once
#ifndef MIQT_QT_GEN_QWIDGET_H
#define MIQT_QT_GEN_QWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAction;
class QActionEvent;
class QBackingStore;
class QBitmap;
class QChildEvent;
class QCloseEvent;
class QContextMenuEvent;
class QCursor;
class QDragEnterEvent;
class QDragLeaveEvent;
class QDragMoveEvent;
class QDropEvent;
class QEvent;
class QFocusEvent;
class QFont;
class QFontInfo;
class QFontMetrics;
class QGraphicsEffect;
class QGraphicsProxyWidget;
class QHideEvent;
class QIcon;
class QInputMethodEvent;
class QKeyEvent;
class QKeySequence;
class QLayout;
class QLocale;
class QMargins;
class QMetaMethod;
class QMetaObject;
class QMouseEvent;
class QMoveEvent;
class QObject;
class QPaintDevice;
class QPaintEngine;
class QPaintEvent;
class QPainter;
class QPalette;
class QPixmap;
class QPoint;
class QRect;
class QRegion;
class QResizeEvent;
class QScreen;
class QShowEvent;
class QSize;
class QSizePolicy;
class QStyle;
class QTabletEvent;
class QTimerEvent;
class QVariant;
class QWheelEvent;
class QWidget;
class QWidgetData;
class QWindow;
#else
typedef struct QAction QAction;
typedef struct QActionEvent QActionEvent;
typedef struct QBackingStore QBackingStore;
typedef struct QBitmap QBitmap;
typedef struct QChildEvent QChildEvent;
typedef struct QCloseEvent QCloseEvent;
typedef struct QContextMenuEvent QContextMenuEvent;
typedef struct QCursor QCursor;
typedef struct QDragEnterEvent QDragEnterEvent;
typedef struct QDragLeaveEvent QDragLeaveEvent;
typedef struct QDragMoveEvent QDragMoveEvent;
typedef struct QDropEvent QDropEvent;
typedef struct QEvent QEvent;
typedef struct QFocusEvent QFocusEvent;
typedef struct QFont QFont;
typedef struct QFontInfo QFontInfo;
typedef struct QFontMetrics QFontMetrics;
typedef struct QGraphicsEffect QGraphicsEffect;
typedef struct QGraphicsProxyWidget QGraphicsProxyWidget;
typedef struct QHideEvent QHideEvent;
typedef struct QIcon QIcon;
typedef struct QInputMethodEvent QInputMethodEvent;
typedef struct QKeyEvent QKeyEvent;
typedef struct QKeySequence QKeySequence;
typedef struct QLayout QLayout;
typedef struct QLocale QLocale;
typedef struct QMargins QMargins;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QMoveEvent QMoveEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPaintEvent QPaintEvent;
typedef struct QPainter QPainter;
typedef struct QPalette QPalette;
typedef struct QPixmap QPixmap;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QRegion QRegion;
typedef struct QResizeEvent QResizeEvent;
typedef struct QScreen QScreen;
typedef struct QShowEvent QShowEvent;
typedef struct QSize QSize;
typedef struct QSizePolicy QSizePolicy;
typedef struct QStyle QStyle;
typedef struct QTabletEvent QTabletEvent;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;
typedef struct QWheelEvent QWheelEvent;
typedef struct QWidget QWidget;
typedef struct QWidgetData QWidgetData;
typedef struct QWindow QWindow;
#endif

QWidgetData* QWidgetData_new(QWidgetData* param1);
void QWidgetData_operatorAssign(QWidgetData* self, QWidgetData* param1);
void QWidgetData_delete(QWidgetData* self);

QWidget* QWidget_new(QWidget* parent);
QWidget* QWidget_new2();
QWidget* QWidget_new3(QWidget* parent, int f);
void QWidget_virtbase(QWidget* src, QObject** outptr_QObject, QPaintDevice** outptr_QPaintDevice);
QMetaObject* QWidget_metaObject(const QWidget* self);
void* QWidget_metacast(QWidget* self, const char* param1);
struct miqt_string QWidget_tr(const char* s);
struct miqt_string QWidget_trUtf8(const char* s);
int QWidget_devType(const QWidget* self);
uintptr_t QWidget_winId(const QWidget* self);
void QWidget_createWinId(QWidget* self);
uintptr_t QWidget_internalWinId(const QWidget* self);
uintptr_t QWidget_effectiveWinId(const QWidget* self);
QStyle* QWidget_style(const QWidget* self);
void QWidget_setStyle(QWidget* self, QStyle* style);
bool QWidget_isTopLevel(const QWidget* self);
bool QWidget_isWindow(const QWidget* self);
bool QWidget_isModal(const QWidget* self);
int QWidget_windowModality(const QWidget* self);
void QWidget_setWindowModality(QWidget* self, int windowModality);
bool QWidget_isEnabled(const QWidget* self);
bool QWidget_isEnabledTo(const QWidget* self, QWidget* param1);
bool QWidget_isEnabledToTLW(const QWidget* self);
void QWidget_setEnabled(QWidget* self, bool enabled);
void QWidget_setDisabled(QWidget* self, bool disabled);
void QWidget_setWindowModified(QWidget* self, bool windowModified);
QRect* QWidget_frameGeometry(const QWidget* self);
QRect* QWidget_geometry(const QWidget* self);
QRect* QWidget_normalGeometry(const QWidget* self);
int QWidget_x(const QWidget* self);
int QWidget_y(const QWidget* self);
QPoint* QWidget_pos(const QWidget* self);
QSize* QWidget_frameSize(const QWidget* self);
QSize* QWidget_size(const QWidget* self);
int QWidget_width(const QWidget* self);
int QWidget_height(const QWidget* self);
QRect* QWidget_rect(const QWidget* self);
QRect* QWidget_childrenRect(const QWidget* self);
QRegion* QWidget_childrenRegion(const QWidget* self);
QSize* QWidget_minimumSize(const QWidget* self);
QSize* QWidget_maximumSize(const QWidget* self);
int QWidget_minimumWidth(const QWidget* self);
int QWidget_minimumHeight(const QWidget* self);
int QWidget_maximumWidth(const QWidget* self);
int QWidget_maximumHeight(const QWidget* self);
void QWidget_setMinimumSize(QWidget* self, QSize* minimumSize);
void QWidget_setMinimumSize2(QWidget* self, int minw, int minh);
void QWidget_setMaximumSize(QWidget* self, QSize* maximumSize);
void QWidget_setMaximumSize2(QWidget* self, int maxw, int maxh);
void QWidget_setMinimumWidth(QWidget* self, int minw);
void QWidget_setMinimumHeight(QWidget* self, int minh);
void QWidget_setMaximumWidth(QWidget* self, int maxw);
void QWidget_setMaximumHeight(QWidget* self, int maxh);
QSize* QWidget_sizeIncrement(const QWidget* self);
void QWidget_setSizeIncrement(QWidget* self, QSize* sizeIncrement);
void QWidget_setSizeIncrement2(QWidget* self, int w, int h);
QSize* QWidget_baseSize(const QWidget* self);
void QWidget_setBaseSize(QWidget* self, QSize* baseSize);
void QWidget_setBaseSize2(QWidget* self, int basew, int baseh);
void QWidget_setFixedSize(QWidget* self, QSize* fixedSize);
void QWidget_setFixedSize2(QWidget* self, int w, int h);
void QWidget_setFixedWidth(QWidget* self, int w);
void QWidget_setFixedHeight(QWidget* self, int h);
QPoint* QWidget_mapToGlobal(const QWidget* self, QPoint* param1);
QPoint* QWidget_mapFromGlobal(const QWidget* self, QPoint* param1);
QPoint* QWidget_mapToParent(const QWidget* self, QPoint* param1);
QPoint* QWidget_mapFromParent(const QWidget* self, QPoint* param1);
QPoint* QWidget_mapTo(const QWidget* self, QWidget* param1, QPoint* param2);
QPoint* QWidget_mapFrom(const QWidget* self, QWidget* param1, QPoint* param2);
QWidget* QWidget_window(const QWidget* self);
QWidget* QWidget_nativeParentWidget(const QWidget* self);
QWidget* QWidget_topLevelWidget(const QWidget* self);
QPalette* QWidget_palette(const QWidget* self);
void QWidget_setPalette(QWidget* self, QPalette* palette);
void QWidget_setBackgroundRole(QWidget* self, int backgroundRole);
int QWidget_backgroundRole(const QWidget* self);
void QWidget_setForegroundRole(QWidget* self, int foregroundRole);
int QWidget_foregroundRole(const QWidget* self);
QFont* QWidget_font(const QWidget* self);
void QWidget_setFont(QWidget* self, QFont* font);
QFontMetrics* QWidget_fontMetrics(const QWidget* self);
QFontInfo* QWidget_fontInfo(const QWidget* self);
QCursor* QWidget_cursor(const QWidget* self);
void QWidget_setCursor(QWidget* self, QCursor* cursor);
void QWidget_unsetCursor(QWidget* self);
void QWidget_setMouseTracking(QWidget* self, bool enable);
bool QWidget_hasMouseTracking(const QWidget* self);
bool QWidget_underMouse(const QWidget* self);
void QWidget_setTabletTracking(QWidget* self, bool enable);
bool QWidget_hasTabletTracking(const QWidget* self);
void QWidget_setMask(QWidget* self, QBitmap* mask);
void QWidget_setMaskWithMask(QWidget* self, QRegion* mask);
QRegion* QWidget_mask(const QWidget* self);
void QWidget_clearMask(QWidget* self);
void QWidget_render(QWidget* self, QPaintDevice* target);
void QWidget_renderWithPainter(QWidget* self, QPainter* painter);
QPixmap* QWidget_grab(QWidget* self);
QGraphicsEffect* QWidget_graphicsEffect(const QWidget* self);
void QWidget_setGraphicsEffect(QWidget* self, QGraphicsEffect* effect);
void QWidget_grabGesture(QWidget* self, int type);
void QWidget_ungrabGesture(QWidget* self, int type);
void QWidget_setWindowTitle(QWidget* self, struct miqt_string windowTitle);
void QWidget_setStyleSheet(QWidget* self, struct miqt_string styleSheet);
struct miqt_string QWidget_styleSheet(const QWidget* self);
struct miqt_string QWidget_windowTitle(const QWidget* self);
void QWidget_setWindowIcon(QWidget* self, QIcon* icon);
QIcon* QWidget_windowIcon(const QWidget* self);
void QWidget_setWindowIconText(QWidget* self, struct miqt_string windowIconText);
struct miqt_string QWidget_windowIconText(const QWidget* self);
void QWidget_setWindowRole(QWidget* self, struct miqt_string windowRole);
struct miqt_string QWidget_windowRole(const QWidget* self);
void QWidget_setWindowFilePath(QWidget* self, struct miqt_string filePath);
struct miqt_string QWidget_windowFilePath(const QWidget* self);
void QWidget_setWindowOpacity(QWidget* self, double level);
double QWidget_windowOpacity(const QWidget* self);
bool QWidget_isWindowModified(const QWidget* self);
void QWidget_setToolTip(QWidget* self, struct miqt_string toolTip);
struct miqt_string QWidget_toolTip(const QWidget* self);
void QWidget_setToolTipDuration(QWidget* self, int msec);
int QWidget_toolTipDuration(const QWidget* self);
void QWidget_setStatusTip(QWidget* self, struct miqt_string statusTip);
struct miqt_string QWidget_statusTip(const QWidget* self);
void QWidget_setWhatsThis(QWidget* self, struct miqt_string whatsThis);
struct miqt_string QWidget_whatsThis(const QWidget* self);
struct miqt_string QWidget_accessibleName(const QWidget* self);
void QWidget_setAccessibleName(QWidget* self, struct miqt_string name);
struct miqt_string QWidget_accessibleDescription(const QWidget* self);
void QWidget_setAccessibleDescription(QWidget* self, struct miqt_string description);
void QWidget_setLayoutDirection(QWidget* self, int direction);
int QWidget_layoutDirection(const QWidget* self);
void QWidget_unsetLayoutDirection(QWidget* self);
void QWidget_setLocale(QWidget* self, QLocale* locale);
QLocale* QWidget_locale(const QWidget* self);
void QWidget_unsetLocale(QWidget* self);
bool QWidget_isRightToLeft(const QWidget* self);
bool QWidget_isLeftToRight(const QWidget* self);
void QWidget_setFocus(QWidget* self);
bool QWidget_isActiveWindow(const QWidget* self);
void QWidget_activateWindow(QWidget* self);
void QWidget_clearFocus(QWidget* self);
void QWidget_setFocusWithReason(QWidget* self, int reason);
int QWidget_focusPolicy(const QWidget* self);
void QWidget_setFocusPolicy(QWidget* self, int policy);
bool QWidget_hasFocus(const QWidget* self);
void QWidget_setTabOrder(QWidget* param1, QWidget* param2);
void QWidget_setFocusProxy(QWidget* self, QWidget* focusProxy);
QWidget* QWidget_focusProxy(const QWidget* self);
int QWidget_contextMenuPolicy(const QWidget* self);
void QWidget_setContextMenuPolicy(QWidget* self, int policy);
void QWidget_grabMouse(QWidget* self);
void QWidget_grabMouseWithQCursor(QWidget* self, QCursor* param1);
void QWidget_releaseMouse(QWidget* self);
void QWidget_grabKeyboard(QWidget* self);
void QWidget_releaseKeyboard(QWidget* self);
int QWidget_grabShortcut(QWidget* self, QKeySequence* key);
void QWidget_releaseShortcut(QWidget* self, int id);
void QWidget_setShortcutEnabled(QWidget* self, int id);
void QWidget_setShortcutAutoRepeat(QWidget* self, int id);
QWidget* QWidget_mouseGrabber();
QWidget* QWidget_keyboardGrabber();
bool QWidget_updatesEnabled(const QWidget* self);
void QWidget_setUpdatesEnabled(QWidget* self, bool enable);
QGraphicsProxyWidget* QWidget_graphicsProxyWidget(const QWidget* self);
void QWidget_update(QWidget* self);
void QWidget_repaint(QWidget* self);
void QWidget_update2(QWidget* self, int x, int y, int w, int h);
void QWidget_updateWithQRect(QWidget* self, QRect* param1);
void QWidget_updateWithQRegion(QWidget* self, QRegion* param1);
void QWidget_repaint2(QWidget* self, int x, int y, int w, int h);
void QWidget_repaintWithQRect(QWidget* self, QRect* param1);
void QWidget_repaintWithQRegion(QWidget* self, QRegion* param1);
void QWidget_setVisible(QWidget* self, bool visible);
void QWidget_setHidden(QWidget* self, bool hidden);
void QWidget_show(QWidget* self);
void QWidget_hide(QWidget* self);
void QWidget_showMinimized(QWidget* self);
void QWidget_showMaximized(QWidget* self);
void QWidget_showFullScreen(QWidget* self);
void QWidget_showNormal(QWidget* self);
bool QWidget_close(QWidget* self);
void QWidget_raise(QWidget* self);
void QWidget_lower(QWidget* self);
void QWidget_stackUnder(QWidget* self, QWidget* param1);
void QWidget_move(QWidget* self, int x, int y);
void QWidget_moveWithQPoint(QWidget* self, QPoint* param1);
void QWidget_resize(QWidget* self, int w, int h);
void QWidget_resizeWithQSize(QWidget* self, QSize* param1);
void QWidget_setGeometry(QWidget* self, int x, int y, int w, int h);
void QWidget_setGeometryWithGeometry(QWidget* self, QRect* geometry);
struct miqt_string QWidget_saveGeometry(const QWidget* self);
bool QWidget_restoreGeometry(QWidget* self, struct miqt_string geometry);
void QWidget_adjustSize(QWidget* self);
bool QWidget_isVisible(const QWidget* self);
bool QWidget_isVisibleTo(const QWidget* self, QWidget* param1);
bool QWidget_isHidden(const QWidget* self);
bool QWidget_isMinimized(const QWidget* self);
bool QWidget_isMaximized(const QWidget* self);
bool QWidget_isFullScreen(const QWidget* self);
int QWidget_windowState(const QWidget* self);
void QWidget_setWindowState(QWidget* self, int state);
void QWidget_overrideWindowState(QWidget* self, int state);
QSize* QWidget_sizeHint(const QWidget* self);
QSize* QWidget_minimumSizeHint(const QWidget* self);
QSizePolicy* QWidget_sizePolicy(const QWidget* self);
void QWidget_setSizePolicy(QWidget* self, QSizePolicy* sizePolicy);
void QWidget_setSizePolicy2(QWidget* self, int horizontal, int vertical);
int QWidget_heightForWidth(const QWidget* self, int param1);
bool QWidget_hasHeightForWidth(const QWidget* self);
QRegion* QWidget_visibleRegion(const QWidget* self);
void QWidget_setContentsMargins(QWidget* self, int left, int top, int right, int bottom);
void QWidget_setContentsMarginsWithMargins(QWidget* self, QMargins* margins);
void QWidget_getContentsMargins(const QWidget* self, int* left, int* top, int* right, int* bottom);
QMargins* QWidget_contentsMargins(const QWidget* self);
QRect* QWidget_contentsRect(const QWidget* self);
QLayout* QWidget_layout(const QWidget* self);
void QWidget_setLayout(QWidget* self, QLayout* layout);
void QWidget_updateGeometry(QWidget* self);
void QWidget_setParent(QWidget* self, QWidget* parent);
void QWidget_setParent2(QWidget* self, QWidget* parent, int f);
void QWidget_scroll(QWidget* self, int dx, int dy);
void QWidget_scroll2(QWidget* self, int dx, int dy, QRect* param3);
QWidget* QWidget_focusWidget(const QWidget* self);
QWidget* QWidget_nextInFocusChain(const QWidget* self);
QWidget* QWidget_previousInFocusChain(const QWidget* self);
bool QWidget_acceptDrops(const QWidget* self);
void QWidget_setAcceptDrops(QWidget* self, bool on);
void QWidget_addAction(QWidget* self, QAction* action);
void QWidget_addActions(QWidget* self, struct miqt_array /* of QAction* */  actions);
void QWidget_insertActions(QWidget* self, QAction* before, struct miqt_array /* of QAction* */  actions);
void QWidget_insertAction(QWidget* self, QAction* before, QAction* action);
void QWidget_removeAction(QWidget* self, QAction* action);
struct miqt_array /* of QAction* */  QWidget_actions(const QWidget* self);
QWidget* QWidget_parentWidget(const QWidget* self);
void QWidget_setWindowFlags(QWidget* self, int type);
int QWidget_windowFlags(const QWidget* self);
void QWidget_setWindowFlag(QWidget* self, int param1);
void QWidget_overrideWindowFlags(QWidget* self, int type);
int QWidget_windowType(const QWidget* self);
QWidget* QWidget_find(uintptr_t param1);
QWidget* QWidget_childAt(const QWidget* self, int x, int y);
QWidget* QWidget_childAtWithQPoint(const QWidget* self, QPoint* p);
void QWidget_setAttribute(QWidget* self, int param1);
bool QWidget_testAttribute(const QWidget* self, int param1);
QPaintEngine* QWidget_paintEngine(const QWidget* self);
void QWidget_ensurePolished(const QWidget* self);
bool QWidget_isAncestorOf(const QWidget* self, QWidget* child);
bool QWidget_autoFillBackground(const QWidget* self);
void QWidget_setAutoFillBackground(QWidget* self, bool enabled);
QBackingStore* QWidget_backingStore(const QWidget* self);
QWindow* QWidget_windowHandle(const QWidget* self);
QScreen* QWidget_screen(const QWidget* self);
QWidget* QWidget_createWindowContainer(QWindow* window);
void QWidget_windowTitleChanged(QWidget* self, struct miqt_string title);
void QWidget_connect_windowTitleChanged(QWidget* self, intptr_t slot);
void QWidget_windowIconChanged(QWidget* self, QIcon* icon);
void QWidget_connect_windowIconChanged(QWidget* self, intptr_t slot);
void QWidget_windowIconTextChanged(QWidget* self, struct miqt_string iconText);
void QWidget_connect_windowIconTextChanged(QWidget* self, intptr_t slot);
void QWidget_customContextMenuRequested(QWidget* self, QPoint* pos);
void QWidget_connect_customContextMenuRequested(QWidget* self, intptr_t slot);
bool QWidget_event(QWidget* self, QEvent* event);
void QWidget_mousePressEvent(QWidget* self, QMouseEvent* event);
void QWidget_mouseReleaseEvent(QWidget* self, QMouseEvent* event);
void QWidget_mouseDoubleClickEvent(QWidget* self, QMouseEvent* event);
void QWidget_mouseMoveEvent(QWidget* self, QMouseEvent* event);
void QWidget_wheelEvent(QWidget* self, QWheelEvent* event);
void QWidget_keyPressEvent(QWidget* self, QKeyEvent* event);
void QWidget_keyReleaseEvent(QWidget* self, QKeyEvent* event);
void QWidget_focusInEvent(QWidget* self, QFocusEvent* event);
void QWidget_focusOutEvent(QWidget* self, QFocusEvent* event);
void QWidget_enterEvent(QWidget* self, QEvent* event);
void QWidget_leaveEvent(QWidget* self, QEvent* event);
void QWidget_paintEvent(QWidget* self, QPaintEvent* event);
void QWidget_moveEvent(QWidget* self, QMoveEvent* event);
void QWidget_resizeEvent(QWidget* self, QResizeEvent* event);
void QWidget_closeEvent(QWidget* self, QCloseEvent* event);
void QWidget_contextMenuEvent(QWidget* self, QContextMenuEvent* event);
void QWidget_tabletEvent(QWidget* self, QTabletEvent* event);
void QWidget_actionEvent(QWidget* self, QActionEvent* event);
void QWidget_dragEnterEvent(QWidget* self, QDragEnterEvent* event);
void QWidget_dragMoveEvent(QWidget* self, QDragMoveEvent* event);
void QWidget_dragLeaveEvent(QWidget* self, QDragLeaveEvent* event);
void QWidget_dropEvent(QWidget* self, QDropEvent* event);
void QWidget_showEvent(QWidget* self, QShowEvent* event);
void QWidget_hideEvent(QWidget* self, QHideEvent* event);
bool QWidget_nativeEvent(QWidget* self, struct miqt_string eventType, void* message, long* result);
void QWidget_changeEvent(QWidget* self, QEvent* param1);
int QWidget_metric(const QWidget* self, int param1);
void QWidget_initPainter(const QWidget* self, QPainter* painter);
QPaintDevice* QWidget_redirected(const QWidget* self, QPoint* offset);
QPainter* QWidget_sharedPainter(const QWidget* self);
void QWidget_inputMethodEvent(QWidget* self, QInputMethodEvent* param1);
QVariant* QWidget_inputMethodQuery(const QWidget* self, int param1);
int QWidget_inputMethodHints(const QWidget* self);
void QWidget_setInputMethodHints(QWidget* self, int hints);
bool QWidget_focusNextPrevChild(QWidget* self, bool next);
struct miqt_string QWidget_tr2(const char* s, const char* c);
struct miqt_string QWidget_tr3(const char* s, const char* c, int n);
struct miqt_string QWidget_trUtf82(const char* s, const char* c);
struct miqt_string QWidget_trUtf83(const char* s, const char* c, int n);
void QWidget_render2(QWidget* self, QPaintDevice* target, QPoint* targetOffset);
void QWidget_render3(QWidget* self, QPaintDevice* target, QPoint* targetOffset, QRegion* sourceRegion);
void QWidget_render4(QWidget* self, QPaintDevice* target, QPoint* targetOffset, QRegion* sourceRegion, int renderFlags);
void QWidget_render22(QWidget* self, QPainter* painter, QPoint* targetOffset);
void QWidget_render32(QWidget* self, QPainter* painter, QPoint* targetOffset, QRegion* sourceRegion);
void QWidget_render42(QWidget* self, QPainter* painter, QPoint* targetOffset, QRegion* sourceRegion, int renderFlags);
QPixmap* QWidget_grab1(QWidget* self, QRect* rectangle);
void QWidget_grabGesture2(QWidget* self, int type, int flags);
int QWidget_grabShortcut2(QWidget* self, QKeySequence* key, int context);
void QWidget_setShortcutEnabled2(QWidget* self, int id, bool enable);
void QWidget_setShortcutAutoRepeat2(QWidget* self, int id, bool enable);
void QWidget_setWindowFlag2(QWidget* self, int param1, bool on);
void QWidget_setAttribute2(QWidget* self, int param1, bool on);
QWidget* QWidget_createWindowContainer2(QWindow* window, QWidget* parent);
QWidget* QWidget_createWindowContainer3(QWindow* window, QWidget* parent, int flags);
bool QWidget_override_virtual_devType(void* self, intptr_t slot);
int QWidget_virtualbase_devType(const void* self);
bool QWidget_override_virtual_setVisible(void* self, intptr_t slot);
void QWidget_virtualbase_setVisible(void* self, bool visible);
bool QWidget_override_virtual_sizeHint(void* self, intptr_t slot);
QSize* QWidget_virtualbase_sizeHint(const void* self);
bool QWidget_override_virtual_minimumSizeHint(void* self, intptr_t slot);
QSize* QWidget_virtualbase_minimumSizeHint(const void* self);
bool QWidget_override_virtual_heightForWidth(void* self, intptr_t slot);
int QWidget_virtualbase_heightForWidth(const void* self, int param1);
bool QWidget_override_virtual_hasHeightForWidth(void* self, intptr_t slot);
bool QWidget_virtualbase_hasHeightForWidth(const void* self);
bool QWidget_override_virtual_paintEngine(void* self, intptr_t slot);
QPaintEngine* QWidget_virtualbase_paintEngine(const void* self);
bool QWidget_override_virtual_event(void* self, intptr_t slot);
bool QWidget_virtualbase_event(void* self, QEvent* event);
bool QWidget_override_virtual_mousePressEvent(void* self, intptr_t slot);
void QWidget_virtualbase_mousePressEvent(void* self, QMouseEvent* event);
bool QWidget_override_virtual_mouseReleaseEvent(void* self, intptr_t slot);
void QWidget_virtualbase_mouseReleaseEvent(void* self, QMouseEvent* event);
bool QWidget_override_virtual_mouseDoubleClickEvent(void* self, intptr_t slot);
void QWidget_virtualbase_mouseDoubleClickEvent(void* self, QMouseEvent* event);
bool QWidget_override_virtual_mouseMoveEvent(void* self, intptr_t slot);
void QWidget_virtualbase_mouseMoveEvent(void* self, QMouseEvent* event);
bool QWidget_override_virtual_wheelEvent(void* self, intptr_t slot);
void QWidget_virtualbase_wheelEvent(void* self, QWheelEvent* event);
bool QWidget_override_virtual_keyPressEvent(void* self, intptr_t slot);
void QWidget_virtualbase_keyPressEvent(void* self, QKeyEvent* event);
bool QWidget_override_virtual_keyReleaseEvent(void* self, intptr_t slot);
void QWidget_virtualbase_keyReleaseEvent(void* self, QKeyEvent* event);
bool QWidget_override_virtual_focusInEvent(void* self, intptr_t slot);
void QWidget_virtualbase_focusInEvent(void* self, QFocusEvent* event);
bool QWidget_override_virtual_focusOutEvent(void* self, intptr_t slot);
void QWidget_virtualbase_focusOutEvent(void* self, QFocusEvent* event);
bool QWidget_override_virtual_enterEvent(void* self, intptr_t slot);
void QWidget_virtualbase_enterEvent(void* self, QEvent* event);
bool QWidget_override_virtual_leaveEvent(void* self, intptr_t slot);
void QWidget_virtualbase_leaveEvent(void* self, QEvent* event);
bool QWidget_override_virtual_paintEvent(void* self, intptr_t slot);
void QWidget_virtualbase_paintEvent(void* self, QPaintEvent* event);
bool QWidget_override_virtual_moveEvent(void* self, intptr_t slot);
void QWidget_virtualbase_moveEvent(void* self, QMoveEvent* event);
bool QWidget_override_virtual_resizeEvent(void* self, intptr_t slot);
void QWidget_virtualbase_resizeEvent(void* self, QResizeEvent* event);
bool QWidget_override_virtual_closeEvent(void* self, intptr_t slot);
void QWidget_virtualbase_closeEvent(void* self, QCloseEvent* event);
bool QWidget_override_virtual_contextMenuEvent(void* self, intptr_t slot);
void QWidget_virtualbase_contextMenuEvent(void* self, QContextMenuEvent* event);
bool QWidget_override_virtual_tabletEvent(void* self, intptr_t slot);
void QWidget_virtualbase_tabletEvent(void* self, QTabletEvent* event);
bool QWidget_override_virtual_actionEvent(void* self, intptr_t slot);
void QWidget_virtualbase_actionEvent(void* self, QActionEvent* event);
bool QWidget_override_virtual_dragEnterEvent(void* self, intptr_t slot);
void QWidget_virtualbase_dragEnterEvent(void* self, QDragEnterEvent* event);
bool QWidget_override_virtual_dragMoveEvent(void* self, intptr_t slot);
void QWidget_virtualbase_dragMoveEvent(void* self, QDragMoveEvent* event);
bool QWidget_override_virtual_dragLeaveEvent(void* self, intptr_t slot);
void QWidget_virtualbase_dragLeaveEvent(void* self, QDragLeaveEvent* event);
bool QWidget_override_virtual_dropEvent(void* self, intptr_t slot);
void QWidget_virtualbase_dropEvent(void* self, QDropEvent* event);
bool QWidget_override_virtual_showEvent(void* self, intptr_t slot);
void QWidget_virtualbase_showEvent(void* self, QShowEvent* event);
bool QWidget_override_virtual_hideEvent(void* self, intptr_t slot);
void QWidget_virtualbase_hideEvent(void* self, QHideEvent* event);
bool QWidget_override_virtual_nativeEvent(void* self, intptr_t slot);
bool QWidget_virtualbase_nativeEvent(void* self, struct miqt_string eventType, void* message, long* result);
bool QWidget_override_virtual_changeEvent(void* self, intptr_t slot);
void QWidget_virtualbase_changeEvent(void* self, QEvent* param1);
bool QWidget_override_virtual_metric(void* self, intptr_t slot);
int QWidget_virtualbase_metric(const void* self, int param1);
bool QWidget_override_virtual_initPainter(void* self, intptr_t slot);
void QWidget_virtualbase_initPainter(const void* self, QPainter* painter);
bool QWidget_override_virtual_redirected(void* self, intptr_t slot);
QPaintDevice* QWidget_virtualbase_redirected(const void* self, QPoint* offset);
bool QWidget_override_virtual_sharedPainter(void* self, intptr_t slot);
QPainter* QWidget_virtualbase_sharedPainter(const void* self);
bool QWidget_override_virtual_inputMethodEvent(void* self, intptr_t slot);
void QWidget_virtualbase_inputMethodEvent(void* self, QInputMethodEvent* param1);
bool QWidget_override_virtual_inputMethodQuery(void* self, intptr_t slot);
QVariant* QWidget_virtualbase_inputMethodQuery(const void* self, int param1);
bool QWidget_override_virtual_focusNextPrevChild(void* self, intptr_t slot);
bool QWidget_virtualbase_focusNextPrevChild(void* self, bool next);
bool QWidget_override_virtual_eventFilter(void* self, intptr_t slot);
bool QWidget_virtualbase_eventFilter(void* self, QObject* watched, QEvent* event);
bool QWidget_override_virtual_timerEvent(void* self, intptr_t slot);
void QWidget_virtualbase_timerEvent(void* self, QTimerEvent* event);
bool QWidget_override_virtual_childEvent(void* self, intptr_t slot);
void QWidget_virtualbase_childEvent(void* self, QChildEvent* event);
bool QWidget_override_virtual_customEvent(void* self, intptr_t slot);
void QWidget_virtualbase_customEvent(void* self, QEvent* event);
bool QWidget_override_virtual_connectNotify(void* self, intptr_t slot);
void QWidget_virtualbase_connectNotify(void* self, QMetaMethod* signal);
bool QWidget_override_virtual_disconnectNotify(void* self, intptr_t slot);
void QWidget_virtualbase_disconnectNotify(void* self, QMetaMethod* signal);
void QWidget_protectedbase_updateMicroFocus(bool* _dynamic_cast_ok, void* self);
void QWidget_protectedbase_create(bool* _dynamic_cast_ok, void* self);
void QWidget_protectedbase_destroy(bool* _dynamic_cast_ok, void* self);
bool QWidget_protectedbase_focusNextChild(bool* _dynamic_cast_ok, void* self);
bool QWidget_protectedbase_focusPreviousChild(bool* _dynamic_cast_ok, void* self);
void QWidget_protectedbase_create1(bool* _dynamic_cast_ok, void* self, uintptr_t param1);
void QWidget_protectedbase_create2(bool* _dynamic_cast_ok, void* self, uintptr_t param1, bool initializeWindow);
void QWidget_protectedbase_create3(bool* _dynamic_cast_ok, void* self, uintptr_t param1, bool initializeWindow, bool destroyOldWindow);
void QWidget_protectedbase_destroy1(bool* _dynamic_cast_ok, void* self, bool destroyWindow);
void QWidget_protectedbase_destroy2(bool* _dynamic_cast_ok, void* self, bool destroyWindow, bool destroySubWindows);
QObject* QWidget_protectedbase_sender(bool* _dynamic_cast_ok, const void* self);
int QWidget_protectedbase_senderSignalIndex(bool* _dynamic_cast_ok, const void* self);
int QWidget_protectedbase_receivers(bool* _dynamic_cast_ok, const void* self, const char* signal);
bool QWidget_protectedbase_isSignalConnected(bool* _dynamic_cast_ok, const void* self, QMetaMethod* signal);
void QWidget_delete(QWidget* self);

#ifdef __cplusplus
} /* extern C */
#endif

#endif
