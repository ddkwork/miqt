#pragma once
#ifndef MIQT_QT_GEN_QDESKTOPSERVICES_H
#define MIQT_QT_GEN_QDESKTOPSERVICES_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QDesktopServices QDesktopServices;
typedef struct QObject QObject;
typedef struct QUrl QUrl;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
bool QDesktopServices_OpenUrl(QUrl* url);
extern __declspec(dllexport) 
void QDesktopServices_SetUrlHandler(struct miqt_string scheme, QObject* receiver, const char* method);
extern __declspec(dllexport) 
void QDesktopServices_UnsetUrlHandler(struct miqt_string scheme);
extern __declspec(dllexport) 
void QDesktopServices_Delete(QDesktopServices* self, bool isSubclass);

}
