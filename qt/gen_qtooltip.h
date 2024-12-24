#pragma once
#ifndef MIQT_QT_GEN_QTOOLTIP_H
#define MIQT_QT_GEN_QTOOLTIP_H

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
class QPalette;
class QPoint;
class QRect;
class QToolTip;
class QWidget;
class _GUID;
class type_info;
#else
typedef struct QFont QFont;
typedef struct QPalette QPalette;
typedef struct QPoint QPoint;
typedef struct QRect QRect;
typedef struct QToolTip QToolTip;
typedef struct QWidget QWidget;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) void QToolTip_ShowText(QPoint* pos, struct miqt_string text);
extern __declspec(dllexport) void QToolTip_HideText();
extern __declspec(dllexport) bool QToolTip_IsVisible();
extern __declspec(dllexport) struct miqt_string QToolTip_Text();
extern __declspec(dllexport) QPalette* QToolTip_Palette();
extern __declspec(dllexport) void QToolTip_SetPalette(QPalette* palette);
extern __declspec(dllexport) QFont* QToolTip_Font();
extern __declspec(dllexport) void QToolTip_SetFont(QFont* font);
extern __declspec(dllexport) void QToolTip_ShowText3(QPoint* pos, struct miqt_string text, QWidget* w);
extern __declspec(dllexport) void QToolTip_ShowText4(QPoint* pos, struct miqt_string text, QWidget* w, QRect* rect);
extern __declspec(dllexport) void QToolTip_ShowText5(QPoint* pos, struct miqt_string text, QWidget* w, QRect* rect, int msecShowTime);
extern __declspec(dllexport) void QToolTip_Delete(QToolTip* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
