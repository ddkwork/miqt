#pragma once
#ifndef MIQT_QT_GEN_QSTRINGMATCHER_H
#define MIQT_QT_GEN_QSTRINGMATCHER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChar QChar;
typedef struct QStringMatcher QStringMatcher;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QStringMatcher* QStringMatcher_new();
extern __declspec(dllexport) 
QStringMatcher* QStringMatcher_new2(struct miqt_string pattern);
extern __declspec(dllexport) 
QStringMatcher* QStringMatcher_new3(QChar* uc, ptrdiff_t lenVal);
extern __declspec(dllexport) 
QStringMatcher* QStringMatcher_new4(QStringMatcher* other);
extern __declspec(dllexport) 
QStringMatcher* QStringMatcher_new5(struct miqt_string pattern, int cs);
extern __declspec(dllexport) 
QStringMatcher* QStringMatcher_new6(QChar* uc, ptrdiff_t lenVal, int cs);
extern __declspec(dllexport) 
void QStringMatcher_OperatorAssign(QStringMatcher* self, QStringMatcher* other);
extern __declspec(dllexport) 
void QStringMatcher_SetPattern(QStringMatcher* self, struct miqt_string pattern);
extern __declspec(dllexport) 
void QStringMatcher_SetCaseSensitivity(QStringMatcher* self, int cs);
extern __declspec(dllexport) 
ptrdiff_t QStringMatcher_IndexIn(const QStringMatcher* self, struct miqt_string str);
extern __declspec(dllexport) 
ptrdiff_t QStringMatcher_IndexIn2(const QStringMatcher* self, QChar* str, ptrdiff_t length);
extern __declspec(dllexport) 
struct miqt_string QStringMatcher_Pattern(const QStringMatcher* self);
extern __declspec(dllexport) 
int QStringMatcher_CaseSensitivity(const QStringMatcher* self);
extern __declspec(dllexport) 
ptrdiff_t QStringMatcher_IndexIn22(const QStringMatcher* self, struct miqt_string str, ptrdiff_t from);
extern __declspec(dllexport) 
ptrdiff_t QStringMatcher_IndexIn3(const QStringMatcher* self, QChar* str, ptrdiff_t length, ptrdiff_t from);
extern __declspec(dllexport) 
void QStringMatcher_Delete(QStringMatcher* self, bool isSubclass);

}
