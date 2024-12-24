#pragma once
#ifndef MIQT_QT_SVG_GEN_QSVGWIDGET_H
#define MIQT_QT_SVG_GEN_QSVGWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEvent QPaintEvent;
typedef struct QSize QSize;
typedef struct QSvgRenderer QSvgRenderer;
typedef struct QSvgWidget QSvgWidget;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QSvgWidget* QSvgWidget_new(QWidget* parent);
extern __declspec(dllexport) 
QSvgWidget* QSvgWidget_new2();
extern __declspec(dllexport) 
QSvgWidget* QSvgWidget_new3(struct miqt_string file);
extern __declspec(dllexport) 
QSvgWidget* QSvgWidget_new4(struct miqt_string file, QWidget* parent);
extern __declspec(dllexport) 
void QSvgWidget_virtbase(QSvgWidget* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QSvgWidget_MetaObject(const QSvgWidget* self);
extern __declspec(dllexport) 
void* QSvgWidget_Metacast(QSvgWidget* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QSvgWidget_Tr(const char* s);
extern __declspec(dllexport) 
QSvgRenderer* QSvgWidget_Renderer(const QSvgWidget* self);
extern __declspec(dllexport) 
QSize* QSvgWidget_SizeHint(const QSvgWidget* self);
extern __declspec(dllexport) 
int QSvgWidget_Options(const QSvgWidget* self);
extern __declspec(dllexport) 
void QSvgWidget_SetOptions(QSvgWidget* self, int options);
extern __declspec(dllexport) 
void QSvgWidget_Load(QSvgWidget* self, struct miqt_string file);
extern __declspec(dllexport) 
void QSvgWidget_LoadWithContents(QSvgWidget* self, struct miqt_string contents);
extern __declspec(dllexport) 
void QSvgWidget_PaintEvent(QSvgWidget* self, QPaintEvent* event);
extern __declspec(dllexport) 
struct miqt_string QSvgWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QSvgWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QSvgWidget_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QSvgWidget_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QSvgWidget_override_virtual_Metacast(void* self, intptr_t slot);
void* QSvgWidget_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QSvgWidget_Delete(QSvgWidget* self, bool isSubclass);

}
