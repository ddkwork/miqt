#pragma once
#ifndef MIQT_QT_GEN_QTEXTDOCUMENTWRITER_H
#define MIQT_QT_GEN_QTEXTDOCUMENTWRITER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QIODevice;
class QTextDocument;
class QTextDocumentFragment;
class QTextDocumentWriter;
class _GUID;
class type_info;
#else
typedef struct QIODevice QIODevice;
typedef struct QTextDocument QTextDocument;
typedef struct QTextDocumentFragment QTextDocumentFragment;
typedef struct QTextDocumentWriter QTextDocumentWriter;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTextDocumentWriter* QTextDocumentWriter_new();
extern __declspec(dllexport) QTextDocumentWriter* QTextDocumentWriter_new2(QIODevice* device, struct miqt_string format);
extern __declspec(dllexport) QTextDocumentWriter* QTextDocumentWriter_new3(struct miqt_string fileName);
extern __declspec(dllexport) QTextDocumentWriter* QTextDocumentWriter_new4(struct miqt_string fileName, struct miqt_string format);
extern __declspec(dllexport) void QTextDocumentWriter_SetFormat(QTextDocumentWriter* self, struct miqt_string format);
extern __declspec(dllexport) struct miqt_string QTextDocumentWriter_Format(const QTextDocumentWriter* self);
extern __declspec(dllexport) void QTextDocumentWriter_SetDevice(QTextDocumentWriter* self, QIODevice* device);
extern __declspec(dllexport) QIODevice* QTextDocumentWriter_Device(const QTextDocumentWriter* self);
extern __declspec(dllexport) void QTextDocumentWriter_SetFileName(QTextDocumentWriter* self, struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QTextDocumentWriter_FileName(const QTextDocumentWriter* self);
extern __declspec(dllexport) bool QTextDocumentWriter_Write(QTextDocumentWriter* self, QTextDocument* document);
extern __declspec(dllexport) bool QTextDocumentWriter_WriteWithFragment(QTextDocumentWriter* self, QTextDocumentFragment* fragment);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QTextDocumentWriter_SupportedDocumentFormats();
extern __declspec(dllexport) void QTextDocumentWriter_Delete(QTextDocumentWriter* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
