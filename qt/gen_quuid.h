#pragma once
#ifndef MIQT_QT_GEN_QUUID_H
#define MIQT_QT_GEN_QUUID_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QUuid__Id128Bytes)
typedef QUuid::Id128Bytes QUuid__Id128Bytes;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_Disambiguated_t)
typedef Qt::Disambiguated_t Disambiguated_t;
typedef struct QAnyStringView QAnyStringView;
typedef struct QByteArrayView QByteArrayView;
typedef struct QUuid QUuid;
typedef struct QUuid__Id128Bytes QUuid__Id128Bytes;
typedef struct Disambiguated_t Disambiguated_t;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) _GUID* _GUID_new();
extern __declspec(dllexport) _GUID* _GUID_new2(_GUID* param1);
extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) QUuid* QUuid_new();
extern __declspec(dllexport) QUuid* QUuid_new2(unsigned int l, uint16_t w1, uint16_t w2, unsigned char b1, unsigned char b2, unsigned char b3, unsigned char b4, unsigned char b5, unsigned char b6, unsigned char b7, unsigned char b8);
extern __declspec(dllexport) QUuid* QUuid_new3(Id128Bytes id128);
extern __declspec(dllexport) QUuid* QUuid_new4(QAnyStringView* stringVal);
extern __declspec(dllexport) QUuid* QUuid_new5(const struct _GUID* guid);
extern __declspec(dllexport) QUuid* QUuid_new6(QUuid* param1);
extern __declspec(dllexport) QUuid* QUuid_new7(Id128Bytes id128, int order);
extern __declspec(dllexport) QUuid* QUuid_FromString(QAnyStringView* stringVal);
extern __declspec(dllexport) struct miqt_string QUuid_ToString(const QUuid* self);
extern __declspec(dllexport) struct miqt_string QUuid_ToByteArray(const QUuid* self);
extern __declspec(dllexport) Id128Bytes QUuid_ToBytes(const QUuid* self);
extern __declspec(dllexport) struct miqt_string QUuid_ToRfc4122(const QUuid* self);
extern __declspec(dllexport) QUuid* QUuid_FromBytes(const void* bytes);
extern __declspec(dllexport) QUuid* QUuid_FromRfc4122(QByteArrayView* param1);
extern __declspec(dllexport) bool QUuid_IsNull(const QUuid* self);
extern __declspec(dllexport) void QUuid_OperatorAssign(QUuid* self, const struct _GUID* guid);
extern __declspec(dllexport) QUuid* QUuid_CreateUuid();
extern __declspec(dllexport) QUuid* QUuid_CreateUuidV5(QUuid* ns, QByteArrayView* baseData);
extern __declspec(dllexport) QUuid* QUuid_CreateUuidV3(QUuid* ns, QByteArrayView* baseData);
extern __declspec(dllexport) QUuid* QUuid_CreateUuidV7();
extern __declspec(dllexport) Variant QUuid_Variant(const QUuid* self);
extern __declspec(dllexport) Version QUuid_Version(const QUuid* self);
extern __declspec(dllexport) void QUuid_OperatorAssignWithQUuid(QUuid* self, QUuid* param1);
extern __declspec(dllexport) struct miqt_string QUuid_ToString1(const QUuid* self, StringFormat mode);
extern __declspec(dllexport) struct miqt_string QUuid_ToByteArray1(const QUuid* self, StringFormat mode);
extern __declspec(dllexport) Id128Bytes QUuid_ToBytes1(const QUuid* self, int order);
extern __declspec(dllexport) QUuid* QUuid_FromBytes2(const void* bytes, int order);
extern __declspec(dllexport) bool QUuid_IsNull1(const QUuid* self, Disambiguated_t* param1);
extern __declspec(dllexport) Variant QUuid_Variant1(const QUuid* self, Disambiguated_t* param1);
extern __declspec(dllexport) Version QUuid_Version1(const QUuid* self, Disambiguated_t* param1);
extern __declspec(dllexport) void QUuid_Delete(QUuid* self, bool isSubclass);

extern __declspec(dllexport) QUuid__Id128Bytes* QUuid__Id128Bytes_new();
extern __declspec(dllexport) QUuid__Id128Bytes* QUuid__Id128Bytes_new2(const Id128Bytes* param1);
extern __declspec(dllexport) void QUuid__Id128Bytes_Delete(QUuid__Id128Bytes* self, bool isSubclass);

} 
