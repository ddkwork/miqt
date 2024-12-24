#pragma once
#ifndef MIQT_QT_GEN_QOBJECTDEFS_H
#define MIQT_QT_GEN_QOBJECTDEFS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QMetaObject__Connection)
typedef QMetaObject::Connection QMetaObject__Connection;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QMetaObject__Data)
typedef QMetaObject::Data QMetaObject__Data;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QMetaObject__SuperData)
typedef QMetaObject::SuperData QMetaObject__SuperData;
typedef struct QGenericArgument QGenericArgument;
typedef struct QGenericReturnArgument QGenericReturnArgument;
typedef struct QMetaClassInfo QMetaClassInfo;
typedef struct QMetaEnum QMetaEnum;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaMethodArgument QMetaMethodArgument;
typedef struct QMetaMethodReturnArgument QMetaMethodReturnArgument;
typedef struct QMetaObject QMetaObject;
typedef struct QMetaObject__Connection QMetaObject__Connection;
typedef struct QMetaObject__Data QMetaObject__Data;
typedef struct QMetaObject__SuperData QMetaObject__SuperData;
typedef struct QMetaProperty QMetaProperty;
typedef struct QMetaType QMetaType;
typedef struct QMethodRawArguments QMethodRawArguments;
typedef struct QObject QObject;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QMethodRawArguments_Delete(QMethodRawArguments* self, bool isSubclass);

extern __declspec(dllexport) QGenericArgument* QGenericArgument_new();
extern __declspec(dllexport) QGenericArgument* QGenericArgument_new2(QGenericArgument* param1);
extern __declspec(dllexport) QGenericArgument* QGenericArgument_new3(const char* aName);
extern __declspec(dllexport) QGenericArgument* QGenericArgument_new4(const char* aName, const void* aData);
extern __declspec(dllexport) void* QGenericArgument_Data(const QGenericArgument* self);
extern __declspec(dllexport) const char* QGenericArgument_Name(const QGenericArgument* self);
extern __declspec(dllexport) void QGenericArgument_Delete(QGenericArgument* self, bool isSubclass);

extern __declspec(dllexport) QGenericReturnArgument* QGenericReturnArgument_new();
extern __declspec(dllexport) QGenericReturnArgument* QGenericReturnArgument_new2(QGenericReturnArgument* param1);
extern __declspec(dllexport) QGenericReturnArgument* QGenericReturnArgument_new3(const char* aName);
extern __declspec(dllexport) QGenericReturnArgument* QGenericReturnArgument_new4(const char* aName, void* aData);
extern __declspec(dllexport) void QGenericReturnArgument_virtbase(QGenericReturnArgument* src, QGenericArgument** outptr_QGenericArgument);
extern __declspec(dllexport) void QGenericReturnArgument_Delete(QGenericReturnArgument* self, bool isSubclass);

extern __declspec(dllexport) void QMetaMethodArgument_Delete(QMetaMethodArgument* self, bool isSubclass);

extern __declspec(dllexport) void QMetaMethodReturnArgument_Delete(QMetaMethodReturnArgument* self, bool isSubclass);

extern __declspec(dllexport) QMetaObject* QMetaObject_new();
extern __declspec(dllexport) QMetaObject* QMetaObject_new2(QMetaObject* param1);
extern __declspec(dllexport) const char* QMetaObject_ClassName(const QMetaObject* self);
extern __declspec(dllexport) QMetaObject* QMetaObject_SuperClass(const QMetaObject* self);
extern __declspec(dllexport) bool QMetaObject_Inherits(const QMetaObject* self, QMetaObject* metaObject);
extern __declspec(dllexport) QObject* QMetaObject_Cast(const QMetaObject* self, QObject* obj);
extern __declspec(dllexport) QObject* QMetaObject_CastWithObj(const QMetaObject* self, QObject* obj);
extern __declspec(dllexport) struct miqt_string QMetaObject_Tr(const QMetaObject* self, const char* s, const char* c);
extern __declspec(dllexport) QMetaType* QMetaObject_MetaType(const QMetaObject* self);
extern __declspec(dllexport) int QMetaObject_MethodOffset(const QMetaObject* self);
extern __declspec(dllexport) int QMetaObject_EnumeratorOffset(const QMetaObject* self);
extern __declspec(dllexport) int QMetaObject_PropertyOffset(const QMetaObject* self);
extern __declspec(dllexport) int QMetaObject_ClassInfoOffset(const QMetaObject* self);
extern __declspec(dllexport) int QMetaObject_ConstructorCount(const QMetaObject* self);
extern __declspec(dllexport) int QMetaObject_MethodCount(const QMetaObject* self);
extern __declspec(dllexport) int QMetaObject_EnumeratorCount(const QMetaObject* self);
extern __declspec(dllexport) int QMetaObject_PropertyCount(const QMetaObject* self);
extern __declspec(dllexport) int QMetaObject_ClassInfoCount(const QMetaObject* self);
extern __declspec(dllexport) int QMetaObject_IndexOfConstructor(const QMetaObject* self, const char* constructor);
extern __declspec(dllexport) int QMetaObject_IndexOfMethod(const QMetaObject* self, const char* method);
extern __declspec(dllexport) int QMetaObject_IndexOfSignal(const QMetaObject* self, const char* signal);
extern __declspec(dllexport) int QMetaObject_IndexOfSlot(const QMetaObject* self, const char* slot);
extern __declspec(dllexport) int QMetaObject_IndexOfEnumerator(const QMetaObject* self, const char* name);
extern __declspec(dllexport) int QMetaObject_IndexOfProperty(const QMetaObject* self, const char* name);
extern __declspec(dllexport) int QMetaObject_IndexOfClassInfo(const QMetaObject* self, const char* name);
extern __declspec(dllexport) QMetaMethod* QMetaObject_Constructor(const QMetaObject* self, int index);
extern __declspec(dllexport) QMetaMethod* QMetaObject_Method(const QMetaObject* self, int index);
extern __declspec(dllexport) QMetaEnum* QMetaObject_Enumerator(const QMetaObject* self, int index);
extern __declspec(dllexport) QMetaProperty* QMetaObject_Property(const QMetaObject* self, int index);
extern __declspec(dllexport) QMetaClassInfo* QMetaObject_ClassInfo(const QMetaObject* self, int index);
extern __declspec(dllexport) QMetaProperty* QMetaObject_UserProperty(const QMetaObject* self);
extern __declspec(dllexport) bool QMetaObject_CheckConnectArgs(const char* signal, const char* method);
extern __declspec(dllexport) bool QMetaObject_CheckConnectArgs2(QMetaMethod* signal, QMetaMethod* method);
extern __declspec(dllexport) struct miqt_string QMetaObject_NormalizedSignature(const char* method);
extern __declspec(dllexport) struct miqt_string QMetaObject_NormalizedType(const char* typeVal);
extern __declspec(dllexport) Connection QMetaObject_Connect(QObject* sender, int signal_index, QObject* receiver, int method_index);
extern __declspec(dllexport) bool QMetaObject_Disconnect(QObject* sender, int signal_index, QObject* receiver, int method_index);
extern __declspec(dllexport) bool QMetaObject_DisconnectOne(QObject* sender, int signal_index, QObject* receiver, int method_index);
extern __declspec(dllexport) void QMetaObject_ConnectSlotsByName(QObject* o);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod(QObject* obj, const char* member, int param3, QGenericReturnArgument* retVal);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod2(QObject* obj, const char* member, QGenericReturnArgument* retVal);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod3(QObject* obj, const char* member, int typeVal, QGenericArgument* val0);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod4(QObject* obj, const char* member, QGenericArgument* val0);
extern __declspec(dllexport) QObject* QMetaObject_NewInstance(const QMetaObject* self, QGenericArgument* val0);
extern __declspec(dllexport) struct miqt_string QMetaObject_Tr3(const QMetaObject* self, const char* s, const char* c, int n);
extern __declspec(dllexport) Connection QMetaObject_Connect5(QObject* sender, int signal_index, QObject* receiver, int method_index, int typeVal);
extern __declspec(dllexport) Connection QMetaObject_Connect6(QObject* sender, int signal_index, QObject* receiver, int method_index, int typeVal, int* types);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod5(QObject* obj, const char* member, int param3, QGenericReturnArgument* retVal, QGenericArgument* val0);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod6(QObject* obj, const char* member, int param3, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod7(QObject* obj, const char* member, int param3, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod8(QObject* obj, const char* member, int param3, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod9(QObject* obj, const char* member, int param3, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod10(QObject* obj, const char* member, int param3, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod11(QObject* obj, const char* member, int param3, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod12(QObject* obj, const char* member, int param3, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod13(QObject* obj, const char* member, int param3, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod14(QObject* obj, const char* member, int param3, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8, QGenericArgument* val9);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod42(QObject* obj, const char* member, QGenericReturnArgument* retVal, QGenericArgument* val0);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod52(QObject* obj, const char* member, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod62(QObject* obj, const char* member, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod72(QObject* obj, const char* member, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod82(QObject* obj, const char* member, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod92(QObject* obj, const char* member, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod102(QObject* obj, const char* member, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod112(QObject* obj, const char* member, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod122(QObject* obj, const char* member, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod132(QObject* obj, const char* member, QGenericReturnArgument* retVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8, QGenericArgument* val9);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod53(QObject* obj, const char* member, int typeVal, QGenericArgument* val0, QGenericArgument* val1);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod63(QObject* obj, const char* member, int typeVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod73(QObject* obj, const char* member, int typeVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod83(QObject* obj, const char* member, int typeVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod93(QObject* obj, const char* member, int typeVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod103(QObject* obj, const char* member, int typeVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod113(QObject* obj, const char* member, int typeVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod123(QObject* obj, const char* member, int typeVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod133(QObject* obj, const char* member, int typeVal, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8, QGenericArgument* val9);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod43(QObject* obj, const char* member, QGenericArgument* val0, QGenericArgument* val1);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod54(QObject* obj, const char* member, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod64(QObject* obj, const char* member, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod74(QObject* obj, const char* member, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod84(QObject* obj, const char* member, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod94(QObject* obj, const char* member, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod104(QObject* obj, const char* member, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod114(QObject* obj, const char* member, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8);
extern __declspec(dllexport) bool QMetaObject_InvokeMethod124(QObject* obj, const char* member, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8, QGenericArgument* val9);
extern __declspec(dllexport) QObject* QMetaObject_NewInstance2(const QMetaObject* self, QGenericArgument* val0, QGenericArgument* val1);
extern __declspec(dllexport) QObject* QMetaObject_NewInstance3(const QMetaObject* self, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2);
extern __declspec(dllexport) QObject* QMetaObject_NewInstance4(const QMetaObject* self, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3);
extern __declspec(dllexport) QObject* QMetaObject_NewInstance5(const QMetaObject* self, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4);
extern __declspec(dllexport) QObject* QMetaObject_NewInstance6(const QMetaObject* self, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5);
extern __declspec(dllexport) QObject* QMetaObject_NewInstance7(const QMetaObject* self, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6);
extern __declspec(dllexport) QObject* QMetaObject_NewInstance8(const QMetaObject* self, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7);
extern __declspec(dllexport) QObject* QMetaObject_NewInstance9(const QMetaObject* self, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8);
extern __declspec(dllexport) QObject* QMetaObject_NewInstance10(const QMetaObject* self, QGenericArgument* val0, QGenericArgument* val1, QGenericArgument* val2, QGenericArgument* val3, QGenericArgument* val4, QGenericArgument* val5, QGenericArgument* val6, QGenericArgument* val7, QGenericArgument* val8, QGenericArgument* val9);
extern __declspec(dllexport) void QMetaObject_Delete(QMetaObject* self, bool isSubclass);

extern __declspec(dllexport) QMetaObject__Connection* QMetaObject__Connection_new();
extern __declspec(dllexport) QMetaObject__Connection* QMetaObject__Connection_new2(const Connection* other);
extern __declspec(dllexport) void QMetaObject__Connection_OperatorAssign(QMetaObject__Connection* self, const Connection* other);
extern __declspec(dllexport) void QMetaObject__Connection_Swap(QMetaObject__Connection* self, Connection* other);
extern __declspec(dllexport) void QMetaObject__Connection_Delete(QMetaObject__Connection* self, bool isSubclass);

extern __declspec(dllexport) QMetaObject__SuperData* QMetaObject__SuperData_new();
extern __declspec(dllexport) QMetaObject__SuperData* QMetaObject__SuperData_new2(QMetaObject* mo);
extern __declspec(dllexport) QMetaObject__SuperData* QMetaObject__SuperData_new3(Getter g);
extern __declspec(dllexport) QMetaObject__SuperData* QMetaObject__SuperData_new4(const SuperData* param1);
extern __declspec(dllexport) QMetaObject* QMetaObject__SuperData_OperatorMinusGreater(const QMetaObject__SuperData* self);
extern __declspec(dllexport) void QMetaObject__SuperData_OperatorAssign(QMetaObject__SuperData* self, const SuperData* param1);
extern __declspec(dllexport) void QMetaObject__SuperData_Delete(QMetaObject__SuperData* self, bool isSubclass);

extern __declspec(dllexport) QMetaObject__Data* QMetaObject__Data_new();
extern __declspec(dllexport) QMetaObject__Data* QMetaObject__Data_new2(const Data* param1);
extern __declspec(dllexport) void QMetaObject__Data_OperatorAssign(QMetaObject__Data* self, const Data* param1);
extern __declspec(dllexport) void QMetaObject__Data_Delete(QMetaObject__Data* self, bool isSubclass);

} 
