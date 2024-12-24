#pragma once
#ifndef MIQT_QT_GEN_QINPUTMETHOD_H
#define MIQT_QT_GEN_QINPUTMETHOD_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QInputMethod QInputMethod;
typedef struct QLocale QLocale;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QRectF QRectF;
typedef struct QTransform QTransform;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QInputMethod_virtbase(QInputMethod* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QInputMethod_MetaObject(const QInputMethod* self);
extern __declspec(dllexport) void* QInputMethod_Metacast(QInputMethod* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QInputMethod_Tr(const char* s);
extern __declspec(dllexport) QTransform* QInputMethod_InputItemTransform(const QInputMethod* self);
extern __declspec(dllexport) void QInputMethod_SetInputItemTransform(QInputMethod* self, QTransform* transform);
extern __declspec(dllexport) QRectF* QInputMethod_InputItemRectangle(const QInputMethod* self);
extern __declspec(dllexport) void QInputMethod_SetInputItemRectangle(QInputMethod* self, QRectF* rect);
extern __declspec(dllexport) QRectF* QInputMethod_CursorRectangle(const QInputMethod* self);
extern __declspec(dllexport) QRectF* QInputMethod_AnchorRectangle(const QInputMethod* self);
extern __declspec(dllexport) QRectF* QInputMethod_KeyboardRectangle(const QInputMethod* self);
extern __declspec(dllexport) QRectF* QInputMethod_InputItemClipRectangle(const QInputMethod* self);
extern __declspec(dllexport) bool QInputMethod_IsVisible(const QInputMethod* self);
extern __declspec(dllexport) void QInputMethod_SetVisible(QInputMethod* self, bool visible);
extern __declspec(dllexport) bool QInputMethod_IsAnimating(const QInputMethod* self);
extern __declspec(dllexport) QLocale* QInputMethod_Locale(const QInputMethod* self);
extern __declspec(dllexport) int QInputMethod_InputDirection(const QInputMethod* self);
extern __declspec(dllexport) QVariant* QInputMethod_QueryFocusObject(int query, QVariant* argument);
extern __declspec(dllexport) void QInputMethod_Show(QInputMethod* self);
extern __declspec(dllexport) void QInputMethod_Hide(QInputMethod* self);
extern __declspec(dllexport) void QInputMethod_Update(QInputMethod* self, int queries);
extern __declspec(dllexport) void QInputMethod_Reset(QInputMethod* self);
extern __declspec(dllexport) void QInputMethod_Commit(QInputMethod* self);
extern __declspec(dllexport) void QInputMethod_InvokeAction(QInputMethod* self, Action a, int cursorPosition);
extern __declspec(dllexport) void QInputMethod_CursorRectangleChanged(QInputMethod* self);
void QInputMethod_connect_CursorRectangleChanged(QInputMethod* self, intptr_t slot);
extern __declspec(dllexport) void QInputMethod_AnchorRectangleChanged(QInputMethod* self);
void QInputMethod_connect_AnchorRectangleChanged(QInputMethod* self, intptr_t slot);
extern __declspec(dllexport) void QInputMethod_KeyboardRectangleChanged(QInputMethod* self);
void QInputMethod_connect_KeyboardRectangleChanged(QInputMethod* self, intptr_t slot);
extern __declspec(dllexport) void QInputMethod_InputItemClipRectangleChanged(QInputMethod* self);
void QInputMethod_connect_InputItemClipRectangleChanged(QInputMethod* self, intptr_t slot);
extern __declspec(dllexport) void QInputMethod_VisibleChanged(QInputMethod* self);
void QInputMethod_connect_VisibleChanged(QInputMethod* self, intptr_t slot);
extern __declspec(dllexport) void QInputMethod_AnimatingChanged(QInputMethod* self);
void QInputMethod_connect_AnimatingChanged(QInputMethod* self, intptr_t slot);
extern __declspec(dllexport) void QInputMethod_LocaleChanged(QInputMethod* self);
void QInputMethod_connect_LocaleChanged(QInputMethod* self, intptr_t slot);
extern __declspec(dllexport) void QInputMethod_InputDirectionChanged(QInputMethod* self, int newDirection);
void QInputMethod_connect_InputDirectionChanged(QInputMethod* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QInputMethod_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QInputMethod_Tr3(const char* s, const char* c, int n);

} 
