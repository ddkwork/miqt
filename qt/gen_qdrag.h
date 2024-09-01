#ifndef GEN_QDRAG_H
#define GEN_QDRAG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QDrag;
class QMetaObject;
class QMimeData;
class QObject;
class QPixmap;
class QPoint;
#else
typedef struct QDrag QDrag;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QObject QObject;
typedef struct QPixmap QPixmap;
typedef struct QPoint QPoint;
#endif

QDrag* QDrag_new(QObject* dragSource);
QMetaObject* QDrag_MetaObject(QDrag* self);
void QDrag_Tr(const char* s, char** _out, int* _out_Strlen);
void QDrag_TrUtf8(const char* s, char** _out, int* _out_Strlen);
void QDrag_SetMimeData(QDrag* self, QMimeData* data);
QMimeData* QDrag_MimeData(QDrag* self);
void QDrag_SetPixmap(QDrag* self, QPixmap* pixmap);
QPixmap* QDrag_Pixmap(QDrag* self);
void QDrag_SetHotSpot(QDrag* self, QPoint* hotspot);
QPoint* QDrag_HotSpot(QDrag* self);
QObject* QDrag_Source(QDrag* self);
QObject* QDrag_Target(QDrag* self);
uintptr_t QDrag_Start(QDrag* self);
uintptr_t QDrag_Exec(QDrag* self);
uintptr_t QDrag_Exec2(QDrag* self, int supportedActions, uintptr_t defaultAction);
void QDrag_SetDragCursor(QDrag* self, QPixmap* cursor, uintptr_t action);
QPixmap* QDrag_DragCursor(QDrag* self, uintptr_t action);
int QDrag_SupportedActions(QDrag* self);
uintptr_t QDrag_DefaultAction(QDrag* self);
void QDrag_Cancel();
void QDrag_ActionChanged(QDrag* self, uintptr_t action);
void QDrag_connect_ActionChanged(QDrag* self, void* slot);
void QDrag_TargetChanged(QDrag* self, QObject* newTarget);
void QDrag_connect_TargetChanged(QDrag* self, void* slot);
void QDrag_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QDrag_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QDrag_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QDrag_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
uintptr_t QDrag_Start1(QDrag* self, int supportedActions);
uintptr_t QDrag_Exec1(QDrag* self, int supportedActions);
void QDrag_Delete(QDrag* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif