#pragma once
#ifndef MIQT_QT_GEN_QMETAOBJECT_H
#define MIQT_QT_GEN_QMETAOBJECT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QByteArrayView QByteArrayView;
typedef struct QGenericArgument QGenericArgument;
typedef struct QGenericReturnArgument QGenericReturnArgument;
typedef struct QMetaClassInfo QMetaClassInfo;
typedef struct QMetaEnum QMetaEnum;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QMetaProperty QMetaProperty;
typedef struct QMetaType QMetaType;
typedef struct QObject QObject;
typedef struct QUntypedBindable QUntypedBindable;
typedef struct QVariant QVariant;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QMetaMethod* QMetaMethod_new();
extern __declspec(dllexport) QMetaMethod* QMetaMethod_new2(QMetaMethod* param1);
extern __declspec(dllexport) struct miqt_string QMetaMethod_MethodSignature(const QMetaMethod* self);
extern __declspec(dllexport) struct miqt_string QMetaMethod_Name(const QMetaMethod* self);
extern __declspec(dllexport) QByteArrayView* QMetaMethod_NameView(const QMetaMethod* self);
extern __declspec(dllexport) const char* QMetaMethod_TypeName(const QMetaMethod* self);
extern __declspec(dllexport) int QMetaMethod_ReturnType(const QMetaMethod* self);
extern __declspec(dllexport) QMetaType* QMetaMethod_ReturnMetaType(const QMetaMethod* self);
extern __declspec(dllexport) int QMetaMethod_ParameterCount(const QMetaMethod* self);
extern __declspec(dllexport) int QMetaMethod_ParameterType(const QMetaMethod* self, int index);
extern __declspec(dllexport) QMetaType* QMetaMethod_ParameterMetaType(const QMetaMethod* self, int index);
extern __declspec(dllexport) void QMetaMethod_GetParameterTypes(const QMetaMethod* self, int* types);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QMetaMethod_ParameterTypes(const QMetaMethod* self);
extern __declspec(dllexport) struct miqt_string QMetaMethod_ParameterTypeName(const QMetaMethod* self, int index);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QMetaMethod_ParameterNames(const QMetaMethod* self);
extern __declspec(dllexport) const char* QMetaMethod_Tag(const QMetaMethod* self);
extern __declspec(dllexport) Access QMetaMethod_Access(const QMetaMethod* self);
extern __declspec(dllexport) MethodType QMetaMethod_MethodType(const QMetaMethod* self);
extern __declspec(dllexport) int QMetaMethod_Attributes(const QMetaMethod* self);
extern __declspec(dllexport) int QMetaMethod_MethodIndex(const QMetaMethod* self);
extern __declspec(dllexport) int QMetaMethod_RelativeMethodIndex(const QMetaMethod* self);
extern __declspec(dllexport) int QMetaMethod_Revision(const QMetaMethod* self);
extern __declspec(dllexport) bool QMetaMethod_IsConst(const QMetaMethod* self);
extern __declspec(dllexport) QMetaObject* QMetaMethod_EnclosingMetaObject(const QMetaMethod* self);
extern __declspec(dllexport) bool QMetaMethod_Invoke(const QMetaMethod* self, QObject* object, int connectionType, QGenericReturnArgument* returnValue);
extern __declspec(dllexport) bool QMetaMethod_Invoke2(const QMetaMethod* self, QObject* object, QGenericReturnArgument* returnValue);
extern __declspec(dllexport) bool QMetaMethod_Invoke3(const QMetaMethod* self, QObject* object, int connectionType, QGenericArgument* val0);
extern __declspec(dllexport) bool QMetaMethod_Invoke4(const QMetaMethod* self, QObject* object, QGenericArgument* val0);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget(const QMetaMethod* self, void* gadget, QGenericReturnArgument* returnValue);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget2(const QMetaMethod* self, void* gadget, QGenericArgument* val0);
extern __declspec(dllexport) bool QMetaMethod_IsValid(const QMetaMethod* self);
extern __declspec(dllexport) bool QMetaMethod_Invoke42(const QMetaMethod* self, QObject* object, int connectionType, QGenericReturnArgument* returnValue, QGenericArgument* val0);
extern __declspec(dllexport) bool QMetaMethod_Invoke5(const QMetaMethod* self, QObject* object, int connectionType, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1);
extern __declspec(dllexport) bool QMetaMethod_Invoke6(const QMetaMethod* self, QObject* object, int connectionType, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2);
extern __declspec(dllexport) bool QMetaMethod_Invoke7(const QMetaMethod* self, QObject* object, int connectionType, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3);
extern __declspec(dllexport) bool QMetaMethod_Invoke8(const QMetaMethod* self, QObject* object, int connectionType, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4);
extern __declspec(dllexport) bool QMetaMethod_Invoke9(const QMetaMethod* self, QObject* object, int connectionType, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5);
extern __declspec(dllexport) bool QMetaMethod_Invoke10(const QMetaMethod* self, QObject* object, int connectionType, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6);
extern __declspec(dllexport) bool QMetaMethod_Invoke11(const QMetaMethod* self, QObject* object, int connectionType, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7);
extern __declspec(dllexport) bool QMetaMethod_Invoke12(const QMetaMethod* self, QObject* object, int connectionType, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8);
extern __declspec(dllexport) bool QMetaMethod_Invoke13(const QMetaMethod* self, QObject* object, int connectionType, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8, QGenericArgument* val9);
extern __declspec(dllexport) bool QMetaMethod_Invoke32(const QMetaMethod* self, QObject* object, QGenericReturnArgument* returnValue, QGenericArgument* val0);
extern __declspec(dllexport) bool QMetaMethod_Invoke43(const QMetaMethod* self, QObject* object, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1);
extern __declspec(dllexport) bool QMetaMethod_Invoke52(const QMetaMethod* self, QObject* object, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2);
extern __declspec(dllexport) bool QMetaMethod_Invoke62(const QMetaMethod* self, QObject* object, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3);
extern __declspec(dllexport) bool QMetaMethod_Invoke72(const QMetaMethod* self, QObject* object, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4);
extern __declspec(dllexport) bool QMetaMethod_Invoke82(const QMetaMethod* self, QObject* object, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5);
extern __declspec(dllexport) bool QMetaMethod_Invoke92(const QMetaMethod* self, QObject* object, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6);
extern __declspec(dllexport) bool QMetaMethod_Invoke102(const QMetaMethod* self, QObject* object, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7);
extern __declspec(dllexport) bool QMetaMethod_Invoke112(const QMetaMethod* self, QObject* object, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8);
extern __declspec(dllexport) bool QMetaMethod_Invoke122(const QMetaMethod* self, QObject* object, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8, QGenericArgument* val9);
extern __declspec(dllexport) bool QMetaMethod_Invoke44(const QMetaMethod* self, QObject* object, int connectionType, QGenericArgument* val0, QGenericArgument* val1);
extern __declspec(dllexport) bool QMetaMethod_Invoke53(const QMetaMethod* self, QObject* object, int connectionType, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2);
extern __declspec(dllexport) bool QMetaMethod_Invoke63(const QMetaMethod* self, QObject* object, int connectionType, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3);
extern __declspec(dllexport) bool QMetaMethod_Invoke73(const QMetaMethod* self, QObject* object, int connectionType, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4);
extern __declspec(dllexport) bool QMetaMethod_Invoke83(const QMetaMethod* self, QObject* object, int connectionType, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5);
extern __declspec(dllexport) bool QMetaMethod_Invoke93(const QMetaMethod* self, QObject* object, int connectionType, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6);
extern __declspec(dllexport) bool QMetaMethod_Invoke103(const QMetaMethod* self, QObject* object, int connectionType, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7);
extern __declspec(dllexport) bool QMetaMethod_Invoke113(const QMetaMethod* self, QObject* object, int connectionType, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8);
extern __declspec(dllexport) bool QMetaMethod_Invoke123(const QMetaMethod* self, QObject* object, int connectionType, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8, QGenericArgument* val9);
extern __declspec(dllexport) bool QMetaMethod_Invoke33(const QMetaMethod* self, QObject* object, QGenericArgument* val0, QGenericArgument* val1);
extern __declspec(dllexport) bool QMetaMethod_Invoke45(const QMetaMethod* self, QObject* object, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2);
extern __declspec(dllexport) bool QMetaMethod_Invoke54(const QMetaMethod* self, QObject* object, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3);
extern __declspec(dllexport) bool QMetaMethod_Invoke64(const QMetaMethod* self, QObject* object, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4);
extern __declspec(dllexport) bool QMetaMethod_Invoke74(const QMetaMethod* self, QObject* object, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5);
extern __declspec(dllexport) bool QMetaMethod_Invoke84(const QMetaMethod* self, QObject* object, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6);
extern __declspec(dllexport) bool QMetaMethod_Invoke94(const QMetaMethod* self, QObject* object, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7);
extern __declspec(dllexport) bool QMetaMethod_Invoke104(const QMetaMethod* self, QObject* object, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8);
extern __declspec(dllexport) bool QMetaMethod_Invoke114(const QMetaMethod* self, QObject* object, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8, QGenericArgument* val9);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget3(const QMetaMethod* self, void* gadget, QGenericReturnArgument* returnValue, QGenericArgument* val0);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget4(const QMetaMethod* self, void* gadget, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget5(const QMetaMethod* self, void* gadget, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget6(const QMetaMethod* self, void* gadget, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget7(const QMetaMethod* self, void* gadget, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget8(const QMetaMethod* self, void* gadget, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget9(const QMetaMethod* self, void* gadget, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget10(const QMetaMethod* self, void* gadget, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget11(const QMetaMethod* self, void* gadget, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget12(const QMetaMethod* self, void* gadget, QGenericReturnArgument* returnValue, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8, QGenericArgument* val9);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget32(const QMetaMethod* self, void* gadget, QGenericArgument* val0, QGenericArgument* val1);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget42(const QMetaMethod* self, void* gadget, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget52(const QMetaMethod* self, void* gadget, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget62(const QMetaMethod* self, void* gadget, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget72(const QMetaMethod* self, void* gadget, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget82(const QMetaMethod* self, void* gadget, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget92(const QMetaMethod* self, void* gadget, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget102(const QMetaMethod* self, void* gadget, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8);
extern __declspec(dllexport) bool QMetaMethod_InvokeOnGadget112(const QMetaMethod* self, void* gadget, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8, QGenericArgument* val9);
extern __declspec(dllexport) void QMetaMethod_Delete(QMetaMethod* self, bool isSubclass);

extern __declspec(dllexport) QMetaEnum* QMetaEnum_new();
extern __declspec(dllexport) QMetaEnum* QMetaEnum_new2(QMetaEnum* param1);
extern __declspec(dllexport) const char* QMetaEnum_Name(const QMetaEnum* self);
extern __declspec(dllexport) const char* QMetaEnum_EnumName(const QMetaEnum* self);
extern __declspec(dllexport) QMetaType* QMetaEnum_MetaType(const QMetaEnum* self);
extern __declspec(dllexport) bool QMetaEnum_IsFlag(const QMetaEnum* self);
extern __declspec(dllexport) bool QMetaEnum_IsScoped(const QMetaEnum* self);
extern __declspec(dllexport) bool QMetaEnum_Is64Bit(const QMetaEnum* self);
extern __declspec(dllexport) int QMetaEnum_KeyCount(const QMetaEnum* self);
extern __declspec(dllexport) const char* QMetaEnum_Key(const QMetaEnum* self, int index);
extern __declspec(dllexport) int QMetaEnum_Value(const QMetaEnum* self, int index);
extern __declspec(dllexport) const char* QMetaEnum_Scope(const QMetaEnum* self);
extern __declspec(dllexport) int QMetaEnum_KeyToValue(const QMetaEnum* self, const char* key);
extern __declspec(dllexport) int QMetaEnum_KeysToValue(const QMetaEnum* self, const char* keys);
extern __declspec(dllexport) const char* QMetaEnum_ValueToKey(const QMetaEnum* self, unsigned long long value);
extern __declspec(dllexport) struct miqt_string QMetaEnum_ValueToKeys(const QMetaEnum* self, unsigned long long value);
extern __declspec(dllexport) QMetaObject* QMetaEnum_EnclosingMetaObject(const QMetaEnum* self);
extern __declspec(dllexport) bool QMetaEnum_IsValid(const QMetaEnum* self);
extern __declspec(dllexport) int QMetaEnum_KeyToValue2(const QMetaEnum* self, const char* key, bool* ok);
extern __declspec(dllexport) int QMetaEnum_KeysToValue2(const QMetaEnum* self, const char* keys, bool* ok);
extern __declspec(dllexport) void QMetaEnum_Delete(QMetaEnum* self, bool isSubclass);

extern __declspec(dllexport) QMetaProperty* QMetaProperty_new();
extern __declspec(dllexport) const char* QMetaProperty_Name(const QMetaProperty* self);
extern __declspec(dllexport) const char* QMetaProperty_TypeName(const QMetaProperty* self);
extern __declspec(dllexport) int QMetaProperty_Type(const QMetaProperty* self);
extern __declspec(dllexport) int QMetaProperty_UserType(const QMetaProperty* self);
extern __declspec(dllexport) int QMetaProperty_TypeId(const QMetaProperty* self);
extern __declspec(dllexport) QMetaType* QMetaProperty_MetaType(const QMetaProperty* self);
extern __declspec(dllexport) int QMetaProperty_PropertyIndex(const QMetaProperty* self);
extern __declspec(dllexport) int QMetaProperty_RelativePropertyIndex(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsReadable(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsWritable(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsResettable(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsDesignable(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsScriptable(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsStored(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsUser(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsConstant(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsFinal(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsRequired(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsBindable(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsFlagType(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsEnumType(const QMetaProperty* self);
extern __declspec(dllexport) QMetaEnum* QMetaProperty_Enumerator(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_HasNotifySignal(const QMetaProperty* self);
extern __declspec(dllexport) QMetaMethod* QMetaProperty_NotifySignal(const QMetaProperty* self);
extern __declspec(dllexport) int QMetaProperty_NotifySignalIndex(const QMetaProperty* self);
extern __declspec(dllexport) int QMetaProperty_Revision(const QMetaProperty* self);
extern __declspec(dllexport) QVariant* QMetaProperty_Read(const QMetaProperty* self, QObject* obj);
extern __declspec(dllexport) bool QMetaProperty_Write(const QMetaProperty* self, QObject* obj, QVariant* value);
extern __declspec(dllexport) bool QMetaProperty_Reset(const QMetaProperty* self, QObject* obj);
extern __declspec(dllexport) QUntypedBindable* QMetaProperty_Bindable(const QMetaProperty* self, QObject* object);
extern __declspec(dllexport) QVariant* QMetaProperty_ReadOnGadget(const QMetaProperty* self, const void* gadget);
extern __declspec(dllexport) bool QMetaProperty_WriteOnGadget(const QMetaProperty* self, void* gadget, QVariant* value);
extern __declspec(dllexport) bool QMetaProperty_ResetOnGadget(const QMetaProperty* self, void* gadget);
extern __declspec(dllexport) bool QMetaProperty_HasStdCppSet(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsAlias(const QMetaProperty* self);
extern __declspec(dllexport) bool QMetaProperty_IsValid(const QMetaProperty* self);
extern __declspec(dllexport) QMetaObject* QMetaProperty_EnclosingMetaObject(const QMetaProperty* self);
extern __declspec(dllexport) void QMetaProperty_Delete(QMetaProperty* self, bool isSubclass);

extern __declspec(dllexport) QMetaClassInfo* QMetaClassInfo_new();
extern __declspec(dllexport) QMetaClassInfo* QMetaClassInfo_new2(QMetaClassInfo* param1);
extern __declspec(dllexport) const char* QMetaClassInfo_Name(const QMetaClassInfo* self);
extern __declspec(dllexport) const char* QMetaClassInfo_Value(const QMetaClassInfo* self);
extern __declspec(dllexport) QMetaObject* QMetaClassInfo_EnclosingMetaObject(const QMetaClassInfo* self);
extern __declspec(dllexport) void QMetaClassInfo_Delete(QMetaClassInfo* self, bool isSubclass);

} 
