#pragma once
#ifndef MIQT_QT_PRINTSUPPORT_GEN_QPRINTERINFO_H
#define MIQT_QT_PRINTSUPPORT_GEN_QPRINTERINFO_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QPageSize;
class QPrinter;
class QPrinterInfo;
class _GUID;
class type_info;
#else
typedef struct QPageSize QPageSize;
typedef struct QPrinter QPrinter;
typedef struct QPrinterInfo QPrinterInfo;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QPrinterInfo* QPrinterInfo_new();
extern __declspec(dllexport) QPrinterInfo* QPrinterInfo_new2(QPrinterInfo* other);
extern __declspec(dllexport) QPrinterInfo* QPrinterInfo_new3(QPrinter* printer);
extern __declspec(dllexport) void QPrinterInfo_OperatorAssign(QPrinterInfo* self, QPrinterInfo* other);
extern __declspec(dllexport) struct miqt_string QPrinterInfo_PrinterName(const QPrinterInfo* self);
extern __declspec(dllexport) struct miqt_string QPrinterInfo_Description(const QPrinterInfo* self);
extern __declspec(dllexport) struct miqt_string QPrinterInfo_Location(const QPrinterInfo* self);
extern __declspec(dllexport) struct miqt_string QPrinterInfo_MakeAndModel(const QPrinterInfo* self);
extern __declspec(dllexport) bool QPrinterInfo_IsNull(const QPrinterInfo* self);
extern __declspec(dllexport) bool QPrinterInfo_IsDefault(const QPrinterInfo* self);
extern __declspec(dllexport) bool QPrinterInfo_IsRemote(const QPrinterInfo* self);
extern __declspec(dllexport) int QPrinterInfo_State(const QPrinterInfo* self);
extern __declspec(dllexport) struct miqt_array /* of QPageSize* */  QPrinterInfo_SupportedPageSizes(const QPrinterInfo* self);
extern __declspec(dllexport) QPageSize* QPrinterInfo_DefaultPageSize(const QPrinterInfo* self);
extern __declspec(dllexport) bool QPrinterInfo_SupportsCustomPageSizes(const QPrinterInfo* self);
extern __declspec(dllexport) QPageSize* QPrinterInfo_MinimumPhysicalPageSize(const QPrinterInfo* self);
extern __declspec(dllexport) QPageSize* QPrinterInfo_MaximumPhysicalPageSize(const QPrinterInfo* self);
extern __declspec(dllexport) struct miqt_array /* of int */  QPrinterInfo_SupportedResolutions(const QPrinterInfo* self);
extern __declspec(dllexport) int QPrinterInfo_DefaultDuplexMode(const QPrinterInfo* self);
extern __declspec(dllexport) struct miqt_array /* of int */  QPrinterInfo_SupportedDuplexModes(const QPrinterInfo* self);
extern __declspec(dllexport) int QPrinterInfo_DefaultColorMode(const QPrinterInfo* self);
extern __declspec(dllexport) struct miqt_array /* of int */  QPrinterInfo_SupportedColorModes(const QPrinterInfo* self);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QPrinterInfo_AvailablePrinterNames();
extern __declspec(dllexport) struct miqt_array /* of QPrinterInfo* */  QPrinterInfo_AvailablePrinters();
extern __declspec(dllexport) struct miqt_string QPrinterInfo_DefaultPrinterName();
extern __declspec(dllexport) QPrinterInfo* QPrinterInfo_DefaultPrinter();
extern __declspec(dllexport) QPrinterInfo* QPrinterInfo_PrinterInfo(struct miqt_string printerName);
extern __declspec(dllexport) void QPrinterInfo_Delete(QPrinterInfo* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
