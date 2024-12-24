#pragma once
#ifndef MIQT_QT_GEN_QSTATICTEXT_H
#define MIQT_QT_GEN_QSTATICTEXT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QFont;
class QSizeF;
class QStaticText;
class QTextOption;
class QTransform;
class _GUID;
class type_info;
#else
typedef struct QFont QFont;
typedef struct QSizeF QSizeF;
typedef struct QStaticText QStaticText;
typedef struct QTextOption QTextOption;
typedef struct QTransform QTransform;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QStaticText* QStaticText_new();
extern __declspec(dllexport) QStaticText* QStaticText_new2(struct miqt_string text);
extern __declspec(dllexport) QStaticText* QStaticText_new3(QStaticText* other);
extern __declspec(dllexport) void QStaticText_OperatorAssign(QStaticText* self, QStaticText* param1);
extern __declspec(dllexport) void QStaticText_Swap(QStaticText* self, QStaticText* other);
extern __declspec(dllexport) void QStaticText_SetText(QStaticText* self, struct miqt_string text);
extern __declspec(dllexport) struct miqt_string QStaticText_Text(const QStaticText* self);
extern __declspec(dllexport) void QStaticText_SetTextFormat(QStaticText* self, int textFormat);
extern __declspec(dllexport) int QStaticText_TextFormat(const QStaticText* self);
extern __declspec(dllexport) void QStaticText_SetTextWidth(QStaticText* self, double textWidth);
extern __declspec(dllexport) double QStaticText_TextWidth(const QStaticText* self);
extern __declspec(dllexport) void QStaticText_SetTextOption(QStaticText* self, QTextOption* textOption);
extern __declspec(dllexport) QTextOption* QStaticText_TextOption(const QStaticText* self);
extern __declspec(dllexport) QSizeF* QStaticText_Size(const QStaticText* self);
extern __declspec(dllexport) void QStaticText_Prepare(QStaticText* self);
extern __declspec(dllexport) void QStaticText_SetPerformanceHint(QStaticText* self, PerformanceHint performanceHint);
extern __declspec(dllexport) PerformanceHint QStaticText_PerformanceHint(const QStaticText* self);
extern __declspec(dllexport) bool QStaticText_OperatorEqual(const QStaticText* self, QStaticText* param1);
extern __declspec(dllexport) bool QStaticText_OperatorNotEqual(const QStaticText* self, QStaticText* param1);
extern __declspec(dllexport) void QStaticText_Prepare1(QStaticText* self, QTransform* matrix);
extern __declspec(dllexport) void QStaticText_Prepare2(QStaticText* self, QTransform* matrix, QFont* font);
extern __declspec(dllexport) void QStaticText_Delete(QStaticText* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
