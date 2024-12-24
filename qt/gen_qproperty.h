#pragma once
#ifndef MIQT_QT_GEN_QPROPERTY_H
#define MIQT_QT_GEN_QPROPERTY_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaType QMetaType;
typedef struct QPropertyBindingError QPropertyBindingError;
typedef struct QPropertyBindingSourceLocation QPropertyBindingSourceLocation;
typedef struct QPropertyNotifier QPropertyNotifier;
typedef struct QPropertyObserver QPropertyObserver;
typedef struct QPropertyObserverBase QPropertyObserverBase;
typedef struct QScopedPropertyUpdateGroup QScopedPropertyUpdateGroup;
typedef struct QUntypedBindable QUntypedBindable;
typedef struct QUntypedPropertyBinding QUntypedPropertyBinding;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QScopedPropertyUpdateGroup* QScopedPropertyUpdateGroup_new();
extern __declspec(dllexport) 
void QScopedPropertyUpdateGroup_Delete(QScopedPropertyUpdateGroup* self, bool isSubclass);

extern __declspec(dllexport) 
QPropertyBindingSourceLocation* QPropertyBindingSourceLocation_new();
extern __declspec(dllexport) 
QPropertyBindingSourceLocation* QPropertyBindingSourceLocation_new2(QPropertyBindingSourceLocation* param1);
extern __declspec(dllexport) 
void QPropertyBindingSourceLocation_Delete(QPropertyBindingSourceLocation* self, bool isSubclass);

extern __declspec(dllexport) 
QPropertyBindingError* QPropertyBindingError_new();
extern __declspec(dllexport) 
QPropertyBindingError* QPropertyBindingError_new2(Type typeVal);
extern __declspec(dllexport) 
QPropertyBindingError* QPropertyBindingError_new3(QPropertyBindingError* other);
extern __declspec(dllexport) 
QPropertyBindingError* QPropertyBindingError_new4(Type typeVal, struct miqt_string description);
extern __declspec(dllexport) 
void QPropertyBindingError_OperatorAssign(QPropertyBindingError* self, QPropertyBindingError* other);
extern __declspec(dllexport) 
bool QPropertyBindingError_HasError(const QPropertyBindingError* self);
extern __declspec(dllexport) 
Type QPropertyBindingError_Type(const QPropertyBindingError* self);
extern __declspec(dllexport) 
struct miqt_string QPropertyBindingError_Description(const QPropertyBindingError* self);
extern __declspec(dllexport) 
void QPropertyBindingError_Delete(QPropertyBindingError* self, bool isSubclass);

extern __declspec(dllexport) 
QUntypedPropertyBinding* QUntypedPropertyBinding_new();
extern __declspec(dllexport) 
QUntypedPropertyBinding* QUntypedPropertyBinding_new2(QMetaType* metaType, const BindingFunctionVTable* vtable, void* function, QPropertyBindingSourceLocation* location);
extern __declspec(dllexport) 
QUntypedPropertyBinding* QUntypedPropertyBinding_new3(QUntypedPropertyBinding* other);
extern __declspec(dllexport) 
void QUntypedPropertyBinding_OperatorAssign(QUntypedPropertyBinding* self, QUntypedPropertyBinding* other);
extern __declspec(dllexport) 
bool QUntypedPropertyBinding_IsNull(const QUntypedPropertyBinding* self);
extern __declspec(dllexport) 
QPropertyBindingError* QUntypedPropertyBinding_Error(const QUntypedPropertyBinding* self);
extern __declspec(dllexport) 
QMetaType* QUntypedPropertyBinding_ValueMetaType(const QUntypedPropertyBinding* self);
extern __declspec(dllexport) 
void QUntypedPropertyBinding_Delete(QUntypedPropertyBinding* self, bool isSubclass);

extern __declspec(dllexport) 
QPropertyObserverBase* QPropertyObserverBase_new();
extern __declspec(dllexport) 
QPropertyObserverBase* QPropertyObserverBase_new2(QPropertyObserverBase* param1);
extern __declspec(dllexport) 
void QPropertyObserverBase_Delete(QPropertyObserverBase* self, bool isSubclass);

extern __declspec(dllexport) 
QPropertyObserver* QPropertyObserver_new();
extern __declspec(dllexport) 
void QPropertyObserver_virtbase(QPropertyObserver* src
, QPropertyObserverBase** outptr_QPropertyObserverBase
);
extern __declspec(dllexport) 
void QPropertyObserver_Delete(QPropertyObserver* self, bool isSubclass);

extern __declspec(dllexport) 
QPropertyNotifier* QPropertyNotifier_new();
extern __declspec(dllexport) 
void QPropertyNotifier_virtbase(QPropertyNotifier* src
, QPropertyObserver** outptr_QPropertyObserver
);
extern __declspec(dllexport) 
void QPropertyNotifier_Delete(QPropertyNotifier* self, bool isSubclass);

extern __declspec(dllexport) 
QUntypedBindable* QUntypedBindable_new();
extern __declspec(dllexport) 
QUntypedBindable* QUntypedBindable_new2(QUntypedBindable* param1);
extern __declspec(dllexport) 
bool QUntypedBindable_IsValid(const QUntypedBindable* self);
extern __declspec(dllexport) 
bool QUntypedBindable_IsBindable(const QUntypedBindable* self);
extern __declspec(dllexport) 
bool QUntypedBindable_IsReadOnly(const QUntypedBindable* self);
extern __declspec(dllexport) 
QUntypedPropertyBinding* QUntypedBindable_MakeBinding(const QUntypedBindable* self);
extern __declspec(dllexport) 
QUntypedPropertyBinding* QUntypedBindable_TakeBinding(QUntypedBindable* self);
extern __declspec(dllexport) 
void QUntypedBindable_Observe(const QUntypedBindable* self, QPropertyObserver* observer);
extern __declspec(dllexport) 
QUntypedPropertyBinding* QUntypedBindable_Binding(const QUntypedBindable* self);
extern __declspec(dllexport) 
bool QUntypedBindable_SetBinding(QUntypedBindable* self, QUntypedPropertyBinding* binding);
extern __declspec(dllexport) 
bool QUntypedBindable_HasBinding(const QUntypedBindable* self);
extern __declspec(dllexport) 
QMetaType* QUntypedBindable_MetaType(const QUntypedBindable* self);
extern __declspec(dllexport) 
QUntypedPropertyBinding* QUntypedBindable_MakeBinding1(const QUntypedBindable* self, QPropertyBindingSourceLocation* location);
extern __declspec(dllexport) 
void QUntypedBindable_Delete(QUntypedBindable* self, bool isSubclass);

}
