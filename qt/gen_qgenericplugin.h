#ifndef GEN_QGENERICPLUGIN_H
#define GEN_QGENERICPLUGIN_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "binding.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QGenericPlugin;
class QMetaObject;
class QObject;
#else
typedef struct QGenericPlugin QGenericPlugin;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
#endif

QMetaObject* QGenericPlugin_MetaObject(const QGenericPlugin* self);
void* QGenericPlugin_Metacast(QGenericPlugin* self, const char* param1);
struct miqt_string* QGenericPlugin_Tr(const char* s);
struct miqt_string* QGenericPlugin_TrUtf8(const char* s);
QObject* QGenericPlugin_Create(QGenericPlugin* self, struct miqt_string* name, struct miqt_string* spec);
struct miqt_string* QGenericPlugin_Tr2(const char* s, const char* c);
struct miqt_string* QGenericPlugin_Tr3(const char* s, const char* c, int n);
struct miqt_string* QGenericPlugin_TrUtf82(const char* s, const char* c);
struct miqt_string* QGenericPlugin_TrUtf83(const char* s, const char* c, int n);
void QGenericPlugin_Delete(QGenericPlugin* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
