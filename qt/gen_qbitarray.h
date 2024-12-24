#pragma once
#ifndef MIQT_QT_GEN_QBITARRAY_H
#define MIQT_QT_GEN_QBITARRAY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QBitArray QBitArray;
typedef struct QBitRef QBitRef;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QBitArray* QBitArray_new();
extern __declspec(dllexport) 
QBitArray* QBitArray_new2(ptrdiff_t size);
extern __declspec(dllexport) 
QBitArray* QBitArray_new3(QBitArray* other);
extern __declspec(dllexport) 
QBitArray* QBitArray_new4(ptrdiff_t size, bool val);
extern __declspec(dllexport) 
void QBitArray_OperatorAssign(QBitArray* self, QBitArray* other);
extern __declspec(dllexport) 
void QBitArray_Swap(QBitArray* self, QBitArray* other);
extern __declspec(dllexport) 
ptrdiff_t QBitArray_Size(const QBitArray* self);
extern __declspec(dllexport) 
ptrdiff_t QBitArray_Count(const QBitArray* self);
extern __declspec(dllexport) 
ptrdiff_t QBitArray_CountWithOn(const QBitArray* self, bool on);
extern __declspec(dllexport) 
bool QBitArray_IsEmpty(const QBitArray* self);
extern __declspec(dllexport) 
bool QBitArray_IsNull(const QBitArray* self);
extern __declspec(dllexport) 
void QBitArray_Resize(QBitArray* self, ptrdiff_t size);
extern __declspec(dllexport) 
void QBitArray_Detach(QBitArray* self);
extern __declspec(dllexport) 
bool QBitArray_IsDetached(const QBitArray* self);
extern __declspec(dllexport) 
void QBitArray_Clear(QBitArray* self);
extern __declspec(dllexport) 
bool QBitArray_TestBit(const QBitArray* self, ptrdiff_t i);
extern __declspec(dllexport) 
void QBitArray_SetBit(QBitArray* self, ptrdiff_t i);
extern __declspec(dllexport) 
void QBitArray_SetBit2(QBitArray* self, ptrdiff_t i, bool val);
extern __declspec(dllexport) 
void QBitArray_ClearBit(QBitArray* self, ptrdiff_t i);
extern __declspec(dllexport) 
bool QBitArray_ToggleBit(QBitArray* self, ptrdiff_t i);
extern __declspec(dllexport) 
bool QBitArray_At(const QBitArray* self, ptrdiff_t i);
extern __declspec(dllexport) 
QBitRef* QBitArray_OperatorSubscript(QBitArray* self, ptrdiff_t i);
extern __declspec(dllexport) 
bool QBitArray_OperatorSubscriptWithQsizetype(const QBitArray* self, ptrdiff_t i);
extern __declspec(dllexport) 
void QBitArray_OperatorBitwiseAndAssign(QBitArray* self, QBitArray* param1);
extern __declspec(dllexport) 
void QBitArray_OperatorBitwiseOrAssign(QBitArray* self, QBitArray* param1);
extern __declspec(dllexport) 
void QBitArray_OperatorBitwiseNotAssign(QBitArray* self, QBitArray* param1);
extern __declspec(dllexport) 
bool QBitArray_Fill(QBitArray* self, bool aval);
extern __declspec(dllexport) 
void QBitArray_Fill2(QBitArray* self, bool val, ptrdiff_t first, ptrdiff_t last);
extern __declspec(dllexport) 
void QBitArray_Truncate(QBitArray* self, ptrdiff_t pos);
extern __declspec(dllexport) 
const char* QBitArray_Bits(const QBitArray* self);
extern __declspec(dllexport) 
QBitArray* QBitArray_FromBits(const char* data, ptrdiff_t lenVal);
extern __declspec(dllexport) 
unsigned int QBitArray_ToUInt32(const QBitArray* self, int endianness);
extern __declspec(dllexport) 
DataPtr* QBitArray_DataPtr(QBitArray* self);
extern __declspec(dllexport) 
const DataPtr* QBitArray_DataPtr2(const QBitArray* self);
extern __declspec(dllexport) 
bool QBitArray_Fill22(QBitArray* self, bool aval, ptrdiff_t asize);
extern __declspec(dllexport) 
unsigned int QBitArray_ToUInt322(const QBitArray* self, int endianness, bool* ok);
extern __declspec(dllexport) 
void QBitArray_Delete(QBitArray* self, bool isSubclass);

extern __declspec(dllexport) 
QBitRef* QBitRef_new(QBitRef* param1);
extern __declspec(dllexport) 
bool QBitRef_OperatorNot(const QBitRef* self);
extern __declspec(dllexport) 
void QBitRef_OperatorAssign(QBitRef* self, QBitRef* val);
extern __declspec(dllexport) 
void QBitRef_OperatorAssignWithVal(QBitRef* self, bool val);
extern __declspec(dllexport) 
void QBitRef_Delete(QBitRef* self, bool isSubclass);

}
