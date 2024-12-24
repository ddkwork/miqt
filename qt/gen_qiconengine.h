#pragma once
#ifndef MIQT_QT_GEN_QICONENGINE_H
#define MIQT_QT_GEN_QICONENGINE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QIconEngine__ScaledPixmapArgument)
typedef QIconEngine::ScaledPixmapArgument QIconEngine__ScaledPixmapArgument;
typedef struct QDataStream QDataStream;
typedef struct QIconEngine QIconEngine;
typedef struct QIconEngine__ScaledPixmapArgument QIconEngine__ScaledPixmapArgument;
typedef struct QPainter QPainter;
typedef struct QPixmap QPixmap;
typedef struct QRect QRect;
typedef struct QSize QSize;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QIconEngine* QIconEngine_new();
extern __declspec(dllexport) void QIconEngine_Paint(QIconEngine* self, QPainter* painter, QRect* rect, int mode, int state);
extern __declspec(dllexport) QSize* QIconEngine_ActualSize(QIconEngine* self, QSize* size, int mode, int state);
extern __declspec(dllexport) QPixmap* QIconEngine_Pixmap(QIconEngine* self, QSize* size, int mode, int state);
extern __declspec(dllexport) void QIconEngine_AddPixmap(QIconEngine* self, QPixmap* pixmap, int mode, int state);
extern __declspec(dllexport) void QIconEngine_AddFile(QIconEngine* self, struct miqt_string fileName, QSize* size, int mode, int state);
extern __declspec(dllexport) struct miqt_string QIconEngine_Key(const QIconEngine* self);
extern __declspec(dllexport) QIconEngine* QIconEngine_Clone(const QIconEngine* self);
extern __declspec(dllexport) bool QIconEngine_Read(QIconEngine* self, QDataStream* in);
extern __declspec(dllexport) bool QIconEngine_Write(const QIconEngine* self, QDataStream* out);
extern __declspec(dllexport) struct miqt_array /* of QSize* */  QIconEngine_AvailableSizes(QIconEngine* self, int mode, int state);
extern __declspec(dllexport) struct miqt_string QIconEngine_IconName(QIconEngine* self);
extern __declspec(dllexport) bool QIconEngine_IsNull(QIconEngine* self);
extern __declspec(dllexport) QPixmap* QIconEngine_ScaledPixmap(QIconEngine* self, QSize* size, int mode, int state, double scale);
extern __declspec(dllexport) void QIconEngine_VirtualHook(QIconEngine* self, int id, void* data);
extern __declspec(dllexport) void QIconEngine_override_virtual_Paint(void* self, intptr_t slot);
void QIconEngine_virtualbase_Paint(void* self, QPainter* painter, QRect* rect, int mode, int state);
extern __declspec(dllexport) void QIconEngine_override_virtual_ActualSize(void* self, intptr_t slot);
QSize* QIconEngine_virtualbase_ActualSize(void* self, QSize* size, int mode, int state);
extern __declspec(dllexport) void QIconEngine_override_virtual_Pixmap(void* self, intptr_t slot);
QPixmap* QIconEngine_virtualbase_Pixmap(void* self, QSize* size, int mode, int state);
extern __declspec(dllexport) void QIconEngine_override_virtual_AddPixmap(void* self, intptr_t slot);
void QIconEngine_virtualbase_AddPixmap(void* self, QPixmap* pixmap, int mode, int state);
extern __declspec(dllexport) void QIconEngine_override_virtual_AddFile(void* self, intptr_t slot);
void QIconEngine_virtualbase_AddFile(void* self, struct miqt_string fileName, QSize* size, int mode, int state);
extern __declspec(dllexport) void QIconEngine_override_virtual_Key(void* self, intptr_t slot);
struct miqt_string QIconEngine_virtualbase_Key(const void* self);
extern __declspec(dllexport) void QIconEngine_override_virtual_Clone(void* self, intptr_t slot);
QIconEngine* QIconEngine_virtualbase_Clone(const void* self);
extern __declspec(dllexport) void QIconEngine_override_virtual_Read(void* self, intptr_t slot);
bool QIconEngine_virtualbase_Read(void* self, QDataStream* in);
extern __declspec(dllexport) void QIconEngine_override_virtual_Write(void* self, intptr_t slot);
bool QIconEngine_virtualbase_Write(const void* self, QDataStream* out);
extern __declspec(dllexport) void QIconEngine_override_virtual_AvailableSizes(void* self, intptr_t slot);
struct miqt_array /* of QSize* */  QIconEngine_virtualbase_AvailableSizes(void* self, int mode, int state);
extern __declspec(dllexport) void QIconEngine_override_virtual_IconName(void* self, intptr_t slot);
struct miqt_string QIconEngine_virtualbase_IconName(void* self);
extern __declspec(dllexport) void QIconEngine_override_virtual_IsNull(void* self, intptr_t slot);
bool QIconEngine_virtualbase_IsNull(void* self);
extern __declspec(dllexport) void QIconEngine_override_virtual_ScaledPixmap(void* self, intptr_t slot);
QPixmap* QIconEngine_virtualbase_ScaledPixmap(void* self, QSize* size, int mode, int state, double scale);
extern __declspec(dllexport) void QIconEngine_override_virtual_VirtualHook(void* self, intptr_t slot);
void QIconEngine_virtualbase_VirtualHook(void* self, int id, void* data);
extern __declspec(dllexport) void QIconEngine_Delete(QIconEngine* self, bool isSubclass);

extern __declspec(dllexport) QIconEngine__ScaledPixmapArgument* QIconEngine__ScaledPixmapArgument_new(const ScaledPixmapArgument* param1);
extern __declspec(dllexport) void QIconEngine__ScaledPixmapArgument_OperatorAssign(QIconEngine__ScaledPixmapArgument* self, const ScaledPixmapArgument* param1);
extern __declspec(dllexport) void QIconEngine__ScaledPixmapArgument_Delete(QIconEngine__ScaledPixmapArgument* self, bool isSubclass);

} 
