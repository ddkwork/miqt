#pragma once
#ifndef MIQT_QT_GEN_QSTACKEDWIDGET_H
#define MIQT_QT_GEN_QSTACKEDWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QEvent;
class QFrame;
class QMetaObject;
class QObject;
class QPaintDevice;
class QPaintEvent;
class QSize;
class QStackedWidget;
class QStyleOptionFrame;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QEvent QEvent;
typedef struct QFrame QFrame;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QSize QSize;
typedef struct QStackedWidget QStackedWidget;
typedef struct QStyleOptionFrame QStyleOptionFrame;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QStackedWidget* QStackedWidget_new(QWidget* parent);
extern __declspec(dllexport) QStackedWidget* QStackedWidget_new2();
extern __declspec(dllexport) void QStackedWidget_virtbase(QStackedWidget* src, QFrame** outptr_QFrame);
extern __declspec(dllexport) QMetaObject* QStackedWidget_MetaObject(const QStackedWidget* self);
extern __declspec(dllexport) void* QStackedWidget_Metacast(QStackedWidget* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QStackedWidget_Tr(const char* s);
extern __declspec(dllexport) int QStackedWidget_AddWidget(QStackedWidget* self, QWidget* w);
extern __declspec(dllexport) int QStackedWidget_InsertWidget(QStackedWidget* self, int index, QWidget* w);
extern __declspec(dllexport) void QStackedWidget_RemoveWidget(QStackedWidget* self, QWidget* w);
extern __declspec(dllexport) QWidget* QStackedWidget_CurrentWidget(const QStackedWidget* self);
extern __declspec(dllexport) int QStackedWidget_CurrentIndex(const QStackedWidget* self);
extern __declspec(dllexport) int QStackedWidget_IndexOf(const QStackedWidget* self, QWidget* param1);
extern __declspec(dllexport) QWidget* QStackedWidget_Widget(const QStackedWidget* self, int param1);
extern __declspec(dllexport) int QStackedWidget_Count(const QStackedWidget* self);
extern __declspec(dllexport) void QStackedWidget_SetCurrentIndex(QStackedWidget* self, int index);
extern __declspec(dllexport) void QStackedWidget_SetCurrentWidget(QStackedWidget* self, QWidget* w);
extern __declspec(dllexport) void QStackedWidget_CurrentChanged(QStackedWidget* self, int param1);
void QStackedWidget_connect_CurrentChanged(QStackedWidget* self, intptr_t slot);
extern __declspec(dllexport) void QStackedWidget_WidgetRemoved(QStackedWidget* self, int index);
void QStackedWidget_connect_WidgetRemoved(QStackedWidget* self, intptr_t slot);
extern __declspec(dllexport) void QStackedWidget_WidgetAdded(QStackedWidget* self, int index);
void QStackedWidget_connect_WidgetAdded(QStackedWidget* self, intptr_t slot);
extern __declspec(dllexport) bool QStackedWidget_Event(QStackedWidget* self, QEvent* e);
extern __declspec(dllexport) struct miqt_string QStackedWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QStackedWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QStackedWidget_override_virtual_Event(void* self, intptr_t slot);
bool QStackedWidget_virtualbase_Event(void* self, QEvent* e);
extern __declspec(dllexport) void QStackedWidget_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QStackedWidget_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QStackedWidget_override_virtual_PaintEvent(void* self, intptr_t slot);
void QStackedWidget_virtualbase_PaintEvent(void* self, QPaintEvent* param1);
extern __declspec(dllexport) void QStackedWidget_override_virtual_ChangeEvent(void* self, intptr_t slot);
void QStackedWidget_virtualbase_ChangeEvent(void* self, QEvent* param1);
extern __declspec(dllexport) void QStackedWidget_override_virtual_InitStyleOption(void* self, intptr_t slot);
void QStackedWidget_virtualbase_InitStyleOption(const void* self, QStyleOptionFrame* option);
extern __declspec(dllexport) void QStackedWidget_Delete(QStackedWidget* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
