#pragma once
#ifndef MIQT_QT_GEN_QBYTEARRAYMATCHER_H
#define MIQT_QT_GEN_QBYTEARRAYMATCHER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QByteArrayMatcher QByteArrayMatcher;
typedef struct QByteArrayView QByteArrayView;
typedef struct QStaticByteArrayMatcherBase QStaticByteArrayMatcherBase;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QByteArrayMatcher* QByteArrayMatcher_new();
extern __declspec(dllexport) 
QByteArrayMatcher* QByteArrayMatcher_new2(struct miqt_string pattern);
extern __declspec(dllexport) 
QByteArrayMatcher* QByteArrayMatcher_new3(QByteArrayView* pattern);
extern __declspec(dllexport) 
QByteArrayMatcher* QByteArrayMatcher_new4(const char* pattern);
extern __declspec(dllexport) 
QByteArrayMatcher* QByteArrayMatcher_new5(QByteArrayMatcher* other);
extern __declspec(dllexport) 
QByteArrayMatcher* QByteArrayMatcher_new6(const char* pattern, ptrdiff_t length);
extern __declspec(dllexport) 
void QByteArrayMatcher_OperatorAssign(QByteArrayMatcher* self, QByteArrayMatcher* other);
extern __declspec(dllexport) 
void QByteArrayMatcher_SetPattern(QByteArrayMatcher* self, struct miqt_string pattern);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayMatcher_IndexIn(const QByteArrayMatcher* self, const char* str, ptrdiff_t lenVal);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayMatcher_IndexInWithData(const QByteArrayMatcher* self, QByteArrayView* data);
extern __declspec(dllexport) 
struct miqt_string QByteArrayMatcher_Pattern(const QByteArrayMatcher* self);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayMatcher_IndexIn3(const QByteArrayMatcher* self, const char* str, ptrdiff_t lenVal, ptrdiff_t from);
extern __declspec(dllexport) 
ptrdiff_t QByteArrayMatcher_IndexIn2(const QByteArrayMatcher* self, QByteArrayView* data, ptrdiff_t from);
extern __declspec(dllexport) 
void QByteArrayMatcher_Delete(QByteArrayMatcher* self, bool isSubclass);


}
