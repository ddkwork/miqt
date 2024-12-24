#pragma once
#ifndef MIQT_QT_PRINTSUPPORT_GEN_QPRINTPREVIEWWIDGET_H
#define MIQT_QT_PRINTSUPPORT_GEN_QPRINTPREVIEWWIDGET_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPrintPreviewWidget QPrintPreviewWidget;
typedef struct QPrinter QPrinter;
typedef struct QWidget QWidget;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QPrintPreviewWidget* QPrintPreviewWidget_new(QWidget* parent);
extern __declspec(dllexport) 
QPrintPreviewWidget* QPrintPreviewWidget_new2(QPrinter* printer);
extern __declspec(dllexport) 
QPrintPreviewWidget* QPrintPreviewWidget_new3();
extern __declspec(dllexport) 
QPrintPreviewWidget* QPrintPreviewWidget_new4(QPrinter* printer, QWidget* parent);
extern __declspec(dllexport) 
QPrintPreviewWidget* QPrintPreviewWidget_new5(QPrinter* printer, QWidget* parent, int flags);
extern __declspec(dllexport) 
QPrintPreviewWidget* QPrintPreviewWidget_new6(QWidget* parent, int flags);
extern __declspec(dllexport) 
void QPrintPreviewWidget_virtbase(QPrintPreviewWidget* src
, QWidget** outptr_QWidget
);
extern __declspec(dllexport) 
QMetaObject* QPrintPreviewWidget_MetaObject(const QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void* QPrintPreviewWidget_Metacast(QPrintPreviewWidget* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QPrintPreviewWidget_Tr(const char* s);
extern __declspec(dllexport) 
double QPrintPreviewWidget_ZoomFactor(const QPrintPreviewWidget* self);
extern __declspec(dllexport) 
int QPrintPreviewWidget_Orientation(const QPrintPreviewWidget* self);
extern __declspec(dllexport) 
ViewMode QPrintPreviewWidget_ViewMode(const QPrintPreviewWidget* self);
extern __declspec(dllexport) 
ZoomMode QPrintPreviewWidget_ZoomMode(const QPrintPreviewWidget* self);
extern __declspec(dllexport) 
int QPrintPreviewWidget_CurrentPage(const QPrintPreviewWidget* self);
extern __declspec(dllexport) 
int QPrintPreviewWidget_PageCount(const QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_SetVisible(QPrintPreviewWidget* self, bool visible);
extern __declspec(dllexport) 
void QPrintPreviewWidget_Print(QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_ZoomIn(QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_ZoomOut(QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_SetZoomFactor(QPrintPreviewWidget* self, double zoomFactor);
extern __declspec(dllexport) 
void QPrintPreviewWidget_SetOrientation(QPrintPreviewWidget* self, int orientation);
extern __declspec(dllexport) 
void QPrintPreviewWidget_SetViewMode(QPrintPreviewWidget* self, ViewMode viewMode);
extern __declspec(dllexport) 
void QPrintPreviewWidget_SetZoomMode(QPrintPreviewWidget* self, ZoomMode zoomMode);
extern __declspec(dllexport) 
void QPrintPreviewWidget_SetCurrentPage(QPrintPreviewWidget* self, int pageNumber);
extern __declspec(dllexport) 
void QPrintPreviewWidget_FitToWidth(QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_FitInView(QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_SetLandscapeOrientation(QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_SetPortraitOrientation(QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_SetSinglePageViewMode(QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_SetFacingPagesViewMode(QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_SetAllPagesViewMode(QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_UpdatePreview(QPrintPreviewWidget* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_PaintRequested(QPrintPreviewWidget* self, QPrinter* printer);
void QPrintPreviewWidget_connect_PaintRequested(QPrintPreviewWidget* self, intptr_t slot);
extern __declspec(dllexport) 
void QPrintPreviewWidget_PreviewChanged(QPrintPreviewWidget* self);
void QPrintPreviewWidget_connect_PreviewChanged(QPrintPreviewWidget* self, intptr_t slot);
extern __declspec(dllexport) 
struct miqt_string QPrintPreviewWidget_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QPrintPreviewWidget_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QPrintPreviewWidget_ZoomIn1(QPrintPreviewWidget* self, double zoom);
extern __declspec(dllexport) 
void QPrintPreviewWidget_ZoomOut1(QPrintPreviewWidget* self, double zoom);
extern __declspec(dllexport) 
void QPrintPreviewWidget_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QPrintPreviewWidget_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QPrintPreviewWidget_override_virtual_Metacast(void* self, intptr_t slot);
void* QPrintPreviewWidget_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QPrintPreviewWidget_Delete(QPrintPreviewWidget* self, bool isSubclass);

}
