#ifndef GEN_QGRIDLAYOUT_H
#define GEN_QGRIDLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QGridLayout;
class QLayout;
class QLayoutItem;
class QMetaObject;
class QRect;
class QSize;
class QWidget;
#else
typedef struct QGridLayout QGridLayout;
typedef struct QLayout QLayout;
typedef struct QLayoutItem QLayoutItem;
typedef struct QMetaObject QMetaObject;
typedef struct QRect QRect;
typedef struct QSize QSize;
typedef struct QWidget QWidget;
#endif

QGridLayout* QGridLayout_new(QWidget* parent);
QGridLayout* QGridLayout_new2();
QMetaObject* QGridLayout_MetaObject(QGridLayout* self);
void QGridLayout_Tr(const char* s, char** _out, int* _out_Strlen);
void QGridLayout_TrUtf8(const char* s, char** _out, int* _out_Strlen);
QSize* QGridLayout_SizeHint(QGridLayout* self);
QSize* QGridLayout_MinimumSize(QGridLayout* self);
QSize* QGridLayout_MaximumSize(QGridLayout* self);
void QGridLayout_SetHorizontalSpacing(QGridLayout* self, int spacing);
int QGridLayout_HorizontalSpacing(QGridLayout* self);
void QGridLayout_SetVerticalSpacing(QGridLayout* self, int spacing);
int QGridLayout_VerticalSpacing(QGridLayout* self);
void QGridLayout_SetSpacing(QGridLayout* self, int spacing);
int QGridLayout_Spacing(QGridLayout* self);
void QGridLayout_SetRowStretch(QGridLayout* self, int row, int stretch);
void QGridLayout_SetColumnStretch(QGridLayout* self, int column, int stretch);
int QGridLayout_RowStretch(QGridLayout* self, int row);
int QGridLayout_ColumnStretch(QGridLayout* self, int column);
void QGridLayout_SetRowMinimumHeight(QGridLayout* self, int row, int minSize);
void QGridLayout_SetColumnMinimumWidth(QGridLayout* self, int column, int minSize);
int QGridLayout_RowMinimumHeight(QGridLayout* self, int row);
int QGridLayout_ColumnMinimumWidth(QGridLayout* self, int column);
int QGridLayout_ColumnCount(QGridLayout* self);
int QGridLayout_RowCount(QGridLayout* self);
QRect* QGridLayout_CellRect(QGridLayout* self, int row, int column);
bool QGridLayout_HasHeightForWidth(QGridLayout* self);
int QGridLayout_HeightForWidth(QGridLayout* self, int param1);
int QGridLayout_MinimumHeightForWidth(QGridLayout* self, int param1);
int QGridLayout_ExpandingDirections(QGridLayout* self);
void QGridLayout_Invalidate(QGridLayout* self);
void QGridLayout_AddWidget(QGridLayout* self, QWidget* w);
void QGridLayout_AddWidget2(QGridLayout* self, QWidget* param1, int row, int column);
void QGridLayout_AddWidget3(QGridLayout* self, QWidget* param1, int row, int column, int rowSpan, int columnSpan);
void QGridLayout_AddLayout(QGridLayout* self, QLayout* param1, int row, int column);
void QGridLayout_AddLayout2(QGridLayout* self, QLayout* param1, int row, int column, int rowSpan, int columnSpan);
void QGridLayout_SetOriginCorner(QGridLayout* self, uintptr_t originCorner);
uintptr_t QGridLayout_OriginCorner(QGridLayout* self);
QLayoutItem* QGridLayout_ItemAt(QGridLayout* self, int index);
QLayoutItem* QGridLayout_ItemAtPosition(QGridLayout* self, int row, int column);
QLayoutItem* QGridLayout_TakeAt(QGridLayout* self, int index);
int QGridLayout_Count(QGridLayout* self);
void QGridLayout_SetGeometry(QGridLayout* self, QRect* geometry);
void QGridLayout_AddItem(QGridLayout* self, QLayoutItem* item, int row, int column);
void QGridLayout_SetDefaultPositioning(QGridLayout* self, int n, uintptr_t orient);
void QGridLayout_GetItemPosition(QGridLayout* self, int idx, int* row, int* column, int* rowSpan, int* columnSpan);
void QGridLayout_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen);
void QGridLayout_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QGridLayout_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen);
void QGridLayout_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen);
void QGridLayout_AddWidget4(QGridLayout* self, QWidget* param1, int row, int column, int param4);
void QGridLayout_AddWidget6(QGridLayout* self, QWidget* param1, int row, int column, int rowSpan, int columnSpan, int param6);
void QGridLayout_AddLayout4(QGridLayout* self, QLayout* param1, int row, int column, int param4);
void QGridLayout_AddLayout6(QGridLayout* self, QLayout* param1, int row, int column, int rowSpan, int columnSpan, int param6);
void QGridLayout_AddItem4(QGridLayout* self, QLayoutItem* item, int row, int column, int rowSpan);
void QGridLayout_AddItem5(QGridLayout* self, QLayoutItem* item, int row, int column, int rowSpan, int columnSpan);
void QGridLayout_AddItem6(QGridLayout* self, QLayoutItem* item, int row, int column, int rowSpan, int columnSpan, int param6);
void QGridLayout_Delete(QGridLayout* self);

#ifdef __cplusplus
} /* extern C */
#endif 

#endif