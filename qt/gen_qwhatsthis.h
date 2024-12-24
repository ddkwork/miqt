#pragma once
#ifndef MIQT_QT_GEN_QWHATSTHIS_H
#define MIQT_QT_GEN_QWHATSTHIS_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAction QAction;
typedef struct QObject QObject;
typedef struct QPoint QPoint;
typedef struct QWhatsThis QWhatsThis;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
void QWhatsThis_EnterWhatsThisMode();
extern __declspec(dllexport) 
bool QWhatsThis_InWhatsThisMode();
extern __declspec(dllexport) 
void QWhatsThis_LeaveWhatsThisMode();
extern __declspec(dllexport) 
void QWhatsThis_ShowText(QPoint* pos, struct miqt_string text);
extern __declspec(dllexport) 
void QWhatsThis_HideText();
extern __declspec(dllexport) 
QAction* QWhatsThis_CreateAction();
extern __declspec(dllexport) 
void QWhatsThis_ShowText3(QPoint* pos, struct miqt_string text, QWidget* w);
extern __declspec(dllexport) 
QAction* QWhatsThis_CreateAction1(QObject* parent);
extern __declspec(dllexport) 
void QWhatsThis_Delete(QWhatsThis* self, bool isSubclass);

}
