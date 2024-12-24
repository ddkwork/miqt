#pragma once
#ifndef MIQT_QT_GEN_QRANDOM_H
#define MIQT_QT_GEN_QRANDOM_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QRandomGenerator QRandomGenerator;
typedef struct QRandomGenerator64 QRandomGenerator64;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QRandomGenerator* QRandomGenerator_new();
extern __declspec(dllexport) QRandomGenerator* QRandomGenerator_new2(const unsigned int* seedBuffer, ptrdiff_t lenVal);
extern __declspec(dllexport) QRandomGenerator* QRandomGenerator_new3(const unsigned int* begin, const unsigned int* end);
extern __declspec(dllexport) QRandomGenerator* QRandomGenerator_new4(QRandomGenerator* other);
extern __declspec(dllexport) QRandomGenerator* QRandomGenerator_new5(unsigned int seedValue);
extern __declspec(dllexport) void QRandomGenerator_OperatorAssign(QRandomGenerator* self, QRandomGenerator* other);
extern __declspec(dllexport) unsigned int QRandomGenerator_Generate(QRandomGenerator* self);
extern __declspec(dllexport) unsigned long long QRandomGenerator_Generate64(QRandomGenerator* self);
extern __declspec(dllexport) double QRandomGenerator_GenerateDouble(QRandomGenerator* self);
extern __declspec(dllexport) double QRandomGenerator_Bounded(QRandomGenerator* self, double highest);
extern __declspec(dllexport) unsigned int QRandomGenerator_BoundedWithHighest(QRandomGenerator* self, unsigned int highest);
extern __declspec(dllexport) unsigned int QRandomGenerator_Bounded2(QRandomGenerator* self, unsigned int lowest, unsigned int highest);
extern __declspec(dllexport) int QRandomGenerator_Bounded3(QRandomGenerator* self, int highest);
extern __declspec(dllexport) int QRandomGenerator_Bounded4(QRandomGenerator* self, int lowest, int highest);
extern __declspec(dllexport) unsigned long long QRandomGenerator_Bounded5(QRandomGenerator* self, unsigned long long highest);
extern __declspec(dllexport) unsigned long long QRandomGenerator_Bounded6(QRandomGenerator* self, unsigned long long lowest, unsigned long long highest);
extern __declspec(dllexport) long long QRandomGenerator_Bounded7(QRandomGenerator* self, long long highest);
extern __declspec(dllexport) long long QRandomGenerator_Bounded8(QRandomGenerator* self, long long lowest, long long highest);
extern __declspec(dllexport) long long QRandomGenerator_Bounded9(QRandomGenerator* self, int lowest, long long highest);
extern __declspec(dllexport) long long QRandomGenerator_Bounded10(QRandomGenerator* self, long long lowest, int highest);
extern __declspec(dllexport) unsigned long long QRandomGenerator_Bounded11(QRandomGenerator* self, unsigned int lowest, unsigned long long highest);
extern __declspec(dllexport) unsigned long long QRandomGenerator_Bounded12(QRandomGenerator* self, unsigned long long lowest, unsigned int highest);
extern __declspec(dllexport) void QRandomGenerator_Generate2(QRandomGenerator* self, unsigned int* begin, unsigned int* end);
extern __declspec(dllexport) result_type QRandomGenerator_OperatorCall(QRandomGenerator* self);
extern __declspec(dllexport) void QRandomGenerator_Seed(QRandomGenerator* self);
extern __declspec(dllexport) void QRandomGenerator_Discard(QRandomGenerator* self, unsigned long long z);
extern __declspec(dllexport) result_type QRandomGenerator_Min();
extern __declspec(dllexport) result_type QRandomGenerator_Max();
extern __declspec(dllexport) QRandomGenerator* QRandomGenerator_System();
extern __declspec(dllexport) QRandomGenerator* QRandomGenerator_Global();
extern __declspec(dllexport) QRandomGenerator* QRandomGenerator_SecurelySeeded();
extern __declspec(dllexport) void QRandomGenerator_Seed1(QRandomGenerator* self, unsigned int s);
extern __declspec(dllexport) void QRandomGenerator_Delete(QRandomGenerator* self, bool isSubclass);

extern __declspec(dllexport) QRandomGenerator64* QRandomGenerator64_new();
extern __declspec(dllexport) QRandomGenerator64* QRandomGenerator64_new2(const unsigned int* seedBuffer, ptrdiff_t lenVal);
extern __declspec(dllexport) QRandomGenerator64* QRandomGenerator64_new3(const unsigned int* begin, const unsigned int* end);
extern __declspec(dllexport) QRandomGenerator64* QRandomGenerator64_new4(QRandomGenerator* other);
extern __declspec(dllexport) QRandomGenerator64* QRandomGenerator64_new5(QRandomGenerator64* param1);
extern __declspec(dllexport) QRandomGenerator64* QRandomGenerator64_new6(unsigned int seedValue);
extern __declspec(dllexport) void QRandomGenerator64_virtbase(QRandomGenerator64* src, QRandomGenerator** outptr_QRandomGenerator);
extern __declspec(dllexport) unsigned long long QRandomGenerator64_Generate(QRandomGenerator64* self);
extern __declspec(dllexport) result_type QRandomGenerator64_OperatorCall(QRandomGenerator64* self);
extern __declspec(dllexport) void QRandomGenerator64_Discard(QRandomGenerator64* self, unsigned long long z);
extern __declspec(dllexport) result_type QRandomGenerator64_Min();
extern __declspec(dllexport) result_type QRandomGenerator64_Max();
extern __declspec(dllexport) QRandomGenerator64* QRandomGenerator64_System();
extern __declspec(dllexport) QRandomGenerator64* QRandomGenerator64_Global();
extern __declspec(dllexport) QRandomGenerator64* QRandomGenerator64_SecurelySeeded();
extern __declspec(dllexport) void QRandomGenerator64_OperatorAssign(QRandomGenerator64* self, QRandomGenerator64* param1);
extern __declspec(dllexport) void QRandomGenerator64_Delete(QRandomGenerator64* self, bool isSubclass);

} 
