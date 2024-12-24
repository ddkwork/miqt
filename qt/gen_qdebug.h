#pragma once
#ifndef MIQT_QT_GEN_QDEBUG_H
#define MIQT_QT_GEN_QDEBUG_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QByteArrayView;
class QChar;
class QDebug;
class QDebugStateSaver;
class QIODevice;
class QIODeviceBase;
class QNoDebug;
class _GUID;
class type_info;
#else
typedef struct QByteArrayView QByteArrayView;
typedef struct QChar QChar;
typedef struct QDebug QDebug;
typedef struct QDebugStateSaver QDebugStateSaver;
typedef struct QIODevice QIODevice;
typedef struct QIODeviceBase QIODeviceBase;
typedef struct QNoDebug QNoDebug;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QDebug* QDebug_new(QIODevice* device);
extern __declspec(dllexport) QDebug* QDebug_new2(QDebug* o);
extern __declspec(dllexport) void QDebug_virtbase(QDebug* src, QIODeviceBase** outptr_QIODeviceBase);
extern __declspec(dllexport) void QDebug_OperatorAssign(QDebug* self, QDebug* other);
extern __declspec(dllexport) void QDebug_Swap(QDebug* self, QDebug* other);
extern __declspec(dllexport) QDebug* QDebug_ResetFormat(QDebug* self);
extern __declspec(dllexport) QDebug* QDebug_Space(QDebug* self);
extern __declspec(dllexport) QDebug* QDebug_Nospace(QDebug* self);
extern __declspec(dllexport) QDebug* QDebug_MaybeSpace(QDebug* self);
extern __declspec(dllexport) QDebug* QDebug_Verbosity(QDebug* self, int verbosityLevel);
extern __declspec(dllexport) int QDebug_Verbosity2(const QDebug* self);
extern __declspec(dllexport) void QDebug_SetVerbosity(QDebug* self, int verbosityLevel);
extern __declspec(dllexport) bool QDebug_AutoInsertSpaces(const QDebug* self);
extern __declspec(dllexport) void QDebug_SetAutoInsertSpaces(QDebug* self, bool b);
extern __declspec(dllexport) bool QDebug_QuoteStrings(const QDebug* self);
extern __declspec(dllexport) void QDebug_SetQuoteStrings(QDebug* self, bool b);
extern __declspec(dllexport) QDebug* QDebug_Quote(QDebug* self);
extern __declspec(dllexport) QDebug* QDebug_Noquote(QDebug* self);
extern __declspec(dllexport) QDebug* QDebug_MaybeQuote(QDebug* self);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeft(QDebug* self, QChar* t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithBool(QDebug* self, bool t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithChar(QDebug* self, char t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithShort(QDebug* self, int16_t t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithUnsignedshort(QDebug* self, uint16_t t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithInt(QDebug* self, int t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithUnsignedint(QDebug* self, unsigned int t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithLong(QDebug* self, long t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithUnsignedlong(QDebug* self, unsigned long t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithQint64(QDebug* self, long long t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithQuint64(QDebug* self, unsigned long long t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithFloat(QDebug* self, float t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithDouble(QDebug* self, double t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeft2(QDebug* self, const char* t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithQString(QDebug* self, struct miqt_string t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithQByteArray(QDebug* self, struct miqt_string t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithQByteArrayView(QDebug* self, QByteArrayView* t);
extern __declspec(dllexport) QDebug* QDebug_OperatorShiftLeftWithVoid(QDebug* self, const void* t);
extern __declspec(dllexport) QDebug* QDebug_MaybeQuote1(QDebug* self, char c);
extern __declspec(dllexport) void QDebug_Delete(QDebug* self, bool isSubclass);

extern __declspec(dllexport) QDebugStateSaver* QDebugStateSaver_new(QDebug* dbg);
extern __declspec(dllexport) void QDebugStateSaver_Delete(QDebugStateSaver* self, bool isSubclass);

extern __declspec(dllexport) QNoDebug* QNoDebug_Space(QNoDebug* self);
extern __declspec(dllexport) QNoDebug* QNoDebug_Nospace(QNoDebug* self);
extern __declspec(dllexport) QNoDebug* QNoDebug_MaybeSpace(QNoDebug* self);
extern __declspec(dllexport) QNoDebug* QNoDebug_Quote(QNoDebug* self);
extern __declspec(dllexport) QNoDebug* QNoDebug_Noquote(QNoDebug* self);
extern __declspec(dllexport) QNoDebug* QNoDebug_MaybeQuote(QNoDebug* self);
extern __declspec(dllexport) QNoDebug* QNoDebug_Verbosity(QNoDebug* self, int param1);
extern __declspec(dllexport) QNoDebug* QNoDebug_MaybeQuote1(QNoDebug* self, const char param1);
extern __declspec(dllexport) void QNoDebug_Delete(QNoDebug* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
