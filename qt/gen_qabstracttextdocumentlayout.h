#pragma once
#ifndef MIQT_QT_GEN_QABSTRACTTEXTDOCUMENTLAYOUT_H
#define MIQT_QT_GEN_QABSTRACTTEXTDOCUMENTLAYOUT_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QAbstractTextDocumentLayout__PaintContext)
typedef QAbstractTextDocumentLayout::PaintContext QAbstractTextDocumentLayout__PaintContext;
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QAbstractTextDocumentLayout__Selection)
typedef QAbstractTextDocumentLayout::Selection QAbstractTextDocumentLayout__Selection;
typedef struct QAbstractTextDocumentLayout QAbstractTextDocumentLayout;
typedef struct QAbstractTextDocumentLayout__PaintContext QAbstractTextDocumentLayout__PaintContext;
typedef struct QAbstractTextDocumentLayout__Selection QAbstractTextDocumentLayout__Selection;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QPaintDevice QPaintDevice;
typedef struct QPainter QPainter;
typedef struct QPointF QPointF;
typedef struct QRectF QRectF;
typedef struct QSizeF QSizeF;
typedef struct QTextBlock QTextBlock;
typedef struct QTextDocument QTextDocument;
typedef struct QTextFormat QTextFormat;
typedef struct QTextFrame QTextFrame;
typedef struct QTextInlineObject QTextInlineObject;
typedef struct QTextObjectInterface QTextObjectInterface;

extern __declspec(dllexport) 
void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) 
void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractTextDocumentLayout* QAbstractTextDocumentLayout_new(QTextDocument* doc);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_virtbase(QAbstractTextDocumentLayout* src
, QObject** outptr_QObject
);
extern __declspec(dllexport) 
QMetaObject* QAbstractTextDocumentLayout_MetaObject(const QAbstractTextDocumentLayout* self);
extern __declspec(dllexport) 
void* QAbstractTextDocumentLayout_Metacast(QAbstractTextDocumentLayout* self, const char* param1);
extern __declspec(dllexport) 
struct miqt_string QAbstractTextDocumentLayout_Tr(const char* s);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_Draw(QAbstractTextDocumentLayout* self, QPainter* painter, const PaintContext* context);
extern __declspec(dllexport) 
int QAbstractTextDocumentLayout_HitTest(const QAbstractTextDocumentLayout* self, QPointF* point, int accuracy);
extern __declspec(dllexport) 
struct miqt_string QAbstractTextDocumentLayout_AnchorAt(const QAbstractTextDocumentLayout* self, QPointF* pos);
extern __declspec(dllexport) 
struct miqt_string QAbstractTextDocumentLayout_ImageAt(const QAbstractTextDocumentLayout* self, QPointF* pos);
extern __declspec(dllexport) 
QTextFormat* QAbstractTextDocumentLayout_FormatAt(const QAbstractTextDocumentLayout* self, QPointF* pos);
extern __declspec(dllexport) 
QTextBlock* QAbstractTextDocumentLayout_BlockWithMarkerAt(const QAbstractTextDocumentLayout* self, QPointF* pos);
extern __declspec(dllexport) 
int QAbstractTextDocumentLayout_PageCount(const QAbstractTextDocumentLayout* self);
extern __declspec(dllexport) 
QSizeF* QAbstractTextDocumentLayout_DocumentSize(const QAbstractTextDocumentLayout* self);
extern __declspec(dllexport) 
QRectF* QAbstractTextDocumentLayout_FrameBoundingRect(const QAbstractTextDocumentLayout* self, QTextFrame* frame);
extern __declspec(dllexport) 
QRectF* QAbstractTextDocumentLayout_BlockBoundingRect(const QAbstractTextDocumentLayout* self, QTextBlock* block);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_SetPaintDevice(QAbstractTextDocumentLayout* self, QPaintDevice* device);
extern __declspec(dllexport) 
QPaintDevice* QAbstractTextDocumentLayout_PaintDevice(const QAbstractTextDocumentLayout* self);
extern __declspec(dllexport) 
QTextDocument* QAbstractTextDocumentLayout_Document(const QAbstractTextDocumentLayout* self);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_RegisterHandler(QAbstractTextDocumentLayout* self, int objectType, QObject* component);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_UnregisterHandler(QAbstractTextDocumentLayout* self, int objectType);
extern __declspec(dllexport) 
QTextObjectInterface* QAbstractTextDocumentLayout_HandlerForObject(const QAbstractTextDocumentLayout* self, int objectType);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_Update(QAbstractTextDocumentLayout* self);
void QAbstractTextDocumentLayout_connect_Update(QAbstractTextDocumentLayout* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_UpdateBlock(QAbstractTextDocumentLayout* self, QTextBlock* block);
void QAbstractTextDocumentLayout_connect_UpdateBlock(QAbstractTextDocumentLayout* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_DocumentSizeChanged(QAbstractTextDocumentLayout* self, QSizeF* newSize);
void QAbstractTextDocumentLayout_connect_DocumentSizeChanged(QAbstractTextDocumentLayout* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_PageCountChanged(QAbstractTextDocumentLayout* self, int newPages);
void QAbstractTextDocumentLayout_connect_PageCountChanged(QAbstractTextDocumentLayout* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_DocumentChanged(QAbstractTextDocumentLayout* self, int from, int charsRemoved, int charsAdded);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_ResizeInlineObject(QAbstractTextDocumentLayout* self, QTextInlineObject* item, int posInDocument, QTextFormat* format);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_PositionInlineObject(QAbstractTextDocumentLayout* self, QTextInlineObject* item, int posInDocument, QTextFormat* format);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_DrawInlineObject(QAbstractTextDocumentLayout* self, QPainter* painter, QRectF* rect, QTextInlineObject* object, int posInDocument, QTextFormat* format);
extern __declspec(dllexport) 
struct miqt_string QAbstractTextDocumentLayout_Tr2(const char* s, const char* c);
extern __declspec(dllexport) 
struct miqt_string QAbstractTextDocumentLayout_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_UnregisterHandler2(QAbstractTextDocumentLayout* self, int objectType, QObject* component);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_Update1(QAbstractTextDocumentLayout* self, QRectF* param1);
void QAbstractTextDocumentLayout_connect_Update1(QAbstractTextDocumentLayout* self, intptr_t slot);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_override_virtual_MetaObject(void* self, intptr_t slot);
QMetaObject* QAbstractTextDocumentLayout_virtualbase_MetaObject(const void* self);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_override_virtual_Metacast(void* self, intptr_t slot);
void* QAbstractTextDocumentLayout_virtualbase_Metacast(void* self, const char* param1);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout_Delete(QAbstractTextDocumentLayout* self, bool isSubclass);

extern __declspec(dllexport) 
QSizeF* QTextObjectInterface_IntrinsicSize(QTextObjectInterface* self, QTextDocument* doc, int posInDocument, QTextFormat* format);
extern __declspec(dllexport) 
void QTextObjectInterface_DrawObject(QTextObjectInterface* self, QPainter* painter, QRectF* rect, QTextDocument* doc, int posInDocument, QTextFormat* format);
extern __declspec(dllexport) 
void QTextObjectInterface_OperatorAssign(QTextObjectInterface* self, QTextObjectInterface* param1);
extern __declspec(dllexport) 
void QTextObjectInterface_Delete(QTextObjectInterface* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractTextDocumentLayout__Selection* QAbstractTextDocumentLayout__Selection_new();
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout__Selection_OperatorAssign(QAbstractTextDocumentLayout__Selection* self, const Selection* param1);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout__Selection_Delete(QAbstractTextDocumentLayout__Selection* self, bool isSubclass);

extern __declspec(dllexport) 
QAbstractTextDocumentLayout__PaintContext* QAbstractTextDocumentLayout__PaintContext_new();
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout__PaintContext_OperatorAssign(QAbstractTextDocumentLayout__PaintContext* self, const PaintContext* param1);
extern __declspec(dllexport) 
void QAbstractTextDocumentLayout__PaintContext_Delete(QAbstractTextDocumentLayout__PaintContext* self, bool isSubclass);

}
