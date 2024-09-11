#ifndef GEN_QTOOLBAR_H
#define GEN_QTOOLBAR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QAction;
class QIcon;
class QMetaObject;
class QPoint;
class QRect;
class QSize;
class QToolBar;
class QWidget;
#else
typedef struct QAction QAction;
typedef struct QIcon QIcon;
typedef struct QMetaObject QMetaObject;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QToolBar QToolBar;
typedef struct QWidget QWidget;
#endif

QToolBar* QToolBar_new(const char* title, size_t title_Strlen);
QToolBar* QToolBar_new2();
QToolBar* QToolBar_new3(const char* title, size_t title_Strlen, QWidget* parent);
QToolBar* QToolBar_new4(QWidget* parent);
QMetaObject* QToolBar_MetaObject(const QToolBar* self);
void QToolBar_Tr(const char* s, char** _out, int* _out_Strlen);
void QToolBar_TrUtf8(const char* s, char** _out, int* _out_Strlen);
void QToolBar_SetMovable(QToolBar* self, bool movable);
bool QToolBar_IsMovable(const QToolBar* self);
void QToolBar_SetAllowedAreas(QToolBar* self, int areas);
int QToolBar_AllowedAreas(const QToolBar* self);
bool QToolBar_IsAreaAllowed(const QToolBar* self, uintptr_t area);
void QToolBar_SetOrientation(QToolBar* self, uintptr_t orientation);
uintptr_t QToolBar_Orientation(const QToolBar* self);
void QToolBar_Clear(QToolBar* self);
QAction* QToolBar_AddAction(QToolBar* self, const char* text, size_t text_Strlen);
QAction* QToolBar_AddAction2(QToolBar* self, QIcon* icon, const char* text, size_t text_Strlen);
QAction* QToolBar_AddSeparator(QToolBar* self);
QAction* QToolBar_InsertSeparator(QToolBar* self, QAction* before);
QAction* QToolBar_AddWidget(QToolBar* self, QWidget* widget);
QAction* QToolBar_InsertWidget(QToolBar* self, QAction* before, QWidget* widget);
QRect* QToolBar_ActionGeometry(const QToolBar* self, QAction* action);
QAction* QToolBar_ActionAt(const QToolBar* self, QPoint* p);
QAction* QToolBar_ActionAt2(const QToolBar* self, int x, int y);
QAction* QToolBar_ToggleViewAction(const QToolBar* self);
QSize* QToolBar_IconSize(const QToolBar* self);
uintptr_t QToolBar_ToolButtonStyle(const QToolBar* self);
QWidget* QToolBar_WidgetForAction(const QToolBar* self, QAction* action);
bool QToolBar_IsFloatable(const QToolBar* self);
void QToolBar_SetFloatable(QToolBar* self, bool floatable);
bool QToolBar_IsFloating(const QToolBar* self);
void QToolBar_SetIconSize(QToolBar* self, QSize* iconSize);
void QToolBar_SetToolButtonStyle(QToolBar* self, uintptr_t toolButtonStyle);
void QToolBar_ActionTriggered(QToolBar* self, QAction* action);
void QToolBar_connect_ActionTriggered(QToolBar* self, void* slot);
void QToolBar_MovableChanged(QToolBar* self, bool movable);
void QToolBar_connect_MovableChanged(QToolBar* self, void* slot);
void QToolBar_AllowedAreasChanged(QToolBar* self, int allowedAreas);
void QToolBar_connect_AllowedAreasChanged(QToolBar* self, void* slot);
void QToolBar_OrientationChanged(QToolBar* self, uintptr_t orientation);
void QToolBar_connect_OrientationChanged(QToolBar* self, void* slot);
void QToolBar_IconSizeChanged(QToolBar* self, QSize* iconSize);
void QToolBar_connect_IconSizeChanged(QToolBar* self, void* slot);
void QToolBar_ToolButtonStyleChanged(QToolBar* self, uintptr_t toolButtonStyle);
void QToolBar_connect_ToolButtonStyleChanged(QToolBar* self, void* slot);
void QToolBar_TopLevelChanged(QToolBar* self, bool topLevel);
void QToolBar_connect_TopLevelChanged(QToolBar* self, void* slot);
void QToolBar_VisibilityChanged(QToolBar* self, bool visible);
void QToolBar_connect_VisibilityChanged(QToolBar* self, void* slot);
void QToolBar_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QToolBar_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QToolBar_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QToolBar_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QToolBar_Delete(QToolBar* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
