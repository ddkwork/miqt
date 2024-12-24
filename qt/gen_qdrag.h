#pragma once
#ifndef MIQT_QT_GEN_QDRAG_H
#define MIQT_QT_GEN_QDRAG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDrag QDrag;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QObject QObject;
typedef struct QPixmap QPixmap;
typedef struct QPoint QPoint;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QDrag* QDrag_new(QObject* dragSource);
extern __declspec(dllexport) 
void QDrag_virtbase(QDrag* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QDrag_MetaObject(const QDrag* self);
extern __declspec(dllexport) 
void* QDrag_Metacast(QDrag* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QDrag_Tr(const char* s);
extern __declspec(dllexport) 
void QDrag_SetMimeData(QDrag* self, QMimeData* data);
extern __declspec(dllexport) 
QMimeData* QDrag_MimeData(const QDrag* self);
extern __declspec(dllexport) 
void QDrag_SetPixmap(QDrag* self, QPixmap* pixmap);
extern __declspec(dllexport) 
QPixmap* QDrag_Pixmap(const QDrag* self);
extern __declspec(dllexport) 
void QDrag_SetHotSpot(QDrag* self, QPoint* hotspot);
extern __declspec(dllexport) 
QPoint* QDrag_HotSpot(const QDrag* self);
extern __declspec(dllexport) 
QObject* QDrag_Source(const QDrag* self);
extern __declspec(dllexport) 
QObject* QDrag_Target(const QDrag* self);
extern __declspec(dllexport) 
int QDrag_Exec(QDrag* self);
extern __declspec(dllexport) 
int QDrag_Exec2(QDrag* self, int supportedActions, int defaultAction);
extern __declspec(dllexport) 
void QDrag_SetDragCursor(QDrag* self, QPixmap* cursor, int action);
extern __declspec(dllexport) 
QPixmap* QDrag_DragCursor(const QDrag* self, int action);
extern __declspec(dllexport) 
int QDrag_SupportedActions(const QDrag* self);
extern __declspec(dllexport) 
int QDrag_DefaultAction(const QDrag* self);
extern __declspec(dllexport) 
void QDrag_Cancel();
extern __declspec(dllexport) 
void QDrag_ActionChanged(QDrag* self, int action);
void QDrag_connect_ActionChanged(QDrag* self, intptr_t slot);
extern __declspec(dllexport) 
void QDrag_TargetChanged(QDrag* self, QObject* newTarget);
void QDrag_connect_TargetChanged(QDrag* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QDrag_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QDrag_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
int QDrag_Exec1(QDrag* self, int supportedActions);
extern __declspec(dllexport) 
void QDrag_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QDrag_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QDrag_override_virtual_Metacast(void* self, intptr_t slot);
void* QDrag_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QDrag_Delete(QDrag* self, bool isSubclass);

}
