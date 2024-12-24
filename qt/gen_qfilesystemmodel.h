#pragma once
#ifndef MIQT_QT_GEN_QFILESYSTEMMODEL_H
#define MIQT_QT_GEN_QFILESYSTEMMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractFileIconProvider QAbstractFileIconProvider;
typedef struct QAbstractItemModel QAbstractItemModel;
typedef struct QDateTime QDateTime;
typedef struct QDir QDir;
typedef struct QEvent QEvent;
typedef struct QFileInfo QFileInfo;
typedef struct QFileSystemModel QFileSystemModel;
typedef struct QIcon QIcon;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QTimeZone QTimeZone;
typedef struct QTimerEvent QTimerEvent;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QFileSystemModel* QFileSystemModel_new();
extern __declspec(dllexport) 
QFileSystemModel* QFileSystemModel_new2(QObject* parent);
extern __declspec(dllexport) 
void QFileSystemModel_virtbase(QFileSystemModel* src
, QAbstractItemModel** outptr_QAbstractItemModel
);
extern __declspec(dllexport) 
QMetaObject* QFileSystemModel_MetaObject(const QFileSystemModel* self);
extern __declspec(dllexport) 
void* QFileSystemModel_Metacast(QFileSystemModel* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QFileSystemModel_Tr(const char* s);
extern __declspec(dllexport) 
void QFileSystemModel_RootPathChanged(QFileSystemModel* self, struct miqt_string newPath);
void QFileSystemModel_connect_RootPathChanged(QFileSystemModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QFileSystemModel_FileRenamed(QFileSystemModel* self, struct miqt_string path, struct miqt_string oldName, struct miqt_string newName);
void QFileSystemModel_connect_FileRenamed(QFileSystemModel* self, intptr_t slot);
extern __declspec(dllexport) 
void QFileSystemModel_DirectoryLoaded(QFileSystemModel* self, struct miqt_string path);
void QFileSystemModel_connect_DirectoryLoaded(QFileSystemModel* self, intptr_t slot);
extern __declspec(dllexport) 
QModelIndex* QFileSystemModel_Index(const QFileSystemModel* self, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
QModelIndex* QFileSystemModel_IndexWithPath(const QFileSystemModel* self, struct miqt_string path);
extern __declspec(dllexport) 
QModelIndex* QFileSystemModel_Parent(const QFileSystemModel* self, QModelIndex* child);
extern __declspec(dllexport) 
QModelIndex* QFileSystemModel_Sibling(const QFileSystemModel* self, int row, int column, QModelIndex* idx);
extern __declspec(dllexport) 
bool QFileSystemModel_HasChildren(const QFileSystemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
bool QFileSystemModel_CanFetchMore(const QFileSystemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
void QFileSystemModel_FetchMore(QFileSystemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
int QFileSystemModel_RowCount(const QFileSystemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
int QFileSystemModel_ColumnCount(const QFileSystemModel* self, QModelIndex* parent);
extern __declspec(dllexport) 
QVariant* QFileSystemModel_MyComputer(const QFileSystemModel* self);
extern __declspec(dllexport) 
QVariant* QFileSystemModel_Data(const QFileSystemModel* self, QModelIndex* index, int role);
extern __declspec(dllexport) 
bool QFileSystemModel_SetData(QFileSystemModel* self, QModelIndex* index, QVariant* value, int role);
extern __declspec(dllexport) 
QVariant* QFileSystemModel_HeaderData(const QFileSystemModel* self, int section, int orientation, int role);
extern __declspec(dllexport) 
int QFileSystemModel_Flags(const QFileSystemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
void QFileSystemModel_Sort(QFileSystemModel* self, int column, int order);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QFileSystemModel_MimeTypes(const QFileSystemModel* self);
extern __declspec(dllexport) 
QMimeData* QFileSystemModel_MimeData(const QFileSystemModel* self, struct miqt_array /* of QModelIndex* */  indexes);
extern __declspec(dllexport) 
bool QFileSystemModel_DropMimeData(QFileSystemModel* self, QMimeData* data, int action, int row, int column, QModelIndex* parent);
extern __declspec(dllexport) 
int QFileSystemModel_SupportedDropActions(const QFileSystemModel* self);
extern __declspec(dllexport) 
struct miqt_map /* of int to struct miqt_string */  QFileSystemModel_RoleNames(const QFileSystemModel* self);
extern __declspec(dllexport) 
QModelIndex* QFileSystemModel_SetRootPath(QFileSystemModel* self, struct miqt_string path);
extern __declspec(dllexport) 
struct miqt_string QFileSystemModel_RootPath(const QFileSystemModel* self);
extern __declspec(dllexport) 
QDir* QFileSystemModel_RootDirectory(const QFileSystemModel* self);
extern __declspec(dllexport) 
void QFileSystemModel_SetIconProvider(QFileSystemModel* self, QAbstractFileIconProvider* provider);
extern __declspec(dllexport) 
QAbstractFileIconProvider* QFileSystemModel_IconProvider(const QFileSystemModel* self);
extern __declspec(dllexport) 
void QFileSystemModel_SetFilter(QFileSystemModel* self, int filters);
extern __declspec(dllexport) 
int QFileSystemModel_Filter(const QFileSystemModel* self);
extern __declspec(dllexport) 
void QFileSystemModel_SetResolveSymlinks(QFileSystemModel* self, bool enable);
extern __declspec(dllexport) 
bool QFileSystemModel_ResolveSymlinks(const QFileSystemModel* self);
extern __declspec(dllexport) 
void QFileSystemModel_SetReadOnly(QFileSystemModel* self, bool enable);
extern __declspec(dllexport) 
bool QFileSystemModel_IsReadOnly(const QFileSystemModel* self);
extern __declspec(dllexport) 
void QFileSystemModel_SetNameFilterDisables(QFileSystemModel* self, bool enable);
extern __declspec(dllexport) 
bool QFileSystemModel_NameFilterDisables(const QFileSystemModel* self);
extern __declspec(dllexport) 
void QFileSystemModel_SetNameFilters(QFileSystemModel* self, struct miqt_array /* of struct miqt_string */  filters);
extern __declspec(dllexport) 
struct miqt_array /* of struct miqt_string */  QFileSystemModel_NameFilters(const QFileSystemModel* self);
extern __declspec(dllexport) 
void QFileSystemModel_SetOption(QFileSystemModel* self, Option option);
extern __declspec(dllexport) 
bool QFileSystemModel_TestOption(const QFileSystemModel* self, Option option);
extern __declspec(dllexport) 
void QFileSystemModel_SetOptions(QFileSystemModel* self, Options options);
extern __declspec(dllexport) 
Options QFileSystemModel_Options(const QFileSystemModel* self);
extern __declspec(dllexport) 
struct miqt_string QFileSystemModel_FilePath(const QFileSystemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
bool QFileSystemModel_IsDir(const QFileSystemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
long long QFileSystemModel_Size(const QFileSystemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_string QFileSystemModel_Type(const QFileSystemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
QDateTime* QFileSystemModel_LastModified(const QFileSystemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
QDateTime* QFileSystemModel_LastModified2(const QFileSystemModel* self, QModelIndex* index, QTimeZone* tz);
extern __declspec(dllexport) 
QModelIndex* QFileSystemModel_Mkdir(QFileSystemModel* self, QModelIndex* parent, struct miqt_string name);
extern __declspec(dllexport) 
bool QFileSystemModel_Rmdir(QFileSystemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
struct miqt_string QFileSystemModel_FileName(const QFileSystemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
QIcon* QFileSystemModel_FileIcon(const QFileSystemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
int QFileSystemModel_Permissions(const QFileSystemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
QFileInfo* QFileSystemModel_FileInfo(const QFileSystemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
bool QFileSystemModel_Remove(QFileSystemModel* self, QModelIndex* index);
extern __declspec(dllexport) 
void QFileSystemModel_TimerEvent(QFileSystemModel* self, QTimerEvent* event);
extern __declspec(dllexport) 
bool QFileSystemModel_Event(QFileSystemModel* self, QEvent* event);
extern __declspec(dllexport) 
struct miqt_string QFileSystemModel_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QFileSystemModel_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
QModelIndex* QFileSystemModel_Index2(const QFileSystemModel* self, struct miqt_string path, int column);
extern __declspec(dllexport) 
QVariant* QFileSystemModel_MyComputer1(const QFileSystemModel* self, int role);
extern __declspec(dllexport) 
void QFileSystemModel_SetOption2(QFileSystemModel* self, Option option, bool on);
extern __declspec(dllexport) 
void QFileSystemModel_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QFileSystemModel_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QFileSystemModel_override_virtual_Metacast(void* self, intptr_t slot);
void* QFileSystemModel_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QFileSystemModel_Delete(QFileSystemModel* self, bool isSubclass);

}
