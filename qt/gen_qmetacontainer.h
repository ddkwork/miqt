#pragma once
#ifndef MIQT_QT_GEN_QMETACONTAINER_H
#define MIQT_QT_GEN_QMETACONTAINER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QMetaAssociation;
class QMetaContainer;
class QMetaSequence;
class QMetaType;
class _GUID;
class type_info;
#else
typedef struct QMetaAssociation QMetaAssociation;
typedef struct QMetaContainer QMetaContainer;
typedef struct QMetaSequence QMetaSequence;
typedef struct QMetaType QMetaType;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QMetaContainer* QMetaContainer_new();
extern __declspec(dllexport) QMetaContainer* QMetaContainer_new2(QMetaContainer* param1);
extern __declspec(dllexport) bool QMetaContainer_HasInputIterator(const QMetaContainer* self);
extern __declspec(dllexport) bool QMetaContainer_HasForwardIterator(const QMetaContainer* self);
extern __declspec(dllexport) bool QMetaContainer_HasBidirectionalIterator(const QMetaContainer* self);
extern __declspec(dllexport) bool QMetaContainer_HasRandomAccessIterator(const QMetaContainer* self);
extern __declspec(dllexport) bool QMetaContainer_HasSize(const QMetaContainer* self);
extern __declspec(dllexport) ptrdiff_t QMetaContainer_Size(const QMetaContainer* self, const void* container);
extern __declspec(dllexport) bool QMetaContainer_CanClear(const QMetaContainer* self);
extern __declspec(dllexport) void QMetaContainer_Clear(const QMetaContainer* self, void* container);
extern __declspec(dllexport) bool QMetaContainer_HasIterator(const QMetaContainer* self);
extern __declspec(dllexport) void* QMetaContainer_Begin(const QMetaContainer* self, void* container);
extern __declspec(dllexport) void* QMetaContainer_End(const QMetaContainer* self, void* container);
extern __declspec(dllexport) void QMetaContainer_DestroyIterator(const QMetaContainer* self, const void* iterator);
extern __declspec(dllexport) bool QMetaContainer_CompareIterator(const QMetaContainer* self, const void* i, const void* j);
extern __declspec(dllexport) void QMetaContainer_CopyIterator(const QMetaContainer* self, void* target, const void* source);
extern __declspec(dllexport) void QMetaContainer_AdvanceIterator(const QMetaContainer* self, void* iterator, ptrdiff_t step);
extern __declspec(dllexport) ptrdiff_t QMetaContainer_DiffIterator(const QMetaContainer* self, const void* i, const void* j);
extern __declspec(dllexport) bool QMetaContainer_HasConstIterator(const QMetaContainer* self);
extern __declspec(dllexport) void* QMetaContainer_ConstBegin(const QMetaContainer* self, const void* container);
extern __declspec(dllexport) void* QMetaContainer_ConstEnd(const QMetaContainer* self, const void* container);
extern __declspec(dllexport) void QMetaContainer_DestroyConstIterator(const QMetaContainer* self, const void* iterator);
extern __declspec(dllexport) bool QMetaContainer_CompareConstIterator(const QMetaContainer* self, const void* i, const void* j);
extern __declspec(dllexport) void QMetaContainer_CopyConstIterator(const QMetaContainer* self, void* target, const void* source);
extern __declspec(dllexport) void QMetaContainer_AdvanceConstIterator(const QMetaContainer* self, void* iterator, ptrdiff_t step);
extern __declspec(dllexport) ptrdiff_t QMetaContainer_DiffConstIterator(const QMetaContainer* self, const void* i, const void* j);
extern __declspec(dllexport) void QMetaContainer_Delete(QMetaContainer* self, bool isSubclass);

extern __declspec(dllexport) QMetaSequence* QMetaSequence_new();
extern __declspec(dllexport) void QMetaSequence_virtbase(QMetaSequence* src, QMetaContainer** outptr_QMetaContainer);
extern __declspec(dllexport) QMetaType* QMetaSequence_ValueMetaType(const QMetaSequence* self);
extern __declspec(dllexport) bool QMetaSequence_IsSortable(const QMetaSequence* self);
extern __declspec(dllexport) bool QMetaSequence_CanAddValueAtBegin(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_AddValueAtBegin(const QMetaSequence* self, void* container, const void* value);
extern __declspec(dllexport) bool QMetaSequence_CanAddValueAtEnd(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_AddValueAtEnd(const QMetaSequence* self, void* container, const void* value);
extern __declspec(dllexport) bool QMetaSequence_CanRemoveValueAtBegin(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_RemoveValueAtBegin(const QMetaSequence* self, void* container);
extern __declspec(dllexport) bool QMetaSequence_CanRemoveValueAtEnd(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_RemoveValueAtEnd(const QMetaSequence* self, void* container);
extern __declspec(dllexport) bool QMetaSequence_CanGetValueAtIndex(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_ValueAtIndex(const QMetaSequence* self, const void* container, ptrdiff_t index, void* result);
extern __declspec(dllexport) bool QMetaSequence_CanSetValueAtIndex(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_SetValueAtIndex(const QMetaSequence* self, void* container, ptrdiff_t index, const void* value);
extern __declspec(dllexport) bool QMetaSequence_CanAddValue(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_AddValue(const QMetaSequence* self, void* container, const void* value);
extern __declspec(dllexport) bool QMetaSequence_CanRemoveValue(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_RemoveValue(const QMetaSequence* self, void* container);
extern __declspec(dllexport) bool QMetaSequence_CanGetValueAtIterator(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_ValueAtIterator(const QMetaSequence* self, const void* iterator, void* result);
extern __declspec(dllexport) bool QMetaSequence_CanSetValueAtIterator(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_SetValueAtIterator(const QMetaSequence* self, const void* iterator, const void* value);
extern __declspec(dllexport) bool QMetaSequence_CanInsertValueAtIterator(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_InsertValueAtIterator(const QMetaSequence* self, void* container, const void* iterator, const void* value);
extern __declspec(dllexport) bool QMetaSequence_CanEraseValueAtIterator(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_EraseValueAtIterator(const QMetaSequence* self, void* container, const void* iterator);
extern __declspec(dllexport) bool QMetaSequence_CanEraseRangeAtIterator(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_EraseRangeAtIterator(const QMetaSequence* self, void* container, const void* iterator1, const void* iterator2);
extern __declspec(dllexport) bool QMetaSequence_CanGetValueAtConstIterator(const QMetaSequence* self);
extern __declspec(dllexport) void QMetaSequence_ValueAtConstIterator(const QMetaSequence* self, const void* iterator, void* result);
extern __declspec(dllexport) void QMetaSequence_Delete(QMetaSequence* self, bool isSubclass);

extern __declspec(dllexport) QMetaAssociation* QMetaAssociation_new();
extern __declspec(dllexport) void QMetaAssociation_virtbase(QMetaAssociation* src, QMetaContainer** outptr_QMetaContainer);
extern __declspec(dllexport) QMetaType* QMetaAssociation_KeyMetaType(const QMetaAssociation* self);
extern __declspec(dllexport) QMetaType* QMetaAssociation_MappedMetaType(const QMetaAssociation* self);
extern __declspec(dllexport) bool QMetaAssociation_CanInsertKey(const QMetaAssociation* self);
extern __declspec(dllexport) void QMetaAssociation_InsertKey(const QMetaAssociation* self, void* container, const void* key);
extern __declspec(dllexport) bool QMetaAssociation_CanRemoveKey(const QMetaAssociation* self);
extern __declspec(dllexport) void QMetaAssociation_RemoveKey(const QMetaAssociation* self, void* container, const void* key);
extern __declspec(dllexport) bool QMetaAssociation_CanContainsKey(const QMetaAssociation* self);
extern __declspec(dllexport) bool QMetaAssociation_ContainsKey(const QMetaAssociation* self, const void* container, const void* key);
extern __declspec(dllexport) bool QMetaAssociation_CanGetMappedAtKey(const QMetaAssociation* self);
extern __declspec(dllexport) void QMetaAssociation_MappedAtKey(const QMetaAssociation* self, const void* container, const void* key, void* mapped);
extern __declspec(dllexport) bool QMetaAssociation_CanSetMappedAtKey(const QMetaAssociation* self);
extern __declspec(dllexport) void QMetaAssociation_SetMappedAtKey(const QMetaAssociation* self, void* container, const void* key, const void* mapped);
extern __declspec(dllexport) bool QMetaAssociation_CanGetKeyAtIterator(const QMetaAssociation* self);
extern __declspec(dllexport) void QMetaAssociation_KeyAtIterator(const QMetaAssociation* self, const void* iterator, void* key);
extern __declspec(dllexport) bool QMetaAssociation_CanGetKeyAtConstIterator(const QMetaAssociation* self);
extern __declspec(dllexport) void QMetaAssociation_KeyAtConstIterator(const QMetaAssociation* self, const void* iterator, void* key);
extern __declspec(dllexport) bool QMetaAssociation_CanGetMappedAtIterator(const QMetaAssociation* self);
extern __declspec(dllexport) void QMetaAssociation_MappedAtIterator(const QMetaAssociation* self, const void* iterator, void* mapped);
extern __declspec(dllexport) bool QMetaAssociation_CanGetMappedAtConstIterator(const QMetaAssociation* self);
extern __declspec(dllexport) void QMetaAssociation_MappedAtConstIterator(const QMetaAssociation* self, const void* iterator, void* mapped);
extern __declspec(dllexport) bool QMetaAssociation_CanSetMappedAtIterator(const QMetaAssociation* self);
extern __declspec(dllexport) void QMetaAssociation_SetMappedAtIterator(const QMetaAssociation* self, const void* iterator, const void* mapped);
extern __declspec(dllexport) bool QMetaAssociation_CanCreateIteratorAtKey(const QMetaAssociation* self);
extern __declspec(dllexport) void* QMetaAssociation_CreateIteratorAtKey(const QMetaAssociation* self, void* container, const void* key);
extern __declspec(dllexport) bool QMetaAssociation_CanCreateConstIteratorAtKey(const QMetaAssociation* self);
extern __declspec(dllexport) void* QMetaAssociation_CreateConstIteratorAtKey(const QMetaAssociation* self, const void* container, const void* key);
extern __declspec(dllexport) void QMetaAssociation_Delete(QMetaAssociation* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
