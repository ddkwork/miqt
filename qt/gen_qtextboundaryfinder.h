#pragma once
#ifndef MIQT_QT_GEN_QTEXTBOUNDARYFINDER_H
#define MIQT_QT_GEN_QTEXTBOUNDARYFINDER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChar;
class QTextBoundaryFinder;
class _GUID;
class type_info;
#else
typedef struct QChar QChar;
typedef struct QTextBoundaryFinder QTextBoundaryFinder;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTextBoundaryFinder* QTextBoundaryFinder_new();
extern __declspec(dllexport) QTextBoundaryFinder* QTextBoundaryFinder_new2(QTextBoundaryFinder* other);
extern __declspec(dllexport) QTextBoundaryFinder* QTextBoundaryFinder_new3(BoundaryType typeVal, struct miqt_string stringVal);
extern __declspec(dllexport) QTextBoundaryFinder* QTextBoundaryFinder_new4(BoundaryType typeVal, QChar* chars, ptrdiff_t length);
extern __declspec(dllexport) QTextBoundaryFinder* QTextBoundaryFinder_new5(BoundaryType typeVal, QChar* chars, ptrdiff_t length, unsigned char* buffer);
extern __declspec(dllexport) QTextBoundaryFinder* QTextBoundaryFinder_new6(BoundaryType typeVal, QChar* chars, ptrdiff_t length, unsigned char* buffer, ptrdiff_t bufferSize);
extern __declspec(dllexport) void QTextBoundaryFinder_OperatorAssign(QTextBoundaryFinder* self, QTextBoundaryFinder* other);
extern __declspec(dllexport) bool QTextBoundaryFinder_IsValid(const QTextBoundaryFinder* self);
extern __declspec(dllexport) BoundaryType QTextBoundaryFinder_Type(const QTextBoundaryFinder* self);
extern __declspec(dllexport) struct miqt_string QTextBoundaryFinder_String(const QTextBoundaryFinder* self);
extern __declspec(dllexport) void QTextBoundaryFinder_ToStart(QTextBoundaryFinder* self);
extern __declspec(dllexport) void QTextBoundaryFinder_ToEnd(QTextBoundaryFinder* self);
extern __declspec(dllexport) ptrdiff_t QTextBoundaryFinder_Position(const QTextBoundaryFinder* self);
extern __declspec(dllexport) void QTextBoundaryFinder_SetPosition(QTextBoundaryFinder* self, ptrdiff_t position);
extern __declspec(dllexport) ptrdiff_t QTextBoundaryFinder_ToNextBoundary(QTextBoundaryFinder* self);
extern __declspec(dllexport) ptrdiff_t QTextBoundaryFinder_ToPreviousBoundary(QTextBoundaryFinder* self);
extern __declspec(dllexport) bool QTextBoundaryFinder_IsAtBoundary(const QTextBoundaryFinder* self);
extern __declspec(dllexport) BoundaryReasons QTextBoundaryFinder_BoundaryReasons(const QTextBoundaryFinder* self);
extern __declspec(dllexport) void QTextBoundaryFinder_Delete(QTextBoundaryFinder* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
