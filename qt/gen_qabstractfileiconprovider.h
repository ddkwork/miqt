#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTFILEICONPROVIDER_H
#define MIQT_QT_GEN_QABSTRACTFILEICONPROVIDER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractFileIconProvider QAbstractFileIconProvider;
typedef struct QFileInfo QFileInfo;
typedef struct QIcon QIcon;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractFileIconProvider* QAbstractFileIconProvider_new();
extern __declspec(dllexport) 
QIcon* QAbstractFileIconProvider_Icon(const QAbstractFileIconProvider* self, IconType param1);
extern __declspec(dllexport) 
QIcon* QAbstractFileIconProvider_IconWithQFileInfo(const QAbstractFileIconProvider* self, QFileInfo* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractFileIconProvider_Type(const QAbstractFileIconProvider* self, QFileInfo* param1);
extern __declspec(dllexport) 
void QAbstractFileIconProvider_SetOptions(QAbstractFileIconProvider* self, Options options);
extern __declspec(dllexport) 
Options QAbstractFileIconProvider_Options(const QAbstractFileIconProvider* self);
extern __declspec(dllexport) 
void QAbstractFileIconProvider_Delete(QAbstractFileIconProvider* self, bool isSubclass);

}
