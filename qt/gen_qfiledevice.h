#ifndef GEN_QFILEDEVICE_H
#define GEN_QFILEDEVICE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QDateTime;
class QFileDevice;
class QMetaObject;
#else
typedef struct QDateTime QDateTime;
typedef struct QFileDevice QFileDevice;
typedef struct QMetaObject QMetaObject;
#endif

QMetaObject* QFileDevice_MetaObject(QFileDevice* self);
void QFileDevice_Tr(const char* s, char** _out, int* _out_Strlen);
void QFileDevice_TrUtf8(const char* s, char** _out, int* _out_Strlen);
uintptr_t QFileDevice_Error(QFileDevice* self);
void QFileDevice_UnsetError(QFileDevice* self);
void QFileDevice_Close(QFileDevice* self);
bool QFileDevice_IsSequential(QFileDevice* self);
int QFileDevice_Handle(QFileDevice* self);
void QFileDevice_FileName(QFileDevice* self, char** _out, int* _out_Strlen);
long long QFileDevice_Pos(QFileDevice* self);
bool QFileDevice_Seek(QFileDevice* self, long long offset);
bool QFileDevice_AtEnd(QFileDevice* self);
bool QFileDevice_Flush(QFileDevice* self);
long long QFileDevice_Size(QFileDevice* self);
bool QFileDevice_Resize(QFileDevice* self, long long sz);
int QFileDevice_Permissions(QFileDevice* self);
bool QFileDevice_SetPermissions(QFileDevice* self, int permissionSpec);
unsigned char* QFileDevice_Map(QFileDevice* self, long long offset, long long size);
bool QFileDevice_Unmap(QFileDevice* self, unsigned char* address);
QDateTime* QFileDevice_FileTime(QFileDevice* self, uintptr_t time);
bool QFileDevice_SetFileTime(QFileDevice* self, QDateTime* newDate, uintptr_t fileTime);
void QFileDevice_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QFileDevice_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QFileDevice_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QFileDevice_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
unsigned char* QFileDevice_Map3(QFileDevice* self, long long offset, long long size, uintptr_t flags);
void QFileDevice_Delete(QFileDevice* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif