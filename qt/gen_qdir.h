#pragma once
#ifndef MIQT_QT_GEN_QDIR_H
#define MIQT_QT_GEN_QDIR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_Disambiguated_t)
typedef Qt::Disambiguated_t Disambiguated_t;
typedef struct QChar QChar;
typedef struct QDir QDir;
typedef struct QFileInfo QFileInfo;
typedef struct Disambiguated_t Disambiguated_t;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QDir* QDir_new(QDir* param1);
extern __declspec(dllexport) QDir* QDir_new2();
extern __declspec(dllexport) QDir* QDir_new3(struct miqt_string path, struct miqt_string nameFilter);
extern __declspec(dllexport) QDir* QDir_new4(struct miqt_string path);
extern __declspec(dllexport) QDir* QDir_new5(struct miqt_string path, struct miqt_string nameFilter, SortFlags sort);
extern __declspec(dllexport) QDir* QDir_new6(struct miqt_string path, struct miqt_string nameFilter, SortFlags sort, Filters filter);
extern __declspec(dllexport) void QDir_OperatorAssign(QDir* self, QDir* param1);
extern __declspec(dllexport) void QDir_Swap(QDir* self, QDir* other);
extern __declspec(dllexport) void QDir_SetPath(QDir* self, struct miqt_string path);
extern __declspec(dllexport) struct miqt_string QDir_Path(const QDir* self);
extern __declspec(dllexport) struct miqt_string QDir_AbsolutePath(const QDir* self);
extern __declspec(dllexport) struct miqt_string QDir_CanonicalPath(const QDir* self);
extern __declspec(dllexport) void QDir_SetSearchPaths(struct miqt_string prefix, struct miqt_array /* of struct miqt_string */  searchPaths);
extern __declspec(dllexport) void QDir_AddSearchPath(struct miqt_string prefix, struct miqt_string path);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QDir_SearchPaths(struct miqt_string prefix);
extern __declspec(dllexport) struct miqt_string QDir_DirName(const QDir* self);
extern __declspec(dllexport) struct miqt_string QDir_FilePath(const QDir* self, struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QDir_AbsoluteFilePath(const QDir* self, struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QDir_RelativeFilePath(const QDir* self, struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QDir_ToNativeSeparators(struct miqt_string pathName);
extern __declspec(dllexport) struct miqt_string QDir_FromNativeSeparators(struct miqt_string pathName);
extern __declspec(dllexport) bool QDir_Cd(QDir* self, struct miqt_string dirName);
extern __declspec(dllexport) bool QDir_CdUp(QDir* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QDir_NameFilters(const QDir* self);
extern __declspec(dllexport) void QDir_SetNameFilters(QDir* self, struct miqt_array /* of struct miqt_string */  nameFilters);
extern __declspec(dllexport) Filters QDir_Filter(const QDir* self);
extern __declspec(dllexport) void QDir_SetFilter(QDir* self, Filters filter);
extern __declspec(dllexport) SortFlags QDir_Sorting(const QDir* self);
extern __declspec(dllexport) void QDir_SetSorting(QDir* self, SortFlags sort);
extern __declspec(dllexport) ptrdiff_t QDir_Count(const QDir* self);
extern __declspec(dllexport) bool QDir_IsEmpty(const QDir* self);
extern __declspec(dllexport) struct miqt_string QDir_OperatorSubscript(const QDir* self, ptrdiff_t param1);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QDir_NameFiltersFromString(struct miqt_string nameFilter);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QDir_EntryList(const QDir* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QDir_EntryListWithNameFilters(const QDir* self, struct miqt_array /* of struct miqt_string */  nameFilters);
extern __declspec(dllexport) struct miqt_array /* of QFileInfo* */  QDir_EntryInfoList(const QDir* self);
extern __declspec(dllexport) struct miqt_array /* of QFileInfo* */  QDir_EntryInfoListWithNameFilters(const QDir* self, struct miqt_array /* of struct miqt_string */  nameFilters);
extern __declspec(dllexport) bool QDir_Mkdir(const QDir* self, struct miqt_string dirName);
extern __declspec(dllexport) bool QDir_Mkdir2(const QDir* self, struct miqt_string dirName, int permissions);
extern __declspec(dllexport) bool QDir_Rmdir(const QDir* self, struct miqt_string dirName);
extern __declspec(dllexport) bool QDir_Mkpath(const QDir* self, struct miqt_string dirPath);
extern __declspec(dllexport) bool QDir_Rmpath(const QDir* self, struct miqt_string dirPath);
extern __declspec(dllexport) bool QDir_RemoveRecursively(QDir* self);
extern __declspec(dllexport) bool QDir_IsReadable(const QDir* self);
extern __declspec(dllexport) bool QDir_Exists(const QDir* self);
extern __declspec(dllexport) bool QDir_IsRoot(const QDir* self);
extern __declspec(dllexport) bool QDir_IsRelativePath(struct miqt_string path);
extern __declspec(dllexport) bool QDir_IsAbsolutePath(struct miqt_string path);
extern __declspec(dllexport) bool QDir_IsRelative(const QDir* self);
extern __declspec(dllexport) bool QDir_IsAbsolute(const QDir* self);
extern __declspec(dllexport) bool QDir_MakeAbsolute(QDir* self);
extern __declspec(dllexport) bool QDir_Remove(QDir* self, struct miqt_string fileName);
extern __declspec(dllexport) bool QDir_Rename(QDir* self, struct miqt_string oldName, struct miqt_string newName);
extern __declspec(dllexport) bool QDir_ExistsWithName(const QDir* self, struct miqt_string name);
extern __declspec(dllexport) struct miqt_array /* of QFileInfo* */  QDir_Drives();
extern __declspec(dllexport) QChar* QDir_ListSeparator();
extern __declspec(dllexport) QChar* QDir_Separator();
extern __declspec(dllexport) bool QDir_SetCurrent(struct miqt_string path);
extern __declspec(dllexport) QDir* QDir_Current();
extern __declspec(dllexport) struct miqt_string QDir_CurrentPath();
extern __declspec(dllexport) QDir* QDir_Home();
extern __declspec(dllexport) struct miqt_string QDir_HomePath();
extern __declspec(dllexport) QDir* QDir_Root();
extern __declspec(dllexport) struct miqt_string QDir_RootPath();
extern __declspec(dllexport) QDir* QDir_Temp();
extern __declspec(dllexport) struct miqt_string QDir_TempPath();
extern __declspec(dllexport) bool QDir_Match(struct miqt_array /* of struct miqt_string */  filters, struct miqt_string fileName);
extern __declspec(dllexport) bool QDir_Match2(struct miqt_string filter, struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_string QDir_CleanPath(struct miqt_string path);
extern __declspec(dllexport) void QDir_Refresh(const QDir* self);
extern __declspec(dllexport) ptrdiff_t QDir_Count1(const QDir* self, Disambiguated_t* param1);
extern __declspec(dllexport) bool QDir_IsEmpty1(const QDir* self, Filters filters);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QDir_EntryList1(const QDir* self, Filters filters);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QDir_EntryList2(const QDir* self, Filters filters, SortFlags sort);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QDir_EntryList22(const QDir* self, struct miqt_array /* of struct miqt_string */  nameFilters, Filters filters);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QDir_EntryList3(const QDir* self, struct miqt_array /* of struct miqt_string */  nameFilters, Filters filters, SortFlags sort);
extern __declspec(dllexport) struct miqt_array /* of QFileInfo* */  QDir_EntryInfoList1(const QDir* self, Filters filters);
extern __declspec(dllexport) struct miqt_array /* of QFileInfo* */  QDir_EntryInfoList2(const QDir* self, Filters filters, SortFlags sort);
extern __declspec(dllexport) struct miqt_array /* of QFileInfo* */  QDir_EntryInfoList22(const QDir* self, struct miqt_array /* of struct miqt_string */  nameFilters, Filters filters);
extern __declspec(dllexport) struct miqt_array /* of QFileInfo* */  QDir_EntryInfoList3(const QDir* self, struct miqt_array /* of struct miqt_string */  nameFilters, Filters filters, SortFlags sort);
extern __declspec(dllexport) void QDir_Delete(QDir* self, bool isSubclass);

} 
