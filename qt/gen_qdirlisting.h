#pragma once
#ifndef MIQT_QT_GEN_QDIRLISTING_H
#define MIQT_QT_GEN_QDIRLISTING_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QDirListing__DirEntry)
typedef QDirListing::DirEntry QDirListing__DirEntry;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QDirListing__const_iterator)
typedef QDirListing::const_iterator QDirListing__const_iterator;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QDirListing__sentinel)
typedef QDirListing::sentinel QDirListing__sentinel;
typedef struct QDateTime QDateTime;
typedef struct QDirListing QDirListing;
typedef struct QDirListing__DirEntry QDirListing__DirEntry;
typedef struct QDirListing__const_iterator QDirListing__const_iterator;
typedef struct QDirListing__sentinel QDirListing__sentinel;
typedef struct QFileInfo QFileInfo;
typedef struct QTimeZone QTimeZone;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QDirListing* QDirListing_new(struct miqt_string path);
extern __declspec(dllexport) 
QDirListing* QDirListing_new2(struct miqt_string path, struct miqt_array /* of struct miqt_string */  nameFilters);
extern __declspec(dllexport) 
QDirListing* QDirListing_new3(struct miqt_string path, IteratorFlags flags);
extern __declspec(dllexport) 
QDirListing* QDirListing_new4(struct miqt_string path, struct miqt_array /* of struct miqt_string */  nameFilters, IteratorFlags flags);
extern __declspec(dllexport) 
void QDirListing_Swap(QDirListing* self, QDirListing* other);
extern __declspec(dllexport) 
struct miqt_string QDirListing_IteratorPath(const QDirListing* self);
extern __declspec(dllexport) 
IteratorFlags QDirListing_IteratorFlags(const QDirListing* self);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QDirListing_NameFilters(const QDirListing* self);
extern __declspec(dllexport) 
const_iterator QDirListing_Begin(const QDirListing* self);
extern __declspec(dllexport) 
const_iterator QDirListing_Cbegin(const QDirListing* self);
extern __declspec(dllexport) 
sentinel QDirListing_End(const QDirListing* self);
extern __declspec(dllexport) 
sentinel QDirListing_Cend(const QDirListing* self);
extern __declspec(dllexport) 
const_iterator QDirListing_ConstBegin(const QDirListing* self);
extern __declspec(dllexport) 
sentinel QDirListing_ConstEnd(const QDirListing* self);
extern __declspec(dllexport) 
void QDirListing_Delete(QDirListing* self, bool isSubclass);

extern __declspec(dllexport) 
QDirListing__DirEntry* QDirListing__DirEntry_new(const DirEntry* param1);
extern __declspec(dllexport) 
QDirListing__DirEntry* QDirListing__DirEntry_new2();
extern __declspec(dllexport) 
struct miqt_string QDirListing__DirEntry_FileName(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
struct miqt_string QDirListing__DirEntry_BaseName(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
struct miqt_string QDirListing__DirEntry_CompleteBaseName(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
struct miqt_string QDirListing__DirEntry_Suffix(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
struct miqt_string QDirListing__DirEntry_BundleName(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
struct miqt_string QDirListing__DirEntry_CompleteSuffix(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
struct miqt_string QDirListing__DirEntry_FilePath(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
bool QDirListing__DirEntry_IsDir(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
bool QDirListing__DirEntry_IsFile(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
bool QDirListing__DirEntry_IsSymLink(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
bool QDirListing__DirEntry_Exists(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
bool QDirListing__DirEntry_IsHidden(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
bool QDirListing__DirEntry_IsReadable(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
bool QDirListing__DirEntry_IsWritable(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
bool QDirListing__DirEntry_IsExecutable(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
QFileInfo* QDirListing__DirEntry_FileInfo(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
struct miqt_string QDirListing__DirEntry_CanonicalFilePath(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
struct miqt_string QDirListing__DirEntry_AbsoluteFilePath(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
struct miqt_string QDirListing__DirEntry_AbsolutePath(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
long long QDirListing__DirEntry_Size(const QDirListing__DirEntry* self);
extern __declspec(dllexport) 
QDateTime* QDirListing__DirEntry_BirthTime(const QDirListing__DirEntry* self, QTimeZone* tz);
extern __declspec(dllexport) 
QDateTime* QDirListing__DirEntry_MetadataChangeTime(const QDirListing__DirEntry* self, QTimeZone* tz);
extern __declspec(dllexport) 
QDateTime* QDirListing__DirEntry_LastModified(const QDirListing__DirEntry* self, QTimeZone* tz);
extern __declspec(dllexport) 
QDateTime* QDirListing__DirEntry_LastRead(const QDirListing__DirEntry* self, QTimeZone* tz);
extern __declspec(dllexport) 
QDateTime* QDirListing__DirEntry_FileTime(const QDirListing__DirEntry* self, int typeVal, QTimeZone* tz);
extern __declspec(dllexport) 
void QDirListing__DirEntry_OperatorAssign(QDirListing__DirEntry* self, const DirEntry* param1);
extern __declspec(dllexport) 
void QDirListing__DirEntry_Delete(QDirListing__DirEntry* self, bool isSubclass);

extern __declspec(dllexport) 
QDirListing__sentinel* QDirListing__sentinel_new();
extern __declspec(dllexport) 
QDirListing__sentinel* QDirListing__sentinel_new2(const sentinel* param1);
extern __declspec(dllexport) 
void QDirListing__sentinel_Delete(QDirListing__sentinel* self, bool isSubclass);

extern __declspec(dllexport) 
QDirListing__const_iterator* QDirListing__const_iterator_new();
extern __declspec(dllexport) 
reference QDirListing__const_iterator_OperatorMultiply(const QDirListing__const_iterator* self);
extern __declspec(dllexport) 
pointer QDirListing__const_iterator_OperatorMinusGreater(const QDirListing__const_iterator* self);
extern __declspec(dllexport) 
const_iterator* QDirListing__const_iterator_OperatorPlusPlus(QDirListing__const_iterator* self);
extern __declspec(dllexport) 
void QDirListing__const_iterator_OperatorPlusPlusWithInt(QDirListing__const_iterator* self, int param1);
extern __declspec(dllexport) 
void QDirListing__const_iterator_Delete(QDirListing__const_iterator* self, bool isSubclass);

}
