#pragma once
#ifndef MIQT_QT_GEN_QTEXTDOCUMENTFRAGMENT_H
#define MIQT_QT_GEN_QTEXTDOCUMENTFRAGMENT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QTextCursor;
class QTextDocument;
class QTextDocumentFragment;
class QTextStream;
class _GUID;
class type_info;
#else
typedef struct QTextCursor QTextCursor;
typedef struct QTextDocument QTextDocument;
typedef struct QTextDocumentFragment QTextDocumentFragment;
typedef struct QTextStream QTextStream;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QTextStream_Delete(QTextStream* self, bool isSubclass);

extern __declspec(dllexport) QTextDocumentFragment* QTextDocumentFragment_new();
extern __declspec(dllexport) QTextDocumentFragment* QTextDocumentFragment_new2(QTextDocument* document);
extern __declspec(dllexport) QTextDocumentFragment* QTextDocumentFragment_new3(QTextCursor* rangeVal);
extern __declspec(dllexport) QTextDocumentFragment* QTextDocumentFragment_new4(QTextDocumentFragment* rhs);
extern __declspec(dllexport) void QTextDocumentFragment_OperatorAssign(QTextDocumentFragment* self, QTextDocumentFragment* rhs);
extern __declspec(dllexport) bool QTextDocumentFragment_IsEmpty(const QTextDocumentFragment* self);
extern __declspec(dllexport) struct miqt_string QTextDocumentFragment_ToPlainText(const QTextDocumentFragment* self);
extern __declspec(dllexport) struct miqt_string QTextDocumentFragment_ToRawText(const QTextDocumentFragment* self);
extern __declspec(dllexport) struct miqt_string QTextDocumentFragment_ToHtml(const QTextDocumentFragment* self);
extern __declspec(dllexport) struct miqt_string QTextDocumentFragment_ToMarkdown(const QTextDocumentFragment* self);
extern __declspec(dllexport) QTextDocumentFragment* QTextDocumentFragment_FromPlainText(struct miqt_string plainText);
extern __declspec(dllexport) QTextDocumentFragment* QTextDocumentFragment_FromHtml(struct miqt_string html);
extern __declspec(dllexport) QTextDocumentFragment* QTextDocumentFragment_FromMarkdown(struct miqt_string markdown);
extern __declspec(dllexport) struct miqt_string QTextDocumentFragment_ToMarkdown1(const QTextDocumentFragment* self, int features);
extern __declspec(dllexport) QTextDocumentFragment* QTextDocumentFragment_FromHtml2(struct miqt_string html, QTextDocument* resourceProvider);
extern __declspec(dllexport) QTextDocumentFragment* QTextDocumentFragment_FromMarkdown2(struct miqt_string markdown, int features);
extern __declspec(dllexport) void QTextDocumentFragment_Delete(QTextDocumentFragment* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
