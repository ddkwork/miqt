#pragma once
#ifndef MIQT_QT_GEN_QICON_H
#define MIQT_QT_GEN_QICON_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

//#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "../libmiqt/libmiqt.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QIcon;
class QIconEngine;
class QPainter;
class QPixmap;
class QRect;
class QSize;
class QWindow;
class _GUID;
class type_info;
#else
typedef struct QIcon QIcon;
typedef struct QIconEngine QIconEngine;
typedef struct QPainter QPainter;
typedef struct QPixmap QPixmap;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QWindow QWindow;
typedef struct _GUID _GUID;
typedef struct type_info type_info;
#endif

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QIcon* QIcon_new();
extern __declspec(dllexport) QIcon* QIcon_new2(QPixmap* pixmap);
extern __declspec(dllexport) QIcon* QIcon_new3(QIcon* other);
extern __declspec(dllexport) QIcon* QIcon_new4(struct miqt_string fileName);
extern __declspec(dllexport) QIcon* QIcon_new5(QIconEngine* engine);
extern __declspec(dllexport) void QIcon_OperatorAssign(QIcon* self, QIcon* other);
extern __declspec(dllexport) void QIcon_Swap(QIcon* self, QIcon* other);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap(const QIcon* self, QSize* size);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap2(const QIcon* self, int w, int h);
extern __declspec(dllexport) QPixmap* QIcon_PixmapWithExtent(const QIcon* self, int extent);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap3(const QIcon* self, QSize* size, double devicePixelRatio);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap4(const QIcon* self, QWindow* window, QSize* size);
extern __declspec(dllexport) QSize* QIcon_ActualSize(const QIcon* self, QSize* size);
extern __declspec(dllexport) QSize* QIcon_ActualSize2(const QIcon* self, QWindow* window, QSize* size);
extern __declspec(dllexport) struct miqt_string QIcon_Name(const QIcon* self);
extern __declspec(dllexport) void QIcon_Paint(const QIcon* self, QPainter* painter, QRect* rect);
extern __declspec(dllexport) void QIcon_Paint2(const QIcon* self, QPainter* painter, int x, int y, int w, int h);
extern __declspec(dllexport) bool QIcon_IsNull(const QIcon* self);
extern __declspec(dllexport) bool QIcon_IsDetached(const QIcon* self);
extern __declspec(dllexport) void QIcon_Detach(QIcon* self);
extern __declspec(dllexport) long long QIcon_CacheKey(const QIcon* self);
extern __declspec(dllexport) void QIcon_AddPixmap(QIcon* self, QPixmap* pixmap);
extern __declspec(dllexport) void QIcon_AddFile(QIcon* self, struct miqt_string fileName);
extern __declspec(dllexport) struct miqt_array /* of QSize* */  QIcon_AvailableSizes(const QIcon* self);
extern __declspec(dllexport) void QIcon_SetIsMask(QIcon* self, bool isMask);
extern __declspec(dllexport) bool QIcon_IsMask(const QIcon* self);
extern __declspec(dllexport) QIcon* QIcon_FromTheme(struct miqt_string name);
extern __declspec(dllexport) QIcon* QIcon_FromTheme2(struct miqt_string name, QIcon* fallback);
extern __declspec(dllexport) bool QIcon_HasThemeIcon(struct miqt_string name);
extern __declspec(dllexport) QIcon* QIcon_FromThemeWithIcon(ThemeIcon icon);
extern __declspec(dllexport) QIcon* QIcon_FromTheme3(ThemeIcon icon, QIcon* fallback);
extern __declspec(dllexport) bool QIcon_HasThemeIconWithIcon(ThemeIcon icon);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QIcon_ThemeSearchPaths();
extern __declspec(dllexport) void QIcon_SetThemeSearchPaths(struct miqt_array /* of struct miqt_string */  searchpath);
extern __declspec(dllexport) struct miqt_array /* of struct miqt_string */  QIcon_FallbackSearchPaths();
extern __declspec(dllexport) void QIcon_SetFallbackSearchPaths(struct miqt_array /* of struct miqt_string */  paths);
extern __declspec(dllexport) struct miqt_string QIcon_ThemeName();
extern __declspec(dllexport) void QIcon_SetThemeName(struct miqt_string path);
extern __declspec(dllexport) struct miqt_string QIcon_FallbackThemeName();
extern __declspec(dllexport) void QIcon_SetFallbackThemeName(struct miqt_string name);
extern __declspec(dllexport) DataPtr* QIcon_DataPtr(QIcon* self);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap22(const QIcon* self, QSize* size, Mode mode);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap32(const QIcon* self, QSize* size, Mode mode, State state);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap33(const QIcon* self, int w, int h, Mode mode);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap42(const QIcon* self, int w, int h, Mode mode, State state);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap23(const QIcon* self, int extent, Mode mode);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap34(const QIcon* self, int extent, Mode mode, State state);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap35(const QIcon* self, QSize* size, double devicePixelRatio, Mode mode);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap43(const QIcon* self, QSize* size, double devicePixelRatio, Mode mode, State state);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap36(const QIcon* self, QWindow* window, QSize* size, Mode mode);
extern __declspec(dllexport) QPixmap* QIcon_Pixmap44(const QIcon* self, QWindow* window, QSize* size, Mode mode, State state);
extern __declspec(dllexport) QSize* QIcon_ActualSize22(const QIcon* self, QSize* size, Mode mode);
extern __declspec(dllexport) QSize* QIcon_ActualSize3(const QIcon* self, QSize* size, Mode mode, State state);
extern __declspec(dllexport) QSize* QIcon_ActualSize32(const QIcon* self, QWindow* window, QSize* size, Mode mode);
extern __declspec(dllexport) QSize* QIcon_ActualSize4(const QIcon* self, QWindow* window, QSize* size, Mode mode, State state);
extern __declspec(dllexport) void QIcon_Paint3(const QIcon* self, QPainter* painter, QRect* rect, int alignment);
extern __declspec(dllexport) void QIcon_Paint4(const QIcon* self, QPainter* painter, QRect* rect, int alignment, Mode mode);
extern __declspec(dllexport) void QIcon_Paint5(const QIcon* self, QPainter* painter, QRect* rect, int alignment, Mode mode, State state);
extern __declspec(dllexport) void QIcon_Paint6(const QIcon* self, QPainter* painter, int x, int y, int w, int h, int alignment);
extern __declspec(dllexport) void QIcon_Paint7(const QIcon* self, QPainter* painter, int x, int y, int w, int h, int alignment, Mode mode);
extern __declspec(dllexport) void QIcon_Paint8(const QIcon* self, QPainter* painter, int x, int y, int w, int h, int alignment, Mode mode, State state);
extern __declspec(dllexport) void QIcon_AddPixmap2(QIcon* self, QPixmap* pixmap, Mode mode);
extern __declspec(dllexport) void QIcon_AddPixmap3(QIcon* self, QPixmap* pixmap, Mode mode, State state);
extern __declspec(dllexport) void QIcon_AddFile2(QIcon* self, struct miqt_string fileName, QSize* size);
extern __declspec(dllexport) void QIcon_AddFile3(QIcon* self, struct miqt_string fileName, QSize* size, Mode mode);
extern __declspec(dllexport) void QIcon_AddFile4(QIcon* self, struct miqt_string fileName, QSize* size, Mode mode, State state);
extern __declspec(dllexport) struct miqt_array /* of QSize* */  QIcon_AvailableSizes1(const QIcon* self, Mode mode);
extern __declspec(dllexport) struct miqt_array /* of QSize* */  QIcon_AvailableSizes2(const QIcon* self, Mode mode, State state);
extern __declspec(dllexport) void QIcon_Delete(QIcon* self, bool isSubclass);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif
