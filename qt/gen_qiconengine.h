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

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QIconEngine* QIconEngine_new();
extern __declspec(dllexport) 
void QIconEngine_Paint(QIconEngine* self, QPainter* painter, QRect* rect, int mode, int state);
extern __declspec(dllexport) 
QSize* QIconEngine_ActualSize(QIconEngine* self, QSize* size, int mode, int state);
extern __declspec(dllexport) 
QPixmap* QIconEngine_Pixmap(QIconEngine* self, QSize* size, int mode, int state);
extern __declspec(dllexport) 
void QIconEngine_AddPixmap(QIconEngine* self, QPixmap* pixmap, int mode, int state);
extern __declspec(dllexport) 
void QIconEngine_AddFile(QIconEngine* self, struct miqt_string fileName, QSize* size, int mode, int state);
extern __declspec(dllexport) 
struct miqt_string QIconEngine_Key(const QIconEngine* self);
extern __declspec(dllexport) 
QIconEngine* QIconEngine_Clone(const QIconEngine* self);
extern __declspec(dllexport) 
bool QIconEngine_Read(QIconEngine* self, QDataStream* in);
extern __declspec(dllexport) 
bool QIconEngine_Write(const QIconEngine* self, QDataStream* out);
extern __declspec(dllexport) 
struct miqt_array /* of QSize* */  QIconEngine_AvailableSizes(QIconEngine* self, int mode, int state);
extern __declspec(dllexport) 
struct miqt_string QIconEngine_IconName(QIconEngine* self);
extern __declspec(dllexport) 
bool QIconEngine_IsNull(QIconEngine* self);
extern __declspec(dllexport) 
QPixmap* QIconEngine_ScaledPixmap(QIconEngine* self, QSize* size, int mode, int state, double scale);
extern __declspec(dllexport) 
void QIconEngine_VirtualHook(QIconEngine* self, int id, void* data);
extern __declspec(dllexport) 
void QIconEngine_Delete(QIconEngine* self, bool isSubclass);

extern __declspec(dllexport) 
QIconEngine__ScaledPixmapArgument* QIconEngine__ScaledPixmapArgument_new(const ScaledPixmapArgument* param1);
extern __declspec(dllexport) 
void QIconEngine__ScaledPixmapArgument_OperatorAssign(QIconEngine__ScaledPixmapArgument* self, const ScaledPixmapArgument* param1);
extern __declspec(dllexport) 
void QIconEngine__ScaledPixmapArgument_Delete(QIconEngine__ScaledPixmapArgument* self, bool isSubclass);

}
