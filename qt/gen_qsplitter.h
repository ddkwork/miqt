#pragma once
#ifndef MIQT_QT_GEN_QSPLITTER_H
#define MIQT_QT_GEN_QSPLITTER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QFrame QFrame;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QResizeEvent QResizeEvent;
typedef struct QSize QSize;
typedef struct QSplitter QSplitter;
typedef struct QSplitterHandle QSplitterHandle;
typedef struct QTextStream QTextStream;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
void QTextStream_Delete(QTextStream* self, bool isSubclass);

extern __declspec(dllexport) 
QSplitter* QSplitter_new(QWidget* parent);
extern __declspec(dllexport) 
QSplitter* QSplitter_new2();
extern __declspec(dllexport) 
QSplitter* QSplitter_new3(int param1);
extern __declspec(dllexport) 
QSplitter* QSplitter_new4(int param1, QWidget* parent);
extern __declspec(dllexport) 
void QSplitter_virtbase(QSplitter* src
, QFrame** outptr_QFrame
);
extern __declspec(dllexport) 
QMetaObject* QSplitter_MetaObject(const QSplitter* self);
extern __declspec(dllexport) 
void* QSplitter_Metacast(QSplitter* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QSplitter_Tr(const char* s);
extern __declspec(dllexport) 
void QSplitter_AddWidget(QSplitter* self, QWidget* widget);
extern __declspec(dllexport) 
void QSplitter_InsertWidget(QSplitter* self, int index, QWidget* widget);
extern __declspec(dllexport) 
QWidget* QSplitter_ReplaceWidget(QSplitter* self, int index, QWidget* widget);
extern __declspec(dllexport) 
void QSplitter_SetOrientation(QSplitter* self, int orientation);
extern __declspec(dllexport) 
int QSplitter_Orientation(const QSplitter* self);
extern __declspec(dllexport) 
void QSplitter_SetChildrenCollapsible(QSplitter* self, bool childrenCollapsible);
extern __declspec(dllexport) 
bool QSplitter_ChildrenCollapsible(const QSplitter* self);
extern __declspec(dllexport) 
void QSplitter_SetCollapsible(QSplitter* self, int index, bool param2);
extern __declspec(dllexport) 
bool QSplitter_IsCollapsible(const QSplitter* self, int index);
extern __declspec(dllexport) 
void QSplitter_SetOpaqueResize(QSplitter* self);
extern __declspec(dllexport) 
bool QSplitter_OpaqueResize(const QSplitter* self);
extern __declspec(dllexport) 
void QSplitter_Refresh(QSplitter* self);
extern __declspec(dllexport) 
QSize* QSplitter_SizeHint(const QSplitter* self);
extern __declspec(dllexport) 
QSize* QSplitter_MinimumSizeHint(const QSplitter* self);
extern __declspec(dllexport) 
struct miqt_array /* of int */  QSplitter_Sizes(const QSplitter* self);
extern __declspec(dllexport) 
void QSplitter_SetSizes(QSplitter* self, struct miqt_array /* of int */  list);
extern __declspec(dllexport) 
struct miqt_string QSplitter_SaveState(const QSplitter* self);
extern __declspec(dllexport) 
bool QSplitter_RestoreState(QSplitter* self, struct miqt_string state);
extern __declspec(dllexport) 
int QSplitter_HandleWidth(const QSplitter* self);
extern __declspec(dllexport) 
void QSplitter_SetHandleWidth(QSplitter* self, int handleWidth);
extern __declspec(dllexport) 
int QSplitter_IndexOf(const QSplitter* self, QWidget* w);
extern __declspec(dllexport) 
QWidget* QSplitter_Widget(const QSplitter* self, int index);
extern __declspec(dllexport) 
int QSplitter_Count(const QSplitter* self);
extern __declspec(dllexport) 
void QSplitter_GetRange(const QSplitter* self, int index, int* param2, int* param3);
extern __declspec(dllexport) 
QSplitterHandle* QSplitter_Handle(const QSplitter* self, int index);
extern __declspec(dllexport) 
void QSplitter_SetStretchFactor(QSplitter* self, int index, int stretch);
extern __declspec(dllexport) 
void QSplitter_SplitterMoved(QSplitter* self, int pos, int index);
void QSplitter_connect_SplitterMoved(QSplitter* self, intptr_t slot);
extern __declspec(dllexport) 
QSplitterHandle* QSplitter_CreateHandle(QSplitter* self);
extern __declspec(dllexport) 
void QSplitter_ChildEvent(QSplitter* self, QChildEvent* param1);
extern __declspec(dllexport) 
bool QSplitter_Event(QSplitter* self, QEvent* param1);
extern __declspec(dllexport) 
void QSplitter_ResizeEvent(QSplitter* self, QResizeEvent* param1);
extern __declspec(dllexport) 
void QSplitter_ChangeEvent(QSplitter* self, QEvent* param1);
extern __declspec(dllexport) 
struct miqt_string QSplitter_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QSplitter_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QSplitter_SetOpaqueResize1(QSplitter* self, bool opaque);
extern __declspec(dllexport) 
void QSplitter_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QSplitter_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QSplitter_override_virtual_Metacast(void* self, intptr_t slot);
void* QSplitter_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QSplitter_Delete(QSplitter* self, bool isSubclass);

extern __declspec(dllexport) 
QSplitterHandle* QSplitterHandle_new(int o, QSplitter* parent);
extern __declspec(dllexport) 
void QSplitterHandle_virtbase(QSplitterHandle* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QSplitterHandle_MetaObject(const QSplitterHandle* self);
extern __declspec(dllexport) 
void* QSplitterHandle_Metacast(QSplitterHandle* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QSplitterHandle_Tr(const char* s);
extern __declspec(dllexport) 
void QSplitterHandle_SetOrientation(QSplitterHandle* self, int o);
extern __declspec(dllexport) 
int QSplitterHandle_Orientation(const QSplitterHandle* self);
extern __declspec(dllexport) 
bool QSplitterHandle_OpaqueResize(const QSplitterHandle* self);
extern __declspec(dllexport) 
QSplitter* QSplitterHandle_Splitter(const QSplitterHandle* self);
extern __declspec(dllexport) 
QSize* QSplitterHandle_SizeHint(const QSplitterHandle* self);
extern __declspec(dllexport) 
void QSplitterHandle_PaintEvent(QSplitterHandle* self, QPaintEvent* param1);
extern __declspec(dllexport) 
void QSplitterHandle_MouseMoveEvent(QSplitterHandle* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QSplitterHandle_MousePressEvent(QSplitterHandle* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QSplitterHandle_MouseReleaseEvent(QSplitterHandle* self, QMouseEvent* param1);
extern __declspec(dllexport) 
void QSplitterHandle_ResizeEvent(QSplitterHandle* self, QResizeEvent* param1);
extern __declspec(dllexport) 
bool QSplitterHandle_Event(QSplitterHandle* self, QEvent* param1);
extern __declspec(dllexport) 
struct miqt_string QSplitterHandle_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QSplitterHandle_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QSplitterHandle_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QSplitterHandle_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QSplitterHandle_override_virtual_Metacast(void* self, intptr_t slot);
void* QSplitterHandle_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QSplitterHandle_Delete(QSplitterHandle* self, bool isSubclass);

}
