#pragma once
#ifndef MIQT_QT_GEN_QTEXTCURSOR_H
#define MIQT_QT_GEN_QTEXTCURSOR_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QImage QImage;
typedef struct QTextBlock QTextBlock;
typedef struct QTextBlockFormat QTextBlockFormat;
typedef struct QTextCharFormat QTextCharFormat;
typedef struct QTextCursor QTextCursor;
typedef struct QTextDocument QTextDocument;
typedef struct QTextDocumentFragment QTextDocumentFragment;
typedef struct QTextFrame QTextFrame;
typedef struct QTextFrameFormat QTextFrameFormat;
typedef struct QTextImageFormat QTextImageFormat;
typedef struct QTextList QTextList;
typedef struct QTextListFormat QTextListFormat;
typedef struct QTextTable QTextTable;
typedef struct QTextTableFormat QTextTableFormat;
typedef struct _GUID _GUID;
typedef struct type_info type_info;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTextCursor* QTextCursor_new();
extern __declspec(dllexport) QTextCursor* QTextCursor_new2(QTextDocument* document);
extern __declspec(dllexport) QTextCursor* QTextCursor_new3(QTextFrame* frame);
extern __declspec(dllexport) QTextCursor* QTextCursor_new4(QTextBlock* block);
extern __declspec(dllexport) QTextCursor* QTextCursor_new5(QTextCursor* cursor);
extern __declspec(dllexport) void QTextCursor_OperatorAssign(QTextCursor* self, QTextCursor* other);
extern __declspec(dllexport) void QTextCursor_Swap(QTextCursor* self, QTextCursor* other);
extern __declspec(dllexport) bool QTextCursor_IsNull(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_SetPosition(QTextCursor* self, int pos);
extern __declspec(dllexport) int QTextCursor_Position(const QTextCursor* self);
extern __declspec(dllexport) int QTextCursor_PositionInBlock(const QTextCursor* self);
extern __declspec(dllexport) int QTextCursor_Anchor(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_InsertText(QTextCursor* self, struct miqt_string text);
extern __declspec(dllexport) void QTextCursor_InsertText2(QTextCursor* self, struct miqt_string text, QTextCharFormat* format);
extern __declspec(dllexport) bool QTextCursor_MovePosition(QTextCursor* self, MoveOperation op);
extern __declspec(dllexport) bool QTextCursor_VisualNavigation(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_SetVisualNavigation(QTextCursor* self, bool b);
extern __declspec(dllexport) void QTextCursor_SetVerticalMovementX(QTextCursor* self, int x);
extern __declspec(dllexport) int QTextCursor_VerticalMovementX(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_SetKeepPositionOnInsert(QTextCursor* self, bool b);
extern __declspec(dllexport) bool QTextCursor_KeepPositionOnInsert(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_DeleteChar(QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_DeletePreviousChar(QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_Select(QTextCursor* self, SelectionType selection);
extern __declspec(dllexport) bool QTextCursor_HasSelection(const QTextCursor* self);
extern __declspec(dllexport) bool QTextCursor_HasComplexSelection(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_RemoveSelectedText(QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_ClearSelection(QTextCursor* self);
extern __declspec(dllexport) int QTextCursor_SelectionStart(const QTextCursor* self);
extern __declspec(dllexport) int QTextCursor_SelectionEnd(const QTextCursor* self);
extern __declspec(dllexport) struct miqt_string QTextCursor_SelectedText(const QTextCursor* self);
extern __declspec(dllexport) QTextDocumentFragment* QTextCursor_Selection(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_SelectedTableCells(const QTextCursor* self, int* firstRow, int* numRows, int* firstColumn, int* numColumns);
extern __declspec(dllexport) QTextBlock* QTextCursor_Block(const QTextCursor* self);
extern __declspec(dllexport) QTextCharFormat* QTextCursor_CharFormat(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_SetCharFormat(QTextCursor* self, QTextCharFormat* format);
extern __declspec(dllexport) void QTextCursor_MergeCharFormat(QTextCursor* self, QTextCharFormat* modifier);
extern __declspec(dllexport) QTextBlockFormat* QTextCursor_BlockFormat(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_SetBlockFormat(QTextCursor* self, QTextBlockFormat* format);
extern __declspec(dllexport) void QTextCursor_MergeBlockFormat(QTextCursor* self, QTextBlockFormat* modifier);
extern __declspec(dllexport) QTextCharFormat* QTextCursor_BlockCharFormat(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_SetBlockCharFormat(QTextCursor* self, QTextCharFormat* format);
extern __declspec(dllexport) void QTextCursor_MergeBlockCharFormat(QTextCursor* self, QTextCharFormat* modifier);
extern __declspec(dllexport) bool QTextCursor_AtBlockStart(const QTextCursor* self);
extern __declspec(dllexport) bool QTextCursor_AtBlockEnd(const QTextCursor* self);
extern __declspec(dllexport) bool QTextCursor_AtStart(const QTextCursor* self);
extern __declspec(dllexport) bool QTextCursor_AtEnd(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_InsertBlock(QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_InsertBlockWithFormat(QTextCursor* self, QTextBlockFormat* format);
extern __declspec(dllexport) void QTextCursor_InsertBlock2(QTextCursor* self, QTextBlockFormat* format, QTextCharFormat* charFormat);
extern __declspec(dllexport) QTextList* QTextCursor_InsertList(QTextCursor* self, QTextListFormat* format);
extern __declspec(dllexport) QTextList* QTextCursor_InsertListWithStyle(QTextCursor* self, int style);
extern __declspec(dllexport) QTextList* QTextCursor_CreateList(QTextCursor* self, QTextListFormat* format);
extern __declspec(dllexport) QTextList* QTextCursor_CreateListWithStyle(QTextCursor* self, int style);
extern __declspec(dllexport) QTextList* QTextCursor_CurrentList(const QTextCursor* self);
extern __declspec(dllexport) QTextTable* QTextCursor_InsertTable(QTextCursor* self, int rows, int cols, QTextTableFormat* format);
extern __declspec(dllexport) QTextTable* QTextCursor_InsertTable2(QTextCursor* self, int rows, int cols);
extern __declspec(dllexport) QTextTable* QTextCursor_CurrentTable(const QTextCursor* self);
extern __declspec(dllexport) QTextFrame* QTextCursor_InsertFrame(QTextCursor* self, QTextFrameFormat* format);
extern __declspec(dllexport) QTextFrame* QTextCursor_CurrentFrame(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_InsertFragment(QTextCursor* self, QTextDocumentFragment* fragment);
extern __declspec(dllexport) void QTextCursor_InsertHtml(QTextCursor* self, struct miqt_string html);
extern __declspec(dllexport) void QTextCursor_InsertMarkdown(QTextCursor* self, struct miqt_string markdown);
extern __declspec(dllexport) void QTextCursor_InsertImage(QTextCursor* self, QTextImageFormat* format, int alignment);
extern __declspec(dllexport) void QTextCursor_InsertImageWithFormat(QTextCursor* self, QTextImageFormat* format);
extern __declspec(dllexport) void QTextCursor_InsertImageWithName(QTextCursor* self, struct miqt_string name);
extern __declspec(dllexport) void QTextCursor_InsertImageWithImage(QTextCursor* self, QImage* image);
extern __declspec(dllexport) void QTextCursor_BeginEditBlock(QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_JoinPreviousEditBlock(QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_EndEditBlock(QTextCursor* self);
extern __declspec(dllexport) bool QTextCursor_OperatorNotEqual(const QTextCursor* self, QTextCursor* rhs);
extern __declspec(dllexport) bool QTextCursor_OperatorLesser(const QTextCursor* self, QTextCursor* rhs);
extern __declspec(dllexport) bool QTextCursor_OperatorLesserOrEqual(const QTextCursor* self, QTextCursor* rhs);
extern __declspec(dllexport) bool QTextCursor_OperatorEqual(const QTextCursor* self, QTextCursor* rhs);
extern __declspec(dllexport) bool QTextCursor_OperatorGreaterOrEqual(const QTextCursor* self, QTextCursor* rhs);
extern __declspec(dllexport) bool QTextCursor_OperatorGreater(const QTextCursor* self, QTextCursor* rhs);
extern __declspec(dllexport) bool QTextCursor_IsCopyOf(const QTextCursor* self, QTextCursor* other);
extern __declspec(dllexport) int QTextCursor_BlockNumber(const QTextCursor* self);
extern __declspec(dllexport) int QTextCursor_ColumnNumber(const QTextCursor* self);
extern __declspec(dllexport) QTextDocument* QTextCursor_Document(const QTextCursor* self);
extern __declspec(dllexport) void QTextCursor_SetPosition2(QTextCursor* self, int pos, MoveMode mode);
extern __declspec(dllexport) bool QTextCursor_MovePosition2(QTextCursor* self, MoveOperation op, MoveMode param2);
extern __declspec(dllexport) bool QTextCursor_MovePosition3(QTextCursor* self, MoveOperation op, MoveMode param2, int n);
extern __declspec(dllexport) void QTextCursor_InsertMarkdown2(QTextCursor* self, struct miqt_string markdown, int features);
extern __declspec(dllexport) void QTextCursor_InsertImage2(QTextCursor* self, QImage* image, struct miqt_string name);
extern __declspec(dllexport) void QTextCursor_Delete(QTextCursor* self, bool isSubclass);

} 
