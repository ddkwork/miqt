#pragma once
#ifndef MIQT_QT_NETWORK_GEN_QFORMDATABUILDER_H
#define MIQT_QT_NETWORK_GEN_QFORMDATABUILDER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAnyStringView QAnyStringView;
typedef struct QByteArrayView QByteArrayView;
typedef struct QFormDataBuilder QFormDataBuilder;
typedef struct QFormDataPartBuilder QFormDataPartBuilder;
typedef struct QHttpHeaders QHttpHeaders;
typedef struct QIODevice QIODevice;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QFormDataPartBuilder* QFormDataPartBuilder_new();
extern __declspec(dllexport) QFormDataPartBuilder* QFormDataPartBuilder_new2(QFormDataPartBuilder* param1);
extern __declspec(dllexport) void QFormDataPartBuilder_Swap(QFormDataPartBuilder* self, QFormDataPartBuilder* other);
extern __declspec(dllexport) QFormDataPartBuilder* QFormDataPartBuilder_SetBody(QFormDataPartBuilder* self, QByteArrayView* data);
extern __declspec(dllexport) QFormDataPartBuilder* QFormDataPartBuilder_SetBodyDevice(QFormDataPartBuilder* self, QIODevice* body);
extern __declspec(dllexport) QFormDataPartBuilder* QFormDataPartBuilder_SetHeaders(QFormDataPartBuilder* self, QHttpHeaders* headers);
extern __declspec(dllexport) QFormDataPartBuilder* QFormDataPartBuilder_SetBody2(QFormDataPartBuilder* self, QByteArrayView* data, QAnyStringView* fileName);
extern __declspec(dllexport) QFormDataPartBuilder* QFormDataPartBuilder_SetBody3(QFormDataPartBuilder* self, QByteArrayView* data, QAnyStringView* fileName, QAnyStringView* mimeType);
extern __declspec(dllexport) QFormDataPartBuilder* QFormDataPartBuilder_SetBodyDevice2(QFormDataPartBuilder* self, QIODevice* body, QAnyStringView* fileName);
extern __declspec(dllexport) QFormDataPartBuilder* QFormDataPartBuilder_SetBodyDevice3(QFormDataPartBuilder* self, QIODevice* body, QAnyStringView* fileName, QAnyStringView* mimeType);
extern __declspec(dllexport) void QFormDataPartBuilder_Delete(QFormDataPartBuilder* self, bool isSubclass);

extern __declspec(dllexport) QFormDataBuilder* QFormDataBuilder_new();
extern __declspec(dllexport) void QFormDataBuilder_Swap(QFormDataBuilder* self, QFormDataBuilder* other);
extern __declspec(dllexport) QFormDataPartBuilder* QFormDataBuilder_Part(QFormDataBuilder* self, QAnyStringView* name);
extern __declspec(dllexport) void QFormDataBuilder_Delete(QFormDataBuilder* self, bool isSubclass);

} 
