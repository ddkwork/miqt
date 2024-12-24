#pragma once
#ifndef MIQT_QT_PRINTSUPPORT_GEN_QPRINTER_H
#define MIQT_QT_PRINTSUPPORT_GEN_QPRINTER_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QMarginsF;
class QPageLayout;
class QPageRanges;
class QPageSize;
class QPagedPaintDevice;
class QPaintDevice;
class QPaintEngine;
class QPrintEngine;
class QPrinter;
class QPrinterInfo;
class QRectF;
class _GUID;
class type_info;
#else
typedef struct QMarginsF QMarginsF;
typedef struct QPageLayout QPageLayout;
typedef struct QPageRanges QPageRanges;
typedef struct QPageSize QPageSize;
typedef struct QPagedPaintDevice QPagedPaintDevice;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPaintEngine QPaintEngine;
typedef struct QPrintEngine QPrintEngine;
typedef struct QPrinter QPrinter;
typedef struct QPrinterInfo QPrinterInfo;
typedef struct QRectF QRectF;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPrinter* QPrinter_new();
extern __declspec(dllexport) QPrinter* QPrinter_new2(QPrinterInfo* printer);
extern __declspec(dllexport) QPrinter* QPrinter_new3(PrinterMode mode);
extern __declspec(dllexport) QPrinter* QPrinter_new4(QPrinterInfo* printer, PrinterMode mode);
extern __declspec(dllexport) void QPrinter_virtbase(QPrinter* src, QPagedPaintDevice** outptr_QPagedPaintDevice);
extern __declspec(dllexport) int QPrinter_DevType(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetOutputFormat(QPrinter* self, OutputFormat format);
extern __declspec(dllexport) OutputFormat QPrinter_OutputFormat(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetPdfVersion(QPrinter* self, PdfVersion version);
extern __declspec(dllexport) PdfVersion QPrinter_PdfVersion(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetPrinterName(QPrinter* self, struct miqt_string printerName);
extern __declspec(dllexport) struct miqt_string QPrinter_PrinterName(const QPrinter* self);
extern __declspec(dllexport) bool QPrinter_IsValid(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetOutputFileName(QPrinter* self, struct miqt_string outputFileName);
extern __declspec(dllexport) struct miqt_string QPrinter_OutputFileName(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetPrintProgram(QPrinter* self, struct miqt_string printProgram);
extern __declspec(dllexport) struct miqt_string QPrinter_PrintProgram(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetDocName(QPrinter* self, struct miqt_string docName);
extern __declspec(dllexport) struct miqt_string QPrinter_DocName(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetCreator(QPrinter* self, struct miqt_string creator);
extern __declspec(dllexport) struct miqt_string QPrinter_Creator(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetPageOrder(QPrinter* self, PageOrder pageOrder);
extern __declspec(dllexport) PageOrder QPrinter_PageOrder(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetResolution(QPrinter* self, int resolution);
extern __declspec(dllexport) int QPrinter_Resolution(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetColorMode(QPrinter* self, ColorMode colorMode);
extern __declspec(dllexport) ColorMode QPrinter_ColorMode(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetCollateCopies(QPrinter* self, bool collate);
extern __declspec(dllexport) bool QPrinter_CollateCopies(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetFullPage(QPrinter* self, bool fullPage);
extern __declspec(dllexport) bool QPrinter_FullPage(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetCopyCount(QPrinter* self, int copyCount);
extern __declspec(dllexport) int QPrinter_CopyCount(const QPrinter* self);
extern __declspec(dllexport) bool QPrinter_SupportsMultipleCopies(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetPaperSource(QPrinter* self, PaperSource paperSource);
extern __declspec(dllexport) PaperSource QPrinter_PaperSource(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetDuplex(QPrinter* self, DuplexMode duplex);
extern __declspec(dllexport) DuplexMode QPrinter_Duplex(const QPrinter* self);
extern __declspec(dllexport) struct miqt_array /* of int */  QPrinter_SupportedResolutions(const QPrinter* self);
extern __declspec(dllexport) struct miqt_array /* of PaperSource */  QPrinter_SupportedPaperSources(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetFontEmbeddingEnabled(QPrinter* self, bool enable);
extern __declspec(dllexport) bool QPrinter_FontEmbeddingEnabled(const QPrinter* self);
extern __declspec(dllexport) QRectF* QPrinter_PaperRect(const QPrinter* self, Unit param1);
extern __declspec(dllexport) QRectF* QPrinter_PageRect(const QPrinter* self, Unit param1);
extern __declspec(dllexport) struct miqt_string QPrinter_PrinterSelectionOption(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetPrinterSelectionOption(QPrinter* self, struct miqt_string printerSelectionOption);
extern __declspec(dllexport) bool QPrinter_NewPage(QPrinter* self);
extern __declspec(dllexport) bool QPrinter_Abort(QPrinter* self);
extern __declspec(dllexport) PrinterState QPrinter_PrinterState(const QPrinter* self);
extern __declspec(dllexport) QPaintEngine* QPrinter_PaintEngine(const QPrinter* self);
extern __declspec(dllexport) QPrintEngine* QPrinter_PrintEngine(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetFromTo(QPrinter* self, int fromPage, int toPage);
extern __declspec(dllexport) int QPrinter_FromPage(const QPrinter* self);
extern __declspec(dllexport) int QPrinter_ToPage(const QPrinter* self);
extern __declspec(dllexport) void QPrinter_SetPrintRange(QPrinter* self, PrintRange rangeVal);
extern __declspec(dllexport) PrintRange QPrinter_PrintRange(const QPrinter* self);
extern __declspec(dllexport) int QPrinter_Metric(const QPrinter* self, PaintDeviceMetric param1);
extern __declspec(dllexport) void QPrinter_override_virtual_DevType(void* self, intptr_t slot);
int QPrinter_virtualbase_DevType(const void* self);
extern __declspec(dllexport) void QPrinter_override_virtual_NewPage(void* self, intptr_t slot);
bool QPrinter_virtualbase_NewPage(void* self);
extern __declspec(dllexport) void QPrinter_override_virtual_PaintEngine(void* self, intptr_t slot);
QPaintEngine* QPrinter_virtualbase_PaintEngine(const void* self);
extern __declspec(dllexport) void QPrinter_override_virtual_Metric(void* self, intptr_t slot);
int QPrinter_virtualbase_Metric(const void* self, PaintDeviceMetric param1);
extern __declspec(dllexport) void QPrinter_override_virtual_SetPageLayout(void* self, intptr_t slot);
bool QPrinter_virtualbase_SetPageLayout(void* self, QPageLayout* pageLayout);
extern __declspec(dllexport) void QPrinter_override_virtual_SetPageSize(void* self, intptr_t slot);
bool QPrinter_virtualbase_SetPageSize(void* self, QPageSize* pageSize);
extern __declspec(dllexport) void QPrinter_override_virtual_SetPageOrientation(void* self, intptr_t slot);
bool QPrinter_virtualbase_SetPageOrientation(void* self, int orientation);
extern __declspec(dllexport) void QPrinter_override_virtual_SetPageMargins(void* self, intptr_t slot);
bool QPrinter_virtualbase_SetPageMargins(void* self, QMarginsF* margins, int units);
extern __declspec(dllexport) void QPrinter_override_virtual_SetPageRanges(void* self, intptr_t slot);
void QPrinter_virtualbase_SetPageRanges(void* self, QPageRanges* ranges);
extern __declspec(dllexport) void QPrinter_Delete(QPrinter* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
