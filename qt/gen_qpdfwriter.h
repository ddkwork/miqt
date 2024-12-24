#pragma once
#ifndef MIQT_QT_GEN_QPDFWRITER_H
#define MIQT_QT_GEN_QPDFWRITER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QChildEvent QChildEvent;
typedef struct QEvent QEvent;
typedef struct QIODevice QIODevice;
typedef struct QMarginsF QMarginsF;
typedef struct QMetaMethod QMetaMethod;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPageLayout QPageLayout;
typedef struct QPageRanges QPageRanges;
typedef struct QPageSize QPageSize;
typedef struct QPagedPaintDevice QPagedPaintDevice;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPdfOutputIntent QPdfOutputIntent;
typedef struct QPdfWriter QPdfWriter;
typedef struct QTimerEvent QTimerEvent;
typedef struct QUuid QUuid;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPdfWriter* QPdfWriter_new(struct miqt_string filename);
extern __declspec(dllexport) QPdfWriter* QPdfWriter_new2(QIODevice* device);
extern __declspec(dllexport) void QPdfWriter_virtbase(QPdfWriter* src, QObject** outptr_QObject, QPagedPaintDevice** outptr_QPagedPaintDevice);
extern __declspec(dllexport) QMetaObject* QPdfWriter_MetaObject(const QPdfWriter* self);
extern __declspec(dllexport) void* QPdfWriter_Metacast(QPdfWriter* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QPdfWriter_Tr(const char* s);
extern __declspec(dllexport) void QPdfWriter_SetPdfVersion(QPdfWriter* self, PdfVersion version);
extern __declspec(dllexport) PdfVersion QPdfWriter_PdfVersion(const QPdfWriter* self);
extern __declspec(dllexport) struct miqt_string QPdfWriter_Title(const QPdfWriter* self);
extern __declspec(dllexport) void QPdfWriter_SetTitle(QPdfWriter* self, struct miqt_string title);
extern __declspec(dllexport) struct miqt_string QPdfWriter_Creator(const QPdfWriter* self);
extern __declspec(dllexport) void QPdfWriter_SetCreator(QPdfWriter* self, struct miqt_string creator);
extern __declspec(dllexport) QUuid* QPdfWriter_DocumentId(const QPdfWriter* self);
extern __declspec(dllexport) void QPdfWriter_SetDocumentId(QPdfWriter* self, QUuid* documentId);
extern __declspec(dllexport) struct miqt_string QPdfWriter_Author(const QPdfWriter* self);
extern __declspec(dllexport) void QPdfWriter_SetAuthor(QPdfWriter* self, struct miqt_string author);
extern __declspec(dllexport) bool QPdfWriter_NewPage(QPdfWriter* self);
extern __declspec(dllexport) void QPdfWriter_SetResolution(QPdfWriter* self, int resolution);
extern __declspec(dllexport) int QPdfWriter_Resolution(const QPdfWriter* self);
extern __declspec(dllexport) void QPdfWriter_SetDocumentXmpMetadata(QPdfWriter* self, struct miqt_string xmpMetadata);
extern __declspec(dllexport) struct miqt_string QPdfWriter_DocumentXmpMetadata(const QPdfWriter* self);
extern __declspec(dllexport) void QPdfWriter_AddFileAttachment(QPdfWriter* self, struct miqt_string fileName, struct miqt_string data);
extern __declspec(dllexport) ColorModel QPdfWriter_ColorModel(const QPdfWriter* self);
extern __declspec(dllexport) void QPdfWriter_SetColorModel(QPdfWriter* self, ColorModel model);
extern __declspec(dllexport) QPdfOutputIntent* QPdfWriter_OutputIntent(const QPdfWriter* self);
extern __declspec(dllexport) void QPdfWriter_SetOutputIntent(QPdfWriter* self, QPdfOutputIntent* intent);
extern __declspec(dllexport) QPaintEngine* QPdfWriter_PaintEngine(const QPdfWriter* self);
extern __declspec(dllexport) int QPdfWriter_Metric(const QPdfWriter* self, PaintDeviceMetric id);
extern __declspec(dllexport) struct miqt_string QPdfWriter_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QPdfWriter_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QPdfWriter_AddFileAttachment3(QPdfWriter* self, struct miqt_string fileName, struct miqt_string data, struct miqt_string mimeType);
extern __declspec(dllexport) void QPdfWriter_override_virtual_NewPage(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_NewPage(void* self);
extern __declspec(dllexport) void QPdfWriter_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QPdfWriter_virtualbase_PaintEngine(const void* self);
extern __declspec(dllexport) void QPdfWriter_override_virtual_Metric(void* self, intptr_t slot);
int QPdfWriter_virtualbase_Metric(const void* self, PaintDeviceMetric id);
extern __declspec(dllexport) void QPdfWriter_override_virtual_Event(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_Event(void* self, QEvent* event);
extern __declspec(dllexport) void QPdfWriter_override_virtual_EventFilter(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_EventFilter(void* self, QObject* watched, QEvent* event);
extern __declspec(dllexport) void QPdfWriter_override_virtual_TimerEvent(void* self, intptr_t slot);
void QPdfWriter_virtualbase_TimerEvent(void* self, QTimerEvent* event);
extern __declspec(dllexport) void QPdfWriter_override_virtual_ChildEvent(void* self, intptr_t slot);
void QPdfWriter_virtualbase_ChildEvent(void* self, QChildEvent* event);
extern __declspec(dllexport) void QPdfWriter_override_virtual_CustomEvent(void* self, intptr_t slot);
void QPdfWriter_virtualbase_CustomEvent(void* self, QEvent* event);
extern __declspec(dllexport) void QPdfWriter_override_virtual_ConnectNotify(void* self, intptr_t slot);
void QPdfWriter_virtualbase_ConnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QPdfWriter_override_virtual_DisconnectNotify(void* self, intptr_t slot);
void QPdfWriter_virtualbase_DisconnectNotify(void* self, QMetaMethod* signal);
extern __declspec(dllexport) void QPdfWriter_override_virtual_SetPageLayout(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_SetPageLayout(void* self, QPageLayout* pageLayout);
extern __declspec(dllexport) void QPdfWriter_override_virtual_SetPageSize(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_SetPageSize(void* self, QPageSize* pageSize);
extern __declspec(dllexport) void QPdfWriter_override_virtual_SetPageOrientation(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_SetPageOrientation(void* self, int orientation);
extern __declspec(dllexport) void QPdfWriter_override_virtual_SetPageMargins(void* self, intptr_t slot);
bool QPdfWriter_virtualbase_SetPageMargins(void* self, QMarginsF* margins, int units);
extern __declspec(dllexport) void QPdfWriter_override_virtual_SetPageRanges(void* self, intptr_t slot);
void QPdfWriter_virtualbase_SetPageRanges(void* self, QPageRanges* ranges);
extern __declspec(dllexport) void QPdfWriter_Delete(QPdfWriter* self, bool isSubclass);

} 
