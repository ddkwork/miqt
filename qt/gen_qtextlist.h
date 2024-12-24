#pragma once
#ifndef MIQT_QT_GEN_QTEXTLIST_H
#define MIQT_QT_GEN_QTEXTLIST_H

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

#include "../libmiqt/libmiqt.h"
extern "C" {
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QTextBlock QTextBlock;
typedef struct QTextBlockGroup QTextBlockGroup;
typedef struct QTextDocument QTextDocument;
typedef struct QTextList QTextList;
typedef struct QTextListFormat QTextListFormat;
typedef struct QTextObject QTextObject;

extern __declspec(dllexport) void _GUID_Delete(_GUID* self, bool isSubclass);

extern __declspec(dllexport) void type_info_Delete(type_info* self, bool isSubclass);

extern __declspec(dllexport) QTextList* QTextList_new(QTextDocument* doc);
extern __declspec(dllexport) void QTextList_virtbase(QTextList* src, QTextBlockGroup** outptr_QTextBlockGroup);
extern __declspec(dllexport) QMetaObject* QTextList_MetaObject(const QTextList* self);
extern __declspec(dllexport) void* QTextList_Metacast(QTextList* self, const char* param1);
extern __declspec(dllexport) struct miqt_string QTextList_Tr(const char* s);
extern __declspec(dllexport) int QTextList_Count(const QTextList* self);
extern __declspec(dllexport) QTextBlock* QTextList_Item(const QTextList* self, int i);
extern __declspec(dllexport) int QTextList_ItemNumber(const QTextList* self, QTextBlock* param1);
extern __declspec(dllexport) struct miqt_string QTextList_ItemText(const QTextList* self, QTextBlock* param1);
extern __declspec(dllexport) void QTextList_RemoveItem(QTextList* self, int i);
extern __declspec(dllexport) void QTextList_Remove(QTextList* self, QTextBlock* param1);
extern __declspec(dllexport) void QTextList_Add(QTextList* self, QTextBlock* block);
extern __declspec(dllexport) void QTextList_SetFormat(QTextList* self, QTextListFormat* format);
extern __declspec(dllexport) QTextListFormat* QTextList_Format(const QTextList* self);
extern __declspec(dllexport) struct miqt_string QTextList_Tr2(const char* s, const char* c);
extern __declspec(dllexport) struct miqt_string QTextList_Tr3(const char* s, const char* c, int n);
extern __declspec(dllexport) void QTextList_override_virtual_BlockInserted(void* self, intptr_t slot);
void QTextList_virtualbase_BlockInserted(void* self, QTextBlock* block);
extern __declspec(dllexport) void QTextList_override_virtual_BlockRemoved(void* self, intptr_t slot);
void QTextList_virtualbase_BlockRemoved(void* self, QTextBlock* block);
extern __declspec(dllexport) void QTextList_override_virtual_BlockFormatChanged(void* self, intptr_t slot);
void QTextList_virtualbase_BlockFormatChanged(void* self, QTextBlock* block);
extern __declspec(dllexport) void QTextList_Delete(QTextList* self, bool isSubclass);

} 
