#pragma once
#ifndef MIQT_QT6_GEN_QHASHFUNCTIONS_H
#define MIQT_QT6_GEN_QHASHFUNCTIONS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QHashSeed;
#else
typedef struct QHashSeed QHashSeed;
#endif

void QHashSeed_new(QHashSeed** outptr_QHashSeed);
void QHashSeed_new2(size_t d, QHashSeed** outptr_QHashSeed);
QHashSeed* QHashSeed_GlobalSeed();
void QHashSeed_SetDeterministicGlobalSeed();
void QHashSeed_ResetRandomGlobalSeed();
void QHashSeed_Delete(QHashSeed* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
