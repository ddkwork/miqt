#ifndef GEN_QFILESYSTEMMODEL_H
#define GEN_QFILESYSTEMMODEL_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QDateTime;
class QDir;
class QFileIconProvider;
class QFileInfo;
class QFileSystemModel;
class QIcon;
class QMetaObject;
class QMimeData;
class QModelIndex;
class QObject;
class QVariant;
#else
typedef struct QDateTime QDateTime;
typedef struct QDir QDir;
typedef struct QFileIconProvider QFileIconProvider;
typedef struct QFileInfo QFileInfo;
typedef struct QFileSystemModel QFileSystemModel;
typedef struct QIcon QIcon;
typedef struct QMetaObject QMetaObject;
typedef struct QMimeData QMimeData;
typedef struct QModelIndex QModelIndex;
typedef struct QObject QObject;
typedef struct QVariant QVariant;
#endif

QFileSystemModel* QFileSystemModel_new();
QFileSystemModel* QFileSystemModel_new2(QObject* parent);
QMetaObject* QFileSystemModel_MetaObject(const QFileSystemModel* self);
void QFileSystemModel_Tr(const char* s, char** _out, int* _out_Strlen);
void QFileSystemModel_TrUtf8(const char* s, char** _out, int* _out_Strlen);
void QFileSystemModel_RootPathChanged(QFileSystemModel* self, const char* newPath, size_t newPath_Strlen);
void QFileSystemModel_connect_RootPathChanged(QFileSystemModel* self, void* slot);
void QFileSystemModel_FileRenamed(QFileSystemModel* self, const char* path, size_t path_Strlen, const char* oldName, size_t oldName_Strlen, const char* newName, size_t newName_Strlen);
void QFileSystemModel_connect_FileRenamed(QFileSystemModel* self, void* slot);
void QFileSystemModel_DirectoryLoaded(QFileSystemModel* self, const char* path, size_t path_Strlen);
void QFileSystemModel_connect_DirectoryLoaded(QFileSystemModel* self, void* slot);
QModelIndex* QFileSystemModel_Index(const QFileSystemModel* self, int row, int column);
QModelIndex* QFileSystemModel_IndexWithPath(const QFileSystemModel* self, const char* path, size_t path_Strlen);
QModelIndex* QFileSystemModel_Parent(const QFileSystemModel* self, QModelIndex* child);
QModelIndex* QFileSystemModel_Sibling(const QFileSystemModel* self, int row, int column, QModelIndex* idx);
bool QFileSystemModel_HasChildren(const QFileSystemModel* self);
bool QFileSystemModel_CanFetchMore(const QFileSystemModel* self, QModelIndex* parent);
void QFileSystemModel_FetchMore(QFileSystemModel* self, QModelIndex* parent);
int QFileSystemModel_RowCount(const QFileSystemModel* self);
int QFileSystemModel_ColumnCount(const QFileSystemModel* self);
QVariant* QFileSystemModel_MyComputer(const QFileSystemModel* self);
QVariant* QFileSystemModel_Data(const QFileSystemModel* self, QModelIndex* index);
bool QFileSystemModel_SetData(QFileSystemModel* self, QModelIndex* index, QVariant* value);
QVariant* QFileSystemModel_HeaderData(const QFileSystemModel* self, int section, uintptr_t orientation);
int QFileSystemModel_Flags(const QFileSystemModel* self, QModelIndex* index);
void QFileSystemModel_Sort(QFileSystemModel* self, int column);
void QFileSystemModel_MimeTypes(const QFileSystemModel* self, char*** _out, int** _out_Lengths, size_t* _out_len);
QMimeData* QFileSystemModel_MimeData(const QFileSystemModel* self, QModelIndex** indexes, size_t indexes_len);
bool QFileSystemModel_DropMimeData(QFileSystemModel* self, QMimeData* data, uintptr_t action, int row, int column, QModelIndex* parent);
int QFileSystemModel_SupportedDropActions(const QFileSystemModel* self);
QModelIndex* QFileSystemModel_SetRootPath(QFileSystemModel* self, const char* path, size_t path_Strlen);
void QFileSystemModel_RootPath(const QFileSystemModel* self, char** _out, int* _out_Strlen);
QDir* QFileSystemModel_RootDirectory(const QFileSystemModel* self);
void QFileSystemModel_SetIconProvider(QFileSystemModel* self, QFileIconProvider* provider);
QFileIconProvider* QFileSystemModel_IconProvider(const QFileSystemModel* self);
void QFileSystemModel_SetFilter(QFileSystemModel* self, int filters);
int QFileSystemModel_Filter(const QFileSystemModel* self);
void QFileSystemModel_SetResolveSymlinks(QFileSystemModel* self, bool enable);
bool QFileSystemModel_ResolveSymlinks(const QFileSystemModel* self);
void QFileSystemModel_SetReadOnly(QFileSystemModel* self, bool enable);
bool QFileSystemModel_IsReadOnly(const QFileSystemModel* self);
void QFileSystemModel_SetNameFilterDisables(QFileSystemModel* self, bool enable);
bool QFileSystemModel_NameFilterDisables(const QFileSystemModel* self);
void QFileSystemModel_SetNameFilters(QFileSystemModel* self, char** filters, uint64_t* filters_Lengths, size_t filters_len);
void QFileSystemModel_NameFilters(const QFileSystemModel* self, char*** _out, int** _out_Lengths, size_t* _out_len);
void QFileSystemModel_SetOption(QFileSystemModel* self, uintptr_t option);
bool QFileSystemModel_TestOption(const QFileSystemModel* self, uintptr_t option);
void QFileSystemModel_SetOptions(QFileSystemModel* self, int options);
int QFileSystemModel_Options(const QFileSystemModel* self);
void QFileSystemModel_FilePath(const QFileSystemModel* self, QModelIndex* index, char** _out, int* _out_Strlen);
bool QFileSystemModel_IsDir(const QFileSystemModel* self, QModelIndex* index);
long long QFileSystemModel_Size(const QFileSystemModel* self, QModelIndex* index);
void QFileSystemModel_Type(const QFileSystemModel* self, QModelIndex* index, char** _out, int* _out_Strlen);
QDateTime* QFileSystemModel_LastModified(const QFileSystemModel* self, QModelIndex* index);
QModelIndex* QFileSystemModel_Mkdir(QFileSystemModel* self, QModelIndex* parent, const char* name, size_t name_Strlen);
bool QFileSystemModel_Rmdir(QFileSystemModel* self, QModelIndex* index);
void QFileSystemModel_FileName(const QFileSystemModel* self, QModelIndex* index, char** _out, int* _out_Strlen);
QIcon* QFileSystemModel_FileIcon(const QFileSystemModel* self, QModelIndex* index);
int QFileSystemModel_Permissions(const QFileSystemModel* self, QModelIndex* index);
QFileInfo* QFileSystemModel_FileInfo(const QFileSystemModel* self, QModelIndex* index);
bool QFileSystemModel_Remove(QFileSystemModel* self, QModelIndex* index);
void QFileSystemModel_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QFileSystemModel_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QFileSystemModel_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QFileSystemModel_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
QModelIndex* QFileSystemModel_Index3(const QFileSystemModel* self, int row, int column, QModelIndex* parent);
QModelIndex* QFileSystemModel_Index2(const QFileSystemModel* self, const char* path, size_t path_Strlen, int column);
bool QFileSystemModel_HasChildren1(const QFileSystemModel* self, QModelIndex* parent);
int QFileSystemModel_RowCount1(const QFileSystemModel* self, QModelIndex* parent);
int QFileSystemModel_ColumnCount1(const QFileSystemModel* self, QModelIndex* parent);
QVariant* QFileSystemModel_MyComputer1(const QFileSystemModel* self, int role);
QVariant* QFileSystemModel_Data2(const QFileSystemModel* self, QModelIndex* index, int role);
bool QFileSystemModel_SetData3(QFileSystemModel* self, QModelIndex* index, QVariant* value, int role);
QVariant* QFileSystemModel_HeaderData3(const QFileSystemModel* self, int section, uintptr_t orientation, int role);
void QFileSystemModel_Sort2(QFileSystemModel* self, int column, uintptr_t order);
void QFileSystemModel_SetOption2(QFileSystemModel* self, uintptr_t option, bool on);
void QFileSystemModel_Delete(QFileSystemModel* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
