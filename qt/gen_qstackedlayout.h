#pragma once
#ifndef MIQT_QT_GEN_QSTACKEDLAYOUT_H
#define MIQT_QT_GEN_QSTACKEDLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChildEvent;
class QLayout;
class QLayoutItem;
class QMetaObject;
class QObject;
class QRect;
class QSize;
class QStackedLayout;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QChildEvent QChildEvent;
typedef struct QLayout QLayout;
typedef struct QLayoutItem QLayoutItem;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QStackedLayout QStackedLayout;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QStackedLayout* QStackedLayout_new(QWidget* parent);
extern __declspec(dllexport) QStackedLayout* QStackedLayout_new2();
extern __declspec(dllexport) QStackedLayout* QStackedLayout_new3(QLayout* parentLayout);
extern __declspec(dllexport) void QStackedLayout_virtbase(QStackedLayout* src, QLayout** outptr_QLayout);
extern __declspec(dllexport) QMetaObject* QStackedLayout_MetaObject(const QStackedLayout* self);
extern __declspec(dllexport) void* QStackedLayout_Metacast(QStackedLayout* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QStackedLayout_Tr(const char* s);
extern __declspec(dllexport) int QStackedLayout_AddWidget(QStackedLayout* self, QWidget* w);
extern __declspec(dllexport) int QStackedLayout_InsertWidget(QStackedLayout* self, int index, QWidget* w);
extern __declspec(dllexport) QWidget* QStackedLayout_CurrentWidget(const QStackedLayout* self);
extern __declspec(dllexport) int QStackedLayout_CurrentIndex(const QStackedLayout* self);
extern __declspec(dllexport) QWidget* QStackedLayout_Widget(const QStackedLayout* self, int param1);
extern __declspec(dllexport) int QStackedLayout_Count(const QStackedLayout* self);
extern __declspec(dllexport) StackingMode QStackedLayout_StackingMode(const QStackedLayout* self);
extern __declspec(dllexport) void QStackedLayout_SetStackingMode(QStackedLayout* self, StackingMode stackingMode);
extern __declspec(dllexport) void QStackedLayout_AddItem(QStackedLayout* self, QLayoutItem* item);
extern __declspec(dllexport) QSize* QStackedLayout_SizeHint(const QStackedLayout* self);
extern __declspec(dllexport) QSize* QStackedLayout_MinimumSize(const QStackedLayout* self);
extern __declspec(dllexport) QLayoutItem* QStackedLayout_ItemAt(const QStackedLayout* self, int param1);
extern __declspec(dllexport) QLayoutItem* QStackedLayout_TakeAt(QStackedLayout* self, int param1);
extern __declspec(dllexport) void QStackedLayout_SetGeometry(QStackedLayout* self, QRect* rect);
extern __declspec(dllexport) bool QStackedLayout_HasHeightForWidth(const QStackedLayout* self);
extern __declspec(dllexport) int QStackedLayout_HeightForWidth(const QStackedLayout* self, int width);
extern __declspec(dllexport) void QStackedLayout_WidgetRemoved(QStackedLayout* self, int index);
void QStackedLayout_connect_WidgetRemoved(QStackedLayout* self, intptr_t slot);
extern __declspec(dllexport) void QStackedLayout_CurrentChanged(QStackedLayout* self, int index);
void QStackedLayout_connect_CurrentChanged(QStackedLayout* self, intptr_t slot);
extern __declspec(dllexport) void QStackedLayout_WidgetAdded(QStackedLayout* self, int index);
void QStackedLayout_connect_WidgetAdded(QStackedLayout* self, intptr_t slot);
extern __declspec(dllexport) void QStackedLayout_SetCurrentIndex(QStackedLayout* self, int index);
extern __declspec(dllexport) void QStackedLayout_SetCurrentWidget(QStackedLayout* self, QWidget* w);
extern __declspec(dllexport) struct miqt_string QStackedLayout_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QStackedLayout_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QStackedLayout_override_virtual_Count(void* self, intptr_t slot);
int QStackedLayout_virtualbase_Count(const void* self);
extern __declspec(dllexport) void QStackedLayout_override_virtual_AddItem(void* self, intptr_t slot);
void QStackedLayout_virtualbase_AddItem(void* self, QLayoutItem* item);
extern __declspec(dllexport) void QStackedLayout_override_virtual_SizeHint(void* self, intptr_t slot);
QSize* QStackedLayout_virtualbase_SizeHint(const void* self);
extern __declspec(dllexport) void QStackedLayout_override_virtual_MinimumSize(void* self, intptr_t slot);
QSize* QStackedLayout_virtualbase_MinimumSize(const void* self);
extern __declspec(dllexport) void QStackedLayout_override_virtual_ItemAt(void* self, intptr_t slot);
QLayoutItem* QStackedLayout_virtualbase_ItemAt(const void* self, int param1);
extern __declspec(dllexport) void QStackedLayout_override_virtual_TakeAt(void* self, intptr_t slot);
QLayoutItem* QStackedLayout_virtualbase_TakeAt(void* self, int param1);
extern __declspec(dllexport) void QStackedLayout_override_virtual_SetGeometry(void* self, intptr_t slot);
void QStackedLayout_virtualbase_SetGeometry(void* self, QRect* rect);
extern __declspec(dllexport) void QStackedLayout_override_virtual_HasHeightForWidth(void* self, intptr_t slot);
bool QStackedLayout_virtualbase_HasHeightForWidth(const void* self);
extern __declspec(dllexport) void QStackedLayout_override_virtual_HeightForWidth(void* self, intptr_t slot);
int QStackedLayout_virtualbase_HeightForWidth(const void* self, int width);
extern __declspec(dllexport) void QStackedLayout_override_virtual_Spacing(void* self, intptr_t slot);
int QStackedLayout_virtualbase_Spacing(const void* self);
extern __declspec(dllexport) void QStackedLayout_override_virtual_SetSpacing(void* self, intptr_t slot);
void QStackedLayout_virtualbase_SetSpacing(void* self, int spacing);
extern __declspec(dllexport) void QStackedLayout_override_virtual_Invalidate(void* self, intptr_t slot);
void QStackedLayout_virtualbase_Invalidate(void* self);
extern __declspec(dllexport) void QStackedLayout_override_virtual_Geometry(void* self, intptr_t slot);
QRect* QStackedLayout_virtualbase_Geometry(const void* self);
extern __declspec(dllexport) void QStackedLayout_override_virtual_ExpandingDirections(void* self, intptr_t slot);
int QStackedLayout_virtualbase_ExpandingDirections(const void* self);
extern __declspec(dllexport) void QStackedLayout_override_virtual_MaximumSize(void* self, intptr_t slot);
QSize* QStackedLayout_virtualbase_MaximumSize(const void* self);
extern __declspec(dllexport) void QStackedLayout_override_virtual_IndexOf(void* self, intptr_t slot);
int QStackedLayout_virtualbase_IndexOf(const void* self, QWidget* param1);
extern __declspec(dllexport) void QStackedLayout_override_virtual_IsEmpty(void* self, intptr_t slot);
bool QStackedLayout_virtualbase_IsEmpty(const void* self);
extern __declspec(dllexport) void QStackedLayout_override_virtual_ControlTypes(void* self, intptr_t slot);
int QStackedLayout_virtualbase_ControlTypes(const void* self);
extern __declspec(dllexport) void QStackedLayout_override_virtual_ReplaceWidget(void* self, intptr_t slot);
QLayoutItem* QStackedLayout_virtualbase_ReplaceWidget(void* self, QWidget* from, QWidget* to, int options);
extern __declspec(dllexport) void QStackedLayout_override_virtual_Layout(void* self, intptr_t slot);
QLayout* QStackedLayout_virtualbase_Layout(void* self);
extern __declspec(dllexport) void QStackedLayout_override_virtual_ChildEvent(void* self, intptr_t slot);
void QStackedLayout_virtualbase_ChildEvent(void* self, QChildEvent* e);
extern __declspec(dllexport) void QStackedLayout_Delete(QStackedLayout* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
