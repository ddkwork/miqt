#pragma once
#ifndef MIQT_QT_GEN_QFILEICONPROVIDER_H
#define MIQT_QT_GEN_QFILEICONPROVIDER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractFileIconProvider QAbstractFileIconProvider;
typedef struct QFileIconProvider QFileIconProvider;
typedef struct QFileInfo QFileInfo;
typedef struct QIcon QIcon;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QFileIconProvider* QFileIconProvider_new();
extern __declspec(dllexport) 
void QFileIconProvider_virtbase(QFileIconProvider* src
, QAbstractFileIconProvider** outptr_QAbstractFileIconProvider
);
extern __declspec(dllexport) 
QIcon* QFileIconProvider_Icon(const QFileIconProvider* self, IconType typeVal);
extern __declspec(dllexport) 
QIcon* QFileIconProvider_IconWithInfo(const QFileIconProvider* self, QFileInfo* info);
extern __declspec(dllexport) 
void QFileIconProvider_Delete(QFileIconProvider* self, bool isSubclass);

}
