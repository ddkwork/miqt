#pragma once
#ifndef MIQT_QT_GEN_QTEXTDOCUMENT_H
#define MIQT_QT_GEN_QTEXTDOCUMENT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QAbstractTextDocumentLayout QAbstractTextDocumentLayout;
typedef struct QAbstractUndoItem QAbstractUndoItem;
typedef struct QChar QChar;
typedef struct QFont QFont;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPagedPaintDevice QPagedPaintDevice;
typedef struct QPainter QPainter;
typedef struct QRectF QRectF;
typedef struct QRegularExpression QRegularExpression;
typedef struct QSizeF QSizeF;
typedef struct QTextBlock QTextBlock;
typedef struct QTextCursor QTextCursor;
typedef struct QTextDocument QTextDocument;
typedef struct QTextFormat QTextFormat;
typedef struct QTextFrame QTextFrame;
typedef struct QTextObject QTextObject;
typedef struct QTextOption QTextOption;
typedef struct QUrl QUrl;
typedef struct QVariant QVariant;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
void QAbstractUndoItem_Undo(QAbstractUndoItem* self);
extern __declspec(dllexport) 
void QAbstractUndoItem_Redo(QAbstractUndoItem* self);
extern __declspec(dllexport) 
void QAbstractUndoItem_OperatorAssign(QAbstractUndoItem* self, QAbstractUndoItem* param1);
extern __declspec(dllexport) 
void QAbstractUndoItem_Delete(QAbstractUndoItem* self, bool isSubclass);

extern __declspec(dllexport) 
QTextDocument* QTextDocument_new();
extern __declspec(dllexport) 
QTextDocument* QTextDocument_new2(struct miqt_string text);
extern __declspec(dllexport) 
QTextDocument* QTextDocument_new3(QObject* parent);
extern __declspec(dllexport) 
QTextDocument* QTextDocument_new4(struct miqt_string text, QObject* parent);
extern __declspec(dllexport) 
void QTextDocument_virtbase(QTextDocument* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QTextDocument_MetaObject(const QTextDocument* self);
extern __declspec(dllexport) 
void* QTextDocument_Metacast(QTextDocument* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QTextDocument_Tr(const char* s);
extern __declspec(dllexport) 
QTextDocument* QTextDocument_Clone(const QTextDocument* self);
extern __declspec(dllexport) 
bool QTextDocument_IsEmpty(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_Clear(QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetUndoRedoEnabled(QTextDocument* self, bool enable);
extern __declspec(dllexport) 
bool QTextDocument_IsUndoRedoEnabled(const QTextDocument* self);
extern __declspec(dllexport) 
bool QTextDocument_IsUndoAvailable(const QTextDocument* self);
extern __declspec(dllexport) 
bool QTextDocument_IsRedoAvailable(const QTextDocument* self);
extern __declspec(dllexport) 
int QTextDocument_AvailableUndoSteps(const QTextDocument* self);
extern __declspec(dllexport) 
int QTextDocument_AvailableRedoSteps(const QTextDocument* self);
extern __declspec(dllexport) 
int QTextDocument_Revision(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetDocumentLayout(QTextDocument* self, QAbstractTextDocumentLayout* layout);
extern __declspec(dllexport) 
QAbstractTextDocumentLayout* QTextDocument_DocumentLayout(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetMetaInformation(QTextDocument* self, MetaInformation info, struct miqt_string param2);
extern __declspec(dllexport) 
struct miqt_string QTextDocument_MetaInformation(const QTextDocument* self, MetaInformation info);
extern __declspec(dllexport) 
struct miqt_string QTextDocument_ToHtml(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetHtml(QTextDocument* self, struct miqt_string html);
extern __declspec(dllexport) 
struct miqt_string QTextDocument_ToMarkdown(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetMarkdown(QTextDocument* self, struct miqt_string markdown);
extern __declspec(dllexport) 
struct miqt_string QTextDocument_ToRawText(const QTextDocument* self);
extern __declspec(dllexport) 
struct miqt_string QTextDocument_ToPlainText(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetPlainText(QTextDocument* self, struct miqt_string text);
extern __declspec(dllexport) 
QChar* QTextDocument_CharacterAt(const QTextDocument* self, int pos);
extern __declspec(dllexport) 
QTextCursor* QTextDocument_Find(const QTextDocument* self, struct miqt_string subString);
extern __declspec(dllexport) 
QTextCursor* QTextDocument_Find2(const QTextDocument* self, struct miqt_string subString, QTextCursor* cursor);
extern __declspec(dllexport) 
QTextCursor* QTextDocument_FindWithExpr(const QTextDocument* self, QRegularExpression* expr);
extern __declspec(dllexport) 
QTextCursor* QTextDocument_Find3(const QTextDocument* self, QRegularExpression* expr, QTextCursor* cursor);
extern __declspec(dllexport) 
QTextFrame* QTextDocument_FrameAt(const QTextDocument* self, int pos);
extern __declspec(dllexport) 
QTextFrame* QTextDocument_RootFrame(const QTextDocument* self);
extern __declspec(dllexport) 
QTextObject* QTextDocument_Object(const QTextDocument* self, int objectIndex);
extern __declspec(dllexport) 
QTextObject* QTextDocument_ObjectForFormat(const QTextDocument* self, QTextFormat* param1);
extern __declspec(dllexport) 
QTextBlock* QTextDocument_FindBlock(const QTextDocument* self, int pos);
extern __declspec(dllexport) 
QTextBlock* QTextDocument_FindBlockByNumber(const QTextDocument* self, int blockNumber);
extern __declspec(dllexport) 
QTextBlock* QTextDocument_FindBlockByLineNumber(const QTextDocument* self, int blockNumber);
extern __declspec(dllexport) 
QTextBlock* QTextDocument_Begin(const QTextDocument* self);
extern __declspec(dllexport) 
QTextBlock* QTextDocument_End(const QTextDocument* self);
extern __declspec(dllexport) 
QTextBlock* QTextDocument_FirstBlock(const QTextDocument* self);
extern __declspec(dllexport) 
QTextBlock* QTextDocument_LastBlock(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetPageSize(QTextDocument* self, QSizeF* size);
extern __declspec(dllexport) 
QSizeF* QTextDocument_PageSize(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetDefaultFont(QTextDocument* self, QFont* font);
extern __declspec(dllexport) 
QFont* QTextDocument_DefaultFont(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetSuperScriptBaseline(QTextDocument* self, double baseline);
extern __declspec(dllexport) 
double QTextDocument_SuperScriptBaseline(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetSubScriptBaseline(QTextDocument* self, double baseline);
extern __declspec(dllexport) 
double QTextDocument_SubScriptBaseline(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetBaselineOffset(QTextDocument* self, double baseline);
extern __declspec(dllexport) 
double QTextDocument_BaselineOffset(const QTextDocument* self);
extern __declspec(dllexport) 
int QTextDocument_PageCount(const QTextDocument* self);
extern __declspec(dllexport) 
bool QTextDocument_IsModified(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_Print(const QTextDocument* self, QPagedPaintDevice* printer);
extern __declspec(dllexport) 
QVariant* QTextDocument_Resource(const QTextDocument* self, int typeVal, QUrl* name);
extern __declspec(dllexport) 
void QTextDocument_AddResource(QTextDocument* self, int typeVal, QUrl* name, QVariant* resource);
extern __declspec(dllexport) 
void QTextDocument_SetResourceProvider(QTextDocument* self, const ResourceProvider* provider);
extern __declspec(dllexport) 
void QTextDocument_SetDefaultResourceProvider(const ResourceProvider* provider);
extern __declspec(dllexport) 
struct miqt_array /* of QTextFormat* */  QTextDocument_AllFormats(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_MarkContentsDirty(QTextDocument* self, int from, int length);
extern __declspec(dllexport) 
void QTextDocument_SetUseDesignMetrics(QTextDocument* self, bool b);
extern __declspec(dllexport) 
bool QTextDocument_UseDesignMetrics(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetLayoutEnabled(QTextDocument* self, bool b);
extern __declspec(dllexport) 
bool QTextDocument_IsLayoutEnabled(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_DrawContents(QTextDocument* self, QPainter* painter);
extern __declspec(dllexport) 
void QTextDocument_SetTextWidth(QTextDocument* self, double width);
extern __declspec(dllexport) 
double QTextDocument_TextWidth(const QTextDocument* self);
extern __declspec(dllexport) 
double QTextDocument_IdealWidth(const QTextDocument* self);
extern __declspec(dllexport) 
double QTextDocument_IndentWidth(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetIndentWidth(QTextDocument* self, double width);
extern __declspec(dllexport) 
double QTextDocument_DocumentMargin(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetDocumentMargin(QTextDocument* self, double margin);
extern __declspec(dllexport) 
void QTextDocument_AdjustSize(QTextDocument* self);
extern __declspec(dllexport) 
QSizeF* QTextDocument_Size(const QTextDocument* self);
extern __declspec(dllexport) 
int QTextDocument_BlockCount(const QTextDocument* self);
extern __declspec(dllexport) 
int QTextDocument_LineCount(const QTextDocument* self);
extern __declspec(dllexport) 
int QTextDocument_CharacterCount(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetDefaultStyleSheet(QTextDocument* self, struct miqt_string sheet);
extern __declspec(dllexport) 
struct miqt_string QTextDocument_DefaultStyleSheet(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_Undo(QTextDocument* self, QTextCursor* cursor);
extern __declspec(dllexport) 
void QTextDocument_Redo(QTextDocument* self, QTextCursor* cursor);
extern __declspec(dllexport) 
void QTextDocument_ClearUndoRedoStacks(QTextDocument* self);
extern __declspec(dllexport) 
int QTextDocument_MaximumBlockCount(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetMaximumBlockCount(QTextDocument* self, int maximum);
extern __declspec(dllexport) 
QTextOption* QTextDocument_DefaultTextOption(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetDefaultTextOption(QTextDocument* self, QTextOption* option);
extern __declspec(dllexport) 
QUrl* QTextDocument_BaseUrl(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetBaseUrl(QTextDocument* self, QUrl* url);
extern __declspec(dllexport) 
int QTextDocument_DefaultCursorMoveStyle(const QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_SetDefaultCursorMoveStyle(QTextDocument* self, int style);
extern __declspec(dllexport) 
void QTextDocument_ContentsChange(QTextDocument* self, int from, int charsRemoved, int charsAdded);
void QTextDocument_connect_ContentsChange(QTextDocument* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextDocument_ContentsChanged(QTextDocument* self);
void QTextDocument_connect_ContentsChanged(QTextDocument* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextDocument_UndoAvailable(QTextDocument* self, bool param1);
void QTextDocument_connect_UndoAvailable(QTextDocument* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextDocument_RedoAvailable(QTextDocument* self, bool param1);
void QTextDocument_connect_RedoAvailable(QTextDocument* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextDocument_UndoCommandAdded(QTextDocument* self);
void QTextDocument_connect_UndoCommandAdded(QTextDocument* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextDocument_ModificationChanged(QTextDocument* self, bool m);
void QTextDocument_connect_ModificationChanged(QTextDocument* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextDocument_CursorPositionChanged(QTextDocument* self, QTextCursor* cursor);
void QTextDocument_connect_CursorPositionChanged(QTextDocument* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextDocument_BlockCountChanged(QTextDocument* self, int newBlockCount);
void QTextDocument_connect_BlockCountChanged(QTextDocument* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextDocument_BaseUrlChanged(QTextDocument* self, QUrl* url);
void QTextDocument_connect_BaseUrlChanged(QTextDocument* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextDocument_DocumentLayoutChanged(QTextDocument* self);
void QTextDocument_connect_DocumentLayoutChanged(QTextDocument* self, intptr_t slot);
extern __declspec(dllexport) 
void QTextDocument_Undo2(QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_Redo2(QTextDocument* self);
extern __declspec(dllexport) 
void QTextDocument_AppendUndoItem(QTextDocument* self, QAbstractUndoItem* param1);
extern __declspec(dllexport) 
void QTextDocument_SetModified(QTextDocument* self);
extern __declspec(dllexport) 
QTextObject* QTextDocument_CreateObject(QTextDocument* self, QTextFormat* f);
extern __declspec(dllexport) 
QVariant* QTextDocument_LoadResource(QTextDocument* self, int typeVal, QUrl* name);
extern __declspec(dllexport) 
struct miqt_string QTextDocument_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QTextDocument_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
QTextDocument* QTextDocument_Clone1(const QTextDocument* self, QObject* parent);
extern __declspec(dllexport) 
struct miqt_string QTextDocument_ToMarkdown1(const QTextDocument* self, MarkdownFeatures features);
extern __declspec(dllexport) 
void QTextDocument_SetMarkdown2(QTextDocument* self, struct miqt_string markdown, MarkdownFeatures features);
extern __declspec(dllexport) 
QTextCursor* QTextDocument_Find22(const QTextDocument* self, struct miqt_string subString, int from);
extern __declspec(dllexport) 
QTextCursor* QTextDocument_Find32(const QTextDocument* self, struct miqt_string subString, int from, FindFlags options);
extern __declspec(dllexport) 
QTextCursor* QTextDocument_Find33(const QTextDocument* self, struct miqt_string subString, QTextCursor* cursor, FindFlags options);
extern __declspec(dllexport) 
QTextCursor* QTextDocument_Find23(const QTextDocument* self, QRegularExpression* expr, int from);
extern __declspec(dllexport) 
QTextCursor* QTextDocument_Find34(const QTextDocument* self, QRegularExpression* expr, int from, FindFlags options);
extern __declspec(dllexport) 
QTextCursor* QTextDocument_Find35(const QTextDocument* self, QRegularExpression* expr, QTextCursor* cursor, FindFlags options);
extern __declspec(dllexport) 
void QTextDocument_DrawContents2(QTextDocument* self, QPainter* painter, QRectF* rect);
extern __declspec(dllexport) 
void QTextDocument_ClearUndoRedoStacks1(QTextDocument* self, Stacks historyToClear);
extern __declspec(dllexport) 
void QTextDocument_SetModified1(QTextDocument* self, bool m);
extern __declspec(dllexport) 
void QTextDocument_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QTextDocument_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QTextDocument_override_virtual_Metacast(void* self, intptr_t slot);
void* QTextDocument_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QTextDocument_Delete(QTextDocument* self, bool isSubclass);

}
