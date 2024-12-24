#pragma once
#ifndef MIQT_QT_MULTIMEDIA_GEN_QABSTRACTVIDEOBUFFER_H
#define MIQT_QT_MULTIMEDIA_GEN_QABSTRACTVIDEOBUFFER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QAbstractVideoBuffer__MapData)
typedef QAbstractVideoBuffer::MapData QAbstractVideoBuffer__MapData;
typedef struct QAbstractVideoBuffer QAbstractVideoBuffer;
typedef struct QAbstractVideoBuffer__MapData QAbstractVideoBuffer__MapData;
typedef struct QVideoFrameFormat QVideoFrameFormat;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
MapData QAbstractVideoBuffer_Map(QAbstractVideoBuffer* self, int mode);
extern __declspec(dllexport) 
void QAbstractVideoBuffer_Unmap(QAbstractVideoBuffer* self);
extern __declspec(dllexport) 
QVideoFrameFormat* QAbstractVideoBuffer_Format(const QAbstractVideoBuffer* self);
extern __declspec(dllexport) 
void QAbstractVideoBuffer_Delete(QAbstractVideoBuffer* self, bool isSubclass);

extern __declspec(dllexport) 
void QAbstractVideoBuffer__MapData_Delete(QAbstractVideoBuffer__MapData* self, bool isSubclass);

}
