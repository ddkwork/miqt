#ifndef GEN_QDIR_H
#define GEN_QDIR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QChar;
class QDir;
class QFileInfo;
#else
typedef struct QChar QChar;
typedef struct QDir QDir;
typedef struct QFileInfo QFileInfo;
#endif

QDir* QDir_new(QDir* param1);
QDir* QDir_new2();
QDir* QDir_new3(const char* path, size_t path_Strlen, const char* nameFilter, size_t nameFilter_Strlen);
QDir* QDir_new4(const char* path, size_t path_Strlen);
QDir* QDir_new5(const char* path, size_t path_Strlen, const char* nameFilter, size_t nameFilter_Strlen, int sort);
QDir* QDir_new6(const char* path, size_t path_Strlen, const char* nameFilter, size_t nameFilter_Strlen, int sort, int filter);
void QDir_OperatorAssign(QDir* self, QDir* param1);
void QDir_OperatorAssignWithPath(QDir* self, const char* path, size_t path_Strlen);
void QDir_Swap(QDir* self, QDir* other);
void QDir_SetPath(QDir* self, const char* path, size_t path_Strlen);
void QDir_Path(const QDir* self, char** _out, int* _out_Strlen);
void QDir_AbsolutePath(const QDir* self, char** _out, int* _out_Strlen);
void QDir_CanonicalPath(const QDir* self, char** _out, int* _out_Strlen);
void QDir_AddResourceSearchPath(const char* path, size_t path_Strlen);
void QDir_SetSearchPaths(const char* prefix, size_t prefix_Strlen, char** searchPaths, uint64_t* searchPaths_Lengths, size_t searchPaths_len);
void QDir_AddSearchPath(const char* prefix, size_t prefix_Strlen, const char* path, size_t path_Strlen);
void QDir_SearchPaths(const char* prefix, size_t prefix_Strlen, char*** _out, int** _out_Lengths, size_t* _out_len);
void QDir_DirName(const QDir* self, char** _out, int* _out_Strlen);
void QDir_FilePath(const QDir* self, const char* fileName, size_t fileName_Strlen, char** _out, int* _out_Strlen);
void QDir_AbsoluteFilePath(const QDir* self, const char* fileName, size_t fileName_Strlen, char** _out, int* _out_Strlen);
void QDir_RelativeFilePath(const QDir* self, const char* fileName, size_t fileName_Strlen, char** _out, int* _out_Strlen);
void QDir_ToNativeSeparators(const char* pathName, size_t pathName_Strlen, char** _out, int* _out_Strlen);
void QDir_FromNativeSeparators(const char* pathName, size_t pathName_Strlen, char** _out, int* _out_Strlen);
bool QDir_Cd(QDir* self, const char* dirName, size_t dirName_Strlen);
bool QDir_CdUp(QDir* self);
void QDir_NameFilters(const QDir* self, char*** _out, int** _out_Lengths, size_t* _out_len);
void QDir_SetNameFilters(QDir* self, char** nameFilters, uint64_t* nameFilters_Lengths, size_t nameFilters_len);
int QDir_Filter(const QDir* self);
void QDir_SetFilter(QDir* self, int filter);
int QDir_Sorting(const QDir* self);
void QDir_SetSorting(QDir* self, int sort);
unsigned int QDir_Count(const QDir* self);
bool QDir_IsEmpty(const QDir* self);
void QDir_OperatorSubscript(const QDir* self, int param1, char** _out, int* _out_Strlen);
void QDir_NameFiltersFromString(const char* nameFilter, size_t nameFilter_Strlen, char*** _out, int** _out_Lengths, size_t* _out_len);
void QDir_EntryList(const QDir* self, char*** _out, int** _out_Lengths, size_t* _out_len);
void QDir_EntryListWithNameFilters(const QDir* self, char** nameFilters, uint64_t* nameFilters_Lengths, size_t nameFilters_len, char*** _out, int** _out_Lengths, size_t* _out_len);
void QDir_EntryInfoList(const QDir* self, QFileInfo*** _out, size_t* _out_len);
void QDir_EntryInfoListWithNameFilters(const QDir* self, char** nameFilters, uint64_t* nameFilters_Lengths, size_t nameFilters_len, QFileInfo*** _out, size_t* _out_len);
bool QDir_Mkdir(const QDir* self, const char* dirName, size_t dirName_Strlen);
bool QDir_Rmdir(const QDir* self, const char* dirName, size_t dirName_Strlen);
bool QDir_Mkpath(const QDir* self, const char* dirPath, size_t dirPath_Strlen);
bool QDir_Rmpath(const QDir* self, const char* dirPath, size_t dirPath_Strlen);
bool QDir_RemoveRecursively(QDir* self);
bool QDir_IsReadable(const QDir* self);
bool QDir_Exists(const QDir* self);
bool QDir_IsRoot(const QDir* self);
bool QDir_IsRelativePath(const char* path, size_t path_Strlen);
bool QDir_IsAbsolutePath(const char* path, size_t path_Strlen);
bool QDir_IsRelative(const QDir* self);
bool QDir_IsAbsolute(const QDir* self);
bool QDir_MakeAbsolute(QDir* self);
bool QDir_OperatorEqual(const QDir* self, QDir* dir);
bool QDir_OperatorNotEqual(const QDir* self, QDir* dir);
bool QDir_Remove(QDir* self, const char* fileName, size_t fileName_Strlen);
bool QDir_Rename(QDir* self, const char* oldName, size_t oldName_Strlen, const char* newName, size_t newName_Strlen);
bool QDir_ExistsWithName(const QDir* self, const char* name, size_t name_Strlen);
void QDir_Drives(QFileInfo*** _out, size_t* _out_len);
QChar* QDir_ListSeparator();
QChar* QDir_Separator();
bool QDir_SetCurrent(const char* path, size_t path_Strlen);
QDir* QDir_Current();
void QDir_CurrentPath(char** _out, int* _out_Strlen);
QDir* QDir_Home();
void QDir_HomePath(char** _out, int* _out_Strlen);
QDir* QDir_Root();
void QDir_RootPath(char** _out, int* _out_Strlen);
QDir* QDir_Temp();
void QDir_TempPath(char** _out, int* _out_Strlen);
bool QDir_Match(char** filters, uint64_t* filters_Lengths, size_t filters_len, const char* fileName, size_t fileName_Strlen);
bool QDir_Match2(const char* filter, size_t filter_Strlen, const char* fileName, size_t fileName_Strlen);
void QDir_CleanPath(const char* path, size_t path_Strlen, char** _out, int* _out_Strlen);
void QDir_Refresh(const QDir* self);
bool QDir_IsEmpty1(const QDir* self, int filters);
void QDir_EntryList1(const QDir* self, int filters, char*** _out, int** _out_Lengths, size_t* _out_len);
void QDir_EntryList2(const QDir* self, int filters, int sort, char*** _out, int** _out_Lengths, size_t* _out_len);
void QDir_EntryList22(const QDir* self, char** nameFilters, uint64_t* nameFilters_Lengths, size_t nameFilters_len, int filters, char*** _out, int** _out_Lengths, size_t* _out_len);
void QDir_EntryList3(const QDir* self, char** nameFilters, uint64_t* nameFilters_Lengths, size_t nameFilters_len, int filters, int sort, char*** _out, int** _out_Lengths, size_t* _out_len);
void QDir_EntryInfoList1(const QDir* self, int filters, QFileInfo*** _out, size_t* _out_len);
void QDir_EntryInfoList2(const QDir* self, int filters, int sort, QFileInfo*** _out, size_t* _out_len);
void QDir_EntryInfoList22(const QDir* self, char** nameFilters, uint64_t* nameFilters_Lengths, size_t nameFilters_len, int filters, QFileInfo*** _out, size_t* _out_len);
void QDir_EntryInfoList3(const QDir* self, char** nameFilters, uint64_t* nameFilters_Lengths, size_t nameFilters_len, int filters, int sort, QFileInfo*** _out, size_t* _out_len);
void QDir_Delete(QDir* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
