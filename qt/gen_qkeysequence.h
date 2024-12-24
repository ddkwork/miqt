#pragma once
#ifndef MIQT_QT_GEN_QKEYSEQUENCE_H
#define MIQT_QT_GEN_QKEYSEQUENCE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QKeyCombination QKeyCombination;
typedef struct QKeySequence QKeySequence;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QKeySequence* QKeySequence_new();
extern __declspec(dllexport) QKeySequence* QKeySequence_new2(struct miqt_string key);
extern __declspec(dllexport) QKeySequence* QKeySequence_new3(int k1);
extern __declspec(dllexport) QKeySequence* QKeySequence_new4(QKeyCombination* k1);
extern __declspec(dllexport) QKeySequence* QKeySequence_new5(QKeySequence* ks);
extern __declspec(dllexport) QKeySequence* QKeySequence_new6(StandardKey key);
extern __declspec(dllexport) QKeySequence* QKeySequence_new7(struct miqt_string key, SequenceFormat format);
extern __declspec(dllexport) QKeySequence* QKeySequence_new8(int k1, int k2);
extern __declspec(dllexport) QKeySequence* QKeySequence_new9(int k1, int k2, int k3);
extern __declspec(dllexport) QKeySequence* QKeySequence_new10(int k1, int k2, int k3, int k4);
extern __declspec(dllexport) QKeySequence* QKeySequence_new11(QKeyCombination* k1, QKeyCombination* k2);
extern __declspec(dllexport) QKeySequence* QKeySequence_new12(QKeyCombination* k1, QKeyCombination* k2, QKeyCombination* k3);
extern __declspec(dllexport) QKeySequence* QKeySequence_new13(QKeyCombination* k1, QKeyCombination* k2, QKeyCombination* k3, QKeyCombination* k4);
extern __declspec(dllexport) int QKeySequence_Count(const QKeySequence* self);
extern __declspec(dllexport) bool QKeySequence_IsEmpty(const QKeySequence* self);
extern __declspec(dllexport) struct miqt_string QKeySequence_ToString(const QKeySequence* self);
extern __declspec(dllexport) QKeySequence* QKeySequence_FromString(struct miqt_string str);
extern __declspec(dllexport) struct miqt_array /* of QKeySequence* */  QKeySequence_ListFromString(struct miqt_string str);
extern __declspec(dllexport) struct miqt_string QKeySequence_ListToString(struct miqt_array /* of QKeySequence* */  list);
extern __declspec(dllexport) SequenceMatch QKeySequence_Matches(const QKeySequence* self, QKeySequence* seq);
extern __declspec(dllexport) QKeySequence* QKeySequence_Mnemonic(struct miqt_string text);
extern __declspec(dllexport) struct miqt_array /* of QKeySequence* */  QKeySequence_KeyBindings(StandardKey key);
extern __declspec(dllexport) QKeyCombination* QKeySequence_OperatorSubscript(const QKeySequence* self, unsigned int i);
extern __declspec(dllexport) void QKeySequence_OperatorAssign(QKeySequence* self, QKeySequence* other);
extern __declspec(dllexport) void QKeySequence_Swap(QKeySequence* self, QKeySequence* other);
extern __declspec(dllexport) bool QKeySequence_OperatorEqual(const QKeySequence* self, QKeySequence* other);
extern __declspec(dllexport) bool QKeySequence_OperatorNotEqual(const QKeySequence* self, QKeySequence* other);
extern __declspec(dllexport) bool QKeySequence_OperatorLesser(const QKeySequence* self, QKeySequence* ks);
extern __declspec(dllexport) bool QKeySequence_OperatorGreater(const QKeySequence* self, QKeySequence* other);
extern __declspec(dllexport) bool QKeySequence_OperatorLesserOrEqual(const QKeySequence* self, QKeySequence* other);
extern __declspec(dllexport) bool QKeySequence_OperatorGreaterOrEqual(const QKeySequence* self, QKeySequence* other);
extern __declspec(dllexport) bool QKeySequence_IsDetached(const QKeySequence* self);
extern __declspec(dllexport) DataPtr* QKeySequence_DataPtr(QKeySequence* self);
extern __declspec(dllexport) struct miqt_string QKeySequence_ToString1(const QKeySequence* self, SequenceFormat format);
extern __declspec(dllexport) QKeySequence* QKeySequence_FromString2(struct miqt_string str, SequenceFormat format);
extern __declspec(dllexport) struct miqt_array /* of QKeySequence* */  QKeySequence_ListFromString2(struct miqt_string str, SequenceFormat format);
extern __declspec(dllexport) struct miqt_string QKeySequence_ListToString2(struct miqt_array /* of QKeySequence* */  list, SequenceFormat format);
extern __declspec(dllexport) void QKeySequence_Delete(QKeySequence* self, bool isSubclass);

} 
