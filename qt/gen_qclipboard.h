#pragma once
#ifndef MIQT_QT_GEN_QCLIPBOARD_H
#define MIQT_QT_GEN_QCLIPBOARD_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QClipboard QClipboard;
typedef struct QImage QImage;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QObject QObject;
typedef struct QPixmap QPixmap;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QClipboard_virtbase(QClipboard* src, QObject** outptr_QObject);
extern __declspec(dllexport) QMetaObject* QClipboard_MetaObject(const QClipboard* self);
extern __declspec(dllexport) void* QClipboard_Metacast(QClipboard* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QClipboard_Tr(const char* s);
extern __declspec(dllexport) void QClipboard_Clear(QClipboard* self);
extern __declspec(dllexport) bool QClipboard_SupportsSelection(const QClipboard* self);
extern __declspec(dllexport) bool QClipboard_SupportsFindBuffer(const QClipboard* self);
extern __declspec(dllexport) bool QClipboard_OwnsSelection(const QClipboard* self);
extern __declspec(dllexport) bool QClipboard_OwnsClipboard(const QClipboard* self);
extern __declspec(dllexport) bool QClipboard_OwnsFindBuffer(const QClipboard* self);
extern __declspec(dllexport) struct miqt_string QClipboard_Text(const QClipboard* self);
extern __declspec(dllexport) struct miqt_string QClipboard_TextWithSubtype(const QClipboard* self, struct miqt_string subtype);
extern __declspec(dllexport) void QClipboard_SetText(QClipboard* self, struct miqt_string param1);
extern __declspec(dllexport) QMimeData* QClipboard_MimeData(const QClipboard* self);
extern __declspec(dllexport) void QClipboard_SetMimeData(QClipboard* self, QMimeData* data);
extern __declspec(dllexport) QImage* QClipboard_Image(const QClipboard* self);
extern __declspec(dllexport) QPixmap* QClipboard_Pixmap(const QClipboard* self);
extern __declspec(dllexport) void QClipboard_SetImage(QClipboard* self, QImage* param1);
extern __declspec(dllexport) void QClipboard_SetPixmap(QClipboard* self, QPixmap* param1);
extern __declspec(dllexport) void QClipboard_Changed(QClipboard* self, int mode);
void QClipboard_connect_Changed(QClipboard* self, intptr_t slot);
extern __declspec(dllexport) void QClipboard_SelectionChanged(QClipboard* self);
void QClipboard_connect_SelectionChanged(QClipboard* self, intptr_t slot);
extern __declspec(dllexport) void QClipboard_FindBufferChanged(QClipboard* self);
void QClipboard_connect_FindBufferChanged(QClipboard* self, intptr_t slot);
extern __declspec(dllexport) void QClipboard_DataChanged(QClipboard* self);
void QClipboard_connect_DataChanged(QClipboard* self, intptr_t slot);
extern __declspec(dllexport) struct miqt_string QClipboard_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QClipboard_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QClipboard_Clear1(QClipboard* self, Mode mode);
extern __declspec(dllexport) struct miqt_string QClipboard_Text1(const QClipboard* self, Mode mode);
extern __declspec(dllexport) struct miqt_string QClipboard_Text2(const QClipboard* self, struct miqt_string subtype, Mode mode);
extern __declspec(dllexport) void QClipboard_SetText2(QClipboard* self, struct miqt_string param1, Mode mode);
extern __declspec(dllexport) QMimeData* QClipboard_MimeData1(const QClipboard* self, Mode mode);
extern __declspec(dllexport) void QClipboard_SetMimeData2(QClipboard* self, QMimeData* data, Mode mode);
extern __declspec(dllexport) QImage* QClipboard_Image1(const QClipboard* self, Mode mode);
extern __declspec(dllexport) QPixmap* QClipboard_Pixmap1(const QClipboard* self, Mode mode);
extern __declspec(dllexport) void QClipboard_SetImage2(QClipboard* self, QImage* param1, Mode mode);
extern __declspec(dllexport) void QClipboard_SetPixmap2(QClipboard* self, QPixmap* param1, Mode mode);

} 
