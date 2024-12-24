#pragma once
#ifndef MIQT_QT_GEN_QSPLASHSCREEN_H
#define MIQT_QT_GEN_QSPLASHSCREEN_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QColor QColor;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QMouseEvent QMouseEvent;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPainter QPainter;
typedef struct QPixmap QPixmap;
typedef struct QScreen QScreen;
typedef struct QSplashScreen QSplashScreen;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSplashScreen* QSplashScreen_new();
extern __declspec(dllexport) 
QSplashScreen* QSplashScreen_new2(QScreen* screen);
extern __declspec(dllexport) 
QSplashScreen* QSplashScreen_new3(QPixmap* pixmap);
extern __declspec(dllexport) 
QSplashScreen* QSplashScreen_new4(QPixmap* pixmap, int f);
extern __declspec(dllexport) 
QSplashScreen* QSplashScreen_new5(QScreen* screen, QPixmap* pixmap);
extern __declspec(dllexport) 
QSplashScreen* QSplashScreen_new6(QScreen* screen, QPixmap* pixmap, int f);
extern __declspec(dllexport) 
void QSplashScreen_virtbase(QSplashScreen* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QSplashScreen_MetaObject(const QSplashScreen* self);
extern __declspec(dllexport) 
void* QSplashScreen_Metacast(QSplashScreen* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QSplashScreen_Tr(const char* s);
extern __declspec(dllexport) 
void QSplashScreen_SetPixmap(QSplashScreen* self, QPixmap* pixmap);
extern __declspec(dllexport) 
QPixmap* QSplashScreen_Pixmap(const QSplashScreen* self);
extern __declspec(dllexport) 
void QSplashScreen_Finish(QSplashScreen* self, QWidget* w);
extern __declspec(dllexport) 
void QSplashScreen_Repaint(QSplashScreen* self);
extern __declspec(dllexport) 
struct miqt_string QSplashScreen_Message(const QSplashScreen* self);
extern __declspec(dllexport) 
void QSplashScreen_ShowMessage(QSplashScreen* self, struct miqt_string message);
extern __declspec(dllexport) 
void QSplashScreen_ClearMessage(QSplashScreen* self);
extern __declspec(dllexport) 
void QSplashScreen_MessageChanged(QSplashScreen* self, struct miqt_string message);
void QSplashScreen_connect_MessageChanged(QSplashScreen* self, intptr_t slot);
extern __declspec(dllexport) 
bool QSplashScreen_Event(QSplashScreen* self, QEvent* e);
extern __declspec(dllexport) 
void QSplashScreen_DrawContents(QSplashScreen* self, QPainter* painter);
extern __declspec(dllexport) 
void QSplashScreen_MousePressEvent(QSplashScreen* self, QMouseEvent* param1);
extern __declspec(dllexport) 
struct miqt_string QSplashScreen_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QSplashScreen_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QSplashScreen_ShowMessage2(QSplashScreen* self, struct miqt_string message, int alignment);
extern __declspec(dllexport) 
void QSplashScreen_ShowMessage3(QSplashScreen* self, struct miqt_string message, int alignment, QColor* color);
extern __declspec(dllexport) 
void QSplashScreen_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QSplashScreen_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QSplashScreen_override_virtual_Metacast(void* self, intptr_t slot);
void* QSplashScreen_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QSplashScreen_Delete(QSplashScreen* self, bool isSubclass);

}
