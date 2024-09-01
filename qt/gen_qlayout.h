#ifndef GEN_QLAYOUT_H
#define GEN_QLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QLayout;
class QLayoutItem;
class QMargins;
class QMetaObject;
class QRect;
class QSize;
class QWidget;
#else
typedef struct QLayout QLayout;
typedef struct QLayoutItem QLayoutItem;
typedef struct QMargins QMargins;
typedef struct QMetaObject QMetaObject;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QWidget QWidget;
#endif

QMetaObject* QLayout_MetaObject(QLayout* self);
void QLayout_Tr(const char* s, char** _out, int* _out_Strlen);
void QLayout_TrUtf8(const char* s, char** _out, int* _out_Strlen);
int QLayout_Margin(QLayout* self);
void QLayout_SetMargin(QLayout* self, int margin);
int QLayout_Spacing(QLayout* self);
void QLayout_SetSpacing(QLayout* self, int spacing);
void QLayout_SetContentsMargins(QLayout* self, int left, int top, int right, int bottom);
void QLayout_SetContentsMarginsWithMargins(QLayout* self, QMargins* margins);
void QLayout_GetContentsMargins(QLayout* self, int* left, int* top, int* right, int* bottom);
QMargins* QLayout_ContentsMargins(QLayout* self);
QRect* QLayout_ContentsRect(QLayout* self);
bool QLayout_SetAlignment(QLayout* self, QWidget* w, int alignment);
bool QLayout_SetAlignment2(QLayout* self, QLayout* l, int alignment);
void QLayout_SetSizeConstraint(QLayout* self, uintptr_t sizeConstraint);
uintptr_t QLayout_SizeConstraint(QLayout* self);
void QLayout_SetMenuBar(QLayout* self, QWidget* w);
QWidget* QLayout_MenuBar(QLayout* self);
QWidget* QLayout_ParentWidget(QLayout* self);
void QLayout_Invalidate(QLayout* self);
QRect* QLayout_Geometry(QLayout* self);
bool QLayout_Activate(QLayout* self);
void QLayout_Update(QLayout* self);
void QLayout_AddWidget(QLayout* self, QWidget* w);
void QLayout_AddItem(QLayout* self, QLayoutItem* param1);
void QLayout_RemoveWidget(QLayout* self, QWidget* w);
void QLayout_RemoveItem(QLayout* self, QLayoutItem* param1);
int QLayout_ExpandingDirections(QLayout* self);
QSize* QLayout_MinimumSize(QLayout* self);
QSize* QLayout_MaximumSize(QLayout* self);
void QLayout_SetGeometry(QLayout* self, QRect* geometry);
QLayoutItem* QLayout_ItemAt(QLayout* self, int index);
QLayoutItem* QLayout_TakeAt(QLayout* self, int index);
int QLayout_IndexOf(QLayout* self, QWidget* param1);
int QLayout_IndexOfWithQLayoutItem(QLayout* self, QLayoutItem* param1);
int QLayout_Count(QLayout* self);
bool QLayout_IsEmpty(QLayout* self);
int QLayout_ControlTypes(QLayout* self);
QLayoutItem* QLayout_ReplaceWidget(QLayout* self, QWidget* from, QWidget* to);
int QLayout_TotalHeightForWidth(QLayout* self, int w);
QSize* QLayout_TotalMinimumSize(QLayout* self);
QSize* QLayout_TotalMaximumSize(QLayout* self);
QSize* QLayout_TotalSizeHint(QLayout* self);
QLayout* QLayout_Layout(QLayout* self);
void QLayout_SetEnabled(QLayout* self, bool enabled);
bool QLayout_IsEnabled(QLayout* self);
QSize* QLayout_ClosestAcceptableSize(QWidget* w, QSize* s);
void QLayout_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QLayout_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QLayout_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QLayout_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
QLayoutItem* QLayout_ReplaceWidget3(QLayout* self, QWidget* from, QWidget* to, int options);
void QLayout_Delete(QLayout* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif