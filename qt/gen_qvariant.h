#pragma once
#ifndef MIQT_QT_GEN_QVARIANT_H
#define MIQT_QT_GEN_QVARIANT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QBitArray;
class QChar;
class QDataStream;
class QDate;
class QDateTime;
class QEasingCurve;
class QJsonArray;
class QJsonDocument;
class QJsonObject;
class QJsonValue;
class QLine;
class QLineF;
class QLocale;
class QMetaType;
class QModelIndex;
class QPartialOrdering;
class QPersistentModelIndex;
class QPoint;
class QPointF;
class QRect;
class QRectF;
class QRegularExpression;
class QSize;
class QSizeF;
class QTime;
class QUrl;
class QUuid;
class QVariant;
class QVariantConstPointer;
class _GUID;
class type_info;
#else
typedef struct QBitArray QBitArray;
typedef struct QChar QChar;
typedef struct QDataStream QDataStream;
typedef struct QDate QDate;
typedef struct QDateTime QDateTime;
typedef struct QEasingCurve QEasingCurve;
typedef struct QJsonArray QJsonArray;
typedef struct QJsonDocument QJsonDocument;
typedef struct QJsonObject QJsonObject;
typedef struct QJsonValue QJsonValue;
typedef struct QLine QLine;
typedef struct QLineF QLineF;
typedef struct QLocale QLocale;
typedef struct QMetaType QMetaType;
typedef struct QModelIndex QModelIndex;
typedef struct QPartialOrdering QPartialOrdering;
typedef struct QPersistentModelIndex QPersistentModelIndex;
typedef struct QPoint QPoint;
typedef struct QPointF QPointF;
typedef struct QRect QRect;
typedef struct QRectF QRectF;
typedef struct QRegularExpression QRegularExpression;
typedef struct QSize QSize;
typedef struct QSizeF QSizeF;
typedef struct QTime QTime;
typedef struct QUrl QUrl;
typedef struct QUuid QUuid;
typedef struct QVariant QVariant;
typedef struct QVariantConstPointer QVariantConstPointer;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QVariant* QVariant_new();
extern __declspec(dllexport) QVariant* QVariant_new2(QMetaType* typeVal);
extern __declspec(dllexport) QVariant* QVariant_new3(QVariant* other);
extern __declspec(dllexport) QVariant* QVariant_new4(int i);
extern __declspec(dllexport) QVariant* QVariant_new5(unsigned int ui);
extern __declspec(dllexport) QVariant* QVariant_new6(long long ll);
extern __declspec(dllexport) QVariant* QVariant_new7(unsigned long long ull);
extern __declspec(dllexport) QVariant* QVariant_new8(bool b);
extern __declspec(dllexport) QVariant* QVariant_new9(double d);
extern __declspec(dllexport) QVariant* QVariant_new10(float f);
extern __declspec(dllexport) QVariant* QVariant_new11(QChar* qchar);
extern __declspec(dllexport) QVariant* QVariant_new12(QDate* date);
extern __declspec(dllexport) QVariant* QVariant_new13(QTime* time);
extern __declspec(dllexport) QVariant* QVariant_new14(QBitArray* bitarray);
extern __declspec(dllexport) QVariant* QVariant_new15(struct miqt_string bytearray);
extern __declspec(dllexport) QVariant* QVariant_new16(QDateTime* datetime);
extern __declspec(dllexport) QVariant* QVariant_new17(struct miqt_map /* of struct miqt_string to QVariant* */  hash);
extern __declspec(dllexport) QVariant* QVariant_new18(QJsonArray* jsonArray);
extern __declspec(dllexport) QVariant* QVariant_new19(QJsonObject* jsonObject);
extern __declspec(dllexport) QVariant* QVariant_new20(QLocale* locale);
extern __declspec(dllexport) QVariant* QVariant_new21(struct miqt_map /* of struct miqt_string to QVariant* */  mapVal);
extern __declspec(dllexport) QVariant* QVariant_new22(QRegularExpression* re);
extern __declspec(dllexport) QVariant* QVariant_new23(struct miqt_string stringVal);
extern __declspec(dllexport) QVariant* QVariant_new24(struct miqt_array /* of struct miqt_string */  stringlist);
extern __declspec(dllexport) QVariant* QVariant_new25(QUrl* url);
extern __declspec(dllexport) QVariant* QVariant_new26(QSize* size);
extern __declspec(dllexport) QVariant* QVariant_new27(QPoint* pt);
extern __declspec(dllexport) QVariant* QVariant_new28(Type typeVal);
extern __declspec(dllexport) QVariant* QVariant_new29(QMetaType* typeVal, const void* copyVal);
extern __declspec(dllexport) void QVariant_OperatorAssign(QVariant* self, QVariant* other);
extern __declspec(dllexport) void QVariant_Swap(QVariant* self, QVariant* other);
extern __declspec(dllexport) int QVariant_UserType(const QVariant* self);
extern __declspec(dllexport) int QVariant_TypeId(const QVariant* self);
extern __declspec(dllexport) const char* QVariant_TypeName(const QVariant* self);
extern __declspec(dllexport) QMetaType* QVariant_MetaType(const QVariant* self);
extern __declspec(dllexport) bool QVariant_CanConvert(const QVariant* self, QMetaType* targetType);
extern __declspec(dllexport) bool QVariant_Convert(QVariant* self, QMetaType* typeVal);
extern __declspec(dllexport) bool QVariant_CanView(const QVariant* self, QMetaType* targetType);
extern __declspec(dllexport) bool QVariant_CanConvertWithTargetTypeId(const QVariant* self, int targetTypeId);
extern __declspec(dllexport) bool QVariant_ConvertWithTargetTypeId(QVariant* self, int targetTypeId);
extern __declspec(dllexport) bool QVariant_IsValid(const QVariant* self);
extern __declspec(dllexport) bool QVariant_IsNull(const QVariant* self);
extern __declspec(dllexport) void QVariant_Clear(QVariant* self);
extern __declspec(dllexport) void QVariant_Detach(QVariant* self);
extern __declspec(dllexport) bool QVariant_IsDetached(const QVariant* self);
extern __declspec(dllexport) int QVariant_ToInt(const QVariant* self);
extern __declspec(dllexport) unsigned int QVariant_ToUInt(const QVariant* self);
extern __declspec(dllexport) long long QVariant_ToLongLong(const QVariant* self);
extern __declspec(dllexport) unsigned long long QVariant_ToULongLong(const QVariant* self);
extern __declspec(dllexport) bool QVariant_ToBool(const QVariant* self);
extern __declspec(dllexport) double QVariant_ToDouble(const QVariant* self);
extern __declspec(dllexport) float QVariant_ToFloat(const QVariant* self);
extern __declspec(dllexport) double QVariant_ToReal(const QVariant* self);
extern __declspec(dllexport) struct miqt_string QVariant_ToByteArray(const QVariant* self);
extern __declspec(dllexport) QBitArray* QVariant_ToBitArray(const QVariant* self);
extern __declspec(dllexport) struct miqt_string QVariant_ToString(const QVariant* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QVariant_ToStringList(const QVariant* self);
extern __declspec(dllexport) QChar* QVariant_ToChar(const QVariant* self);
extern __declspec(dllexport) QDate* QVariant_ToDate(const QVariant* self);
extern __declspec(dllexport) QTime* QVariant_ToTime(const QVariant* self);
extern __declspec(dllexport) QDateTime* QVariant_ToDateTime(const QVariant* self);
extern __declspec(dllexport) struct miqt_map /* of struct miqt_string to QVariant* */  QVariant_ToMap(const QVariant* self);
extern __declspec(dllexport) struct miqt_map /* of struct miqt_string to QVariant* */  QVariant_ToHash(const QVariant* self);
extern __declspec(dllexport) QPoint* QVariant_ToPoint(const QVariant* self);
extern __declspec(dllexport) QPointF* QVariant_ToPointF(const QVariant* self);
extern __declspec(dllexport) QRect* QVariant_ToRect(const QVariant* self);
extern __declspec(dllexport) QSize* QVariant_ToSize(const QVariant* self);
extern __declspec(dllexport) QSizeF* QVariant_ToSizeF(const QVariant* self);
extern __declspec(dllexport) QLine* QVariant_ToLine(const QVariant* self);
extern __declspec(dllexport) QLineF* QVariant_ToLineF(const QVariant* self);
extern __declspec(dllexport) QRectF* QVariant_ToRectF(const QVariant* self);
extern __declspec(dllexport) QLocale* QVariant_ToLocale(const QVariant* self);
extern __declspec(dllexport) QRegularExpression* QVariant_ToRegularExpression(const QVariant* self);
extern __declspec(dllexport) QEasingCurve* QVariant_ToEasingCurve(const QVariant* self);
extern __declspec(dllexport) QUuid* QVariant_ToUuid(const QVariant* self);
extern __declspec(dllexport) QUrl* QVariant_ToUrl(const QVariant* self);
extern __declspec(dllexport) QJsonValue* QVariant_ToJsonValue(const QVariant* self);
extern __declspec(dllexport) QJsonObject* QVariant_ToJsonObject(const QVariant* self);
extern __declspec(dllexport) QJsonArray* QVariant_ToJsonArray(const QVariant* self);
extern __declspec(dllexport) QJsonDocument* QVariant_ToJsonDocument(const QVariant* self);
extern __declspec(dllexport) QModelIndex* QVariant_ToModelIndex(const QVariant* self);
extern __declspec(dllexport) QPersistentModelIndex* QVariant_ToPersistentModelIndex(const QVariant* self);
extern __declspec(dllexport) void QVariant_Load(QVariant* self, QDataStream* ds);
extern __declspec(dllexport) void QVariant_Save(const QVariant* self, QDataStream* ds);
extern __declspec(dllexport) Type QVariant_Type(const QVariant* self);
extern __declspec(dllexport) const char* QVariant_TypeToName(int typeId);
extern __declspec(dllexport) Type QVariant_NameToType(const char* name);
extern __declspec(dllexport) void* QVariant_Data(QVariant* self);
extern __declspec(dllexport) const void* QVariant_ConstData(const QVariant* self);
extern __declspec(dllexport) const void* QVariant_Data2(const QVariant* self);
extern __declspec(dllexport) void QVariant_SetValue(QVariant* self, QVariant* avalue);
extern __declspec(dllexport) QVariant* QVariant_FromMetaType(QMetaType* typeVal);
extern __declspec(dllexport) QPartialOrdering* QVariant_Compare(QVariant* lhs, QVariant* rhs);
extern __declspec(dllexport) DataPtr* QVariant_DataPtr(QVariant* self);
extern __declspec(dllexport) const DataPtr* QVariant_DataPtr2(const QVariant* self);
extern __declspec(dllexport) int QVariant_ToInt1(const QVariant* self, bool* ok);
extern __declspec(dllexport) unsigned int QVariant_ToUInt1(const QVariant* self, bool* ok);
extern __declspec(dllexport) long long QVariant_ToLongLong1(const QVariant* self, bool* ok);
extern __declspec(dllexport) unsigned long long QVariant_ToULongLong1(const QVariant* self, bool* ok);
extern __declspec(dllexport) double QVariant_ToDouble1(const QVariant* self, bool* ok);
extern __declspec(dllexport) float QVariant_ToFloat1(const QVariant* self, bool* ok);
extern __declspec(dllexport) double QVariant_ToReal1(const QVariant* self, bool* ok);
extern __declspec(dllexport) QVariant* QVariant_FromMetaType2(QMetaType* typeVal, const void* copyVal);
extern __declspec(dllexport) void QVariant_Delete(QVariant* self, bool isSubclass);

extern __declspec(dllexport) QVariantConstPointer* QVariantConstPointer_new(QVariant* variant);
extern __declspec(dllexport) QVariantConstPointer* QVariantConstPointer_new2(QVariantConstPointer* param1);
extern __declspec(dllexport) QVariant* QVariantConstPointer_OperatorMultiply(const QVariantConstPointer* self);
extern __declspec(dllexport) QVariant* QVariantConstPointer_OperatorMinusGreater(const QVariantConstPointer* self);
extern __declspec(dllexport) void QVariantConstPointer_OperatorAssign(QVariantConstPointer* self, QVariantConstPointer* param1);
extern __declspec(dllexport) void QVariantConstPointer_Delete(QVariantConstPointer* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
