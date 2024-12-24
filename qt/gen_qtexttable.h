#pragma once
#ifndef MIQT_QT_GEN_QTEXTTABLE_H
#define MIQT_QT_GEN_QTEXTTABLE_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QTextFrame__iterator)
typedef QTextFrame::iterator QTextFrame__iterator;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTextCharFormat QTextCharFormat;
typedef struct QTextCursor QTextCursor;
typedef struct QTextDocument QTextDocument;
typedef struct QTextFrame QTextFrame;
typedef struct QTextFrame__iterator QTextFrame__iterator;
typedef struct QTextObject QTextObject;
typedef struct QTextTable QTextTable;
typedef struct QTextTableCell QTextTableCell;
typedef struct QTextTableFormat QTextTableFormat;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTextTableCell* QTextTableCell_new();
extern __declspec(dllexport) QTextTableCell* QTextTableCell_new2(QTextTableCell* o);
extern __declspec(dllexport) void QTextTableCell_OperatorAssign(QTextTableCell* self, QTextTableCell* o);
extern __declspec(dllexport) void QTextTableCell_SetFormat(QTextTableCell* self, QTextCharFormat* format);
extern __declspec(dllexport) QTextCharFormat* QTextTableCell_Format(const QTextTableCell* self);
extern __declspec(dllexport) int QTextTableCell_Row(const QTextTableCell* self);
extern __declspec(dllexport) int QTextTableCell_Column(const QTextTableCell* self);
extern __declspec(dllexport) int QTextTableCell_RowSpan(const QTextTableCell* self);
extern __declspec(dllexport) int QTextTableCell_ColumnSpan(const QTextTableCell* self);
extern __declspec(dllexport) bool QTextTableCell_IsValid(const QTextTableCell* self);
extern __declspec(dllexport) QTextCursor* QTextTableCell_FirstCursorPosition(const QTextTableCell* self);
extern __declspec(dllexport) QTextCursor* QTextTableCell_LastCursorPosition(const QTextTableCell* self);
extern __declspec(dllexport) int QTextTableCell_FirstPosition(const QTextTableCell* self);
extern __declspec(dllexport) int QTextTableCell_LastPosition(const QTextTableCell* self);
extern __declspec(dllexport) bool QTextTableCell_OperatorEqual(const QTextTableCell* self, QTextTableCell* other);
extern __declspec(dllexport) bool QTextTableCell_OperatorNotEqual(const QTextTableCell* self, QTextTableCell* other);
extern __declspec(dllexport) QTextFrame__iterator* QTextTableCell_Begin(const QTextTableCell* self);
extern __declspec(dllexport) QTextFrame__iterator* QTextTableCell_End(const QTextTableCell* self);
extern __declspec(dllexport) int QTextTableCell_TableCellFormatIndex(const QTextTableCell* self);
extern __declspec(dllexport) void QTextTableCell_Delete(QTextTableCell* self, bool isSubclass);

extern __declspec(dllexport) QTextTable* QTextTable_new(QTextDocument* doc);
extern __declspec(dllexport) void QTextTable_virtbase(QTextTable* src, QTextFrame** outptr_QTextFrame);
extern __declspec(dllexport) QMetaObject* QTextTable_MetaObject(const QTextTable* self);
extern __declspec(dllexport) void* QTextTable_Metacast(QTextTable* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QTextTable_Tr(const char* s);
extern __declspec(dllexport) void QTextTable_Resize(QTextTable* self, int rows, int cols);
extern __declspec(dllexport) void QTextTable_InsertRows(QTextTable* self, int pos, int num);
extern __declspec(dllexport) void QTextTable_InsertColumns(QTextTable* self, int pos, int num);
extern __declspec(dllexport) void QTextTable_AppendRows(QTextTable* self, int count);
extern __declspec(dllexport) void QTextTable_AppendColumns(QTextTable* self, int count);
extern __declspec(dllexport) void QTextTable_RemoveRows(QTextTable* self, int pos, int num);
extern __declspec(dllexport) void QTextTable_RemoveColumns(QTextTable* self, int pos, int num);
extern __declspec(dllexport) void QTextTable_MergeCells(QTextTable* self, int row, int col, int numRows, int numCols);
extern __declspec(dllexport) void QTextTable_MergeCellsWithCursor(QTextTable* self, QTextCursor* cursor);
extern __declspec(dllexport) void QTextTable_SplitCell(QTextTable* self, int row, int col, int numRows, int numCols);
extern __declspec(dllexport) int QTextTable_Rows(const QTextTable* self);
extern __declspec(dllexport) int QTextTable_Columns(const QTextTable* self);
extern __declspec(dllexport) QTextTableCell* QTextTable_CellAt(const QTextTable* self, int row, int col);
extern __declspec(dllexport) QTextTableCell* QTextTable_CellAtWithPosition(const QTextTable* self, int position);
extern __declspec(dllexport) QTextTableCell* QTextTable_CellAtWithQTextCursor(const QTextTable* self, QTextCursor* c);
extern __declspec(dllexport) QTextCursor* QTextTable_RowStart(const QTextTable* self, QTextCursor* c);
extern __declspec(dllexport) QTextCursor* QTextTable_RowEnd(const QTextTable* self, QTextCursor* c);
extern __declspec(dllexport) void QTextTable_SetFormat(QTextTable* self, QTextTableFormat* format);
extern __declspec(dllexport) QTextTableFormat* QTextTable_Format(const QTextTable* self);
extern __declspec(dllexport) struct miqt_string QTextTable_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QTextTable_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QTextTable_Delete(QTextTable* self, bool isSubclass);

} 
